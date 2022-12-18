package helpers_test

import (
	"testing"

	"snaLinkShortener/pkg/helpers"

	"github.com/stretchr/testify/require"
)

func TestGetConfig(t *testing.T) {
	config := helpers.GetConfig()
	t.Setenv("SHORT_LINK_BASE", "http://localhost:15001/")
	t.Setenv("ALLOW_CORS_FROM", "localhost:3000")
	t.Setenv("POSTGRES_CONN_STRING", "postgres://user:password@localhost:5432/linkShortener")
	require.Equal(t, "localhost:3000", config.CorsAllowFrom)
	require.Equal(t, "http://localhost:15001/", config.ShortLinkBase)
	require.Equal(t, "postgres://user:password@localhost:5432/linkShortener", config.DbConnString)
}
