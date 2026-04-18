package config

import (
    "strings"
    "testing"
    "github.com/stretchr/testify/assert"
    "github.com/stretchr/testify/require"
)

func TestParseValidConfig(t *testing.T) {
    yamlData := `
service:
  provider:
    type: open-meteo
  location:
    lat: 53.6688
    long: 23.8223
`
    r := strings.NewReader(yamlData)
    cfg, err := Parse(r)
    require.NoError(t, err)

    assert.Equal(t, "open-meteo", cfg.P.Type)
    assert.Equal(t, 53.6688, cfg.L.Lat)
    assert.Equal(t, 23.8223, cfg.L.Long)
}

func TestParseInvalidYAML(t *testing.T) {
    invalidYAML := `service: [provider: type:`
    r := strings.NewReader(invalidYAML)
    _, err := Parse(r)
    assert.Error(t, err)
}
