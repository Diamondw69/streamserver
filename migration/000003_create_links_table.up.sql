CREATE TABLE IF NOT EXISTS links (
                                     id bigserial PRIMARY KEY,
                                     name text NOT NULL,
                                     link text UNIQUE NOT NULL,
                                     status bool default true,
                                     device text NOT NULL
                                     );