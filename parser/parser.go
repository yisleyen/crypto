package parser

import (
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
)

func ParseWebPage(cryptocurrency string) map[string]string {
	cryptocurrencys := make(map[string]string)

	parseElementDollar := ".text-white[data-socket-key='" + strings.ToLower(cryptocurrency) + "']"
	parseElementTry := ".flex .justify-between .mr-16"

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
		log.Fatalf("Status Code Error: %d %s", res.StatusCode, res.Status)
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)

	if err != nil {
		log.Fatal(err)
	}

	dollarValue := doc.Find(parseElementDollar).Text()
	dollarValue = strings.Replace(dollarValue, ".", "", 1)
	cryptocurrencys["dollar"] = strings.Replace(dollarValue, "$", "", 1)

	data := doc.Find(parseElementTry)
	data.Each(func(i int, s *goquery.Selection) {
		if i == 0 {
			tryValue := s.Find(".text-md").Text()
			tryValue = strings.Replace(tryValue, ".", "", 1)
			cryptocurrencys["try"] = strings.Replace(tryValue, "₺", "", 1)
		}
	})

	return cryptocurrencys
}
