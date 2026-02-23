package blanknode

import (
	"codeberg.org/reiver/go-erorr"
)

const (
	ErrLabelFirstCharacterNotAllowed = erorr.Error("blank-node-label first character not allowed")
	ErrLabelLastCharacterNotAllowed  = erorr.Error("blank-node-label last character not allowed")
	ErrEmptyLabel                    = erorr.Error("empty blank-node-label")
	ErrEmptyString                   = erorr.Error("empty string")
	ErrNilReceiver                   = erorr.Error("nil receiver")
)
