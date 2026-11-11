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
}

type Request struct {
	Url	string
	Id string
	//SesionCookie string
}

type Result struct {
	StatusCode int
	Body []byte
	ContentLength int
	Url string
}

func (b Baseline) Evaluate(res Result) bool {
	/* this is a new thing i learned Golang idiomatic style bool instead of 
	if res.StatusCode != b.StatusCode || res.ContentLength != b.ContentLength {
		return true
	}                          |  
	return false i can do this V*/
	return res.StatusCode != b.StatusCode || !bytes.Equal(res.Body, b.Body) 
}

func (req Request) fetch(client *http.Client) (Result, error) {

	maxAllowedbytes := int64(4 * MB)

	resp, err := client.Get(req.Url)
	if err != nil {
		return Result{}, fmt.Errorf("unexpected error has occurred: %w", err)
	}

	defer resp.Body.Close()

	limitedReader := io.LimitReader(resp.Body, maxAllowedbytes)
	b, err := io.ReadAll(limitedReader)
	
	return Result {
		StatusCode: resp.StatusCode,
		Body: b,
		ContentLength: len(b),
		Url: req.Url,
	}, nil
}

func main() {
	url := flag.String("url", "", "Target URL template (e.g. https://target.com/user/{ID})")
	start := flag.Int("start", 1, "Starting ID")
	end := flag.Int("end", 100, "Ending ID")
	baselineID := flag.String("baseline", "999999", "Known invalid ID for baseline")
	verbose := flag.Bool("v", false, "Enable verbose output")
	_ = verbose

	flag.Parse()

	client := &http.Client{Timeout: 10 * time.Second}

	baseUrl := strings.ReplaceAll(*url, "{ID}", *baselineID)
	baseReq := Request{
		Url: baseUrl,
		Id: *baselineID,
	}

	res, err := baseReq.fetch(client)
	if err != nil {											// here is where the color are used its called ANSI ecape Codes 
		fmt.Printf("%s unexpected error has occurred:%s %s%v%s\n", Red, Reset, BoldWhite, err, Reset)
		return 
	}

	baseline := Baseline{
		StatusCode: res.StatusCode,
		ContentLength: res.ContentLength,
		Body: res.Body,
	}

	for i := *start; i <= *end; i++ {
		idStr := strconv.Itoa(i)
		targetUrl := strings.ReplaceAll(*url, "{ID}", idStr)
		req := Request{
			Url: targetUrl,
			Id: idStr,
		}

		res, err = req.fetch(client)
		if err != nil {											// here is where the color are used its called ANSI ecape Codes 
			fmt.Printf("%s unexpected error has occurred:%s %s%v%s\n", Red, Reset, BoldWhite, err, Reset)
		}

		if baseline.Evaluate(res) {									// here is where the color are used its called ANSI ecape Codes 
			fmt.Printf("%s{Hit}%s Found ID: %d Status: %d Length: %d\n",Green, Reset, i, res.StatusCode, res.ContentLength)
		}
	}
}