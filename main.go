package main

import (
	"fmt"
	"os"
)

func main() {
	cli()
	if *setup {
		if err := setupUserProfile(); err != nil {
			panic(err)
		}

		fmt.Printf("CBlog configure init complete. File path: %s \n", globalConfigPath)
		os.Exit(0)
	}

	if err := run(); err != nil {
		panic(err)
	}

	if *preview {
		startWEB(*port)
	}

	if err := upload(); err != nil {
		panic(err)
	}
}
