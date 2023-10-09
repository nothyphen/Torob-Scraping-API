package torobproduct

import (
	"fmt"
	//"io"
	"log"
	"net/http"
	"strings"

	"github.com/gocolly/colly/v2"
)

func Torob(product_name string) (string, error) {
	// create varuiable
	var url 	string

	// create own query for torob website
	query := product_name
	words := strings.Split(query, " ")
	url = "https://torob.com/search/?query=" + strings.Replace(query, " ","%20", len(words)-1)

	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal("Not Found !")
	}
	req.Header.Set("Referer", "https://www.google.com/")
	req.Header.Set("Upgrade-Insecure-Requests", "1")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/117.0.0.0 Safari/537.36")
	req.Header.Set("sec-ch-ua", `"Google Chrome";v="117", "Not;A=Brand";v="8", "Chromium";v="117"`)
	req.Header.Set("sec-ch-ua-mobile", "?0")
	req.Header.Set("sec-ch-ua-platform", `"Windows"`)

	res, err := client.Do(req)
	if err != nil {
		log.Fatal("Cant Not Send Request")
	}

	defer res.Body.Close()
	//body, err := io.ReadAll(res.Body)
	//if err != nil {
	//	log.Fatal("Cant read body page")
	//}

	//fmt.Printf("%s\n", body)

	c := colly.NewCollector()
	c.OnHTML(".jsx-2514672dc9197d80 product-name", func(h *colly.HTMLElement) {
		name := h.Index
		fmt.Println(name)
	})

	c.Visit(url)

	return "ok", nil
}
