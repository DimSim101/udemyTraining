package vis

// MyName is exported because it starts with a capital letter
var MyName = "Todd"
var yourName = "Future Rock Star Programmer"

// externally visibility naming scheme applies to constants too

// externally visible
const Q = 42

// not visible
const q = 42

type nonVisiblePerson struct {
	first string
	last string
	age int
}

type VisiblePerson struct {
	first string
	last string
	age int
}
