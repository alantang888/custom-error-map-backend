package main

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
)

type ErrorMapping struct {
	ErrorMapping map[int]int `yaml:"error_mapping"`
}

func readMapping(mapping *ErrorMapping) {
	configPath := os.Getenv("CONFIG_PATH")

	// Set default value if env var is empty
	if configPath == "" {
		configPath = "error_mapping.yaml"
	}

	configByte, err := ioutil.ReadFile(configPath)
	if err != nil {
		fmt.Printf("Can't read config file: %v\n", err)
		return
	}

	err = yaml.Unmarshal(configByte, mapping)
	if err != nil {
		fmt.Printf("Can't prase config YAML: %v\n", err)
		return
	}
}

func main() {
	var mapping ErrorMapping
	readMapping(&mapping)

	listenPortEnv := os.Getenv("LISTEN_PORT")
	listenPort, err := strconv.Atoi(listenPortEnv)
	if err != nil {
		fmt.Printf("Can't parse \"%s\", use default 8080.\n", listenPortEnv)
		listenPort = 8080
	}

	defaultReturnEnv := os.Getenv("DEFAULT_RETURN")
	defaultReturn, err := strconv.Atoi(defaultReturnEnv)
	if err != nil {
		fmt.Printf("Can't parse \"%s\", use default 404.\n", defaultReturnEnv)
		defaultReturn = 404
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("Requested header: %s\n", r.Header)

		returnStatus := defaultReturn

		originalErrorCodeSlice, ok := r.Header["X-Code"]
		if ok {
			if len(originalErrorCodeSlice) == 1 {
				originalErrorCode, err := strconv.Atoi(originalErrorCodeSlice[0])
				if err == nil {
					newErrorCode, ok := mapping.ErrorMapping[originalErrorCode]
					if ok {
						returnStatus = newErrorCode
					} else {
						returnStatus = originalErrorCode
					}
				}
			}
		}

		w.WriteHeader(returnStatus)
		w.Write([]byte(fmt.Sprintf("Error %d!!", returnStatus)))
	})

	fmt.Printf("Server started on port %d\n", listenPort)
	http.ListenAndServe(fmt.Sprintf(":%d", listenPort), nil)
}
