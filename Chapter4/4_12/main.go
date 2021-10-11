package main

import (
	"encoding/json"
	"exercise/Chapter4/4_12/xkcd"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func IsFileExists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	return true
}

var Index []*xkcd.Comic

func init() {
	//  File not exists
	if !IsFileExists("index.json") {
		for i := 1; i <= 100; i++ {
			comic, err := xkcd.FetchComic(i)
			if err != nil {
				log.Fatal(err)
			}
			Index = append(Index, comic)
		}
		jsonContent, _ := json.MarshalIndent(Index, "", " ")
		ioutil.WriteFile("index.json", jsonContent, 0666)
	} else {
		loadJson()
	}
}

func loadJson() {
	file, _ := os.Open("index.json")
	err := json.NewDecoder(file).Decode(&Index)
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	flag.Parse()
	keywords := flag.Args()
	var result []*xkcd.Comic
	for _, comic := range Index {
		flag := true
		for _, keyword := range keywords {
			if !comic.Match(keyword) {
				flag = false
				break
			}
		}
		if flag {
			result = append(result, comic)
		}
	}
	for _, comic := range result {
		fmt.Printf("%s\thttps://xkcd.com/%d/info.0.json\n", comic.Title, comic.Num)
	}
}
