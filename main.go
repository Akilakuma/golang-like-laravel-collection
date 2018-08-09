package main

import (
	"log"
	"time"

	collection "./collection"
)

func main() {
	// 將input convert 成coll

	inputData4 := collection.CollSets{
		{
			"word": "hello man!",
			"code": "321234",
		},
		{
			"word": "hello man!",
			"code": "321234",
		},
		{
			"word": "hello man!",
			"code": "123456",
		},
		{
			"word": "hello girl!",
			"code": "789102",
		},
	}

	tS := time.Now()
	inputData4.WhereStr("word", "hello man!").Groupby("code").OrderBy().Get()
	td := time.Since(tS)
	log.Print("處理時間(奈秒)： ")
	log.Println(td.Nanoseconds())
}
