package sage100c

type TypeM struct {
	Type                  string `fixed:"1,1"`
	AccountNumber         string `fixed:"2,9"`
	JournalCode           string `fixed:"10,11"`
	FolioNumber           string `fixed:"12,14"`
	Date                  Date   `fixed:"15,20"` // JJMMAA
	DescriptionCode       string `fixed:"21,21"`
	Description           string `fixed:"22,41"`
	Side                  string `fixed:"42,42"`
	Amount                Amount `fixed:"43,55"`
	ContraAccount         string `fixed:'56,63'`
	DueDate               Date   `fixed:"64,69"` // JJMMAA
	LetteringCode         string `fixed:"70,71"`
	StatisticsCode        string `fixed:"72,74"`
	DocumentCode5         string `fixed:"75,79"`
	CompanyCode           string `fixed:"80,89"`
	Quantity              string `fixed:"90,99"`
	DocumentCode8         string `fixed:"100,107"`
	Currency              string `fixed:"108,110"`
	JournalCode3          string `fixed:"111,113"`
	VATInclusive          Flag   `fixed:"114,114"`
	VATCode               string `fixed:"115,115"`
	VATCalculation        string `fixed:"116,116"`
	AccountingDescription string `fixed:"117,147"`
	Reference             string `fixed:"149,158"`
	Reserved1             string `fixed:"159,168"`
	FCAmount              Amount `fixed:"169,181"`
	AttachmentFilename    string `fixed:"182,193"`
	Quantity2             int    `fixed:"194,203,omitempty"`
	NumUnique             string `fixed:"204,213"`
	OperatorCode          string `fixed:"214,217"`
	SystemDate            string `fixed:"218,231"`
}
