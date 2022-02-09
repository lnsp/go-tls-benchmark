package main

import (
	"flag"
	"log"
	"net/http"
	"os"
	"os/signal"
	"runtime/pprof"
	"syscall"
)

var (
	tlsKeyFile = flag.String("tls-key", "", "tls private key file")
	tlsCertFile = flag.String("tls-cert", "", "tls certificate file")
	cpuProfile = flag.String("cpu-profile", "", "write cpu profile to `file`")
)

func main() {
	flag.Parse()
	// Setup profiling
	if *cpuProfile != "" {
        f, err := os.Create(*cpuProfile)
        if err != nil {
            log.Fatal("could not create CPU profile: ", err)
        }
        defer f.Close() // error handling omitted for example
        if err := pprof.StartCPUProfile(f); err != nil {
            log.Fatal("could not start CPU profile: ", err)
        }
        defer pprof.StopCPUProfile()
    }
	// Setup HTTP
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello, world\n"))
	})
	server := &http.Server{
		Addr: ":8443",
		Handler: handler,
	}
	// Setup signal handler
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	go func() { <-sigs; server.Close() }()
	// And wait for connections
	if err := server.ListenAndServeTLS(*tlsCertFile, *tlsKeyFile); err != http.ErrServerClosed {
		log.Fatal(err)
	}
}
