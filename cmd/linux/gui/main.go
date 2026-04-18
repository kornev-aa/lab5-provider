package main

import (
    "fmt"
    "os"
    "time"
    "github.com/kornev-aa/lab5-tests/internal/pkg/gui"
    "github.com/kornev-aa/lab5-tests/pkg/cache"
    "github.com/kornev-aa/lab5-tests/pkg/config"
    "github.com/kornev-aa/lab5-tests/pkg/logger"
    "github.com/kornev-aa/lab5-tests/pkg/storage"
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

    var store storage.LocationStorage
    switch cfg.StorageType {
    case "file":
        store = storage.NewFileStorage(cfg.FilePath)
        log.Info("Используется файловое хранилище")
    default:
        store = storage.NewFileStorage("./location.json")
        log.Info("Используется файловое хранилище по умолчанию")
    }

    // Создаём кэш в памяти
    memCache := cache.NewMemoryCache()
    cacheTTL := 5 * time.Minute

    app := gui.NewGUIApp(log, store, memCache, cacheTTL)
    if err := app.Run(); err != nil {
        log.Error("Ошибка GUI", err)
        os.Exit(1)
    }
}
