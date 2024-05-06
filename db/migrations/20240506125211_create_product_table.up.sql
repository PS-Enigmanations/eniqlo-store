-- Create table products
CREATE TABLE "public"."products" (
    "id" UUID NOT NULL,
    "name" varchar(30) NOT NULL,
    "sku" varchar(30) NOT NULL,
    "category" varchar(20) NOT NULL,
    "image_url" varchar(200) NOT NULL,
    "notes" varchar(200) NOT NULL,
    "price" numeric NOT NULL,
    "stock" integer NOT NULL,
    "location" varchar(200) NOT NULL,
    "is_available" boolean NOT NULL,
    "created_at" timestamptz NOT NULL DEFAULT now(),
    "updated_at" timestamptz NULL,
	"deleted_at" timestamptz NULL,
    CONSTRAINT products_pkey PRIMARY KEY (id)
);