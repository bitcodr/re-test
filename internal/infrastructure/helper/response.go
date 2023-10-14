package helper

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/bitcodr/re-test/internal/domain/model"
)

func ResponseError(res http.ResponseWriter, message string, err error) {
	if err != nil {
		log.Println(err)
	}

	res.Header().Add("Content-Type", "application/json")
	res.WriteHeader(http.StatusInternalServerError)
	_, err = res.Write([]byte(message))
	if err != nil {
		log.Println(err)
	}
}

type Success interface {
	string | []int | *model.Order
}

func ResponseSuccess[T Success](res http.ResponseWriter, message T) {
	data, err := json.Marshal(message)
	if err != nil {
		ResponseError(res, "error in encoding", err)
		return
	}

	res.Header().Add("Content-Type", "application/json")
	res.WriteHeader(http.StatusOK)
	if _, err = res.Write(data); err != nil {
		ResponseError(res, "error in writing response", err)
	}
}
