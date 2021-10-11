package main

import (
	"exercise/Chapter8/links"
	"fmt"
	"log"
	"os"
)

var tokens = make(chan struct{}, 20)

func crawl(url string) []string {
	fmt.Println(url)
	// 获取 token 是发送一个空结构体看上去，有点反直觉，其实效果是一样的，毕竟放满了就会阻塞，从而控制速率
	tokens <- struct{}{} // acquire a token
	list, err := links.Extract(url)
	<-tokens // release the token
	if err != nil {
		log.Print(err)
	}
	return list
}

func main() {
	worklist := make(chan []string)

	go func() {
		worklist <- os.Args[1:]
	}()

	// Crawl the web concurrently
	seen := make(map[string]bool)
	for list := range worklist {
		for _, link := range list {
			if !seen[link] {
				seen[link] = true
				go func(link string) {
					worklist <- crawl(link)
				}(link)
			}
		}
	}
}
