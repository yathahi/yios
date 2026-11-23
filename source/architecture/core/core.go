package main

import (
	"log/slog"
	"os"
	"yathahi/source/library"
	packagemanager "yathahi/source/library/package-manager"
)

func main() {
	slog.SetDefault(slog.New(slog.NewTextHandler(os.Stdout, nil)))
	requestsChannel := make(chan library.EntryData)

	go library.UdsServer(library.EntrySocketPath, requestsChannel)

	go Containersetup()

	go func() {

		for {
			receivedData := <-requestsChannel
			packagemanager.PackageManager(receivedData.Service)
		}
	}()

	select {}
}
