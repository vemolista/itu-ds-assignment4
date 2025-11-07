package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"os/signal"

	"github.com/vemolista/itu-ds-assignment4/node"
)

func main() {
	index := flag.Int("i", 0, "Zero-based index of the node to start (based on the array of defined nodes in config.json)")
	flag.Parse()

	fmt.Printf("starting node with id %d\n", *index)

	configFile, err := os.Open("config.json")
	if err != nil {
		log.Fatalf("cannot open config file: %v", err)
	}
	defer configFile.Close()

	var config node.Config
	bytes, err := io.ReadAll(configFile)
	if err != nil {
		log.Fatalf("cannot read bytes: %v", err)
	}
	json.Unmarshal(bytes, &config)

	n, _ := node.NewNode(*index, &config)
	n.Start()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c
}
