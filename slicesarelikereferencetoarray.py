package main
import "fmt"

 

func main() {
    names := [4]string{
        "John",
        "Paul",
        "George",
        "Ringo",
    }
    fmt.Println(names)

 

    a := names[0:3]  //index 0 to index (2-1)
    b := names[1:4]  //index 1 to index (3-1)
    fmt.Println(a, b)

 

    b[0], b[1] = "XXX", "YYY"
    fmt.Println(a, b)
    fmt.Println(names)
}
