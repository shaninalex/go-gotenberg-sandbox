package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
)

type DataGathering struct {
	CompanyOverview *CompanyOverviewPayload
	Wallstreetbets  *WallstreetbetsPayload
}

type WallstreetbetsResponse struct {
	Data   any `json:"data"`
	Status any `json:"status"`
}

type WallstreetbetsPayload struct {
	Date      string
	Sentiment float32
}

type CompanyOverviewPayload struct {
	Name   string
	Symbol string
}

type PdfPayload struct {
	PdfLogo            string
	CompanyName        string
	Ticker             string
	WallstreetbetsJson string
}

type Url struct {
	Link string
	Name string
}

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("./templates/*")

	router.GET("/", getData)
	router.Run(":8080")
}

func getData(c *gin.Context) {

	// GET DATA FROM DATASOURCE
	links := []Url{
		{Link: "http://localhost:9010/company_overview", Name: "CompanyOverview"},
		{Link: "http://localhost:9010/wallstreetbets", Name: "Wallstreetbets"},
	}

	var wg sync.WaitGroup
	wg.Add(len(links))

	var dataGathering DataGathering

	for _, link := range links {
		go func(link Url) {
			defer wg.Done()
			response, err := http.Get(link.Link)
			if err != nil {
				log.Printf("Error downloading %s: %s\n", link.Name, err.Error())
				return
			}
			defer response.Body.Close()
			bodyBytes, err := io.ReadAll(response.Body)
			if err != nil {
				log.Fatal(err)
			}
			if link.Name == "CompanyOverview" {
				var companyOverview CompanyOverviewPayload
				if err := json.Unmarshal(bodyBytes, &companyOverview); err != nil {
					log.Println(err)
					return
				}
				dataGathering.CompanyOverview = &companyOverview
			}

			if link.Name == "Wallstreetbets" {
				var wallstreetbets WallstreetbetsResponse
				if err := json.Unmarshal(bodyBytes, &wallstreetbets); err != nil {
					log.Println(err)
					return
				}
				var a WallstreetbetsPayload
				if err := json.Unmarshal(wallstreetbets.Data, &a); err != nil {
					log.Println(err)
					return
				}
				dataGathering.Wallstreetbets = &a
			}

			// if link.Name == "Balancesheet" {
			// 	dataGathering.Balancesheet = &decodedData
			// }

			// if link.Name == "Cashflow" {
			// 	dataGathering.Cashflow = &decodedData
			// }

			log.Printf("Downloaded %s\n", link)
		}(link)
	}

	wg.Wait()

	pdfPayload := PdfPayload{
		PdfLogo: "http://localhost:9020/images/logo.png",
	}

	if dataGathering.CompanyOverview != nil {
		pdfPayload.CompanyName = dataGathering.CompanyOverview.Name
		pdfPayload.Ticker = dataGathering.CompanyOverview.Symbol
	}

	if dataGathering.Wallstreetbets != nil {
		b, _ := json.Marshal(dataGathering.Wallstreetbets)
		pdfPayload.WallstreetbetsJson = string(b)
		log.Println(pdfPayload.WallstreetbetsJson)
	}
	c.HTML(http.StatusOK, "index.tmpl", pdfPayload)
}
