package output

import (
	"fmt"
	"reflect"

	"github.com/cedi/kkpctl/pkg/utils"
)

const (
	// Text as output parameter specifies the output format as human readable text
	Text string = "text"

	// JSON as output parameter specifies the output format as JSON
	JSON string = "json"

	// YAML as output parameter specifies the output format as YAML (Experimental)
	YAML string = "yaml"

	// Name as sortBy parameter specified the output (if in a list) should be sorted by name
	Name string = "name"

	// Date as sortBy parameter specified the output (if in a list) should be sorted by Date
	Date string = "date"
)

// make the parser factory a singleton
var parserFactory *ParserFactory

// GetParserFactory returns a singleton instance for the parser factory
func GetParserFactory() *ParserFactory {
	if parserFactory == nil {
		parserFactory = NewParserFactory()
	}

	return parserFactory
}

// ParseOutput takes any KKP Object as an input and then parses it to the appropriate output format
func ParseOutput(object interface{}, output string, sortBy string) (string, error) {

	err := validateOutput(output)
	if err != nil {
		return "", err
	}

	err = validateSorting(sortBy)
	if err != nil {
		return "", err
	}

	parser := GetParserFactory()

	collectionsParser, ok := parser.GetCollectionParser(reflect.TypeOf(object))
	if ok {
		return collectionsParser.ParseCollection(object, output, sortBy)
	}

	objectParser, ok := parser.GetObjectParser(reflect.TypeOf(object))
	if ok {
		return objectParser.ParseObject(object, output)
	}

	return fmt.Sprintf("%v\n", object), fmt.Errorf("unable to determine proper output type")
}

func validateOutput(output string) error {
	if !utils.IsOneOf(output, Text, JSON, YAML) {
		return fmt.Errorf("the output type '%s' is not a valid output", output)
	}
	return nil
}

func validateSorting(sort string) error {
	if !utils.IsOneOf(sort, Name, Date) {
		return fmt.Errorf("the sort parameter '%s' is not a valid sorting criteria", sort)
	}

	return nil
}
