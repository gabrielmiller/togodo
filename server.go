package main

import(
    "fmt"
    "net/http"
    "html/template"
    "log"
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

func Log(handler http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        log.Printf("%s %s %s", r.RemoteAddr, r.Method, r.URL)
    handler.ServeHTTP(w, r)
    })
}

func main(){
    fmt.Println(" > HTTP Server running...")
    http.Handle("/static/", http.FileServer(http.Dir("./static/"))) // y u no werk
    http.HandleFunc("/", root)
    //http.HandleFunc("/item", item)
    //http.Handle("/static/", http.FileServer(http.Dir("./static")))

    serveFile("/sitemap.xml", "./sitemap.xml")
    serveFile("/favicon.ico", "./favicon.ico")
    serveFile("/robots.txt", "./robots.txt")

    http.ListenAndServe(":1234", Log(http.DefaultServeMux))
    //err := http.ListenAndServe(":1234", nil)
    //if err != nil {
    //    log.Fatal("ListenAndServe: ", err)
    //}
}
