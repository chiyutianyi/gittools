package main

import (
	"fmt"
	"os"

	"github.com/go-git/go-git/v5/plumbing/format/idxfile"
)

func main() {
	args := os.Args
	if len(args) < 2 {
		fmt.Println("Usage: idxfile <idx file>")
		os.Exit(1)
	}

	idx := args[1]

	f, err := os.Open(idx)
	if err != nil {
		fmt.Fprintf(os.Stderr, "can not open %v: %v\n", idx, err)
		return
	}
	defer f.Close()

	i := idxfile.NewMemoryIndex()
	d := idxfile.NewDecoder(f)
	if err = d.Decode(i); err != nil {
		fmt.Fprintf(os.Stderr, "decode err: %v\n", err)
		return
	}

	fmt.Printf("Version: %v\nFanout: %v\nFanoutMapping: %v\nNames: %v\n"+
		"Offset32: %v\nOffset64: %v\nCRC32: %v\nPackfileChecksum: %x\nIdxChecksum: %x\n",
		i.Version,
		len(i.Fanout), len(i.FanoutMapping),
		len(i.Names),
		len(i.Offset32),
		len(i.Offset64),
		len(i.CRC32),
		i.PackfileChecksum,
		i.IdxChecksum)

	for i, d := range i.Fanout {
		fmt.Printf("Fanout[%02x] = %v\n", i, d)
	}

	for i, d := range i.FanoutMapping {
		fmt.Printf("FanoutMapping[%02x] = %v\n", i, d)
	}

	for i, d := range i.Names {
		fmt.Printf("Names[%v] = %x\n", i, d)
	}

	for i, d := range i.Offset32 {
		fmt.Printf("Offset32[%v] = %x\n", i, d)
	}

	for i, d := range i.CRC32 {
		fmt.Printf("CRC32[%v] = %x\n", i, d)
	}
}
