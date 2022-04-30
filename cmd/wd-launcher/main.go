package main

import (
	"bufio"
	"flag"
	"io"
	"io/ioutil"
	"log"
	"os"

	"github.com/erauer/wd-launcher/internal/api"
	"github.com/erauer/wd-launcher/internal/ipc"
	"github.com/erauer/wd-launcher/internal/warp"
)

func main() {
	var logFileName string

	flag.StringVar(&logFileName, "l", "", "Log file name")
	flag.Parse()

	logWriter := ioutil.Discard
	if len(logFileName) > 0 {
		logFile, err := os.OpenFile("/tmp/debug.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
		if nil != err {
			panic(err)
		}
		defer logFile.Close()

		logWriter = logFile
	}

	logger := log.New(logWriter, "wd ", log.LstdFlags)

	projects, err := warp.Load()
	if err != nil {
		logger.Printf("Error: %+v\n", err)
		panic(err)
	}

	multi := io.MultiWriter(os.Stdout, logger.Writer())
	responder := ipc.NewResponder(multi)

	handler := api.NewHandler(*ipc.NewParser(), responder, projects, logger)

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		message := scanner.Text()
		if err := handler.Process(message); err != nil {
			logger.Printf("Error: %+v\n", err)
		}
	}
}
