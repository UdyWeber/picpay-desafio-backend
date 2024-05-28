CREATE TABLE "common_user"
(
    "id"         bigserial PRIMARY KEY,
    "full_name"  varchar     NOT NULL,
    "cpf"        varchar     NOT NULL,
    "email"      varchar     NOT NULL,
    "created_at" timestamptz NOT NULL DEFAULT (now()),
    "deleted_at" timestamptz
);

CREATE TABLE "shoopkeepers"
(
    "common_user_id" bigserial NOT NULL,
    "cnpj"           varchar   NOT NULL,
    FOREIGN KEY (common_user_id) REFERENCES common_user (id)
);

