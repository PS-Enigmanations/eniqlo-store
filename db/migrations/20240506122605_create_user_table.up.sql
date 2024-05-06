-- Create table users
CREATE TABLE "public"."users" (
    "id" UUID NOT NULL,
    "name" varchar(50) NOT NULL,
    "phone_number" varchar(16) NOT NULL,
    "password" varchar(150) NOT NULL,
    "created_at" timestamptz NOT NULL DEFAULT now(),
    "updated_at" timestamptz NULL,
	"deleted_at" timestamptz NULL,
    CONSTRAINT users_pkey PRIMARY KEY (id)
);
CREATE UNIQUE INDEX idx_credentials_phone_number ON "public"."users" USING btree (phone_number);