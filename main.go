package main

import (
	"V3/packages/config"
	"V3/packages/database"
	"V3/packages/routes"
	"log"
	"os"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	requestCount = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "http_requests_total",
			Help: "Total number of HTTP requests.",
		},
		[]string{"method", "endpoint"},
	)

	requestDuration = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name: "http_request_duration_seconds",
			Help: "Duration of HTTP requests in seconds.",
		},
		[]string{"method", "endpoint"},
	)
)

func init() {

	prometheus.MustRegister(requestCount)
	prometheus.MustRegister(requestDuration)
}
func main() {
	var err error
	config.EthClient, err = ethclient.Dial("https://rpc.ankr.com/bsc")
	if err != nil {
		log.Println(err)
	}
	database.DBSet()
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	router := gin.Default()
	router.GET("/metrics", gin.WrapH(promhttp.Handler()))

	// Register Prometheus metrics

	config := cors.DefaultConfig()
	config.AllowHeaders = append(config.AllowHeaders, "userid")
	config.AllowHeaders = append(config.AllowHeaders, "authorization")

	routes.Routes(router)

	log.Println("running")
	router.Run(":" + port)
}
