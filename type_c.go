package sage100c

type TypeC struct {
	Type                      string      `fixed:"1,1"`
	AccountNumber             string      `fixed:"2,9"`
	Description               string      `fixed:"10,39"`
	AlphaKey                  string      `fixed:"40,46"`
	DebitAmount               Amount      `fixed:"47,59"`
	CreditAmount              Amount      `fixed:"60,72"`
	DebitAmount2              Amount      `fixed:"73,85"`
	CreditAmount2             Amount      `fixed:"86,98"`
	ControlAccount            string      `fixed:"99,106"`
	AddressLine1              string      `fixed:"107,136"`
	AddressLine2              string      `fixed:"137,166"`
	Telephone                 string      `fixed:"167,196"`
	UpdateExistingAccountFlag Flag        `fixed:"217,217"`
	AccountType               AccountType `fixed:"218,218"`
	CentralizeAccount         string      `fixed:"219,219"`
	Bank                      string      `fixed:"220,249"`
	RIB                       string      `fixed:"250,279"`
	PaymentMethod             string      `fixed:"280,281"`
	PaymentDays               string      `fixed:"282,283"`
	// PaymentTerms              string      `fixed:"284,285"`
	// IssuedDate                Date        `fixed:"286,287"`
	VATCode         string `fixed:"288,289"`
	ContraAccount   string `fixed:"290,297"`
	PaymentDays3    int    `fixed:"298,300"`
	TaxInvoiceFlag  Flag   `fixed:"301,301"`
	Fax             string `fixed:"302,321"`
	PaymentMode     string `fixed:"322,325"`
	Group4          string `fixed:"326,333"`
	SIRET           string `fixed:"334,347"`
	EditM2          string `fixed:"348,348"`
	Profession      string `fixed:"349,378"`
	Country         string `fixed:"379,428"`
	CashJournalCode string `fixed:"429,431"`
	Company         Flag   `fixed:"432,432"`
	Creditworthy    Flag   `fixed:"433,433"`
	IBAN            string `fixed:"434,437"`
	BIC             string `fixed:"438,448"`
	FreshChargeCode int    `fixed:"449,450"`
	SEPA            string `fixed:"451,453"`
}
