# README

Playing around with http servers in Go.

## What

### Mux
- Creating HTTP servers using the default http mux.

- Creating HTTP servers using a user-defined mux.

### Tests
 Testing HTTP servers using

- Unit Tests: Testing the handlers using _http ResponseRecorders_ and mock requests via _httptest_.

- Integration Tests: Testing the server itself using a mock server, passing to it the mux and hitting it with actual requests. 


### URL

Exploring the URL package
    
- Some important functions
    
- URL struct and important fields/methods
    
### Streaming

Create a small 'streaming' server using _io.Pipe()_

You can check it out by

- Running the function in _main_

- Since this request is streaming:
    - First establish a TCP connection using _nc localhost 8082_
    - Send an HTTP request in the format:
    ```
    GET /stream HTTP/1.1
    Host: localhost:8082
    <CR>
    <CR>
    ```

- You'll get a response as _Chunked_ encoding i.e. stream size followed by stream data.
