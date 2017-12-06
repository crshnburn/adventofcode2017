package main

import (
	"testing"
	"github.com/stretchr/testify/require"
)

func TestMemoryWalk(t *testing.T) {
	require.Equal(t, 5, MemoryWalk([]int{0, 3, 0, 1, -3}))
}

func TestMemoryWalk2(t *testing.T) {
	require.Equal(t, 10, MemoryWalk2([]int{0, 3, 0, 1, -3}))
}