package service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

func Post(jsonResult string, path string, nameFile string, count int) {
	data := []byte(jsonResult)
	uri := "http://localhost:4080/api/_bulk"
	h := &http.Client{}
	req, _ := http.NewRequest("POST", uri, bytes.NewBuffer(data))
	req.SetBasicAuth("admin", "Complexpass#123")
	rs, err := h.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	if rs.Body != nil {
		defer rs.Body.Close()
	}
	body, readErr := ioutil.ReadAll(rs.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}
	var result Response
	json.Unmarshal(body, &result)
	if result.Record_count == 0 {
		fmt.Println("Error Indexed : " + path + " : file " + nameFile + " : Count " + strconv.Itoa(count))
		WriteFile(jsonResult, path, count, nameFile)
	} else {
		fmt.Println("Success: " + path + " : file " + nameFile + " : Count " + strconv.Itoa(count))

	}

}

type Response struct {
	Message      string `json:"message"`
	Record_count int    `json:"record_count"`
}
