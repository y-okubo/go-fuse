// Mounts MemNodeFs for testing purposes.

package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/y-okubo/go-fuse/fuse"
	"github.com/y-okubo/go-fuse/fuse/nodefs"
)

func main() {
	// Scans the arg list and sets up flags
	debug := flag.Bool("debug", false, "print debugging messages.")
	flag.Parse()
	if flag.NArg() < 2 {
		// TODO - where to get program name?
		fmt.Println("usage: main MOUNTPOINT BACKING-PREFIX")
		os.Exit(2)
	}

	mountPoint := flag.Arg(0)
	prefix := flag.Arg(1)
	root := nodefs.NewMemNodeFSRoot(prefix)
	conn := nodefs.NewFileSystemConnector(root, nil)
	server, err := fuse.NewServer(conn.RawFS(), mountPoint, nil)
	if err != nil {
		fmt.Printf("Mount fail: %v\n", err)
		os.Exit(1)
	}
	server.SetDebug(*debug)
	fmt.Println("Mounted!")
	server.Serve()
}
