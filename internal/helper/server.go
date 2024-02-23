package helper

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func WriteResponse(rw http.ResponseWriter, statusCode int, response any) {
	resJson, _ := json.Marshal(response)

	rw.Header().Add("content-type", "application/json")

	rw.WriteHeader(statusCode)

	fmt.Fprint(rw, string(resJson))
}
