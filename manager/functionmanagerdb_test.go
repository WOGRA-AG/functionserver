package manager

import (
	"testing"
)

func TestInitDbFunctionManager(t *testing.T) {
	var dbConf DbConfig

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

func generateTestFunction() *FunctionDescription {
	var newFunction = new(FunctionDescription)
	newFunction.Name = "test"
	newFunction.BotId = "1"
	newFunction.Code = "return 2 + 2"
	newFunction.AppId = "1"

	return newFunction
}

func TestAddAndDeleteFunctionInDB(t *testing.T) {
	var fm = NewDb()

	if fm != nil {

		var testFunction = generateTestFunction()

		if fm.AddFunction(testFunction) == nil {
			t.Fatalf("Test of TestAddAndDeleteFunctionInDB failed. (Add call)")
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
		}
	}
}

func TestExecuteFunctionInDB(t *testing.T) {
	var fm = NewDb()

	if fm != nil {

		var testFunction = generateTestFunction()

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
		}
	}
}

func TestUpdateFunctionInDB(t *testing.T) {
	var fm = NewDb()

	if fm != nil {

		var testFunction = generateTestFunction()

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
		}
	}
}
