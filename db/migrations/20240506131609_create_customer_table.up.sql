-- Create table customers
CREATE TABLE "public"."customers" (
    "id" varchar(50) NOT NULL,
    "name" varchar(50) NOT NULL,
    "phone_number" varchar(16) NOT NULL,
    "created_at" timestamptz NOT NULL DEFAULT now(),
    "updated_at" timestamptz NULL,
	"deleted_at" timestamptz NULL,
    CONSTRAINT customers_pkey PRIMARY KEY (id)
);
CREATE UNIQUE INDEX idx_customers_phone_number ON "public"."customers" USING btree (phone_number);