package out

import (
	"io"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/journald"
)

type logWriter struct {
	logger zerolog.Logger
}

// NewLogWriter creates a new io.Writer that writes to the out log
func NewLogWriter() io.Writer {
	return &logWriter{
		logger: zerolog.New(journald.NewJournalDWriter()).With().Timestamp().Logger(),
	}
}

func (lw logWriter) Write(p []byte) (n int, err error) {
	lw.logger.Print(string(p))
	return len(p), nil
}
