package requests

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

const host = "http://localhost:9993"
const zerotierToken = "a9bxlzwe8lwrh8pe64m07h3b"
const headerName = "X-ZT1-Auth"

func ZtRequestGet(url string) string {
	req, err := http.NewRequest(http.MethodGet, host+url, nil)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	req.Header.Set(headerName, zerotierToken)
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	body, _ := io.ReadAll(res.Body)
	return string(body)
}

func commonRequest(method string, url string, data string) string {
	content := strings.NewReader(data)
	req, err := http.NewRequest(method, host+url, content)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	req.Header.Set(headerName, zerotierToken)
	req.Header.Set("Content-Type", "application/json")
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	body, _ := io.ReadAll(res.Body)
	return string(body)
}

func ZtRequestPost(url string, data string) string {
	return commonRequest(http.MethodPost, url, data)
}

func ZtRequestDelete(url string, data string) string {
	return commonRequest(http.MethodDelete, url, data)
}
