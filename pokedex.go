package main

import (
	"github.com/VirgiVaRu/pokedexcli/internal/PokeAPI"
)

type Pokedex struct {
	caughtPokemon map[string]PokeAPI.Pokemon
}

func (p Pokedex) GetEntry(pokemonName string) (PokeAPI.Pokemon, bool) {
	pokemon, found := p.caughtPokemon[pokemonName]

	if found {
		return pokemon, true
	} else {
		return pokemon, false
	}
}