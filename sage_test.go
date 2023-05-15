package sage100c_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"testing"

	"github.com/ianlopshire/go-fixedwidth"
	"github.com/omniboost/go-sage100c"
)

func TestTypeC(t *testing.T) {
	// data, _ := ioutil.ReadFile("/home/leon/Desktop/Exemple journal 01-2019 ASCII")
	// pieces := strings.Split(string(data), "\r\n")
	// for k, v := range pieces {
	// 	fmt.Printf("%d: %s\n", k, v)
	// }
	// os.Exit(4)

	data, err := ioutil.ReadFile("H6342_OD_20221026.pnm")
	if err != nil {
		t.Error(err)
	}

	// r, err := os.Open("H6342_OD_20221026.pnm")
	// if err != nil {
	// 	t.Error(err)
	// }

	splits := bytes.SplitN(data, []byte("\r\n"), 2)

	header := sage100c.Header{}
	r := bytes.NewReader(splits[0])
	dec := fixedwidth.NewDecoder(r)
	dec.SetLineTerminator([]byte("\r\n"))
	dec.SetUseCodepointIndices(true)
	err = dec.Decode(&header)
	if err != nil {
		t.Error(err)
	}

	b, err := json.MarshalIndent(header, "", "  ")
	if err != nil {
		t.Error(err)
	}

	log.Println(string(b))

	lines := sage100c.Lines{}
	r = bytes.NewReader(splits[1])
	dec = fixedwidth.NewDecoder(r)
	dec.SetLineTerminator([]byte("\r\n"))
	dec.SetUseCodepointIndices(true)
	err = dec.Decode(&lines)
	if err != nil {
		t.Error(err)
	}

	b, err = json.MarshalIndent(lines, "", "  ")
	if err != nil {
		t.Error(err)
	}

	log.Println(string(b))

	buf := bytes.NewBuffer(nil)
	enc := fixedwidth.NewEncoder(buf)
	enc.SetLineTerminator([]byte("\r\n"))
	enc.SetUseCodepointIndices(true)
	err = enc.Encode(lines)
	if err != nil {
		t.Error(err)
	}
	buf.WriteString("\r\n")

	if !bytes.Equal(buf.Bytes(), splits[1]) {
		fmt.Println(string(splits[1]))
		fmt.Println(buf.String())
		t.Error("Input & output not equal")
	}
}
