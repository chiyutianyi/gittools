package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/go-git/go-git/v5/plumbing/format/idxfile"
	"github.com/spf13/cobra"
)

type formatCmd struct {
}

func (cmd *formatCmd) Run(_ *cobra.Command, args []string) {
	if len(args) < 1 {
		fmt.Println("Usage: gittools idx-format <idx file>")
		os.Exit(1)
	}

	idx := args[0]

	if !strings.HasSuffix(idx, ".idx") {
		fmt.Fprintf(os.Stderr, "packfile index %v not end with .idx\n", idx)
		return
	}

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
		if d < 0 {
			continue
		}
		fmt.Printf("FanoutMapping[%02x] = %v\n", i, d)
	}

	for i, d := range i.Names {
		for j := 0; j < len(d)-19; j += 20 {
			fmt.Printf("Names[%v,%d] = %x\n", i, j/20, d[j:j+20])
		}
	}

	for i, d := range i.Offset32 {
		for j := 0; j < len(d)-3; j += 4 {
			offset := fmt.Sprintf("%x", d[j:j+4])
			n, err := strconv.ParseUint(offset, 16, 32)
			if err != nil {
				fmt.Fprintf(os.Stderr, "parse %x: %v\n", offset, err)
				return
			}
			fmt.Printf("Offset[%v,%d] = %d\n", i, j>>2, n)
		}
	}

	for i, d := range i.CRC32 {
		fmt.Printf("CRC32[%v] = %x\n", i, d)
	}
}

func init() {
	format := &formatCmd{}

	cmd := &cobra.Command{
		Use:   "idx-format",
		Short: "format index file",
		Run:   format.Run,
	}

	Cmd.AddCommand(cmd)
}
