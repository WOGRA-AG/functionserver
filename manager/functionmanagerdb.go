package manager

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
	"wogra.com/config"
	"wogra.com/executor"
)

var db *sql.DB

type FunctionManagerDb struct {
	//db *sql.DB
}

func NewDb() *FunctionManagerDb {

	var fm = new(FunctionManagerDb)
	if fm.init(config.ReadDatabaseConfiguration()) {
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

func (fm FunctionManagerDb) init(dbConfig config.DbConfig) bool {

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
	if fm.validateAppId(fd.AppId, fd.BotId) {
		return fm.findFunction(fd.Name, fd.BotId)
	} else {
		return nil
	}

	return fd
}

func (fm FunctionManagerDb) AddFunction(fd *FunctionDescription) *FunctionDescription {

	if fm.validateAppId(fd.AppId, fd.BotId) {
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

func (fm FunctionManagerDb) validateAppId(appId string, botId string) bool {
	rows, err := db.Query("SELECT count(*) FROM appid_bots WHERE appId = ? and botid = ?", appId, botId)

	defer rows.Close()

	if err != nil {
		log.Printf("validateAppId error. %v", err)
		return false
	} else {
		for rows.Next() {
			var count int
			if err := rows.Scan(&count); err != nil {
				log.Printf("validateAppId error. %v", err)
			} else {
				if count > 0 {
					return true
				}
			}
		}
	}

	return fm.validateAppIdSuperUser(appId)
}

func (fm FunctionManagerDb) validateAppIdSuperUser(appId string) bool {
	rows, err := db.Query("SELECT count(*) FROM appids WHERE appId = ? and superuser = 1", appId)

	defer rows.Close()

	if err != nil {
		log.Printf("validateAppIdSuperUser error. %v", err)
		return false
	} else {
		for rows.Next() {
			var count int
			if err := rows.Scan(&count); err != nil {
				log.Printf("validateAppIdSuperUser error. %v", err)
			} else {
				if count > 0 {
					return true
				}
			}
		}
	}

	return false
}

func (fm FunctionManagerDb) ExecuteFunction(call *FunctionCall) (string, error) {
	if fm.validateAppId(call.AppId, call.BotId) {
		fd := fm.findFunction(call.Name, call.BotId)

		if fd != nil {
			result, err := executor.ExecuteFunction(call.BotId, fd.Code, call.Params)
			return result, err
		} else {
			return "", fmt.Errorf("function %q not found for bot %q", call.Name, call.BotId)
		}
	} else {
		return "", fmt.Errorf("appid not valid %q", call.AppId)
	}
}

func (fm FunctionManagerDb) UpdateFunction(fd *FunctionDescription) bool {

	if fm.validateAppId(fd.AppId, fd.BotId) {
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

	if fm.validateAppId(fd.AppId, fd.BotId) {
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

func (fm FunctionManagerDb) GetFunctionList(appId string, botId string) []*FunctionDescription {

	if fm.validateAppId(appId, botId) {
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

func (fm FunctionManagerDb) CreateAppId(appId string, owner string, contact string) string {
	if fm.validateAppIdSuperUser(appId) {
		return fm.createAppId(owner, contact, false)
	}

	return ""
}

func (fm FunctionManagerDb) CreateAppIdSuperUser(appId string, owner string, contact string) string {
	if fm.validateAppIdSuperUser(appId) {
		return fm.createAppId(owner, contact, true)
	}

	return ""
}

func (fm FunctionManagerDb) createAppId(owner string, contact string, superuser bool) string {

	if owner != "" && contact != "" {
		uuid := uuid.New().String()
		_, err := db.Exec("insert into appids (appid, owner, contact, superuser) values(?,?,?,?)", uuid, owner, contact, superuser)

		if err != nil {
			log.Printf("AppId creation error. %v", err)
		} else {
			return uuid
		}
	} else {
		log.Printf("AppId creation error. Missing data owner %q or contact %q", owner, contact)
	}

	return ""
}

func (fm FunctionManagerDb) DeleteAppId(appId string, appIdToDelete string) bool {
	if fm.validateAppIdSuperUser(appId) {
		_, err := db.Exec("delete from appids where appid = ?", appIdToDelete)

		if err != nil {
			log.Printf("AppId creation error. %v", err)
		} else {
			return true
		}
	}

	return false
}

func (fm FunctionManagerDb) CheckCredentials(appId string, botId string) bool {
	return fm.validateAppId(appId, botId)
}
func (fm FunctionManagerDb) AddCredentials(superuserAppId string, appId string, botId string) bool {
	if fm.validateAppIdSuperUser(superuserAppId) {
		_, err := db.Exec("insert into appid_bots (appid, botid) values(?,?)", appId, botId)

		if err != nil {
			log.Printf("Credential creation error. %v", err)
		} else {
			return true
		}
	}

	return false
}

func (fm FunctionManagerDb) DeleteCredentials(superuserAppId string, appId string, botId string) bool {
	if fm.validateAppIdSuperUser(superuserAppId) {
		_, err := db.Exec("delete from appid_bots where appid = ? and botid = ?", appId, botId)

		if err != nil {
			log.Printf("Credential deletion error. %v", err)
		} else {
			return true
		}
	}

	return false
}
