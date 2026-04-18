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
    // Используем координаты из конфига как источник "хранилища"
    store = storage.NewMemoryStorage(cfg.L.Lat, cfg.L.Long)
    logg.Info("Using memory storage from config")

    // Создаём кэш в памяти
    memCache := cache.NewMemoryCache()
    cacheTTL := 5 * time.Minute

    handlers := httphandlers.NewHandlers(logg, store, memCache, cacheTTL)

    router := chi.NewRouter()
    router.Use(middleware.Logger)
    router.Use(middleware.Recoverer)

    router.Get("/weather", handlers.GetWeather)
    router.Post("/location", handlers.SaveLocation)
    router.Get("/location", handlers.GetLocation)

    logg.Info(fmt.Sprintf("HTTP server starting on :8080"))
    http.ListenAndServe(":8080", router)
}
