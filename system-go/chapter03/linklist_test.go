package chapter03

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAdd(t *testing.T) {
	var head = new(Node)

	add(head, 10)
	assert.NotNil(t, head)
	assert.Equal(t, 10, head.Value)
	add(head, 20)
	assert.Equal(t, 10, head.Next.Value)
}
