package main

//import protobuf "github.com/cuelang.org/protobuf/proto"
import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"

	"cuelang.org/go/cue/format"
	"cuelang.org/go/encoding/protobuf"
)

func main() {

	args := os.Args[1:]
	d, f := filepath.Split(args[0])
	cwd, _ := os.Getwd()
	var paths = []string{}
	// currenet dir
	paths = append(paths, cwd)
	// referernced location
	paths = append(paths, filepath.Join(d, "../"))
	paths = append(paths, filepath.Join(d, "../../"))
	paths = append(paths, filepath.Join(d, "../../../"))
	paths = append(paths, filepath.Join(d, "../../../../"))
	paths = append(paths, filepath.Join(d, "../../../_extschema/schema/proto/"))    // mauice
	paths = append(paths, filepath.Join(d, "../../../../_extschema/schema/proto/")) // pikachu
	// on the golang root
	paths = append(paths, filepath.Join(os.Getenv("GOPATH"), "/"))
	paths = append(paths, filepath.Join(os.Getenv("GOPATH"), "src/"))
	paths = append(paths, filepath.Dir("/usr/local/Cellar/protobuf/3.10.1/include/"))

	fmt.Printf(strings.Join(paths, "\n"))
	fmt.Printf("\n\n")

	file, err := protobuf.Extract(args[0], nil, &protobuf.Config{
		Paths: paths,
	})

	if err != nil {
		log.Fatal(err, "")
	}

	b, _ := format.Node(file)
	ioutil.WriteFile(filepath.Join("out", strings.Join([]string{f, "cue"}, ".")), b, 0644)
}
