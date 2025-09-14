-- name: CreateRoom :one
INSERT INTO room (
  room_id,
  room_name,
  hotel_id,
  floor,
  type_id,
  max_capacity,
  rate,
  description,
  price,
  created_at,
  created_by,
  update_at,
  update_by
) VALUES (
  $1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13
) RETURNING *;

-- name: GetRoom :one
SELECT * FROM room
WHERE room_id = $1 LIMIT 1;

-- name: ListRooms :many
SELECT * FROM room
ORDER BY room_id
LIMIT $1
OFFSET $2;

-- name: ListRoomsByHotel :many
SELECT * FROM room
WHERE hotel_id = $1
ORDER BY floor, room_name
LIMIT $2
OFFSET $3;

-- name: GetAvailableRooms :many
SELECT r.* FROM room r
WHERE r.hotel_id = $1
  AND r.room_id NOT IN (
    SELECT res.room_id FROM reservation res
    WHERE res.room_id = r.room_id
      AND res.status = 'confirmed'
      AND (
        (res.start_date <= $2 AND res.end_date > $2) OR
        (res.start_date < $3 AND res.end_date >= $3) OR
        (res.start_date >= $2 AND res.end_date <= $3)
      )
  )
ORDER BY r.price
LIMIT $4
OFFSET $5;

-- name: UpdateRoom :one
UPDATE room
SET 
  room_name = $2,
  hotel_id = $3,
  floor = $4,
  type_id = $5,
  max_capacity = $6,
  rate = $7,
  description = $8,
  price = $9,
  update_at = $10,
  update_by = $11
WHERE room_id = $1
RETURNING *;

-- name: DeleteRoom :exec
DELETE FROM room
WHERE room_id = $1;