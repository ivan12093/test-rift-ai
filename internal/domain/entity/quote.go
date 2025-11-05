package entity

type Quote struct {
	Text string
}

func NewQuote(text string) *Quote {
	return &Quote{Text: text}
}
