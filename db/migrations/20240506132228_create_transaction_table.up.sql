-- Create table transactions
CREATE TABLE "public"."transactions" (
    "id" serial NOT NULL,
    "customer_id" UUID NOT NULL,
    "total" numeric NOT NULL,
    "paid" numeric NOT NULL,
    "change" numeric NOT NULL,
    "created_at" timestamptz NOT NULL DEFAULT now(),
    "updated_at" timestamptz NULL,
    CONSTRAINT "fk_transactions_customer" FOREIGN KEY ("customer_id") REFERENCES "public"."customers"("id"),
    CONSTRAINT transactions_pkey PRIMARY KEY (id)
);