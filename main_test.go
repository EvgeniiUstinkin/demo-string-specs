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

}

func TestWholeStory(t *testing.T) {

}

func TestStoryStat(t *testing.T) {

}
