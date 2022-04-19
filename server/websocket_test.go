package server_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/mu-box/mist/auth"
	"github.com/mu-box/mist/server"
)

// TestWSStart tests to ensure a server will start
func TestWSStart(t *testing.T) {
	fmt.Println("Starting WS test...")

	// ensure authentication is disabled
	auth.Start("")

	go func() {
		if err := server.Start([]string{"ws://127.0.0.1:8888"}, ""); err != nil {
			t.Fatalf("Unexpected error - %s", err.Error())
		}
	}()
	<-time.After(time.Second)
}

// TestWSSStart tests to ensure a server will start
func TestWSSStart(t *testing.T) {
	fmt.Println("Starting WSS test...")

	// ensure authentication is disabled
	auth.Start("")

	go func() {
		if err := server.Start([]string{"wss://127.0.0.1:8988"}, ""); err != nil {
			t.Fatalf("Unexpected error - %s", err.Error())
		}
	}()
	<-time.After(time.Second)
}
