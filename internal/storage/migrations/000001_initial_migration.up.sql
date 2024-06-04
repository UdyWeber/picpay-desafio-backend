CREATE TABLE "common_user"
(
    "id"         bigserial PRIMARY KEY,
    "full_name"  varchar     NOT NULL,
    "cpf"        varchar     NOT NULL,
    "email"      varchar     NOT NULL,
    "created_at" timestamptz NOT NULL DEFAULT (now()),
    "cnpj"       varchar,
    "deleted_at" timestamptz
);
