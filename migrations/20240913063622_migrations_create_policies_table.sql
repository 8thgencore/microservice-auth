-- +goose Up
-- +goose StatementBegin
CREATE TABLE policies (
    id serial primary key,
    endpoint text not null unique,
    allowed_roles role[]
);

INSERT INTO policies(endpoint, allowed_roles) VALUES
    ('/access_v1.AccessV1/AddRoleEndpoint', ARRAY ['ADMIN']::role[]),
    ('/access_v1.AccessV1/UpdateRoleEndpoint', ARRAY ['ADMIN']::role[]),
    ('/access_v1.AccessV1/DeleteRoleEndpoint', ARRAY ['ADMIN']::role[]),
    ('/access_v1.AccessV1/GetRoleEndpoints', ARRAY ['ADMIN']::role[]);

INSERT INTO policies(endpoint, allowed_roles) VALUES
    ('/chat_v1.ChatV1/Create', ARRAY ['ADMIN']::role[]),
    ('/chat_v1.ChatV1/Delete', ARRAY ['ADMIN']::role[]),
    ('/chat_v1.ChatV1/SendMessage', ARRAY ['ADMIN', 'USER']::role[]),
    ('/chat_v1.ChatV1/Connect', ARRAY ['ADMIN', 'USER']::role[]);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS policies;
-- +goose StatementEnd