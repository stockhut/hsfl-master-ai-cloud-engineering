package main

import (
	"fmt"
	"github.com/stockhut/hsfl-master-ai-cloud-engineering/reverse-proxy"
	"github.com/stockhut/hsfl-master-ai-cloud-engineering/reverse-proxy/config"
	"log"
	"net/http"
	"os"
)

const Port int = 5000
const Host string = "localhost"

func main() {

	c, err := config.FromFile("config.yml")
	if err != nil {
		panic(err)
	}
	fmt.Println(c)

	logger := log.New(os.Stdout, "", 0)

	proxy := reverse_proxy.New(logger, c.Services)

	addr := fmt.Sprintf("%s:%d", Host, Port)
	log.Fatal(http.ListenAndServe(addr, proxy))
}
