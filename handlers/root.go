package handlers

import (
	"net/http"
	"encoding/json"
	"github.com/managef/models/log"
)

type Api struct{
	Name    string
}
func GetApi(w http.ResponseWriter, r *http.Request) {
	log.Info("get Api request")
	json.NewEncoder(w).Encode(Api{Name: "Hello to ManageF"})
}
