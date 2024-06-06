# README

## What 

A cleaner HTTP server with separation of concerns between

- Middleware
- Handlers i.e. business logic and route registration
- Any injections that might be required e.g. DB connections, loggers





## Packages

We will divide our app into the following packages

### Config

- An _AppConfig_ struct that contains all: 'injections' e.g. Loggers, Database Connections

- _App_ struct:
    
    <u>Fields</u>: 
    - _Config *AppConfig_:
    - _Handler func(w http.ResponseWriter, req *http.Request, config *AppConfig)_

    <u>Methods</u>: 
    - _ServeHTTP(w http.ResponseWriter, req *http.Request)_:
        - Satisfies _Handler_ interface.
        - _ServeHTTP_ will call our _Handler_ which is how we inject any requirements

- Note: Each path registered with our top level server's mux using _mux.Handle()_ will use an instance of _App_ as its _Handler_ parameter.

### Middleware

- Contains middlewares with signature of the form: func _LoggingMiddleware(handler http.Handler, config *app.AppConfig) http.HandlerFunc_

- Receive a _Handler_ interface type, and return a _HandleFunc_ concrete type function which first does something i.e. a middleware action like logging then calls the _serveHTTP_ of the _Handler_ it is wrapping.

### Handlers

- Defines the business logic for handling each route i.e. the actual handlers for each route.

- Register each path with an instance of _App_ and the business logic which is contained in the _Handler_ field of the _App_


## Tests

Each package (except _App_) contains a basic unit test that tests the component in isolation

Can also write an integration test that brings it all together in main_test.


## Self notes

- I think this kind of a structure separating concerns is similar across programming languages(?). I've done it with Javascript and a bit of Python and it was pretty similar.