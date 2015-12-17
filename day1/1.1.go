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

    for i := 0; i < len(data); i++ {
        if data[i] == '(' {
		    floor++
        } else if data[i] == ')' {
            floor--
        }
    }

    fmt.Printf("%d\n", floor)
}