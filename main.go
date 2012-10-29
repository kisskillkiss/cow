package main

import (
	"flag"
	"log"
	"os"
	"os/signal"
	"runtime"
	"runtime/pprof"
)

var c = make(chan os.Signal, 1)

var cpuprofile = flag.String("cpuprofile", "", "write cpu profile to file")

func main() {
	flag.Parse()
	if *cpuprofile != "" {
		f, err := os.Create(*cpuprofile)
		if err != nil {
			log.Fatal(err)
		}
		pprof.StartCPUProfile(f)
		signal.Notify(c, os.Interrupt)
		go func() {
			for sig := range c {
				info.Printf("captured %v, stopping profiler and exiting..", sig)
				pprof.StopCPUProfile()
				os.Exit(0)
			}
		}()
	}

	runtime.GOMAXPROCS(2)
	loadConfig()

	py := NewProxy("localhost:9000")
	py.Serve()
}
