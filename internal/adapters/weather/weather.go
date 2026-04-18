package weather

import (
    "encoding/json"
    "errors"
    "fmt"
    "io"
    "net/http"
    "time"
    "github.com/kornev-aa/lab5-provider/internal/domain/models"
    "github.com/kornev-aa/lab5-provider/pkg/cache"
)

const apiURL = "https://api.open-meteo.com/v1/forecast"

type Logger interface {
    Info(msg string)
    Debug(msg string)
    Error(msg string, err error)
}

type WeatherInfo struct {
    c        current
    l        Logger
    cache    cache.Cache
    cacheTTL time.Duration
    isLoaded bool
}

type current struct {
    Temp float64 `json:"temperature_2m"`
}

type response struct {
    Curr current `json:"current"`
}

func New(l Logger, cache cache.Cache, cacheTTL time.Duration) *WeatherInfo {
    return &WeatherInfo{
        l:        l,
        cache:    cache,
        cacheTTL: cacheTTL,
    }
}

func (wi *WeatherInfo) getWeatherInfo(lat, lon float64) error {
    cacheKey := fmt.Sprintf("weather:%.4f:%.4f", lat, lon)

    if cached, found := wi.cache.Get(cacheKey); found {
        var resp response
        if err := json.Unmarshal(cached, &resp); err == nil {
            wi.c = resp.Curr
            wi.isLoaded = true
            wi.l.Debug("Данные получены из кэша")
            return nil
        }
    }

    params := fmt.Sprintf(
        "latitude=%f&longitude=%f&current=temperature_2m",
        lat, lon,
    )
    url := fmt.Sprintf("%s?%s", apiURL, params)

    wi.l.Debug(fmt.Sprintf("Запрос к API: %s", url))

    resp, err := http.Get(url)
    if err != nil {
        wi.l.Error("Ошибка запроса к API", err)
        return errors.Join(errors.New("failed to fetch weather"), err)
    }
    defer resp.Body.Close()

    data, err := io.ReadAll(resp.Body)
    if err != nil {
        wi.l.Error("Ошибка чтения ответа", err)
        return errors.Join(errors.New("failed to read response"), err)
    }

    var apiResp response
    if err := json.Unmarshal(data, &apiResp); err != nil {
        wi.l.Error("Ошибка парсинга JSON", err)
        return errors.Join(errors.New("failed to parse response"), err)
    }

    wi.c = apiResp.Curr
    wi.isLoaded = true
    wi.cache.Set(cacheKey, data, wi.cacheTTL)
    wi.l.Debug("Данные сохранены в кэш")

    return nil
}

func (wi *WeatherInfo) GetTemperature(lat, lon float64) (models.TempInfo, error) {
    err := wi.getWeatherInfo(lat, lon)
    return models.TempInfo{Temp: wi.c.Temp}, err
}
