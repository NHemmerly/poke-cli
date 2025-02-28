package main

import "fmt"

type Pokedex struct {
	count   int
	Pokemon map[string]Pokemon
}

func NewPokedex() Pokedex {
	pokedex := Pokedex{count: 0, Pokemon: make(map[string]Pokemon)}
	return pokedex
}

type DexInterface interface {
	Add(poke Pokemon)
	Get(name string) (Pokemon, bool)
}

func (p *Pokedex) Add(poke *Pokemon) {
	p.Pokemon[poke.Name] = *poke
}

func (p *Pokedex) Get(name string) (*Pokemon, bool) {
	if val, ok := p.Pokemon[name]; ok {
		return &val, true
	}
	return nil, false
}

func PrintPokemon(poke *Pokemon) {
	fmt.Printf(`Name: %s
Height: %d
Weight: %d
Stats:
%s
Types:
%s`, poke.Name, poke.Height, poke.Weight, GetStats(poke), GetTypes(poke))
}

func GetStats(poke *Pokemon) string {
	var out string
	for _, stat := range poke.Stats {
		out += fmt.Sprintf("\t-%s: %d\n", stat.Stat.Name, stat.BaseStat)
	}
	return out
}

func GetTypes(poke *Pokemon) string {
	var out string
	for _, typ := range poke.Types {
		out += fmt.Sprintf("\t- %s\n", typ.Type.Name)
	}
	return out
}
