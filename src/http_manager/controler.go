package http_manager

import (
    "../mongo_manager"
    "gopkg.in/mgo.v2/bson"
    "net/http"
    "log"
)

var MongoDb = new(mongo_manager.ProjectConfig);

func init(){
    mongoUrl, dbName := ReadConf("conf/main.conf")
    mongo_manager.Init(MongoDb, mongoUrl, dbName)
}

// 任务查看
func TaskList(w http.ResponseWriter, r *http.Request){
    w.Write([]byte("Hello"))
}

// 增加任务
func TaskAdd(w http.ResponseWriter, r *http.Request){
    w.Write([]byte("Hello"))
}

// 容器查看
func DockerList(w http.ResponseWriter, r *http.Request){
    w.Write([]byte("Hello"))
}

// 增加容器
func DockerAdd(w http.ResponseWriter, r *http.Request){
    w.Write([]byte("Hello"))
}

// 代码库查看
func CodeList(w http.ResponseWriter, r *http.Request){
    log.Println(r.Method, r.URL, r.Proto, r.ContentLength, r.Host)
    result := MongoDb.Find(MongoDb.CodeCol, bson.M{})
    log.Println(result)
    w.Write([]byte("Hello"))
}

// 增加代码库
func CodeAdd(w http.ResponseWriter, r *http.Request){
    log.Println(r)
    w.Write([]byte("Hello"))
}

// 生成提测邮件
func MakingTestEmail(w http.ResponseWriter, r *http.Request){
    w.Write([]byte("Hello"))
}

// 查看和配置邮件列表
func ConfigEmailList(w http.ResponseWriter, r *http.Request){
    w.Write([]byte("Hello"))
}
