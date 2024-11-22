CREATE TYPE trait AS (
  name TEXT,
  description TEXT
);

CREATE TYPE ability AS (
  name TEXT,
  description TEXT,
  type TEXT
);

CREATE TYPE attack AS (
  cost TEXT[],
  name TEXT,
  description TEXT,
  damage TEXT,
  converted_energy_cost INTEGER
);

CREATE TYPE type_relationship AS (
  type TEXT,
  value TEXT
);


CREATE TABLE IF NOT EXISTS cards (
  id TEXT NOT NULL PRIMARY KEY,
  name TEXT NOT NULL,
  set_id TEXT,
  number INTEGER,
  artist TEXT,
  rarity TEXT,
  national_pokedex_number INTEGER[],
  image_small TEXT,
  image_large TEXT,
  supertype TEXT NOT NULL,
  subtypes TEXT[],
  level TEXT,
  hp INT,
  types TEXT[],
  evolves_from TEXT,
  evolves_to TEXT[],
  rules TEXT,
  ancient_trait trait[],
  abilities ability[],
  attacks attack[],
  weaknesses type_relationship[],
  resistances type_relationship[],
  retreat_cost TEXT[],
  converted_retreat_cost INTEGER
);
