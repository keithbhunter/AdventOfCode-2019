package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSuccessfulPasswords(t *testing.T) {
	assert.Equal(t, 1135, numberOfSuccessfulPasswordsInRange(172851, 675869))
}

func Test111111MeetsCriteria(t *testing.T) {
	assert.False(t, passwordMeetsCriteria("111111"))
}

func Test223450MeetsCriteria(t *testing.T) {
	assert.False(t, passwordMeetsCriteria("223450"))
}

func Test123789MeetsCriteria(t *testing.T) {
	assert.False(t, passwordMeetsCriteria("123789"))
}

func Test112233MeetsCriteria(t *testing.T) {
	assert.True(t, passwordMeetsCriteria("112233"))
}

func Test123444MeetsCriteria(t *testing.T) {
	assert.False(t, passwordMeetsCriteria("123444"))
}

func Test111122MeetsCriteria(t *testing.T) {
	assert.True(t, passwordMeetsCriteria("111122"))
}

func Test223333MeetsCriteria(t *testing.T) {
	assert.True(t, passwordMeetsCriteria("223333"))
}

func Test788999MeetsCriteria(t *testing.T) {
	assert.True(t, passwordMeetsCriteria("788999"))
}

func Test111223MeetsCriteria(t *testing.T) {
	assert.True(t, passwordMeetsCriteria("111223"))
}

func Test111222MeetsCriteria(t *testing.T) {
	assert.False(t, passwordMeetsCriteria("111222"))
}
