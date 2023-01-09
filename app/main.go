package main

import (
	"context"
	"crypto/rand"
	"flag"
	"fmt"
	"io"
	"os"
	"time"

	//"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func genOut(len, count int) {
	var err error

	for i := 0; i < count; i++ {
		b := make([]byte, len)
		_, err = rand.Read(b)
		if err != nil {
			break
		}

		// Convert the slice of bytes to a string.
		m := fmt.Sprintf("%x", b)

		log.Info().Str("name", "fbit.logger").Msgf("%s", m)
	}

	if err != nil {
		log.Error().Str("name", "fbit.logger").Err(err).Msg("failed to logging")
	}
}

func runApp(ctx context.Context) {
	l := ctx.Value("len")
	c := ctx.Value("count")

	genOut(l.(int), c.(int))
}

type Log struct {
	Name     string
	Age      int
	Contents string
}

func main() {
	out := flag.String("o", "", "log out file")
	len := flag.Int("l", 1024*64, "log length")
	count := flag.Int("c", 4, "log count")

	flag.Parse()

	if *out != "" {
		f, err := os.Create(*out)
		if err != nil {
			log.Fatal().Err(err).Msgf("failed to create log file %v", *out)
		}
		defer f.Close()

		zerolog.TimestampFieldName = "logtime"
		zerolog.TimeFieldFormat = time.RFC3339Nano

		var w []io.Writer
		w = append(w, os.Stdout)
		w = append(w, f)

		// Set the log output to the log writer.
		log.Logger = log.Output(io.MultiWriter(w...))
	}

	// docker run -e FOO=bar my_image
	//foo := os.Getenv("FOO")

	ctx := context.Background()
	ctx = context.WithValue(ctx, "len", (int)((*len)/2))
	ctx = context.WithValue(ctx, "count", (int)(*count))

	runApp(ctx)
}
