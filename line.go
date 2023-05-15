package sage100c

type Lines []Line

type Line struct {
	CodeJournal string `fixed:"1,3"`
	// JJMMAA
	DatePiece Date `fixed:"4,9"`
	// FC: Facture client
	// AC: Avoir client
	// RC: Reglement client
	// FF: Facture fournisseur
	// AF: Avoir fournisseur
	// RF: Reglement fourn
	// OD: Operations diverses
	// Blanc
	TypePiece     string `fixed:"10,11"`
	CompteGeneral string `fixed:"12,24"`
	// G ou Blanc: General
	// A: Analytique
	// X: Auxiliaire
	// E: Échéance
	TypeCompte                   string `fixed:"25,25"`
	CompteAuxiliaireOuAnalytique string `fixed:"26,38"`
	// Reference Piece ou Numero de piece selon option en import
	ReferenceEcriture string `fixed:"39,51"`
	LibelleEcriture   string `fixed:"52,76"`
	// E: Espece
	// C: Cheque
	// T: Traite
	// U: Carte bleue
	// L: LCR
	// V: Virement
	// P: Prelevement
	// S: Aucun
	ModePaiement string `fixed:"77,77"`
	// JJMMAA
	DateEcheance string `fixed:"78,83"`
	// Montant en monnaie de tenue de compte
	MontantCompte Amount `fixed:"84,104,right,0"`
	// N: Normale
	// A: A-Noveau
	// S: Simulation
	TypeEcriture string `fixed:"105,105"`
	// Non Utilise
	NumeroPiece string `fixed:"106,112"`
	// Variable selon les produits
	ZoneReservee string `fixed:113,138`
	// Code ISO de la devise dans laquelle est exprimee le montant de
	// l'ecriture.
	CodeISO string `fixed:"139,141"`
	// Montant en devise
	MontantDevise Amount `fixed:142,161,right,0"`
	CodeISODevise string `fixed:162,164`
	// Numero du Plan analytique dans la compatabilite 100
	AxeAnalytique string `fixed:"165,174"`
	// Non utilise
	SectionAnalytique string `fixed:"175,199"`
}
