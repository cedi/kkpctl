package output

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"reflect"
	"sort"

	"github.com/lensesio/tableprinter"
	"gopkg.in/yaml.v3"
)

// VersionRender is a intermediate struct to make use of lensesio/tableprinter, which relies on the header anotation
type VersionRender struct {
	Component string `header:"Component"`
	Version   string `header:"Version"`
	Date      string `header:"Date"`
	Commit    string `header:"Commit"`
	BuiltBy   string `header:"BuiltBy"`
}

// ParseCollection parses the version of KKPCTL
func (r VersionRender) ParseCollection(inputObj interface{}, output string, sortBy string) (string, error) {
	var err error
	var parsedOutput []byte

	objects, ok := inputObj.([]VersionRender)
	if !ok {
		return "", fmt.Errorf("inputObj is not a []VersionRender")
	}

	switch output {
	case JSON:
		parsedOutput, err = json.MarshalIndent(objects, "", "  ")

	case YAML:
		parsedOutput, err = yaml.Marshal(objects)

	case Text:
		sort.Slice(objects, func(i, j int) bool {
			return objects[j].Component < objects[i].Component
		})

		var bodyBuf io.ReadWriter
		bodyBuf = new(bytes.Buffer)

		tableprinter.Print(bodyBuf, objects)
		parsedOutput, err = ioutil.ReadAll(bodyBuf)

	default:
		return "", fmt.Errorf("unable to parse objects")
	}

	return string(parsedOutput), err
}

func init() {
	parser := GetParserFactory()
	parser.AddCollectionParser(reflect.TypeOf([]VersionRender{}), VersionRender{})
}
