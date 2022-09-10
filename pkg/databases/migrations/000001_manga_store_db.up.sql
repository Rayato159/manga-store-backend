BEGIN;

SET TIME ZONE 'Asia/Bangkok';

CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE OR REPLACE FUNCTION set_updated_at_column()
RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = now();
    RETURN NEW;   
END;
$$ language 'plpgsql';

CREATE TYPE "user_role" AS ENUM (
    'user',
    'admin'
);

CREATE TYPE "coupon_type" AS ENUM (
    'percent',
    'unit'
);

CREATE TYPE "order_status" AS ENUM (
    'waiting',
    'shipping',
    'success',
    'canceled'
);

CREATE TABLE "users" (
  "id" uuid NOT NULL UNIQUE PRIMARY KEY DEFAULT uuid_generate_v4(),
  "password" VARCHAR NOT NULL,
  "username" VARCHAR NOT NULL UNIQUE,
  "role" VARCHAR NOT NULL DEFAULT 'user',
  "refresh_token" VARCHAR NOT NULL DEFAULT '',
  "created_at" TIMESTAMP NOT NULL DEFAULT now(),
  "updated_at" TIMESTAMP NOT NULL DEFAULT now()
);

CREATE TABLE "shipping_addresses" (
  "id" uuid NOT NULL UNIQUE PRIMARY KEY DEFAULT uuid_generate_v4(),
  "user_id" uuid NOT NULL,
  "contract" VARCHAR NOT NULL DEFAULT '',
  "address" VARCHAR NOT NULL DEFAULT ''
);

CREATE TABLE "books" (
  "id" uuid NOT NULL UNIQUE PRIMARY KEY DEFAULT uuid_generate_v4(),
  "title" VARCHAR NOT NULL,
  "volume" INT NOT NULL DEFAULT 0,
  "description" VARCHAR NOT NULL DEFAULT '',
  "price" DOUBLE PRECISION NOT NULL DEFAULT 0,
  "author" VARCHAR NOT NULL,
  "publisher" VARCHAR NOT NULL,
  "qty" INT NOT NULL DEFAULT 0,
  "created_at" TIMESTAMP NOT NULL DEFAULT now(),
  "updated_at" TIMESTAMP NOT NULL DEFAULT now()
);

CREATE TABLE "coupons" (
  "id" uuid NOT NULL UNIQUE PRIMARY KEY DEFAULT uuid_generate_v4(),
  "code" VARCHAR NOT NULL UNIQUE,
  "discount" DOUBLE PRECISION NOT NULL DEFAULT 0,
  "type" coupon_type NOT NULL,
  "qty" INT NOT NULL DEFAULT 0,
  "created_at" TIMESTAMP NOT NULL DEFAULT now(),
  "updated_at" TIMESTAMP NOT NULL DEFAULT now()
);

CREATE TABLE "coupon_usages" (
  "id" INT NOT NULL PRIMARY KEY,
  "coupon_id" uuid NOT NULL,
  "user_id" uuid NOT NULL,
  "created_at" TIMESTAMP NOT NULL DEFAULT now(),
  "updated_at" TIMESTAMP NOT NULL DEFAULT now()
);

CREATE TABLE "categories" (
  "id" INT NOT NULL PRIMARY KEY,
  "name" VARCHAR NOT NULL,
  "description" VARCHAR NOT NULL DEFAULT '',
  "created_at" TIMESTAMP NOT NULL DEFAULT now(),
  "updated_at" TIMESTAMP NOT NULL DEFAULT now()
);

CREATE TABLE "books_categories" (
  "id" INT NOT NULL PRIMARY KEY,
  "book_id" uuid NOT NULL,
  "category_id" INT NOT NULL
);

CREATE TABLE "carts" (
  "id" INT NOT NULL PRIMARY KEY,
  "book_id" uuid NOT NULL,
  "user_id" uuid NOT NULL,
  "qty" INT NOT NULL DEFAULT 0,
  "created_at" TIMESTAMP NOT NULL DEFAULT now(),
  "updated_at" TIMESTAMP NOT NULL DEFAULT now()
);

CREATE TABLE "orders" (
  "id" uuid NOT NULL UNIQUE PRIMARY KEY DEFAULT uuid_generate_v4(),
  "user_id" uuid NOT NULL,
  "status" order_status NOT NULL DEFAULT 'waiting',
  "contract" VARCHAR NOT NULL,
  "shipping_address" VARCHAR NOT NULL,
  "sub_total" DOUBLE PRECISION NOT NULL,
  "discount" DOUBLE PRECISION NOT NULL DEFAULT 0,
  "grand_total" DOUBLE PRECISION NOT NULL,
  "coupon_id" uuid,
  "description" VARCHAR NOT NULL DEFAULT '',
  "created_at" TIMESTAMP NOT NULL DEFAULT now(),
  "updated_at" TIMESTAMP NOT NULL DEFAULT now()
);

CREATE TABLE "order_items" (
  "id" uuid NOT NULL UNIQUE PRIMARY KEY DEFAULT uuid_generate_v4(),
  "order_id" uuid NOT NULL,
  "book_id" uuid NOT NULL,
  "book_title" VARCHAR NOT NULL,
  "book_author" VARCHAR NOT NULL,
  "book_publisher" VARCHAR NOT NULL,
  "qty" INT NOT NULL DEFAULT 1,
  "grand_total" DOUBLE PRECISION NOT NULL,
  "sub_total" DOUBLE PRECISION NOT NULL DEFAULT 0,
  "discount" DOUBLE PRECISION NOT NULL,
  "created_at" TIMESTAMP NOT NULL DEFAULT now(),
  "updated_at" TIMESTAMP NOT NULL DEFAULT now()
);

CREATE TABLE "attachments" (
  "id" uuid NOT NULL UNIQUE PRIMARY KEY DEFAULT uuid_generate_v4(),
  "book_id" uuid,
  "order_id" uuid,
  "url" VARCHAR NOT NULL DEFAULT '',
  "name" VARCHAR NOT NULL DEFAULT '',
  "extension" VARCHAR NOT NULL DEFAULT '',
  "purpose" VARCHAR NOT NULL DEFAULT '',
  "created_at" TIMESTAMP NOT NULL DEFAULT now(),
  "updated_at" TIMESTAMP NOT NULL DEFAULT now()
);

ALTER TABLE "shipping_addresses" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id") ON DELETE CASCADE;
ALTER TABLE "carts" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id") ON DELETE CASCADE;
ALTER TABLE "carts" ADD FOREIGN KEY ("book_id") REFERENCES "books" ("id") ON DELETE CASCADE;
ALTER TABLE "books_categories" ADD FOREIGN KEY ("book_id") REFERENCES "books" ("id") ON DELETE CASCADE;
ALTER TABLE "books_categories" ADD FOREIGN KEY ("category_id") REFERENCES "categories" ("id") ON DELETE CASCADE;
ALTER TABLE "attachments" ADD FOREIGN KEY ("book_id") REFERENCES "books" ("id") ON DELETE CASCADE;
ALTER TABLE "orders" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");
ALTER TABLE "order_items" ADD FOREIGN KEY ("order_id") REFERENCES "orders" ("id") ON DELETE CASCADE;
ALTER TABLE "coupon_usages" ADD FOREIGN KEY ("coupon_id") REFERENCES "coupons" ("id");
ALTER TABLE "coupon_usages" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");
ALTER TABLE "orders" ADD FOREIGN KEY ("coupon_id") REFERENCES "coupons" ("id");

CREATE TRIGGER set_updated_at_timestamp_users_table BEFORE UPDATE ON "users" FOR EACH ROW EXECUTE PROCEDURE set_updated_at_column();
CREATE TRIGGER set_updated_at_timestamp_books_table BEFORE UPDATE ON "books" FOR EACH ROW EXECUTE PROCEDURE set_updated_at_column();
CREATE TRIGGER set_updated_at_timestamp_coupons_table BEFORE UPDATE ON "coupons" FOR EACH ROW EXECUTE PROCEDURE set_updated_at_column();
CREATE TRIGGER set_updated_at_timestamp_coupon_usages_table BEFORE UPDATE ON "coupon_usages" FOR EACH ROW EXECUTE PROCEDURE set_updated_at_column();
CREATE TRIGGER set_updated_at_timestamp_categories_table BEFORE UPDATE ON "categories" FOR EACH ROW EXECUTE PROCEDURE set_updated_at_column();
CREATE TRIGGER set_updated_at_timestamp_carts_table BEFORE UPDATE ON "carts" FOR EACH ROW EXECUTE PROCEDURE set_updated_at_column();
CREATE TRIGGER set_updated_at_timestamp_orders_table BEFORE UPDATE ON "orders" FOR EACH ROW EXECUTE PROCEDURE set_updated_at_column();
CREATE TRIGGER set_updated_at_timestamp_order_items_table BEFORE UPDATE ON "order_items" FOR EACH ROW EXECUTE PROCEDURE set_updated_at_column();
CREATE TRIGGER set_updated_at_timestamp_attachments_table BEFORE UPDATE ON "attachments" FOR EACH ROW EXECUTE PROCEDURE set_updated_at_column();

COMMIT;