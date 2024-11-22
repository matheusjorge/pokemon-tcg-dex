package models

type CardJson struct {
	Id                    string              `json:"id"`
	Name                  string              `json:"name"`
	Supertype             string              `json:"supertype"`
	Subtypes              []string            `json:"subtypes"`
	Level                 string              `json:"level"`
	HP                    int                 `json:"hp"`
	Types                 []string            `json:"types"`
	EvolvesFrom           string              `json:"evolvesFrom"`
	EvolvesTo             []string            `json:"evolvesTo"`
	Rules                 []string            `json:"rules"`
	AncientTrait          map[string]string   `json:"ancientTrait"`
	Abilities             []map[string]string `json:"abilities"`
	Attacks               []map[string]string `json:"attacks"`
	Weaknessess           []map[string]string `json:"weaknessess"`
	Resistances           []map[string]string `json:"resistances"`
	Set                   map[string]string   `json:"set"`
	Number                int                 `json:"number"`
	Artist                string              `json:"artist"`
	Rarity                string              `json:"rarity"`
	NationalPokedexNumber int                 `json:"nationalPokedexNumber"`
}

type Card struct {
	Id                    string              `json:"id"`
	Name                  string              `json:"name"`
	Supertype             string              `json:"supertype"`
	Subtypes              []string            `json:"subtypes"`
	Level                 string              `json:"level"`
	HP                    int                 `json:"hp"`
	Types                 []string            `json:"types"`
	EvolvesFrom           string              `json:"evolvesFrom"`
	EvolvesTo             []string            `json:"evolvesTo"`
	Rules                 []string            `json:"rules"`
	AncientTrait          map[string]string   `json:"ancientTrait"`
	Abilities             []map[string]string `json:"abilities"`
	Attacks               []map[string]string `json:"attacks"`
	Weaknessess           []map[string]string `json:"weaknessess"`
	Resistances           []map[string]string `json:"resistances"`
	SetId                 string              `json:"set"`
	Number                int                 `json:"number"`
	Artist                string              `json:"artist"`
	Rarity                string              `json:"rarity"`
	NationalPokedexNumber int                 `json:"nationalPokedexNumber"`
}
