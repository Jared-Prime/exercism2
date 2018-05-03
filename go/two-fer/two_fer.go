// Package twofer makes it nice to share
package twofer

import "fmt"

// ShareWith shares with the person named.
func ShareWith(name string) string {
	if name == "" {
		name = "you"
	}

	return fmt.Sprintf("One for %s, one for me.", name)
}
