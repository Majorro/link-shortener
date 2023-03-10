package helpers_test

import (
	"testing"

	"snaLinkShortener/pkg/helpers"

	"github.com/stretchr/testify/require"
)

func TestGetConfig(t *testing.T) {
	config := helpers.GetConfig()
	t.Setenv("SHORT_LINK_BASE", "http://localhost:15001/")
	t.Setenv("ALLOW_CORS_FROM", "http://localhost:3000")
	t.Setenv("POSTGRES_CONN_STRING", "postgres://user:password@localhost:5432/linkShortener")
	require.Regexp(t, `http://\w+:\d+/`, config.ShortLinkBase)
	require.Regexp(t, `postgres://\w+:\w+@\w+:\d+/\w+`, config.DbConnString)
	require.Regexp(t, `http://\w+:\d+`, config.CorsAllowFrom)
}
