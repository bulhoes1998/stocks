package main

import "github.com/bulhoes1998/stock/cmd/api/internal/app"

func main() {
	if err := run(); err != nil {
		panic(err)
	}
}

func run() error {
	app.BuildApplication()

	return nil
}
