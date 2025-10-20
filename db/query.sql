-- name: Singup :one
INSERT INTO users(username,email,password)
VALUES ($1,$2,$3)
RETURNING id, email, username,picture;

-- name: Googlelogin :one
INSERT INTO users(username,email,picture,google_id,login_option)
VALUES ($1,$2,$3,$4,'google')
RETURNING id, email, username,picture;

-- name: EditUser :exec
UPDATE users
SET 
  picture = COALESCE(NULLIF($1,''),picture),
  bio = COALESCE(NULLIF($2,''),bio),
  phone_number = COALESCE(NULLIF($3,''),phone_number),
  updated_at = NOW()
WHERE id = $4;


-- name: EditPicture :exec
UPDATE users
SET 
 picture = COALESCE(NULLIF($1,''),picture),
 updated_at = NOW()
WHERE id = $2;


-- name: FindById :one
SELECT id,email,username,picture
FROM users
WHERE id = $1;

-- name: FindByEmail :one
SELECT id,email,username,picture
FROM users
WHERE email = $1;

-- name: FindByUserName :one
SELECT id,email,username,picture
FROM users
WHERE username = $1;

-- name: SeeRevoke :one
SELECT revoked 
FROM users
WHERE id = $1;

-- name: EditRevoke :exec
UPDATE users
SET 
 revoked = $1,
 updated_at = NOW()
WHERE id = $2;

-- name: DeleteUser :exec
DELETE FROM users
WHERE id = $1;


