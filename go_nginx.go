package main

import (
	"github.com/ravitezu/nginxplay"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		os.Args = append(os.Args, "-h")
		nginxplay.CheckArgs(os.Args)
	} else {
		nginxplay.CheckArgs(os.Args)
	}
}
