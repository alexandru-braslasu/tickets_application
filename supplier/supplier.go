package supplier

type Supplier interface {
	ReadLines() ([]string, error)
	WriteResult(data any) error
	GetCapitalFrom() string
	GetCapitalTo() string
}