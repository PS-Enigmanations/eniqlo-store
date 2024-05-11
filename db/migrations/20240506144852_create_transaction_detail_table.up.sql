-- Create table transaction_details
CREATE TABLE "public"."transaction_details" (
    "id" UUID NOT NULL,
    "transaction_id" UUID NOT NULL,
    "product_id" UUID NOT NULL,
    "quantity" int NOT NULL,
    "created_at" timestamptz NOT NULL DEFAULT now(),
    "updated_at" timestamptz NULL,
    CONSTRAINT "fk_transaction_details_transactions" FOREIGN KEY ("transaction_id") REFERENCES "public"."transactions"("id"),
    CONSTRAINT "fk_transaction_details_products" FOREIGN KEY ("product_id") REFERENCES "public"."products"("id"),
    CONSTRAINT transaction_details_pkey PRIMARY KEY (id)
);