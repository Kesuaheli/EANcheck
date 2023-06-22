package ean

import "eancheck/gs1"

// EAN13 represents an EAN13 code
type EAN13 struct {
	Prefix gs1.Prefix
	Data   string
	Parity int
}
