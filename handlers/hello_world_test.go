package handlers_test

import (
	"testing"

	"github.com/keaaa/hello-world/handlers"
)

func Test_hello_world_output(t *testing.T) {
	msg := handlers.GetResponse()

	if msg != "Hello world!" {
		t.Error()
	}
}
