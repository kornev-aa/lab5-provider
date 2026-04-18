package pogodaby

import (
    "encoding/json"
    "net/http"
    "github.com/kornev-aa/lab5-provider/internal/domain/models"
)

const apiURL = "https://pogoda.by/api/v2/weather-fact?station=26820"

type Logger interface {
    Info(msg string)
    Debug(msg string)
    Error(msg string, err error)
}

type resp struct {
    Temp float32 `json:"t"`
}

type Pogoda struct {
    l Logger
}

func New(l Logger) *Pogoda {
    return &Pogoda{l: l}
}

func (p *Pogoda) GetTemperature(lat, lon float64) (models.TempInfo, error) {
    p.l.Debug("Запрос к pogoda.by API")

    response, err := http.Get(apiURL)
    if err != nil {
        p.l.Error("Ошибка запроса к pogoda.by", err)
        return models.TempInfo{}, err
    }
    defer func() {
        if err := response.Body.Close(); err != nil {
            p.l.Error("Ошибка закрытия body", err)
        }
    }()

    var r resp
    if err := json.NewDecoder(response.Body).Decode(&r); err != nil {
        p.l.Error("Ошибка декодирования JSON", err)
        return models.TempInfo{}, err
    }

    p.l.Debug("Данные получены от pogoda.by")
    return models.TempInfo{Temp: float64(r.Temp)}, nil
}
