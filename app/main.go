package main

import (
	"context"
	"crypto/rand"
	"fmt"
	"io"
	"os"
	"time"

	"github.com/caarlos0/env/v6"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func genLog(len, count int) {
	var err error

	for i := 0; i < count; i++ {
		b := make([]byte, len)
		_, err = rand.Read(b)
		if err != nil {
			break
		}

		// Convert the slice of bytes to a string.
		m := fmt.Sprintf("%x", b)

		log.Info().Str("name", "fluentlogger").Msgf("%s", m)
	}

	if err != nil {
		log.Error().Str("name", "fluentlogger").Err(err).Msg("failed to logging")
	}
}

func runApp(ctx context.Context) {
	l := ctx.Value("len")
	c := ctx.Value("count")

	genLog(l.(int), c.(int))
}

type LogConfig struct {
	Out   string `env:"LOG_OUT" envDefault:""`
	Len   int    `env:"LOG_LEN" envDefault:"64"`
	Count int    `env:"LOG_COUNT" envDefault:"10"`
}

func main() {
	var w []io.Writer
	w = append(w, os.Stdout)
	zerolog.TimestampFieldName = "logtime"
	zerolog.TimeFieldFormat = time.RFC3339Nano

	cfg := LogConfig{}
	if err := env.Parse(&cfg); err != nil {
		fmt.Printf("failed to parse env with error %+v", err)
		log.Fatal().Err(err).Msgf("failed to parse env")
	}

	// init log with config
	if cfg.Out != "" {
		f, err := os.Create(cfg.Out)
		if err != nil {
			fmt.Printf("failed to create log out %v", cfg.Out)
			log.Fatal().Err(err).Msgf("failed to create log out %v", cfg.Out)
		}
		defer f.Close()

		w = append(w, f)
	}

	// Set the log output to the log writer.
	log.Logger = log.Output(io.MultiWriter(w...))

	ctx := context.Background()
	ctx = context.WithValue(ctx, "len", (int)((cfg.Len)/2))
	ctx = context.WithValue(ctx, "count", (int)(cfg.Count))

	runApp(ctx)
}
