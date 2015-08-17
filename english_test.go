package english_test

import (
	"testing"

	"github.com/bfontaine/english"
	"github.com/stretchr/testify/assert"
)

func TestBool(t *testing.T) {
	assert.Equal(t, "yes", english.Bool(true))
	assert.Equal(t, "no", english.Bool(false))
}

func TestOrdinalLiteral(t *testing.T) {
	assert.Equal(t, "0th", english.OrdinalLiteral(0))

	assert.Equal(t, "1st", english.OrdinalLiteral(1))
	assert.Equal(t, "2nd", english.OrdinalLiteral(2))
	assert.Equal(t, "3rd", english.OrdinalLiteral(3))
	assert.Equal(t, "4th", english.OrdinalLiteral(4))

	assert.Equal(t, "11th", english.OrdinalLiteral(11))
	assert.Equal(t, "12th", english.OrdinalLiteral(12))
	assert.Equal(t, "13th", english.OrdinalLiteral(13))
	assert.Equal(t, "14th", english.OrdinalLiteral(14))

	assert.Equal(t, "811th", english.OrdinalLiteral(811))
	assert.Equal(t, "912th", english.OrdinalLiteral(912))

	assert.Equal(t, "942nd", english.OrdinalLiteral(942))
	assert.Equal(t, "12345th", english.OrdinalLiteral(12345))
	assert.Equal(t, "123451st", english.OrdinalLiteral(123451))

	assert.Equal(t, "-1st", english.OrdinalLiteral(-1))
	assert.Equal(t, "-2nd", english.OrdinalLiteral(-2))
	assert.Equal(t, "-3rd", english.OrdinalLiteral(-3))
	assert.Equal(t, "-4th", english.OrdinalLiteral(-4))
}
