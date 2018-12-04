package main

import (
	"flag"
	"log"
	"math"
	"os"
	"path"
	"runtime"
	"sync"

	"github.com/joho/godotenv"
)

var (
	flagListenAddr   = flag.String("listen", ":8008", "the http listen address")
	flagWorkersCount = flag.Int("workers", int(math.Pow(float64(runtime.NumCPU()), 4)), "the number of workers")
	flagEnvFile      = flag.String("env", path.Join(path.Dir(os.Args[0]), ".env"), "the plugins directory")
)

var (
	currentVersion = "1.0.0-alpha"
	globalContext  = new(sync.Map)
)

func init() {
	flag.Parse()

	runtime.GOMAXPROCS(*flagWorkersCount)

	if err := godotenv.Overload(*flagEnvFile); err != nil {
		log.Fatal("[.env] ", err.Error())
	}
}
