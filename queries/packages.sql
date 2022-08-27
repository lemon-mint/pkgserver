-- name: GetPackageWithID :one
SELECT * FROM packages
WHERE id = $1 LIMIT 1;

-- name: GetPackageWithName :one
SELECT * FROM packages
WHERE pkg_name = $1 LIMIT 1;

-- name: SearchPackages :many
SELECT * FROM packages
WHERE to_tsvector('english', description) @@ to_tsquery('english', $1)
OR pkg_name = $1;
