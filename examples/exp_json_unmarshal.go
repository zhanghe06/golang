package main

import (
	"encoding/json"
	"fmt"
)

type Source struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type Target struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type Item struct {
	Code   string  `json:"code"`
	Count  int     `json:"count"`
	Source *Source `json:"source,omitempty"`
	Target *Target `json:"target,omitempty"`
}

type Items struct {
	Items []Item `json:"items"`
}

func translate(testData []byte) {
	var items Items
	err := json.Unmarshal(testData, &items)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	result, err := json.Marshal(&items)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Printf("%s\n", result)

}

func main() {
	testDataA := []byte(`
{
	"items": [
		{
			"code": "DA-001",
			"count": 10,
			"source": {
				"id": "1",
				"name": "a"
			},
			"target": {
				"id": "2",
				"name": "b"
			},
			"other": "..."
		}
	]
}
`)
	testDataB := []byte(`
{
	"items": [
		{
			"code": "DA-001",
			"count": 10,
			"source": {
				"id": "1",
				"name": "a"
			},
			"target": {
				"id": "2",
				"name": "b"
			}
		}
	]
}
`)
	testDataC := []byte(`
{
	"items": [
		{
			"code": "DA-001",
			"count": 10,
			"source": {
				"id": "1",
				"name": "a"
			},
			"target": null
		}
	]
}
`)
	testDataD := []byte(`
{
	"items": [
		{
			"code": "DA-001",
			"count": 10,
			"source": {
				"id": "1",
				"name": "a"
			}
		}
	]
}
`)
	var testDataE []byte

	translate(testDataA)
	translate(testDataB)
	translate(testDataC)
	translate(testDataD)
	translate(testDataE)
}
