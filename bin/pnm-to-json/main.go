package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/ianlopshire/go-fixedwidth"
	"github.com/omniboost/go-sage100c"
)

func main() {
	// stdin is "just" a file so stat it
	stdin := os.Stdin
	fi, err := stdin.Stat()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	// var input io.Reader
	var input []byte
	if fi.Mode()&os.ModeNamedPipe == 0 {
		args := os.Args[1:len(os.Args)]
		// no pipe: check for positional argument
		if len(args) == 0 {
			fmt.Fprintln(os.Stderr, "Provide pnm file via stdin or argument")
			os.Exit(1)
		}
		// input = strings.NewReader(args[0])
		// input = []byte(args[0])
		input, err = os.ReadFile(args[0])
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	} else {
		// input = stdin
		input, err = ioutil.ReadAll(stdin)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	}

	splits := bytes.SplitN(input, []byte("\r\n"), 2)

	// header := sage100c.Header{}
	// r := bytes.NewReader(splits[0])
	// dec := fixedwidth.NewDecoder(r)
	// dec.SetLineTerminator([]byte("\r\n"))
	// dec.SetUseCodepointIndices(true)
	// err = dec.Decode(&header)
	// if err != nil {
	// 	fmt.Fprintln(os.Stderr, err)
	// 	os.Exit(1)
	// }

	// b, err := json.MarshalIndent(header, "", "  ")
	// if err != nil {
	// 	t.Error(err)
	// }

	lines := sage100c.Lines{}
	r := bytes.NewReader(splits[1])
	dec := fixedwidth.NewDecoder(r)
	dec.SetLineTerminator([]byte("\r\n"))
	dec.SetUseCodepointIndices(true)
	err = dec.Decode(&lines)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	b, err := json.MarshalIndent(lines, "", "  ")
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	fmt.Println(string(b))
}
