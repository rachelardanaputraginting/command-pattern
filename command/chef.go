package command

type FriedRiceParams struct {
	Spiciness int
	Saltyness int
	ExtraEgg bool
}

type Cooker interface {
	CookFriedRice(params FriedRiceParams)
}

type LordAdi struct {

}

func (la LordAdi) CookFriedRice(params FriedRiceCommand){
	fnt.Println("Lord adi cooks fried rice %y", params)
}