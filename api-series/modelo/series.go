package modelo

type Series struct {
	Id        int    `json:"id"`
	Nombre    string `json:"nombre"`
	Genero    string `json:"genero"`
	Capitulos int    `json:"capitulos"`
	Portada   string `json:"portada"`
}
