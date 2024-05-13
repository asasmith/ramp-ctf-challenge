package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

const url = "https://tns4lpgmziiypnxxzel5ss5nyu0nftol.lambda-url.us-east-1.on.aws/challenge"

func fetchAndParseHtml(url string) []string {
	var values []string

	res, err := http.Get(url)

	if err != nil {
		log.Fatal(err)
	}

	defer res.Body.Close()

	if res.StatusCode != 200 {
		log.Fatalf("Recieved non 200 status code: %d %s\n", res.StatusCode, res.Status)
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)

	if err != nil {
		log.Fatal(err)
	}

	doc.Find("code[data-class]").Each(func(index int, code *goquery.Selection) {
		dataClass, _ := code.Attr("data-class")
		if strings.HasPrefix(dataClass, "23") { // Check if data-class starts with "23"
			code.Find("div[data-tag]").Each(func(index int, div *goquery.Selection) {
				dataTag, _ := div.Attr("data-tag")
				if strings.HasSuffix(dataTag, "93") { // Check if data-tag ends with "93"
					div.Find("span[data-id]").Each(func(index int, span *goquery.Selection) {
						dataId, _ := span.Attr("data-id")
						if strings.Contains(dataId, "21") { // Check if data-id contains "21"
							span.Find("i.char").Each(func(index int, i *goquery.Selection) {
								value, exists := i.Attr("value")
								if exists { // Directly append the value of <i class="char">
									values = append(values, value)
								}
							})
						}
					})
				}
			})
		}
	})

	return values
}

func main() {
	url := fetchAndParseHtml(url)

	fmt.Println(strings.Join(url, ""))
}
