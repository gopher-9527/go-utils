package main

import (
	"fmt"
	"os"

	switcher "github.com/hix-k8s-seed/aws-account-switcher"
)

func main() {
	if err := switcher.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
