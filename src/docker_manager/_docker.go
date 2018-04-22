package main

import (
    "fmt"
    "github.com/docker/docker/api/types"
    "github.com/docker/docker/client"
    "golang.org/x/net/context"
)

func ListDockers(c *client.Client, ctx context.Context)([]types.Container){
    containers, err := c.ContainerList(
        ctx, types.ContainerListOptions{})

    if err != nil{
        panic(err.Error())
    }
    return containers
}

func StopContainer(c *client.Client, ctxContext context.Context, containerID string){

    err := c.ContainerStop(ctxContext, containerID, nil)
    if err != nil{
        fmt.Println(err.Error())
    }
}
func main(){
    ctx := context.Background()
    cli, err := client.NewEnvClient()
    if err != nil{
        panic(err)
    }
    var sID  string
   for _, container := range ListDockers(cli, ctx){
       fmt.Println(container.ImageID, container.Names[0])
       sID = container.ID
   }
   StopContainer(cli, ctx, sID)
   fmt.Println(sID)
}
