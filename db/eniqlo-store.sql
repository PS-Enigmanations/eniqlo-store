--
-- PostgreSQL database dump
--

-- Dumped from database version 13.8
-- Dumped by pg_dump version 13.8

-- Started on 2024-05-10 10:04:09 WITA

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET client_encoding = 'EUC_KR';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET xmloption = content;
SET client_min_messages = warning;
SET row_security = off;

--
-- TOC entry 3216 (class 1262 OID 3580419)
-- Name: eniqlo-store; Type: DATABASE; Schema: -; Owner: postgres
--

CREATE DATABASE "eniqlo-store" WITH TEMPLATE = template0 ENCODING = 'UTF8' LOCALE = 'C';


ALTER DATABASE "eniqlo-store" OWNER TO postgres;

\connect -reuse-previous=on "dbname='eniqlo-store'"

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET client_encoding = 'EUC_KR';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET xmloption = content;
SET client_min_messages = warning;
SET row_security = off;

--
-- TOC entry 3 (class 2615 OID 2200)
-- Name: public; Type: SCHEMA; Schema: -; Owner: postgres
--

CREATE SCHEMA public;


ALTER SCHEMA public OWNER TO postgres;

--
-- TOC entry 3217 (class 0 OID 0)
-- Dependencies: 3
-- Name: SCHEMA public; Type: COMMENT; Schema: -; Owner: postgres
--

COMMENT ON SCHEMA public IS 'standard public schema';


SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- TOC entry 203 (class 1259 OID 3580471)
-- Name: customers; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.customers (
    id uuid NOT NULL,
    name character varying(50) NOT NULL,
    phone_number character varying(16) NOT NULL,
    created_at timestamp with time zone DEFAULT now() NOT NULL,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone
);


ALTER TABLE public.customers OWNER TO postgres;

--
-- TOC entry 202 (class 1259 OID 3580456)
-- Name: products; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.products (
    id uuid NOT NULL,
    name character varying(30) NOT NULL,
    sku character varying(30) NOT NULL,
    category character varying(20) NOT NULL,
    image_url character varying(200) NOT NULL,
    notes character varying(200) NOT NULL,
    price numeric NOT NULL,
    stock integer NOT NULL,
    location character varying(200) NOT NULL,
    is_available boolean NOT NULL,
    _search tsvector,
    created_at timestamp with time zone DEFAULT now() NOT NULL,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone
);


ALTER TABLE public.products OWNER TO postgres;

--
-- TOC entry 200 (class 1259 OID 3580438)
-- Name: schema_migrations; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.schema_migrations (
    version bigint NOT NULL,
    dirty boolean NOT NULL
);


ALTER TABLE public.schema_migrations OWNER TO postgres;

--
-- TOC entry 205 (class 1259 OID 3580500)
-- Name: transaction_details; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.transaction_details (
    id uuid NOT NULL,
    transaction_id uuid NOT NULL,
    product_id uuid NOT NULL,
    quantity integer NOT NULL,
    total numeric NOT NULL,
    created_at timestamp with time zone DEFAULT now() NOT NULL,
    updated_at timestamp with time zone
);


ALTER TABLE public.transaction_details OWNER TO postgres;

--
-- TOC entry 204 (class 1259 OID 3580482)
-- Name: transactions; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.transactions (
    id uuid NOT NULL,
    customer_id uuid NOT NULL,
    total numeric NOT NULL,
    paid numeric NOT NULL,
    change numeric NOT NULL,
    created_at timestamp with time zone DEFAULT now() NOT NULL,
    updated_at timestamp with time zone
);


ALTER TABLE public.transactions OWNER TO postgres;

--
-- TOC entry 201 (class 1259 OID 3580445)
-- Name: users; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.users (
    id uuid NOT NULL,
    name character varying(50) NOT NULL,
    phone_number character varying(16) NOT NULL,
    password character varying(150) NOT NULL,
    created_at timestamp with time zone DEFAULT now() NOT NULL,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone
);


ALTER TABLE public.users OWNER TO postgres;

--
-- TOC entry 3208 (class 0 OID 3580471)
-- Dependencies: 203
-- Data for Name: customers; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO public.customers VALUES ('c5c0e5c7-62c5-4e6e-8f6d-09cc12345678', 'John Doe', '1234567890', '2024-05-09 14:07:09.516022+08', NULL, NULL);
INSERT INTO public.customers VALUES ('2c1b08f9-3c43-42d2-b0f2-b55e98765432', 'Jane Smith', '9876543210', '2024-05-09 14:07:09.516022+08', NULL, NULL);
INSERT INTO public.customers VALUES ('6f704b33-eb4b-49b1-af32-0ae655432167', 'Alice Johnson', '5551234567', '2024-05-09 14:07:09.516022+08', NULL, NULL);


--
-- TOC entry 3207 (class 0 OID 3580456)
-- Dependencies: 202
-- Data for Name: products; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO public.products VALUES ('0f8fad5b-d9cb-469f-a165-70867728950e', 'Blue T-Shirt', 'TS001', 'Clothing', 'https://example.com/image1.jpg', 'Comfortable cotton material', 15.99, 50, 'Warehouse A', true, '''blue'':1 ''shirt'':4 ''t-shirt'':2', '2024-05-09 14:07:09.516022+08', NULL, NULL);
INSERT INTO public.products VALUES ('7c9e6679-7425-40de-944b-e07fc1f90ae7', 'Black Leather Belt', 'BLB002', 'Accessories', 'https://example.com/image2.jpg', 'Genuine leather, adjustable buckle', 29.99, 30, 'Warehouse B', true, '''belt'':3 ''black'':1 ''leather'':2', '2024-05-09 14:07:09.516022+08', NULL, NULL);
INSERT INTO public.products VALUES ('8f14e45f-ceea-167a-5a36-dedd4bea2543', 'Running Shoes', 'SH003', 'Footwear', 'https://example.com/image3.jpg', 'Breathable mesh, cushioned sole', 49.99, 100, 'Warehouse C', true, '''run'':1 ''shoe'':2', '2023-05-09 14:07:09.516+08', NULL, NULL);
INSERT INTO public.products VALUES ('c9f0f895-fb98-ab91-59f5-1fd0297e236d', 'Green Tea', 'GT004', 'Beverages', 'https://example.com/image4.jpg', 'Organic green tea leaves', 7.99, 0, 'Warehouse D', false, '''green'':1 ''tea'':2', '2024-06-09 14:07:09.516+08', NULL, NULL);


--
-- TOC entry 3205 (class 0 OID 3580438)
-- Dependencies: 200
-- Data for Name: schema_migrations; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO public.schema_migrations VALUES (20240506152354, false);


--
-- TOC entry 3210 (class 0 OID 3580500)
-- Dependencies: 205
-- Data for Name: transaction_details; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO public.transaction_details VALUES ('7b52009b-8b3c-4dc4-8f11-345bea012345', '33e19244-91a0-4e7b-90c8-41bf8e0514e1', '0f8fad5b-d9cb-469f-a165-70867728950e', 2, 31.98, '2024-05-06 12:00:00+08', NULL);
INSERT INTO public.transaction_details VALUES ('60e2dece-087b-40ab-a4c9-7865dc987654', 'de9fb3b0-7b98-4c90-bd6b-35b8f75432fc', '7c9e6679-7425-40de-944b-e07fc1f90ae7', 1, 29.99, '2024-05-06 12:00:00+08', NULL);
INSERT INTO public.transaction_details VALUES ('3f7b709b-41b2-4c5c-ae5f-1f669d012345', 'd4d7b6f2-ae92-4ee4-bfa7-f789f0d12345', '8f14e45f-ceea-167a-5a36-dedd4bea2543', 3, 149.97, '2024-05-06 12:00:00+08', NULL);


--
-- TOC entry 3209 (class 0 OID 3580482)
-- Dependencies: 204
-- Data for Name: transactions; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO public.transactions VALUES ('33e19244-91a0-4e7b-90c8-41bf8e0514e1', 'c5c0e5c7-62c5-4e6e-8f6d-09cc12345678', 150.99, 160.00, 9.01, '2024-05-06 12:00:00+08', NULL);
INSERT INTO public.transactions VALUES ('de9fb3b0-7b98-4c90-bd6b-35b8f75432fc', '2c1b08f9-3c43-42d2-b0f2-b55e98765432', 75.50, 80.00, 4.50, '2024-05-06 12:00:00+08', NULL);
INSERT INTO public.transactions VALUES ('d4d7b6f2-ae92-4ee4-bfa7-f789f0d12345', '6f704b33-eb4b-49b1-af32-0ae655432167', 200.00, 200.00, 0, '2024-05-06 12:00:00+08', NULL);


--
-- TOC entry 3206 (class 0 OID 3580445)
-- Dependencies: 201
-- Data for Name: users; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO public.users VALUES ('7a5da8e5-8f41-4e22-bdb9-d3c63f0b2f6e', 'John Doe', '1234567890', '$2a$08$byUr0FmFYtz8zVp7RzsU8.ASjdwAKGAwL6n.nPU6J4g6VNpDx/utu', '2024-05-09 14:07:09.516022+08', NULL, NULL);
INSERT INTO public.users VALUES ('b19ab0d0-0c49-47ff-8575-4a34a72b0e17', 'Jane Smith', '9876543210', '$2a$08$byUr0FmFYtz8zVp7RzsU8.ASjdwAKGAwL6n.nPU6J4g6VNpDx/utu', '2024-05-09 14:07:09.516022+08', NULL, NULL);
INSERT INTO public.users VALUES ('f2bb0d18-8ef3-4d7a-a2fc-0744f13e32b7', 'Alice Johnson', '5551234567', '$2a$08$byUr0FmFYtz8zVp7RzsU8.ASjdwAKGAwL6n.nPU6J4g6VNpDx/utu', '2024-05-09 14:07:09.516022+08', NULL, NULL);


--
-- TOC entry 3065 (class 2606 OID 3580476)
-- Name: customers customers_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.customers
    ADD CONSTRAINT customers_pkey PRIMARY KEY (id);


--
-- TOC entry 3062 (class 2606 OID 3580464)
-- Name: products products_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.products
    ADD CONSTRAINT products_pkey PRIMARY KEY (id);


--
-- TOC entry 3057 (class 2606 OID 3580442)
-- Name: schema_migrations schema_migrations_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.schema_migrations
    ADD CONSTRAINT schema_migrations_pkey PRIMARY KEY (version);


--
-- TOC entry 3070 (class 2606 OID 3580508)
-- Name: transaction_details transaction_details_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.transaction_details
    ADD CONSTRAINT transaction_details_pkey PRIMARY KEY (id);


--
-- TOC entry 3068 (class 2606 OID 3580490)
-- Name: transactions transactions_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.transactions
    ADD CONSTRAINT transactions_pkey PRIMARY KEY (id);


--
-- TOC entry 3060 (class 2606 OID 3580450)
-- Name: users users_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_pkey PRIMARY KEY (id);


--
-- TOC entry 3058 (class 1259 OID 3580451)
-- Name: idx_credentials_phone_number; Type: INDEX; Schema: public; Owner: postgres
--

CREATE UNIQUE INDEX idx_credentials_phone_number ON public.users USING btree (phone_number);


--
-- TOC entry 3066 (class 1259 OID 3580477)
-- Name: idx_customers_phone_number; Type: INDEX; Schema: public; Owner: postgres
--

CREATE UNIQUE INDEX idx_customers_phone_number ON public.customers USING btree (phone_number);


--
-- TOC entry 3063 (class 1259 OID 3580465)
-- Name: products_search; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX products_search ON public.products USING gin (_search);


--
-- TOC entry 3074 (class 2620 OID 3580466)
-- Name: products products_vector_update; Type: TRIGGER; Schema: public; Owner: postgres
--

CREATE TRIGGER products_vector_update BEFORE INSERT OR UPDATE ON public.products FOR EACH ROW EXECUTE FUNCTION tsvector_update_trigger('_search', 'pg_catalog.english', 'name');


--
-- TOC entry 3073 (class 2606 OID 3580514)
-- Name: transaction_details fk_transaction_details_products; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.transaction_details
    ADD CONSTRAINT fk_transaction_details_products FOREIGN KEY (product_id) REFERENCES public.products(id);


--
-- TOC entry 3072 (class 2606 OID 3580509)
-- Name: transaction_details fk_transaction_details_transactions; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.transaction_details
    ADD CONSTRAINT fk_transaction_details_transactions FOREIGN KEY (transaction_id) REFERENCES public.transactions(id);


--
-- TOC entry 3071 (class 2606 OID 3580491)
-- Name: transactions fk_transactions_customer; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.transactions
    ADD CONSTRAINT fk_transactions_customer FOREIGN KEY (customer_id) REFERENCES public.customers(id);


-- Completed on 2024-05-10 10:04:10 WITA

--
-- PostgreSQL database dump complete
--

