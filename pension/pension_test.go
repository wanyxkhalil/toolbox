package pension

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBasicPension(t *testing.T) {
	assert.Equal(t, 720., basicPension(6000, 0.6, 15))
	assert.Equal(t, 990., basicPension(6000, 1.2, 15))
}

func TestPrivatePension(t *testing.T) {
	assert.Equal(t, 517.92, privatePension(5000, 15, 139))
}
