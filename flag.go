package util

import (
	"flag"
	"os"
)

func Usage() {
	flag.Usage()
	os.Exit(1)
}

func UsageF(flag *flag.FlagSet) {
	flag.Usage()
	os.Exit(1)
}
