package main

import "fmt"

type Set map[string]struct{}

func NewSet() Set {
	return make(Set)
}

func (s Set) Add(element string) {
	s[element] = struct{}{}
}

func (s Set) Remove(element string) {
	delete(s, element)
}

func (s Set) Contains(element string) bool {
	_, exists := s[element]
	return exists
}

func (s Set) Size() int {
	return len(s)
}

func main() {
	s := NewSet()
	s.Add("abc")
	s.Add("xyz")
	s.Add("mnk")

	fmt.Println("Set contains 'abc':", s.Contains("abc"))
	fmt.Println("Set size:", s.Size())

	s.Remove("xyz")
	fmt.Println("Set contains 'xyz':", s.Contains("banana"))
	fmt.Println("Set size:", s.Size())
}
