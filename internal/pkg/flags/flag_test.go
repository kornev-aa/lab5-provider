package flags

import (
    "flag"
    "os"
    "testing"
    "github.com/stretchr/testify/assert"
)

func TestParseDefaultConfigPath(t *testing.T) {
    resetFlags()
    
    oldArgs := os.Args
    defer func() { os.Args = oldArgs }()

    os.Args = []string{"cmd"}
    flags := Parse()
    assert.Equal(t, "./config/config.yaml", flags.Path)
}

func TestParseCustomConfigPath(t *testing.T) {
    resetFlags()
    
    oldArgs := os.Args
    defer func() { os.Args = oldArgs }()

    os.Args = []string{"cmd", "-config", "/custom/path.yaml"}
    flags := Parse()
    assert.Equal(t, "/custom/path.yaml", flags.Path)
}

func resetFlags() {
    flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.ExitOnError)
}
