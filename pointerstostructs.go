package main
import "fmt"

 

type Vertex struct {
    X int
    Y int
}

 

func main() {
    v := Vertex{1, 2}
    p := &v
    fmt.Println(v, p, p.X, (*p).X, p.Y, (*p).Y)

 

    (*p).X = 2e9 //2*10^9
    //p.X =12
    fmt.Println(v)
}
