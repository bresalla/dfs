package ware_test

import (
	"fmt"
	"io"
	"net/http"
	"testing"
)

func Test(t *testing.T) {

	url := "http://localhost:8080/config/ac7ad792-66ee-4360-b291-7580fac6f1a9"

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("Accept", "application/json, application/xml")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)

	fmt.Println(res)
	fmt.Println(string(body))
}
