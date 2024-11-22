package internal

import "os"

type Config struct {
	PokemonTCGAPIKey    string
	PokemonTCGAPIURL    string
	PokemonTCGImagesURL string
}

const (
	PokemonTCGAPIURL_DEFAULT    = "https://api.pokemontcg.io"
	PokemonTCGImagesURL_DEFAULT = "https://images.pokemontcg.io"
)

func LoadConfigs() (*Config, error) {
	config := &Config{}

	config.PokemonTCGAPIKey = os.Getenv("POKEMON_TCG_API_KEY")
	config.PokemonTCGAPIURL = loadEnvVar("POKEMON_TCG_API_URL", PokemonTCGAPIURL_DEFAULT)
	config.PokemonTCGImagesURL = loadEnvVar("POKEMON_TCG_IMAGES_URL", PokemonTCGImagesURL_DEFAULT)

	return config, nil
}

func loadEnvVar(key string, dft string) string {
	value, hasValue := os.LookupEnv(key)
	if !hasValue {
		return dft
	}

	return value
}
