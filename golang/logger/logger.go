package logger

import (
	"log/slog"
	"os"
)

func SetupLogger(level slog.Level) *slog.Logger {
	handler := slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		Level: level,
		ReplaceAttr: func(groups []string, a slog.Attr) slog.Attr {
			switch a.Key {
			case slog.TimeKey: // Format time as YYYY-MM-DD HH:MM:SS
				return slog.String(a.Key, a.Value.Time().Format("2006-01-02 15:04:05"))
			case slog.LevelKey: // Shorten log level names
				levelMap := map[slog.Level]string{
					slog.LevelDebug: "D",
					slog.LevelInfo:  "I",
					slog.LevelWarn:  "W",
					slog.LevelError: "E",
				}
				return slog.String(a.Key, levelMap[a.Value.Any().(slog.Level)])
			default:
				return a
			}
		},
	})

	return slog.New(handler)
}
