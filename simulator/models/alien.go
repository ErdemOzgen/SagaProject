package models

import (
	"fmt"
	"sagaAlienInvasion/models"
)

// Alien data struct contains model intruder with composition
// And contains city for invadesion
type Alien struct {
	models.Intruder
	city *City
}

// NewAlien returns a new Alien with taken name
func NewAlien(name string) Alien {
	// FlagDead is defaulted to false
	return Alien{
		Intruder: models.NewIntruder(name),
	}
}

//InvadeCity invades the city for selected alien
func (a *Alien) InvadeCity(city *City) {
	a.Vertex = &city.Vertex
	a.city = city

}

// City returns City for current Alien is occupying
func (a *Alien) City() *City {
	return a.city
}

// IsDead checks if Alien died
func (a *Alien) IsDead() bool {
	return a.SimConditions[DeadCond]
}

// Kill Alien makes alien dead
func (a *Alien) Kill() {
	a.SimConditions[DeadCond] = true
}

//IsInvading checks if Alien is invading a city
func (a *Alien) IsInvading() bool {
	return a.Vertex != nil
}

// IsTrapped checks if Alien is trapped
// For some case aliens are trapped when they are not invading a city becuase of road links blown
func (a *Alien) IsTrapped() bool {
	if !a.IsInvading() {
		return false
	}
	for _, v := range a.City().Vertexs {
		c := City{Vertex: *v}
		if !c.IsDestroyed() {
			return false
		}
	}
	return true
}

//String representation for Alien
func (a *Alien) String() string {
	return fmt.Sprintf("nameofAllien=%s currentVertexofAlien={%s}\n", a.Name, a.Vertex)
}
