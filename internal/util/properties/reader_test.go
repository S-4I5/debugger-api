package properties

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestReadProperties(t *testing.T) {
	actual, err := ReadProperties(".//test//testReadProperties.properties")

	expected := map[string]string{
		"something.like.this":      "and like this",
		"something.like.this.also": "and like this also",
		"something.like.this.not":  "and like this also",
	}

	require.NoError(t, err)
	require.Equal(t, expected, actual)
}

func TestReadPropertiesWithIncorrectPropertiesFormat(t *testing.T) {
	_, err := ReadProperties(".//test2//TestReadPropertiesWithIncorrectPropertiesFormat.properties")

	require.Error(t, err)
}
