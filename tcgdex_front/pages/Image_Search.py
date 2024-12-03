import io
import streamlit as st
import requests

from PIL import ImageOps, Image

from tcgdex_front.internal.caching import cached_image
import tcgdex_front.internal.repositories.duckdb_repository as duckdb


conn = duckdb.connect()
_, request_col, _ = st.columns([2, 6, 2], vertical_alignment="top")


def format_similarity(similarity: str) -> str:
    return f"{float(similarity):.1%}"


def render_file_uploader():
    if uploaded_image is not None:
        image = ImageOps.exif_transpose(Image.open(uploaded_image))
        if image is not None:
            _, c, _ = st.columns([2, 6, 2])
            with c:
                st.image(image, width=500)


def get_similars():
    if search_button and uploaded_image is not None:
        form_data = {"n_similar": 5}
        image = ImageOps.exif_transpose(Image.open(uploaded_image))
        img_byte_arr = io.BytesIO()
        if image is not None:
            image.save(img_byte_arr, format="JPEG")
        img_byte_arr.seek(0)
        files = {"image": ("payload.jpeg", img_byte_arr, "image/jpeg")}

        res = requests.post(
            "http://0.0.0.0:8001/v1/cards/similar", files=files, data=form_data
        )
        if res.status_code == 200:
            st.session_state["card_number"] = 0
            st.session_state["cards"] = res.json()["cards"]
            st.session_state["search/card_pills"] = 0


def render_similar_card():
    if st.session_state.get("search/card_pills") is not None:
        cards = st.session_state.get("cards", [])
        _card_number = st.session_state.get("card_number", 0)
        if len(cards) > _card_number:
            _, c, _ = st.columns([2, 6, 2])
            with c:
                card_number = st.segmented_control(
                    "Card",
                    options=range(len(cards)),
                    key="search/card_pills",
                    selection_mode="single",
                )
                if card_number is None:
                    card_number = _card_number
                st.session_state["card_number"] = card_number
                st.image(
                    cached_image(cards[card_number]["image"]),
                    width=600,
                )
                st.metric(
                    "Similarity",
                    format_similarity(cards[card_number]["similarity"]),
                )
                st.toggle(
                    "Collected",
                    value=conn.sql(
                        f"""
                            SELECT
                                collected
                            FROM 
                                collection
                            WHERE id = '{cards[card_number]["id"]}'
                            """
                    )
                    .to_df()["collected"][0]
                    .item(),
                    on_change=duckdb.update_collected,
                    kwargs={"card_id": cards[card_number]["id"]},
                )

            with st.popover("Card Info"):
                _, c, _ = st.columns([2, 6, 2])
                cards = st.session_state.get("cards", [])
                card_number = st.session_state.get("card_number", 0)
                with c:
                    st.image(
                        cached_image(
                            cards[card_number]["image"].replace("_hires", ""),
                        ),
                        width=300,
                    )
                    res = requests.get(
                        f"http://0.0.0.0:8001/v1/cards/get/{cards[card_number]['id']}"
                    )
                    if res.status_code == 200:
                        data = res.json()
                        html_code = f"""
                            <div style="display: flex; flex-direction: row; gap: 10px; flex-wrap: wrap;">
                                <span style="background-color: #E63946; color: white; padding: 5px 10px; border-radius: 5px; font-weight: bold;">
                                    Name: {data['name']}
                                </span>
                                <span style="background-color: #4682B4; color: white; padding: 5px 10px; border-radius: 5px; font-weight: bold;">
                                    Pokedex Number: {data['national_pokedex_number']}
                                </span>
                                <span style="background-color: #F4F4F4; color: #333; padding: 5px 10px; border-radius: 5px; font-weight: bold;">
                                    HP: {data['hp']}
                                </span>
                                <span style="background-color: #F4F4F4; color: #333; padding: 5px 10px; border-radius: 5px; font-weight: bold;">
                                    Types: {", ".join(data['types'])}
                                </span>
                                <span style="background-color: #F4F4F4; color: #333; padding: 5px 10px; border-radius: 5px; font-weight: bold;">
                                    Evolves To: {", ".join(data['evolves_to']) if data['evolves_to'] else "None"}
                                </span>
                                <span style="background-color: #F4F4F4; color: #333; padding: 5px 10px; border-radius: 5px; font-weight: bold;">
                                    Weaknesses: {", ".join([f"{w['type']} {w['value']}" for w in data['weaknesses']]) if data['weaknesses'] else "None"}
                                </span>
                                <span style="background-color: #F4F4F4; color: #333; padding: 5px 10px; border-radius: 5px; font-weight: bold;">
                                    Rarity: {data['rarity']}
                                </span>
                                <span style="background-color: #F4F4F4; color: #333; padding: 5px 10px; border-radius: 5px; font-weight: bold;">
                                    Artist: {data['artist']}
                                </span>
                                <span style="background-color: #F4F4F4; color: #333; padding: 5px 10px; border-radius: 5px; font-weight: bold;">
                                    Set: {data['set_id']}
                                </span>
                            </div>
                            """

                        with st.container():
                            st.markdown(html_code, unsafe_allow_html=True)


with request_col:
    search_expander = st.expander("Search", expanded=True)
    with search_expander:
        _, upload_col, _ = st.columns([2, 6, 2], vertical_alignment="center")
        with upload_col:
            option_map = {
                0: ":material/photo_camera:",
                1: ":material/upload_file:",
            }
            upload_method = st.segmented_control(
                "Upload method",
                options=option_map.keys(),
                format_func=lambda option: option_map[option],
                selection_mode="single",
                default=0,
            )
            with st.popover("Upload Image"):
                if upload_method == 0:
                    uploaded_image = st.camera_input(
                        "Upload image for search", key="search/image_uploader"
                    )
                else:
                    uploaded_image = st.file_uploader(
                        "Upload image for search", key="search/image_uploader"
                    )
                search_button = st.button(
                    "Search", key="search/search_button", type="primary", icon="🔍"
                )
        render_file_uploader()

    similar_cards_expander = st.expander("Similar Cards", expanded=True)
    with similar_cards_expander:
        get_similars()
        render_similar_card()
