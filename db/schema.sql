--
-- PostgreSQL database dump
--

-- Dumped from database version 12.5
-- Dumped by pg_dump version 12.5

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

--
-- Name: trace; Type: SCHEMA; Schema: -; Owner: -
--

CREATE SCHEMA trace;


SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- Name: goose_db_version; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.goose_db_version (
    id integer NOT NULL,
    version_id bigint NOT NULL,
    is_applied boolean NOT NULL,
    tstamp timestamp without time zone DEFAULT now()
);


--
-- Name: goose_db_version_id_seq; Type: SEQUENCE; Schema: public; Owner: -
--

CREATE SEQUENCE public.goose_db_version_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


--
-- Name: goose_db_version_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: -
--

ALTER SEQUENCE public.goose_db_version_id_seq OWNED BY public.goose_db_version.id;


--
-- Name: call; Type: TABLE; Schema: trace; Owner: -
--

CREATE TABLE trace.call (
    id integer NOT NULL,
    opcode bytea NOT NULL,
    src character varying(66) NOT NULL,
    dst character varying(66) NOT NULL,
    input bytea NOT NULL,
    output bytea NOT NULL,
    value numeric NOT NULL,
    gas_used numeric NOT NULL,
    transaction_id integer NOT NULL
);


--
-- Name: TABLE call; Type: COMMENT; Schema: trace; Owner: -
--

COMMENT ON TABLE trace.call IS 'Internal calls';


--
-- Name: COLUMN call.opcode; Type: COMMENT; Schema: trace; Owner: -
--

COMMENT ON COLUMN trace.call.opcode IS 'Solidity Opcode';


--
-- Name: COLUMN call.src; Type: COMMENT; Schema: trace; Owner: -
--

COMMENT ON COLUMN trace.call.src IS 'sender of internal tx';


--
-- Name: COLUMN call.input; Type: COMMENT; Schema: trace; Owner: -
--

COMMENT ON COLUMN trace.call.input IS 'Input of internal transaction. First 4 bytes are keccak256 hash of method signature';


--
-- Name: COLUMN call.output; Type: COMMENT; Schema: trace; Owner: -
--

COMMENT ON COLUMN trace.call.output IS 'Result of internal transaction';


--
-- Name: call_id_seq; Type: SEQUENCE; Schema: trace; Owner: -
--

CREATE SEQUENCE trace.call_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


--
-- Name: call_id_seq; Type: SEQUENCE OWNED BY; Schema: trace; Owner: -
--

ALTER SEQUENCE trace.call_id_seq OWNED BY trace.call.id;


--
-- Name: transaction; Type: TABLE; Schema: trace; Owner: -
--

CREATE TABLE trace.transaction (
    id integer NOT NULL,
    tx_hash character varying(66) NOT NULL,
    index integer NOT NULL,
    block_number integer NOT NULL,
    block_hash character varying(66) NOT NULL
);


--
-- Name: COLUMN transaction.tx_hash; Type: COMMENT; Schema: trace; Owner: -
--

COMMENT ON COLUMN trace.transaction.tx_hash IS 'Transction hash';


--
-- Name: COLUMN transaction.index; Type: COMMENT; Schema: trace; Owner: -
--

COMMENT ON COLUMN trace.transaction.index IS 'Transaction index';


--
-- Name: transaction_id_seq; Type: SEQUENCE; Schema: trace; Owner: -
--

CREATE SEQUENCE trace.transaction_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


--
-- Name: transaction_id_seq; Type: SEQUENCE OWNED BY; Schema: trace; Owner: -
--

ALTER SEQUENCE trace.transaction_id_seq OWNED BY trace.transaction.id;


--
-- Name: goose_db_version id; Type: DEFAULT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.goose_db_version ALTER COLUMN id SET DEFAULT nextval('public.goose_db_version_id_seq'::regclass);


--
-- Name: call id; Type: DEFAULT; Schema: trace; Owner: -
--

ALTER TABLE ONLY trace.call ALTER COLUMN id SET DEFAULT nextval('trace.call_id_seq'::regclass);


--
-- Name: transaction id; Type: DEFAULT; Schema: trace; Owner: -
--

ALTER TABLE ONLY trace.transaction ALTER COLUMN id SET DEFAULT nextval('trace.transaction_id_seq'::regclass);


--
-- Name: goose_db_version goose_db_version_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.goose_db_version
    ADD CONSTRAINT goose_db_version_pkey PRIMARY KEY (id);


--
-- Name: call call_pk; Type: CONSTRAINT; Schema: trace; Owner: -
--

ALTER TABLE ONLY trace.call
    ADD CONSTRAINT call_pk PRIMARY KEY (id);


--
-- Name: transaction transaction_pk; Type: CONSTRAINT; Schema: trace; Owner: -
--

ALTER TABLE ONLY trace.transaction
    ADD CONSTRAINT transaction_pk PRIMARY KEY (id);


--
-- Name: transaction transaction_tx_hash_key; Type: CONSTRAINT; Schema: trace; Owner: -
--

ALTER TABLE ONLY trace.transaction
    ADD CONSTRAINT transaction_tx_hash_key UNIQUE (tx_hash);


--
-- Name: call call_transaction_id_fk; Type: FK CONSTRAINT; Schema: trace; Owner: -
--

ALTER TABLE ONLY trace.call
    ADD CONSTRAINT call_transaction_id_fk FOREIGN KEY (transaction_id) REFERENCES trace.transaction(id);


--
-- PostgreSQL database dump complete
--

