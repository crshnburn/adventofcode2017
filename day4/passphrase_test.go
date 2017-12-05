package main

import (
	"testing"
	"github.com/stretchr/testify/require"
)

func TestCheckPassPhrase(t *testing.T){
	require.Equal(t, true, CheckPassPhrase("aa bb cc dd ee"))
	require.False(t, CheckPassPhrase("aa bb cc dd aa"))
	require.True(t, CheckPassPhrase("aa bb cc dd aaa"))
}