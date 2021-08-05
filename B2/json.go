package main

import (
	"encoding/json"
	"fmt"
)
type Detail struct {
	Height float64
	Weight float64
}

type Person struct {
	Name string
	Age int
	Detail Detail
}

func main() {

	jsonString :=`[{
	"name":"Tung",
	"age":23,
	"detail": {
		"height":175,
		"weight":65
	}
	},
	{
		"name":"Hien",
		"age":19,
		"detail":{
			"height":175,
			"weight":65
			}
	}]`
var student []Person
json.Unmarshal([]byte(jsonString), &student)
for _,v:= range student {
fmt.Println(v.Name,v.Age, v.Detail.Height, v.Detail.Weight)
}
}