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
	Response interface{}
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
		body, err := ioutil.ReadAll(c.Request.Body)
		if err != nil {
			panic(err)
		}
		fmt.Println(string(body))

		apiResp := APIResponse{}
		err = json.Unmarshal(body, &apiResp)
		if err != nil {
			panic(err)
		}
		fmt.Println(apiResp.Response)

		c.HTML(http.StatusOK, "index.tmpl.html", gin.H{
			"response": apiResp,
		})
	})

	router.GET("/callback", func(c *gin.Context) {
		body, err := ioutil.ReadAll(c.Request.Body)
		if err != nil {
			panic(err)
		}
		fmt.Println(string(body))

		apiResp := APIResponse{}
		err = json.Unmarshal(body, &apiResp)
		if err != nil {
			panic(err)
		}
		fmt.Println(apiResp.Response)

		c.HTML(http.StatusOK, "index.tmpl.html", gin.H{
			"response": apiResp,
		})
	})

	router.Run(":" + port)
}
