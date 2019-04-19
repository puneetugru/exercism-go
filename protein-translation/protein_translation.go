// Package protein translates RNA sequences into proteins.
package protein

import (
	"errors"
)

var ErrStop = errors.New("stop")

var ErrInvalidBase = errors.New("invalid base")

func IsEndCodon(c string) bool {
	return c == "UAA" || c == "UAG" || c == "UGA"

}

func FromCodon(c string) (string, error) {
	if IsEndCodon(c) {
		return "", ErrStop
	}

	switch c {
	case
		"AUG":
		return "Methionine", nil
	case
		"UUU",
		"UUC":
		return "Phenylalanine", nil
	case
		"UUA",
		"UUG":
		return "Leucine", nil
	case
		"UCU",
		"UCC",
		"UCA",
		"UCG":
		return "Serine", nil
	case
		"UAU",
		"UAC":
		return "Tyrosine", nil
	case
		"UGU",
		"UGC":
		return "Cysteine", nil
	case
		"UGG":
		return "Tryptophan", nil
	default:
		return "", ErrInvalidBase
	}
}

func FromRNA(s string) ([]string, error) {
	var proteins []string
	bases := []rune(s)
	for i := 0; i < len(bases); i += 3 {
		p, err := FromCodon(string(bases[i : i+3]))
		switch err {
		case ErrStop:
			return proteins, nil
		case ErrInvalidBase:
			return proteins, err
		default:
			proteins = append(proteins, p)
		}
	}
	return proteins, nil
}
