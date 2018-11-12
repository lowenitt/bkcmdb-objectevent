package main

import (
	"encoding/json"
	"log"
	"net/http"
)

// response response info struct
type response struct {
	Result bool        `json:"result"`
	Code   int         `json:"bk_error_code"`
	ErrMsg string      `json:"bk_error_msg"`
	Data   interface{} `json:"data"`
}

// WriteJSON respone json to client
func WriteJSON(httpCode int, rsp *response, w http.ResponseWriter) {

	byteBody, err := json.Marshal(rsp)
	if err != nil {
		log.Printf("json marshal error, error:%s \n", err.Error())
		return
	}
	w.WriteHeader(httpCode)
	_, err = w.Write(byteBody)
	if err != nil {
		log.Printf("response %s error, error:%s \n", string(byteBody), err.Error())
		return
	}
	return
}

// GetResponseErrorBody get response error struct
func GetResponseErrorBody(errCode int, errMsg string) *response {
	return &response{
		Result: false,
		Code:   errCode,
		ErrMsg: errMsg,
		Data:   nil,
	}

}

// GetResponseSuccBody get response success struct
func GetResponseSuccBody(data interface{}) *response {
	return &response{
		Result: true,
		Data:   nil,
	}

}
