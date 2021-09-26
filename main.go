package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/chickenzord/reqstash/pkg/reqstash"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

func main() {
	godotenv.Overload()
	var cfg reqstash.Config
	if err := envconfig.Process("reqstash", &cfg); err != nil {
		panic(err)
	}
	fmt.Println(cfg)

	s := reqstash.MemoryStorage{
		Capacity: 2,
		TTL:      10 * time.Second,
	}

	srv := gin.Default()
	srv.GET("/", func(c *gin.Context) {
		reqs, err := s.ListAll()
		if err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"requests": reqs,
		})
	})
	srv.Any("/record", func(c *gin.Context) {
		req := reqstash.NewRequest(c.Request)
		s.Put(req)

		c.JSON(http.StatusAccepted, gin.H{
			"request": req,
		})
	})

	if err := http.ListenAndServe(":8080", srv); err != nil {
		panic(err)
	}
}
