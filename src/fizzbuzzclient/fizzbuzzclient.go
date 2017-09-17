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
fizzbuzzclient sends http fizzbuzz requests.
`)+"\n")
		flag.PrintDefaults()
	}
	jobs := flag.Int("jobs", 1, "concurrent jobs")
	number := flag.Int("number", 1, "number request per jobs")
	port := flag.Uint("port", 8084, "server port")
	limit := flag.Int("limit", 16, "fizzbuzz limit")
	flag.Parse()

	count := int32(0)
	ticker := time.NewTicker(time.Second)
	go func() {
		for _ = range ticker.C {
			log.Printf("%d requests succeeded\n", atomic.LoadInt32(&count))
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
				_, err := client.FizzBuzz(3, 5, *limit, "fizz", "buzz")
				if err != nil {
					log.Fatalf("error: %v", err)
				}
				atomic.AddInt32(&count, 1)
			}
		}()
	}
	wg.Wait()
	ticker.Stop()
	log.Printf("%d requests in: %v\n", count, time.Since(start))
}
