package gcm

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFoo(t *testing.T) {

}

func getAddress(m *Mux) (string, error) {
	host, err := m.String("host")
	if err != nil {
		return "", err
	}

	port, err := m.Int("port")
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%s:%d", host, port), nil
}

func TestUsingTestMux(t *testing.T) {
	m := NewTestMux(map[string]interface{}{
		"host": "localhost",
		"port": 8080,
	})

	addr, err := getAddress(m)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, "localhost:8080", addr)
}
