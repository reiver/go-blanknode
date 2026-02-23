package blanknode

import (
	"encoding"
	"strings"
	"unsafe"
)

// IdentifierPrefix is the prefix at the beginning of all blank-node-identifiers.
//
// It makes blank-node-identifiers look like a URL/URI/IRI with a scheme of "_".
//
// You can see the prefix in these example blank-node-identifiers:
//
//	_:b0
//	_:address84
//	_:n1
//	_:ed7ba470-8e54-465e-825c-99712043e01c
//	_:label123
const IdentifierPrefix string = "_:"

// HasIdentifierPrefix return whether a string beings with a "_:" or not.
//
// All blank-node-identifiers begin with a "_:".
func HasIdentifierPrefix(value string) bool {
	return strings.HasPrefix(value, IdentifierPrefix)

}

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
		if !HasIdentifierPrefix(value) {
			return Identifier{}, ErrIdentifierPrefixNotFound
		}

		str = str[len(IdentifierPrefix):]
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

func (receiver Identifier) Get() (string, bool) {
	label, found := receiver.label.Get()
	if !found {
		return "", false
	}

	return IdentifierPrefix + label, true
}

func (receiver Identifier) IsNothing() bool {
	return receiver.label.IsNothing()
}

// Label returns the blank-node-label for the blank-node-identifier.
func (receiver Identifier) Label() (Label, bool) {
	return receiver.label, !receiver.label.IsNothing()
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
	return IdentifierPrefix + receiver.label.String()
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
