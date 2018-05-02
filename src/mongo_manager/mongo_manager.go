package mongo_manager

import (
    "fmt"
    "gopkg.in/mgo.v2"
    "gopkg.in/mgo.v2/bson"
)

type ProjectConfig struct{
    Session     *mgo.Session
    Db          *mgo.Database
    CodeCol     *mgo.Collection
    DockerCol   *mgo.Collection
    PmCol       *mgo.Collection
}

func (p ProjectConfig) Close(){
    p.Session.Close()
}

func Init(p *ProjectConfig, mongoUrl, dbName string){
    session, err := mgo.Dial(mongoUrl)
    if err != nil{
        fmt.Println(err.Error())
    }
    p.Db = session.DB(dbName)
    p.CodeCol = p.Db.C("code_info")
    p.DockerCol = p.Db.C("docker_info")
    p.PmCol = p.Db.C("pm_info")
}

func (p ProjectConfig) Remove(c *mgo.Collection, query *bson.M) (*mgo.ChangeInfo, error){
    return c.RemoveAll(query)
}

func (p ProjectConfig) Find(c *mgo.Collection, query bson.M) []bson.M{
    s := []bson.M{}
    fmt.Println( c.Find(query))
    c.Find(query).All(&s)
    return s
}
