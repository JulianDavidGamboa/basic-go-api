package model

// Community estructura de una comunidad
type Community struct {
	Name string `json:"name"`
}

// Communities slice de community
type Communities []Community

// Person estructura de una persona
type Person struct {
	Name string `json:"name"`
	Age  uint8  `json:"age"`
	// Communities comunidades a las que pertenece una persona
	Communities Communities `json:"communities"`
}

// Persons slice de personas
type Persons []Person
