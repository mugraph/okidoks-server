// logger/logger.go
package logger

import (
	"io"
	"log/slog"
	"os"
	"path/filepath"
)

var Log *slog.Logger

func init() {
	Log = logger(os.Stdout, slog.LevelInfo)
}

func logger(w io.Writer, level slog.Level) *slog.Logger {
	return slog.New(slog.NewTextHandler(w, &slog.HandlerOptions{
		AddSource:   true,
		Level:       &level,
		ReplaceAttr: ReplaceSourceAttr,
	}))
}

// Provide shortfile in logger
func ReplaceSourceAttr(groups []string, a slog.Attr) slog.Attr {
	if a.Key == slog.SourceKey {
		source, _ := a.Value.Any().(*slog.Source)
		if source != nil {
			source.File = filepath.Base(source.File)
		}
	}
	return a
}
