package config

import (
	"sync"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

var (
	onceLog sync.Once
	Log     zerolog.Logger
)

func init() {
	onceLog.Do(func() {
		Log = log.With().Caller().Logger()
	})
}
