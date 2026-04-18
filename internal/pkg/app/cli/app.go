package cli

import (
    "fmt"
    "github.com/kornev-aa/lab5-tests/internal/domain/models"
    "github.com/kornev-aa/lab5-tests/pkg/config"
)

type Logger interface {
    Info(msg string)
    Debug(msg string)
    Error(msg string, err error)
}

type WeatherInfo interface {
    GetTemperature(lat, lon float64) models.TempInfo
}

type cliApp struct {
    l   Logger
    wi  WeatherInfo
    cfg config.Config
}

func New(l Logger, wi WeatherInfo, cfg config.Config) *cliApp {
    return &cliApp{
        l:   l,
        wi:  wi,
        cfg: cfg,
    }
}

func (c *cliApp) Run() error {
    c.l.Info("Запуск приложения")

    lat := c.cfg.L.Lat
    lon := c.cfg.L.Long

    c.l.Info(fmt.Sprintf("Координаты из конфига: широта=%.4f, долгота=%.4f", lat, lon))

    temp := c.wi.GetTemperature(lat, lon)

    fmt.Printf("Температура воздуха - %.2f градусов цельсия\n", temp.Temp)
    return nil
}
