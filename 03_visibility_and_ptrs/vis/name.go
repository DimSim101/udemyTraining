package vis

// MyName is exported because it starts with a capital letter
var MyName = "Todd"
var yourName = "Future Rock Star Programmer"

// externally visibility naming scheme applies to constants too

// Q is Externally visible
const Q = 42

// Q is not Externally visible
const q = 42

type nonVisiblePerson struct {
	first string
	last  string
	age   int
}

// VisiblePerson is Externally visible for use in an example showing visibility.
type VisiblePerson struct {
	first string
	last  string
	age   int
}
