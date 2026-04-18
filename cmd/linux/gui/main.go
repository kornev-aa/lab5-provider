package main

import (
    "fmt"
    "os"
    "time"
    "github.com/kornev-aa/lab5-provider/internal/pkg/gui"
    "github.com/kornev-aa/lab5-provider/pkg/cache"
    "github.com/kornev-aa/lab5-provider/pkg/config"
    "github.com/kornev-aa/lab5-provider/pkg/logger"
    "github.com/kornev-aa/lab5-provider/pkg/storage"
)

func main() {
    // Загружаем конфиг из YAML файла
    r, err := os.Open("./config/config.yaml")
    if err != nil {
        fmt.Printf("Ошибка открытия конфига: %s\n", err.Error())
        os.Exit(1)
    }
    defer r.Close()

    cfg, err := config.Parse(r)
    if err != nil {
        fmt.Printf("Ошибка парсинга конфига: %s\n", err.Error())
        os.Exit(1)
    }

    log := logger.New()

    // Используем координаты из конфига
    store := storage.NewMemoryStorage(cfg.L.Lat, cfg.L.Long)
    log.Info("Используется хранилище из конфига")

    // Создаём кэш в памяти
    memCache := cache.NewMemoryCache()
    cacheTTL := 5 * time.Minute

    app := gui.NewGUIApp(log, store, memCache, cacheTTL)
    if err := app.Run(); err != nil {
        log.Error("Ошибка GUI", err)
        os.Exit(1)
    }
}
