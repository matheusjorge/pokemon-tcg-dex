# Pokémon TCGDex

Pokémon TCGDex is a management system for Pokémon TCG collections, designed to make it easier to organize, track, and explore your cards. 
Whether you're a casual collector or a dedicated TCG enthusiast, TCGDex helps you manage your collection effortlessly.

I created the system because I thought it would be nice to horn my skills in different languages. The projects is in constant evolution and any feedback great!

## Features

- **Collection Management**: Keep track of your Pokémon cards and organize them with ease.
- **Image-Based Search**: Snap a picture of your card and search the database for instant identification.

## Requirements

- **Golang**: Version 1.22 or higher
- **Python**: Version 3.11 or higher

## Getting Started

Follow these steps to set up and run the application.

### Setup

1. Install dependencies:
   ```bash
   pip install uv
   ```

2. Bring up the application services:
   ```bash
   make docker-up
   ```

3. Apply database migrations:
   ```bash
   make migrate-up
   ```

4. Build the application:
   ```bash
   make build
   ```

5. In a separate terminal, start the sidecar service (it is needed to compute the embeddings):
   ```bash
   make sidecar
   ```

6. Initialize the database:
   ```bash
   make db-setup
   ```

### Running the Application

In separate terminal windows, run the following commands:

1. Start the sidecar service (if it is not already started):
   ```bash
   make sidecar
   ```

2. Launch the API:
   ```bash
   make api
   ```

3. Run the front-end:
   ```bash
   make front
   ```

---

With TCGDex, managing your Pokémon TCG collection has never been easier. Happy collecting!

