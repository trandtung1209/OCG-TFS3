package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
    dat := []byte("1 3 5 7 9")
    err := ioutil.WriteFile("myfile.txt", dat, 0777)
    if err != nil {
        fmt.Println(err)
    }
    data, err := ioutil.ReadFile("myfile.txt")
    if err != nil {
        fmt.Println(err)
    }
    fmt.Println(data)
}