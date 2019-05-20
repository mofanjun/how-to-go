package engine

type ParseFunc func([]byte, string) ParseResult

type Request struct {
	Url	string
	ParseFunc ParseFunc
}

type ParseResult struct {
	Requests []Request
	Items	[]Item
}

type Item struct {
	Id		string
	Type 	string
	Url		string
	PayLoad interface{}
}

func NilParser([]byte) ParseResult{
	return ParseResult{}
}