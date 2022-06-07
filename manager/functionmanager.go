package manager

type FunctionDescription struct {
	Name    string `json:"name"`
	BotId   string `json:"botId"`
	Code    string `json:"code"`
	Version int    `json:"version"`
	AppId   string `json:"appId"`
}

type FunctionCall struct {
	Name   string            `json:"name"`
	BotId  string            `json:"botId"`
	AppId  string            `json:"appId"`
	Params map[string]string `json:"params"`
}

type FunctionManager interface {
	FindFunction(fd *FunctionDescription) *FunctionDescription
	AddFunction(fd *FunctionDescription) *FunctionDescription
	ExecuteFunction(call *FunctionCall) (string, error)
	UpdateFunction(fd *FunctionDescription) bool
	DeleteFunction(fd *FunctionDescription) bool
	GetFunctionList(botId string, appId string) []*FunctionDescription
}
