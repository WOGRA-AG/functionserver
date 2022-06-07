package manager

import (
	"testing"
)

func TestAddAndFindFunction(t *testing.T) {

	fm := NewMemory()

	var testFunction = generateTestFunction()

	fm.AddFunction(testFunction)

	if fm.FindFunction(testFunction) == nil {
		t.Fatalf("Test of AddFunction failed. Added Function not found")
	}
}

func TestSearchWithInvalidBotId(t *testing.T) {

	fm := NewMemory()

	var testFunction = generateTestFunction()
	testFunction.BotId = "5000"

	fm.AddFunction(testFunction)

	testFunction.BotId = "1"

	if fm.FindFunction(testFunction) != nil {
		t.Fatalf("Test of TestSearchWithInvalidBotIdFunction failed. not null returned")
	}
}

func TestSearchInvalidFunctionName(t *testing.T) {

	fm := NewMemory()

	var testFunction = generateTestFunction()

	fm.AddFunction(testFunction)

	testFunction.Name = "blubb"

	if fm.FindFunction(testFunction) != nil {
		t.Fatalf("Test of TestSearchWithInvalidBotIdFunction failed. not null returned")
	}
}

func TestExecuteFunction(t *testing.T) {

	fm := NewMemory()
	var testFunction = generateTestFunction()

	fm.AddFunction(testFunction)

	call := new(FunctionCall)
	call.AppId = testFunction.AppId
	call.BotId = testFunction.BotId
	call.Name = testFunction.Name

	result, err := fm.ExecuteFunction(call)

	if err == nil {
		if result != "4" {
			t.Fatalf("Test of TestExecuteFunction failed. Expected: %q, got %q.", "4", result)
		}
	} else {
		t.Fatalf("Test of TestExecuteFunction failed. Error returned %q.", err)
	}
}
