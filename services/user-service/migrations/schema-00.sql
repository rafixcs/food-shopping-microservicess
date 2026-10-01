-- user-service: identidade e endereços de clientes, restaurantes e entregadores
-- Postgres. Pensado para ser aplicado com golang-migrate (divida em arquivos
-- numerados .up.sql / .down.sql conforme a convenção do projeto).

CREATE EXTENSION IF NOT EXISTS pgcrypto;

CREATE TABLE users (
    id             UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name           VARCHAR(150) NOT NULL,
    email          VARCHAR(255) NOT NULL UNIQUE,
    password_hash  VARCHAR(255) NOT NULL,
    role           VARCHAR(20)  NOT NULL CHECK (role IN ('customer', 'restaurant', 'courier')),
    phone          VARCHAR(20),
    created_at     TIMESTAMPTZ  NOT NULL DEFAULT now(),
    updated_at     TIMESTAMPTZ  NOT NULL DEFAULT now()
);

-- e-mail já é único (UNIQUE cria índice); role é consultado com frequência
-- pelo gateway ao autorizar rotas restritas (ex: PATCH /orders/{id}/status).
CREATE INDEX idx_users_role ON users (role);

CREATE TABLE addresses (
    id          UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id     UUID NOT NULL REFERENCES users (id) ON DELETE CASCADE,
    label       VARCHAR(50),
    street      VARCHAR(200) NOT NULL,
    number      VARCHAR(20),
    city        VARCHAR(100) NOT NULL,
    state       CHAR(2)      NOT NULL,
    zip         VARCHAR(10)  NOT NULL,
    lat         DOUBLE PRECISION,
    lng         DOUBLE PRECISION,
    is_default  BOOLEAN      NOT NULL DEFAULT false,
    created_at  TIMESTAMPTZ  NOT NULL DEFAULT now()
);

CREATE INDEX idx_addresses_user_id ON addresses (user_id);

-- Garante no máximo um endereço padrão por usuário sem precisar de lógica
-- de aplicação para desmarcar os outros -- índice único parcial.
CREATE UNIQUE INDEX uq_addresses_default_per_user
    ON addresses (user_id)
    WHERE is_default = true;

-- Trigger genérica para manter updated_at sempre atual em UPDATEs.
CREATE OR REPLACE FUNCTION set_updated_at()
RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at := now();
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER trg_users_updated_at
    BEFORE UPDATE ON users
    FOR EACH ROW EXECUTE FUNCTION set_updated_at();