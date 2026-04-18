package main

import (
    "os"
    "time"
    pogodaby "github.com/kornev-aa/lab5-provider/internal/adapters/pogoda_by"
    "github.com/kornev-aa/lab5-provider/internal/adapters/weather"
    "github.com/kornev-aa/lab5-provider/internal/pkg/app/cli"
    "github.com/kornev-aa/lab5-provider/internal/pkg/flags"
    "github.com/kornev-aa/lab5-provider/pkg/cache"
    "github.com/kornev-aa/lab5-provider/pkg/config"
    "github.com/kornev-aa/lab5-provider/pkg/logger"
)

func getProvider(cfg config.Config, log *logger.Logger, cache cache.Cache, ttl time.Duration) cli.WeatherInfo {
    switch cfg.P.Type {
    case "open-meteo":
        return weather.New(log, cache, ttl)
    case "pogoda":
        return pogodaby.New(log)
    default:
        return weather.New(log, cache, ttl)
    }
}

func RunApp(configPath string) error {
    r, err := os.Open(configPath)
    if err != nil {
        return err
    }
    defer r.Close()

    cfg, err := config.Parse(r)
    if err != nil {
        return err
    }

    log := logger.New()
    memCache := cache.NewMemoryCache()
    cacheTTL := 5 * time.Minute
    wi := getProvider(cfg, log, memCache, cacheTTL)
    app := cli.New(log, wi, cfg)

    return app.Run()
}

func main() {
    args := flags.Parse()
    if err := RunApp(args.Path); err != nil {
        log := logger.New()
        log.Error("Приложение завершилось с ошибкой", err)
        os.Exit(1)
    }
}
