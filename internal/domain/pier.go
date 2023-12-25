package domain

import (
	"fmt"
	"sync"
	"time"
)

type Pier struct {
	Name     string
	Fruit    Fruit
	Tunnel   *Tunnel
	wg       *sync.WaitGroup
	Quantity int
}

func NewPier(name string, fruit Fruit, tunnel *Tunnel, wg *sync.WaitGroup) Pier {
	return Pier{
		Name:     name,
		Fruit:    fruit,
		Tunnel:   tunnel,
		Quantity: 0,
		wg:       wg,
	}

}

func (p *Pier) StartLoading() {
	for {
		time.Sleep(time.Millisecond * 500)
		ship := p.Tunnel.GetAndDeleteShipByFruit(p.Fruit)
		if ship != nil {
			if !ship.checkLoading() {
				println(fmt.Sprintf("%s: Начало загрузки %s...: %v", p.Name, p.Fruit.String(), ship))

				p.LoadingShip(ship)
				println(fmt.Sprintf("%s: Погрузка %s завершена: %v", p.Name, p.Fruit.String(), ship))
			}

		}
	}
	p.wg.Done()

}

func (p *Pier) LoadingShip(ship *Ship) {
	for i := 0; i < ship.Size; i += 10 {
		ship.loadingProduct(10)
		//println(fmt.Sprintf("loading ship: %v", ship))
		time.Sleep(time.Second)
	}
}
