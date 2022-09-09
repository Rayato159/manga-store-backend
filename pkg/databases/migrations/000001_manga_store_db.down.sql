BEGIN;

DROP TRIGGER set_updated_at_timestamp_users_table ON "users";
DROP TRIGGER set_updated_at_timestamp_books_table ON "books";
DROP TRIGGER set_updated_at_timestamp_coupons_table ON "coupons";
DROP TRIGGER set_updated_at_timestamp_coupon_usages_table ON "coupon_usages";
DROP TRIGGER set_updated_at_timestamp_categories_table ON "categories";
DROP TRIGGER set_updated_at_timestamp_carts_table ON "carts";
DROP TRIGGER set_updated_at_timestamp_orders_table ON "orders";
DROP TRIGGER set_updated_at_timestamp_order_items_table ON "order_items";
DROP TRIGGER set_updated_at_timestamp_attachments_table ON "attachments";

DROP FUNCTION set_updated_at_column();

DROP TABLE "attachments";
DROP TABLE "order_items";
DROP TABLE "coupon_usages";
DROP TABLE "orders";
DROP TABLE "coupons";
DROP TABLE "shipping_addresses";
DROP TABLE "books_categories";
DROP TABLE "categories";
DROP TABLE "carts";
DROP TABLE "books";
DROP TABLE "users";

DROP TYPE "user_role";
DROP TYPE "coupon_type";
DROP TYPE "order_status";

COMMIT;