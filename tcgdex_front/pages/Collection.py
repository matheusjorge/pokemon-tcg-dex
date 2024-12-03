from streamlit_extras.grid import grid
import streamlit as st

from tcgdex_front.internal.caching import cached_image
import tcgdex_front.internal.repositories.duckdb_repository as duckdb


page_key = "collection/page"
set_id_key = "collection/set_id"

duckdb.setup()


@st.cache_data
def set_options():
    return duckdb.conn.sql("SELECT * FROM set_ids").to_df()["set_id"].tolist()


def on_change_set_id():
    st.session_state[page_key] = 0


def on_change_page():
    st.session_state[page_key] = st.session_state["collection/pills"]


@st.cache_data
def filtered_df(set_id: str):
    return duckdb.conn.sql(
        f"""
        SELECT 
              cards.id
            , cards.images
            , collected
        FROM
            cards
        LEFT JOIN
            collection
            ON (cards.id = collection.id)
        WHERE 
            set_id = '{set_id}' 
        ORDER BY
            number
        """
    ).to_df()


@st.cache_data
def filtered_images(set_id: str):
    return (filtered_df(set_id)["images"].apply(lambda x: x["small"])).tolist()


@st.cache_data
def filtered_collection(set_id: str):
    return filtered_df(set_id)["collected"].tolist()


@st.cache_data
def filtered_ids(set_id: str):
    return filtered_df(set_id)["id"].tolist()


if st.session_state.get(set_id_key):
    default_set_id = set_options().index(st.session_state[set_id_key])
else:
    default_set_id = 0

left, _ = st.columns([2, 8])
with left:
    set_id = st.selectbox(
        "Select Set",
        options=set_options(),
        index=default_set_id,
        placeholder="Select a set",
        on_change=on_change_set_id,
    )
    st.session_state[set_id_key] = set_id
collection_layout = grid(5, vertical_align="center")

images = filtered_images(set_id)
collected = filtered_collection(set_id)
ids = filtered_ids(set_id)

page_size = 20

_, c, _ = st.columns([6, 6, 2])
with c:
    pages = st.pills(
        "Pages",
        label_visibility="hidden",
        options=range((len(images) // page_size) + 1),
        default=st.session_state.get(page_key, 0),
        selection_mode="single",
        key="collection/pills",
        on_change=on_change_page,
    )
    if pages is None:
        pages = 0
    # st.session_state[page_key] = pages

for i_card in range(page_size * pages, page_size * (pages + 1)):
    _, c, _ = collection_layout.columns([2, 6, 2])
    with c:
        if i_card < len(images):
            st.image(cached_image(images[i_card]), width=200)
            st.toggle(
                f"toggle-{ids[i_card]}",
                label_visibility="hidden",
                value=duckdb.conn.sql(
                    f"""
                        SELECT
                            collected
                        FROM 
                            collection
                        WHERE id = '{ids[i_card]}'
                        """
                )
                .to_df()["collected"][0]
                .item(),
                on_change=duckdb.update_collected,
                kwargs={"card_id": ids[i_card]},
            )
