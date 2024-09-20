-- +goose Up
-- +goose StatementBegin
CREATE TABLE policies (
    id uuid primary key,
    endpoint text not null unique,
    allowed_roles role[]
);

INSERT INTO policies(id, endpoint, allowed_roles) VALUES
    (gen_random_uuid(), '/access_v1.AccessV1/AddRoleEndpoint', ARRAY ['ADMIN']::role[]),
    (gen_random_uuid(), '/access_v1.AccessV1/UpdateRoleEndpoint', ARRAY ['ADMIN']::role[]),
    (gen_random_uuid(), '/access_v1.AccessV1/DeleteRoleEndpoint', ARRAY ['ADMIN']::role[]),
    (gen_random_uuid(), '/access_v1.AccessV1/GetRoleEndpoints', ARRAY ['ADMIN']::role[]);

INSERT INTO policies(id, endpoint, allowed_roles) VALUES
    (gen_random_uuid(), '/chat_v1.ChatV1/Create', ARRAY ['ADMIN']::role[]),
    (gen_random_uuid(), '/chat_v1.ChatV1/Delete', ARRAY ['ADMIN']::role[]),
    (gen_random_uuid(), '/chat_v1.ChatV1/SendMessage', ARRAY ['ADMIN', 'USER']::role[]),
    (gen_random_uuid(), '/chat_v1.ChatV1/Connect', ARRAY ['ADMIN', 'USER']::role[]);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS policies;
-- +goose StatementEnd