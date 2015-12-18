package main
    
import (
    "fmt"
    "io/ioutil"
)

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func main() {
    data, err := ioutil.ReadFile("input.txt")
    check(err)
    
    var floor int = 0
    var floor_position int = 1

    for i := 0; i < len(data); i++ {
        if data[i] == '(' {
            floor++
        } else if data[i] == ')' {
            floor--
        }

        if floor < 0 {
            break
        }

        floor_position++
    }

    fmt.Printf("%d\n", floor_position)
}