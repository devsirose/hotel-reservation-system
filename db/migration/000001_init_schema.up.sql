CREATE EXTENSION IF NOT EXISTS postgis;

CREATE TABLE "reservation" (
  "reservation_id" uuid PRIMARY KEY,
  "room_id" uuid,
  "user_id" varchar,
  "start_date" TIMESTAMPTZ,
  "end_date" TIMESTAMPTZ,
  "status" varchar,
  "created_at" TIMESTAMPTZ,
  "created_by" uuid,
  "update_at" TIMESTAMPTZ,
  "update_by" uuid
);

CREATE TABLE "hotel" (
  "hotel_id" uuid PRIMARY KEY,
  "destination_id" uuid,
  "type_id" varchar,
  "total_room" integer,
  "rating" float
);

CREATE TABLE "room" (
  "room_id" uuid PRIMARY KEY,
  "room_name" varchar,
  "hotel_id" uuid,
  "floor" integer,
  "type_id" varchar,
  "max_capacity" integer,
  "rate" float,
  "description" varchar,
  "price" integer,
  "created_at" TIMESTAMPTZ,
  "created_by" uuid,
  "update_at" TIMESTAMPTZ,
  "update_by" uuid
);

CREATE TABLE "type" (
  "type_code" varchar PRIMARY KEY,
  "description" varchar,
  "created_at" TIMESTAMPTZ,
  "created_by" uuid,
  "update_at" TIMESTAMPTZ,
  "update_by" uuid
);

CREATE TABLE "media" (
  "media_id" uuid PRIMARY KEY,
  "room_id" uuid,
  "url" varchar,
  "type" varchar,
  "description" varchar,
  "is_primary" boolean
);

CREATE TABLE "room_amenity" (
  "room_id" uuid,
  "amenity_code" varchar,
  PRIMARY KEY ("room_id", "amenity_code")
);

CREATE TABLE "amenity" (
  "amenity_code" varchar PRIMARY KEY,
  "description" varchar
);

CREATE TABLE "user" (
  "username" varchar(30) PRIMARY KEY,
  "role" varchar
);

CREATE TABLE "role" (
  "role_code" varchar PRIMARY KEY,
  "desciption" varchar
);

CREATE TABLE "destination" (
  "destination_id" uuid PRIMARY KEY,
  "address" varchar,
  "country" varchar,
  "type" varchar,
  "location" "GEOGRAPHY(POINT,4326)",
  "boundary" "GEOGRAPHY(POLYGON,4326)"
);

CREATE TABLE "rate" (
  "room_id" uuid,
  "user_id" varchar,
  "score" float,
  "comment" text,
  PRIMARY KEY ("room_id", "user_id")
);

ALTER TABLE "reservation" ADD FOREIGN KEY ("room_id") REFERENCES "room" ("room_id");

ALTER TABLE "reservation" ADD FOREIGN KEY ("user_id") REFERENCES "user" ("username");

ALTER TABLE "hotel" ADD FOREIGN KEY ("destination_id") REFERENCES "destination" ("destination_id");

ALTER TABLE "hotel" ADD FOREIGN KEY ("type_id") REFERENCES "type" ("type_code");

ALTER TABLE "room" ADD FOREIGN KEY ("hotel_id") REFERENCES "hotel" ("hotel_id");

ALTER TABLE "room" ADD FOREIGN KEY ("type_id") REFERENCES "type" ("type_code");

ALTER TABLE "media" ADD FOREIGN KEY ("room_id") REFERENCES "room" ("room_id");

ALTER TABLE "room_amenity" ADD FOREIGN KEY ("room_id") REFERENCES "room" ("room_id");

ALTER TABLE "room_amenity" ADD FOREIGN KEY ("amenity_code") REFERENCES "amenity" ("amenity_code");

ALTER TABLE "user" ADD FOREIGN KEY ("role") REFERENCES "role" ("role_code");

ALTER TABLE "rate" ADD FOREIGN KEY ("room_id") REFERENCES "room" ("room_id");

ALTER TABLE "rate" ADD FOREIGN KEY ("user_id") REFERENCES "user" ("username");
