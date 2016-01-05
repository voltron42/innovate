package model

type Method int

const (
	GET Method = iota
	PUT
	POST
	DELETE
)

var methods = []string{"GET","PUT","POST","DELETE"}

func (m Method) String() string {
	return methods[m]
}
