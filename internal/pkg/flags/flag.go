package flags

import "flag"

type Flags struct {
    Path string
}

func Parse() *Flags {
    configPath := flag.String("config", "./config/config.yaml", "path to config file")
    flag.Parse()
    return &Flags{
        Path: *configPath,
    }
}
