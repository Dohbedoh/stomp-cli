package main

import (
	"github.com/dohbedoh/stomp-cli/cmd"
	"github.com/rs/zerolog/pkgerrors"

	"github.com/rs/zerolog"
)

func main() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	zerolog.ErrorStackMarshaler = pkgerrors.MarshalStack
	cmd.Execute()
}
