package utils

import (
	"fmt"
	"os"
	"time"

	"log/slog"
)

func GetLogger(serviceName string) *slog.Logger {
	// logLevel := slog.LevelInfo // accept [Info, Warn, Error]
	logLevel := slog.LevelDebug // accept all log levels

	// load timezone location
	location, err := time.LoadLocation("America/Argentina/Buenos_Aires")

	if err != nil {
		fmt.Println("[GetLogger] [Error]:", err)
	}

	//get the current date and time in the Argentina timezone
	currentTime := time.Now().In(location)

	opts := &slog.HandlerOptions{
		Level: logLevel,
		ReplaceAttr: func(groups []string, attr slog.Attr) slog.Attr {
			if attr.Key == slog.TimeKey {
				attr.Value = slog.StringValue(currentTime.Format("2006-01-02 15:04:05"))
			}
			return attr
		},
	}

	jsonHandler := slog.NewJSONHandler(os.Stdout, opts).WithAttrs([]slog.Attr{
		slog.String("service", serviceName),
	})
	logger := slog.New(jsonHandler)
	return logger
}
