package main

import (
	"testing"
	"github.com/stretchr/testify/require"
)

func TestChecksum(t *testing.T){
	require.Equal(t, 18, CalculateChecksum("5\t1\t9\t5",
		"7\t5\t3",
		"2\t4\t6\t8"))
}

func TestGetEvenDivisors(t *testing.T){
	require.Equal(t, 4, GetEvenDivisors("5\t9\t2\t8"))
	require.Equal(t, 3,	GetEvenDivisors("9\t4\t7\t3"))
	require.Equal(t, 2,	GetEvenDivisors("3\t8\t6\t5"))
}

func TestGetEvenDivisorsSum(t *testing.T){
	require.Equal(t, 9, GetEvenDivisorsSum("5\t9\t2\t8",
		"9\t4\t7\t3",
		"3\t8\t6\t5"))
}