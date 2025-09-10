-- Drop constraints first (to avoid dependency errors)
ALTER TABLE IF EXISTS "reservation" DROP CONSTRAINT IF EXISTS reservation_room_id_fkey;
ALTER TABLE IF EXISTS "reservation" DROP CONSTRAINT IF EXISTS reservation_user_id_fkey;
ALTER TABLE IF EXISTS "hotel" DROP CONSTRAINT IF EXISTS hotel_destination_id_fkey;
ALTER TABLE IF EXISTS "hotel" DROP CONSTRAINT IF EXISTS hotel_type_id_fkey;
ALTER TABLE IF EXISTS "room" DROP CONSTRAINT IF EXISTS room_hotel_id_fkey;
ALTER TABLE IF EXISTS "room" DROP CONSTRAINT IF EXISTS room_type_id_fkey;
ALTER TABLE IF EXISTS "media" DROP CONSTRAINT IF EXISTS media_room_id_fkey;
ALTER TABLE IF EXISTS "room_amenity" DROP CONSTRAINT IF EXISTS room_amenity_room_id_fkey;
ALTER TABLE IF EXISTS "room_amenity" DROP CONSTRAINT IF EXISTS room_amenity_amenity_code_fkey;
ALTER TABLE IF EXISTS "user" DROP CONSTRAINT IF EXISTS user_role_fkey;
ALTER TABLE IF EXISTS "rate" DROP CONSTRAINT IF EXISTS rate_room_id_fkey;
ALTER TABLE IF EXISTS "rate" DROP CONSTRAINT IF EXISTS rate_user_id_fkey;

-- Drop tables in reverse dependency order
DROP TABLE IF EXISTS "rate";
DROP TABLE IF EXISTS "reservation";
DROP TABLE IF EXISTS "media";
DROP TABLE IF EXISTS "room_amenity";
DROP TABLE IF EXISTS "amenity";
DROP TABLE IF EXISTS "room";
DROP TABLE IF EXISTS "hotel";
DROP TABLE IF EXISTS "destination";
DROP TABLE IF EXISTS "type";
DROP TABLE IF EXISTS "user";
DROP TABLE IF EXISTS "role";
