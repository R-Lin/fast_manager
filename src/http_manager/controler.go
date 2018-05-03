package http_manager

import (
    "../mongo_manager"
    "os"
    "gopkg.in/mgo.v2/bson"
    "net/http"
    "log"
)

var MongoDb = new(mongo_manager.ProjectConfig);
var logger *log.Logger

func init(){
    mongoUrl, dbName := ReadConf("conf/main.conf")
    mongo_manager.Init(MongoDb, mongoUrl, dbName)
    logger = log.New(os.Stdout, "[INFO] ", log.LstdFlags | log.Lshortfile)
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
    logger.Println(r.Method, r.URL, r.Proto, r.ContentLength, r.Host, r.PostForm)
    logger.Println(13, r.PostForm)
    w.Write([]byte("sss"))
}

// 代码库查看
func CodeList(w http.ResponseWriter, r *http.Request){
    logger.Println(r.Method, r.URL, r.Proto, r.ContentLength, r.Host)
    result := MongoDb.Find(MongoDb.CodeCol, bson.M{})
    logger.Println(result)
    w.Write([]byte("Hello"))
}

// 增加代码库
func CodeAdd(w http.ResponseWriter, r *http.Request){
    if r.Method != "POST"{
        http.Error(w, "Method Not Allow", 400)
    }else{
        codeInfo := GetPostData(r)
        insertRecord := bson.M{}
        for _, keyName := range []string{"codename", "codepath"}{
            if value, ok := codeInfo[keyName];ok{
                insertRecord[keyName] = value.(string)
            }
        
        }
        err := MongoDb.CodeCol.Insert(insertRecord)
        if err != nil{
            http.Error(w, "Add Failed: " + err.Error(), 500)
        }else{
            logger.Println(insertRecord)
            w.Write([]byte("Add Record successfully"))
        
        }
    }

}

// 生成提测邮件
func MakingTestEmail(w http.ResponseWriter, r *http.Request){
    w.Write([]byte("Hello"))
}

// 查看和配置邮件列表
func ConfigEmailList(w http.ResponseWriter, r *http.Request){
    w.Write([]byte("Hello"))
}
