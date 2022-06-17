package manager

import (
	"testing"
	"wogra.com/config"
)

func TestInitDbFunctionManager(t *testing.T) {
	var dbConf config.DbConfig

	dbConf.DatabaseName = "maxbot"
	dbConf.DatabaseUrl = "127.0.0.1"
	dbConf.DatabasePort = "3306"
	dbConf.Login = "root"
	dbConf.Password = "root"

	var fm = new(FunctionManagerDb)

	if !fm.init(dbConf) {
		t.Fatalf("Test of TestInitDbFunctionManager failed.")
	}
}

func generateTestFunction(t *testing.T, fm *FunctionManagerDb) *FunctionDescription {
	var newFunction = new(FunctionDescription)
	newFunction.Name = "test"
	newFunction.BotId = "1"
	newFunction.Code = "return 2 + 2"
	newFunction.AppId = createStandardUserAppId(t)

	appId := createSuperUserAppId(t)

	if fm.AddCredentials(appId, newFunction.AppId, newFunction.BotId) == false {
		t.Fatalf("Assignment of credentials failed in generateTestFunction superuser %q, appId %q, botId %q", appId, newFunction.AppId, newFunction.BotId)
		return nil
	}

	if fm.DeleteAppId(appId, appId) == false {
		t.Fatalf("generateTestFunction failed. (Delete call)")
	}

	return newFunction
}

func generateTestFunctionWithInvalidCredentialAssignments(t *testing.T, fm *FunctionManagerDb) *FunctionDescription {
	var newFunction = new(FunctionDescription)
	newFunction.Name = "test"
	newFunction.BotId = "1"
	newFunction.Code = "return 2 + 2"
	newFunction.AppId = createStandardUserAppId(t)

	appId := createSuperUserAppId(t)

	if fm.AddCredentials(appId, newFunction.AppId, "blubb") == false {
		t.Fatalf("Assignment of credentials failed in generateTestFunction superuser %q, appId %q, botId %q", appId, newFunction.AppId, newFunction.BotId)
		return nil
	}

	if fm.DeleteAppId(appId, appId) == false {
		t.Fatalf("generateTestFunction failed. (Delete call)")
	}

	return newFunction
}

func deleteAppId(appId string, fm *FunctionManagerDb, t *testing.T) {
	superuserappId := createSuperUserAppId(t)
	fm.DeleteAppId(superuserappId, appId)
	fm.DeleteAppId(superuserappId, superuserappId)
}

func TestCreateSuperUserAppId(t *testing.T) {
	var fm = NewDb()
	appId := createSuperUserAppId(t)
	if "" == appId {
		t.Fatalf("Test of TestCreateSuperUserAppId failed. (Create call)")
	} else {

		if fm.validateAppIdSuperUser(appId) == false {
			t.Fatalf("Test of TestCreateSuperUserAppId failed. (Validate call)")
		} else {
			if fm.DeleteAppId(appId, appId) == false {
				t.Fatalf("Test of TestCreateSuperUserAppId failed. (Delete call)")
			}
		}
	}
}

func TestCreateStandardUserAppId(t *testing.T) {
	var fm = NewDb()
	appId := createStandardUserAppId(t)

	if "" == appId {
		t.Fatalf("Test of TestCreateStandardUserAppId failed. (Create call)")
	} else {

		superuserAppId := createSuperUserAppId(t)
		if fm.DeleteAppId(superuserAppId, appId) == false {
			t.Fatalf("Test of TestCreateStandardUserAppId failed. (Delete call)")
		}

		if fm.DeleteAppId(superuserAppId, superuserAppId) == false {
			t.Fatalf("Test of TestCreateStandardUserAppId failed. (Delete call)")
		}
	}
}

func createSuperUserAppId(t *testing.T) string {
	return createAppId(t, true)
}

func createStandardUserAppId(t *testing.T) string {
	return createAppId(t, false)
}

func createAppId(t *testing.T, superuser bool) string {
	var fm = NewDb()

	if fm != nil {
		appId := fm.createAppId("test", "test", superuser)

		if appId == "" {
			t.Fatalf("Test of TestAddAndDeleteFunctionInDB failed. (Add call)")
			return ""
		}

		return appId
	}

	return ""
}

func TestAddAndDeleteFunctionInDB(t *testing.T) {
	var fm = NewDb()

	if fm != nil {

		var testFunction = generateTestFunction(t, fm)

		if fm.AddFunction(testFunction) == nil {
			t.Fatalf("Test of TestAddAndDeleteFunctionInDB failed. (Add call) %v", testFunction)
		} else {

			fd := fm.FindFunction(testFunction)

			if fd == nil {
				t.Fatalf("Test of TestAddAndDeleteFunctionInDB failed. (Find call)")
			} else {

				if fd.Code != testFunction.Code {
					t.Fatalf("Test of TestAddAndDeleteFunctionInDB failed. Different content. Expected: %q got: %q", testFunction.Code, fd.Code)
				}

				if fm.DeleteFunction(testFunction) == false {
					t.Fatalf("Test of TestAddAndDeleteFunctionInDB failed. (Delete call)")
				}
			}

			deleteAppId(testFunction.AppId, fm, t)
		}
	}
}

func TestAddAndDeleteFunctionInDBWithInvalidCredentials(t *testing.T) {
	var fm = NewDb()

	if fm != nil {

		var testFunction = generateTestFunctionWithInvalidCredentialAssignments(t, fm)

		if fm.AddFunction(testFunction) != nil {
			t.Fatalf("Test of TestAddAndDeleteFunctionInDBWithInvalidCredentials failed. (Add call) %v", testFunction)
		}

		deleteAppId(testFunction.AppId, fm, t)

	}
}

func TestExecuteFunctionInDB(t *testing.T) {
	var fm = NewDb()

	if fm != nil {

		var testFunction = generateTestFunction(t, fm)

		if fm.AddFunction(testFunction) == nil {
			t.Fatalf("Test of TestExecuteFunctionInDB failed. (Add call)")
		} else {

			fd := fm.FindFunction(testFunction)

			if fd == nil {
				t.Fatalf("Test of TestExecuteFunctionInDB failed. (Find call)")
			} else {

				fc := new(FunctionCall)
				fc.Name = testFunction.Name
				fc.BotId = testFunction.BotId
				fc.AppId = testFunction.AppId
				expectation := "4"
				result, _ := fm.ExecuteFunction(fc)
				if result != expectation {
					t.Fatalf("Execution delivers unexpected result. Expected: %q Got: %q", expectation, result)
				}

				if fm.DeleteFunction(testFunction) == false {
					t.Fatalf("Test of TestExecuteFunctionInDB failed. (Delete call)")
				}
			}

			deleteAppId(testFunction.AppId, fm, t)
		}
	}
}

func TestUpdateFunctionInDB(t *testing.T) {
	var fm = NewDb()

	if fm != nil {

		var testFunction = generateTestFunction(t, fm)

		if fm.AddFunction(testFunction) == nil {
			t.Fatalf("Test of TestUpdateFunctionInDB failed. (Add call)")
		} else {

			fd := fm.FindFunction(testFunction)

			if fd == nil {
				t.Fatalf("Test of TestUpdateFunctionInDB failed. (Find call)")
			} else {

				testFunction.Code = "return 4 + 4"

				if fm.UpdateFunction(testFunction) == false {
					t.Fatalf("Update function failed!")
				}

				fd := fm.FindFunction(testFunction)

				if fd == nil {
					t.Fatalf("Test of TestUpdateFunctionInDB failed. (Second Find call)")
				} else {
					if fd.Code != testFunction.Code {
						t.Fatalf("Expectatioin error. Expected: %q Got: %q", testFunction.Code, fd.Code)
					}
				}

				if fm.DeleteFunction(testFunction) == false {
					t.Fatalf("Test of TestUpdateFunctionInDB failed. (Delete call)")
				}
			}

			deleteAppId(testFunction.AppId, fm, t)
		}
	}
}

func TestGetEmptyFunctionListInDBWithSuperUser(t *testing.T) {
	var fm = NewDb()

	if fm != nil {
		appId := createSuperUserAppId(t)
		list := fm.GetFunctionList(appId, "bluub")

		if len(list) > 0 {
			t.Fatalf("Test of TestGetEmptyFunctionListInDBWithSuperUser failed. (Delete call) elements: %d", len(list))
		}

		deleteAppId(appId, fm, t)
	}
}

func TestGetEmptyFunctionListInDBWithStandardUser(t *testing.T) {
	var fm = NewDb()

	if fm != nil {
		appId := createStandardUserAppId(t)
		list := fm.GetFunctionList(appId, "bluub")

		if len(list) > 0 {
			t.Fatalf("Test of TestGetEmptyFunctionListInDBWithSuperUser failed. (Delete call) elements: %d", len(list))
		}

		deleteAppId(appId, fm, t)
	}
}

func TestGetFunctionListWithTwoEntries(t *testing.T) {
	var fm = NewDb()

	if fm != nil {
		var testFunction1 = generateTestFunction(t, fm)
		var testFunction2 = generateTestFunction(t, fm)

		testFunction1.Name = testFunction1.Name + "11"
		testFunction2.Name = testFunction1.Name + "22"
		fm.AddFunction(testFunction1)
		fm.AddFunction(testFunction2)

		list := fm.GetFunctionList(testFunction1.AppId, testFunction1.BotId)

		if len(list) != 2 {
			t.Fatalf("Test of TestGetEmptyFunctionListInDBWithSuperUser failed. (Delete call) elements: %d", len(list))
		}

		fm.DeleteFunction(testFunction1)
		fm.DeleteFunction(testFunction2)

		deleteAppId(testFunction2.AppId, fm, t)
		deleteAppId(testFunction1.AppId, fm, t)
	}
}

/*

 */
