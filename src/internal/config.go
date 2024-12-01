package internal

import (
	"os"
	"strconv"
)

type Config struct {
	PokemonTCGAPIKey       string
	PokemonTCGAPIURL       string
	PokemonTCGImagesURL    string
	PostgresURL            string
	RemoteDataURL          string
	DataPath               string
	DataConfigFile         string
	ImageDownloaderWorkers int
	SidecarPort            int
	ServerPort             int
}

const (
	PokemonTCGAPIURL_DEFAULT       = "https://api.pokemontcg.io"
	PokemonTCGImagesURL_DEFAULT    = "https://images.pokemontcg.io"
	PostgresURL_DEFAUT             = "postgresql://ash:pikachu@0.0.0.0:5432/pokedex"
	RemoteDataURL_DEFAULT          = "https://raw.githubusercontent.com/PokemonTCG/pokemon-tcg-data/refs/heads/master"
	DataPath_DEFAULT               = "data"
	DataConfigFile_DEFAULT         = "data_config.json"
	ImageDownloaderWorkers_DEFAULT = "5"
	SidecarPort_DEFAULT            = "8000"
	ServerPort_DEFAULT             = "8001"
)

func LoadConfigs() (*Config, error) {
	config := &Config{}

	config.PokemonTCGAPIKey = os.Getenv("POKEMON_TCG_API_KEY")
	config.PokemonTCGAPIURL = loadEnvVar("POKEMON_TCG_API_URL", PokemonTCGAPIURL_DEFAULT)
	config.PokemonTCGImagesURL = loadEnvVar("POKEMON_TCG_IMAGES_URL", PokemonTCGImagesURL_DEFAULT)
	config.PostgresURL = loadEnvVar("POSTGRES_URL", PostgresURL_DEFAUT)
	config.RemoteDataURL = loadEnvVar("REMOTE_DATA_URL", RemoteDataURL_DEFAULT)
	config.DataPath = loadEnvVar("DATA_PATH", DataPath_DEFAULT)
	config.DataConfigFile = loadEnvVar("DATA_CONFIG_FILE", DataConfigFile_DEFAULT)
	config.ImageDownloaderWorkers, _ = strconv.Atoi(loadEnvVar("IMAGE_DOWNLOADER_WORKERS", ImageDownloaderWorkers_DEFAULT))
	config.SidecarPort, _ = strconv.Atoi(loadEnvVar("SIDECAR_PORT", SidecarPort_DEFAULT))
	config.ServerPort, _ = strconv.Atoi(loadEnvVar("SERVER_PORT", ServerPort_DEFAULT))

	return config, nil
}

func loadEnvVar(key string, dft string) string {
	value, hasValue := os.LookupEnv(key)
	if !hasValue {
		return dft
	}

	return value
}
