# new-todo
Go Crud Api Project

create table
-- Table: public.tags

-- DROP TABLE IF EXISTS public.tags;

CREATE TABLE IF NOT EXISTS public.tags
(
    id bigint NOT NULL DEFAULT nextval('tags_id_seq'::regclass),
    item_name character varying(225) COLLATE pg_catalog."default",
    status character varying(225) COLLATE pg_catalog."default",
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    CONSTRAINT tags_pkey PRIMARY KEY (id)
)

TABLESPACE pg_default;

ALTER TABLE IF EXISTS public.tags
    OWNER to postgres;
-- Index: idx_tags_deleted_at

-- DROP INDEX IF EXISTS public.idx_tags_deleted_at;

CREATE INDEX IF NOT EXISTS idx_tags_deleted_at
    ON public.tags USING btree
    (deleted_at ASC NULLS LAST)
    TABLESPACE pg_default;


insert script 
