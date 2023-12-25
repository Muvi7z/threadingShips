package domain

import (
	"fmt"
	"math/rand"
	"time"
)

type Harbor struct {
	Ships    []Ship
	Tunnel   *Tunnel
	ShipChan chan Ship
}

func NewHarbor(tunnel *Tunnel) Harbor {
	return Harbor{
		Ships:    nil,
		ShipChan: make(chan Ship),
		Tunnel:   tunnel,
	}
}

func (h *Harbor) GenerateShips(count int) {
	for i := count; i > 0 || count == -1; i-- {
		time.Sleep(200 * time.Millisecond)
		fruit := rand.Intn(3)
		size := 1 + rand.Intn(10)

		ship := Ship{
			Fruit: Fruit(fruit),
			Size:  (size) * 10,
			Count: 0,
		}
		println(fmt.Sprintf("harbor ships len: %v", len(h.Ships)))
		h.Ships = append(h.Ships, ship)
		_ = h.Tunnel.AddShip(&ship)

		//h.ShipChan <- ship
	}
}

func (h *Harbor) addToTunnel() {
	ship := h.getShip()
	for {
		err := h.Tunnel.AddShip(ship)
		if err == nil {
			ship = h.getShip()
		}

		time.Sleep(time.Millisecond * 500)
	}

}

func (h *Harbor) getShip() *Ship {
	if len(h.Ships) > 0 {
		idShip := 1 + rand.Intn(len(h.Ships))
		ship := h.Ships[idShip]

		//newShip := <-t.Harbor.ShipChan
		//t.AddShip(&newShip)
		return &ship
	}
	return nil
}

func (h *Harbor) deleteShipByID(id int) {
	h.Ships = append(h.Ships[:id], h.Ships[id+1:]...)

}
