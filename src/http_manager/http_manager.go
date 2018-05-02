package http_manager

import (
    "log"
    "net/http"
)

type httpHandlerFunc func(http.ResponseWriter, *http.Request)
var UrlRoute map[string] httpHandlerFunc = map[string]httpHandlerFunc{
    "/": TaskList,
    "/tasklist": TaskList,
    "/taskadd": TaskAdd,
    "/dockerlist": DockerList,
    "/dockeradd": DockerAdd,
    "/codelist": CodeList,
    "/codeadd": CodeAdd,
    "/makingtestemail": MakingTestEmail,
    "/configemaillist": ConfigEmailList,
}

func InitServer(addr string){
    for url, handler := range UrlRoute{
        http.HandleFunc(url, handler)
    }
    log.Printf("Server start, Listen on %s\n", addr)
    http.ListenAndServe(addr, nil)
}
