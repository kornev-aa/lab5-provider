package gui

import (
    "fmt"
    "strconv"
    "time"
    "fyne.io/fyne/v2"
    "fyne.io/fyne/v2/app"
    "fyne.io/fyne/v2/container"
    "fyne.io/fyne/v2/widget"
    "github.com/kornev-aa/lab5-provider/internal/adapters/weather"
    "github.com/kornev-aa/lab5-provider/pkg/cache"
    "github.com/kornev-aa/lab5-provider/pkg/logger"
    "github.com/kornev-aa/lab5-provider/pkg/storage"
)

type GUIApp struct {
    log     *logger.Logger
    storage storage.LocationStorage
    weather *weather.WeatherInfo
}

func NewGUIApp(log *logger.Logger, storage storage.LocationStorage, cache cache.Cache, cacheTTL time.Duration) *GUIApp {
    return &GUIApp{
        log:     log,
        storage: storage,
        weather: weather.New(log, cache, cacheTTL),
    }
}

func (g *GUIApp) Run() error {
    g.log.Info("Запуск GUI приложения")

    myApp := app.New()
    myWindow := myApp.NewWindow("Погода")
    myWindow.Resize(fyne.NewSize(400, 300))

    lat, _ := g.storage.GetLatitude()
    lon, _ := g.storage.GetLongitude()

    latEntry := widget.NewEntry()
    latEntry.SetText(fmt.Sprintf("%.4f", lat))
    lonEntry := widget.NewEntry()
    lonEntry.SetText(fmt.Sprintf("%.4f", lon))

    resultLabel := widget.NewLabel("Нажмите 'Получить погоду'")
    resultLabel.Wrapping = fyne.TextWrapWord

    getWeatherBtn := widget.NewButton("Получить погоду", func() {
        g.log.Debug("Кнопка 'Получить погоду' нажата")

        latVal, err1 := strconv.ParseFloat(latEntry.Text, 64)
        lonVal, err2 := strconv.ParseFloat(lonEntry.Text, 64)

        if err1 != nil || err2 != nil {
            resultLabel.SetText("Ошибка: введите корректные координаты")
            g.log.Error("Неверные координаты", fmt.Errorf("lat=%s, lon=%s", latEntry.Text, lonEntry.Text))
            return
        }

        g.log.Debug(fmt.Sprintf("Запрос погоды для координат: %.4f, %.4f", latVal, lonVal))

        temp := g.weather.GetTemperature(latVal, lonVal)

        resultLabel.SetText(fmt.Sprintf("Температура: %.2f°C", temp.Temp))
        g.log.Info(fmt.Sprintf("Получена погода: %.2f°C", temp.Temp))
    })

    saveLocationBtn := widget.NewButton("Сохранить координаты", func() {
        g.log.Debug("Кнопка 'Сохранить координаты' нажата")

        latVal, err1 := strconv.ParseFloat(latEntry.Text, 64)
        lonVal, err2 := strconv.ParseFloat(lonEntry.Text, 64)

        if err1 != nil || err2 != nil {
            resultLabel.SetText("Ошибка: введите корректные координаты")
            g.log.Error("Неверные координаты", fmt.Errorf("lat=%s, lon=%s", latEntry.Text, lonEntry.Text))
            return
        }

        if err := g.storage.SaveLocation(latVal, lonVal); err != nil {
            resultLabel.SetText(fmt.Sprintf("Ошибка сохранения: %s", err.Error()))
            g.log.Error("Ошибка сохранения координат", err)
            return
        }

        resultLabel.SetText("Координаты сохранены!")
        g.log.Info(fmt.Sprintf("Координаты сохранены: %.4f, %.4f", latVal, lonVal))
    })

    content := container.NewVBox(
        widget.NewLabel("Широта:"),
        latEntry,
        widget.NewLabel("Долгота:"),
        lonEntry,
        container.NewHBox(getWeatherBtn, saveLocationBtn),
        widget.NewSeparator(),
        resultLabel,
    )

    myWindow.SetContent(content)
    myWindow.ShowAndRun()
    return nil
}
