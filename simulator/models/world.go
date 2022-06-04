package models

import "fmt"

//World data types for simulation maps for  cities names to cities
type World map[string]*City

//AddCity adds a city to the world
//takes city name and insert in world map
func (w World) AddCity(city City) *City {
	w[city.Name] = &city
	return &city
}

// AddNewCity function is more clean than AddNewCity2 function
//func (w World) AddNewCity2(name string) *City {
//	city := NewCity(name)
//	//w[city.Name] = &city
//	//return &city
//	return w.AddCity(city)
//}

//AddNewCity takes city name and insert in world map
func (w World) AddNewCity(name string) *City {
	return w.AddCity(NewCity(name))
}

// Strings representation of world map
func (w World) String() string {
	var cities string
	for _, city := range w {
		cities += fmt.Sprintf("%s\n", city)
	}
	return cities
}
