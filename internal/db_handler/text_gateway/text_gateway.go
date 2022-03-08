package text_gateway

type TextGateway interface {
	Append(string) error
	Get(int) (string, error)
}