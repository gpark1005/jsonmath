package handlers

import (
	"encoding/json"
	"jsonmath/data"
	"net/http"
)

func PostNewData(w http.ResponseWriter, r *http.Request) {
	od := data.OuterData{}

	_ = json.NewDecoder(r.Body).Decode(&od)

	od.SetId()
	err := od.DoConversion()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Num field should only have 2 numbers"))
	}

	jsonBytes, _ := json.Marshal(od)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonBytes)

}

func UpdateData(w http.ResponseWriter, r *http.Request) {
	od := data.OuterData{}

	_ = json.NewDecoder(r.Body).Decode(&od)

	err := od.DoConversion()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Num field should only have 2 numbers"))
	}

	jsonBytes, _ := json.Marshal(od)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonBytes)

}
