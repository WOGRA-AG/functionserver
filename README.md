# Functionserver
The Functionserver executes jaavscript functions on serverside. The javascript code can be added over a rest server to the server. The javascript snippets will be stored in a database table (currently mysql). over the another rest call the javascript can be executed. 

To make live easier we added two simple http request function to javascript (httpget and httppost) which executes the http commands. there are some config files which must be edited to make the service available. they all are stored in the service directory.

## Used frameworks

- restservice was implemented by the gin framework (github.com/gin-gonic/gin)
- javascript will be executed over the otto js framework (github.com/robertkrimen/otto)

## Further plans
-To the service we will implement a commandline tool to test the javascript code locally. also we want to make it possible to push data to the server over the commandline tool.
- Beside jaavscript we want to execute webassembly at the server.
- A further step is to make it possible, that the user can load and save json documents in a table to store data over the functions.


Here is still work in progress don't use it at the moment in prodution.
