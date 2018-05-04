package docker_manager

import (
    "fmt"
    "io"
    "github.com/docker/docker/client"
    "golang.org/x/net/context"
    "github.com/docker/docker/api/types"
)

var CLIENT, err = client.NewEnvClient()
var CTX context.Context = context.Background()

// 每个容器的结构
type ContainerType struct{
    ID          string      // 用来操作容器
    Names       []string
    Created     int64
    Ports       []types.Port
    State       string
    Status      string
}

func ShowContainerList()[]map[string]ContainerType{
    containers, err := CLIENT.ContainerList(CTX, types.ContainerListOptions{})
    if err != nil{
        fmt.Println(err.Error())
    }
    formateContains := make([]map[string]ContainerType, 0)
    for _, obj := range containers {
        container := ContainerType{
            obj.ID,
            obj.Names,
            obj.Created,
            obj.Ports,
            obj.State,
            obj.Status,
        }
        formateContains = append(formateContains, map[string]ContainerType{
            obj.ID: container,
        })
    }
    return formateContains
}

func StartContainer(id string){
    err := CLIENT.ContainerStart(CTX, id, types.ContainerStartOptions{})
    if err != nil{
        fmt.Println(err.Error())
    }
}
func StopContainer(id string){
    err := CLIENT.ContainerStop(CTX, id, nil)
    if err != nil{
        fmt.Println(err.Error())
    }
}

func ShowContainerLogs(id, tailNum string)(io.ReadCloser){
   options := types.ContainerLogsOptions{
       ShowStdout: true,
       Timestamps: true,
       Tail: tailNum}

   out, err := CLIENT.ContainerLogs(CTX, id, options)
   if err != nil{
    fmt.Println(err.Error())
   }
   return out
}

func main(){
    s := ShowContainerList()
    for i, container := range s{
        fmt.Println(i, container)
    }
}
