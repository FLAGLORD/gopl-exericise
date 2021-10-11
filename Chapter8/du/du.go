package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"sync"
	"time"
)

var sema = make(chan struct{}, 20)

func dirents(dir string) []os.FileInfo {
	sema <- struct{}{}
	defer func() {
		<-sema
	}()
	entries, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Errorf("du: %v\n", err)
		return nil
	}
	return entries
}

func walkDir(dir string, fileSize chan<- int64) {
	defer n.Done()
	for _, entry := range dirents(dir) {
		if entry.IsDir() {
			n.Add(1)
			subdir := filepath.Join(dir, entry.Name())
			go walkDir(subdir, fileSize)
		} else {
			fileSize <- entry.Size()
		}
	}
}

var verbose = flag.Bool("v", false, "show verbose progress messages")
var n sync.WaitGroup

func main() {
	flag.Parse()
	roots := flag.Args()

	var tick <-chan time.Time
	if *verbose {
		tick = time.Tick(500 * time.Millisecond)
	}
	if len(roots) == 0 {
		roots = []string{"."}
	}
	fileSize := make(chan int64)

	for _, root := range roots {
		n.Add(1)
		go walkDir(root, fileSize)
	}

	go func() {
		n.Wait()
		close(fileSize)
	}()

	var nfiles, nbytes int64
loop:
	for {
		select {
		case size, ok := <-fileSize:
			if !ok {
				break loop
			}
			nfiles++
			nbytes += size
		case <-tick:
			printDiskUsage(nfiles, nbytes)

		}
	}
	printDiskUsage(nfiles, nbytes)
}

func printDiskUsage(nfiles, nbytes int64) {
	fmt.Printf("%d files %.1f GB\n", nfiles, float64(nbytes)/1e9)
}
