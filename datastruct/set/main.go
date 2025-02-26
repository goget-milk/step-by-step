package main

import "fmt"

type Set struct {
	Elements map[string]struct{}
}

func NewSet() *Set {
	var set Set
	set.Elements = make(map[string]struct{})
	return &set
}

func (s *Set) Add(element string) {
	s.Elements[element] = struct{}{}
}

func (s *Set) Delete(element string) error {
	if _, ok := s.Elements[element]; !ok {
		return fmt.Errorf("element %s not found", element)
	}
	delete(s.Elements, element)
	return nil
}

func (s *Set) Contains(element string) bool {
	_, ok := s.Elements[element]
	return ok
}

func (s *Set) List() {
	for element := range s.Elements {
		fmt.Println(element)
	}
}

func main() {
	set := NewSet()
	set.Add("earth")
	set.Add("venus")
	set.Add("mars")
	set.Add("earth")
	set.Delete("venus")

	set.List()
	fmt.Println(set.Contains("earth"))
}
