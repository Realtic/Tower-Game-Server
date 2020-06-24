package resapi

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

var (
	statusError = "error"
	statusOk    = "ok"
)

// ResponseAPI ...
type ResponseAPI struct {
	Writer http.ResponseWriter
}

// ResponseBlob ...
type ResponseBlob struct {
	Status string
	Data   interface{}
}

func (r *ResponseAPI) writeBlob(blob *ResponseBlob) {
	// Here or in Error()?
	// w.WriteHeader(http.StatusInternalServerError)

	json, err := json.Marshal(blob)
	if err != nil {
		log.Print(err)
		return
	}

	fmt.Fprint(r.Writer, string(json))
}
