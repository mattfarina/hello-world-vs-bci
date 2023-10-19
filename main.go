package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"

	"log/slog"
)

const out = `Hello World`

var debug = flag.Bool("debug", false, "display debug output")

func handler(w http.ResponseWriter, r *http.Request) {
	slog.Info("Printing message")
	slog.Debug("Request %v", r)
	fmt.Fprint(w, out)
}

func main() {

	flag.Parse()

	var logger *slog.Logger
	if *debug {
		logger = slog.New(slog.NewTextHandler(os.Stderr, &slog.HandlerOptions{
			Level: slog.LevelDebug,
		}))
	} else {
		logger = slog.New(slog.NewTextHandler(os.Stderr, nil))
	}
	slog.SetDefault(logger)

	slog.Info("Starting Hello World")

	http.HandleFunc("/", handler)
	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}
}
