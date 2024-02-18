package main

import (
	"net/http"
)

/*
 * 错误示例
 */
func httpGetBad() error {
	res, err := http.Get("http://notexists")
	defer res.Body.Close()
	if err != nil {
		return err
	}
	// ..code...
	return nil
}

func httpGetGood() error {
	res, err := http.Get("http://notexists")
	if res != nil {
		defer res.Body.Close()
	}
	if err != nil {
		return err
	}
	// ..code...
	return nil
}

func main() {
	//httpGetBad()
	httpGetGood()
}
