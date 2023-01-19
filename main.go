package main

import (
	"index/enron/dominio"
	"log"
	"net/http"
	_ "net/http/pprof"
	"sync"
)

func main() {
	go func() {
		log.Println(http.ListenAndServe("localhost:6060", nil))
	}()
	var wg sync.WaitGroup
	wg.Add(1)
	go dominio.Search()
	wg.Wait()

}
