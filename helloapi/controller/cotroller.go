package controller

import (
	"encoding/json"
	"net/http"
)

func HomePage(w http.ResponseWriter, r *http.Request) {

	queryValues, ok := r.URL.Query()["hi"]
	if !ok || len(queryValues) < 1 {
		json.NewEncoder(w).Encode(Response{
			Status: 400,
			Res:    "Not found",
		})
	} else {
		json.NewEncoder(w).Encode(Response{
			Status: 200,
			Res:    "Hello",
		})
	}
}
