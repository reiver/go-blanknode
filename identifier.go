package blanknode

import (
	"encoding"
	"strings"
	"unsafe"
)

const prefix string = "_:"

// Identifier represents a blank-node-identifier from RDF (resource description framework) technologies, such as:
// JSON-LD,
// N-Quads,
// N-Triples,
// RDF/XML,
// RDFa,
// TriG,
// Turtle,
// etc.
//
// Here are some example blank-node-identifiers:
//
//	_:b0
//	_:address84
//	_:n1
//	_:ed7ba470-8e54-465e-825c-99712043e01c
//	_:label123
type Identifier struct {
	label Label
}

var (
	_ encoding.TextMarshaler   = Identifier{}
	_ encoding.TextUnmarshaler = &Identifier{}
)

// NoLIdentifier returns an empty [Identifier].
//
// [Identifier] is an optional-type
// —
// also known as a option-type or maybe-type elsewhere.
func NoIdentifier() Identifier {
	return Identifier{}
}

func someIdentifier(label Label) Identifier {
	return Identifier{label:label}
}

// ParseIdentifierString parses the string for a Blank Node Identifier.
//
// A Blank Node Identifier looks like these:
//
//	_:b0
//	_:address84
//	_:n1
//	_:ed7ba470-8e54-465e-825c-99712043e01c
//	_:label123
func ParseIdentifierString(value string) (Identifier, error) {
	if "" == value {
		return Identifier{}, ErrEmptyString
	}

	var str string

	{
		if !strings.HasPrefix(value, prefix) {
			return Identifier{}, ErrIdentifierPrefixNotFound
		}

		str = str[len(prefix):]
	}

	label, err := ParseLabelString(str)
	if nil != err {
		return Identifier{}, err
	}

	return someIdentifier(label), nil
}

// ParseIdentifierBytes parses the []byte for a Blank Node Identifier.
//
// A Blank Node Identifier looks like these:
//
//	_:b0
//	_:address84
//	_:n1
//	_:ed7ba470-8e54-465e-825c-99712043e01c
//	_:label123
func ParseIdentifierBytes(value []byte) (Identifier, error) {
	var str string = unsafe.String(unsafe.SliceData(value), len(value))

	return ParseIdentifierString(str)
}

func (receiver Identifier) IsNothing() bool {
	return receiver.label.IsNothing()
}

// MarshalText makes [Identifier] fit [encoding.TextMarshaler].
func (receiver Identifier) MarshalText() (text []byte, err error) {
	if receiver.label.IsNothing() {
		return nil, ErrEmptyIdentifier
	}

	return []byte(receiver.String()), nil
}

// String makes [Identifier] fit [fmt.Stringer].
func (receiver Identifier) String() string {
	return prefix + receiver.label.String()
}

// UnmarshalText makes [Identifier] fit [encoding.TextUnmarshaler].
func (receiver *Identifier) UnmarshalText(text []byte) error {
	if nil == receiver {
		panic(ErrNilReceiver)
	}

	result, err := ParseIdentifierBytes(text)
	if nil != err {
		return err
	}

	*receiver = result
	return nil
}
