package main

import (
    "io"
    "log"
    "os"
    "net/http"
    "fmt"
)

func hello(rw http.ResponseWriter, req *http.Request){
    fmt.Println(req)
    io.WriteString(rw, "Fuck")
}

func main(){
    logger := log.New(os.Stderr, "", log.Lshortfile | log.Ldate | log.Ltime)
    http.HandleFunc("/", hello)
    logger.Println("服务启动, 监听地址: 0.0.0.0:999")
    http.ListenAndServe(":9999", nil)
}
