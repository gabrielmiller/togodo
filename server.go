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

func main(){
    fmt.Println(" > HTTP Server running...")
    http.Handle("/static/", http.FileServer(http.Dir("./static/")))
    http.HandleFunc("/", root)
    http.HandleFunc("/item", item)
    err := http.ListenAndServe(":1234", nil)
    if err != nil {
        log.Fatal("ListenAndServe: ", err)
    }
}
