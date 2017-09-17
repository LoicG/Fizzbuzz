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
	flag.Parse()

	count := int32(0)
	start := time.Now()
	wg := sync.WaitGroup{}
	for i := 0; i != *jobs; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j != *number; j++ {
				client := client.NewClient(uint32(*port))
				_, err := client.FizzBuzz(3, 5, 16, "fizz", "buzz")
				if err != nil {
					log.Fatalf("error: %v", err)
				}
				atomic.AddInt32(&count, 1)
			}
		}()
	}
	wg.Wait()
	log.Printf("%d requests in: %v\n", count, time.Since(start))
}
