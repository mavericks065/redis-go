package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHandleRequests_WhenPing(t *testing.T) {
	// Given
	request := "*1\r\n$4\r\nping\r\n"

	// When
	result, _ := HandleRequests(request)

	// Then
	assert.Equal(t, []byte("+PONG\r\n"), result)
}

func TestHandleRequests_WhenHandlingCommand(t *testing.T) {
	// Given
	request := "*2\r\n$7\r\nCOMMAND\r\n$4\r\nDOCS"

	// When
	result, _ := HandleRequests(request)

	// Then
	assert.Equal(t, []byte("+PONG\r\n"), result)
}

func TestHandleRequests_ReturnsAnErrorForWrongCmdIfNotMatchingWhatExpected(t *testing.T) {
	// Given
	request := "*1\r\n$4\r\nbla\r\n"

	// When
	_, err := HandleRequests(request)

	// Then
	assert.Equal(t, "command not accepted", err.Error())
}
