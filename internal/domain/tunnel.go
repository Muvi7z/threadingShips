package domain

import (
	"errors"
	"sync"
)

type Tunnel struct {
	Ships []*Ship
	mu    *sync.Mutex
	//Harbor *Harbor
}

func NewTunnel(mutex *sync.Mutex) Tunnel {
	return Tunnel{
		Ships: nil,
		mu:    mutex,
		//Harbor: h,
	}
}

func (t *Tunnel) AddShip(s *Ship) error {
	if len(t.Ships) < 5 {
		t.mu.Lock()
		t.Ships = append(t.Ships, s)
		defer t.mu.Unlock()
		return nil
	}
	return errors.New("Error, tunnel is full")
}

func (t *Tunnel) GetAndDeleteShipByFruit(fruit Fruit) *Ship {
	t.mu.Lock()
	defer t.mu.Unlock()

	if len(t.Ships) > 0 {
		for i, ship := range t.Ships {
			if ship.Fruit == fruit {
				t.Ships = append(t.Ships[:i], t.Ships[i+1:]...)
				//newShip := <-t.Harbor.ShipChan
				//t.AddShip(&newShip)
				return ship
			}
		}
	}

	return nil
}
