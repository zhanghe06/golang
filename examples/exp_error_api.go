package main

import (
	"encoding/json"
	"log"

	"github.com/pkg/errors"
)

type ApiResponse struct {
	Name string
}

type ApiError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func (e *ApiError) Error() string {
	es, _ := json.Marshal(e)
	return string(es)
}

func apiClient(req bool) (res *ApiResponse, err error) {
	defer func() {
		if err != nil {
			//log.Println("api call error")
			err = errors.Wrap(err, "api call error")
		}
	}()

	if req {
		res = &ApiResponse{
			Name: "A",
		}
	} else {
		err = &ApiError{
			Code:    400000001,
			Message: "this is a error",
		}
	}
	return
}

func apiMain() (err error) {
	var res *ApiResponse
	res, err = apiClient(false) // true/false
	if err != nil {
		return
	}
	log.Println(res)
	return
}

func main() {
	err := apiMain()

	//switch t := interface{}(err).(type) {
	switch t := interface{}(errors.Cause(err)).(type) { // 处理被errors.Wrap封装的error
	case *ApiError:
		log.Println("ApiError:", t)
		log.Println("err:", err)
	default:
		log.Println("err:", t)
	}
}
