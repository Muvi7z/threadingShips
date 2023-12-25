package domain

import "fmt"

type Ship struct {
	Fruit Fruit
	Size  int
	Count int
}

func NewShip(fruit Fruit, size int) Ship {
	return Ship{
		Fruit: fruit,
		Size:  size,
	}
}

func (s *Ship) loadingProduct(count int) string {
	if s.Count+count <= s.Size {
		s.Count += count
		return fmt.Sprintf("На корабль успешно загружен товар %d", count)
	}
	return "Корабль полностью загружен"
}

func (s *Ship) checkLoading() bool {
	return s.Count >= s.Size
}
