package parser

import (
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
)

func ParseWebPage(cryptocurrency string) string {
	parseElement := ".text-white[data-socket-key='" + strings.ToLower(cryptocurrency) + "']"

	url := "https://www.doviz.com/kripto-paralar/" + strings.ToLower(cryptocurrency)

	client := &http.Client{
		Timeout: 30 * time.Second,
	}
	res, err := client.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	defer res.Body.Close()

	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
		return "0"
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)

	if err != nil {
		log.Fatal(err)
	}

	data := doc.Find(parseElement).Text()

	return data
}
