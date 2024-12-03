import requests
import streamlit as st
import duckdb
import pandas as pd


@st.cache_resource
def connect():
    return duckdb.connect("tmp/duck.db")


conn = connect()


@st.cache_data
def setup():
    res = requests.get("http://0.0.0.0:8001/v1/cards/get/all")
    if res.status_code == 200:
        df = pd.DataFrame(res.json())
        df["number"] = pd.to_numeric(df["number"], errors="coerce")
        # Drops
        conn.sql("DROP TABLE IF EXISTS cards")
        conn.sql("DROP VIEW IF EXISTS set_ids")
        conn.sql("DROP VIEW IF EXISTS artists")
        conn.sql("DROP VIEW IF EXISTS pokemon")

        # Create views and tables
        conn.sql("CREATE TABLE cards AS SELECT * FROM df")
        conn.sql(
            "CREATE VIEW set_ids AS SELECT DISTINCT set_id FROM cards ORDER BY set_id"
        )
        conn.sql(
            "CREATE VIEW artists AS SELECT DISTINCT artist FROM cards ORDER BY artist"
        )
        conn.sql(
            "CREATE VIEW pokemon AS SELECT DISTINCT name FROM cards ORDER BY national_pokedex_number"
        )
        conn.sql(
            "CREATE TABLE IF NOT EXISTS collection (id TEXT PRIMARY KEY, collected BOOL)"
        )
        conn.sql(
            "INSERT INTO collection SELECT id, false AS collected FROM cards ON CONFLICT DO NOTHING"
        )


def update_collected(card_id: str):
    conn.sql(
        f"""
        UPDATE collection
        SET collected =
        (SELECT NOT collected FROM collection WHERE id = '{card_id}')
        WHERE id='{card_id}'
        """
    )
