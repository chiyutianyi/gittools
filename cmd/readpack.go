package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/spf13/cobra"

	"github.com/chiyutianyi/gittools/pkg/git"
	"github.com/chiyutianyi/gittools/pkg/utils"
)

type readPackCmd struct {
}

func (cmd *readPackCmd) Run(_ *cobra.Command, args []string) {
	if len(args) < 3 {
		fmt.Println("Usage: gittools readpack <packfile> <offset> <length>")
		os.Exit(1)
	}

	packfile := args[0]
	if !strings.HasSuffix(packfile, ".pack") {
		fmt.Fprintf(os.Stderr, "packfile %v not end with .pack\n", packfile)
		return
	}
	offset, err := strconv.ParseInt(args[1], 10, 64)
	if err != nil || offset < 0 {
		fmt.Fprintf(os.Stderr, "error offset %v\n", args[2])
		return
	}
	length, err := strconv.ParseInt(args[2], 10, 64)
	if err != nil || length < 0 {
		fmt.Fprintf(os.Stderr, "error length %v\n", args[3])
		return
	}

	pack, err := os.Open(packfile)
	if err != nil {
		fmt.Fprintf(os.Stderr, "can not open %v: %v\n", packfile, err)
		return
	}
	defer pack.Close()

	content := make([]byte, length)
	pack.ReadAt(content, offset)

	t := git.ParseType(content[0])
	if t == 0 || t > 7 {
		fmt.Fprintf(os.Stderr, "error type %s\n", t)
	} else {
		fmt.Printf("Type = %s\n", t)
	}

	size := int64(content[0] & git.MaskFirstLength)

	i := 0
	c := content[0]
	shift := git.FirstLengthBits
	for c&git.MaskContinue > 0 && i < len(content) {
		i++
		c = content[i]

		size += int64(c&git.MaskLength) << shift
		shift += git.LengthBits
	}

	fmt.Printf("Size = %v\n", size)

	if size == 0 {
		return
	}

	if t < 5 {
		content, err = utils.ZlibUncompress(content[i+1:])
		if err != nil {
			fmt.Fprintf(os.Stderr, "uncompress %v err: %v\n", packfile, err)
		}
		fmt.Printf("Content bytes length = %v\n", len(content))
		fmt.Printf("Content bytes= \n%v\n", content)
		fmt.Printf("Content = \n%v\n", string(content))
		return
	}

	if i+int(size) > len(content) {
		fmt.Printf("Content bytes length = %v\n", len(content))
		fmt.Printf("Content bytes= \n%v\n", content)
	} else {
		fmt.Printf("Content bytes length = %v\n", size)
		fmt.Printf("Content bytes= \n%v\n", content[i:i+int(size)])
	}
}

func init() {
	readPack := &readPackCmd{}

	cmd := &cobra.Command{
		Use:   "read-pack",
		Short: "read a packfile at given position",
		Run:   readPack.Run,
	}

	Cmd.AddCommand(cmd)
}
