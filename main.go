// package main

// import (
// 	"context"
// 	"fmt"
// 	"log"
// 	"time"

// 	"github.com/chromedp/chromedp"
// )

// func main() {

// 	ctx, cancel := chromedp.NewContext(
// 		context.Background(),
// 	)
// 	defer cancel()

// 	ctx, cancel = context.WithTimeout(ctx, 6*time.Second)
// 	defer cancel()

// 	url := "http://localhost:80/DVWA-master/vulnerabilities/brute/"
// 	var res string

// 	err := chromedp.Run(ctx,
// 		chromedp.Navigate(url),
// 		chromedp.SendKeys("input[name=username]", "admin"),
// 		chromedp.SendKeys("input[name=password]", "password"),
// 		// chromedp.Click("button", chromedp.NodeVisible),
// 		chromedp.Submit("input[name=username]"),
// 		chromedp.WaitVisible(`//*[contains(., 'Welcome to the password protected area ')]`),
// 		chromedp.Text(`(//*[contains(., 'Welcome to the password protected area ')]`, &res),
// 	)

// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	fmt.Println(res)
// }
