package main

import (
	"context"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lmicroseconds | log.Lshortfile)

	do := func() {
		req, e := http.NewRequest("GET", "http://127.0.0.1:8080", nil)
		if e != nil {
			log.Print(e.Error())
			return
		}

		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()
		req = req.WithContext(ctx)

		resp, e := http.DefaultClient.Do(req)
		if e != nil {
			log.Print(e.Error())
			return
		}
		_, e = io.Copy(ioutil.Discard, resp.Body)
		if e != nil {
			log.Print(e.Error())
			return
		}
		e = resp.Body.Close()
		if e != nil {
			log.Print(e.Error())
			return
		}
	}

	const N = 1e2
	const C = 1e1
	sem := make(chan struct{}, C)
	for i := 0; i < N; i++ {
		sem <- struct{}{}
		go func() {
			defer func() {
				<-sem
			}()
			do()
		}()
	}
	for i := 0; i < C; i++ {
		sem <- struct{}{}
	}

	<-time.After(time.Hour)
}
