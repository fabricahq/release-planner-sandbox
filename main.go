// Command greet prints a greeting. It exists to demonstrate Release Planner.
package main

import (
	"flag"
	"fmt"
	"os"
)

// version is set at build time with -ldflags "-X main.version=<version>".
var version = "dev"

func main() {
	showVersion := flag.Bool("version", false, "print the version and exit")
	flag.Parse()
	if *showVersion {
		fmt.Println(version)
		return
	}
	name := "world"
	if flag.NArg() > 0 {
		name = flag.Arg(0)
	}
	fmt.Fprintln(os.Stdout, greeting(name))
}

// greeting returns the message greet prints for name.
func greeting(name string) string {
	return "Hello, " + name + "!"
}
