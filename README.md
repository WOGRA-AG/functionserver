# WOGRA Function Server
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

### /addFunction (Post)

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
  
  
### /getFunction (Post)

returns the function data of the given function in database. If the function wasn't found an 404 error code will be returned.

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
  
  ### /updateFunction (Post)

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
  
### /deleteFunction (Post)

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
  
  
### /executeFunction (Post)

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
  
  **params**
  The params which will be injected in the runtime. In the add example above the code in the script engine is return parseInt(p1) + parseInt(p2), because every parameter will be added as string to the script engine. the result will always be returned as string as well.

### /getFunctionList (Post)

returns all functions of the given botId.

#### Sample call

https://localhost:8080/getFunctionList

#### JSON Body:
``{
    "botId": "<botId of Function>",
    "appId": "<provided appId>"
}``

#### Sample JSON Body
``{
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
},
{
    "name": "test23",
    "botId": "1",
    "code": "2+2;",
    "version": 0,
    "appId": "jakfhakjdfhurfueinakn76283vdjkdksvhkjd"
}``

#### Explanation of the JSON input values:
  
 **botId**
  The id of the bot which needs this function. This service was introduce as simple functionserver  our max chatbot system. So the bot developer can easily add custom functions to bots. So a botId is always mandantory. If you want to use this service for other purposes. set the botId to 0. 
  
  **appId**
  the service provider will support you with an appId. Only with an valid appId you are able to add functions to the server. 

### /ping (Get)

returns pong with http status 200. Is only to check if the service is started.

#### Sample call

``https://localhost:8080/ping``
  

## Supported javascript

The javascript engine of otto supports ECMA Script 5. We added further function which enhances the script engine. The code which should be executed will be loaded from database. The code will be surrounded by a function head and footer and a function call. The function should always return something.

**Example:**

``function wrapper() {
<code from database>
}
wrapper();``
    
To prevent endless loops in function code, the execution of a function has a maximum duration of 10 seconds. If a function doesn't finish in this time, the execution will be cancelled.

### Enhanced HTTP Requests
We added httpGet and httpPost request to the javascript engine which returns the result as string. it is very simple to call this requests.

#### httpGet(address)

returns the result of a get call to the given address.

##### Example

``httpGet("https://wogra.com")``

**address**
the server address whcih should be called like https://localhost:8080

#### httpPost(address, type, body)


##### Example

``httpPost("https://wogra.com/dosomehting", "application/json", "{\"a\" : \"b\"}")``

returns the result of a post call to the given address.

**address**
the server address whcih should be called like https://localhost:8080

**type**
the type of the body for example application/json

**body**
the body which will be sent via post to the server.

### Integrated key value store
To load and save data we integrated a key value store for storing and loading data. The data will be stored regarding to the botid. so the key must be unique with the botid. the value will be stored as long text. So you are able to store json documents or any other strings. We offer three function for the key value store:
- loadDocument(string key)
- saveDocument(string key, string content)
- deleteDcoument(string key)
    
#### string loadDocument(string key)
loads the document with given key and returns the content as string.
    
#### bool saveDocument(string key, string content)
 saves the document conent with the given key. returns true if storing was successfull. if the key already exists, the content will be overridden by the function.
    
#### bool deleteDocument(string key)
deletes the document with the given key.
  
## Used frameworks

- restservice was implemented by the gin framework (github.com/gin-gonic/gin)
- javascript will be executed over the otto js framework (github.com/robertkrimen/otto)  

## Further plans
see issue list for further plans.
