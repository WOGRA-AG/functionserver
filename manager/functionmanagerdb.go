package manager

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/go-sql-driver/mysql"
	"wogra.com/executor"
)

var db *sql.DB

type FunctionManagerDb struct {
	//db *sql.DB
}

func NewDb() *FunctionManagerDb {

	var fm = new(FunctionManagerDb)
	if fm.init(ReadDatabaseConfiguration()) {
		return fm
	}

	return nil
}

func (fm FunctionManagerDb) pingDb() bool {
	if db != nil {
		pingErr := db.Ping()
		if pingErr != nil {
			log.Fatal(pingErr)
			return false
		}
	} else {
		log.Print("DB is nil")
	}

	return true
}

func (fm FunctionManagerDb) init(dbConfig DbConfig) bool {

	if db == nil {
		cfg := mysql.Config{
			User:                 dbConfig.Login,
			Passwd:               dbConfig.Password,
			Net:                  "tcp",
			Addr:                 dbConfig.DatabaseUrl + ":" + dbConfig.DatabasePort,
			DBName:               dbConfig.DatabaseName,
			AllowNativePasswords: true,
		}

		var err error
		db, err = sql.Open("mysql", cfg.FormatDSN())
		if err != nil {
			log.Fatal(err)
			return false
		}
	}

	return fm.pingDb()
}

func (fm FunctionManagerDb) findFunction(name string, botId string) *FunctionDescription {

	fd := new(FunctionDescription)

	if db != nil {
		row := db.QueryRow("SELECT name, botid, code, version FROM functions WHERE name = ? and botid = ?", name, botId)
		err := row.Scan(&fd.Name, &fd.BotId, &fd.Code, &fd.Version)
		if err != nil {
			if err == sql.ErrNoRows {
				log.Printf("No function %q with BotId %q found.", name, botId)
				return nil
			}

			log.Printf("SQL Error for function search %q with BotId %q Error: %q", name, botId, err)
			return nil
		}
	} else {
		log.Print("DB is nil")
	}

	return fd
}

func (fm FunctionManagerDb) FindFunction(fd *FunctionDescription) *FunctionDescription {
	if fm.validateAppId(fd.AppId) {
		return fm.findFunction(fd.Name, fd.BotId)
	} else {
		return nil
	}

	return fd
}

func (fm FunctionManagerDb) AddFunction(fd *FunctionDescription) *FunctionDescription {

	if fm.validateAppId(fd.AppId) {
		if fm.FindFunction(fd) == nil {
			_, err := db.Exec("INSERT INTO functions (name, botid, code, version) VALUES (?, ?, ?, 0)", fd.Name, fd.BotId, fd.Code)
			if err != nil {
				log.Printf("Insert error. %v", err)
				return nil
			}

		} else {
			log.Printf("Function %q in Bot %q already exists.", fd.Name, fd.BotId)
			return nil
		}

		fd.Version = 0
	} else {
		return nil
	}
	return fd
}

func (fm FunctionManagerDb) validateAppId(appId string) bool {
	return HasAccessToken(appId)
}

func (fm FunctionManagerDb) ExecuteFunction(call *FunctionCall) (string, error) {
	if fm.validateAppId(call.AppId) {
		fd := fm.findFunction(call.Name, call.BotId)

		if fd != nil {
			result, err := executor.WrapAndExecuteJSFunction(fd.Code, call.Params)
			return result, err
		} else {
			return "", fmt.Errorf("function %q not found for bot %q", call.Name, call.BotId)
		}
	} else {
		return "", fmt.Errorf("appid not valid %q", call.AppId)
	}
}

func (fm FunctionManagerDb) UpdateFunction(fd *FunctionDescription) bool {

	if fm.validateAppId(fd.AppId) {
		_, err := db.Exec("UPDATE functions set code = ? where name = ? and botid = ?", fd.Code, fd.Name, fd.BotId)
		if err != nil {
			log.Printf("Insert error. %v", err)
			return false
		}
	} else {
		return false
	}
	return true
}

func (fm FunctionManagerDb) DeleteFunction(fd *FunctionDescription) bool {

	if fm.validateAppId(fd.AppId) {
		_, err := db.Exec("DELETE FROM functions where name = ? and botid = ?", fd.Name, fd.BotId)
		if err != nil {
			log.Printf("Delete error. %v", err)
			return false
		}
	} else {
		return false
	}
	return true
}

func (fm FunctionManagerDb) GetFunctionList(botId string, appId string) []*FunctionDescription {

	if fm.validateAppId(appId) {
		rows, err := db.Query("SELECT name, botid, code, version FROM functions WHERE botid = ?", botId)
		if err != nil {
			log.Printf("GetFunctionList error. %v", err)
			return nil
		}

		defer rows.Close()

		var functionDescriptions []*FunctionDescription

		// Loop through rows, using Scan to assign column data to struct fields.
		for rows.Next() {
			fd := new(FunctionDescription)
			err := rows.Scan(&fd.Name, &fd.BotId, &fd.Code, &fd.Version)

			if err != nil {
				log.Printf("GetFunctionList error. %v", err)
				return nil
			}

			functionDescriptions = append(functionDescriptions, fd)
		}

		err = rows.Err()

		if err != nil {
			log.Printf("GetFunctionList error. %v", err)
			return nil
		}

		return functionDescriptions
	}

	return nil
}
