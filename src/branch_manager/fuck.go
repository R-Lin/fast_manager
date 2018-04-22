package main
import (
    "fmt"
    "time"
)

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
    fmt.Println(time.Now())
    loc, _ := time.LoadLocation("Europe/Berlin")
    fmt.Println(time.Date(2013, 03, 06, 14, 30, 0, 0, loc))
}
