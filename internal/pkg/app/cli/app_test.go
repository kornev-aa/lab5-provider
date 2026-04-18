package cli

import (
    "testing"
    "github.com/stretchr/testify/assert"
    "go.uber.org/mock/gomock"
    "github.com/kornev-aa/lab5-tests/internal/domain/models"
    "github.com/kornev-aa/lab5-tests/internal/pkg/app/cli/mocks"
    "github.com/kornev-aa/lab5-tests/pkg/config"
)

type testLogger struct{}

func (l *testLogger) Info(msg string)  {}
func (l *testLogger) Debug(msg string) {}
func (l *testLogger) Error(msg string, err error) {}

func TestCliApp_Run(t *testing.T) {
    ctrl := gomock.NewController(t)
    defer ctrl.Finish()

    log := &testLogger{}
    
    mockWI := mocks.NewMockWeatherInfo(ctrl)
    mockWI.EXPECT().GetTemperature(53.6688, 23.8223).Return(models.TempInfo{Temp: 18.5})

    cfg := config.Config{
        P: config.Provider{Type: "test"},
        L: config.Location{Lat: 53.6688, Long: 23.8223},
    }

    app := New(log, mockWI, cfg)
    err := app.Run()
    
    assert.NoError(t, err)
}
