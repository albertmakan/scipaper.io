package helpers

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func JSONResponse(w http.ResponseWriter, status int, data interface{}) {
	w.WriteHeader(status)
	if data == nil {
		return
	}
	resp, _ := json.Marshal(data)
	w.Header().Set("Content-Type", "application/json")
	w.Write(resp)
}

func ReadJSONBody(r *http.Request, data any) error {
	body, err := ReadBody(r)
	if err != nil {
		return err
	}
	return json.Unmarshal(body, data)
}

func ReadBody(r *http.Request) ([]byte, error) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}
	r.Body.Close()
	r.Body = ioutil.NopCloser(bytes.NewBuffer(body))
	return body, nil
}