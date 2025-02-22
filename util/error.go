package util

import "github.com/rs/zerolog/log"

func Unwrap(err error, msg string) {
	if err != nil {
		log.Fatal().Err(err).Msg("")
	}
}
