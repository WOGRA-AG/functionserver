# Functionserver
The Functionserver executes jaavscript functions on serverside. The javascript code can be added over a rest server to the server. The javascript snippets will be stored in a database table (currently mysql). over the another rest call the javascript can be executed. 

To make live easier we added two simple http request function to javascript (httpget and httppost) which executes the http commands. there are some config files which must be edited to make the service available. they all are stored in the service directory.

## REST API Descirption

Except of the ping call, all requests are post requests with json bodies. There are 7 different requests:

-  /addFunction (post)
-  /getFunction (post)
-  /updateFunction (post)
-  /deleteFunction (post)
-  /executeFunction (post)
-  /getFunctionList (post)
-  /ping (get)

All calls will be described here:

### /addFunction

adds a new function to the server and stores it in the database. This call doesn't executes and validates the fucntion. It returns an error if the function already exists for the given botId. If the function was successfully added to the database. the function returns the stored function object with http 200.

#### Sample call

https://localhost:8080/addFunction

#### JSON Body:

``{
    "name" : "<name of function>",  
    "botId" : "<the botId>",  
    "code" : "<the javascript code>",  
    "appId" : "<the given appId or accesstoken which allows you to add functions on the server>"  
}``
  
#### Sample JSON Body
`` {  
    "name" : "test",  
    "botId" : "1",  
    "code" : "2+2;",  
    "appId" : "jakfhakjdfhurfueinakn76283vdjkdksvhkjd"  
}``

#### Sample result

``{
    "name": "test",
    "botId": "1",
    "code": "2+2;",
    "version": 0,
    "appId": "jakfhakjdfhurfueinakn76283vdjkdksvhkjd"
}``

#### Explanation of the JSON input values:
 
 **name**
  the unique name of the function corresponding to the botid. 
  
 **botId**
  The id of the bot which needs this function. This service was introduce as simple functionserver  our max chatbot system. So the bot developer can easily add custom functions to bots. So a botId is always mandantory. If you want to use this service for other purposes. set the botId to 0. 

  **code**
  The javascript code of the function. Which code is supported will be described in the supported javascript section.
  
  **appId**
  the service provider will support you with an appId. Only with an valid appId you are able to add functions to the server. 
  
  
### /getFunction

returns nthe function data of the given function in database. If the function wasn't found an 404 error code will be returned.

#### Sample call

https://localhost:8080/getFunction

#### JSON Body:
``{
    "name": "<name of function>",
    "botId": "<botId of Function>",
    "appId": "<provided appId>"
}``

#### Sample JSON Body
``{
    "name": "test",
    "botId": "1",
    "appId": "jakfhakjdfhurfueinakn76283vdjkdksvhkjd"
}``
    
#### Sample result

``{
    "name": "test",
    "botId": "1",
    "code": "2+2;",
    "version": 0,
    "appId": "jakfhakjdfhurfueinakn76283vdjkdksvhkjd"
}``

#### Explanation of the JSON input values:
 
 **name**
  the unique name of the function corresponding to the botid. 
  
 **botId**
  The id of the bot which needs this function. This service was introduce as simple functionserver  our max chatbot system. So the bot developer can easily add custom functions to bots. So a botId is always mandantory. If you want to use this service for other purposes. set the botId to 0. 
  
  **appId**
  the service provider will support you with an appId. Only with an valid appId you are able to add functions to the server. 
  
  ### /updateFunction

updates a given function on the functionserver. The code will be replaced in the database. If successfull, the function returns the stored function object with http 200.

#### Sample call

https://localhost:8080/updateFunction

#### JSON Body:

``{
    "name" : "<name of function>",  
    "botId" : "<the botId>",  
    "code" : "<the javascript code>",  
    "appId" : "<the given appId or accesstoken which allows you to add functions on the server>"  
}``
  
#### Sample JSON Body
`` {  
    "name" : "test",  
    "botId" : "1",  
    "code" : "5+5;",  
    "appId" : "jakfhakjdfhurfueinakn76283vdjkdksvhkjd"  
}``

#### Sample result

``{
    "name": "test",
    "botId": "1",
    "code": "2+2;",
    "version": 0,
    "appId": "jakfhakjdfhurfueinakn76283vdjkdksvhkjd"
}``

#### Explanation of the JSON input values:
 
 **name**
  the unique name of the function corresponding to the botid. 
  
 **botId**
  The id of the bot which needs this function. This service was introduce as simple functionserver  our max chatbot system. So the bot developer can easily add custom functions to bots. So a botId is always mandantory. If you want to use this service for other purposes. set the botId to 0. 

  **code**
  The javascript code of the function. Which code is supported will be described in the supported javascript section.
  
  **appId**
  the service provider will support you with an appId. Only with an valid appId you are able to add functions to the server. 
  
### /deleteFunction

delets the given function in database. After deletion execution is not possible anymore. If successfull code 200 will be returned.

#### Sample call

https://localhost:8080/deleteFunction

#### JSON Body:
``{
    "name": "<name of function>",
    "botId": "<botId of Function>",
    "appId": "<provided appId>"
}``

#### Sample JSON Body
``{
    "name": "test",
    "botId": "1",
    "appId": "jakfhakjdfhurfueinakn76283vdjkdksvhkjd"
}``
    
#### Sample result
{
    "result": "function deleted"
}


#### Explanation of the JSON input values:
 
 **name**
  the unique name of the function corresponding to the botid. 
  
 **botId**
  The id of the bot which needs this function. This service was introduce as simple functionserver  our max chatbot system. So the bot developer can easily add custom functions to bots. So a botId is always mandantory. If you want to use this service for other purposes. set the botId to 0. 
  
  **appId**
  the service provider will support you with an appId. Only with an valid appId you are able to add functions to the server. 
  
  
### /executeFunction

delets the given function in database. After deletion execution is not possible anymore. If successfull code 200 will be returned.

#### Sample call

https://localhost:8080/executeFunction

#### JSON Body:
``{
    "name": "<name of function>",
    "botId": "<botId of Function>",
    "appId": "<provided appId>"
    "params": "<paramter name and value map>"
}``

#### Sample JSON Body
``{
    "name": "add",
    "botId": "1",
    "appId": "1",
    "params": {"p1" : "4", "p2" : "6"}
}``
    
#### Sample result
{
    "result": "10"
}


#### Explanation of the JSON input values:
 
 **name**
  the unique name of the function corresponding to the botid. 
  
 **botId**
  The id of the bot which needs this function. This service was introduce as simple functionserver  our max chatbot system. So the bot developer can easily add custom functions to bots. So a botId is always mandantory. If you want to use this service for other purposes. set the botId to 0. 
  
  **appId**
  the service provider will support you with an appId. Only with an valid appId you are able to add functions to the server.   
  
  
  
  

## Supported javascript

The javascript engine of otto supports ECMA Script 5. We added further function which enhances the script engine.

### Enhanced HTTP Requests
We added httpGet and httpPost request to the javascript engine which returns the result as string. it is very simple to call this requests.

#### httpGet(address)

returns the result of a get call to the given address.

##### Example

httpGet("https://wogra.com")

**address**
the server address whcih should be called like https://localhost:8080

#### httpPost(address, type, body)


##### Example

httpPost("https://wogra.com/dosomehting", "application/json", "{\"a\" : \"b\"}")

returns the result of a post call to the given address.

**address**
the server address whcih should be called like https://localhost:8080

**type**
the type of the body for example application/json

**body**
the body which will be sent via post to the server.

  
## Used frameworks

- restservice was implemented by the gin framework (github.com/gin-gonic/gin)
- javascript will be executed over the otto js framework (github.com/robertkrimen/otto)  

## Further plans
-To the service we will implement a commandline tool to test the javascript code locally. also we want to make it possible to push data to the server over the commandline tool.
- Beside jaavscript we want to execute webassembly at the server.
- A further step is to make it possible, that the user can load and save json documents in a table to store data over the functions.
