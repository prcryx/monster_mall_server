package utils

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/prcryx/monster_mall/internal/common/constants"
	"github.com/prcryx/monster_mall/internal/common/exception"
)

func ResponseWithJSONData(w http.ResponseWriter, code int, payload interface{}) {
	data, err := json.Marshal(payload)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Add(constants.ContentType, constants.ApplicationJson)
	w.WriteHeader(code)
	w.Write(data)
}

func ResponseWithError(w http.ResponseWriter, code int, msg string) {
	if code > 499 {
		log.Printf("5xx Error: %v", msg)
	}

	ResponseWithJSONData(w, code, exception.ErrorResponse{
		Error: msg,
	})
}
