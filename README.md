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

adds a new function to the server and stores it in the database. This call doesn't executes and validates the fucntion. It returns an error if the function already exists for the given botId.

#### Sample call

https://localhost:8080/addFunction

#### JSON Body:

``{
    "name" : "<name of function>",
    "botId" : "<the botId>",
    "code" : "<the javascript code>",
    "version" : 0,
    "appId" : "<the given appId or accesstoken which allows you to add functions on the server>"
}``
  
#### Sample JSON BOdy
  
 #### Explanation of the values:
 
 **name**
  the unique name of the function corresponding to the botid. 
  
 **botId**
  The id of the bot which needs this function. This service was introduce as simple functionserver  our max chatbot system. So the bot developer can easily add custom functions to bots. So a botId is always mandantory. If you want to use this service for other purposes. set the botId to 0. 

  **code**
  The javascript code of the function. Which code is supported will be described in the supported javascript section.
  
 **version**
  can always be 0. At the moment this has no effect. In the future we want to add version control and then we need this.
  
  **appId**
  the service provider will support you with an appId. Only with an valid appId you are able to add functions to the server. 
  
## Used frameworks

- restservice was implemented by the gin framework (github.com/gin-gonic/gin)
- javascript will be executed over the otto js framework (github.com/robertkrimen/otto)
  
  ## Supported javascript

## Further plans
-To the service we will implement a commandline tool to test the javascript code locally. also we want to make it possible to push data to the server over the commandline tool.
- Beside jaavscript we want to execute webassembly at the server.
- A further step is to make it possible, that the user can load and save json documents in a table to store data over the functions.


Here is still work in progress don't use it at the moment in prodution.
