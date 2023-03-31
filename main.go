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
	"os"
)

// Declare default environment variables
var (
	// AppName is the name of the application
	AppName = "http-server"
	// AppVersion is the version of the application
	AppVersion = "0.1.0"
	// AppCommit is the commit hash of the application
	AppCommit = "unknown"
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
}

func getRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("got / request\n")
	io.WriteString(w, fmt.Sprintf(
		"Hello! I'm the %s service, serving version %s from commit %s\n",
		AppName, AppVersion, AppCommit))
}
func getHello(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("got /hello request\n")
	io.WriteString(w, "Hello, HTTP!\n")
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
