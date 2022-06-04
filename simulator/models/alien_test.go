package models

import (
	"reflect"
	"sagaAlienInvasion/models"
	"testing"
)

func TestNewAlien(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name string
		args args
		want Alien
	}{
		// TODO: Add test cases.
		{"alien1 test", args{"alien1"}, NewAlien("alien1")},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewAlien(tt.args.name); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewAlien() = %v, want %v", got, tt.want)
			}
		})
	}
}

func NewCityP(name string) *City {
	return &City{
		Vertex:    models.NewVertex(name),
		RoadNames: make(map[string]string),
	}
}

func TestAlien_InvadeCity(t *testing.T) {

	type fields struct {
		Intruder models.Intruder
		city     *City
	}
	type args struct {
		city *City
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
		{"Invade City test", fields{models.NewIntruder("alien1"), NewCityP("alien1city")}, args{city: NewCityP("alien1city")}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &Alien{
				Intruder: tt.fields.Intruder,
				city:     tt.fields.city,
			}
			if a.city.Name != tt.args.city.Name {
				t.Errorf("Alien.InvadeCity() city.Name = %v, want %v", a.city.Name, tt.args.city.Name)
			}
			a.InvadeCity(tt.args.city)
			t.Log("Invade City test is correct")
		})
	}
}
