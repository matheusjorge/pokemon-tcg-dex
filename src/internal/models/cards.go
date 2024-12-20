package models

import (
	"strconv"
)

type Attack struct {
	Cost                []string `json:"cost"`
	ConvertedEnergyCost int      `json:"convertedEnergyCost"`
	Damage              string   `json:"damage"`
	Name                string   `json:"name"`
	Text                string   `json:"text"`
}

type TypeRelationship struct {
	Type  string `json:"type"`
	Value string `json:"value"`
}

type AncientTrait struct {
	Name string `json:"name"`
	Text string `json:"text"`
}

type Ability struct {
	Name string `json:"name"`
	Text string `json:"text"`
	Type string `json:"type"`
}

type Images struct {
	Small string `json:"small"`
	Large string `json:"large"`
}

type CardJson struct {
	Id                    string             `json:"id"`
	Name                  string             `json:"name"`
	Supertype             string             `json:"supertype"`
	Subtypes              []string           `json:"subtypes"`
	Level                 string             `json:"level"`
	HP                    string             `json:"hp"`
	Types                 []string           `json:"types"`
	EvolvesFrom           string             `json:"evolvesFrom"`
	EvolvesTo             []string           `json:"evolvesTo"`
	Rules                 []string           `json:"rules"`
	AncientTrait          AncientTrait       `json:"ancientTrait"`
	Abilities             []Ability          `json:"abilities"`
	Attacks               []Attack           `json:"attacks"`
	Weaknessess           []TypeRelationship `json:"weaknesses"`
	Resistances           []TypeRelationship `json:"resistances"`
	Number                string             `json:"number"`
	Artist                string             `json:"artist"`
	Rarity                string             `json:"rarity"`
	NationalPokedexNumber []int              `json:"nationalPokedexNumbers"`
	RetreatCost           []string           `json:"retreatCost"`
	ConvertedRetreatCost  int                `json:"convertedRetreatCost"`
	Images                Images             `json:"images"`
}

type Card struct {
	Id                    string             `json:"id"`
	Name                  string             `json:"name"`
	Supertype             string             `json:"supertype"`
	Subtypes              []string           `json:"subtypes"`
	Level                 string             `json:"level"`
	HP                    int                `json:"hp"`
	Types                 []string           `json:"types"`
	EvolvesFrom           string             `json:"evolves_from"`
	EvolvesTo             []string           `json:"evolves_to"`
	Rules                 []string           `json:"rules"`
	AncientTrait          AncientTrait       `json:"ancient_trait"`
	Abilities             []Ability          `json:"abilities"`
	Attacks               []Attack           `json:"attack"`
	Weaknessess           []TypeRelationship `json:"weaknesses"`
	Resistances           []TypeRelationship `json:"resistances"`
	SetId                 string             `json:"set_id"`
	Number                string             `json:"number"`
	Artist                string             `json:"artist"`
	Rarity                string             `json:"rarity"`
	NationalPokedexNumber int                `json:"national_pokedex_number"`
	RetreatCost           []string           `json:"retreat_cost"`
	ConvertedRetreatCost  int                `json:"converted_retreat_cost"`
	Images                Images             `json:"images"`
}

func FromJsonToPg(cardJson *CardJson) (*Card, error) {
	hp, err := strconv.Atoi(cardJson.HP)
	if err != nil {
		hp = 0
	}

	var pokedexNumber int
	if len(cardJson.NationalPokedexNumber) > 0 {
		pokedexNumber = cardJson.NationalPokedexNumber[0]
	} else {
		pokedexNumber = -1
	}

	card := Card{
		Id:                    cardJson.Id,
		Name:                  cardJson.Name,
		Supertype:             cardJson.Supertype,
		Subtypes:              cardJson.Subtypes,
		Level:                 cardJson.Level,
		HP:                    hp,
		Types:                 cardJson.Types,
		EvolvesFrom:           cardJson.EvolvesFrom,
		EvolvesTo:             cardJson.EvolvesTo,
		Rules:                 cardJson.Rules,
		AncientTrait:          cardJson.AncientTrait,
		Abilities:             cardJson.Abilities,
		Attacks:               cardJson.Attacks,
		Weaknessess:           cardJson.Weaknessess,
		Resistances:           cardJson.Resistances,
		Number:                cardJson.Number,
		Artist:                cardJson.Artist,
		Rarity:                cardJson.Rarity,
		NationalPokedexNumber: pokedexNumber,
		RetreatCost:           cardJson.RetreatCost,
		ConvertedRetreatCost:  cardJson.ConvertedRetreatCost,
		Images:                cardJson.Images,
	}

	return &card, nil

}
