package storage

type MemoryStorage struct {
    lat  float64
    lon  float64
}

func NewMemoryStorage(lat, lon float64) *MemoryStorage {
    return &MemoryStorage{
        lat: lat,
        lon: lon,
    }
}

func (m *MemoryStorage) GetLatitude() (float64, error) {
    return m.lat, nil
}

func (m *MemoryStorage) GetLongitude() (float64, error) {
    return m.lon, nil
}

func (m *MemoryStorage) SaveLocation(lat, lon float64) error {
    m.lat = lat
    m.lon = lon
    return nil
}
