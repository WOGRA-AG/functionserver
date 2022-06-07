package executor

import (
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"

	"github.com/robertkrimen/otto"
)

var (
	functionName string = "executorFunctionWrapper"
)

func ExecuteJSFunction(functionCode string, functionParams map[string]string) (string, error) {

	vm := otto.New()
	addEnhancedFeaturesToScriptEngine(vm)

	for keyParam, valueParam := range functionParams {
		vm.Set(keyParam, valueParam)
	}

	result, err := vm.Run(functionCode)

	if err != nil {
		log.Print(err)
	}

	return result.String(), err
}

func WrapJsFunction(functionCode string) (string, error) {

	if len(functionCode) == 0 {
		return "", errors.New("length of code is zero")
	}

	functionCode = fmt.Sprintf("function %s()\n{\n%s\n}\n%s();", functionName, functionCode, functionName)
	return functionCode, nil
}

func WrapAndExecuteJSFunction(functionCode string, functionParams map[string]string) (string, error) {

	newFunctionCode, err := WrapJsFunction(functionCode)

	if err == nil {
		return ExecuteJSFunction(newFunctionCode, functionParams)
	} else {
		return "", err
	}
}

func addEnhancedFeaturesToScriptEngine(engine *otto.Otto) {
	if engine != nil {
		addHttpFeaturesToScriptEngine(engine)
	}
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
			response, err := http.Get(call.Argument(0).String())

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
		})
	}
}

func addHttpPostFeatureToScriptEngine(engine *otto.Otto) {
	if engine != nil {
		engine.Set("httpPost", func(call otto.FunctionCall) otto.Value {

			myReader := strings.NewReader(call.Argument(2).String())
			response, err := http.Post(call.Argument(0).String(), call.Argument(1).String(), myReader)

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
		})
	}
}

func ExecuteWasmFunction(functionCode string, functionParams map[string]string) (string, error) {

	// todo
	return "", errors.New("not yet supported ;)")
}
