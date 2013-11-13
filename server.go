package main

import(
    "fmt"
    "net/http"
)

func root(w http.ResponseWriter, req *http.Request){
    http.ServeFile(w, req, "./views/index.html")
}

func items(w http.ResponseWriter, req *http.Request){
    // API for CRUD
    fmt.Println("/items hit");
}

func serveFile(pattern string, filename string) {
    http.HandleFunc(pattern, func(w http.ResponseWriter, req *http.Request) {
        http.ServeFile(w, req, filename)
    })
}

func printOutput(handler http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
        fmt.Printf("%s %s %s\n", req.RemoteAddr, req.Method, req.URL)
        handler.ServeHTTP(w, req)
    })
}

func main(){
    fmt.Println(" > HTTP Server running...")

    http.HandleFunc("/static/", func(w http.ResponseWriter, req *http.Request) {
        http.ServeFile(w, req, req.URL.Path[1:])
    })

    http.HandleFunc("/items", items)

    http.HandleFunc("/", root)

    serveFile("/favicon.ico", "./favicon.ico")

    serveFile("/robots.txt", "./robots.txt")

    http.ListenAndServe(":1234", printOutput(http.DefaultServeMux))
}
