package blanknodeid

import (
	"encoding"
	_ "fmt"
	"unicode/utf8"
	"unsafe"

	"codeberg.org/reiver/go-erorr"
	"github.com/reiver/go-opt"
	"github.com/reiver/go-ord/en"
)

// Label represents a blank-node-label from RDF (resource description framework) technologies, such as:
// JSON-LD,
// N-Quads,
// N-Triples,
// RDF/XML,
// RDFa,
// TriG,
// Turtle,
// etc.
//
// For each of the Blank Node Identifiers:
//
//	_:b0
//	_:address84
//	_:n1
//	_:ed7ba470-8e54-465e-825c-99712043e01c
//	_:label123
//
// Their Blank Node Labels are as follows:
//
//	b0
//	address84
//	n1
//	ed7ba470-8e54-465e-825c-99712043e01c
//	label123
//
// I.e., a Blank Node Label is the part of the Blank Node Identifier after the "_:".
type Label struct {

	optional opt.Optional[string]
}

var (
	_ encoding.TextMarshaler   = Label{}
	_ encoding.TextUnmarshaler = &Label{}
)

// NoLabel returns an empty [Label].
//
// [Label] is an optional-type
// —
// also known as a option-type or maybe-type elsewhere.
func NoLabel() Label {
	return Label{}
}

func someLabel(value string) Label {
	return Label{optional:opt.Something(value)}
}

// ParseLabelBytes parses a []byte for a blank-node-label from RDF (resource description framework) technologies, such as:
// JSON-LD,
// N-Quads,
// N-Triples,
// RDF/XML,
// RDFa,
// TriG,
// Turtle,
// etc.
//
// A blank-node-label is what comes after the "_:" is a blank-node-identifier.
//
// For example, if the blank-node-identifier is:
//
//	_:abcdef12345
//
// Then its blank-node-label is:
//
//	abcdef12345
//
//
// This uses the definition found in RDF 1.1 Turtle (https://www.w3.org/TR/turtle/):
//
//	RDF blank nodes in Turtle are expressed as _: followed by a blank node label which is a series of name characters.
//	The characters in the label are built upon PN_CHARS_BASE, liberalized as follows:
//	
//	• The characters _ and digits may appear anywhere in a blank node label.
//	• The character . may appear anywhere except the first or last character.
//	• The characters -, U+00B7, U+0300 to U+036F and U+203F to U+2040 are permitted anywhere except the first character.
//
// ( https://www.w3.org/TR/turtle/#BNodes )
//
// Where, PN_CHARS_BASE is defined as:
//
//	PN_CHARS_BASE ::= [A-Z]           |
//	                  [a-z]           |
//	                  [#x00C0-#x00D6] |
//	                  [#x00D8-#x00F6] |
//	                  [#x00F8-#x02FF] |
//	                  [#x0370-#x037D] |
//	                  [#x037F-#x1FFF] |
//	                  [#x200C-#x200D] |
//	                  [#x2070-#x218F] |
//	                  [#x2C00-#x2FEF] |
//	                  [#x3001-#xD7FF] |
//	                  [#xF900-#xFDCF] |
//	                  [#xFDF0-#xFFFD] |
//	                  [#x10000-#xEFFFF]
//
// ( https://www.w3.org/TR/turtle/#grammar-production-PN_CHARS_BASE )
//
// The specification seems to have an error in it.
// As it is missing [0-9], so that was added, too.
//
// See also: [ParseLabelString].
func ParseLabelBytes(value []byte) (Label, error) {
	var str string = unsafe.String(unsafe.SliceData(value), len(value))

	return ParseLabelString(str)
}

// ParseLabelString parses a string for a blank-node-label from RDF (resource description framework) technologies, such as:
// JSON-LD,
// N-Quads,
// N-Triples,
// RDF/XML,
// RDFa,
// TriG,
// Turtle,
// etc.
//
// A blank-node-label is what comes after the "_:" is a blank-node-identifier.
//
// For example, if the blank-node-identifier is:
//
//	_:abcdef12345
//
// Then its blank-node-label is:
//
//	abcdef12345
//
// This uses the definition found in RDF 1.1 Turtle (https://www.w3.org/TR/turtle/):
//
//	 RDF blank nodes in Turtle are expressed as _: followed by a blank node label which is a series of name characters.
//	The characters in the label are built upon PN_CHARS_BASE, liberalized as follows:
//	
//	• The characters _ and digits may appear anywhere in a blank node label.
//	• The character . may appear anywhere except the first or last character.
//	• The characters -, U+00B7, U+0300 to U+036F and U+203F to U+2040 are permitted anywhere except the first character.
//
// ( https://www.w3.org/TR/turtle/#BNodes )
//
// Where, PN_CHARS_BASE is defined as:
//
//	PN_CHARS_BASE ::= [A-Z]           |
//	                  [a-z]           |
//	                  [#x00C0-#x00D6] |
//	                  [#x00D8-#x00F6] |
//	                  [#x00F8-#x02FF] |
//	                  [#x0370-#x037D] |
//	                  [#x037F-#x1FFF] |
//	                  [#x200C-#x200D] |
//	                  [#x2070-#x218F] |
//	                  [#x2C00-#x2FEF] |
//	                  [#x3001-#xD7FF] |
//	                  [#xF900-#xFDCF] |
//	                  [#xFDF0-#xFFFD] |
//	                  [#x10000-#xEFFFF]
//
// ( https://www.w3.org/TR/turtle/#grammar-production-PN_CHARS_BASE )
//
// The specification seems to have an error in it.
// As it is missing [0-9], so that was added, too.
//
// See also: [ParseLabelString].
func ParseLabelString(value string) (Label, error) {
	if "" == value {
		return Label{}, ErrEmptyString
	}

	{
		r0, _ := utf8.DecodeRuneInString(value)

		switch {
		case '\u002E' == r0                    || // '.' Full Stop
		     '\u002D' == r0                    || // '-' Hyphen-Minus
		     '\u00B7' == r0                    || // '·' Middle Dot
		    ('\u0300' <= r0 && r0 <= '\u036F') ||
		    ('\u203F' <= r0 && r0 <= '\u2040'):
			return Label{}, erorr.Errorf("failed to parse blank-node-label %q due to first character %q (%U): %w", value, r0, r0, ErrLabelFirstCharacterNotAllowed)
		}
	}

	{
		rLast, _ := utf8.DecodeLastRuneInString(value)

		switch rLast {
		case '\u002E': // '.' Full Stop
			return Label{}, erorr.Errorf("failed to parse blank-node-label %q due to last character %q (%U): %w", value, rLast, rLast, ErrLabelLastCharacterNotAllowed)
		}

	}

	for index, r := range value {
		switch {
		case '_' == r:
			// nothing here.
		case '.' == r:
			// nothing here.
		case '-' == r:
			// nothing here.
		case '\u00B7' == r:
			// nothing here.
		case     '\u0300' <= r && r <= '\u036F':
			// nothing here.
		case     '\u203F' <= r && r <= '\u2040':
			// nothing here.



		case          'A' <= r && r <= 'Z':
			// nothing here.
		case          'a' <= r && r <= 'z':
			// nothing here.
		case     '\u00C0' <= r && r <= '\u00D6':
			// nothing here.
		case     '\u00D8' <= r && r <= '\u00F6':
			// nothing here.
		case     '\u00F8' <= r && r <= '\u02FF':
			// nothing here.
		case     '\u0370' <= r && r <= '\u037D':
			// nothing here.
		case     '\u037F' <= r && r <= '\u1FFF':
			// nothing here.
		case     '\u200C' <= r && r <= '\u200D':
			// nothing here.
		case     '\u2070' <= r && r <= '\u218F':
			// nothing here.
		case     '\u2C00' <= r && r <= '\u2FEF':
			// nothing here.
		case     '\u3001' <= r && r <= '\uD7FF':
			// nothing here.
		case     '\uF900' <= r && r <= '\uFDCF':
			// nothing here.
		case     '\uFDF0' <= r && r <= '\uFFFD':
			// nothing here.
		case '\U00010000' <= r && r <= '\U000EFFFF':
			// nothing here.



		case          '0' <= r && r <= '9':
			// nothing here.



		default:
			return Label{}, erorr.Errorf("failed to parse blank-node-label %q due to %s character %q (%U): %w", value, orden.FormatInt64(int64(1+index)), r, r, ErrLabelFirstCharacterNotAllowed)
		}
	}


	return someLabel(value), nil
}

// MarshalText makes [Label] fit [encoding.TextMarshaler].
func (receiver Label) MarshalText() (text []byte, err error) {
	value, found := receiver.optional.Get()
	if !found {
		return nil, ErrEmptyLabel
	}

	return []byte(value), nil
}

// String makes [Label] if [fmt.Stringer].
func (receiver Label) String() string {
	return receiver.optional.GetElse("")
}

// UnmarshalText makes [Label] fit [encoding.TextUnmarshaler].
func (receiver *Label) UnmarshalText(text []byte) error {
	if nil == receiver {
		panic(ErrNilReceiver)
	}

	result, err := ParseLabelBytes(text)
	if nil != err {
		return err
	}

	*receiver = result
	return nil
}
