package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHandleRequests_WhenPing(t *testing.T) {
	// Given
	request := "*1\n$4\nping\n"

	// When
	result, _ := HandleRequests(request)

	// Then
	assert.Equal(t, []byte("+PONG\r\n"), result)
}

func TestHandleRequests_ReturnsAnErrorForWrongCmdIfNotMatchingWhatExpected(t *testing.T) {
	// Given
	request := "*1\n$4\nbla\n"

	// When
	_, err := HandleRequests(request)

	// Then
	assert.Equal(t, "command not accepted", err.Error())
}
