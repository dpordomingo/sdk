package uast

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNodeTokens(t *testing.T) {
	require := require.New(t)

	f, err := getFixture("java_example_1.json")
	require.NoError(err)

	var c ToNoder = &BaseToNoder{
		InternalTypeKey: "internalClass",
		LineKey:         "line",
		OffsetKey:       "startPosition",
		TokenKeys: map[string]bool{
			"identifier":        true, // SimpleName
			"escapedValue":      true, // StringLiteral
			"keyword":           true, // Modifier
			"primitiveTypeCode": true, // ?
		},
	}
	n, err := c.ToNode(f)
	require.NoError(err)
	require.NotNil(n)

	tokens := Tokens(n)
	require.True(len(tokens) > 0)
}
