package main

import (
    "fmt"
    "io/ioutil"
    "strconv"
    "gopkg.in/yaml.v2"
    "gopkg.in/mgo.v2/bson"
    "./mongo_manager"
)

var ProjectDb mongo_manager.ProjectConfig

func main(){
    var c map[string]interface{}
    yamlFile, err := ioutil.ReadFile("conf/main.conf")
    if err != nil{
        fmt.Println("Fuck", err.Error())
    }
    err = yaml.Unmarshal(yamlFile, &c)
    if err != nil{
        fmt.Println("fuck", err.Error())
    }
    mongoHost := c["mongo"].(string)
    mongoPort := strconv.Itoa(c["port"].(int))
    mongoDb := c["db"].(string)
    mongoUrl := mongoHost + ":" + mongoPort
    fmt.Println(mongoUrl)
    mongo_manager.Init(&ProjectDb, mongoUrl, mongoDb)
    ProjectDb.PmCol.Insert(&bson.M{
        "fuck": 3,
    })
    fmt.Print(ProjectDb.Find(ProjectDb.PmCol, &bson.M{}))
    ProjectDb.Remove(ProjectDb.PmCol, &bson.M{})
}
