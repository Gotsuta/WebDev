package api

import (
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

var services = []Service{
	{
		Id:    1,
		Image: "/image/White.jpg",
		Name:  "White",
		Gb:    10,
		Min:   100,
		Sms:   100,
		Price: 100,
	},
	{
		Id:    2,
		Image: "/image/Gray.jpg",
		Name:  "Gray",
		Gb:    20,
		Min:   200,
		Sms:   200,
		Price: 200,
	},
	{
		Id:    3,
		Image: "/image/Black.jpg",
		Name:  "Black",
		Gb:    30,
		Min:   300,
		Sms:   300,
		Price: 300,
	},
	{
		Id:    4,
		Image: "/image/Total black.jpg",
		Name:  "Total Black",
		Gb:    40,
		Min:   400,
		Sms:   400,
		Price: 400,
	},
}

func StartServer() {
	log.Println("Server start up")

	r := gin.Default()
	r.LoadHTMLGlob("templates/*")

	r.GET("/", func(c *gin.Context) {
		searchQuery := c.DefaultQuery("fsearch", "")
		var result []Service

		for _, service := range services {
			if strings.Contains(strings.ToLower(service.Name), strings.ToLower(searchQuery)) {
				result = append(result, service)
			}
		}

		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"services": result,
			"fsearch":  searchQuery,
		})

	})

	r.GET("/service/:id", func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))

		if err != nil {
			// Обработка ошибки
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid id"})
			return
		}

		service := services[id-1]
		c.HTML(http.StatusOK, "rate.tmpl", service)
	})
	r.Static("/image", "./resources/image")
	r.Static("/css", "./resources/css")

	r.Run(":8081") // listen and serve on 0.0.0.0:8081 (for windows "localhost:8081")

	log.Println("Server down")
}
