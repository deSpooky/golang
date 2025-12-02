package main

import (
	"encoding/json"
	"net/http"
	"strconv"
)

func GetCardHandler(cache *Cache) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idStr := r.URL.Query().Get("id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			http.Error(w, "invalid id", http.StatusBadRequest)
			return
		}

		card, ok := cache.Get(id)
		if !ok {
			http.Error(w, "card not found", http.StatusNotFound)
			return
		}

		json.NewEncoder(w).Encode(card)
	}
}