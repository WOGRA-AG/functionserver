package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"wogra.com/config"
	"wogra.com/manager"
)

var functionManager *manager.FunctionManagerDb

func initWebService() error {

	log.Println("Starting functionmanager service")
	functionManager = manager.NewDb()

	if functionManager != nil {

		router := gin.Default()
		log.Println("Adding routes")
		router.POST("/getFunctionList/", getFunctionList)
		router.POST("/getFunction/", findFunction)
		router.POST("/addFunction/", addFunction)
		router.POST("/deleteFunction/", deleteFunction)
		router.POST("/updateFunction/", updateFunction)
		router.POST("/executeFunction/", executeFunction)
		router.POST("/createAppId/", createAppId)
		router.POST("/deleteAppId/", deleteAppId)
		router.POST("/checkCredentials/", checkCredentials)
		router.POST("/addCredentials/", addCredentials)
		router.POST("/deleteCredentials/", deleteCredentials)
		router.GET("/ping", func(c *gin.Context) {
			c.String(http.StatusOK, "pong")
		})

		log.Println("Run router")

		restConfig := config.ReadRestConfiguration()
		connection := restConfig.Host + ":" + restConfig.Port

		router.Run(connection)
	} else {
		return fmt.Errorf("can not init functionmanager")
	}

	return nil
}

func findFunction(c *gin.Context) {
	var fd manager.FunctionDescription
	err := c.BindJSON(&fd)

	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"result": "json invalid"})
	}

	fdFound := functionManager.FindFunction(&fd)

	if fdFound != nil {
		c.IndentedJSON(http.StatusOK, fdFound)
		return
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"result": "function not found"})
}

func addFunction(c *gin.Context) {
	var fd manager.FunctionDescription
	err := c.BindJSON(&fd)

	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"result": "json invalid"})
	}

	createdFunction := functionManager.AddFunction(&fd)

	if createdFunction != nil {
		c.IndentedJSON(http.StatusCreated, createdFunction)
	} else {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"result": "internal server error"})
	}
}

func deleteFunction(c *gin.Context) {
	var fd manager.FunctionDescription
	err := c.BindJSON(&fd)

	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"result": "json invalid"})
	}

	if functionManager.DeleteFunction(&fd) {
		c.IndentedJSON(http.StatusOK, gin.H{"result": "function deleted"})
	} else {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"result": "internal server error"})
	}
}

func updateFunction(c *gin.Context) {
	var fd manager.FunctionDescription
	err := c.BindJSON(&fd)

	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"result": "json invalid"})
	}

	if functionManager.UpdateFunction(&fd) {
		c.IndentedJSON(http.StatusOK, gin.H{"result": "function updated"})
	} else {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"result": "internal server error"})
	}
}

func executeFunction(c *gin.Context) {
	var fc manager.FunctionCall
	err := c.BindJSON(&fc)

	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"result": "json invalid"})
	}

	result, err := functionManager.ExecuteFunction(&fc)

	if err == nil {
		c.IndentedJSON(http.StatusOK, gin.H{"result": result})
	} else {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"result": err.Error})
	}
}

func getFunctionList(c *gin.Context) {
	json := make(map[string]string) // note the content accepted by the structure
	c.BindJSON(&json)

	fds := functionManager.GetFunctionList(json["appId"], json["botId"])

	if fds != nil {
		c.IndentedJSON(http.StatusOK, fds)
		return
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"result": "list is empty"})
}

func createAppId(c *gin.Context) {
	json := make(map[string]string) // note the content accepted by the structure
	c.BindJSON(&json)

	fds := functionManager.CreateAppId(json["superUserAppId"], json["owner"], json["contact"])

	if fds != "" {
		c.IndentedJSON(http.StatusOK, fds)
		return
	} else {
		c.IndentedJSON(http.StatusBadRequest, "create AppId failed")
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"result": "list is empty"})
}

func deleteAppId(c *gin.Context) {
	json := make(map[string]string) // note the content accepted by the structure
	c.BindJSON(&json)

	fds := functionManager.DeleteAppId(json["superUserAppId"], json["appId"])

	if fds == true {
		c.IndentedJSON(http.StatusOK, "AppId deleted")
		return
	} else {
		c.IndentedJSON(http.StatusBadRequest, "deletion of AppId failed")
		return
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"result": "list is empty"})
}

func checkCredentials(c *gin.Context) {
	json := make(map[string]string) // note the content accepted by the structure
	c.BindJSON(&json)

	fds := functionManager.CheckCredentials(json["appId"], json["botId"])

	if fds == true {
		c.IndentedJSON(http.StatusOK, "Access granted")
		return
	} else {
		c.IndentedJSON(http.StatusForbidden, "Access failed")
		return
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"result": "list is empty"})
}

func addCredentials(c *gin.Context) {
	json := make(map[string]string) // note the content accepted by the structure
	c.BindJSON(&json)

	fds := functionManager.AddCredentials(json["superUserAppId"], json["appId"], json["botId"])

	if fds == true {
		c.IndentedJSON(http.StatusOK, "Access granted")
		return
	} else {
		c.IndentedJSON(http.StatusBadRequest, "Access failed")
		return
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"result": "list is empty"})
}

func deleteCredentials(c *gin.Context) {
	json := make(map[string]string) // note the content accepted by the structure
	c.BindJSON(&json)

	fds := functionManager.DeleteCredentials(json["superUserAppId"], json["appId"], json["botId"])

	if fds == true {
		c.IndentedJSON(http.StatusOK, "Access deleted")
		return
	} else {
		c.IndentedJSON(http.StatusBadRequest, "Access failed")
		return
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"result": "list is empty"})
}
