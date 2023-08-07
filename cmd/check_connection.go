package cmd

import (
	"context"
	"crypto/tls"
	"net"
	"strings"
	"time"

	"github.com/go-stomp/stomp"
	zlog "github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

var transport string
var skipTls bool
var username string
var password string
var timeout int64

func NewCheckConnectionCommand() *cobra.Command {
	command := &cobra.Command{
		Use:     "connect",
		Short:   "Check connection to a STOMP server and returns the version",
		Aliases: []string{"conn"},
		Args:    cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			version, err := TestConnection(args[0], &ConnectOptions{
				Transport: transport,
				SkipTls:   skipTls,
				Username:  username,
				Password:  password,
				Timeout:   timeout,
			})
			if err != nil {
				zlog.Error().Err(err).Msg("")
			} else {
				zlog.Info().Msg(version)
			}
		},
	}
	command.Flags().StringVarP(&transport, "transport", "t", "tls", "Transport type (default|tls). Default to 'tls'")
	command.Flags().BoolVarP(&skipTls, "skipTls", "k", false, "Skip TLS (optional)")
	command.Flags().StringVarP(&username, "username", "u", "", "Username (optional)")
	command.Flags().StringVarP(&password, "password", "p", "", "Password (optional)")
	command.Flags().Int64VarP(&timeout, "connectTimeout", "c", 10, "Connection timeout in seconds (optional)")
	return command
}

type ConnectOptions struct {
	Transport          string
	SkipTls            bool
	Username, Password string
	Timeout            int64
}

// Example of connecting to a STOMP server using an existing network connection.
func TestConnection(address string, options *ConnectOptions) (string, error) {
	if strings.EqualFold(options.Transport, "tls") {
		return TestTLSConnection(address, options)
	} else {
		return TestDefaultConnection(address, options)
	}
}

// Example of connecting to a STOMP server using an existing network connection.
func TestDefaultConnection(address string, options *ConnectOptions) (string, error) {

	// Create the TLS Connection
	netConn, err := net.DialTimeout("tcp", address, time.Duration(options.Timeout)*time.Second)
	if err != nil {
		return "", err
	}
	defer netConn.Close()

	// Create the Stomp Connection
	var stompConn *stomp.Conn
	if options.Username != "" {
		stompConn, err = stomp.Connect(netConn,
			stomp.ConnOpt.Login(options.Username, options.Password))
	} else {
		stompConn, err = stomp.Connect(netConn)
	}
	if err != nil {
		return "", err
	}
	defer stompConn.Disconnect()
	version := stompConn.Version().String()

	defer stompConn.Disconnect()

	return version, nil
}

// Example of connecting to a STOMP server using an existing network connection.
func TestTLSConnection(address string, options *ConnectOptions) (string, error) {

	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(options.Timeout)*time.Second)
	d := tls.Dialer{
		Config: &tls.Config{
			InsecureSkipVerify: options.SkipTls,
		},
	}
	netConn, err := d.DialContext(ctx, "tcp", address)
	cancel()
	if err != nil {
		return "", err
	}
	defer netConn.Close()

	// Create the Stomp Connection
	var stompConn *stomp.Conn
	if options.Username != "" {
		stompConn, err = stomp.Connect(netConn,
			stomp.ConnOpt.Login(options.Username, options.Password))
	} else {
		stompConn, err = stomp.Connect(netConn)
	}
	if err != nil {
		return "", err
	}
	defer stompConn.Disconnect()
	version := stompConn.Version().String()

	defer stompConn.Disconnect()

	return version, nil
}
