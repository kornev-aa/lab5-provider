package main

import (
    "fmt"
    "os"
    "log"
    "net/http"
    "time"
    "github.com/go-chi/chi/v5"
    "github.com/go-chi/chi/v5/middleware"
    httphandlers "github.com/kornev-aa/lab5-provider/internal/pkg/http"
    "github.com/kornev-aa/lab5-provider/pkg/cache"
    "github.com/kornev-aa/lab5-provider/pkg/config"
    "github.com/kornev-aa/lab5-provider/pkg/logger"
    "github.com/kornev-aa/lab5-provider/pkg/storage"
)

func main() {
    // Загружаем конфиг из YAML файла
    r, err := os.Open("./config/config.yaml")
    if err != nil {
        log.Fatal("Failed to open config:", err)
    }
    defer r.Close()

    cfg, err := config.Parse(r)
    if err != nil {
        log.Fatal("Failed to parse config:", err)
    }

    logg := logger.New()

    var store storage.LocationStorage
    switch cfg.StorageType {
    case "file":
        store = storage.NewFileStorage(cfg.FilePath)
        logg.Info("Using file storage")
    default:
        store = storage.NewFileStorage("./location.json")
        logg.Info("Using default file storage")
    }

    // Создаём кэш в памяти
    memCache := cache.NewMemoryCache()
    cacheTTL := 5 * time.Minute

    handlers := httphandlers.NewHandlers(logg, store, memCache, cacheTTL)

    r := chi.NewRouter()
    r.Use(middleware.Logger)
    r.Use(middleware.Recoverer)

    r.Get("/weather", handlers.GetWeather)
    r.Post("/location", handlers.SaveLocation)
    r.Get("/location", handlers.GetLocation)

    logg.Info(fmt.Sprintf("HTTP server starting on :8080"))
    http.ListenAndServe(":8080", r)
}
