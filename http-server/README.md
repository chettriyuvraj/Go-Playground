# README

We are doing a few things related to http-servers here


- Creating HTTP servers using the default http mux.

- Creating HTTP servers using a user-defined mux.

- Testing HTTP servers with small:
    - Unit Tests: Testing the handlers using http ResponseRecorders and mock requests via httptest.
    - Integration Tests: Testing the server itself using a mock server and passing to it the mux and hitting it with actual requests. 

- Exploring the URL package:
    - Important funcs
    - URL struct and important fields/methods
    - Values struct i.e. Query string in map form and important fields/methods
    - Also exists: UserInfo (simple enough), similar to Values (not a map), a struct specifically for username password

- Create a small 'streaming' server using io.Pipe()