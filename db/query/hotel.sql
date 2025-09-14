-- name: CreateHotel :one
INSERT INTO hotel (
  hotel_id,
  destination_id,
  type_id,
  total_room,
  rating
) VALUES (
  $1, $2, $3, $4, $5
) RETURNING *;

-- name: GetHotel :one
SELECT * FROM hotel
WHERE hotel_id = $1 LIMIT 1;

-- name: ListHotels :many
SELECT * FROM hotel
ORDER BY hotel_id
LIMIT $1
OFFSET $2;

-- name: ListHotelsByDestination :many
SELECT * FROM hotel
WHERE destination_id = $1
ORDER BY rating DESC
LIMIT $2
OFFSET $3;

-- name: UpdateHotel :one
UPDATE hotel
SET 
  destination_id = $2,
  type_id = $3,
  total_room = $4,
  rating = $5
WHERE hotel_id = $1
RETURNING *;

-- name: DeleteHotel :exec
DELETE FROM hotel
WHERE hotel_id = $1;