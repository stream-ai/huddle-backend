package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestIntegration(t *testing.T) {
	// Start the server in a separate goroutine
	go func() {
		err := run(context.Background(), nil, nil)
		if err != nil {
			fmt.Fprintf(os.Stderr, "error running server: %s\n", err)
		}
	}()

	// Wait for the server to start
	time.Sleep(1 * time.Second)

	// path, expectedStatusCode, expectedBody
	// Make a request to the server
	resp, err := http.Get("http://localhost:80/healthz")
	assert.NoError(t, err)
	defer resp.Body.Close()
	// Read the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("error reading response body: %v", err)
	}
	assert.Equal(t, "OK", string(body))
	// Assert the response status code
	assert.Equal(t, http.StatusOK, resp.StatusCode)
}
