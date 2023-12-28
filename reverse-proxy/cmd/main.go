package main

import (
	"fmt"
	"github.com/stockhut/hsfl-master-ai-cloud-engineering/common/environment"
	"github.com/stockhut/hsfl-master-ai-cloud-engineering/reverse-proxy"
	"github.com/stockhut/hsfl-master-ai-cloud-engineering/reverse-proxy/config"
	"log"
	"net/http"
	"os"
)

const Port int = 5000
const Host string = ""

const ConfigFileEnvKey = "CONFIG_FILE"

func main() {

	logger := log.New(os.Stdout, "", 0)

	configPath := environment.GetRequiredEnvVar(ConfigFileEnvKey)

	c, err := config.FromFile(configPath)
	if err != nil {
		log.Fatalf("Failed to load configuration file '%s': %s", configPath, err)
	}
	prettyPrintConfig(logger, c)

	proxy := reverse_proxy.New(logger, c.Services)

	addr := fmt.Sprintf("%s:%d", Host, Port)

	logger.Printf("Listening on %s", addr)
	log.Fatal(http.ListenAndServe(addr, proxy))
}
func prettyPrintConfig(logger *log.Logger, c config.Config) {
	logger.Printf("Loaded %d service configurations:\n", len(c.Services))
	for _, s := range c.Services {
		logger.Printf("%s: %s ➡️ %s\n", s.Name, s.Route, s.TargetHost)
	}
}
