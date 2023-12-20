package utils

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

const zerotierToken = "w5l5b11t4x668x2o4ku2c4cg"
const headerName = "X-ZT1-Auth"

const host = "http://localhost:9993"

func GetPostData[K string, V any](ctx *gin.Context) map[K]V {
	dataMap := make(map[K]V)
	if err := ctx.ShouldBind(&dataMap); err != nil {
		fmt.Println("数据错误")
		ctx.JSON(http.StatusOK, map[string]interface{}{
			"msg": "提交的数据格式错误",
		})
		return nil
	}
	return dataMap
}

func ZerotierGet(path string, body string) (*http.Response, bool) {
	client := &http.Client{}
	content := strings.NewReader(body)
	req, err := http.NewRequest(http.MethodGet, host+path, content)
	if err != nil {
		fmt.Println(err)
		return nil, false
	}
	req.Header.Set(headerName, zerotierToken)
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return nil, false
	}
	return res, true
}

func ZerotierPost(path string, body string) (*http.Response, bool) {
	client := &http.Client{}
	content := strings.NewReader(body)
	req, err := http.NewRequest(http.MethodPost, host+path, content)
	if err != nil {
		fmt.Println(err)
		return nil, false
	}
	req.Header.Set(headerName, zerotierToken)
	req.Header.Set("Content-Type", "application/json")
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return nil, false
	}
	return res, true
}

func ZerotierDelete(path string) (*http.Response, bool) {
	client := &http.Client{}
	req, err := http.NewRequest(http.MethodDelete, host+path, nil)
	if err != nil {
		fmt.Println(err)
		return nil, false
	}
	req.Header.Set(headerName, zerotierToken)
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return nil, false
	}
	return res, true
}
