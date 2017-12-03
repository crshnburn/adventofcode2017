package main

import (
	"testing"
	"github.com/stretchr/testify/require"
)

func Test1122(t *testing.T) {
	require.Equal(t, 3, ReverseCaptcha("1122"), "Captcha incorrect")
}

func Test1111(t *testing.T) {
	require.Equal(t, 4, ReverseCaptcha("1111"), "Captcha incorrect")
}

func Test1234(t *testing.T) {
	require.Equal(t, 0, ReverseCaptcha("1234"), "Captcha incorrect")
}

func Test91212129(t *testing.T) {
	require.Equal(t, 9, ReverseCaptcha("91212129"), "Captcha incorrect")
}

func Test1212(t *testing.T){
	require.Equal(t, 6, ReverseCaptcha2("1212"), "Captcha incorrect")
}

