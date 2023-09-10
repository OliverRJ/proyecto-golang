package entities

type Cliente struct {
	Nombre          string `json:"nombre"`
	FechaNacimiento string `json:"fechaNacimiento"`
	Edad            int64  `json:"edad"`
}
