package weather

import (
    "testing"
    "time"
    "github.com/stretchr/testify/assert"
    "go.uber.org/mock/gomock"
    "github.com/kornev-aa/lab5-provider/pkg/cache/mocks"
)

type testLogger struct{}

func (l *testLogger) Info(msg string)  {}
func (l *testLogger) Debug(msg string) {}
func (l *testLogger) Error(msg string, err error) {}

func TestWeatherInfo_GetTemperature_FromCache(t *testing.T) {
    ctrl := gomock.NewController(t)
    defer ctrl.Finish()

    log := &testLogger{}
    mockCache := mocks.NewMockCache(ctrl)
    
    cachedData := []byte(`{"current":{"temperature_2m":25.5}}`)
    mockCache.EXPECT().Get("weather:53.6688:23.8223").Return(cachedData, true)

    wi := New(log, mockCache, 5*time.Minute)
    temp := wi.GetTemperature(53.6688, 23.8223)

    assert.Equal(t, 25.5, temp.Temp)
}

func TestWeatherInfo_GetTemperature_CacheMiss_Success(t *testing.T) {
    ctrl := gomock.NewController(t)
    defer ctrl.Finish()

    log := &testLogger{}
    mockCache := mocks.NewMockCache(ctrl)
    
    mockCache.EXPECT().Get("weather:53.6688:23.8223").Return(nil, false)
    mockCache.EXPECT().Set(gomock.Any(), gomock.Any(), 5*time.Minute).Return()

    wi := New(log, mockCache, 5*time.Minute)
    temp := wi.GetTemperature(53.6688, 23.8223)

    assert.NotEqual(t, 0.0, temp.Temp)
}
