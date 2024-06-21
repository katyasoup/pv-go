package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

var pvBaseUrl = "https://phinvads.cdc.gov/baseStu3"

type App struct {
	Port   string
	Router *gin.Engine
}

func main() {
	s := NewServer()
	s.setUpRoutes()

	r := s.Router
	r.Run(s.Port) // Default if blank is :8080
}

// NewServer creates a new instance of App with defined Port and Router
func NewServer() *App {
	server := &App{
		Port:   ":8000",
		Router: gin.Default(),
	}

	return server
}

// setUpRoutes defines the endpoints, attaches them to a server, and starts listening and serving HTTP requests
func (a *App) setUpRoutes() {
	r := a.Router

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message:": "pong",
		})
	})

	pv := r.Group("phinvads")

	pv.GET("/ValueSet/:id", func(c *gin.Context) {
		url := formatUrl(c)
		result, err := get(url, c)

		if err != nil {
			handleErrorResponse(err)
		}

		handleJsonResponse(result, c)
	})
}

func formatUrl(c *gin.Context) string {
	id := c.Param("id")
	path := c.Request.URL
	path.Opaque = "opaque"

	url := fmt.Sprintf("%s/ValueSet/%s", pvBaseUrl, id)

	queryParams := strings.Split(c.Request.URL.String(), "opaque")[1]

	if queryParams != "" {
		url += queryParams
	}

	return url
}

func get(url string, c *gin.Context) (result PhinVadsResponse, err error) {
	resp, err := http.Get(url)
	checkResponse(err, resp, c)

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	checkResponse(err, resp, c)

	respType := resp.Header.Get("Content-Type")

	if respType == "application/json" {
		return unmarshalResponse(body)

	} else {
		result := string(body[:])
		return result, nil
	}

}

func checkResponse(err error, resp *http.Response, c *gin.Context) {
	if err != nil {
		c.JSON(resp.StatusCode, handleErrorResponse(err))
		return
	}
}

func unmarshalResponse(b []byte) (result PhinVadsResponse, err error) {
	var data PhinVadsResponse
	err = json.Unmarshal(b, &data)

	if err != nil {
		return nil, err
	}

	return data, nil
}

func handleErrorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}

func handleJsonResponse(result PhinVadsResponse, c *gin.Context) {
	c.Header("Content-Type", "application/json")
	c.JSON(http.StatusOK, result)
}

type PhinVadsResponse interface {
	any
}
