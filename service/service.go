package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
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
		router.GET("/ping", func(c *gin.Context) {
			c.String(http.StatusOK, "pong")
		})

		log.Println("Run router")

		restConfig := manager.ReadRestConfiguration()
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

	fds := functionManager.GetFunctionList(json["botId"], json["appId"])

	if fds != nil {
		c.IndentedJSON(http.StatusOK, fds)
		return
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"result": "list is empty"})
}
