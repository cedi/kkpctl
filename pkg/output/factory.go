package output

import "reflect"

// CollectionsParser is the interface used to parse a collection of items
type CollectionsParser interface {
	ParseCollection(any, string, string) (string, error)
}

// ObjectParser is the interface used to parse a single item
type ObjectParser interface {
	ParseObject(any, string) (string, error)
}

// ParserFactory is the factory to keep track of all parsers
type ParserFactory struct {
	collectionParsers map[reflect.Type]CollectionsParser
	objectParsers     map[reflect.Type]ObjectParser
}

// NewParserFactory returns a new ParserFactory
func NewParserFactory() *ParserFactory {
	return &ParserFactory{
		collectionParsers: make(map[reflect.Type]CollectionsParser),
		objectParsers:     make(map[reflect.Type]ObjectParser),
	}
}

// AddCollectionParser adds a parser for a collection
func (o *ParserFactory) AddCollectionParser(objType reflect.Type, outputParser CollectionsParser) {
	o.collectionParsers[objType] = outputParser
}

// AddObjectParser adds a parser for a object
func (o *ParserFactory) AddObjectParser(objType reflect.Type, outputParser ObjectParser) {
	o.objectParsers[objType] = outputParser
}

// GetCollectionParser gets a parser for a collection
func (o *ParserFactory) GetCollectionParser(objType reflect.Type) (CollectionsParser, bool) {
	parser, ok := o.collectionParsers[objType]
	return parser, ok
}

// GetObjectParser gets a parser for a object
func (o *ParserFactory) GetObjectParser(objType reflect.Type) (ObjectParser, bool) {
	parser, ok := o.objectParsers[objType]
	return parser, ok
}
