package main

import (
	"client"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
	"sync"
	"sync/atomic"
	"time"
)

func main() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, strings.TrimSpace(`
client sends http requests to the fizz-buzz server [OPTIONS]
`)+"\n")
		flag.PrintDefaults()
	}
	jobs := flag.Int("jobs", 1, "concurrent jobs")
	number := flag.Int("number", 1, "number request per jobs")
	port := flag.Uint("port", 8084, "server port")
	limit := flag.Int("limit", 16, "fizzbuzz limit")
	invalid := flag.Bool("error", false, "if true sends invalid request")
	flag.Parse()

	fizz := 3
	buzz := 5
	if *invalid {
		fizz = -1
	}

	count := int64(0)
	ticker := time.NewTicker(time.Second)
	go func() {
		for _ = range ticker.C {
			log.Printf("%d requests\n", atomic.LoadInt64(&count))
		}
	}()

	start := time.Now()
	wg := sync.WaitGroup{}
	for i := 0; i != *jobs; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			client := client.NewClient(uint32(*port), *jobs)
			for j := 0; j != *number; j++ {
				client.FizzBuzz(fizz, buzz, *limit, "fizz", "buzz")
				atomic.AddInt64(&count, 1)
			}
		}()
	}
	wg.Wait()
	ticker.Stop()
	end := time.Since(start)
	log.Printf("%d Requests in: %v\n", count, end)
	seconds := int64(end.Seconds())
	if seconds != 0 {
		log.Printf("%d Requests per Second", count/seconds)
	}
}
