package main

import (
	"fmt"
	"time"

	"github.com/alecthomas/kong"
	"github.com/xvello/letsblockit/src/server"
)

func main() {
	start := time.Now()
	options := &server.Options{}
	k := kong.Parse(options)
	err := server.NewServer(options).Start()

	if err == server.ErrDryRunFinished {
		fmt.Printf("Dry-run checks finished in %s\n", time.Since(start))
	} else {
		k.FatalIfErrorf(err)
	}
}
