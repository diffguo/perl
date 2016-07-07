package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"
)

const (
	HTTP_ADDR = "http://127.0.0.1:8100"
	USER_ID   = "a3156ff8-5cff-4661-9e44-e53bc8ce847a"
)

func HttpGet(url string) (string, error) {
	rsp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer rsp.Body.Close()

	d, err := ioutil.ReadAll(rsp.Body)
	if err != nil {
		return "", err
	}

	return string(d), nil
}

type HiRsp struct {
	Rsp string `json:"rsp"`
}

func TestHi(t *testing.T) {
	url := fmt.Sprintf("%s/hi?hi=Daniel", HTTP_ADDR)
	content, err := HttpGet(url)
	if err != nil {
		t.Error(err)
		return
	}

	rsp := &HiRsp{}
	err = json.Unmarshal([]byte(content), rsp)
	if err != nil {
		t.Error(err)
		return
	}

	if rsp.Rsp != "hi:Daniel" {
		fmt.Println(rsp)
		t.Error("rsp incorrect")
		return
	}
}

type GroupCountRsp struct {
	Count int `json:"count"`
}

func TestGoupCount(t *testing.T) {
	url := fmt.Sprintf("%s/groupcount?uid=%s", HTTP_ADDR, USER_ID)
	content, err := HttpGet(url)
	if err != nil {
		t.Error(err)
		return
	}

	rsp := &GroupCountRsp{}
	err = json.Unmarshal([]byte(content), rsp)
	if err != nil {
		fmt.Println("content:", content)
		t.Error(err)
		return
	}

	if rsp.Count == 0 {
		fmt.Println(rsp)
		t.Error("count incorrect")
		return
	}
}
