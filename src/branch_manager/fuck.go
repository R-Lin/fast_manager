package main
import "fmt"

type student struct{
    num int
    name string
}
func main(){
    s := student{4, "s"}
    fmt.Println(s)
    for i, v := range [3]int{4, 4, 5}{
        fmt.Println(i, v)
    }
}
