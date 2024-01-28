package util

import (
	"encoding/json"
	"io"
	"log"
)

func ReadRequestBody(reader io.Reader, model interface{}) (err error) {
	bytes, err := io.ReadAll(reader)
	if err != nil {
		log.Println("err in reading request body: ", err)
		return
	}

	err = json.Unmarshal(bytes, model)
	if err != nil {
		log.Println("err in unmarshalling body into model", err)
		return
	}

	return
}
