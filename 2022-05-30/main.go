package may302022

import "fmt"

func Tasks() {
	s := "Cassidy"
	ok := allUnique(s)
	fmt.Printf("The chars in the string '%s' are unique: %t.\n", s, ok)
}

func allUnique(str string) bool {
	m := make(map[int32]struct{})
	for _, c := range str {
		if _, ok := m[c]; ok {
			return false
		}
		m[c] = struct{}{}
	}
	return true
}
