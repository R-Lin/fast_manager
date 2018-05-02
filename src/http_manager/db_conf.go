package http_manager

import (
    "fmt"
    "io/ioutil"
    "strconv"
    "gopkg.in/yaml.v2"
)

func ReadConf(confFile string) (mongoUrl, dbName string){
    var c map[string]interface{}
    yamlFile, err := ioutil.ReadFile(confFile)
    if err != nil{
        fmt.Println("Fuck", err.Error())
    }
    err = yaml.Unmarshal(yamlFile, &c)
    if err != nil{
        fmt.Println("fuck", err.Error())
    }
    mongoMap := c["mongo"].(map[interface{}]interface{})
    mongoHost := mongoMap["host"].(string)
    mongoPort := strconv.Itoa(mongoMap["port"].(int))
    mongoDb := mongoMap["db"].(string)
    mongoUrl = mongoHost + ":" + mongoPort
    return mongoUrl, mongoDb
}
