--
-- PostgreSQL database dump
--

\restrict q3AsFTUwOQOdqocyacH1u1VMKLR7W3pofren69MmTbCNMaRoRXxMQvR3REYg93A

-- Dumped from database version 14.5 (Debian 14.5-2.pgdg110+2)
-- Dumped by pg_dump version 14.19 (Homebrew)

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET xmloption = content;
SET client_min_messages = warning;
SET row_security = off;

SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- Name: schema_migration; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.schema_migration (
    version character varying(14) NOT NULL
);


ALTER TABLE public.schema_migration OWNER TO postgres;

--
-- Name: url_shorteners; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.url_shorteners (
    id integer NOT NULL,
    full_path character varying(255) NOT NULL,
    shortcut character varying(255) NOT NULL,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL
);


ALTER TABLE public.url_shorteners OWNER TO postgres;

--
-- Name: url_shorteners_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.url_shorteners_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.url_shorteners_id_seq OWNER TO postgres;

--
-- Name: url_shorteners_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.url_shorteners_id_seq OWNED BY public.url_shorteners.id;


--
-- Name: url_shorteners id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.url_shorteners ALTER COLUMN id SET DEFAULT nextval('public.url_shorteners_id_seq'::regclass);


--
-- Name: schema_migration schema_migration_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.schema_migration
    ADD CONSTRAINT schema_migration_pkey PRIMARY KEY (version);


--
-- Name: url_shorteners url_shorteners_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.url_shorteners
    ADD CONSTRAINT url_shorteners_pkey PRIMARY KEY (id);


--
-- Name: schema_migration_version_idx; Type: INDEX; Schema: public; Owner: postgres
--

CREATE UNIQUE INDEX schema_migration_version_idx ON public.schema_migration USING btree (version);


--
-- PostgreSQL database dump complete
--

\unrestrict q3AsFTUwOQOdqocyacH1u1VMKLR7W3pofren69MmTbCNMaRoRXxMQvR3REYg93A

