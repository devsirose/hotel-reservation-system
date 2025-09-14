-- name: CreateReservation :one
INSERT INTO reservation (
  reservation_id,
  room_id,
  user_id,
  start_date,
  end_date,
  status,
  created_at,
  created_by,
  update_at,
  update_by
) VALUES (
  $1, $2, $3, $4, $5, $6, $7, $8, $9, $10
) RETURNING *;

-- name: GetReservation :one
SELECT * FROM reservation
WHERE reservation_id = $1 LIMIT 1;

-- name: ListReservations :many
SELECT * FROM reservation
ORDER BY created_at DESC
LIMIT $1
OFFSET $2;

-- name: ListReservationsByUser :many
SELECT * FROM reservation
WHERE user_id = $1
ORDER BY start_date DESC
LIMIT $2
OFFSET $3;

-- name: ListReservationsByRoom :many
SELECT * FROM reservation
WHERE room_id = $1
ORDER BY start_date
LIMIT $2
OFFSET $3;

-- name: GetReservationsByDateRange :many
SELECT * FROM reservation
WHERE room_id = $1
  AND status = $2
  AND (
    (start_date <= $3 AND end_date > $3) OR
    (start_date < $4 AND end_date >= $4) OR
    (start_date >= $3 AND end_date <= $4)
  )
ORDER BY start_date;

-- name: UpdateReservationStatus :one
UPDATE reservation
SET 
  status = $2,
  update_at = $3,
  update_by = $4
WHERE reservation_id = $1
RETURNING *;

-- name: UpdateReservation :one
UPDATE reservation
SET 
  room_id = $2,
  user_id = $3,
  start_date = $4,
  end_date = $5,
  status = $6,
  update_at = $7,
  update_by = $8
WHERE reservation_id = $1
RETURNING *;

-- name: DeleteReservation :exec
DELETE FROM reservation
WHERE reservation_id = $1;