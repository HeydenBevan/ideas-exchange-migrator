package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

const (
	sleepingState = true
)

type config struct {
	Keys   map[string]string
	Client *http.Client
}

func main() {
	c := setConfig()
	fmt.Println(c.Keys["client-timeout"])
}

func setConfig() *config {
	var c config
	c.Keys = keysFromEnv()

	timeout, err := time.ParseDuration(c.Keys["client-timeout"])
	if err != nil {
		log.Fatal(err)
	}

	c.Client = &http.Client{Timeout: timeout}

	return &c
}

const (
	keyName  = 0
	keyValue = 1
)

func keysFromEnv() map[string]string {
	m := make(map[string]string)
	for _, e := range environ() {
		kv := strings.Split(e, "=")
		key := kv[keyName]
		val := kv[keyValue]
		m[key] = val
	}
	return m
}

func environ() []string {
	var k []string
	f, err := os.Open(".env")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		ln := scanner.Text()
		if ln == "" {
			break
		}
		k = append(k, ln)
	}
	return k
}
