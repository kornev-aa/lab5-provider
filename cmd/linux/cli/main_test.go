package main

import (
    "os"
    "testing"
    "github.com/stretchr/testify/assert"
)

func TestRunAppWithValidConfig(t *testing.T) {
    // Создаём временный конфиг
    tmpFile, err := os.CreateTemp("", "config*.yaml")
    assert.NoError(t, err)
    defer os.Remove(tmpFile.Name())

    configContent := `
service:
  provider:
    type: open-meteo
  location:
    lat: 53.6688
    long: 23.8223
`
    _, err = tmpFile.WriteString(configContent)
    assert.NoError(t, err)
    tmpFile.Close()

    err = RunApp(tmpFile.Name())
    assert.NoError(t, err)
}

func TestRunAppWithInvalidConfig(t *testing.T) {
    tmpFile, err := os.CreateTemp("", "invalid*.yaml")
    assert.NoError(t, err)
    defer os.Remove(tmpFile.Name())

    _, err = tmpFile.WriteString("invalid: yaml: [")
    assert.NoError(t, err)
    tmpFile.Close()

    err = RunApp(tmpFile.Name())
    assert.Error(t, err)
}

func TestRunAppWithNonExistentFile(t *testing.T) {
    err := RunApp("/nonexistent/file.yaml")
    assert.Error(t, err)
}
