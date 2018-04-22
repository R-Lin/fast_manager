package tty
import (
    "github.com/kr/pty"
    "io"
    "os"
    "os/exec"
)

func main(){
    c := exec.Command("/bin/ls", "-l", "tty.go")
    f, err := pty.Start(c)
    if err != nil{
        panic(err)
    }
    io.Copy(os.Stdout, f)
}
