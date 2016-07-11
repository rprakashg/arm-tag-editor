package main

import (
	"github.com/rprakashg/tageditor/server"
)

// initialize
func init() {
	log.SetFlags(log.LstdFlags | log.Lmicroseconds | log.Lshortfile)
}

// main entry point.
func main() {
	server.Start()
}
