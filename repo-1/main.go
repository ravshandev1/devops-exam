package main

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
)

type Message struct {
	Msg string `json:"message"`
}

type ConfigMessage struct {
	LogLevel    string `json:"log_level"`
	GRPCPort    string `json:"grpc_port"`
	Environment string `json:"environment"`
	PodIP       string `json:"pod_ip"`
	DBURL       string `json:"db_url"`
}

type DomainMessage struct {
	Endpoint string `json:"endpoint"`
	Domain   string `json:"domain"`
	LogLevel string `json:"log_level"`
}

type LoadCapabilityMessage struct {
	CPU    string `json:"cpus"`
	MEMORY string `json:"memorys"`
}

func getEnvVariable(key string) string {
	value := os.Getenv(key)
	return value
}

func message(c *gin.Context) {
	message := Message{
		Msg: "Welcome to your Kubernetes Exam!",
	}
	c.JSON(200, message)
}

func getHealth(c *gin.Context) {
	message := Message{
		health: "true",
	}
	c.JSON(200, message)
}

func getConfig(c *gin.Context) {

	logLevel := getEnvVariable("LOG_LEVEL")
	grpcPort := getEnvVariable("GRPC_PORT")
	environment := getEnvVariable("ENVIRONMENT")
	dbURL := getEnvVariable("DB_URL")
	podIP := getEnvVariable("POD_IP")

	message := ConfigMessage{
		LogLevel:    logLevel,
		GRPCPort:    grpcPort,
		Environment: environment,
		DBURL:       dbURL,
		PodIP:       podIP,
	}
	c.JSON(200, message)
}

func getDomain(c *gin.Context) {
	host := c.Request.Host
	fmt.Printf("Host: %s\n", host)

	logLevel := getEnvVariable("LOG_LEVEL")

	message := DomainMessage{
		Endpoint: "/domain",
		Domain:   host,
		LogLevel: logLevel,
	}

	c.JSON(200, message)
}

func getLoad(c *gin.Context) {
	cpus := getEnvVariable("CPU")
	memorys := getEnvVariable("MEMORY")

	message := LoadCapabilityMessage{
		CPU:    cpus,
		MEMORY: memorys,
	}

	c.JSON(200, message)
}

func main() {
	r := gin.Default()

	r.GET("/", message)
	r.GET("/config", getConfig)
	r.GET("/domain", getDomain)
	r.GET("/load-capability", getLoad)
	r.GET("/health", getHealth)

	port := ":80"
	fmt.Printf("Server is listening on port %s...\n", port)
	r.Run(port)
}
