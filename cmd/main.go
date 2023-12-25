package main

import (
	"fmt"
	"sync"
	"threadingShips/internal/domain"
)

func main() {
	mu := sync.Mutex{}
	wg := &sync.WaitGroup{}
	t := domain.NewTunnel(&mu)

	h := domain.NewHarbor(&t)

	pBanana := domain.NewPier("pBanana", domain.Banana, &t, wg)
	pBread := domain.NewPier("pBread", domain.Bread, &t, wg)
	pClothes := domain.NewPier("pClothes", domain.Clothes, &t, wg)
	//ship1 := domain.Ship{
	//	Fruit: domain.Banana, Size: 40, Count: 0,
	//}
	//ship2 := domain.Ship{
	//	Fruit: domain.Banana, Size: 80, Count: 0,
	//}
	//ship3 := domain.Ship{
	//	Fruit: domain.Clothes, Size: 10, Count: 0,
	//}
	//t.AddShip(&ship1)
	//t.AddShip(&ship2)
	//t.AddShip(&ship3)
	//ship := t.GetAndDeleteShipByFruit(domain.Banana)

	wg.Add(4)

	go h.GenerateShips(-1)

	go pBanana.StartLoading()
	go pBread.StartLoading()
	go pClothes.StartLoading()
	wg.Wait()
	fmt.Println(fmt.Sprintf("tunnel ships: %v", t.Ships))

}
