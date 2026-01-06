# Creating Web Servers using Go.

* You would need 3 core items :
    * Handlers : Think of them like controllers in MVC architecture patterns, they process incoming HTTP requests and generate HTTP responses.

    * Routers : These are used to map URL patterns with their corresponding hanlders, http.ServeMux is one implementation of a router.

    * WebServer : We can create a webserver using http.ListenAndServe() in go, and we would need to pass 2 arguments
        * TCP-Address : We generally follow the format of "host:port", if you omit the lost, like we did in this example the server will listen to available network interfaces on the computer.

        * Router : This is just the server mux we created earlier on, in the file.

* Handlers have two parameters :
    * http.ResponseWriter : This has methods that allow you to write http response and send it to the user

    * *http.Request : This is a pointer to an http request and contains information about like request like http method, url etc.