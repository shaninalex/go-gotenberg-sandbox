package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"mime/multipart"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/", createPDF)
	router.Run(":8090")
}

func createPDF(c *gin.Context) {
	url := "http://gotenberg:3000/forms/chromium/convert/url"
	var b bytes.Buffer
	w := multipart.NewWriter(&b)

	// Add a form field with name "url" and value "http://urlGenerator:8080"
	fw, err := w.CreateFormField("url")
	if err != nil {
		fmt.Println(err)
		return
	}
	fw.Write([]byte("http://urlgenerator:8080"))

	// Close the form
	w.Close()

	req, err := http.NewRequest("POST", url, &b)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Something went wrong..."})
	}
	req.Header.Set("Content-Type", w.FormDataContentType())

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err})
	}
	defer resp.Body.Close()

	pdfData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.Data(http.StatusOK, "application/pdf", pdfData)
}
