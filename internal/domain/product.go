package domain

type Product struct {
	Name string `json:"name"`
}

type Fruit int

const (
	Banana Fruit = iota
	Bread
	Clothes
)

func (f Fruit) String() string {
	switch f {
	case Banana:
		return "Банан"
	case Bread:
		return "Хлеб"
	case Clothes:
		return "Одежда"
	}
	return "unknown"
}
