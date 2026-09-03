package main
//http handler is part of your go app that receives http requests and sends http responses. its responsible for handling and processing incoming requests, executing the appropriate logic,.. its the intermediary btn the client and the server. it can also handle errors, validate input, and manage flow of data. it can be used to implement routing, authentication, and other middleware functionalities. in go httphandlers are typically implemented using the net/http package, which provides a simple and efficient way to create web servers and handle http requests. you can define your own handler functions or use the built-in http.Handler interface to create custom handlers that respond to specific routes or endpoints. overall, http handlers play a crucial role in building web applications and services in go, enabling developers to create robust and scalable server-side solutions.
// w -> response write- we use it to send a response back to the client. allows us to write data to the response body, set headers, and control the status code of the response, its an instance of http.ResponseWriter, which is an interface provided by the net/http package in go. its passed as a parameter to the handler function, allowing us to write the response data that will be sent back to the client. we can use methods like w.Write() to write the response body, w.Header().Set() to set headers, and w.WriteHeader() to set the status code of the response. 
// r -> request contains in4 of what the client sent to the server. its an instance of http.Request, which is a struct provided by the net/http package in go. contains information about the incoming http request, e.g the request method (GET, POST, etc.), request headers, query parameters, request body, and other relevant details. we can access these properties to extract data from the request and perform necessary operations based on the client's input. the http.Request struct provides various fields and methods that allow us to retrieve information about the request, such as r.Method for the request method, r.URL for the requested URL, r.Header for accessing headers, and r.Body for reading the request body. by utilizing the information in the http.Request object, we can process the client's request and generate an appropriate response.
import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World!")
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil) // this starts your http server running on port 8080. 1st argument is the port nu to litsen to, the 2nd is the handler to use for incoming requests. if nil is provided, the default serve mux is used. using
}

//HTTP STaTUS CODES: no that tells the client the result of its request
// 200 OK: request was succesfully created and the server has returned the requested data.
// 201 Created: request successfully created a new resource on the server.
// 400 Bad Request: request invalid
// 401 Unauthorized: request requires authentication or the provided credentials are invalid.
// 403 Forbidden: server understood the request but refuses to authorize it.
// 404 Not Found: requested resource could not be found on the server.
// 500 Internal server error: server encountered an unexpected condition/problem that prevented it from fulfilling the request.
// 502 Bad Gateway: server received an invalid response from an upstream server while acting as a gateway proxy.
