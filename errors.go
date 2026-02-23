package blanknode

import (
	"codeberg.org/reiver/go-erorr"
)

const (
	ErrIdentifierPrefixNotFound      = erorr.Error("blank-node-identifier prefix (\"_:\") not found")
	ErrLabelFirstCharacterNotAllowed = erorr.Error("blank-node-label first character not allowed")
	ErrLabelLastCharacterNotAllowed  = erorr.Error("blank-node-label last character not allowed")
	ErrEmptyIdentifier               = erorr.Error("empty blank-node-identifier")
	ErrEmptyLabel                    = erorr.Error("empty blank-node-label")
	ErrEmptyString                   = erorr.Error("empty string")
	ErrNilReceiver                   = erorr.Error("nil receiver")
)
