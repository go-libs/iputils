package iputils

import (
	"github.com/bmizerany/assert"
	"testing"
)

func TestIP2Long(t *testing.T) {
	var ip uint32 = 2346851510
	str := IP2Long("139.226.28.182")
	assert.Equal(t, ip, str)
}

func TestLong2IP(t *testing.T) {
	ip := Long2IP(134744072)
	assert.Equal(t, ip, "8.8.8.8")
}
