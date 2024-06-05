# README

## What 

A 'cleaner' HTTP-server with separation of concerns between

- Business logic
- Routing
- Middleware


## Packages

We will divide our app into the following packages

### App
-  An _App_ struct that would handle all 'injections' e.g. Loggers, Database connections
- This would be the top-level _Handler_ to our server
- Struct would contain:
    - Anything that has to be injected as the fields
    - A field called handler which would be an instance of a func with params same as handleFunc(RespWrit, Reqest, ...Logger, DB connection etc.)
    - ServeHTTP func which would make it satisfy the _Handler_ interface
- Each path registration would carry an instance of this _App_

### Middleware
- This contains functions e.g. _LoggingMiddleWare_ which receive a _Handler_ interface type, and return a _HandleFunc_ concrete type function which first does something then calls the _serveHTTP_ of the thing that it contains

### Handlers

- Defines the business logic for handling each route

- Also registers each path with an instance of _App_ and the correct handler