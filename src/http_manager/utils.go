package http_manager

import (
    "io/ioutil"
    "fmt"
    "encoding/json"
    "net/http"
)

func GetPostData(r *http.Request) map[string]interface{}{
    fmt.Println(r.Header.Get("Content-Type"))
    body, _ := ioutil.ReadAll(r.Body)
    fmt.Println(string(body))
    result := make(map[string]interface{})
    json.Unmarshal(body, &result)
    return result


}
