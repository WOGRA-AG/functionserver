package executor

import (
	"fmt"
	"testing"
)

func TestExecuteJSExpressionWithoutParams(t *testing.T) {

	code := "2 + 2;"
	expectation := "4"
	var emptyParams map[string]string
	myResult, err := executeJSFunction("1", code, emptyParams)

	if err != nil || myResult != expectation {
		t.Fatalf("Test of ExecuteFunction failed. Expected: %q, Got: %q Error: %q", expectation, myResult, err)
	}
}

func TestExecuteFunctionWithoutParams(t *testing.T) {

	code := "function myFunc() { return 2 + 2; } myFunc();"
	expectation := "4"
	var emptyParams map[string]string
	myResult, err := executeJSFunction("1", code, emptyParams)

	if err != nil || myResult != expectation {
		t.Fatalf("Test of ExecuteFunction failed. Expected: %q, Got: %q Error: %q", expectation, myResult, err)
	}
}

func TestExecuteFunctionWithSyntaxErrors(t *testing.T) {

	code := "function myFunc() { returnadhsjdghjd; 2 + 2; } myFunc();"
	var emptyParams map[string]string
	myResult, err := executeJSFunction("1", code, emptyParams)

	if err == nil {
		t.Fatalf("Test of TestExecuteFunctionWithSyntaxErrors failed. Expected error got nil Result: %q", myResult)
	}
}

func TestWrapAndExecuteFunctionWithParams(t *testing.T) {

	code := "return a;"
	apple := "apple"
	banana := "banana"
	params := map[string]string{"a": apple, "b": banana}

	result, err := ExecuteFunction("1", code, params)

	if err != nil || result != apple {
		t.Fatalf("Test of TestWrapAndExecuteFunctionWithParams failed. Expected: %q, Got: %q Error: %q", apple, result, err)
	}
}

func TestWrapAndExecuteJSExpressionWithoutParams(t *testing.T) {

	code := "return 2 + 2;"
	expectation := "4"
	var emptyParams map[string]string
	myResult, err := ExecuteFunction("1", code, emptyParams)

	if err != nil || myResult != expectation {
		t.Fatalf("Test of ExecuteFunction failed. Expected: %q, Got: %q Error: %q", expectation, myResult, err)
	}
}

func TestExecuteJSExpressionGetRequestHttps(t *testing.T) {

	code := "httpGet(\"https://wogra.com\")"

	var emptyParams map[string]string
	myResult, err := executeJSFunction("1", code, emptyParams)

	if err != nil || len(myResult) == 0 {
		t.Fatalf("Test of TestExecuteJSExpressionGetRequestHttps failed. Result %q ResultError: %q", myResult, err)
	}
}

func TestExecuteJSExpressionGetRequestHttp(t *testing.T) {

	code := "httpGet(\"http://wogra.com\")"

	var emptyParams map[string]string
	myResult, err := executeJSFunction("1", code, emptyParams)

	if err != nil || len(myResult) == 0 {
		t.Fatalf("Test of TestExecuteJSExpressionGetRequestHttp failed. Result %q ResultError: %q", myResult, err)
	}
}

func TestExecuteJSExpressionGetRequestLocalHostHttp(t *testing.T) {

	code := "httpGet(\"http://localhost\")"

	var emptyParams map[string]string
	myResult, err := executeJSFunction("1", code, emptyParams)

	if myResult != "not allowed" {
		t.Fatalf("Test of TestExecuteJSExpressionGetRequestHttp failed. Result %q ResultError: %q", myResult, err)
	}
}

func TestExecuteJSExpressionGetRequestLocalHostIpHttp(t *testing.T) {

	code := "httpGet(\"http://127.0.0.1\")"

	var emptyParams map[string]string
	myResult, err := executeJSFunction("1", code, emptyParams)

	if myResult != "not allowed" {
		t.Fatalf("Test of TestExecuteJSExpressionGetRequestHttp failed. Result %q ResultError: %q", myResult, err)
	}
}

func TestExecuteJSExpressionPostRequestLocalHostHttp(t *testing.T) {

	code := "httpPost(\"http://localhost\", \"json\", \"{ \\\"Test\\\" : \\\"blubb\\\"}\")"

	var emptyParams map[string]string
	myResult, err := executeJSFunction("1", code, emptyParams)

	if myResult != "not allowed" {
		t.Fatalf("Test of TestExecuteJSExpressionGetRequestHttp failed. Result %q ResultError: %q", myResult, err)
	}
}

func TestExecuteJSExpressionPostRequestLocalHostIpHttp(t *testing.T) {

	code := "httpPost(\"http://127.0.0.1\", \"json\", \"{ \\\"Test\\\" : \\\"blubb\\\"}\")"

	var emptyParams map[string]string
	myResult, err := executeJSFunction("1", code, emptyParams)

	if myResult != "not allowed" {
		t.Fatalf("Test of TestExecuteJSExpressionGetRequestHttp failed. Result %q ResultError: %q", myResult, err)
	}
}

func TestExecuteJSExpressionGetRequestUnknownUrlHttps(t *testing.T) {

	code := "httpGet(\"https://hjkkjhhkhkj.de\")"

	var emptyParams map[string]string
	myResult, err := executeJSFunction("1", code, emptyParams)

	if len(myResult) > 0 {
		t.Fatalf("Test of TestExecuteJSExpressionGetRequestUnknownUrlHttps failed. Result %q ResultError: %q", myResult, err)
	}
}

func TestExecuteJSExpressionGetRequestUnknownUrlHttp(t *testing.T) {

	code := "httpGet(\"http://hjkkjhhkhkj.de\")"

	var emptyParams map[string]string
	myResult, err := executeJSFunction("1", code, emptyParams)

	if len(myResult) > 0 {
		t.Fatalf("Test of TestExecuteJSExpressionGetRequestUnknownUrlHttp failed. Result %q ResultError: %q", myResult, err)
	}
}

func TestExecuteJSExpressionGetRequestUnknownProtocol(t *testing.T) {

	code := "httpGet(\"tps://wogra.com\")"

	var emptyParams map[string]string
	myResult, err := executeJSFunction("1", code, emptyParams)

	if len(myResult) > 0 {
		t.Fatalf("Test of TestExecuteJSExpressionGetRequestUnknownProtocol failed. Result %q ResultError: %q", myResult, err)
	}
}

func TestExecuteJSExpressionPostRequestHttps(t *testing.T) {

	code := "httpPost(\"https://wogra.com\", \"json\", \"{ \\\"Test\\\" : \\\"blubb\\\"}\")"

	var emptyParams map[string]string
	myResult, err := executeJSFunction("1", code, emptyParams)

	if err != nil || len(myResult) == 0 {
		t.Fatalf("Test of TestExecuteJSExpressionPostRequestHttps failed. Result %q ResultError: %q", myResult, err)
	}
}

func TestExecuteJSExpressionPostRequestUnknownProtocol(t *testing.T) {

	code := "httpPost(\"tps://wogra.com\", \"json\", \"{ \\\"Test\\\" : \\\"blubb\\\"}\")"

	var emptyParams map[string]string
	myResult, err := executeJSFunction("1", code, emptyParams)

	if len(myResult) > 0 {
		t.Fatalf("Test of TestExecuteJSExpressionPostRequestUnknownProtocol failed. Result %q ResultError: %q", myResult, err)
	}
}

func TestSaveAndLoadDocument(t *testing.T) {

	key := "k1"
	value := "v1"
	code := fmt.Sprintf("saveDocument(\"%s\",\"%s\"); var result = loadDocument(\"%s\"); return result;", key, value, key)

	t.Logf("Code %q", code)
	var emptyParams map[string]string
	myResult, err := ExecuteFunction("2", code, emptyParams)

	if err != nil {
		t.Fatalf("Test of TestSaveAndLoadDocument failed.Error occured: %q", err)
	}

	if myResult != value {
		t.Fatalf("Test of TestSaveAndLoadDocument failed. Expected Result: %q, Got: %q", value, myResult)
	}
}

func TestSaveAndLoadAndDeleteDocument(t *testing.T) {

	key := "k2"
	value := "v2"
	code := fmt.Sprintf("saveDocument(\"%s\",\"%s\"); deleteDocument(\"%s\"); var result = loadDocument(\"%s\"); return result;", key, value, key, key)

	t.Logf("Code %q", code)
	var emptyParams map[string]string
	myResult, err := ExecuteFunction("2", code, emptyParams)

	if err != nil {
		t.Fatalf("Test of TestSaveAndLoadDocument failed.Error occured: %q", err)
	}

	if myResult != "" {
		t.Fatalf("Test of TestSaveAndLoadDocument failed. Expected Result: %q, Got: %q", "", myResult)
	}
}

func TestExecuteJSEndlessLoopFor(t *testing.T) {

	code := "var a = 0; for(;;) {a = a+1;}"

	var emptyParams map[string]string
	myResult, _ := executeJSFunction("1", code, emptyParams)

	if len(myResult) != 0 {
		t.Fatal("Test of TestExecuteJSEndlessLoop myResult is not nil")
	}
}

func TestExecuteJSEndlessLoopDoWhile(t *testing.T) {

	code := "var a = 0; do {a = a+1;}while(true);"

	var emptyParams map[string]string
	myResult, _ := executeJSFunction("1", code, emptyParams)

	if len(myResult) != 0 {
		t.Fatal("Test of TestExecuteJSEndlessLoop myResult is not nil")
	}
}
