package manager

import (
	"fmt"
	"log"

	"golang.org/x/exp/maps"

	"wogra.com/executor"
)

type FunctionManagerMemory struct {
	Functions map[string]*FunctionDescription
}

func NewMemory() *FunctionManagerMemory {

	var fm = new(FunctionManagerMemory)
	fm.Functions = make(map[string]*FunctionDescription)
	return fm
}

func (fm FunctionManagerMemory) generateFunctionIdentifier(name string, botId string) string {
	return botId + "_" + name
}

func (fm FunctionManagerMemory) findFunction(name string, botId string) *FunctionDescription {

	identifier := fm.generateFunctionIdentifier(name, botId)
	return fm.Functions[identifier]
}

func (fm FunctionManagerMemory) FindFunction(fd *FunctionDescription) *FunctionDescription {

	if fm.validateAppId(fd.AppId) {
		return fm.findFunction(fd.Name, fd.BotId)
	}

	return nil
}

func (fm FunctionManagerMemory) AddFunction(fd *FunctionDescription) *FunctionDescription {

	if fm.validateAppId(fd.AppId) {
		identifier := fm.generateFunctionIdentifier(fd.Name, fd.BotId)

		if fm.Functions[identifier] == nil {
			var newFunction = new(FunctionDescription)
			newFunction.Name = fd.Name
			newFunction.BotId = fd.BotId
			newFunction.Code = fd.Code
			fm.Functions[identifier] = newFunction
			return newFunction
		} else {
			log.Printf("Function %q already exists for Bot %q", fd.Name, fd.BotId)
		}
	}

	return nil
}

func (fm FunctionManagerMemory) validateAppId(appId string) bool {
	return HasAccessToken(appId)
}

func (fm FunctionManagerMemory) ExecuteFunction(call *FunctionCall) (string, error) {

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

func (fm FunctionManagerMemory) UpdateFunction(fd *FunctionDescription) bool {

	if fm.validateAppId(fd.AppId) {
		fdFund := fm.FindFunction(fd)

		if fd != nil {
			fdFund.Code = fd.Code
			return true
		}
	}

	return false
}
func (fm FunctionManagerMemory) DeleteFunction(fd *FunctionDescription) bool {

	if fm.validateAppId(fd.AppId) {
		identifier := fm.generateFunctionIdentifier(fd.Name, fd.BotId)
		fd = fm.FindFunction(fd)

		if fd != nil {
			delete(fm.Functions, identifier)
			return true
		}
	}

	return false
}

func (fm FunctionManagerMemory) GetFunctionList(botId string, appId string) []*FunctionDescription {
	if fm.validateAppId(appId) {
		return maps.Values(fm.Functions)
	}

	return nil
}
