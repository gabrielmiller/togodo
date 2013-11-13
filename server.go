package main

import(
    "fmt"
    "net/http"
    "html/template"
)


func root(w http.ResponseWriter, req *http.Request){
    indexTemplate := template.Must(template.New("index").ParseFiles("./templates/index.html"))
    indexTemplate.ExecuteTemplate(w, "base", "index")
}

func item(w http.ResponseWriter, req *http.Request){
    // API for CRUD
}

func serveFile(pattern string, filename string) {
    http.HandleFunc(pattern, func(w http.ResponseWriter, r *http.Request) {
        http.ServeFile(w, r, filename)
    })
}

func printOutput(handler http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        fmt.Printf("%s %s %s\n", r.RemoteAddr, r.Method, r.URL)
    handler.ServeHTTP(w, r)
    })
}

func main(){
    fmt.Println(" > HTTP Server running...")

    http.HandleFunc("/static/", func(w http.ResponseWriter, r *http.Request) {
        http.ServeFile(w, r, r.URL.Path[1:])
    })

    http.HandleFunc("/", root)
    //http.HandleFunc("/item", item)
    //http.Handle("/static/", http.FileServer(http.Dir("./static")))

    serveFile("/favicon.ico", "./favicon.ico")
    serveFile("/robots.txt", "./robots.txt")

    http.ListenAndServe(":1234", printOutput(http.DefaultServeMux))
}
