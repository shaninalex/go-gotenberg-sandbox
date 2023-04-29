package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/gin-contrib/cache"
	"github.com/gin-contrib/cache/persistence"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	store := persistence.NewInMemoryStore(time.Minute)

	router.GET("/balancesheet", cache.CachePage(store, time.Minute, balancesheet))
	router.GET("/cashflow", cache.CachePage(store, time.Minute, cashflow))
	router.GET("/company_overview", cache.CachePage(store, time.Minute, company_overview))
	router.GET("/earnings_calendar", cache.CachePage(store, time.Minute, earnings_calendar))
	router.GET("/summary", cache.CachePage(store, time.Minute, summary))
	router.GET("/congresstrading", cache.CachePage(store, time.Minute, congresstrading))
	router.GET("/housetrading", cache.CachePage(store, time.Minute, housetrading))
	router.GET("/offexchange", cache.CachePage(store, time.Minute, offexchange))
	router.GET("/senatetrading", cache.CachePage(store, time.Minute, senatetrading))
	router.GET("/twitter", cache.CachePage(store, time.Minute, twitter))
	router.GET("/wallstreetbets", cache.CachePage(store, time.Minute, wallstreetbets))

	router.Run(":9010")
}

func readJSONFile(filename string) (json.RawMessage, error) {
	contents, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return json.RawMessage(contents), nil
}

func handleRequest(c *gin.Context, filename string) {
	rawMsg, err := readJSONFile(filename)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "cant get data"})
		return
	}
	var data interface{}
	if err := json.Unmarshal(rawMsg, &data); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "cant unmarshal filedata"})
		return
	}
	c.JSON(http.StatusOK, data)
}

func balancesheet(c *gin.Context) {
	handleRequest(c, "/source/a_balancesheet.json")
}
func cashflow(c *gin.Context) {
	handleRequest(c, "/source/a_cashflow.json")
}
func company_overview(c *gin.Context) {
	handleRequest(c, "/source/a_company_overview.json")
}
func earnings_calendar(c *gin.Context) {
	handleRequest(c, "/source/a_earnings_calendar.json")
}
func summary(c *gin.Context) {
	handleRequest(c, "/source/g_summary.json")
}
func congresstrading(c *gin.Context) {
	handleRequest(c, "/source/q_congresstrading.json")
}
func housetrading(c *gin.Context) {
	handleRequest(c, "/source/q_housetrading.json")
}
func offexchange(c *gin.Context) {
	handleRequest(c, "/source/q_offexchange.json")
}
func senatetrading(c *gin.Context) {
	handleRequest(c, "/source/q_senatetrading.json")
}
func twitter(c *gin.Context) {
	handleRequest(c, "/source/q_twitter.json")
}
func wallstreetbets(c *gin.Context) {
	handleRequest(c, "/source/q_wallstreetbets.json")
}
