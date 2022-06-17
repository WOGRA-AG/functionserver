package executor

import (
	"database/sql"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/go-sql-driver/mysql"
	"github.com/robertkrimen/otto"
	"wogra.com/config"
)

var (
	functionName string = "executorFunctionWrapper"
	errHalt             = errors.New("function timeout exception")
)

func executeJSFunction(botId string, functionCode string, functionParams map[string]string) (string, error) {

	start := time.Now()
	vm := otto.New()

	addEnhancedFeaturesToScriptEngine(botId, vm)

	for keyParam, valueParam := range functionParams {
		vm.Set(keyParam, valueParam)
	}

	defer runtimeCheck(start)
	addTimeoutCheckToEngine(vm)
	//log.Printf("Code: %q", functionCode)

	result, err := vm.Run(functionCode)

	if err != nil {
		log.Print(err)
	}

	return result.String(), err
}

func addTimeoutCheckToEngine(engine *otto.Otto) {
	engine.Interrupt = make(chan func(), 1) // The buffer prevents blocking

	go func() {
		time.Sleep(10 * time.Second) // Stop after ten seconds
		engine.Interrupt <- func() {
			panic(errHalt)
		}
	}()
}

func runtimeCheck(start time.Time) {
	duration := time.Since(start)
	halt := recover()
	if halt != nil {
		if halt == errHalt {
			log.Printf("Some code took to long! Stopping after: %v\n", duration)
			return
		}

		panic(halt) // Something else happened, repanic!
	}
}

func wrapJsFunction(functionCode string) (string, error) {

	if len(functionCode) == 0 {
		return "", errors.New("length of code is zero")
	}

	functionCode = fmt.Sprintf("function %s()\n{\n%s\n}\n%s();", functionName, functionCode, functionName)
	return functionCode, nil
}

func ExecuteFunction(botId string, functionCode string, functionParams map[string]string) (string, error) {

	newFunctionCode, err := wrapJsFunction(functionCode)

	if err == nil {
		return executeJSFunction(botId, newFunctionCode, functionParams)
	} else {
		return "", err
	}
}

func addEnhancedFeaturesToScriptEngine(botId string, engine *otto.Otto) {
	if engine != nil {
		addHttpFeaturesToScriptEngine(engine)
		addDocumentFeaturesToScriptEngine(botId, engine)
	}
}

func addDocumentFeaturesToScriptEngine(botId string, engine *otto.Otto) {
	if engine != nil {
		addDocumentLoadFeatureToScriptEngine(botId, engine)
		addDocumentSaveFeatureToScriptEngine(botId, engine)
		addDocumentDeleteFeatureToScriptEngine(botId, engine)
	}
}

func addDocumentLoadFeatureToScriptEngine(botId string, engine *otto.Otto) {
	if engine != nil {
		engine.Set("loadDocument", func(call otto.FunctionCall) otto.Value {
			content := loadDocument(botId, call.Argument(0).String())
			result, _ := engine.ToValue(content)
			return result
		})
	}
}

func addDocumentSaveFeatureToScriptEngine(botId string, engine *otto.Otto) {
	if engine != nil {
		engine.Set("saveDocument", func(call otto.FunctionCall) otto.Value {

			success := saveDocument(botId, call.Argument(0).String(), call.Argument(1).String())
			result, _ := engine.ToValue(success)
			return result
		})
	}
}

func addDocumentDeleteFeatureToScriptEngine(botId string, engine *otto.Otto) {
	if engine != nil {
		engine.Set("deleteDocument", func(call otto.FunctionCall) otto.Value {

			success := deleteDocument(botId, call.Argument(0).String())
			result, _ := engine.ToValue(success)
			return result
		})
	}
}

func createDbObject() *sql.DB {
	dbConfig := config.ReadDatabaseConfiguration()

	cfg := mysql.Config{
		User:                 dbConfig.Login,
		Passwd:               dbConfig.Password,
		Net:                  "tcp",
		Addr:                 dbConfig.DatabaseUrl + ":" + dbConfig.DatabasePort,
		DBName:               dbConfig.DatabaseName,
		AllowNativePasswords: true,
	}

	var err error
	db, err := sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
		return nil
	} else {
		return db
	}

	return nil
}

func saveDocument(botId string, key string, value string) bool {

	db := createDbObject()

	if db != nil {
		rows, err := db.Query("SELECT count(*) FROM key_value_store WHERE botId = ? and identifier = ?", botId, key)

		defer rows.Close()

		if err != nil {
			log.Printf("check if key exists error. %v", err)
			return false
		} else {
			for rows.Next() {
				var count int
				if err := rows.Scan(&count); err != nil {
					log.Printf("check if key exists error. %v", err)
				} else {
					if count == 0 {
						// insert
						_, err = db.Query("insert into key_value_store (botid, identifier, value) VALUES (?, ?, ?)", botId, key, value)

						if err != nil {
							log.Printf("Insert error in key value store. %v", err)
						} else {
							return true
						}
					} else {
						// update
						_, err = db.Query("update key_value_store set value = ? where botId = ? and identifier = ?", value, botId, key)

						if err != nil {
							log.Printf("Update error in key value store. %v", err)
						} else {
							return true
						}
					}
				}
			}
		}
	} else {
		log.Print("Could not open DB connection for saving document")
	}

	return false
}

func loadDocument(botId string, key string) string {
	db := createDbObject()

	if db != nil {
		rows, err := db.Query("SELECT value FROM key_value_store WHERE botId = ? and identifier = ?", botId, key)

		defer rows.Close()

		if err != nil {
			log.Printf("loadKeyValue error. %v", err)
			return ""
		} else {
			for rows.Next() {
				var data string
				if err := rows.Scan(&data); err != nil {
					log.Printf("loadKeyValue error. %v", err)
				} else {
					return data
				}
			}
		}
	} else {
		log.Print("Could not open DB connection for loading document")
	}

	return ""
}

func deleteDocument(botId string, key string) bool {
	db := createDbObject()

	if db != nil {
		_, err := db.Exec("delete FROM key_value_store WHERE botId = ? and identifier = ?", botId, key)

		if err != nil {
			log.Printf("deleteDocument error. %v", err)
		} else {
			return true
		}
	} else {
		log.Print("Could not open DB connection for document deletion")
	}

	return false
}

func addHttpFeaturesToScriptEngine(engine *otto.Otto) {
	if engine != nil {
		addHttpGetFeatureToScriptEngine(engine)
		addHttpPostFeatureToScriptEngine(engine)
	}
}

func addHttpGetFeatureToScriptEngine(engine *otto.Otto) {
	if engine != nil {
		engine.Set("httpGet", func(call otto.FunctionCall) otto.Value {

			address := call.Argument(0).String()

			if validateRequestAddress(address) {
				response, err := http.Get(address)

				if err == nil {
					body, _ := io.ReadAll(response.Body)
					response.Body.Close()
					result, _ := engine.ToValue(string(body))
					return result
				} else {
					log.Printf("Request Error occured. Error: %q", err)
					result, _ := engine.ToValue("")
					return result
				}
			} else {
				log.Printf("Request address not allowed: %q", address)
				result, _ := engine.ToValue("not allowed")
				return result
			}
		})
	}
}

func addHttpPostFeatureToScriptEngine(engine *otto.Otto) {
	if engine != nil {
		engine.Set("httpPost", func(call otto.FunctionCall) otto.Value {

			address := call.Argument(0).String()

			if validateRequestAddress(address) {
				myReader := strings.NewReader(call.Argument(2).String())
				response, err := http.Post(address, call.Argument(1).String(), myReader)

				if err == nil {
					body, _ := io.ReadAll(response.Body)
					response.Body.Close()
					result, _ := engine.ToValue(string(body))
					return result
				} else {
					log.Printf("Request Error occured. Error: %q", err)
					result, _ := engine.ToValue("")
					return result
				}
			} else {
				log.Printf("Request address not allowed: %q", address)
				result, _ := engine.ToValue("not allowed")
				return result
			}
		})
	}
}

func validateRequestAddress(address string) bool {
	if strings.Contains(address, "://localhost") || strings.Contains(address, "://127.0.0.1") {
		return false
	}

	return true
}

func ExecuteWasmFunction(functionCode string, functionParams map[string]string) (string, error) {

	// todo
	return "", errors.New("not yet supported ;)")
}
