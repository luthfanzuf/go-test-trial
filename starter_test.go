package starter

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSayHello(t *testing.T) {
	greeting := SayHello("William")
	assert.Equal(t, "Hello William. Welcome!", greeting)

}
