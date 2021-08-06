package handlers

import (
	"net/http"
)

func Cal(w http.ResponseWriter, req *http.Request) {
	//đoạn này xem lại phần req,res rồi code tiếp hic
}

func Calculate(operator string, a,b float64) float64{
	switch operator {
	case "sum":
		return a+b
	case "sub":
		return a-b
	case "mul":
		return a*b
	case "div":
		if b==0 {
			return 0
		}
		return a/b
	default: return 0
	}
}