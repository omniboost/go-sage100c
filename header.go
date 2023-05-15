package sage100c

type Header struct {
	// nom de la société (30 caractères maximum)
	Societe string `fixed:"1,30"`
}
