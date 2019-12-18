package quadracompta_test

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"

	"github.com/ianlopshire/go-fixedwidth"
	"github.com/omniboost/go-quadracompta"
)

func TestTypeC(t *testing.T) {
	// data, _ := ioutil.ReadFile("/home/leon/Desktop/Exemple journal 01-2019 ASCII")
	// pieces := strings.Split(string(data), "\r\n")
	// for k, v := range pieces {
	// 	fmt.Printf("%d: %s\n", k, v)
	// }
	// os.Exit(4)

	r, err := os.Open("/home/leon/Desktop/Exemple journal 01-2019 ASCII")
	if err != nil {
		t.Error(err)
	}

	lines := quadracompta.Lines{}
	dec := fixedwidth.NewDecoder(r)
	dec.SetLineTerminator([]byte("\r\n"))
	dec.SetUseCodepointIndices(true)
	err = dec.Decode(&lines)
	if err != nil {
		t.Error(err)
	}

	// for _, line := range lines {
	// 	log.Println(line.Type())
	// }

	b, err := json.MarshalIndent(lines, "", "  ")
	if err != nil {
		t.Error(err)
	}

	fmt.Println(string(b))
}
