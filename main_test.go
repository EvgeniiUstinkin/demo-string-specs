package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestTestValidity(t *testing.T) {
	t.Run("valid string", func(t *testing.T) {
		require.Equal(t, true, testValidity("23-ab-48-caba-56-haha"))
	})
	t.Run("invalid string", func(t *testing.T) {
		require.Equal(t, false, testValidity("23 ab 48 caba 56 haha"))
	})
	t.Run("empty string", func(t *testing.T) {
		require.Equal(t, false, testValidity(""))
	})
}

func TestAverageNumber(t *testing.T) {
	t.Run("string without numbers", func(t *testing.T) {
		require.Equal(t, uint16(0), averageNumber("oo-ab-bb-caba-pp-haha"))
	})
	t.Run("string with numbers", func(t *testing.T) {
		require.Equal(t, uint16(42), averageNumber("23-ab-48-caba-56-haha"))
	})
}

func TestWholeStory(t *testing.T) {
	t.Run("empty string", func(t *testing.T) {
		require.Equal(t, "", wholeStory(""))
	})
	t.Run("string without numbers", func(t *testing.T) {
		require.Equal(t, "oo ab bb caba pp haha", wholeStory("oo-ab-bb-caba-pp-haha"))
	})
	t.Run("string with numbers", func(t *testing.T) {
		require.Equal(t, "ab caba haha", wholeStory("23-ab-48-caba-56-haha"))
	})
}

func TestStoryStat(t *testing.T) {

}
