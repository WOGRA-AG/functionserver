package executor

import (
	"testing"
)

func TestExecuteJSExpressionWithoutParams(t *testing.T) {

	code := "2 + 2;"
	expectation := "4"
	myResult, err := ExecuteJSFunction(code, nil)

	if err != nil || myResult != expectation {
		t.Fatalf("Test of ExecuteJSFunction failed. Expected: %q, Got: %q Error: %q", expectation, myResult, err)
	}
}

func TestExecuteJSFunctionWithoutParams(t *testing.T) {

	code := "function myFunc() { return 2 + 2; } myFunc();"
	expectation := "4"
	myResult, err := ExecuteJSFunction(code, nil)

	if err != nil || myResult != expectation {
		t.Fatalf("Test of ExecuteJSFunction failed. Expected: %q, Got: %q Error: %q", expectation, myResult, err)
	}
}

func TestExecuteJSFunctionWithSyntaxErrors(t *testing.T) {

	code := "function myFunc() { returnadhsjdghjd; 2 + 2; } myFunc();"
	result, err := ExecuteJSFunction(code, nil)

	if err == nil {
		t.Fatalf("Test of TestExecuteJSFunctionWithSyntaxErrors failed. Expected error got nil Result: %q", result)
	}
}

func TestExecuteJSFunctionWithParams(t *testing.T) {

	code := "a;"
	apple := "apple"
	banana := "banana"
	params := map[string]string{"a": apple, "b": banana}

	result, err := ExecuteJSFunction(code, params)

	if err != nil || result != apple {
		t.Fatalf("Test of ExecuteJSFunction failed. Expected: %q, Got: %q Error: %q", params, result, err)
	}
}

func TestWrapAndExecuteJSFunctionWithParams(t *testing.T) {

	code := "return a;"
	apple := "apple"
	banana := "banana"
	params := map[string]string{"a": apple, "b": banana}

	result, err := WrapAndExecuteJSFunction(code, params)

	if err != nil || result != apple {
		t.Fatalf("Test of WrapAndExecuteJSFunction failed. Expected: %q, Got: %q Error: %q", apple, result, err)
	}
}

func TestWrapAndExecuteJSExpressionWithoutParams(t *testing.T) {

	code := "return 2 + 2;"
	expectation := "4"
	myResult, err := WrapAndExecuteJSFunction(code, nil)

	if err != nil || myResult != expectation {
		t.Fatalf("Test of WrapAndExecuteJSFunction failed. Expected: %q, Got: %q Error: %q", expectation, myResult, err)
	}
}

func TestExecuteJSExpressionGetRequestHttps(t *testing.T) {

	code := "httpGet(\"https://wogra.com\")"

	myResult, err := ExecuteJSFunction(code, nil)

	if err != nil || len(myResult) == 0 {
		t.Fatalf("Test of TestExecuteJSExpressionGetRequestHttps failed. Result %q ResultError: %q", myResult, err)
	}
}

func TestExecuteJSExpressionGetRequestHttp(t *testing.T) {

	code := "httpGet(\"http://wogra.com\")"

	myResult, err := ExecuteJSFunction(code, nil)

	if err != nil || len(myResult) == 0 {
		t.Fatalf("Test of TestExecuteJSExpressionGetRequestHttp failed. Result %q ResultError: %q", myResult, err)
	}
}

func TestExecuteJSExpressionGetRequestUnknownUrlHttps(t *testing.T) {

	code := "httpGet(\"https://hjkkjhhkhkj.de\")"

	myResult, err := ExecuteJSFunction(code, nil)

	if len(myResult) > 0 {
		t.Fatalf("Test of TestExecuteJSExpressionGetRequestUnknownUrlHttps failed. Result %q ResultError: %q", myResult, err)
	}
}

func TestExecuteJSExpressionGetRequestUnknownUrlHttp(t *testing.T) {

	code := "httpGet(\"http://hjkkjhhkhkj.de\")"

	myResult, err := ExecuteJSFunction(code, nil)

	if len(myResult) > 0 {
		t.Fatalf("Test of TestExecuteJSExpressionGetRequestUnknownUrlHttp failed. Result %q ResultError: %q", myResult, err)
	}
}

func TestExecuteJSExpressionGetRequestUnknownProtocol(t *testing.T) {

	code := "httpGet(\"tps://wogra.com\")"

	myResult, err := ExecuteJSFunction(code, nil)

	if len(myResult) > 0 {
		t.Fatalf("Test of TestExecuteJSExpressionGetRequestUnknownProtocol failed. Result %q ResultError: %q", myResult, err)
	}
}

func TestExecuteJSExpressionPostRequestHttps(t *testing.T) {

	code := "httpPost(\"https://wogra.com\", \"json\", \"{ \\\"Test\\\" : \\\"blubb\\\"}\")"

	myResult, err := ExecuteJSFunction(code, nil)

	if err != nil || len(myResult) == 0 {
		t.Fatalf("Test of TestExecuteJSExpressionPostRequestHttps failed. Result %q ResultError: %q", myResult, err)
	}
}

func TestExecuteJSExpressionPostRequestUnknownProtocol(t *testing.T) {

	code := "httpPost(\"tps://wogra.com\", \"json\", \"{ \\\"Test\\\" : \\\"blubb\\\"}\")"

	myResult, err := ExecuteJSFunction(code, nil)

	if len(myResult) > 0 {
		t.Fatalf("Test of TestExecuteJSExpressionPostRequestUnknownProtocol failed. Result %q ResultError: %q", myResult, err)
	}
}
