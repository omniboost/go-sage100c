package sage100c

type TypeI struct {
	Type       string `fixed:"1,1"`
	Percentage int    `fixed:"2,6"`
	Amount     int    `fixed:"7,12"`
	CodeCentre string `fixed:"20,29"`
	CodeNature string `fixed:"30,39"`
}
