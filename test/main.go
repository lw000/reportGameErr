// test project main.go
package main

import (
	"encoding/json"
	"fmt"
	"github.com/lw000/gocommon/utils"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
	"sync"
	"sync/atomic"
	"time"
)

type Config struct {
	Count  int    `json:"count"`
	Method string `json:"method"`
	Url    string `json:"url"`
	Data   string `json:"data"`
}

func NewConfig() *Config {
	return &Config{}
}

func (c *Config) LoadConfig() bool {
	data, err := ioutil.ReadFile("./conf.json")
	if err != nil {
		log.Println(err.Error())
		return false
	}

	err = json.Unmarshal(data, c)
	if err != nil {
		log.Println(err.Error())
		return false
	}

	return true
}

var (
	wg           *sync.WaitGroup
	successCount uint64
)

func HttpPost(w *sync.WaitGroup, url string, data string) {
	go func(w *sync.WaitGroup) {
		w.Add(1)
		defer w.Done()

		resp, err := http.Post(url, "application/x-www-form-urlencoded", strings.NewReader(data))
		if err != nil {
			log.Println(err.Error())
			return
		}
		defer resp.Body.Close()

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Println(err.Error())
			return
		}

		if len(body) > 0 {

		}

		log.Printf("[%d] %s", atomic.AddUint64(&successCount, 1), string(body))
	}(w)
}

func HttpGet(w *sync.WaitGroup, url string) {
	w.Add(1)
	defer w.Done()

	resp, err := http.Get(url)
	if err != nil {
		log.Println(err.Error())
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err.Error())
		return
	}
	if len(body) > 0 {

	}
	log.Printf("[%d] %s", atomic.AddUint64(&successCount, 1), string(body))
}

func LogTest(u string, module string, eventId string, info string) {
	v := url.Values{}
	v.Set("module", module)
	v.Set("eventId", eventId)
	v.Set("info", info)
	s := v.Encode()
	go HttpGet(wg, u+"?"+s)
}

func main() {
	cfg := NewConfig()
	if !cfg.LoadConfig() {
		return
	}

	wg = &sync.WaitGroup{}

	start := time.Now()

	for i := 0; i < cfg.Count; i++ {
		LogTest(cfg.Url, "10101", tyutils.UUID(), tyutils.UUID())
		time.Sleep(time.Millisecond * time.Duration(10))
	}

	wg.Wait()

	end := time.Now()
	fmt.Printf("%d, %v\n", successCount, end.Sub(start))
}
