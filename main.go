package main

import (
    "fmt"
    "html"
    "log"
    "net/http"
    "database/sql"
    _ "github.com/go-sql-driver/mysql"
)

func main() {

    fmt.Println("Simple App")
    db, err := sql.Open("mysql", "<username>:<password>@tcp(database_server_url:port_number)")
    if err != nil {
        panic(err.Error())
    }
    defer db.Close()
    
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
    })
    log.Fatal(http.ListenAndServe(":8081", nil))
}
