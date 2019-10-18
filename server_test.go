package indiserver_test

import (
	"os"
	"os/signal"
	"time"

	"github.com/rickbassham/goexec"
	"github.com/rickbassham/logging"
	"github.com/spf13/afero"
	"github.com/goastro/indiserver"
)

func Example() {
	logger := logging.NewLogger(os.Stdout, logging.JSONFormatter{}, logging.LogLevelInfo)
	fs := afero.NewOsFs()
	port := ""
	s := indiserver.NewINDIServer(logger, fs, port, goexec.ExecCommand{})

	logger.WithField("drivers", s.Drivers()).Info("drivers")

	s.StartServer()

	time.Sleep(1 * time.Second)

	s.StartDriver("indi_asi_ccd", "CCD 1")

	println("Server Running. Press CTRL-C to stop.")

	// Wait for a CTRL-C to stop the server.
	sigchan := make(chan os.Signal, 1)
	signal.Notify(sigchan, os.Interrupt)
	<-sigchan

	s.StopDriver("indi_asi_ccd", "CCD 1")
	time.Sleep(1 * time.Second)
	s.StopServer()
}
