package repositories

type Email struct {
	To      []string
	From    string
	Bcc     []string
	cc      []string
	Title   string
	Message string
}
