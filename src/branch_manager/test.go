package main

import (
    "fmt"
    "github.com/kr/pty"
    "io"
    "os"
    "os/exec"
)

var PATH string = "/Users/lin_r/Desktop/company_pro/detect/detect-ui"
var BASH, err = pty.Start(exec.Command("bash"))

func ExitBash(){
    BASH.Write([]byte("exit\n"))
}

func ShowBranch(path string)(string){
    BASH.Write([]byte("cd " + path + "\n"))
    BASH.Write([]byte("git branch\n"))
    fmt.Println(123)
    io.Copy(os.Stdout, BASH)
    return "123"
}

func main(){
    ShowBranch(PATH)
    defer ExitBash()
}
