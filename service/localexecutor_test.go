package main

import (
	"testing"
)

func TestParamsReader(t *testing.T) {

	args := []string{"p1=v1", "p2=v2"}

	params := readParams(args)

	if len(params) != 2 {
		t.Fatalf("Test of TestParamsReader failed. Expected Params: 2 Got %d", len(params))
	}

	if params["p1"] != "v1" {
		t.Fatalf("Test of TestParamsReader failed. Expected first parameter: v1 Got %q", params["p1"])
	}
}

func TestParamsReaderWithoutContent(t *testing.T) {
	params := readParams(nil)

	if params != nil {
		t.Fatalf("Test of TestParamsReader failed. Expected Params: nil Got %d", len(params))
	}
}

func TestExecuteFileWithoutParams(t *testing.T) {
	args := []string{"-execute", "./TestAddWithoutParams.js"}
	result := runAsLocalExecutor(args)

	if result != true {
		t.Fatalf("Test of TestExecuteFileWithoutParams failed.")
	}
}

func TestExecuteFileWithoutParamsCheckResult(t *testing.T) {
	result := executeLocal("./TestAddWithoutParams.js", nil)

	if result != "4" {
		t.Fatalf("Test of TestExecuteFileWithoutParams failed. Expected 4 Got: %q", result)
	}
}

func TestExecuteFileWithParams(t *testing.T) {
	args := []string{"-execute", "./TestAddWithParams.js", "p1=2", "p2=3"}
	result := runAsLocalExecutor(args)

	if result != true {
		t.Fatalf("Test of TestExecuteFileWithoutParams failed.")
	}
}

func TestExecuteFileWithParamsCheckResult(t *testing.T) {

	paramMap := make(map[string]string)
	paramMap["p1"] = "2"
	paramMap["p2"] = "3"

	result := executeLocal("./TestAddWithParams.js", paramMap)

	if result != "5" {
		t.Fatalf("Test of TestExecuteFileWithoutParams failed. Expected 5 Got: %q", result)
	}
}

func TestExecuteFileWithEndlessLoop(t *testing.T) {

	result := executeLocal("./TestEndlessLoop.js", nil)

	if result != "" {
		t.Fatalf("Test of TestExecuteFileWithEndlessLoop failed. Expected ? Got: %q", result)
	}
}

func TestExecuteFileWithHttpGet(t *testing.T) {

	result := executeLocal("./TestHttpGet.js", nil)

	if len(result) == 0 {
		t.Fatalf("Test of TestExecuteFileWithHttpGet failed. Expected not empty string Got: nothing")
	}
}

func TestExecuteFileWithHttpPost(t *testing.T) {

	result := executeLocal("./TestHttpPost.js", nil)

	if len(result) == 0 {
		t.Fatalf("Test of TestExecuteFileWithHttpGet failed. Expected not empty string Got: nothing")
	}
}

func TestExecuteFileWithForbiddenHttpGet(t *testing.T) {

	result := executeLocal("./TestForbiddenHttpGet.js", nil)

	if result != "not allowed" {
		t.Fatalf("Test of TestExecuteFileWithForbiddenHttpGet failed. Expected not allowed Got: %q", result)
	}
}

func TestExecuteFileWithForbiddenHttpPost(t *testing.T) {

	result := executeLocal("./TestForbiddenHttpPost.js", nil)

	if result != "not allowed" {
		t.Fatalf("Test of TestExecuteFileWithForbiddenHttpPost failed. Expected  not allowed Got: %q", result)
	}
}
