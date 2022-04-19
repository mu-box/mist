package server_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/mu-box/mist/auth"
	"github.com/mu-box/mist/server"
)

// TestTCPStart tests to ensure a server will start
func TestTCPStart(t *testing.T) {
	fmt.Println("Starting TCP test...")

	// ensure authentication is disabled
	auth.Start("")

	go func() {
		if err := server.Start([]string{"tcp://127.0.0.1:1445"}, ""); err != nil {
			t.Fatalf("Unexpected error - %s", err.Error())
		}
	}()
	<-time.After(time.Second)
}
