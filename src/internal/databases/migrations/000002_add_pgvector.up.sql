CREATE EXTENSION vector;

ALTER TABLE cards ADD image_embedding VECTOR(1280);

CREATE INDEX ON cards USING hnsw (image_embedding vector_cosine_ops);
