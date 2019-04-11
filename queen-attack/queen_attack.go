package queenattack

import (
	"fmt"
)

func sqVal(s string) error {
	if len(s) != 2 ||
		s[0] < 'a' || s[0] > 'h' ||
		s[1] < '1' || s[1] > '8' {
		return fmt.Errorf("Square value invalid: %q", s)
	}

	return nil
}

func CanQueenAttack(w, b string) (attack bool, err error) {
	if err = sqVal(w); err != nil {
		return false, err
	}
	if err = sqVal(b); err != nil {
		return false, err
	}
	if w == b {
		return false, fmt.Errorf("both queens are on same square")
	}
	if w[0] == b[0] || w[1] == b[1] {
		return true, nil
	}
	df := w[0] - b[0]
	dr := w[1] - b[1]
	return df == dr || df+dr == 0,
		nil
}
