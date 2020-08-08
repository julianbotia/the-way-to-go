package http

import (
	"bytes"
	"encoding/json"
	"log"
)

func checkError(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

func jsonPrettyPrint(in string) string {
	var out bytes.Buffer
	err := json.Indent(&out, []byte(in), "", "\t")
	if err != nil {
		return in
	}
	return out.String()
}
