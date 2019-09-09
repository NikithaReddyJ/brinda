package main
import "fmt"

 

type Vertex struct {
    X float32
    Y string
}
func main() {
    st := Vertex{X:1, Y:"niki"}
    st2 := Vertex{X:2, Y:"niki"}
    //st := Vertex{1, "niki"}
    fmt.Println(st, st2)
    fmt.Println(Vertex{X:1, Y:"niki"})
    
    ar := []Vertex{
        Vertex{1,"Go"},
        Vertex{2,"K8s"},
        }  
        fmt.Print(ar, len(ar))
}
