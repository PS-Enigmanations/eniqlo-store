-- Create table transaction_details
CREATE TABLE "public"."transaction_details" (
    "id" serial NOT NULL,
    "product_id" varchar(50) NOT NULL,
    "quantity" int NOT NULL,
    "total" numeric NOT NULL,
    "created_at" timestamptz NOT NULL DEFAULT now(),
    "updated_at" timestamptz NULL,
    CONSTRAINT "fk_transaction_details_products" FOREIGN KEY ("product_id") REFERENCES "public"."products"("id"),
    CONSTRAINT transaction_details_pkey PRIMARY KEY (id)
);