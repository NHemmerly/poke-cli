package main

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
