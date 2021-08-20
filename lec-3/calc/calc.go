package calc

import (
	"encoding/json"
	"net/http"
	"strconv"
)
type Data struct {
	Num1 string `json:"num1"`
	Num2 string	`json:"num2"`
	Opr string	`json:"opr"`
}

func AddCorsHeader(res http.ResponseWriter) {
	headers := res.Header()
	headers.Set("Access-Control-Allow-Origin", "*")
	headers.Set("Access-Control-Allow-Headers", "Content-Type")
}

func Submit(w http.ResponseWriter, req *http.Request) {
	AddCorsHeader(w)
	if req.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}

	var x Data
	err := json.NewDecoder(req.Body).Decode(&x)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	a, _ := strconv.ParseFloat(x.Num1, 64)
	b, _ := strconv.ParseFloat(x.Num2, 64)
	r:=0.0
	switch x.Opr {
		case "÷": 
			r = a/b
		case "×":
			r = a*b
		case "+":
			r = a+b
		case "-":
			r = a-b
		}
	err = json.NewEncoder(w).Encode(r)
	if err != nil {
		return 
	}
}