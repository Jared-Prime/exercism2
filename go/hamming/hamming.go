package hamming

import "errors"

// Distance returns the number of different codes
// between two strings
func Distance(a, b string) (int, error) {
	var distance int

	if len(a) != len(b) {
		return distance, errors.New("unequal strands")
	}

	for i, acid := range []byte(a) {
		if acid != b[i] {
			distance++
		}
	}

	return distance, nil
}
