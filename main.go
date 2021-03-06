package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/heroku/x/hmetrics/onload"
)

type APIResponse struct {
	Response interface{} `json:"response"`
}

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("$PORT must be set")
	}

	router := gin.New()
	router.Use(gin.Logger())
	router.LoadHTMLGlob("templates/*.tmpl.html")
	router.Static("/static", "static")

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl.html", nil)
	})

	router.GET("/callback", func(c *gin.Context) {
		log.Println("req body", c.Request.Body)
		// log.Println("req.resp body", c.Request.Response.Body)
		body, err := ioutil.ReadAll(c.Request.Body)
		if err != nil {
			log.Panic("ioutil read error: ", err)
		}
		log.Println("body (string) :", string(body))
		log.Println("body (byte) :", body)

		apiResp := APIResponse{}

		err = json.Unmarshal(body, &apiResp)
		if err != nil {
			log.Panic("unmarshal error: ", err)
		}
		fmt.Println(apiResp.Response)
		log.Println(apiResp.Response)

		c.JSON(http.StatusOK, apiResp.Response)
		// c.HTML(http.StatusOK, "response.tmpl.html", gin.H{
		// 	"response": apiResp,
		// })
	})

	http.HandleFunc("/callback2", func(w http.ResponseWriter, r *http.Request) {
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Panic("ioutil read error: ", err)
		}
		log.Println("body (string) :", string(body))

		apiResp := APIResponse{}
		err = json.Unmarshal(body, &apiResp)
		if err != nil {
			log.Panic("unmarshal error: ", err)
		}
		fmt.Println(apiResp.Response)
		log.Println(apiResp.Response)
	})

	router.Run(":" + port)
}
