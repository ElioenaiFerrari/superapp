package logger

import "github.com/rs/zerolog"

var L = zerolog.New(zerolog.NewConsoleWriter()).With().Timestamp().Logger()
