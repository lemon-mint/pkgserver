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

-- name: CreatePackage :exec
INSERT INTO packages (
    pkg_name   ,
    pkg_type   ,
    vcs        ,
    url        ,
    description
) VALUES ($1, $2, $3, $4, $5);

-- name: DeletePackage :exec
DELETE FROM packages
WHERE pkg_name = $1;
