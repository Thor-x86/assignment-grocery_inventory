-- ** Database generated with pgModeler (PostgreSQL Database Modeler).
-- ** pgModeler version: 2.0.0-beta1
-- ** PostgreSQL version: 18.0
-- ** Project Site: pgmodeler.io
-- ** Model Author: Athaariq Ardhiansyah <foss@athaariq.my.id>
-- object: app | type: ROLE --
-- DROP ROLE IF EXISTS app;
CREATE ROLE app WITH 
	LOGIN
	 PASSWORD 'hWPyWuUgmvfTzPUJ';
-- ddl-end --


SET search_path TO pg_catalog,public;
-- ddl-end --

-- object: public.type_status | type: TYPE --
-- DROP TYPE IF EXISTS public.type_status CASCADE;
CREATE TYPE public.type_status AS
ENUM ('unknown','active','backordered','discontinued');
-- ddl-end --
ALTER TYPE public.type_status OWNER TO app;
-- ddl-end --

-- object: public.product | type: TABLE --
-- DROP TABLE IF EXISTS public.product CASCADE;
CREATE TABLE public.product (
	id bigint NOT NULL,
	supplier_id bigint NOT NULL,
	warehouse_id bigint NOT NULL,
	received_at timestamp NOT NULL,
	last_order_at timestamp NOT NULL,
	expired_at timestamp NOT NULL,
	name varchar(32) NOT NULL,
	status public.type_status NOT NULL,
	unit_price_usd float4 NOT NULL,
	percentage float4 NOT NULL,
	stock_qty int2 NOT NULL,
	reorder_level int2 NOT NULL,
	reorder_qty int2 NOT NULL,
	sales_volume int2 NOT NULL,
	inventory_turnover_ratio int2 NOT NULL,
	CONSTRAINT product_pk PRIMARY KEY (id)
);
-- ddl-end --
ALTER TABLE public.product OWNER TO app;
-- ddl-end --

-- object: public.supplier | type: TABLE --
-- DROP TABLE IF EXISTS public.supplier CASCADE;
CREATE TABLE public.supplier (
	id bigint NOT NULL,
	name varchar(32) NOT NULL,
	CONSTRAINT supplier_pk PRIMARY KEY (id)
);
-- ddl-end --
ALTER TABLE public.supplier OWNER TO app;
-- ddl-end --

-- object: public.warehouse | type: TABLE --
-- DROP TABLE IF EXISTS public.warehouse CASCADE;
CREATE TABLE public.warehouse (
	id bigint NOT NULL,
	name varchar(32) NOT NULL,
	code int4 NOT NULL,
	CONSTRAINT warehouse_pk PRIMARY KEY (id),
	CONSTRAINT uq_warehouse_name UNIQUE (name)
);
-- ddl-end --
ALTER TABLE public.warehouse OWNER TO app;
-- ddl-end --

-- object: idx_product_reorder_best | type: INDEX --
-- DROP INDEX IF EXISTS public.idx_product_reorder_best CASCADE;
CREATE INDEX idx_product_reorder_best ON public.product
USING btree
(
	reorder_level DESC NULLS LAST
)
WHERE (status = 'active');
-- ddl-end --

-- object: idx_product_reorder_worst | type: INDEX --
-- DROP INDEX IF EXISTS public.idx_product_reorder_worst CASCADE;
CREATE INDEX idx_product_reorder_worst ON public.product
USING btree
(
	reorder_level ASC NULLS LAST
)
WHERE (status = 'active');
-- ddl-end --

-- object: idx_product_reorder_best_by_warehouse | type: INDEX --
-- DROP INDEX IF EXISTS public.idx_product_reorder_best_by_warehouse CASCADE;
CREATE INDEX idx_product_reorder_best_by_warehouse ON public.product
USING btree
(
	warehouse_id ASC NULLS LAST,
	reorder_level DESC NULLS LAST
)
WHERE (status = 'active');
-- ddl-end --

-- object: idx_product_reorder_worst_by_warehouse | type: INDEX --
-- DROP INDEX IF EXISTS public.idx_product_reorder_worst_by_warehouse CASCADE;
CREATE INDEX idx_product_reorder_worst_by_warehouse ON public.product
USING btree
(
	warehouse_id ASC NULLS LAST,
	reorder_level ASC NULLS LAST
)
WHERE (status = 'active');
-- ddl-end --

-- object: fk_product_supplier_id | type: CONSTRAINT --
-- ALTER TABLE public.product DROP CONSTRAINT IF EXISTS fk_product_supplier_id CASCADE;
ALTER TABLE public.product ADD CONSTRAINT fk_product_supplier_id FOREIGN KEY (supplier_id)
REFERENCES public.supplier (id) MATCH SIMPLE
ON DELETE RESTRICT ON UPDATE RESTRICT;
-- ddl-end --

-- object: fk_product_warehouse_id | type: CONSTRAINT --
-- ALTER TABLE public.product DROP CONSTRAINT IF EXISTS fk_product_warehouse_id CASCADE;
ALTER TABLE public.product ADD CONSTRAINT fk_product_warehouse_id FOREIGN KEY (warehouse_id)
REFERENCES public.warehouse (id) MATCH SIMPLE
ON DELETE RESTRICT ON UPDATE RESTRICT;
-- ddl-end --


