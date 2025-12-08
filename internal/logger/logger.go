package logger

import (
	"errors"
	"log/slog"
	"os"
	"project/internal/application/enums"
)

func Init(mode string) (*slog.Logger, error) {
	var handler slog.Handler
	var level slog.Level

	switch mode {
	case enums.LocalApplicationMode.Value():
		level = slog.LevelDebug
		handler = slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
			Level: level,
		})
	case enums.DevApplicationMode.Value():
		level = slog.LevelDebug
		handler = slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
			Level: level,
		})
	case enums.StagingApplicationMode.Value():
		level = slog.LevelDebug
		handler = slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
			Level: level,
		})
	case enums.ProdApplicationMode.Value():
		level = slog.LevelInfo
		handler = slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
			Level: level,
		})
	default:
		return nil, errors.New("unknown logger mode: " + mode)
	}

	return slog.New(handler), nil
}
