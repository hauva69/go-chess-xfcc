package xfcc

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOngoing(t *testing.T) {
	expected := "*"
	var result Result = "Ongoing"
	assert.Equal(t, expected, result.PGN())
}

func TestWhiteWins(t *testing.T) {
	expected := "1-0"
	var result Result = "WhiteWins"
	assert.Equal(t, expected, result.PGN())
}

func TestBlackWins(t *testing.T) {
	expected := "0-1"
	var result Result = "BlackWins"
	assert.Equal(t, expected, result.PGN())
}

func TestWhiteWinAbjudicated(t *testing.T) {
	expected := "1-0"
	var result Result = "WhiteWinAbjudicated"
	assert.Equal(t, expected, result.PGN())
}

func TestBlackWinAbjudicated(t *testing.T) {
	expected := "0-1"
	var result Result = "BlackWinAbjudicated"
	assert.Equal(t, expected, result.PGN())
}

func TestDrawAbjudicated(t *testing.T) {
	expected := "1/2-1/2"
	var result Result = "DrawAbjudicated"
	assert.Equal(t, expected, result.PGN())
}

func TestWhiteDefaulted(t *testing.T) {
	expected := "0-1"
	var result Result = "WhiteDefaulted"
	assert.Equal(t, expected, result.PGN())
}

func TestBothDefaulted(t *testing.T) {
	expected := "*"
	var result Result = "BothDefaulted"
	assert.Equal(t, expected, result.PGN())
}

func TestCancelled(t *testing.T) {
	expected := "*"
	var result Result = "Cancelled"
	assert.Equal(t, expected, result.PGN())
}

func TestAbjudicationPending(t *testing.T) {
	expected := "*"
	var result Result = "AbjudicationPending"
	assert.Equal(t, expected, result.PGN())
}

func TestDraw(t *testing.T) {
	expected := "1/2-1/2"
	var result Result = "Draw"
	assert.Equal(t, expected, result.PGN())
}
