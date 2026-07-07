CREATE TABLE texts (
    id BIGSERIAL PRIMARY KEY,
    title TEXT NOT NULL,
    content TEXT NOT NULL,
    created TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    image_ref_id BIGINT
);

CREATE TABLE images_ref (
    id BIGSERIAL PRIMARY KEY,
    title TEXT NOT NULL,
    content TEXT NOT NULL,
    created TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    text_id BIGINT
);

-- Apply foreign keys after all tables exist
ALTER TABLE texts 
    ADD CONSTRAINT fk_image_ref FOREIGN KEY (image_ref_id) REFERENCES images_ref(id) ON DELETE SET NULL;

ALTER TABLE images_ref 
    ADD CONSTRAINT fk_text FOREIGN KEY (text_id) REFERENCES texts(id) ON DELETE CASCADE;
