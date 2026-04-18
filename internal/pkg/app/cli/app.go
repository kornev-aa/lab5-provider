package cli

import (
    "fmt"
    "github.com/kornev-aa/lab5-provider/internal/domain/models"
    "github.com/kornev-aa/lab5-provider/pkg/config"
)

type Logger interface {
    Info(msg string)
    Debug(msg string)
    Error(msg string, err error)
}

type WeatherInfo interface {
    GetTemperature(lat, lon float64) (models.TempInfo, error)
}

type cliApp struct {
    l    Logger
    wi   WeatherInfo
    conf config.Config
}

func New(l Logger, wi WeatherInfo, c config.Config) *cliApp {
    return &cliApp{
        l:    l,
        wi:   wi,
        conf: c,
    }
}

func (c *cliApp) Run() error {
    c.l.Info("Запуск приложения")
    c.l.Info(fmt.Sprintf("Координаты из конфига: широта=%.4f, долгота=%.4f", c.conf.L.Lat, c.conf.L.Long))

    tempInfo, err := c.wi.GetTemperature(c.conf.L.Lat, c.conf.L.Long)
    if err != nil {
        c.l.Error("Ошибка получения температуры", err)
        return err
    }

    fmt.Printf("Температура воздуха - %.2f градусов цельсия\n", tempInfo.Temp)
    return nil
}
