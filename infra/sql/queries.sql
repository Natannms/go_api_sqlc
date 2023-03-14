-- name: CreateUser :exec
INSERT INTO users (name, email, password) VALUES (?, ?, ?);

-- name: GetUsers :many
SELECT * FROM users;

-- name: GetUserById :one
SELECT * FROM users WHERE id = ?;

-- name: UpdateUser :exec
UPDATE users SET email = ?, password = ? WHERE name = ?;

-- name: DeleteUser :exec
DELETE FROM users WHERE name = ?;

-- name: CreateMetadata :exec
INSERT INTO metadata (user_id, user_type, photo, description, status, address)
VALUES (?, ?, ?, ?, ?, ?);

-- name: GetMetadataById :one
SELECT * FROM metadata WHERE id = ?;

-- name: GetMetadata :many
SELECT * FROM metadata;

-- name: UpdateMetadata :exec
UPDATE metadata SET user_id = ?, user_type = ?, photo = ?, description = ?, status = ?, address = ? WHERE id = ? ;

-- name: DeleteMetadata :exec
DELETE FROM metadata WHERE id = ?;

-- name: CreateInstrument :exec
INSERT INTO instrument (user_id, is_active, instrument_name)
VALUES (?, ?, ?);

-- name: GetInstrumentByUserId :one
SELECT * FROM instrument WHERE user_id = ?;

-- name: GetAllInstruments :many
SELECT * FROM instrument;

-- name: UpdateInstrumentByUserId :exec
UPDATE instrument SET is_active = ?, instrument_name = ? WHERE user_id = ? ;

-- name: DeleteInstrumentByUserId :exec
DELETE FROM instrument WHERE user_id = ?;


-- name: CreateInvite :exec
INSERT INTO invites (user_id, remittee, status, invite_type)
VALUES (?, ?, ?, ?) ;

-- name: GetInviteByUserId :one
SELECT * FROM invites WHERE user_id = ?;

-- name: GetAllInvites :many
SELECT * FROM invites;

-- name: UpdateInviteByUserId :exec
UPDATE invites SET remittee = ?, status = ?, invite_type = ? WHERE user_id = ? ;

-- name: DeleteInviteByUserId :exec
DELETE FROM invites WHERE user_id = ?;


-- name: CreateSolicitation :exec
INSERT INTO solicitation (user_id, remittee, status, invite_type)
VALUES (?, ?, ?, ?);

-- name: GetSolicitationByUserId :one
SELECT * FROM solicitation WHERE user_id = ?;

-- name: GetAllSolicitations :many
SELECT * FROM solicitation;

-- name: UpdateSolicitationByUserId :exec
UPDATE solicitation SET remittee = ?, status = ?, invite_type = ? WHERE user_id = ?;

-- name: DeleteSolicitationByUserId :exec
DELETE FROM solicitation WHERE user_id = ?;


-- name: InsertRepertoire :exec
INSERT INTO repertoire (user_id, name, is_active, music_list_id)
VALUES (?, ?, ?, ?);

-- name: SelectAllRepertoire :many
SELECT * FROM repertoire;

-- name: SelectRepertoireById :one
SELECT * FROM repertoire WHERE user_id = ?;

-- name: UpdateRepertoire :exec
UPDATE repertoire
SET name = ?, is_active = ?, music_list_id = ?
WHERE user_id = ?;

-- name: DeleteRepertoire :exec
DELETE FROM repertoire WHERE user_id = ?;
