/*Romans 8:18 "I consider that out present sufferings are not worth comparing with the glory that will be revealed in us."*/

package main 

import (
	"fmt"
	"net/http"
	"io"
	"bytes"
	"flag"
	"time"
	"strings"
	"strconv"
	"sync"
)

const (
	KB = 1024
	MB = 1024 * KB
)

/* this is a new thing i learn to it is more of an astetic thing im color blind so red and green aren't going to help that much 
but it might help someone else using this */
const (
	Reset = "\033[0m"
	BoldWhite = "\033[1;37m"
	Red = "\033[31m"
	Green = "\033[32m"
	Yellow = "\033[33m"
	Blue = "\033[34m"
)

type Evaluator interface {
	Evaluate(res Result) bool
}

type Baseline struct {
	StatusCode	int
	ContentLength int
	Body []byte
	SessionCookie string
}

type Target  struct {
	Url	string
	Id string
	SessionCookie string
}

type Result struct {
	StatusCode int
	Body []byte
	ContentLength int
	Url string
}

type Limiter struct {
	Mu sync.Mutex
	LastRequest	time.Time
	Interval time.Duration
}

func NewLimiter(rps int) *Limiter {
	interval := time.Second / time.Duration(rps)
	return &Limiter{
		Interval: interval,
	}
}

func (l *Limiter) wait() {
	l.Mu.Lock()
	wait := l.Interval - time.Since(l.LastRequest)

	if wait > 0 {
		l.LastRequest = time.Now().Add(wait)
	} else {
		l.LastRequest = time.Now()
	}

	l.Mu.Unlock()

	if wait > 0 {
		time.Sleep(wait)
	}
}

func (b Baseline) Evaluate(res Result) bool {
	/* this is a new thing i learned Golang idiomatic style bool instead of 
	if res.StatusCode != b.StatusCode || res.ContentLength != b.ContentLength {
		return true
	}                          |  
	return false i can do this V*/
	return res.StatusCode != b.StatusCode || !bytes.Equal(res.Body, b.Body) 
}

func (target Target) Fetch(client *http.Client) (Result, error) {

	maxAllowedbytes := int64(4 * MB)

	req, err := http.NewRequest("GET", target.Url, nil)
	if err != nil {
		return Result{}, fmt.Errorf("unexpected error has occurred: %w\n", err)
	}

	if target.SessionCookie != "" {
		req.Header.Set("Cookie", target.SessionCookie)
	}

	resp, err := client.Do(req)
	if err != nil {
		return Result{}, fmt.Errorf("unexpected error has occurred: %w\n", err)
	}

	defer resp.Body.Close()

	limitedReader := io.LimitReader(resp.Body, maxAllowedbytes)
	b, err := io.ReadAll(limitedReader)
	
	return Result {
		StatusCode: resp.StatusCode,
		Body: b,
		ContentLength: len(b),
		Url: target.Url,
	}, nil
}

func worker(jobs <-chan int, wg *sync.WaitGroup, urlTemplate string, cookie string, baseline Baseline, client *http.Client, limiter *Limiter) {

	defer wg.Done()

	for i := range jobs {
		idStr := strconv.Itoa(i)
		targetUrl := strings.ReplaceAll(urlTemplate, "{ID}", idStr)
		target := Target {
			Url: targetUrl,
			Id: idStr,
			SessionCookie: cookie,
		}

		limiter.wait() 

		res, err := target.Fetch(client)
		if err != nil {											// here is where the color are used its called ANSI ecape Codes 
			fmt.Printf("%s unexpected error has occurred:%s %s%v%s\n", Red, Reset, BoldWhite, err, Reset)
			continue
		}

		if baseline.Evaluate(res) {									// here is where the color are used its called ANSI ecape Codes 
			fmt.Printf("%s{Hit}%s Found ID: %d Status: %d Length: %d\n",Green, Reset, i, res.StatusCode, res.ContentLength)
		}
	}
}

func main() {
	url := flag.String("url", "", "Target URL template (e.g. https://target.com/user/{ID})")
	sessionCookie := flag.String("cookie", "", "Session Cookie (e.g. session=ZZjPhwBbonLVK3kiSHlI7mUWoN94GQjg)")
	start := flag.Int("start", 1, "Starting ID")
	end := flag.Int("end", 100, "Ending ID")
	workers := flag.Int("workers", 2, "Concurrency")
	baselineID := flag.String("baseline", "999999", "Known invalid ID for baseline")
	verbose := flag.Bool("v", false, "Enable verbose output")
	rps := flag.Int("rps", 10, "Requests per second limit")
	_ = verbose

	flag.Parse()

	limiter := NewLimiter(*rps)

	var wg sync.WaitGroup
	jobs := make(chan int)

	client := &http.Client{Timeout: 10 * time.Second}

	baseUrl := strings.ReplaceAll(*url, "{ID}", *baselineID)
	baseReq := Target {
		Url: baseUrl,
		Id: *baselineID,
		SessionCookie: *sessionCookie,
	}

	res, err := baseReq.Fetch(client)
	if err != nil {											// here is where the color are used its called ANSI ecape Codes 
		fmt.Printf("%s unexpected error has occurred:%s %s%v%s\n", Red, Reset, BoldWhite, err, Reset)
		return 
	}

	baseline := Baseline{
		StatusCode: res.StatusCode,
		ContentLength: res.ContentLength,
		Body: res.Body,
	}

	for w := 0; w < *workers; w++ {
		wg.Add(1)
		go worker(jobs, &wg, *url, *sessionCookie, baseline, client, limiter)
	}

	for i := *start; i <= *end; i++ {
		jobs <- i
	}
	close(jobs)

	wg.Wait()
}

