package main

type Object struct {
	name string
	next []*Object
}

func (s Object) findNode(name string) *Object {
	if s.name == name {
		return &s
	} else {
		for _, object := range s.next {
			o := object.findNode(name)
			if o != nil {
				return o
			}
		}
	}

	return nil
}
