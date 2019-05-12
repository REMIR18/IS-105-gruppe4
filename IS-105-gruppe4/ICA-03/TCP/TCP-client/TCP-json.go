package tcp

import (
	"encoding/json"
	"net/http"
)

type Person struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

var human []Person

func tcp() {
	human = append(human, Person{Name: "Henriette", Email: "henrf18@uia.no"})
	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(human)
	})
	http.ListenAndServe(":5000", nil)
}
