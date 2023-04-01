/*
Copyright 2023.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package main

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"strings"
)

// Declare default environment variables
var (
	// AppName is the name of the application
	AppName = "http-server"
	// AppVersion is the version of the application
	AppVersion = "0.1.0"
	// AppCommit is the commit hash of the application
	AppCommit = "unknown"
	// AppBackendServices is a comma-separated list of backend services
	AppBackendServices []string = []string{"localhost:8080"}
)

// Read environment variables
func readEnv() {
	if n := os.Getenv("APP_NAME"); n != "" {
		AppName = n
	}
	if v := os.Getenv("APP_VERSION"); v != "" {
		AppVersion = v
	}
	if c := os.Getenv("APP_COMMIT"); c != "" {
		AppCommit = c
	}
	if b := os.Getenv("APP_BACKEND_SERVICES"); b != "" {
		b := strings.Split(b, ",")
		AppBackendServices = b
	}
}

var client http.Client

var HelloMessage string = fmt.Sprintf(
	"Hello from the %s service with version %s from commit %s\n",
	AppName, AppVersion, AppCommit)

func getRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("got / request\n")

	io.WriteString(w, fmt.Sprintf(
		"This html is served by the %s service! It fetches data from the backend services.\n\n", AppName))

	for _, backend := range AppBackendServices {
		fmt.Printf("fetching %s\n", backend)
		url, err := url.Parse(fmt.Sprintf("http://%s/hello", backend))
		if err != nil {
			fmt.Println(err)
		}
		resp, err := client.Get(url.String())
		if err != nil {
			fmt.Printf("error fetching %s : %s\n", url.String(), err)
		}
		defer resp.Body.Close()
		
		if resp.StatusCode == http.StatusOK {
			bodyBytes, err := io.ReadAll(resp.Body)
			if err != nil {
				fmt.Println(err)
			}
			bodyString := string(bodyBytes)
			io.WriteString(w, bodyString)
		}
	}
}

func getHello(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("got /hello request\n")
	io.WriteString(w, HelloMessage)
}

func main() {
	readEnv()

	http.HandleFunc("/", getRoot)
	http.HandleFunc("/hello", getHello)

	err := http.ListenAndServe(":8080", nil)
	if errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("server closed\n")
	} else if err != nil {
		fmt.Printf("error starting server: %s\n", err)
		os.Exit(1)
	}
}
