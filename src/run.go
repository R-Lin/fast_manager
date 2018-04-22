package main

import (
    "./docker_manager"
    // "github.com/kr/pty"
    "io/ioutil"
    "time"
     "io"
     "os"
    // "os/exec"
    "fmt"
)

func main(){
    fmt.Println(docker_manager.ShowContainerList())
    fmt.Println("123")
    out := docker_manager.ShowContainerLogs(
    "5464f862262849a204b5681530730516e81a73a87c2d82d1ff4a4bf05b1c788d", "3")
    mesg, err := ioutil.ReadAll(out)
    fmt.Println(string(mesg), err)
    fmt.Println(time.Now())
    // c := exec.Command("/bin/bash")
    // f, err := pty.Start(c)
    // if err != nil{
    //     fmt.Println(err.Error())
    // }
    // f.Write([]byte("ls\n"))
    io.Copy(os.Stdout, out)
}
