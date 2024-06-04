CREATE TABLE "common_user"
(
    "id"         bigserial PRIMARY KEY,
    "full_name"  varchar     NOT NULL,
    "cpf"        varchar     NOT NULL UNIQUE,
    "email"      varchar     NOT NULL UNIQUE,
    "created_at" timestamptz NOT NULL DEFAULT (now()),
    "cnpj"       varchar,
    "deleted_at" timestamptz
);
