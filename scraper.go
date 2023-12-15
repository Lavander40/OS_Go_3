package main

import (
	"fmt"

	"github.com/gocolly/colly"
)

func main() {
	c := colly.NewCollector()

	c.OnHTML("div.vulnerable_code_area", func(h *colly.HTMLElement) {
		fmt.Println("FOUND")
	})

	c.Visit("http://localhost/DVWA-master/vulnerabilities/brute/")
}

func generator(maxChar int) []string {
	alphabet := "abcdefghijklmnopqrstuvwxyz"
	var result []string = []string{""}
	var last, next int
	for i := 1; i <= maxChar; i++ {
		last = len(result)
		for _, str := range result[next:] {
			for _, char := range alphabet {
				result = append(result, str+string(char))
			}
		}
		next = last
	}
	return result
}
