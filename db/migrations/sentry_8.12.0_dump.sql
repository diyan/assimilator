--
-- PostgreSQL database dump
--

-- Dumped from database version 9.6.1
-- Dumped by pg_dump version 9.6.1

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SET check_function_bodies = false;
SET client_min_messages = warning;
SET row_security = off;

--
-- Name: plpgsql; Type: EXTENSION; Schema: -; Owner: 
--

CREATE EXTENSION IF NOT EXISTS plpgsql WITH SCHEMA pg_catalog;


--
-- Name: EXTENSION plpgsql; Type: COMMENT; Schema: -; Owner: 
--

COMMENT ON EXTENSION plpgsql IS 'PL/pgSQL procedural language';


SET search_path = public, pg_catalog;

--
-- Name: sentry_increment_project_counter(bigint, integer); Type: FUNCTION; Schema: public; Owner: sentry
--

CREATE FUNCTION sentry_increment_project_counter(project bigint, delta integer) RETURNS integer
    LANGUAGE plpgsql
    AS $$
                declare
                  new_val int;
                begin
                  loop
                    update sentry_projectcounter set value = value + delta
                     where project_id = project
                       returning value into new_val;
                    if found then
                      return new_val;
                    end if;
                    begin
                      insert into sentry_projectcounter(project_id, value)
                           values (project, delta)
                        returning value into new_val;
                      return new_val;
                    exception when unique_violation then
                    end;
                  end loop;
                end
                $$;


ALTER FUNCTION public.sentry_increment_project_counter(project bigint, delta integer) OWNER TO sentry;

SET default_tablespace = '';

SET default_with_oids = false;

--
-- Name: auth_authenticator; Type: TABLE; Schema: public; Owner: sentry
--

CREATE TABLE auth_authenticator (
    id integer NOT NULL,
    user_id integer NOT NULL,
    created_at timestamp with time zone NOT NULL,
    last_used_at timestamp with time zone,
    type integer NOT NULL,
    config text NOT NULL,
    CONSTRAINT auth_authenticator_type_check CHECK ((type >= 0))
);


ALTER TABLE auth_authenticator OWNER TO sentry;

--
-- Name: auth_authenticator_id_seq; Type: SEQUENCE; Schema: public; Owner: sentry
--

CREATE SEQUENCE auth_authenticator_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE auth_authenticator_id_seq OWNER TO sentry;

--
-- Name: auth_authenticator_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: sentry
--

ALTER SEQUENCE auth_authenticator_id_seq OWNED BY auth_authenticator.id;


--
-- Name: auth_group; Type: TABLE; Schema: public; Owner: sentry
--

CREATE TABLE auth_group (
    id integer NOT NULL,
    name character varying(80) NOT NULL
);


ALTER TABLE auth_group OWNER TO sentry;

--
-- Name: auth_group_id_seq; Type: SEQUENCE; Schema: public; Owner: sentry
--

CREATE SEQUENCE auth_group_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE auth_group_id_seq OWNER TO sentry;

--
-- Name: auth_group_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: sentry
--

ALTER SEQUENCE auth_group_id_seq OWNED BY auth_group.id;


--
-- Name: auth_group_permissions; Type: TABLE; Schema: public; Owner: sentry
--

CREATE TABLE auth_group_permissions (
    id integer NOT NULL,
    group_id integer NOT NULL,
    permission_id integer NOT NULL
);


ALTER TABLE auth_group_permissions OWNER TO sentry;

--
-- Name: auth_group_permissions_id_seq; Type: SEQUENCE; Schema: public; Owner: sentry
--

CREATE SEQUENCE auth_group_permissions_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE auth_group_permissions_id_seq OWNER TO sentry;

--
-- Name: auth_group_permissions_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: sentry
--

ALTER SEQUENCE auth_group_permissions_id_seq OWNED BY auth_group_permissions.id;


--
-- Name: auth_permission; Type: TABLE; Schema: public; Owner: sentry
--

CREATE TABLE auth_permission (
    id integer NOT NULL,
    name character varying(50) NOT NULL,
    content_type_id integer NOT NULL,
    codename character varying(100) NOT NULL
);


ALTER TABLE auth_permission OWNER TO sentry;

--
-- Name: auth_permission_id_seq; Type: SEQUENCE; Schema: public; Owner: sentry
--

CREATE SEQUENCE auth_permission_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE auth_permission_id_seq OWNER TO sentry;

--
-- Name: auth_permission_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: sentry
--

ALTER SEQUENCE auth_permission_id_seq OWNED BY auth_permission.id;


--
-- Name: auth_user; Type: TABLE; Schema: public; Owner: sentry
--

CREATE TABLE auth_user (
    password character varying(128) NOT NULL,
    last_login timestamp with time zone NOT NULL,
    id integer NOT NULL,
    username character varying(128) NOT NULL,
    first_name character varying(200) NOT NULL,
    email character varying(75) NOT NULL,
    is_staff boolean NOT NULL,
    is_active boolean NOT NULL,
    is_superuser boolean NOT NULL,
    date_joined timestamp with time zone NOT NULL,
    is_managed boolean NOT NULL,
    is_password_expired boolean NOT NULL,
    last_password_change timestamp with time zone,
    session_nonce character varying(12)
);


ALTER TABLE auth_user OWNER TO sentry;

--
-- Name: auth_user_id_seq; Type: SEQUENCE; Schema: public; Owner: sentry
--

CREATE SEQUENCE auth_user_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE auth_user_id_seq OWNER TO sentry;

--
-- Name: auth_user_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: sentry
--

ALTER SEQUENCE auth_user_id_seq OWNED BY auth_user.id;


--
-- Name: django_admin_log; Type: TABLE; Schema: public; Owner: sentry
--

CREATE TABLE django_admin_log (
    id integer NOT NULL,
    action_time timestamp with time zone NOT NULL,
    user_id integer NOT NULL,
    content_type_id integer,
    object_id text,
    object_repr character varying(200) NOT NULL,
    action_flag smallint NOT NULL,
    change_message text NOT NULL,
    CONSTRAINT django_admin_log_action_flag_check CHECK ((action_flag >= 0))
);


ALTER TABLE django_admin_log OWNER TO sentry;

--
-- Name: django_admin_log_id_seq; Type: SEQUENCE; Schema: public; Owner: sentry
--

CREATE SEQUENCE django_admin_log_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE django_admin_log_id_seq OWNER TO sentry;

--
-- Name: django_admin_log_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: sentry
--

ALTER SEQUENCE django_admin_log_id_seq OWNED BY django_admin_log.id;


--
-- Name: django_content_type; Type: TABLE; Schema: public; Owner: sentry
--

CREATE TABLE django_content_type (
    id integer NOT NULL,
    name character varying(100) NOT NULL,
    app_label character varying(100) NOT NULL,
    model character varying(100) NOT NULL
);


ALTER TABLE django_content_type OWNER TO sentry;

--
-- Name: django_content_type_id_seq; Type: SEQUENCE; Schema: public; Owner: sentry
--

CREATE SEQUENCE django_content_type_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE django_content_type_id_seq OWNER TO sentry;

--
-- Name: django_content_type_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: sentry
--

ALTER SEQUENCE django_content_type_id_seq OWNED BY django_content_type.id;


--
-- Name: django_session; Type: TABLE; Schema: public; Owner: sentry
--

CREATE TABLE django_session (
    session_key character varying(40) NOT NULL,
    session_data text NOT NULL,
    expire_date timestamp with time zone NOT NULL
);


ALTER TABLE django_session OWNER TO sentry;

--
-- Name: django_site; Type: TABLE; Schema: public; Owner: sentry
--

CREATE TABLE django_site (
    id integer NOT NULL,
    domain character varying(100) NOT NULL,
    name character varying(50) NOT NULL
);


ALTER TABLE django_site OWNER TO sentry;

--
-- Name: django_site_id_seq; Type: SEQUENCE; Schema: public; Owner: sentry
--

CREATE SEQUENCE django_site_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE django_site_id_seq OWNER TO sentry;

--
-- Name: django_site_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: sentry
--

ALTER SEQUENCE django_site_id_seq OWNED BY django_site.id;


--
-- Name: nodestore_node; Type: TABLE; Schema: public; Owner: sentry
--

CREATE TABLE nodestore_node (
    id character varying(40) NOT NULL,
    data text NOT NULL,
    "timestamp" timestamp with time zone NOT NULL
);


ALTER TABLE nodestore_node OWNER TO sentry;

--
-- Name: sentry_activity; Type: TABLE; Schema: public; Owner: sentry
--

CREATE TABLE sentry_activity (
    id integer NOT NULL,
    project_id integer NOT NULL,
    group_id integer,
    type integer NOT NULL,
    ident character varying(64),
    user_id integer,
    datetime timestamp with time zone NOT NULL,
    data text,
    CONSTRAINT sentry_activity_type_check CHECK ((type >= 0))
);


ALTER TABLE sentry_activity OWNER TO sentry;

--
-- Name: sentry_activity_id_seq; Type: SEQUENCE; Schema: public; Owner: sentry
--

CREATE SEQUENCE sentry_activity_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE sentry_activity_id_seq OWNER TO sentry;

--
-- Name: sentry_activity_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: sentry
--

ALTER SEQUENCE sentry_activity_id_seq OWNED BY sentry_activity.id;


--
-- Name: sentry_apikey; Type: TABLE; Schema: public; Owner: sentry
--

CREATE TABLE sentry_apikey (
    id integer NOT NULL,
    organization_id integer NOT NULL,
    label character varying(64) NOT NULL,
    key character varying(32) NOT NULL,
    scopes bigint NOT NULL,
    status integer NOT NULL,
    date_added timestamp with time zone NOT NULL,
    allowed_origins text,
    CONSTRAINT sentry_apikey_status_check CHECK ((status >= 0))
);


ALTER TABLE sentry_apikey OWNER TO sentry;

--
-- Name: sentry_apikey_id_seq; Type: SEQUENCE; Schema: public; Owner: sentry
--

CREATE SEQUENCE sentry_apikey_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE sentry_apikey_id_seq OWNER TO sentry;

--
-- Name: sentry_apikey_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: sentry
--

ALTER SEQUENCE sentry_apikey_id_seq OWNED BY sentry_apikey.id;


--
-- Name: sentry_apitoken; Type: TABLE; Schema: public; Owner: sentry
--

CREATE TABLE sentry_apitoken (
    id integer NOT NULL,
    key_id integer,
    user_id integer NOT NULL,
    token character varying(64) NOT NULL,
    scopes bigint NOT NULL,
    date_added timestamp with time zone NOT NULL
);


ALTER TABLE sentry_apitoken OWNER TO sentry;

--
-- Name: sentry_apitoken_id_seq; Type: SEQUENCE; Schema: public; Owner: sentry
--

CREATE SEQUENCE sentry_apitoken_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE sentry_apitoken_id_seq OWNER TO sentry;

--
-- Name: sentry_apitoken_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: sentry
--

ALTER SEQUENCE sentry_apitoken_id_seq OWNED BY sentry_apitoken.id;


--
-- Name: sentry_auditlogentry; Type: TABLE; Schema: public; Owner: sentry
--

CREATE TABLE sentry_auditlogentry (
    id integer NOT NULL,
    organization_id integer NOT NULL,
    actor_id integer,
    target_object integer,
    target_user_id integer,
    event integer NOT NULL,
    data text NOT NULL,
    datetime timestamp with time zone NOT NULL,
    ip_address inet,
    actor_label character varying(64),
    actor_key_id integer,
    CONSTRAINT sentry_auditlogentry_event_check CHECK ((event >= 0)),
    CONSTRAINT sentry_auditlogentry_target_object_check CHECK ((target_object >= 0))
);


ALTER TABLE sentry_auditlogentry OWNER TO sentry;

--
-- Name: sentry_auditlogentry_id_seq; Type: SEQUENCE; Schema: public; Owner: sentry
--

CREATE SEQUENCE sentry_auditlogentry_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE sentry_auditlogentry_id_seq OWNER TO sentry;

--
-- Name: sentry_auditlogentry_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: sentry
--

ALTER SEQUENCE sentry_auditlogentry_id_seq OWNED BY sentry_auditlogentry.id;


--
-- Name: sentry_authidentity; Type: TABLE; Schema: public; Owner: sentry
--

CREATE TABLE sentry_authidentity (
    id integer NOT NULL,
    user_id integer NOT NULL,
    auth_provider_id integer NOT NULL,
    ident character varying(128) NOT NULL,
    data text NOT NULL,
    date_added timestamp with time zone NOT NULL,
    last_verified timestamp with time zone NOT NULL,
    last_synced timestamp with time zone NOT NULL
);


ALTER TABLE sentry_authidentity OWNER TO sentry;

--
-- Name: sentry_authidentity_id_seq; Type: SEQUENCE; Schema: public; Owner: sentry
--

CREATE SEQUENCE sentry_authidentity_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE sentry_authidentity_id_seq OWNER TO sentry;

--
-- Name: sentry_authidentity_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: sentry
--

ALTER SEQUENCE sentry_authidentity_id_seq OWNED BY sentry_authidentity.id;


--
-- Name: sentry_authprovider; Type: TABLE; Schema: public; Owner: sentry
--

CREATE TABLE sentry_authprovider (
    id integer NOT NULL,
    organization_id integer NOT NULL,
    provider character varying(128) NOT NULL,
    config text NOT NULL,
    date_added timestamp with time zone NOT NULL,
    sync_time integer,
    last_sync timestamp with time zone,
    default_role integer NOT NULL,
    default_global_access boolean NOT NULL,
    flags bigint NOT NULL,
    CONSTRAINT sentry_authprovider_default_role_check CHECK ((default_role >= 0)),
    CONSTRAINT sentry_authprovider_sync_time_check CHECK ((sync_time >= 0))
);


ALTER TABLE sentry_authprovider OWNER TO sentry;

--
-- Name: sentry_authprovider_default_teams; Type: TABLE; Schema: public; Owner: sentry
--

CREATE TABLE sentry_authprovider_default_teams (
    id integer NOT NULL,
    authprovider_id integer NOT NULL,
    team_id integer NOT NULL
);


ALTER TABLE sentry_authprovider_default_teams OWNER TO sentry;

--
-- Name: sentry_authprovider_default_teams_id_seq; Type: SEQUENCE; Schema: public; Owner: sentry
--

CREATE SEQUENCE sentry_authprovider_default_teams_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE sentry_authprovider_default_teams_id_seq OWNER TO sentry;

--
-- Name: sentry_authprovider_default_teams_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: sentry
--

ALTER SEQUENCE sentry_authprovider_default_teams_id_seq OWNED BY sentry_authprovider_default_teams.id;


--
-- Name: sentry_authprovider_id_seq; Type: SEQUENCE; Schema: public; Owner: sentry
--

CREATE SEQUENCE sentry_authprovider_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE sentry_authprovider_id_seq OWNER TO sentry;

--
-- Name: sentry_authprovider_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: sentry
--

ALTER SEQUENCE sentry_authprovider_id_seq OWNED BY sentry_authprovider.id;


--
-- Name: sentry_broadcast; Type: TABLE; Schema: public; Owner: sentry
--

CREATE TABLE sentry_broadcast (
    id integer NOT NULL,
    message character varying(256) NOT NULL,
    link character varying(200),
    is_active boolean NOT NULL,
    date_added timestamp with time zone NOT NULL,
    title character varying(32) NOT NULL,
    upstream_id character varying(32),
    date_expires timestamp with time zone
);


ALTER TABLE sentry_broadcast OWNER TO sentry;

--
-- Name: sentry_broadcast_id_seq; Type: SEQUENCE; Schema: public; Owner: sentry
--

CREATE SEQUENCE sentry_broadcast_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE sentry_broadcast_id_seq OWNER TO sentry;

--
-- Name: sentry_broadcast_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: sentry
--

ALTER SEQUENCE sentry_broadcast_id_seq OWNED BY sentry_broadcast.id;


--
-- Name: sentry_broadcastseen; Type: TABLE; Schema: public; Owner: sentry
--

CREATE TABLE sentry_broadcastseen (
    id integer NOT NULL,
    broadcast_id integer NOT NULL,
    user_id integer NOT NULL,
    date_seen timestamp with time zone NOT NULL
);


ALTER TABLE sentry_broadcastseen OWNER TO sentry;

--
-- Name: sentry_broadcastseen_id_seq; Type: SEQUENCE; Schema: public; Owner: sentry
--

CREATE SEQUENCE sentry_broadcastseen_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE sentry_broadcastseen_id_seq OWNER TO sentry;

--
-- Name: sentry_broadcastseen_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: sentry
--

ALTER SEQUENCE sentry_broadcastseen_id_seq OWNED BY sentry_broadcastseen.id;


--
-- Name: sentry_commit; Type: TABLE; Schema: public; Owner: sentry
--

CREATE TABLE sentry_commit (
    id integer NOT NULL,
    organization_id integer NOT NULL,
    repository_id integer NOT NULL,
    key character varying(64) NOT NULL,
    date_added timestamp with time zone NOT NULL,
    author_id integer,
    message text,
    CONSTRAINT sentry_commit_organization_id_check CHECK ((organization_id >= 0)),
    CONSTRAINT sentry_commit_repository_id_check CHECK ((repository_id >= 0))
);


ALTER TABLE sentry_commit OWNER TO sentry;

--
-- Name: sentry_commit_id_seq; Type: SEQUENCE; Schema: public; Owner: sentry
--

CREATE SEQUENCE sentry_commit_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE sentry_commit_id_seq OWNER TO sentry;

--
-- Name: sentry_commit_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: sentry
--

ALTER SEQUENCE sentry_commit_id_seq OWNED BY sentry_commit.id;


--
-- Name: sentry_commitauthor; Type: TABLE; Schema: public; Owner: sentry
--

CREATE TABLE sentry_commitauthor (
    id integer NOT NULL,
    organization_id integer NOT NULL,
    name character varying(128),
    email character varying(75) NOT NULL,
    CONSTRAINT sentry_commitauthor_organization_id_check CHECK ((organization_id >= 0))
);


ALTER TABLE sentry_commitauthor OWNER TO sentry;

--
-- Name: sentry_commitauthor_id_seq; Type: SEQUENCE; Schema: public; Owner: sentry
--

CREATE SEQUENCE sentry_commitauthor_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE sentry_commitauthor_id_seq OWNER TO sentry;

--
-- Name: sentry_commitauthor_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: sentry
--

ALTER SEQUENCE sentry_commitauthor_id_seq OWNED BY sentry_commitauthor.id;


--
-- Name: sentry_commitfilechange; Type: TABLE; Schema: public; Owner: sentry
--

CREATE TABLE sentry_commitfilechange (
    id integer NOT NULL,
    organization_id integer NOT NULL,
    commit_id integer NOT NULL,
    filename character varying(255) NOT NULL,
    type character varying(1) NOT NULL,
    CONSTRAINT sentry_commitfilechange_organization_id_check CHECK ((organization_id >= 0))
);


ALTER TABLE sentry_commitfilechange OWNER TO sentry;

--
-- Name: sentry_commitfilechange_id_seq; Type: SEQUENCE; Schema: public; Owner: sentry
--

CREATE SEQUENCE sentry_commitfilechange_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE sentry_commitfilechange_id_seq OWNER TO sentry;

--
-- Name: sentry_commitfilechange_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: sentry
--

ALTER SEQUENCE sentry_commitfilechange_id_seq OWNED BY sentry_commitfilechange.id;


--
-- Name: sentry_dsymbundle; Type: TABLE; Schema: public; Owner: sentry
--

CREATE TABLE sentry_dsymbundle (
    id integer NOT NULL,
    sdk_id integer NOT NULL,
    object_id integer NOT NULL
);


ALTER TABLE sentry_dsymbundle OWNER TO sentry;

--
-- Name: sentry_dsymbundle_id_seq; Type: SEQUENCE; Schema: public; Owner: sentry
--

CREATE SEQUENCE sentry_dsymbundle_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE sentry_dsymbundle_id_seq OWNER TO sentry;

--
-- Name: sentry_dsymbundle_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: sentry
--

ALTER SEQUENCE sentry_dsymbundle_id_seq OWNED BY sentry_dsymbundle.id;


--
-- Name: sentry_dsymobject; Type: TABLE; Schema: public; Owner: sentry
--

CREATE TABLE sentry_dsymobject (
    id integer NOT NULL,
    cpu_name character varying(40) NOT NULL,
    object_path text NOT NULL,
    uuid character varying(36) NOT NULL,
    vmaddr integer,
    vmsize integer
);


ALTER TABLE sentry_dsymobject OWNER TO sentry;

--
-- Name: sentry_dsymobject_id_seq; Type: SEQUENCE; Schema: public; Owner: sentry
--

CREATE SEQUENCE sentry_dsymobject_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE sentry_dsymobject_id_seq OWNER TO sentry;

--
-- Name: sentry_dsymobject_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: sentry
--

ALTER SEQUENCE sentry_dsymobject_id_seq OWNED BY sentry_dsymobject.id;


--
-- Name: sentry_dsymsdk; Type: TABLE; Schema: public; Owner: sentry
--

CREATE TABLE sentry_dsymsdk (
    id integer NOT NULL,
    dsym_type character varying(20) NOT NULL,
    sdk_name character varying(20) NOT NULL,
    version_major integer NOT NULL,
    version_minor integer NOT NULL,
    version_patchlevel integer NOT NULL,
    version_build character varying(40) NOT NULL
);


ALTER TABLE sentry_dsymsdk OWNER TO sentry;

--
-- Name: sentry_dsymsdk_id_seq; Type: SEQUENCE; Schema: public; Owner: sentry
--

CREATE SEQUENCE sentry_dsymsdk_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE sentry_dsymsdk_id_seq OWNER TO sentry;

--
-- Name: sentry_dsymsdk_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: sentry
--

ALTER SEQUENCE sentry_dsymsdk_id_seq OWNED BY sentry_dsymsdk.id;


--
-- Name: sentry_dsymsymbol; Type: TABLE; Schema: public; Owner: sentry
--

CREATE TABLE sentry_dsymsymbol (
    id integer NOT NULL,
    object_id integer NOT NULL,
    address integer NOT NULL,
    symbol text NOT NULL
);


ALTER TABLE sentry_dsymsymbol OWNER TO sentry;

--
-- Name: sentry_dsymsymbol_id_seq; Type: SEQUENCE; Schema: public; Owner: sentry
--

CREATE SEQUENCE sentry_dsymsymbol_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE sentry_dsymsymbol_id_seq OWNER TO sentry;

--
-- Name: sentry_dsymsymbol_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: sentry
--

ALTER SEQUENCE sentry_dsymsymbol_id_seq OWNED BY sentry_dsymsymbol.id;


--
-- Name: sentry_environment; Type: TABLE; Schema: public; Owner: sentry
--

CREATE TABLE sentry_environment (
    id integer NOT NULL,
    project_id integer NOT NULL,
    name character varying(64) NOT NULL,
    date_added timestamp with time zone NOT NULL,
    CONSTRAINT sentry_environment_project_id_check CHECK ((project_id >= 0))
);


ALTER TABLE sentry_environment OWNER TO sentry;

--
-- Name: sentry_environment_id_seq; Type: SEQUENCE; Schema: public; Owner: sentry
--

CREATE SEQUENCE sentry_environment_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE sentry_environment_id_seq OWNER TO sentry;

--
-- Name: sentry_environment_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: sentry
--

ALTER SEQUENCE sentry_environment_id_seq OWNED BY sentry_environment.id;


--
-- Name: sentry_environmentrelease; Type: TABLE; Schema: public; Owner: sentry
--

CREATE TABLE sentry_environmentrelease (
    id integer NOT NULL,
    project_id integer NOT NULL,
    release_id integer NOT NULL,
    environment_id integer NOT NULL,
    first_seen timestamp with time zone NOT NULL,
    last_seen timestamp with time zone NOT NULL,
    organization_id integer,
    CONSTRAINT ck_organization_id_pstv_4f21eb33d1f59511 CHECK ((organization_id >= 0)),
    CONSTRAINT sentry_environmentrelease_environment_id_check CHECK ((environment_id >= 0)),
    CONSTRAINT sentry_environmentrelease_organization_id_check CHECK ((organization_id >= 0)),
    CONSTRAINT sentry_environmentrelease_project_id_check CHECK ((project_id >= 0)),
    CONSTRAINT sentry_environmentrelease_release_id_check CHECK ((release_id >= 0))
);


ALTER TABLE sentry_environmentrelease OWNER TO sentry;

--
-- Name: sentry_environmentrelease_id_seq; Type: SEQUENCE; Schema: public; Owner: sentry
--

CREATE SEQUENCE sentry_environmentrelease_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE sentry_environmentrelease_id_seq OWNER TO sentry;

--
-- Name: sentry_environmentrelease_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: sentry
--

ALTER SEQUENCE sentry_environmentrelease_id_seq OWNED BY sentry_environmentrelease.id;


--
-- Name: sentry_eventmapping; Type: TABLE; Schema: public; Owner: sentry
--

CREATE TABLE sentry_eventmapping (
    id integer NOT NULL,
    project_id integer NOT NULL,
    group_id integer NOT NULL,
    event_id character varying(32) NOT NULL,
    date_added timestamp with time zone NOT NULL
);


ALTER TABLE sentry_eventmapping OWNER TO sentry;

--
-- Name: sentry_eventmapping_id_seq; Type: SEQUENCE; Schema: public; Owner: sentry
--

CREATE SEQUENCE sentry_eventmapping_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE sentry_eventmapping_id_seq OWNER TO sentry;

--
-- Name: sentry_eventmapping_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: sentry
--

ALTER SEQUENCE sentry_eventmapping_id_seq OWNED BY sentry_eventmapping.id;


--
-- Name: sentry_eventtag; Type: TABLE; Schema: public; Owner: sentry
--

CREATE TABLE sentry_eventtag (
    id integer NOT NULL,
    project_id integer NOT NULL,
    event_id integer NOT NULL,
    key_id integer NOT NULL,
    value_id integer NOT NULL,
    date_added timestamp with time zone NOT NULL,
    group_id integer
);


ALTER TABLE sentry_eventtag OWNER TO sentry;

--
-- Name: sentry_eventtag_id_seq; Type: SEQUENCE; Schema: public; Owner: sentry
--

CREATE SEQUENCE sentry_eventtag_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE sentry_eventtag_id_seq OWNER TO sentry;

--
-- Name: sentry_eventtag_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: sentry
--

ALTER SEQUENCE sentry_eventtag_id_seq OWNED BY sentry_eventtag.id;


--
-- Name: sentry_eventuser; Type: TABLE; Schema: public; Owner: sentry
--

CREATE TABLE sentry_eventuser (
    id integer NOT NULL,
    project_id integer NOT NULL,
    ident character varying(128),
    email character varying(75),
    username character varying(128),
    ip_address inet,
    date_added timestamp with time zone NOT NULL,
    hash character varying(32) NOT NULL
);


ALTER TABLE sentry_eventuser OWNER TO sentry;

--
-- Name: sentry_eventuser_id_seq; Type: SEQUENCE; Schema: public; Owner: sentry
--

CREATE SEQUENCE sentry_eventuser_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE sentry_eventuser_id_seq OWNER TO sentry;

--
-- Name: sentry_eventuser_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: sentry
--

ALTER SEQUENCE sentry_eventuser_id_seq OWNED BY sentry_eventuser.id;


--
-- Name: sentry_file; Type: TABLE; Schema: public; Owner: sentry
--

CREATE TABLE sentry_file (
    id integer NOT NULL,
    name character varying(128) NOT NULL,
    path text,
    type character varying(64) NOT NULL,
    size integer,
    "timestamp" timestamp with time zone NOT NULL,
    checksum character varying(40),
    headers text NOT NULL,
    blob_id integer,
    CONSTRAINT sentry_file_size_check CHECK ((size >= 0))
);


ALTER TABLE sentry_file OWNER TO sentry;

--
-- Name: sentry_file_id_seq; Type: SEQUENCE; Schema: public; Owner: sentry
--

CREATE SEQUENCE sentry_file_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE sentry_file_id_seq OWNER TO sentry;

--
-- Name: sentry_file_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: sentry
--

ALTER SEQUENCE sentry_file_id_seq OWNED BY sentry_file.id;


--
-- Name: sentry_fileblob; Type: TABLE; Schema: public; Owner: sentry
--

CREATE TABLE sentry_fileblob (
    id integer NOT NULL,
    path text,
    size integer,
    checksum character varying(40) NOT NULL,
    "timestamp" timestamp with time zone NOT NULL,
    CONSTRAINT sentry_fileblob_size_check CHECK ((size >= 0))
);


ALTER TABLE sentry_fileblob OWNER TO sentry;

--
-- Name: sentry_fileblob_id_seq; Type: SEQUENCE; Schema: public; Owner: sentry
--

CREATE SEQUENCE sentry_fileblob_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE sentry_fileblob_id_seq OWNER TO sentry;

--
-- Name: sentry_fileblob_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: sentry
--

ALTER SEQUENCE sentry_fileblob_id_seq OWNED BY sentry_fileblob.id;


--
-- Name: sentry_fileblobindex; Type: TABLE; Schema: public; Owner: sentry
--

CREATE TABLE sentry_fileblobindex (
    id integer NOT NULL,
    file_id integer NOT NULL,
    blob_id integer NOT NULL,
    "offset" integer NOT NULL,
    CONSTRAINT sentry_fileblobindex_offset_check CHECK (("offset" >= 0))
);


ALTER TABLE sentry_fileblobindex OWNER TO sentry;

--
-- Name: sentry_fileblobindex_id_seq; Type: SEQUENCE; Schema: public; Owner: sentry
--

CREATE SEQUENCE sentry_fileblobindex_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE sentry_fileblobindex_id_seq OWNER TO sentry;

--
-- Name: sentry_fileblobindex_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: sentry
--

ALTER SEQUENCE sentry_fileblobindex_id_seq OWNED BY sentry_fileblobindex.id;


--
-- Name: sentry_filterkey; Type: TABLE; Schema: public; Owner: sentry
--

CREATE TABLE sentry_filterkey (
    id integer NOT NULL,
    project_id integer NOT NULL,
    key character varying(32) NOT NULL,
    values_seen integer NOT NULL,
    label character varying(64),
    status integer NOT NULL,
    CONSTRAINT ck_status_pstv_56aaa5973127b013 CHECK ((status >= 0)),
    CONSTRAINT ck_values_seen_pstv_12eab0d3ff94a35c CHECK ((values_seen >= 0)),
    CONSTRAINT sentry_filterkey_status_check CHECK ((status >= 0)),
    CONSTRAINT sentry_filterkey_values_seen_check CHECK ((values_seen >= 0))
);


ALTER TABLE sentry_filterkey OWNER TO sentry;

--
-- Name: sentry_filterkey_id_seq; Type: SEQUENCE; Schema: public; Owner: sentry
--

CREATE SEQUENCE sentry_filterkey_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE sentry_filterkey_id_seq OWNER TO sentry;

--
-- Name: sentry_filterkey_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: sentry
--

ALTER SEQUENCE sentry_filterkey_id_seq OWNED BY sentry_filterkey.id;


--
-- Name: sentry_filtervalue; Type: TABLE; Schema: public; Owner: sentry
--

CREATE TABLE sentry_filtervalue (
    id integer NOT NULL,
    key character varying(32) NOT NULL,
    value character varying(200) NOT NULL,
    project_id integer,
    times_seen integer NOT NULL,
    last_seen timestamp with time zone,
    first_seen timestamp with time zone,
    data text,
    CONSTRAINT ck_times_seen_pstv_10c4372f28cef967 CHECK ((times_seen >= 0)),
    CONSTRAINT sentry_filtervalue_times_seen_check CHECK ((times_seen >= 0))
);


ALTER TABLE sentry_filtervalue OWNER TO sentry;

--
-- Name: sentry_filtervalue_id_seq; Type: SEQUENCE; Schema: public; Owner: sentry
--

CREATE SEQUENCE sentry_filtervalue_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE sentry_filtervalue_id_seq OWNER TO sentry;

--
-- Name: sentry_filtervalue_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: sentry
--

ALTER SEQUENCE sentry_filtervalue_id_seq OWNED BY sentry_filtervalue.id;


--
-- Name: sentry_globaldsymfile; Type: TABLE; Schema: public; Owner: sentry
--

CREATE TABLE sentry_globaldsymfile (
    id integer NOT NULL,
    file_id integer NOT NULL,
    object_name text NOT NULL,
    cpu_name character varying(40) NOT NULL,
    uuid character varying(36) NOT NULL
);


ALTER TABLE sentry_globaldsymfile OWNER TO sentry;

--
-- Name: sentry_globaldsymfile_id_seq; Type: SEQUENCE; Schema: public; Owner: sentry
--

CREATE SEQUENCE sentry_globaldsymfile_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE sentry_globaldsymfile_id_seq OWNER TO sentry;

--
-- Name: sentry_globaldsymfile_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: sentry
--

ALTER SEQUENCE sentry_globaldsymfile_id_seq OWNED BY sentry_globaldsymfile.id;


--
-- Name: sentry_groupasignee; Type: TABLE; Schema: public; Owner: sentry
--

CREATE TABLE sentry_groupasignee (
    id integer NOT NULL,
    project_id integer NOT NULL,
    group_id integer NOT NULL,
    user_id integer NOT NULL,
    date_added timestamp with time zone NOT NULL
);


ALTER TABLE sentry_groupasignee OWNER TO sentry;

--
-- Name: sentry_groupasignee_id_seq; Type: SEQUENCE; Schema: public; Owner: sentry
--

CREATE SEQUENCE sentry_groupasignee_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE sentry_groupasignee_id_seq OWNER TO sentry;

--
-- Name: sentry_groupasignee_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: sentry
--

ALTER SEQUENCE sentry_groupasignee_id_seq OWNED BY sentry_groupasignee.id;


--
-- Name: sentry_groupbookmark; Type: TABLE; Schema: public; Owner: sentry
--

CREATE TABLE sentry_groupbookmark (
    id integer NOT NULL,
    project_id integer NOT NULL,
    group_id integer NOT NULL,
    user_id integer NOT NULL,
    date_added timestamp with time zone
);


ALTER TABLE sentry_groupbookmark OWNER TO sentry;

--
-- Name: sentry_groupbookmark_id_seq; Type: SEQUENCE; Schema: public; Owner: sentry
--

CREATE SEQUENCE sentry_groupbookmark_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE sentry_groupbookmark_id_seq OWNER TO sentry;

--
-- Name: sentry_groupbookmark_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: sentry
--

ALTER SEQUENCE sentry_groupbookmark_id_seq OWNED BY sentry_groupbookmark.id;


--
-- Name: sentry_groupedmessage; Type: TABLE; Schema: public; Owner: sentry
--

CREATE TABLE sentry_groupedmessage (
    id integer NOT NULL,
    logger character varying(64) NOT NULL,
    level integer NOT NULL,
    message text NOT NULL,
    view character varying(200),
    status integer NOT NULL,
    times_seen integer NOT NULL,
    last_seen timestamp with time zone NOT NULL,
    first_seen timestamp with time zone NOT NULL,
    data text,
    score integer NOT NULL,
    project_id integer,
    time_spent_total integer NOT NULL,
    time_spent_count integer NOT NULL,
    resolved_at timestamp with time zone,
    active_at timestamp with time zone,
    is_public boolean,
    platform character varying(64),
    num_comments integer,
    first_release_id integer,
    short_id integer,
    CONSTRAINT ck_num_comments_pstv_44851d4d5d739eab CHECK ((num_comments >= 0)),
    CONSTRAINT sentry_groupedmessage_level_check CHECK ((level >= 0)),
    CONSTRAINT sentry_groupedmessage_num_comments_check CHECK ((num_comments >= 0)),
    CONSTRAINT sentry_groupedmessage_status_check CHECK ((status >= 0)),
    CONSTRAINT sentry_groupedmessage_times_seen_check CHECK ((times_seen >= 0))
);


ALTER TABLE sentry_groupedmessage OWNER TO sentry;

--
-- Name: sentry_groupedmessage_id_seq; Type: SEQUENCE; Schema: public; Owner: sentry
--

CREATE SEQUENCE sentry_groupedmessage_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE sentry_groupedmessage_id_seq OWNER TO sentry;

--
-- Name: sentry_groupedmessage_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: sentry
--

ALTER SEQUENCE sentry_groupedmessage_id_seq OWNED BY sentry_groupedmessage.id;


--
-- Name: sentry_groupemailthread; Type: TABLE; Schema: public; Owner: sentry
--

CREATE TABLE sentry_groupemailthread (
    id integer NOT NULL,
    email character varying(75) NOT NULL,
    project_id integer NOT NULL,
    group_id integer NOT NULL,
    msgid character varying(100) NOT NULL,
    date timestamp with time zone NOT NULL
);


ALTER TABLE sentry_groupemailthread OWNER TO sentry;

--
-- Name: sentry_groupemailthread_id_seq; Type: SEQUENCE; Schema: public; Owner: sentry
--

CREATE SEQUENCE sentry_groupemailthread_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE sentry_groupemailthread_id_seq OWNER TO sentry;

--
-- Name: sentry_groupemailthread_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: sentry
--

ALTER SEQUENCE sentry_groupemailthread_id_seq OWNED BY sentry_groupemailthread.id;


--
-- Name: sentry_grouphash; Type: TABLE; Schema: public; Owner: sentry
--

CREATE TABLE sentry_grouphash (
    id integer NOT NULL,
    project_id integer,
    hash character varying(32) NOT NULL,
    group_id integer
);


ALTER TABLE sentry_grouphash OWNER TO sentry;

--
-- Name: sentry_grouphash_id_seq; Type: SEQUENCE; Schema: public; Owner: sentry
--

CREATE SEQUENCE sentry_grouphash_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE sentry_grouphash_id_seq OWNER TO sentry;

--
-- Name: sentry_grouphash_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: sentry
--

ALTER SEQUENCE sentry_grouphash_id_seq OWNED BY sentry_grouphash.id;


--
-- Name: sentry_groupmeta; Type: TABLE; Schema: public; Owner: sentry
--

CREATE TABLE sentry_groupmeta (
    id integer NOT NULL,
    group_id integer NOT NULL,
    key character varying(64) NOT NULL,
    value text NOT NULL
);


ALTER TABLE sentry_groupmeta OWNER TO sentry;

--
-- Name: sentry_groupmeta_id_seq; Type: SEQUENCE; Schema: public; Owner: sentry
--

CREATE SEQUENCE sentry_groupmeta_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE sentry_groupmeta_id_seq OWNER TO sentry;

--
-- Name: sentry_groupmeta_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: sentry
--

ALTER SEQUENCE sentry_groupmeta_id_seq OWNED BY sentry_groupmeta.id;


--
-- Name: sentry_groupredirect; Type: TABLE; Schema: public; Owner: sentry
--

CREATE TABLE sentry_groupredirect (
    id integer NOT NULL,
    group_id integer NOT NULL,
    previous_group_id integer NOT NULL
);


ALTER TABLE sentry_groupredirect OWNER TO sentry;

--
-- Name: sentry_groupredirect_id_seq; Type: SEQUENCE; Schema: public; Owner: sentry
--

CREATE SEQUENCE sentry_groupredirect_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE sentry_groupredirect_id_seq OWNER TO sentry;

--
-- Name: sentry_groupredirect_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: sentry
--

ALTER SEQUENCE sentry_groupredirect_id_seq OWNED BY sentry_groupredirect.id;


--
-- Name: sentry_grouprelease; Type: TABLE; Schema: public; Owner: sentry
--

CREATE TABLE sentry_grouprelease (
    id integer NOT NULL,
    project_id integer NOT NULL,
    group_id integer NOT NULL,
    release_id integer NOT NULL,
    environment character varying(64) NOT NULL,
    first_seen timestamp with time zone NOT NULL,
    last_seen timestamp with time zone NOT NULL,
    CONSTRAINT sentry_grouprelease_group_id_check CHECK ((group_id >= 0)),
    CONSTRAINT sentry_grouprelease_project_id_check CHECK ((project_id >= 0)),
    CONSTRAINT sentry_grouprelease_release_id_check CHECK ((release_id >= 0))
);


ALTER TABLE sentry_grouprelease OWNER TO sentry;

--
-- Name: sentry_grouprelease_id_seq; Type: SEQUENCE; Schema: public; Owner: sentry
--

CREATE SEQUENCE sentry_grouprelease_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE sentry_grouprelease_id_seq OWNER TO sentry;

--
-- Name: sentry_grouprelease_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: sentry
--

ALTER SEQUENCE sentry_grouprelease_id_seq OWNED BY sentry_grouprelease.id;


--
-- Name: sentry_groupresolution; Type: TABLE; Schema: public; Owner: sentry
--

CREATE TABLE sentry_groupresolution (
    id integer NOT NULL,
    group_id integer NOT NULL,
    release_id integer NOT NULL,
    datetime timestamp with time zone NOT NULL,
    status integer NOT NULL,
    CONSTRAINT ck_status_pstv_375a4efcf0df73b9 CHECK ((status >= 0)),
    CONSTRAINT sentry_groupresolution_status_check CHECK ((status >= 0))
);


ALTER TABLE sentry_groupresolution OWNER TO sentry;

--
-- Name: sentry_groupresolution_id_seq; Type: SEQUENCE; Schema: public; Owner: sentry
--

CREATE SEQUENCE sentry_groupresolution_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE sentry_groupresolution_id_seq OWNER TO sentry;

--
-- Name: sentry_groupresolution_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: sentry
--

ALTER SEQUENCE sentry_groupresolution_id_seq OWNED BY sentry_groupresolution.id;


--
-- Name: sentry_grouprulestatus; Type: TABLE; Schema: public; Owner: sentry
--

CREATE TABLE sentry_grouprulestatus (
    id integer NOT NULL,
    project_id integer NOT NULL,
    rule_id integer NOT NULL,
    group_id integer NOT NULL,
    status smallint NOT NULL,
    date_added timestamp with time zone NOT NULL,
    last_active timestamp with time zone,
    CONSTRAINT sentry_grouprulestatus_status_check CHECK ((status >= 0))
);


ALTER TABLE sentry_grouprulestatus OWNER TO sentry;

--
-- Name: sentry_grouprulestatus_id_seq; Type: SEQUENCE; Schema: public; Owner: sentry
--

CREATE SEQUENCE sentry_grouprulestatus_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE sentry_grouprulestatus_id_seq OWNER TO sentry;

--
-- Name: sentry_grouprulestatus_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: sentry
--

ALTER SEQUENCE sentry_grouprulestatus_id_seq OWNED BY sentry_grouprulestatus.id;


--
-- Name: sentry_groupseen; Type: TABLE; Schema: public; Owner: sentry
--

CREATE TABLE sentry_groupseen (
    id integer NOT NULL,
    project_id integer NOT NULL,
    group_id integer NOT NULL,
    user_id integer NOT NULL,
    last_seen timestamp with time zone NOT NULL
);


ALTER TABLE sentry_groupseen OWNER TO sentry;

--
-- Name: sentry_groupseen_id_seq; Type: SEQUENCE; Schema: public; Owner: sentry
--

CREATE SEQUENCE sentry_groupseen_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE sentry_groupseen_id_seq OWNER TO sentry;

--
-- Name: sentry_groupseen_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: sentry
--

ALTER SEQUENCE sentry_groupseen_id_seq OWNED BY sentry_groupseen.id;


--
-- Name: sentry_groupsnooze; Type: TABLE; Schema: public; Owner: sentry
--

CREATE TABLE sentry_groupsnooze (
    id integer NOT NULL,
    group_id integer NOT NULL,
    until timestamp with time zone NOT NULL
);


ALTER TABLE sentry_groupsnooze OWNER TO sentry;

--
-- Name: sentry_groupsnooze_id_seq; Type: SEQUENCE; Schema: public; Owner: sentry
--

CREATE SEQUENCE sentry_groupsnooze_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE sentry_groupsnooze_id_seq OWNER TO sentry;

--
-- Name: sentry_groupsnooze_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: sentry
--

ALTER SEQUENCE sentry_groupsnooze_id_seq OWNED BY sentry_groupsnooze.id;


--
-- Name: sentry_groupsubscription; Type: TABLE; Schema: public; Owner: sentry
--

CREATE TABLE sentry_groupsubscription (
    id integer NOT NULL,
    project_id integer NOT NULL,
    group_id integer NOT NULL,
    user_id integer NOT NULL,
    is_active boolean NOT NULL,
    reason integer NOT NULL,
    date_added timestamp with time zone,
    CONSTRAINT sentry_groupsubscription_reason_check CHECK ((reason >= 0))
);


ALTER TABLE sentry_groupsubscription OWNER TO sentry;

--
-- Name: sentry_groupsubscription_id_seq; Type: SEQUENCE; Schema: public; Owner: sentry
--

CREATE SEQUENCE sentry_groupsubscription_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE sentry_groupsubscription_id_seq OWNER TO sentry;

--
-- Name: sentry_groupsubscription_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: sentry
--

ALTER SEQUENCE sentry_groupsubscription_id_seq OWNED BY sentry_groupsubscription.id;


--
-- Name: sentry_grouptagkey; Type: TABLE; Schema: public; Owner: sentry
--

CREATE TABLE sentry_grouptagkey (
    id integer NOT NULL,
    project_id integer,
    group_id integer NOT NULL,
    key character varying(32) NOT NULL,
    values_seen integer NOT NULL,
    CONSTRAINT sentry_grouptagkey_values_seen_check CHECK ((values_seen >= 0))
);


ALTER TABLE sentry_grouptagkey OWNER TO sentry;

--
-- Name: sentry_grouptagkey_id_seq; Type: SEQUENCE; Schema: public; Owner: sentry
--

CREATE SEQUENCE sentry_grouptagkey_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE sentry_grouptagkey_id_seq OWNER TO sentry;

--
-- Name: sentry_grouptagkey_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: sentry
--

ALTER SEQUENCE sentry_grouptagkey_id_seq OWNED BY sentry_grouptagkey.id;


--
-- Name: sentry_lostpasswordhash; Type: TABLE; Schema: public; Owner: sentry
--

CREATE TABLE sentry_lostpasswordhash (
    id integer NOT NULL,
    user_id integer NOT NULL,
    hash character varying(32) NOT NULL,
    date_added timestamp with time zone NOT NULL
);


ALTER TABLE sentry_lostpasswordhash OWNER TO sentry;

--
-- Name: sentry_lostpasswordhash_id_seq; Type: SEQUENCE; Schema: public; Owner: sentry
--

CREATE SEQUENCE sentry_lostpasswordhash_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE sentry_lostpasswordhash_id_seq OWNER TO sentry;

--
-- Name: sentry_lostpasswordhash_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: sentry
--

ALTER SEQUENCE sentry_lostpasswordhash_id_seq OWNED BY sentry_lostpasswordhash.id;


--
-- Name: sentry_message; Type: TABLE; Schema: public; Owner: sentry
--

CREATE TABLE sentry_message (
    id integer NOT NULL,
    message text NOT NULL,
    datetime timestamp with time zone NOT NULL,
    data text,
    group_id integer,
    message_id character varying(32),
    project_id integer,
    time_spent integer,
    platform character varying(64)
);


ALTER TABLE sentry_message OWNER TO sentry;

--
-- Name: sentry_message_id_seq; Type: SEQUENCE; Schema: public; Owner: sentry
--

CREATE SEQUENCE sentry_message_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE sentry_message_id_seq OWNER TO sentry;

--
-- Name: sentry_message_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: sentry
--

ALTER SEQUENCE sentry_message_id_seq OWNED BY sentry_message.id;


--
-- Name: sentry_messagefiltervalue; Type: TABLE; Schema: public; Owner: sentry
--

CREATE TABLE sentry_messagefiltervalue (
    id integer NOT NULL,
    group_id integer NOT NULL,
    times_seen integer NOT NULL,
    key character varying(32) NOT NULL,
    value character varying(200) NOT NULL,
    project_id integer,
    last_seen timestamp with time zone,
    first_seen timestamp with time zone,
    CONSTRAINT sentry_messagefiltervalue_times_seen_check CHECK ((times_seen >= 0))
);


ALTER TABLE sentry_messagefiltervalue OWNER TO sentry;

--
-- Name: sentry_messagefiltervalue_id_seq; Type: SEQUENCE; Schema: public; Owner: sentry
--

CREATE SEQUENCE sentry_messagefiltervalue_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE sentry_messagefiltervalue_id_seq OWNER TO sentry;

--
-- Name: sentry_messagefiltervalue_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: sentry
--

ALTER SEQUENCE sentry_messagefiltervalue_id_seq OWNED BY sentry_messagefiltervalue.id;


--
-- Name: sentry_messageindex; Type: TABLE; Schema: public; Owner: sentry
--

CREATE TABLE sentry_messageindex (
    id integer NOT NULL,
    object_id integer NOT NULL,
    "column" character varying(32) NOT NULL,
    value character varying(128) NOT NULL,
    CONSTRAINT sentry_messageindex_object_id_check CHECK ((object_id >= 0))
);


ALTER TABLE sentry_messageindex OWNER TO sentry;

--
-- Name: sentry_messageindex_id_seq; Type: SEQUENCE; Schema: public; Owner: sentry
--

CREATE SEQUENCE sentry_messageindex_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE sentry_messageindex_id_seq OWNER TO sentry;

--
-- Name: sentry_messageindex_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: sentry
--

ALTER SEQUENCE sentry_messageindex_id_seq OWNED BY sentry_messageindex.id;


--
-- Name: sentry_option; Type: TABLE; Schema: public; Owner: sentry
--

CREATE TABLE sentry_option (
    id integer NOT NULL,
    key character varying(64) NOT NULL,
    value text NOT NULL,
    last_updated timestamp with time zone NOT NULL
);


ALTER TABLE sentry_option OWNER TO sentry;

--
-- Name: sentry_option_id_seq; Type: SEQUENCE; Schema: public; Owner: sentry
--

CREATE SEQUENCE sentry_option_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE sentry_option_id_seq OWNER TO sentry;

--
-- Name: sentry_option_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: sentry
--

ALTER SEQUENCE sentry_option_id_seq OWNED BY sentry_option.id;


--
-- Name: sentry_organization; Type: TABLE; Schema: public; Owner: sentry
--

CREATE TABLE sentry_organization (
    id integer NOT NULL,
    name character varying(64) NOT NULL,
    status integer NOT NULL,
    date_added timestamp with time zone NOT NULL,
    slug character varying(50) NOT NULL,
    flags bigint NOT NULL,
    default_role character varying(32) NOT NULL,
    CONSTRAINT sentry_organization_status_check CHECK ((status >= 0))
);


ALTER TABLE sentry_organization OWNER TO sentry;

--
-- Name: sentry_organization_id_seq; Type: SEQUENCE; Schema: public; Owner: sentry
--

CREATE SEQUENCE sentry_organization_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE sentry_organization_id_seq OWNER TO sentry;

--
-- Name: sentry_organization_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: sentry
--

ALTER SEQUENCE sentry_organization_id_seq OWNED BY sentry_organization.id;


--
-- Name: sentry_organizationaccessrequest; Type: TABLE; Schema: public; Owner: sentry
--

CREATE TABLE sentry_organizationaccessrequest (
    id integer NOT NULL,
    team_id integer NOT NULL,
    member_id integer NOT NULL
);


ALTER TABLE sentry_organizationaccessrequest OWNER TO sentry;

--
-- Name: sentry_organizationaccessrequest_id_seq; Type: SEQUENCE; Schema: public; Owner: sentry
--

CREATE SEQUENCE sentry_organizationaccessrequest_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE sentry_organizationaccessrequest_id_seq OWNER TO sentry;

--
-- Name: sentry_organizationaccessrequest_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: sentry
--

ALTER SEQUENCE sentry_organizationaccessrequest_id_seq OWNED BY sentry_organizationaccessrequest.id;


--
-- Name: sentry_organizationmember; Type: TABLE; Schema: public; Owner: sentry
--

CREATE TABLE sentry_organizationmember (
    id integer NOT NULL,
    organization_id integer NOT NULL,
    user_id integer,
    type integer NOT NULL,
    date_added timestamp with time zone NOT NULL,
    email character varying(75),
    has_global_access boolean NOT NULL,
    flags bigint NOT NULL,
    role character varying(32) NOT NULL,
    token character varying(64),
    CONSTRAINT sentry_organizationmember_type_check CHECK ((type >= 0))
);


ALTER TABLE sentry_organizationmember OWNER TO sentry;

--
-- Name: sentry_organizationmember_id_seq; Type: SEQUENCE; Schema: public; Owner: sentry
--

CREATE SEQUENCE sentry_organizationmember_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE sentry_organizationmember_id_seq OWNER TO sentry;

--
-- Name: sentry_organizationmember_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: sentry
--

ALTER SEQUENCE sentry_organizationmember_id_seq OWNED BY sentry_organizationmember.id;


--
-- Name: sentry_organizationmember_teams; Type: TABLE; Schema: public; Owner: sentry
--

CREATE TABLE sentry_organizationmember_teams (
    id integer NOT NULL,
    organizationmember_id integer NOT NULL,
    team_id integer NOT NULL,
    is_active boolean NOT NULL
);


ALTER TABLE sentry_organizationmember_teams OWNER TO sentry;

--
-- Name: sentry_organizationmember_teams_id_seq; Type: SEQUENCE; Schema: public; Owner: sentry
--

CREATE SEQUENCE sentry_organizationmember_teams_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE sentry_organizationmember_teams_id_seq OWNER TO sentry;

--
-- Name: sentry_organizationmember_teams_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: sentry
--

ALTER SEQUENCE sentry_organizationmember_teams_id_seq OWNED BY sentry_organizationmember_teams.id;


--
-- Name: sentry_organizationonboardingtask; Type: TABLE; Schema: public; Owner: sentry
--

CREATE TABLE sentry_organizationonboardingtask (
    id integer NOT NULL,
    organization_id integer NOT NULL,
    user_id integer,
    task integer NOT NULL,
    status integer NOT NULL,
    date_completed timestamp with time zone NOT NULL,
    project_id integer,
    data text NOT NULL,
    CONSTRAINT sentry_organizationonboardingtask_status_check CHECK ((status >= 0)),
    CONSTRAINT sentry_organizationonboardingtask_task_check CHECK ((task >= 0))
);


ALTER TABLE sentry_organizationonboardingtask OWNER TO sentry;

--
-- Name: sentry_organizationonboardingtask_id_seq; Type: SEQUENCE; Schema: public; Owner: sentry
--

CREATE SEQUENCE sentry_organizationonboardingtask_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE sentry_organizationonboardingtask_id_seq OWNER TO sentry;

--
-- Name: sentry_organizationonboardingtask_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: sentry
--

ALTER SEQUENCE sentry_organizationonboardingtask_id_seq OWNED BY sentry_organizationonboardingtask.id;


--
-- Name: sentry_organizationoptions; Type: TABLE; Schema: public; Owner: sentry
--

CREATE TABLE sentry_organizationoptions (
    id integer NOT NULL,
    organization_id integer NOT NULL,
    key character varying(64) NOT NULL,
    value text NOT NULL
);


ALTER TABLE sentry_organizationoptions OWNER TO sentry;

--
-- Name: sentry_organizationoptions_id_seq; Type: SEQUENCE; Schema: public; Owner: sentry
--

CREATE SEQUENCE sentry_organizationoptions_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE sentry_organizationoptions_id_seq OWNER TO sentry;

--
-- Name: sentry_organizationoptions_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: sentry
--

ALTER SEQUENCE sentry_organizationoptions_id_seq OWNED BY sentry_organizationoptions.id;


--
-- Name: sentry_project; Type: TABLE; Schema: public; Owner: sentry
--

CREATE TABLE sentry_project (
    id integer NOT NULL,
    name character varying(200) NOT NULL,
    public boolean NOT NULL,
    date_added timestamp with time zone NOT NULL,
    status integer NOT NULL,
    slug character varying(50),
    team_id integer NOT NULL,
    organization_id integer NOT NULL,
    first_event timestamp with time zone,
    forced_color character varying(6),
    CONSTRAINT ck_status_pstv_3af8360b8a37db73 CHECK ((status >= 0)),
    CONSTRAINT sentry_project_status_check CHECK ((status >= 0))
);


ALTER TABLE sentry_project OWNER TO sentry;

--
-- Name: sentry_project_id_seq; Type: SEQUENCE; Schema: public; Owner: sentry
--

CREATE SEQUENCE sentry_project_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE sentry_project_id_seq OWNER TO sentry;

--
-- Name: sentry_project_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: sentry
--

ALTER SEQUENCE sentry_project_id_seq OWNED BY sentry_project.id;


--
-- Name: sentry_projectbookmark; Type: TABLE; Schema: public; Owner: sentry
--

CREATE TABLE sentry_projectbookmark (
    id integer NOT NULL,
    project_id integer,
    user_id integer NOT NULL,
    date_added timestamp with time zone
);


ALTER TABLE sentry_projectbookmark OWNER TO sentry;

--
-- Name: sentry_projectbookmark_id_seq; Type: SEQUENCE; Schema: public; Owner: sentry
--

CREATE SEQUENCE sentry_projectbookmark_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE sentry_projectbookmark_id_seq OWNER TO sentry;

--
-- Name: sentry_projectbookmark_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: sentry
--

ALTER SEQUENCE sentry_projectbookmark_id_seq OWNED BY sentry_projectbookmark.id;


--
-- Name: sentry_projectcounter; Type: TABLE; Schema: public; Owner: sentry
--

CREATE TABLE sentry_projectcounter (
    id integer NOT NULL,
    project_id integer NOT NULL,
    value integer NOT NULL
);


ALTER TABLE sentry_projectcounter OWNER TO sentry;

--
-- Name: sentry_projectcounter_id_seq; Type: SEQUENCE; Schema: public; Owner: sentry
--

CREATE SEQUENCE sentry_projectcounter_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE sentry_projectcounter_id_seq OWNER TO sentry;

--
-- Name: sentry_projectcounter_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: sentry
--

ALTER SEQUENCE sentry_projectcounter_id_seq OWNED BY sentry_projectcounter.id;


--
-- Name: sentry_projectdsymfile; Type: TABLE; Schema: public; Owner: sentry
--

CREATE TABLE sentry_projectdsymfile (
    id integer NOT NULL,
    file_id integer NOT NULL,
    object_name text NOT NULL,
    cpu_name character varying(40) NOT NULL,
    project_id integer,
    uuid character varying(36) NOT NULL
);


ALTER TABLE sentry_projectdsymfile OWNER TO sentry;

--
-- Name: sentry_projectdsymfile_id_seq; Type: SEQUENCE; Schema: public; Owner: sentry
--

CREATE SEQUENCE sentry_projectdsymfile_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE sentry_projectdsymfile_id_seq OWNER TO sentry;

--
-- Name: sentry_projectdsymfile_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: sentry
--

ALTER SEQUENCE sentry_projectdsymfile_id_seq OWNED BY sentry_projectdsymfile.id;


--
-- Name: sentry_projectkey; Type: TABLE; Schema: public; Owner: sentry
--

CREATE TABLE sentry_projectkey (
    id integer NOT NULL,
    project_id integer NOT NULL,
    public_key character varying(32),
    secret_key character varying(32),
    date_added timestamp with time zone,
    roles bigint NOT NULL,
    label character varying(64),
    status integer NOT NULL,
    CONSTRAINT ck_status_pstv_1f17c0d00e89ed63 CHECK ((status >= 0)),
    CONSTRAINT sentry_projectkey_status_check CHECK ((status >= 0))
);


ALTER TABLE sentry_projectkey OWNER TO sentry;

--
-- Name: sentry_projectkey_id_seq; Type: SEQUENCE; Schema: public; Owner: sentry
--

CREATE SEQUENCE sentry_projectkey_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE sentry_projectkey_id_seq OWNER TO sentry;

--
-- Name: sentry_projectkey_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: sentry
--

ALTER SEQUENCE sentry_projectkey_id_seq OWNED BY sentry_projectkey.id;


--
-- Name: sentry_projectoptions; Type: TABLE; Schema: public; Owner: sentry
--

CREATE TABLE sentry_projectoptions (
    id integer NOT NULL,
    project_id integer NOT NULL,
    key character varying(64) NOT NULL,
    value text NOT NULL
);


ALTER TABLE sentry_projectoptions OWNER TO sentry;

--
-- Name: sentry_projectoptions_id_seq; Type: SEQUENCE; Schema: public; Owner: sentry
--

CREATE SEQUENCE sentry_projectoptions_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE sentry_projectoptions_id_seq OWNER TO sentry;

--
-- Name: sentry_projectoptions_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: sentry
--

ALTER SEQUENCE sentry_projectoptions_id_seq OWNED BY sentry_projectoptions.id;


--
-- Name: sentry_projectplatform; Type: TABLE; Schema: public; Owner: sentry
--

CREATE TABLE sentry_projectplatform (
    id integer NOT NULL,
    project_id integer NOT NULL,
    platform character varying(64) NOT NULL,
    date_added timestamp with time zone NOT NULL,
    last_seen timestamp with time zone NOT NULL
);


ALTER TABLE sentry_projectplatform OWNER TO sentry;

--
-- Name: sentry_projectplatform_id_seq; Type: SEQUENCE; Schema: public; Owner: sentry
--

CREATE SEQUENCE sentry_projectplatform_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE sentry_projectplatform_id_seq OWNER TO sentry;

--
-- Name: sentry_projectplatform_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: sentry
--

ALTER SEQUENCE sentry_projectplatform_id_seq OWNED BY sentry_projectplatform.id;


--
-- Name: sentry_release; Type: TABLE; Schema: public; Owner: sentry
--

CREATE TABLE sentry_release (
    id integer NOT NULL,
    project_id integer NOT NULL,
    version character varying(64) NOT NULL,
    date_added timestamp with time zone NOT NULL,
    date_released timestamp with time zone,
    ref character varying(64),
    url character varying(200),
    date_started timestamp with time zone,
    data text NOT NULL,
    new_groups integer NOT NULL,
    owner_id integer,
    organization_id integer,
    CONSTRAINT ck_new_groups_pstv_2cb74b3445ff4f0c CHECK ((new_groups >= 0)),
    CONSTRAINT sentry_release_new_groups_check CHECK ((new_groups >= 0))
);


ALTER TABLE sentry_release OWNER TO sentry;

--
-- Name: sentry_release_id_seq; Type: SEQUENCE; Schema: public; Owner: sentry
--

CREATE SEQUENCE sentry_release_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE sentry_release_id_seq OWNER TO sentry;

--
-- Name: sentry_release_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: sentry
--

ALTER SEQUENCE sentry_release_id_seq OWNED BY sentry_release.id;


--
-- Name: sentry_release_project; Type: TABLE; Schema: public; Owner: sentry
--

CREATE TABLE sentry_release_project (
    id integer NOT NULL,
    project_id integer NOT NULL,
    release_id integer NOT NULL
);


ALTER TABLE sentry_release_project OWNER TO sentry;

--
-- Name: sentry_release_project_id_seq; Type: SEQUENCE; Schema: public; Owner: sentry
--

CREATE SEQUENCE sentry_release_project_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE sentry_release_project_id_seq OWNER TO sentry;

--
-- Name: sentry_release_project_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: sentry
--

ALTER SEQUENCE sentry_release_project_id_seq OWNED BY sentry_release_project.id;


--
-- Name: sentry_releasecommit; Type: TABLE; Schema: public; Owner: sentry
--

CREATE TABLE sentry_releasecommit (
    id integer NOT NULL,
    project_id integer NOT NULL,
    release_id integer NOT NULL,
    commit_id integer NOT NULL,
    "order" integer NOT NULL,
    organization_id integer,
    CONSTRAINT ck_organization_id_pstv_63c72b7b5009246 CHECK ((organization_id >= 0)),
    CONSTRAINT sentry_releasecommit_order_check CHECK (("order" >= 0)),
    CONSTRAINT sentry_releasecommit_organization_id_check CHECK ((organization_id >= 0)),
    CONSTRAINT sentry_releasecommit_project_id_check CHECK ((project_id >= 0))
);


ALTER TABLE sentry_releasecommit OWNER TO sentry;

--
-- Name: sentry_releasecommit_id_seq; Type: SEQUENCE; Schema: public; Owner: sentry
--

CREATE SEQUENCE sentry_releasecommit_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE sentry_releasecommit_id_seq OWNER TO sentry;

--
-- Name: sentry_releasecommit_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: sentry
--

ALTER SEQUENCE sentry_releasecommit_id_seq OWNED BY sentry_releasecommit.id;


--
-- Name: sentry_releasefile; Type: TABLE; Schema: public; Owner: sentry
--

CREATE TABLE sentry_releasefile (
    id integer NOT NULL,
    project_id integer NOT NULL,
    release_id integer NOT NULL,
    file_id integer NOT NULL,
    ident character varying(40) NOT NULL,
    name text NOT NULL,
    organization_id integer
);


ALTER TABLE sentry_releasefile OWNER TO sentry;

--
-- Name: sentry_releasefile_id_seq; Type: SEQUENCE; Schema: public; Owner: sentry
--

CREATE SEQUENCE sentry_releasefile_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE sentry_releasefile_id_seq OWNER TO sentry;

--
-- Name: sentry_releasefile_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: sentry
--

ALTER SEQUENCE sentry_releasefile_id_seq OWNED BY sentry_releasefile.id;


--
-- Name: sentry_repository; Type: TABLE; Schema: public; Owner: sentry
--

CREATE TABLE sentry_repository (
    id integer NOT NULL,
    organization_id integer NOT NULL,
    name character varying(200) NOT NULL,
    date_added timestamp with time zone NOT NULL,
    url character varying(200),
    provider character varying(64),
    external_id character varying(64),
    config text NOT NULL,
    status integer NOT NULL,
    CONSTRAINT ck_status_pstv_562b3ff4dae47f6b CHECK ((status >= 0)),
    CONSTRAINT sentry_repository_organization_id_check CHECK ((organization_id >= 0)),
    CONSTRAINT sentry_repository_status_check CHECK ((status >= 0))
);


ALTER TABLE sentry_repository OWNER TO sentry;

--
-- Name: sentry_repository_id_seq; Type: SEQUENCE; Schema: public; Owner: sentry
--

CREATE SEQUENCE sentry_repository_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE sentry_repository_id_seq OWNER TO sentry;

--
-- Name: sentry_repository_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: sentry
--

ALTER SEQUENCE sentry_repository_id_seq OWNED BY sentry_repository.id;


--
-- Name: sentry_rule; Type: TABLE; Schema: public; Owner: sentry
--

CREATE TABLE sentry_rule (
    id integer NOT NULL,
    project_id integer NOT NULL,
    label character varying(64) NOT NULL,
    data text NOT NULL,
    date_added timestamp with time zone NOT NULL,
    status integer NOT NULL,
    CONSTRAINT ck_status_pstv_64efa876e92cb76d CHECK ((status >= 0)),
    CONSTRAINT sentry_rule_status_check CHECK ((status >= 0))
);


ALTER TABLE sentry_rule OWNER TO sentry;

--
-- Name: sentry_rule_id_seq; Type: SEQUENCE; Schema: public; Owner: sentry
--

CREATE SEQUENCE sentry_rule_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE sentry_rule_id_seq OWNER TO sentry;

--
-- Name: sentry_rule_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: sentry
--

ALTER SEQUENCE sentry_rule_id_seq OWNED BY sentry_rule.id;


--
-- Name: sentry_savedsearch; Type: TABLE; Schema: public; Owner: sentry
--

CREATE TABLE sentry_savedsearch (
    id integer NOT NULL,
    project_id integer NOT NULL,
    name character varying(128) NOT NULL,
    query text NOT NULL,
    date_added timestamp with time zone NOT NULL,
    is_default boolean NOT NULL
);


ALTER TABLE sentry_savedsearch OWNER TO sentry;

--
-- Name: sentry_savedsearch_id_seq; Type: SEQUENCE; Schema: public; Owner: sentry
--

CREATE SEQUENCE sentry_savedsearch_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE sentry_savedsearch_id_seq OWNER TO sentry;

--
-- Name: sentry_savedsearch_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: sentry
--

ALTER SEQUENCE sentry_savedsearch_id_seq OWNED BY sentry_savedsearch.id;


--
-- Name: sentry_savedsearch_userdefault; Type: TABLE; Schema: public; Owner: sentry
--

CREATE TABLE sentry_savedsearch_userdefault (
    id integer NOT NULL,
    savedsearch_id integer NOT NULL,
    project_id integer NOT NULL,
    user_id integer NOT NULL
);


ALTER TABLE sentry_savedsearch_userdefault OWNER TO sentry;

--
-- Name: sentry_savedsearch_userdefault_id_seq; Type: SEQUENCE; Schema: public; Owner: sentry
--

CREATE SEQUENCE sentry_savedsearch_userdefault_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE sentry_savedsearch_userdefault_id_seq OWNER TO sentry;

--
-- Name: sentry_savedsearch_userdefault_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: sentry
--

ALTER SEQUENCE sentry_savedsearch_userdefault_id_seq OWNED BY sentry_savedsearch_userdefault.id;


--
-- Name: sentry_team; Type: TABLE; Schema: public; Owner: sentry
--

CREATE TABLE sentry_team (
    id integer NOT NULL,
    slug character varying(50) NOT NULL,
    name character varying(64) NOT NULL,
    date_added timestamp with time zone,
    status integer NOT NULL,
    organization_id integer NOT NULL,
    CONSTRAINT ck_status_pstv_1772e42d30eba7ba CHECK ((status >= 0)),
    CONSTRAINT sentry_team_status_check CHECK ((status >= 0))
);


ALTER TABLE sentry_team OWNER TO sentry;

--
-- Name: sentry_team_id_seq; Type: SEQUENCE; Schema: public; Owner: sentry
--

CREATE SEQUENCE sentry_team_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE sentry_team_id_seq OWNER TO sentry;

--
-- Name: sentry_team_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: sentry
--

ALTER SEQUENCE sentry_team_id_seq OWNED BY sentry_team.id;


--
-- Name: sentry_useravatar; Type: TABLE; Schema: public; Owner: sentry
--

CREATE TABLE sentry_useravatar (
    id integer NOT NULL,
    user_id integer NOT NULL,
    file_id integer,
    ident character varying(32) NOT NULL,
    avatar_type smallint NOT NULL,
    CONSTRAINT sentry_useravatar_avatar_type_check CHECK ((avatar_type >= 0))
);


ALTER TABLE sentry_useravatar OWNER TO sentry;

--
-- Name: sentry_useravatar_id_seq; Type: SEQUENCE; Schema: public; Owner: sentry
--

CREATE SEQUENCE sentry_useravatar_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE sentry_useravatar_id_seq OWNER TO sentry;

--
-- Name: sentry_useravatar_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: sentry
--

ALTER SEQUENCE sentry_useravatar_id_seq OWNED BY sentry_useravatar.id;


--
-- Name: sentry_useremail; Type: TABLE; Schema: public; Owner: sentry
--

CREATE TABLE sentry_useremail (
    id integer NOT NULL,
    user_id integer NOT NULL,
    email character varying(75) NOT NULL,
    validation_hash character varying(32) NOT NULL,
    date_hash_added timestamp with time zone NOT NULL,
    is_verified boolean NOT NULL
);


ALTER TABLE sentry_useremail OWNER TO sentry;

--
-- Name: sentry_useremail_id_seq; Type: SEQUENCE; Schema: public; Owner: sentry
--

CREATE SEQUENCE sentry_useremail_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE sentry_useremail_id_seq OWNER TO sentry;

--
-- Name: sentry_useremail_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: sentry
--

ALTER SEQUENCE sentry_useremail_id_seq OWNED BY sentry_useremail.id;


--
-- Name: sentry_useroption; Type: TABLE; Schema: public; Owner: sentry
--

CREATE TABLE sentry_useroption (
    id integer NOT NULL,
    user_id integer NOT NULL,
    project_id integer,
    key character varying(64) NOT NULL,
    value text NOT NULL
);


ALTER TABLE sentry_useroption OWNER TO sentry;

--
-- Name: sentry_useroption_id_seq; Type: SEQUENCE; Schema: public; Owner: sentry
--

CREATE SEQUENCE sentry_useroption_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE sentry_useroption_id_seq OWNER TO sentry;

--
-- Name: sentry_useroption_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: sentry
--

ALTER SEQUENCE sentry_useroption_id_seq OWNED BY sentry_useroption.id;


--
-- Name: sentry_userreport; Type: TABLE; Schema: public; Owner: sentry
--

CREATE TABLE sentry_userreport (
    id integer NOT NULL,
    project_id integer NOT NULL,
    group_id integer,
    event_id character varying(32) NOT NULL,
    name character varying(128) NOT NULL,
    email character varying(75) NOT NULL,
    comments text NOT NULL,
    date_added timestamp with time zone NOT NULL
);


ALTER TABLE sentry_userreport OWNER TO sentry;

--
-- Name: sentry_userreport_id_seq; Type: SEQUENCE; Schema: public; Owner: sentry
--

CREATE SEQUENCE sentry_userreport_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE sentry_userreport_id_seq OWNER TO sentry;

--
-- Name: sentry_userreport_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: sentry
--

ALTER SEQUENCE sentry_userreport_id_seq OWNED BY sentry_userreport.id;


--
-- Name: social_auth_usersocialauth; Type: TABLE; Schema: public; Owner: sentry
--

CREATE TABLE social_auth_usersocialauth (
    id integer NOT NULL,
    user_id integer NOT NULL,
    provider character varying(32) NOT NULL,
    uid character varying(255) NOT NULL,
    extra_data text NOT NULL
);


ALTER TABLE social_auth_usersocialauth OWNER TO sentry;

--
-- Name: social_auth_usersocialauth_id_seq; Type: SEQUENCE; Schema: public; Owner: sentry
--

CREATE SEQUENCE social_auth_usersocialauth_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE social_auth_usersocialauth_id_seq OWNER TO sentry;

--
-- Name: social_auth_usersocialauth_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: sentry
--

ALTER SEQUENCE social_auth_usersocialauth_id_seq OWNED BY social_auth_usersocialauth.id;


--
-- Name: south_migrationhistory; Type: TABLE; Schema: public; Owner: sentry
--

CREATE TABLE south_migrationhistory (
    id integer NOT NULL,
    app_name character varying(255) NOT NULL,
    migration character varying(255) NOT NULL,
    applied timestamp with time zone NOT NULL
);


ALTER TABLE south_migrationhistory OWNER TO sentry;

--
-- Name: south_migrationhistory_id_seq; Type: SEQUENCE; Schema: public; Owner: sentry
--

CREATE SEQUENCE south_migrationhistory_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE south_migrationhistory_id_seq OWNER TO sentry;

--
-- Name: south_migrationhistory_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: sentry
--

ALTER SEQUENCE south_migrationhistory_id_seq OWNED BY south_migrationhistory.id;


--
-- Name: auth_authenticator id; Type: DEFAULT; Schema: public; Owner: sentry
--

ALTER TABLE ONLY auth_authenticator ALTER COLUMN id SET DEFAULT nextval('auth_authenticator_id_seq'::regclass);


--
-- Name: auth_group id; Type: DEFAULT; Schema: public; Owner: sentry
--

ALTER TABLE ONLY auth_group ALTER COLUMN id SET DEFAULT nextval('auth_group_id_seq'::regclass);


--
-- Name: auth_group_permissions id; Type: DEFAULT; Schema: public; Owner: sentry
--

ALTER TABLE ONLY auth_group_permissions ALTER COLUMN id SET DEFAULT nextval('auth_group_permissions_id_seq'::regclass);


--
-- Name: auth_permission id; Type: DEFAULT; Schema: public; Owner: sentry
--

ALTER TABLE ONLY auth_permission ALTER COLUMN id SET DEFAULT nextval('auth_permission_id_seq'::regclass);


--
-- Name: auth_user id; Type: DEFAULT; Schema: public; Owner: sentry
--

ALTER TABLE ONLY auth_user ALTER COLUMN id SET DEFAULT nextval('auth_user_id_seq'::regclass);


--
-- Name: django_admin_log id; Type: DEFAULT; Schema: public; Owner: sentry
--

ALTER TABLE ONLY django_admin_log ALTER COLUMN id SET DEFAULT nextval('django_admin_log_id_seq'::regclass);


--
-- Name: django_content_type id; Type: DEFAULT; Schema: public; Owner: sentry
--

ALTER TABLE ONLY django_content_type ALTER COLUMN id SET DEFAULT nextval('django_content_type_id_seq'::regclass);


--
-- Name: django_site id; Type: DEFAULT; Schema: public; Owner: sentry
--

ALTER TABLE ONLY django_site ALTER COLUMN id SET DEFAULT nextval('django_site_id_seq'::regclass);


--
-- Name: sentry_activity id; Type: DEFAULT; Schema: public; Owner: sentry
--

ALTER TABLE ONLY sentry_activity ALTER COLUMN id SET DEFAULT nextval('sentry_activity_id_seq'::regclass);


--
-- Name: sentry_apikey id; Type: DEFAULT; Schema: public; Owner: sentry
--

ALTER TABLE ONLY sentry_apikey ALTER COLUMN id SET DEFAULT nextval('sentry_apikey_id_seq'::regclass);


--
-- Name: sentry_apitoken id; Type: DEFAULT; Schema: public; Owner: sentry
--

ALTER TABLE ONLY sentry_apitoken ALTER COLUMN id SET DEFAULT nextval('sentry_apitoken_id_seq'::regclass);


--
-- Name: sentry_auditlogentry id; Type: DEFAULT; Schema: public; Owner: sentry
--

ALTER TABLE ONLY sentry_auditlogentry ALTER COLUMN id SET DEFAULT nextval('sentry_auditlogentry_id_seq'::regclass);


--
-- Name: sentry_authidentity id; Type: DEFAULT; Schema: public; Owner: sentry
--

ALTER TABLE ONLY sentry_authidentity ALTER COLUMN id SET DEFAULT nextval('sentry_authidentity_id_seq'::regclass);


--
-- Name: sentry_authprovider id; Type: DEFAULT; Schema: public; Owner: sentry
--

ALTER TABLE ONLY sentry_authprovider ALTER COLUMN id SET DEFAULT nextval('sentry_authprovider_id_seq'::regclass);


--
-- Name: sentry_authprovider_default_teams id; Type: DEFAULT; Schema: public; Owner: sentry
--

ALTER TABLE ONLY sentry_authprovider_default_teams ALTER COLUMN id SET DEFAULT nextval('sentry_authprovider_default_teams_id_seq'::regclass);


--
-- Name: sentry_broadcast id; Type: DEFAULT; Schema: public; Owner: sentry
--

ALTER TABLE ONLY sentry_broadcast ALTER COLUMN id SET DEFAULT nextval('sentry_broadcast_id_seq'::regclass);


--
-- Name: sentry_broadcastseen id; Type: DEFAULT; Schema: public; Owner: sentry
--

ALTER TABLE ONLY sentry_broadcastseen ALTER COLUMN id SET DEFAULT nextval('sentry_broadcastseen_id_seq'::regclass);


--
-- Name: sentry_commit id; Type: DEFAULT; Schema: public; Owner: sentry
--

ALTER TABLE ONLY sentry_commit ALTER COLUMN id SET DEFAULT nextval('sentry_commit_id_seq'::regclass);


--
-- Name: sentry_commitauthor id; Type: DEFAULT; Schema: public; Owner: sentry
--

ALTER TABLE ONLY sentry_commitauthor ALTER COLUMN id SET DEFAULT nextval('sentry_commitauthor_id_seq'::regclass);


--
-- Name: sentry_commitfilechange id; Type: DEFAULT; Schema: public; Owner: sentry
--

ALTER TABLE ONLY sentry_commitfilechange ALTER COLUMN id SET DEFAULT nextval('sentry_commitfilechange_id_seq'::regclass);


--
-- Name: sentry_dsymbundle id; Type: DEFAULT; Schema: public; Owner: sentry
--

ALTER TABLE ONLY sentry_dsymbundle ALTER COLUMN id SET DEFAULT nextval('sentry_dsymbundle_id_seq'::regclass);


--
-- Name: sentry_dsymobject id; Type: DEFAULT; Schema: public; Owner: sentry
--

ALTER TABLE ONLY sentry_dsymobject ALTER COLUMN id SET DEFAULT nextval('sentry_dsymobject_id_seq'::regclass);


--
-- Name: sentry_dsymsdk id; Type: DEFAULT; Schema: public; Owner: sentry
--

ALTER TABLE ONLY sentry_dsymsdk ALTER COLUMN id SET DEFAULT nextval('sentry_dsymsdk_id_seq'::regclass);


--
-- Name: sentry_dsymsymbol id; Type: DEFAULT; Schema: public; Owner: sentry
--

ALTER TABLE ONLY sentry_dsymsymbol ALTER COLUMN id SET DEFAULT nextval('sentry_dsymsymbol_id_seq'::regclass);


--
-- Name: sentry_environment id; Type: DEFAULT; Schema: public; Owner: sentry
--

ALTER TABLE ONLY sentry_environment ALTER COLUMN id SET DEFAULT nextval('sentry_environment_id_seq'::regclass);


--
-- Name: sentry_environmentrelease id; Type: DEFAULT; Schema: public; Owner: sentry
--

ALTER TABLE ONLY sentry_environmentrelease ALTER COLUMN id SET DEFAULT nextval('sentry_environmentrelease_id_seq'::regclass);


--
-- Name: sentry_eventmapping id; Type: DEFAULT; Schema: public; Owner: sentry
--

ALTER TABLE ONLY sentry_eventmapping ALTER COLUMN id SET DEFAULT nextval('sentry_eventmapping_id_seq'::regclass);


--
-- Name: sentry_eventtag id; Type: DEFAULT; Schema: public; Owner: sentry
--

ALTER TABLE ONLY sentry_eventtag ALTER COLUMN id SET DEFAULT nextval('sentry_eventtag_id_seq'::regclass);


--
-- Name: sentry_eventuser id; Type: DEFAULT; Schema: public; Owner: sentry
--

ALTER TABLE ONLY sentry_eventuser ALTER COLUMN id SET DEFAULT nextval('sentry_eventuser_id_seq'::regclass);


--
-- Name: sentry_file id; Type: DEFAULT; Schema: public; Owner: sentry
--

ALTER TABLE ONLY sentry_file ALTER COLUMN id SET DEFAULT nextval('sentry_file_id_seq'::regclass);


--
-- Name: sentry_fileblob id; Type: DEFAULT; Schema: public; Owner: sentry
--

ALTER TABLE ONLY sentry_fileblob ALTER COLUMN id SET DEFAULT nextval('sentry_fileblob_id_seq'::regclass);


--
-- Name: sentry_fileblobindex id; Type: DEFAULT; Schema: public; Owner: sentry
--

ALTER TABLE ONLY sentry_fileblobindex ALTER COLUMN id SET DEFAULT nextval('sentry_fileblobindex_id_seq'::regclass);


--
-- Name: sentry_filterkey id; Type: DEFAULT; Schema: public; Owner: sentry
--

ALTER TABLE ONLY sentry_filterkey ALTER COLUMN id SET DEFAULT nextval('sentry_filterkey_id_seq'::regclass);


--
-- Name: sentry_filtervalue id; Type: DEFAULT; Schema: public; Owner: sentry
--

ALTER TABLE ONLY sentry_filtervalue ALTER COLUMN id SET DEFAULT nextval('sentry_filtervalue_id_seq'::regclass);


--
-- Name: sentry_globaldsymfile id; Type: DEFAULT; Schema: public; Owner: sentry
--

ALTER TABLE ONLY sentry_globaldsymfile ALTER COLUMN id SET DEFAULT nextval('sentry_globaldsymfile_id_seq'::regclass);


--
-- Name: sentry_groupasignee id; Type: DEFAULT; Schema: public; Owner: sentry
--

ALTER TABLE ONLY sentry_groupasignee ALTER COLUMN id SET DEFAULT nextval('sentry_groupasignee_id_seq'::regclass);


--
-- Name: sentry_groupbookmark id; Type: DEFAULT; Schema: public; Owner: sentry
--

ALTER TABLE ONLY sentry_groupbookmark ALTER COLUMN id SET DEFAULT nextval('sentry_groupbookmark_id_seq'::regclass);


--
-- Name: sentry_groupedmessage id; Type: DEFAULT; Schema: public; Owner: sentry
--

ALTER TABLE ONLY sentry_groupedmessage ALTER COLUMN id SET DEFAULT nextval('sentry_groupedmessage_id_seq'::regclass);


--
-- Name: sentry_groupemailthread id; Type: DEFAULT; Schema: public; Owner: sentry
--

ALTER TABLE ONLY sentry_groupemailthread ALTER COLUMN id SET DEFAULT nextval('sentry_groupemailthread_id_seq'::regclass);


--
-- Name: sentry_grouphash id; Type: DEFAULT; Schema: public; Owner: sentry
--

ALTER TABLE ONLY sentry_grouphash ALTER COLUMN id SET DEFAULT nextval('sentry_grouphash_id_seq'::regclass);


--
-- Name: sentry_groupmeta id; Type: DEFAULT; Schema: public; Owner: sentry
--

ALTER TABLE ONLY sentry_groupmeta ALTER COLUMN id SET DEFAULT nextval('sentry_groupmeta_id_seq'::regclass);


--
-- Name: sentry_groupredirect id; Type: DEFAULT; Schema: public; Owner: sentry
--

ALTER TABLE ONLY sentry_groupredirect ALTER COLUMN id SET DEFAULT nextval('sentry_groupredirect_id_seq'::regclass);


--
-- Name: sentry_grouprelease id; Type: DEFAULT; Schema: public; Owner: sentry
--

ALTER TABLE ONLY sentry_grouprelease ALTER COLUMN id SET DEFAULT nextval('sentry_grouprelease_id_seq'::regclass);


--
-- Name: sentry_groupresolution id; Type: DEFAULT; Schema: public; Owner: sentry
--

ALTER TABLE ONLY sentry_groupresolution ALTER COLUMN id SET DEFAULT nextval('sentry_groupresolution_id_seq'::regclass);


--
-- Name: sentry_grouprulestatus id; Type: DEFAULT; Schema: public; Owner: sentry
--

ALTER TABLE ONLY sentry_grouprulestatus ALTER COLUMN id SET DEFAULT nextval('sentry_grouprulestatus_id_seq'::regclass);


--
-- Name: sentry_groupseen id; Type: DEFAULT; Schema: public; Owner: sentry
--

ALTER TABLE ONLY sentry_groupseen ALTER COLUMN id SET DEFAULT nextval('sentry_groupseen_id_seq'::regclass);


--
-- Name: sentry_groupsnooze id; Type: DEFAULT; Schema: public; Owner: sentry
--

ALTER TABLE ONLY sentry_groupsnooze ALTER COLUMN id SET DEFAULT nextval('sentry_groupsnooze_id_seq'::regclass);


--
-- Name: sentry_groupsubscription id; Type: DEFAULT; Schema: public; Owner: sentry
--

ALTER TABLE ONLY sentry_groupsubscription ALTER COLUMN id SET DEFAULT nextval('sentry_groupsubscription_id_seq'::regclass);


--
-- Name: sentry_grouptagkey id; Type: DEFAULT; Schema: public; Owner: sentry
--

ALTER TABLE ONLY sentry_grouptagkey ALTER COLUMN id SET DEFAULT nextval('sentry_grouptagkey_id_seq'::regclass);


--
-- Name: sentry_lostpasswordhash id; Type: DEFAULT; Schema: public; Owner: sentry
--

ALTER TABLE ONLY sentry_lostpasswordhash ALTER COLUMN id SET DEFAULT nextval('sentry_lostpasswordhash_id_seq'::regclass);


--
-- Name: sentry_message id; Type: DEFAULT; Schema: public; Owner: sentry
--

ALTER TABLE ONLY sentry_message ALTER COLUMN id SET DEFAULT nextval('sentry_message_id_seq'::regclass);


--
-- Name: sentry_messagefiltervalue id; Type: DEFAULT; Schema: public; Owner: sentry
--

ALTER TABLE ONLY sentry_messagefiltervalue ALTER COLUMN id SET DEFAULT nextval('sentry_messagefiltervalue_id_seq'::regclass);


--
-- Name: sentry_messageindex id; Type: DEFAULT; Schema: public; Owner: sentry
--

ALTER TABLE ONLY sentry_messageindex ALTER COLUMN id SET DEFAULT nextval('sentry_messageindex_id_seq'::regclass);


--
-- Name: sentry_option id; Type: DEFAULT; Schema: public; Owner: sentry
--

ALTER TABLE ONLY sentry_option ALTER COLUMN id SET DEFAULT nextval('sentry_option_id_seq'::regclass);


--
-- Name: sentry_organization id; Type: DEFAULT; Schema: public; Owner: sentry
--

ALTER TABLE ONLY sentry_organization ALTER COLUMN id SET DEFAULT nextval('sentry_organization_id_seq'::regclass);


--
-- Name: sentry_organizationaccessrequest id; Type: DEFAULT; Schema: public; Owner: sentry
--

ALTER TABLE ONLY sentry_organizationaccessrequest ALTER COLUMN id SET DEFAULT nextval('sentry_organizationaccessrequest_id_seq'::regclass);


--
-- Name: sentry_organizationmember id; Type: DEFAULT; Schema: public; Owner: sentry
--

ALTER TABLE ONLY sentry_organizationmember ALTER COLUMN id SET DEFAULT nextval('sentry_organizationmember_id_seq'::regclass);


--
-- Name: sentry_organizationmember_teams id; Type: DEFAULT; Schema: public; Owner: sentry
--

ALTER TABLE ONLY sentry_organizationmember_teams ALTER COLUMN id SET DEFAULT nextval('sentry_organizationmember_teams_id_seq'::regclass);


--
-- Name: sentry_organizationonboardingtask id; Type: DEFAULT; Schema: public; Owner: sentry
--

ALTER TABLE ONLY sentry_organizationonboardingtask ALTER COLUMN id SET DEFAULT nextval('sentry_organizationonboardingtask_id_seq'::regclass);


--
-- Name: sentry_organizationoptions id; Type: DEFAULT; Schema: public; Owner: sentry
--

ALTER TABLE ONLY sentry_organizationoptions ALTER COLUMN id SET DEFAULT nextval('sentry_organizationoptions_id_seq'::regclass);


--
-- Name: sentry_project id; Type: DEFAULT; Schema: public; Owner: sentry
--

ALTER TABLE ONLY sentry_project ALTER COLUMN id SET DEFAULT nextval('sentry_project_id_seq'::regclass);


--
-- Name: sentry_projectbookmark id; Type: DEFAULT; Schema: public; Owner: sentry
--

ALTER TABLE ONLY sentry_projectbookmark ALTER COLUMN id SET DEFAULT nextval('sentry_projectbookmark_id_seq'::regclass);


--
-- Name: sentry_projectcounter id; Type: DEFAULT; Schema: public; Owner: sentry
--

ALTER TABLE ONLY sentry_projectcounter ALTER COLUMN id SET DEFAULT nextval('sentry_projectcounter_id_seq'::regclass);


--
-- Name: sentry_projectdsymfile id; Type: DEFAULT; Schema: public; Owner: sentry
--

ALTER TABLE ONLY sentry_projectdsymfile ALTER COLUMN id SET DEFAULT nextval('sentry_projectdsymfile_id_seq'::regclass);


--
-- Name: sentry_projectkey id; Type: DEFAULT; Schema: public; Owner: sentry
--

ALTER TABLE ONLY sentry_projectkey ALTER COLUMN id SET DEFAULT nextval('sentry_projectkey_id_seq'::regclass);


--
-- Name: sentry_projectoptions id; Type: DEFAULT; Schema: public; Owner: sentry
--

ALTER TABLE ONLY sentry_projectoptions ALTER COLUMN id SET DEFAULT nextval('sentry_projectoptions_id_seq'::regclass);


--
-- Name: sentry_projectplatform id; Type: DEFAULT; Schema: public; Owner: sentry
--

ALTER TABLE ONLY sentry_projectplatform ALTER COLUMN id SET DEFAULT nextval('sentry_projectplatform_id_seq'::regclass);


--
-- Name: sentry_release id; Type: DEFAULT; Schema: public; Owner: sentry
--

ALTER TABLE ONLY sentry_release ALTER COLUMN id SET DEFAULT nextval('sentry_release_id_seq'::regclass);


--
-- Name: sentry_release_project id; Type: DEFAULT; Schema: public; Owner: sentry
--

ALTER TABLE ONLY sentry_release_project ALTER COLUMN id SET DEFAULT nextval('sentry_release_project_id_seq'::regclass);


--
-- Name: sentry_releasecommit id; Type: DEFAULT; Schema: public; Owner: sentry
--

ALTER TABLE ONLY sentry_releasecommit ALTER COLUMN id SET DEFAULT nextval('sentry_releasecommit_id_seq'::regclass);


--
-- Name: sentry_releasefile id; Type: DEFAULT; Schema: public; Owner: sentry
--

ALTER TABLE ONLY sentry_releasefile ALTER COLUMN id SET DEFAULT nextval('sentry_releasefile_id_seq'::regclass);


--
-- Name: sentry_repository id; Type: DEFAULT; Schema: public; Owner: sentry
--

ALTER TABLE ONLY sentry_repository ALTER COLUMN id SET DEFAULT nextval('sentry_repository_id_seq'::regclass);


--
-- Name: sentry_rule id; Type: DEFAULT; Schema: public; Owner: sentry
--

ALTER TABLE ONLY sentry_rule ALTER COLUMN id SET DEFAULT nextval('sentry_rule_id_seq'::regclass);


--
-- Name: sentry_savedsearch id; Type: DEFAULT; Schema: public; Owner: sentry
--

ALTER TABLE ONLY sentry_savedsearch ALTER COLUMN id SET DEFAULT nextval('sentry_savedsearch_id_seq'::regclass);


--
-- Name: sentry_savedsearch_userdefault id; Type: DEFAULT; Schema: public; Owner: sentry
--

ALTER TABLE ONLY sentry_savedsearch_userdefault ALTER COLUMN id SET DEFAULT nextval('sentry_savedsearch_userdefault_id_seq'::regclass);


--
-- Name: sentry_team id; Type: DEFAULT; Schema: public; Owner: sentry
--

ALTER TABLE ONLY sentry_team ALTER COLUMN id SET DEFAULT nextval('sentry_team_id_seq'::regclass);


--
-- Name: sentry_useravatar id; Type: DEFAULT; Schema: public; Owner: sentry
--

ALTER TABLE ONLY sentry_useravatar ALTER COLUMN id SET DEFAULT nextval('sentry_useravatar_id_seq'::regclass);


--
-- Name: sentry_useremail id; Type: DEFAULT; Schema: public; Owner: sentry
--

ALTER TABLE ONLY sentry_useremail ALTER COLUMN id SET DEFAULT nextval('sentry_useremail_id_seq'::regclass);


--
-- Name: sentry_useroption id; Type: DEFAULT; Schema: public; Owner: sentry
--

ALTER TABLE ONLY sentry_useroption ALTER COLUMN id SET DEFAULT nextval('sentry_useroption_id_seq'::regclass);


--
-- Name: sentry_userreport id; Type: DEFAULT; Schema: public; Owner: sentry
--

ALTER TABLE ONLY sentry_userreport ALTER COLUMN id SET DEFAULT nextval('sentry_userreport_id_seq'::regclass);


--
-- Name: social_auth_usersocialauth id; Type: DEFAULT; Schema: public; Owner: sentry
--

ALTER TABLE ONLY social_auth_usersocialauth ALTER COLUMN id SET DEFAULT nextval('social_auth_usersocialauth_id_seq'::regclass);


--
-- Name: south_migrationhistory id; Type: DEFAULT; Schema: public; Owner: sentry
--

ALTER TABLE ONLY south_migrationhistory ALTER COLUMN id SET DEFAULT nextval('south_migrationhistory_id_seq'::regclass);


--
-- Data for Name: auth_authenticator; Type: TABLE DATA; Schema: public; Owner: sentry
--

COPY auth_authenticator (id, user_id, created_at, last_used_at, type, config) FROM stdin;
\.


--
-- Name: auth_authenticator_id_seq; Type: SEQUENCE SET; Schema: public; Owner: sentry
--

SELECT pg_catalog.setval('auth_authenticator_id_seq', 1, false);


--
-- Data for Name: auth_group; Type: TABLE DATA; Schema: public; Owner: sentry
--

COPY auth_group (id, name) FROM stdin;
\.


--
-- Name: auth_group_id_seq; Type: SEQUENCE SET; Schema: public; Owner: sentry
--

SELECT pg_catalog.setval('auth_group_id_seq', 1, false);


--
-- Data for Name: auth_group_permissions; Type: TABLE DATA; Schema: public; Owner: sentry
--

COPY auth_group_permissions (id, group_id, permission_id) FROM stdin;
\.


--
-- Name: auth_group_permissions_id_seq; Type: SEQUENCE SET; Schema: public; Owner: sentry
--

SELECT pg_catalog.setval('auth_group_permissions_id_seq', 1, false);


--
-- Data for Name: auth_permission; Type: TABLE DATA; Schema: public; Owner: sentry
--

COPY auth_permission (id, name, content_type_id, codename) FROM stdin;
1	Can add log entry	1	add_logentry
2	Can change log entry	1	change_logentry
3	Can delete log entry	1	delete_logentry
4	Can add permission	2	add_permission
5	Can change permission	2	change_permission
6	Can delete permission	2	delete_permission
7	Can add group	3	add_group
8	Can change group	3	change_group
9	Can delete group	3	delete_group
10	Can add content type	4	add_contenttype
11	Can change content type	4	change_contenttype
12	Can delete content type	4	delete_contenttype
13	Can add session	5	add_session
14	Can change session	5	change_session
15	Can delete session	5	delete_session
16	Can add site	6	add_site
17	Can change site	6	change_site
18	Can delete site	6	delete_site
19	Can add migration history	7	add_migrationhistory
20	Can change migration history	7	change_migrationhistory
21	Can delete migration history	7	delete_migrationhistory
22	Can add activity	8	add_activity
23	Can change activity	8	change_activity
24	Can delete activity	8	delete_activity
25	Can add api key	9	add_apikey
26	Can change api key	9	change_apikey
27	Can delete api key	9	delete_apikey
28	Can add api token	10	add_apitoken
29	Can change api token	10	change_apitoken
30	Can delete api token	10	delete_apitoken
31	Can add audit log entry	11	add_auditlogentry
32	Can change audit log entry	11	change_auditlogentry
33	Can delete audit log entry	11	delete_auditlogentry
34	Can add authenticator	12	add_authenticator
35	Can change authenticator	12	change_authenticator
36	Can delete authenticator	12	delete_authenticator
37	Can add auth identity	13	add_authidentity
38	Can change auth identity	13	change_authidentity
39	Can delete auth identity	13	delete_authidentity
40	Can add auth provider	14	add_authprovider
41	Can change auth provider	14	change_authprovider
42	Can delete auth provider	14	delete_authprovider
43	Can add broadcast	15	add_broadcast
44	Can change broadcast	15	change_broadcast
45	Can delete broadcast	15	delete_broadcast
46	Can add broadcast seen	16	add_broadcastseen
47	Can change broadcast seen	16	change_broadcastseen
48	Can delete broadcast seen	16	delete_broadcastseen
49	Can add commit	17	add_commit
50	Can change commit	17	change_commit
51	Can delete commit	17	delete_commit
52	Can add commit author	18	add_commitauthor
53	Can change commit author	18	change_commitauthor
54	Can delete commit author	18	delete_commitauthor
55	Can add commit file change	19	add_commitfilechange
56	Can change commit file change	19	change_commitfilechange
57	Can delete commit file change	19	delete_commitfilechange
58	Can add counter	20	add_counter
59	Can change counter	20	change_counter
60	Can delete counter	20	delete_counter
61	Can add file blob	21	add_fileblob
62	Can change file blob	21	change_fileblob
63	Can delete file blob	21	delete_fileblob
64	Can add file	22	add_file
65	Can change file	22	change_file
66	Can delete file	22	delete_file
67	Can add file blob index	23	add_fileblobindex
68	Can change file blob index	23	change_fileblobindex
69	Can delete file blob index	23	delete_fileblobindex
70	Can add d sym sdk	24	add_dsymsdk
71	Can change d sym sdk	24	change_dsymsdk
72	Can delete d sym sdk	24	delete_dsymsdk
73	Can add d sym object	25	add_dsymobject
74	Can change d sym object	25	change_dsymobject
75	Can delete d sym object	25	delete_dsymobject
76	Can add d sym bundle	26	add_dsymbundle
77	Can change d sym bundle	26	change_dsymbundle
78	Can delete d sym bundle	26	delete_dsymbundle
79	Can add d sym symbol	27	add_dsymsymbol
80	Can change d sym symbol	27	change_dsymsymbol
81	Can delete d sym symbol	27	delete_dsymsymbol
82	Can add project d sym file	28	add_projectdsymfile
83	Can change project d sym file	28	change_projectdsymfile
84	Can delete project d sym file	28	delete_projectdsymfile
85	Can add global d sym file	29	add_globaldsymfile
86	Can change global d sym file	29	change_globaldsymfile
87	Can delete global d sym file	29	delete_globaldsymfile
88	Can add environment	30	add_environment
89	Can change environment	30	change_environment
90	Can delete environment	30	delete_environment
91	Can add message	31	add_event
92	Can change message	31	change_event
93	Can delete message	31	delete_event
94	Can add event mapping	32	add_eventmapping
95	Can change event mapping	32	change_eventmapping
96	Can delete event mapping	32	delete_eventmapping
97	Can add event tag	33	add_eventtag
98	Can change event tag	33	change_eventtag
99	Can delete event tag	33	delete_eventtag
100	Can add event user	34	add_eventuser
101	Can change event user	34	change_eventuser
102	Can delete event user	34	delete_eventuser
103	Can add grouped message	35	add_group
104	Can change grouped message	35	change_group
105	Can delete grouped message	35	delete_group
106	Can view	35	can_view
107	Can add group assignee	36	add_groupassignee
108	Can change group assignee	36	change_groupassignee
109	Can delete group assignee	36	delete_groupassignee
110	Can add group bookmark	37	add_groupbookmark
111	Can change group bookmark	37	change_groupbookmark
112	Can delete group bookmark	37	delete_groupbookmark
113	Can add group email thread	38	add_groupemailthread
114	Can change group email thread	38	change_groupemailthread
115	Can delete group email thread	38	delete_groupemailthread
116	Can add group hash	39	add_grouphash
117	Can change group hash	39	change_grouphash
118	Can delete group hash	39	delete_grouphash
119	Can add group meta	40	add_groupmeta
120	Can change group meta	40	change_groupmeta
121	Can delete group meta	40	delete_groupmeta
122	Can add group redirect	41	add_groupredirect
123	Can change group redirect	41	change_groupredirect
124	Can delete group redirect	41	delete_groupredirect
125	Can add group release	42	add_grouprelease
126	Can change group release	42	change_grouprelease
127	Can delete group release	42	delete_grouprelease
128	Can add group resolution	43	add_groupresolution
129	Can change group resolution	43	change_groupresolution
130	Can delete group resolution	43	delete_groupresolution
131	Can add group rule status	44	add_grouprulestatus
132	Can change group rule status	44	change_grouprulestatus
133	Can delete group rule status	44	delete_grouprulestatus
134	Can add group seen	45	add_groupseen
135	Can change group seen	45	change_groupseen
136	Can delete group seen	45	delete_groupseen
137	Can add group snooze	46	add_groupsnooze
138	Can change group snooze	46	change_groupsnooze
139	Can delete group snooze	46	delete_groupsnooze
140	Can add group subscription	47	add_groupsubscription
141	Can change group subscription	47	change_groupsubscription
142	Can delete group subscription	47	delete_groupsubscription
143	Can add group tag key	48	add_grouptagkey
144	Can change group tag key	48	change_grouptagkey
145	Can delete group tag key	48	delete_grouptagkey
146	Can add group tag value	49	add_grouptagvalue
147	Can change group tag value	49	change_grouptagvalue
148	Can delete group tag value	49	delete_grouptagvalue
149	Can add lost password hash	50	add_lostpasswordhash
150	Can change lost password hash	50	change_lostpasswordhash
151	Can delete lost password hash	50	delete_lostpasswordhash
152	Can add option	51	add_option
153	Can change option	51	change_option
154	Can delete option	51	delete_option
155	Can add organization	52	add_organization
156	Can change organization	52	change_organization
157	Can delete organization	52	delete_organization
158	Can add organization access request	53	add_organizationaccessrequest
159	Can change organization access request	53	change_organizationaccessrequest
160	Can delete organization access request	53	delete_organizationaccessrequest
161	Can add organization member team	54	add_organizationmemberteam
162	Can change organization member team	54	change_organizationmemberteam
163	Can delete organization member team	54	delete_organizationmemberteam
164	Can add organization member	55	add_organizationmember
165	Can change organization member	55	change_organizationmember
166	Can delete organization member	55	delete_organizationmember
167	Can add organization onboarding task	56	add_organizationonboardingtask
168	Can change organization onboarding task	56	change_organizationonboardingtask
169	Can delete organization onboarding task	56	delete_organizationonboardingtask
170	Can add organization option	57	add_organizationoption
171	Can change organization option	57	change_organizationoption
172	Can delete organization option	57	delete_organizationoption
173	Can add project	58	add_project
174	Can change project	58	change_project
175	Can delete project	58	delete_project
176	Can add project bookmark	59	add_projectbookmark
177	Can change project bookmark	59	change_projectbookmark
178	Can delete project bookmark	59	delete_projectbookmark
179	Can add project key	60	add_projectkey
180	Can change project key	60	change_projectkey
181	Can delete project key	60	delete_projectkey
182	Can add project option	61	add_projectoption
183	Can change project option	61	change_projectoption
184	Can delete project option	61	delete_projectoption
185	Can add project platform	62	add_projectplatform
186	Can change project platform	62	change_projectplatform
187	Can delete project platform	62	delete_projectplatform
188	Can add release project	63	add_releaseproject
189	Can change release project	63	change_releaseproject
190	Can delete release project	63	delete_releaseproject
191	Can add release	64	add_release
192	Can change release	64	change_release
193	Can delete release	64	delete_release
194	Can add release commit	65	add_releasecommit
195	Can change release commit	65	change_releasecommit
196	Can delete release commit	65	delete_releasecommit
197	Can add release environment	66	add_releaseenvironment
198	Can change release environment	66	change_releaseenvironment
199	Can delete release environment	66	delete_releaseenvironment
200	Can add release file	67	add_releasefile
201	Can change release file	67	change_releasefile
202	Can delete release file	67	delete_releasefile
203	Can add repository	68	add_repository
204	Can change repository	68	change_repository
205	Can delete repository	68	delete_repository
206	Can add rule	69	add_rule
207	Can change rule	69	change_rule
208	Can delete rule	69	delete_rule
209	Can add saved search	70	add_savedsearch
210	Can change saved search	70	change_savedsearch
211	Can delete saved search	70	delete_savedsearch
212	Can add saved search user default	71	add_savedsearchuserdefault
213	Can change saved search user default	71	change_savedsearchuserdefault
214	Can delete saved search user default	71	delete_savedsearchuserdefault
215	Can add tag key	72	add_tagkey
216	Can change tag key	72	change_tagkey
217	Can delete tag key	72	delete_tagkey
218	Can add tag value	73	add_tagvalue
219	Can change tag value	73	change_tagvalue
220	Can delete tag value	73	delete_tagvalue
221	Can add team	74	add_team
222	Can change team	74	change_team
223	Can delete team	74	delete_team
224	Can add user	75	add_user
225	Can change user	75	change_user
226	Can delete user	75	delete_user
227	Can add user avatar	76	add_useravatar
228	Can change user avatar	76	change_useravatar
229	Can delete user avatar	76	delete_useravatar
230	Can add user email	77	add_useremail
231	Can change user email	77	change_useremail
232	Can delete user email	77	delete_useremail
233	Can add user option	78	add_useroption
234	Can change user option	78	change_useroption
235	Can delete user option	78	delete_useroption
236	Can add user report	79	add_userreport
237	Can change user report	79	change_userreport
238	Can delete user report	79	delete_userreport
239	Can add node	80	add_node
240	Can change node	80	change_node
241	Can delete node	80	delete_node
242	Can add user social auth	81	add_usersocialauth
243	Can change user social auth	81	change_usersocialauth
244	Can delete user social auth	81	delete_usersocialauth
\.


--
-- Name: auth_permission_id_seq; Type: SEQUENCE SET; Schema: public; Owner: sentry
--

SELECT pg_catalog.setval('auth_permission_id_seq', 244, true);


--
-- Data for Name: auth_user; Type: TABLE DATA; Schema: public; Owner: sentry
--

COPY auth_user (password, last_login, id, username, first_name, email, is_staff, is_active, is_superuser, date_joined, is_managed, is_password_expired, last_password_change, session_nonce) FROM stdin;
pbkdf2_sha256$12000$GrqCKrh4gpuI$PLLnjVsHTgSDCcAv6ql0rJ2X/5RE9oNoaHHc8D/WTtE=	2017-01-25 12:40:45.021856+00	1	admin		alexey.diyan@gmail.com	t	t	t	2017-01-25 12:40:45.021856+00	f	f	2017-01-25 12:40:45.05685+00	\N
\.


--
-- Name: auth_user_id_seq; Type: SEQUENCE SET; Schema: public; Owner: sentry
--

SELECT pg_catalog.setval('auth_user_id_seq', 1, true);


--
-- Data for Name: django_admin_log; Type: TABLE DATA; Schema: public; Owner: sentry
--

COPY django_admin_log (id, action_time, user_id, content_type_id, object_id, object_repr, action_flag, change_message) FROM stdin;
\.


--
-- Name: django_admin_log_id_seq; Type: SEQUENCE SET; Schema: public; Owner: sentry
--

SELECT pg_catalog.setval('django_admin_log_id_seq', 1, false);


--
-- Data for Name: django_content_type; Type: TABLE DATA; Schema: public; Owner: sentry
--

COPY django_content_type (id, name, app_label, model) FROM stdin;
1	log entry	admin	logentry
2	permission	auth	permission
3	group	auth	group
4	content type	contenttypes	contenttype
5	session	sessions	session
6	site	sites	site
7	migration history	south	migrationhistory
8	activity	sentry	activity
9	api key	sentry	apikey
10	api token	sentry	apitoken
11	audit log entry	sentry	auditlogentry
12	authenticator	sentry	authenticator
13	auth identity	sentry	authidentity
14	auth provider	sentry	authprovider
15	broadcast	sentry	broadcast
16	broadcast seen	sentry	broadcastseen
17	commit	sentry	commit
18	commit author	sentry	commitauthor
19	commit file change	sentry	commitfilechange
20	counter	sentry	counter
21	file blob	sentry	fileblob
22	file	sentry	file
23	file blob index	sentry	fileblobindex
24	d sym sdk	sentry	dsymsdk
25	d sym object	sentry	dsymobject
26	d sym bundle	sentry	dsymbundle
27	d sym symbol	sentry	dsymsymbol
28	project d sym file	sentry	projectdsymfile
29	global d sym file	sentry	globaldsymfile
30	environment	sentry	environment
31	message	sentry	event
32	event mapping	sentry	eventmapping
33	event tag	sentry	eventtag
34	event user	sentry	eventuser
35	grouped message	sentry	group
36	group assignee	sentry	groupassignee
37	group bookmark	sentry	groupbookmark
38	group email thread	sentry	groupemailthread
39	group hash	sentry	grouphash
40	group meta	sentry	groupmeta
41	group redirect	sentry	groupredirect
42	group release	sentry	grouprelease
43	group resolution	sentry	groupresolution
44	group rule status	sentry	grouprulestatus
45	group seen	sentry	groupseen
46	group snooze	sentry	groupsnooze
47	group subscription	sentry	groupsubscription
48	group tag key	sentry	grouptagkey
49	group tag value	sentry	grouptagvalue
50	lost password hash	sentry	lostpasswordhash
51	option	sentry	option
52	organization	sentry	organization
53	organization access request	sentry	organizationaccessrequest
54	organization member team	sentry	organizationmemberteam
55	organization member	sentry	organizationmember
56	organization onboarding task	sentry	organizationonboardingtask
57	organization option	sentry	organizationoption
58	project	sentry	project
59	project bookmark	sentry	projectbookmark
60	project key	sentry	projectkey
61	project option	sentry	projectoption
62	project platform	sentry	projectplatform
63	release project	sentry	releaseproject
64	release	sentry	release
65	release commit	sentry	releasecommit
66	release environment	sentry	releaseenvironment
67	release file	sentry	releasefile
68	repository	sentry	repository
69	rule	sentry	rule
70	saved search	sentry	savedsearch
71	saved search user default	sentry	savedsearchuserdefault
72	tag key	sentry	tagkey
73	tag value	sentry	tagvalue
74	team	sentry	team
75	user	sentry	user
76	user avatar	sentry	useravatar
77	user email	sentry	useremail
78	user option	sentry	useroption
79	user report	sentry	userreport
80	node	nodestore	node
81	user social auth	social_auth	usersocialauth
\.


--
-- Name: django_content_type_id_seq; Type: SEQUENCE SET; Schema: public; Owner: sentry
--

SELECT pg_catalog.setval('django_content_type_id_seq', 81, true);


--
-- Data for Name: django_session; Type: TABLE DATA; Schema: public; Owner: sentry
--

COPY django_session (session_key, session_data, expire_date) FROM stdin;
\.


--
-- Data for Name: django_site; Type: TABLE DATA; Schema: public; Owner: sentry
--

COPY django_site (id, domain, name) FROM stdin;
1	example.com	example.com
\.


--
-- Name: django_site_id_seq; Type: SEQUENCE SET; Schema: public; Owner: sentry
--

SELECT pg_catalog.setval('django_site_id_seq', 1, true);


--
-- Data for Name: nodestore_node; Type: TABLE DATA; Schema: public; Owner: sentry
--

COPY nodestore_node (id, data, "timestamp") FROM stdin;
N0XToXlhRvyCOfYh1sRpZA==	eJytWetvnEgS/z5/Beu7FUzkYLp5Wx7vSdmNFGk3OZ2jfEkizEAPZsMAAsax5fP/fvWAgRlPktM40a7trn5VV9WvXhhpLWZXeqvKrrk387JTzSpOVGv+pdo2zpQ+q+XMSGsbFq23JAdG72/yVoP/Yq1Tbaf1k1qmStXEnUq1TZuXmXZ93cS3qqRF19ew15217ZWu7romhpGHh/vIwX1rxk12C7RgZhR1OLs60c82bXNWVElcnC3z8oxO0k9mtbBmMUzjmTQUNLzpuvr87Mz37NgPYi90Hemo1Av8ZBUs/TSRXmoLT5x/fJ0X8E6Vfv4XnX1Ttd15aFnWmaTT5CwGDosqTuPbDPgRNjIknNlry7Qca/cfXA1kWx6iunTQplUNnuLii5qqYpY9EkPUqFV0q5o2r0pc48/eyBnQVyA61dQNKATJJBEBItEfHrRUreJN0WmPj6gci+7Ac3Ak+v1P9XnVxckXkHlCKiWdSlTqqolBdUhz8BLp0oyHM5sy6ZgviRpabvIijdYtikQGeEu8bKM67m6QgMxN1FXky7P6vrupSmn6Z23eqZc1MAAW0rIaz5Zxq8z6HvbaFh5WNypKKmD4Dp9sC+TGlnCqtvePOdYWWotPwrX2LP4AFrCzkiajvFxVsDJTXTQSDNzj7G8Yjz7FeffQPMivbFdVs1bNolXFytwSaI9HuriNG5Sn7RNuAnhBvVkWeRJ9UfTacPYWVqWAEUSSRSMFEumi7r4mmiBLQSmZNDGiESzHIQUnmwKsAyXl2HRCF2d4q+NMzstTpJDdpYCBRKQiSWXgpPZquVKr1dKx03TpJ9bKp5M9PPnLV4AhHUUPcALcvp7cT8g8Dv2w3UVln+gFcFjQGE3WIjCALmHm264IV8PjM/IZro2LB5DiFNmv6x6H0hOGKZ7jTXHq+sQaKht9F2rzgjVD9vuqyIFdrVr+rZJOizvNuvNXgeet4qUQoXWJe0KUapevVdTWivDssdJ3Qe4JIt7EZVqQw/DkeNmuGTy9L1g5bprSfZ6N9zWqBR+BQwel5bk/FK1HYvX8WUYa3hyvYi8gFdcxoKklQoiq8VnziJa4g/044+Oj22ff6EvSEgQPjCb285/gO2hnzDYErKznU2Lgu+i3V82gB/BuTe9feuc01YxIlg5pxid8AfAx8PmELp/QxcEQrw2RGICgnv2CQJC0yS3gkDx+YOOzSNGB83wpBe5WSoG3lRLCpffMAcQzQBmS1lW6KdC7BYihEUJICTnkFaoE4QEhtIYlkygRCooSEKsnYSKUlCvYe2FijlNPPDxK3tzU6HmNB1yBPv5k38cz8xwpzydh5BQeHHr7Zz7SVT4FIGAvL6O4rpEUwMsprvXMRkVe0tvCAyEtietuA/GPYmfLoWWXhvFF8IF4UFnRGHynjXKJQaHCkrOMAnW/k1agwp8Tm4WFSOBoLCz3AO+NgssoebGeSIce0qikatJI3SWqxnwiapUqDRiSWOe009+L4KgqjPC4f5t3GLQ02L9kjJ2ntI9+wt+jyz3VCGGnrE2YAVAs8AdLNQTVZRyuBWSW2ffjr4Bss49BAlLFjOOOgCTxmMCT7YYbAVkm+Jlj4owQiEaK/xkFmMzpn+T1SD8e6AISU+BqCNXZkLX0OqTbA7wSvBr83MI+ewp2YBeOOohwgYdnDGkBqeyORZApQABLjRcvUMtkN5DHkt1sQczGSGbwcUyCPlNWz1kiXNBDM3sCxRcvOPnhs9HwCWIe2BxDDLLjHYiNRZKQ3nORJv0RaTKYsHZyckLz4VQerxoFKgNdlgwAyF2rNaiuV+r1tUmVi7UnoMvLSy0hYzJ332Dof933J/29AeO4AT8GdpH+os9ReraYXt5zZMsJcmwMu8+zMhtFfpTt22T7PSNg7xxIBbA24BMS8Z+BT8jeW4yrHOGQAlggY8ckG0M70TAQbvGSQfDFuHgIDZjNfwMNjj2iAWC9RQPUf0PA+GPwqQbCA90cA3LxtirBA27t+ZzOcw+o0PF+bFQ4Gi4iq3L8Q1DaB+E0ihmH3enpYB2L/veEZ5JjMGDQxzSaMAiujDGIziBC64nGxoTA8uIIGLZJk9cdjDYlWGmvAFeMgHSnlSgVL4uiyjIwZPPN29fvKJBASTJ1WGghi/fNRvGkM4UhRaCKJNqaUJ0anKidag+Pc0w0hLujLApfiwc6x9tPWrhyOacid5OnRn+AP8EmYQrLK33sfdBhlHmxP8CxZx2BPQ/FRI7/Su+fRGTKOAWUI2z/GaOCapN11VGxv8qpv+I547VINXmF+R/69YrWHap5pAf6Zh7cvWjTK9RkhdISbxdpB1Xu+SPmvKkPfiQlelsXTCHCn3pXbGC1XVptOvMrVOZg8n8QhN78rn2NW+3X5tOnUtd+1Yw8xZRkTgYOGfP3kETqHWpc1nA/6rWMhQ8DJBjSQCx/CB/rOOeekvPzIEHFUA8J3+vlQxWscaKdQObVliwZfxAUTzIxmIiLjQ5yPDYxA3aeanmZFJtUUVOpXXxkTemfWVRhL6repgNUZd+qEAGzBIY2bQsG9k/tC0LJhDk+1LV9Ry/AfPgCbJ5I5jsy/n/TNDwFHECZqInBhsCf67HBBlwLtqQhyJd+JpvkNI+CMldiFOxgRGC5GNzcn/BbNU/3h+CrfG49CK7RQCCklBC1cfEhLjb7FXEIAfdce+CAea5xtGIXyKPHS3wJ1HU/QnXo/B+oDt0R1VDCbSPpP7T3735/Z6QJFu3N/FyDQgWqla6419pNXVcNvLO9LxMNildV8Ju2ho3lhdkWStWGTeYdbiuTfDXkWGAD4MfSPI1WcV4YHIjDcOrCn/oN1TRV8wu6CzR8SE0OZ61P4p/Bl55qqAAzitI86aJoTl1iMfgJLFTIUchtvXjBtfklLTxUMI69d1wxFoQSCkISJqUJh1Sk5WsSY++LJJaI5AJARFGE5X4UaYuFpkcRLoki/Zye7E+9Kn4W+Gh9Bl/RgMQ3S6PRjZd8z6dPoOKvv/0Xfqs7Nf/tnxBCdfh/so1lGIyeQ1r4vCjCurIDmMKl2Dm30FhZEto4C5OaQX+/zMs5WqUUZORRhA2Lfq/83scKKWze0DvYfg83aoFRGrnT25H45FaP23q02p+uBhrnad/27w3mlInOBwXMDQufjqPicKsBXCSp0zBoTeInoovhewApU/u2yncitBvH5BqklHxtWiV8q7SppUbo3t6N9C2id8xOuqMFzw7CAVSu7vLOwKOMOZk9FmUxNVXbCOY7tY7G7x3SH1o0TIuqdU59SZwLOFVp0y80DPl7CXLFIIty7PJIGyUDlZsJpYUpiELvHD/pSBubUy6UGKZFQwfn+0aXtN3Bc71kbRGRvw2BJ1L5LfMDlvsasmHXdkCsZH/kJOgdNn0dkliZYPmPf0NNYmCC3AAf0XAXtjV05duhstNlEihiF+qPDn5CoIQN3JRHKiZkfWUvoW7AJfxxg42Ovxcxmhz6wrFlHjsNOticKWwcUrtf32rRFTvCwHY+uBsSkadTn5w/gUhs7+v9ly4ak9jWqov7/ql0+UsV1md6l3cFb/NnH46qP3EvAKM1/weVhHSP	2017-01-25 12:51:02.120997+00
\.


--
-- Data for Name: sentry_activity; Type: TABLE DATA; Schema: public; Owner: sentry
--

COPY sentry_activity (id, project_id, group_id, type, ident, user_id, datetime, data) FROM stdin;
\.


--
-- Name: sentry_activity_id_seq; Type: SEQUENCE SET; Schema: public; Owner: sentry
--

SELECT pg_catalog.setval('sentry_activity_id_seq', 1, false);


--
-- Data for Name: sentry_apikey; Type: TABLE DATA; Schema: public; Owner: sentry
--

COPY sentry_apikey (id, organization_id, label, key, scopes, status, date_added, allowed_origins) FROM stdin;
\.


--
-- Name: sentry_apikey_id_seq; Type: SEQUENCE SET; Schema: public; Owner: sentry
--

SELECT pg_catalog.setval('sentry_apikey_id_seq', 1, false);


--
-- Data for Name: sentry_apitoken; Type: TABLE DATA; Schema: public; Owner: sentry
--

COPY sentry_apitoken (id, key_id, user_id, token, scopes, date_added) FROM stdin;
\.


--
-- Name: sentry_apitoken_id_seq; Type: SEQUENCE SET; Schema: public; Owner: sentry
--

SELECT pg_catalog.setval('sentry_apitoken_id_seq', 1, false);


--
-- Data for Name: sentry_auditlogentry; Type: TABLE DATA; Schema: public; Owner: sentry
--

COPY sentry_auditlogentry (id, organization_id, actor_id, target_object, target_user_id, event, data, datetime, ip_address, actor_label, actor_key_id) FROM stdin;
1	2	1	1	1	6	eJwVjDEOgzAMAHd/JCxEBVpYuzJHYo3cxEKW4jYiUMHvcda70zUxd+AMF49h5z8ZyD3Mjw6KMyTIScEACyY66bKRL/y+18pt+AnkZ+12QvElHau2L22DUFsZ5LFqIfnQ5jmqnmDWtb0BoiMk/Q==	2017-01-25 14:06:03.105315+00	192.169.100.1	admin	\N
\.


--
-- Name: sentry_auditlogentry_id_seq; Type: SEQUENCE SET; Schema: public; Owner: sentry
--

SELECT pg_catalog.setval('sentry_auditlogentry_id_seq', 1, true);


--
-- Data for Name: sentry_authidentity; Type: TABLE DATA; Schema: public; Owner: sentry
--

COPY sentry_authidentity (id, user_id, auth_provider_id, ident, data, date_added, last_verified, last_synced) FROM stdin;
\.


--
-- Name: sentry_authidentity_id_seq; Type: SEQUENCE SET; Schema: public; Owner: sentry
--

SELECT pg_catalog.setval('sentry_authidentity_id_seq', 1, false);


--
-- Data for Name: sentry_authprovider; Type: TABLE DATA; Schema: public; Owner: sentry
--

COPY sentry_authprovider (id, organization_id, provider, config, date_added, sync_time, last_sync, default_role, default_global_access, flags) FROM stdin;
\.


--
-- Data for Name: sentry_authprovider_default_teams; Type: TABLE DATA; Schema: public; Owner: sentry
--

COPY sentry_authprovider_default_teams (id, authprovider_id, team_id) FROM stdin;
\.


--
-- Name: sentry_authprovider_default_teams_id_seq; Type: SEQUENCE SET; Schema: public; Owner: sentry
--

SELECT pg_catalog.setval('sentry_authprovider_default_teams_id_seq', 1, false);


--
-- Name: sentry_authprovider_id_seq; Type: SEQUENCE SET; Schema: public; Owner: sentry
--

SELECT pg_catalog.setval('sentry_authprovider_id_seq', 1, false);


--
-- Data for Name: sentry_broadcast; Type: TABLE DATA; Schema: public; Owner: sentry
--

COPY sentry_broadcast (id, message, link, is_active, date_added, title, upstream_id, date_expires) FROM stdin;
\.


--
-- Name: sentry_broadcast_id_seq; Type: SEQUENCE SET; Schema: public; Owner: sentry
--

SELECT pg_catalog.setval('sentry_broadcast_id_seq', 1, false);


--
-- Data for Name: sentry_broadcastseen; Type: TABLE DATA; Schema: public; Owner: sentry
--

COPY sentry_broadcastseen (id, broadcast_id, user_id, date_seen) FROM stdin;
\.


--
-- Name: sentry_broadcastseen_id_seq; Type: SEQUENCE SET; Schema: public; Owner: sentry
--

SELECT pg_catalog.setval('sentry_broadcastseen_id_seq', 1, false);


--
-- Data for Name: sentry_commit; Type: TABLE DATA; Schema: public; Owner: sentry
--

COPY sentry_commit (id, organization_id, repository_id, key, date_added, author_id, message) FROM stdin;
\.


--
-- Name: sentry_commit_id_seq; Type: SEQUENCE SET; Schema: public; Owner: sentry
--

SELECT pg_catalog.setval('sentry_commit_id_seq', 1, false);


--
-- Data for Name: sentry_commitauthor; Type: TABLE DATA; Schema: public; Owner: sentry
--

COPY sentry_commitauthor (id, organization_id, name, email) FROM stdin;
\.


--
-- Name: sentry_commitauthor_id_seq; Type: SEQUENCE SET; Schema: public; Owner: sentry
--

SELECT pg_catalog.setval('sentry_commitauthor_id_seq', 1, false);


--
-- Data for Name: sentry_commitfilechange; Type: TABLE DATA; Schema: public; Owner: sentry
--

COPY sentry_commitfilechange (id, organization_id, commit_id, filename, type) FROM stdin;
\.


--
-- Name: sentry_commitfilechange_id_seq; Type: SEQUENCE SET; Schema: public; Owner: sentry
--

SELECT pg_catalog.setval('sentry_commitfilechange_id_seq', 1, false);


--
-- Data for Name: sentry_dsymbundle; Type: TABLE DATA; Schema: public; Owner: sentry
--

COPY sentry_dsymbundle (id, sdk_id, object_id) FROM stdin;
\.


--
-- Name: sentry_dsymbundle_id_seq; Type: SEQUENCE SET; Schema: public; Owner: sentry
--

SELECT pg_catalog.setval('sentry_dsymbundle_id_seq', 1, false);


--
-- Data for Name: sentry_dsymobject; Type: TABLE DATA; Schema: public; Owner: sentry
--

COPY sentry_dsymobject (id, cpu_name, object_path, uuid, vmaddr, vmsize) FROM stdin;
\.


--
-- Name: sentry_dsymobject_id_seq; Type: SEQUENCE SET; Schema: public; Owner: sentry
--

SELECT pg_catalog.setval('sentry_dsymobject_id_seq', 1, false);


--
-- Data for Name: sentry_dsymsdk; Type: TABLE DATA; Schema: public; Owner: sentry
--

COPY sentry_dsymsdk (id, dsym_type, sdk_name, version_major, version_minor, version_patchlevel, version_build) FROM stdin;
\.


--
-- Name: sentry_dsymsdk_id_seq; Type: SEQUENCE SET; Schema: public; Owner: sentry
--

SELECT pg_catalog.setval('sentry_dsymsdk_id_seq', 1, false);


--
-- Data for Name: sentry_dsymsymbol; Type: TABLE DATA; Schema: public; Owner: sentry
--

COPY sentry_dsymsymbol (id, object_id, address, symbol) FROM stdin;
\.


--
-- Name: sentry_dsymsymbol_id_seq; Type: SEQUENCE SET; Schema: public; Owner: sentry
--

SELECT pg_catalog.setval('sentry_dsymsymbol_id_seq', 1, false);


--
-- Data for Name: sentry_environment; Type: TABLE DATA; Schema: public; Owner: sentry
--

COPY sentry_environment (id, project_id, name, date_added) FROM stdin;
1	2		2017-01-25 12:51:02.113843+00
\.


--
-- Name: sentry_environment_id_seq; Type: SEQUENCE SET; Schema: public; Owner: sentry
--

SELECT pg_catalog.setval('sentry_environment_id_seq', 1, true);


--
-- Data for Name: sentry_environmentrelease; Type: TABLE DATA; Schema: public; Owner: sentry
--

COPY sentry_environmentrelease (id, project_id, release_id, environment_id, first_seen, last_seen, organization_id) FROM stdin;
\.


--
-- Name: sentry_environmentrelease_id_seq; Type: SEQUENCE SET; Schema: public; Owner: sentry
--

SELECT pg_catalog.setval('sentry_environmentrelease_id_seq', 1, false);


--
-- Data for Name: sentry_eventmapping; Type: TABLE DATA; Schema: public; Owner: sentry
--

COPY sentry_eventmapping (id, project_id, group_id, event_id, date_added) FROM stdin;
1	2	1	dcf8c1d1cd284d3fbfeffb43ddb7c0f7	2017-01-25 12:51:02.107047+00
\.


--
-- Name: sentry_eventmapping_id_seq; Type: SEQUENCE SET; Schema: public; Owner: sentry
--

SELECT pg_catalog.setval('sentry_eventmapping_id_seq', 1, true);


--
-- Data for Name: sentry_eventtag; Type: TABLE DATA; Schema: public; Owner: sentry
--

COPY sentry_eventtag (id, project_id, event_id, key_id, value_id, date_added, group_id) FROM stdin;
1	2	1	1	1	2017-01-25 12:51:02.165648+00	1
2	2	1	2	2	2017-01-25 12:51:02.182397+00	1
\.


--
-- Name: sentry_eventtag_id_seq; Type: SEQUENCE SET; Schema: public; Owner: sentry
--

SELECT pg_catalog.setval('sentry_eventtag_id_seq', 2, true);


--
-- Data for Name: sentry_eventuser; Type: TABLE DATA; Schema: public; Owner: sentry
--

COPY sentry_eventuser (id, project_id, ident, email, username, ip_address, date_added, hash) FROM stdin;
\.


--
-- Name: sentry_eventuser_id_seq; Type: SEQUENCE SET; Schema: public; Owner: sentry
--

SELECT pg_catalog.setval('sentry_eventuser_id_seq', 1, false);


--
-- Data for Name: sentry_file; Type: TABLE DATA; Schema: public; Owner: sentry
--

COPY sentry_file (id, name, path, type, size, "timestamp", checksum, headers, blob_id) FROM stdin;
\.


--
-- Name: sentry_file_id_seq; Type: SEQUENCE SET; Schema: public; Owner: sentry
--

SELECT pg_catalog.setval('sentry_file_id_seq', 1, false);


--
-- Data for Name: sentry_fileblob; Type: TABLE DATA; Schema: public; Owner: sentry
--

COPY sentry_fileblob (id, path, size, checksum, "timestamp") FROM stdin;
\.


--
-- Name: sentry_fileblob_id_seq; Type: SEQUENCE SET; Schema: public; Owner: sentry
--

SELECT pg_catalog.setval('sentry_fileblob_id_seq', 1, false);


--
-- Data for Name: sentry_fileblobindex; Type: TABLE DATA; Schema: public; Owner: sentry
--

COPY sentry_fileblobindex (id, file_id, blob_id, "offset") FROM stdin;
\.


--
-- Name: sentry_fileblobindex_id_seq; Type: SEQUENCE SET; Schema: public; Owner: sentry
--

SELECT pg_catalog.setval('sentry_fileblobindex_id_seq', 1, false);


--
-- Data for Name: sentry_filterkey; Type: TABLE DATA; Schema: public; Owner: sentry
--

COPY sentry_filterkey (id, project_id, key, values_seen, label, status) FROM stdin;
1	2	server_name	0	\N	0
2	2	level	0	\N	0
\.


--
-- Name: sentry_filterkey_id_seq; Type: SEQUENCE SET; Schema: public; Owner: sentry
--

SELECT pg_catalog.setval('sentry_filterkey_id_seq', 2, true);


--
-- Data for Name: sentry_filtervalue; Type: TABLE DATA; Schema: public; Owner: sentry
--

COPY sentry_filtervalue (id, key, value, project_id, times_seen, last_seen, first_seen, data) FROM stdin;
1	server_name	e739e3dbc8e1	2	0	2017-01-25 12:51:02.160658+00	2017-01-25 12:51:02.160672+00	\N
2	level	info	2	0	2017-01-25 12:51:02.177787+00	2017-01-25 12:51:02.177796+00	\N
\.


--
-- Name: sentry_filtervalue_id_seq; Type: SEQUENCE SET; Schema: public; Owner: sentry
--

SELECT pg_catalog.setval('sentry_filtervalue_id_seq', 2, true);


--
-- Data for Name: sentry_globaldsymfile; Type: TABLE DATA; Schema: public; Owner: sentry
--

COPY sentry_globaldsymfile (id, file_id, object_name, cpu_name, uuid) FROM stdin;
\.


--
-- Name: sentry_globaldsymfile_id_seq; Type: SEQUENCE SET; Schema: public; Owner: sentry
--

SELECT pg_catalog.setval('sentry_globaldsymfile_id_seq', 1, false);


--
-- Data for Name: sentry_groupasignee; Type: TABLE DATA; Schema: public; Owner: sentry
--

COPY sentry_groupasignee (id, project_id, group_id, user_id, date_added) FROM stdin;
\.


--
-- Name: sentry_groupasignee_id_seq; Type: SEQUENCE SET; Schema: public; Owner: sentry
--

SELECT pg_catalog.setval('sentry_groupasignee_id_seq', 1, false);


--
-- Data for Name: sentry_groupbookmark; Type: TABLE DATA; Schema: public; Owner: sentry
--

COPY sentry_groupbookmark (id, project_id, group_id, user_id, date_added) FROM stdin;
\.


--
-- Name: sentry_groupbookmark_id_seq; Type: SEQUENCE SET; Schema: public; Owner: sentry
--

SELECT pg_catalog.setval('sentry_groupbookmark_id_seq', 1, false);


--
-- Data for Name: sentry_groupedmessage; Type: TABLE DATA; Schema: public; Owner: sentry
--

COPY sentry_groupedmessage (id, logger, level, message, view, status, times_seen, last_seen, first_seen, data, score, project_id, time_spent_total, time_spent_count, resolved_at, active_at, is_public, platform, num_comments, first_release_id, short_id) FROM stdin;
1		20	This is a test message generated using ``raven test`` __main__ in <module>	__main__ in <module>	0	1	2017-01-25 12:51:01+00	2017-01-25 12:51:01+00	eJwdykEOgjAUBND9P0V3sDKpQO0JvADErf2xY20CpOF/SLy9xWQ2M/PaWCyNzcyizw0v5AOxoXKlu+390PXeOUsyNvotqHtXbcSb91lr689ngXJk5doHamNxlWjW+eQ3ekyfLKaGjULULBDhBJOwYmNFNLvkNZkQNj6w/lEIVDyJXH6p7jGr	1485348661	2	0	0	\N	2017-01-25 12:51:01+00	f	python	0	\N	1
\.


--
-- Name: sentry_groupedmessage_id_seq; Type: SEQUENCE SET; Schema: public; Owner: sentry
--

SELECT pg_catalog.setval('sentry_groupedmessage_id_seq', 1, true);


--
-- Data for Name: sentry_groupemailthread; Type: TABLE DATA; Schema: public; Owner: sentry
--

COPY sentry_groupemailthread (id, email, project_id, group_id, msgid, date) FROM stdin;
\.


--
-- Name: sentry_groupemailthread_id_seq; Type: SEQUENCE SET; Schema: public; Owner: sentry
--

SELECT pg_catalog.setval('sentry_groupemailthread_id_seq', 1, false);


--
-- Data for Name: sentry_grouphash; Type: TABLE DATA; Schema: public; Owner: sentry
--

COPY sentry_grouphash (id, project_id, hash, group_id) FROM stdin;
1	2	61c945c7d8c3df5d0ebb23ae715f6a67	1
\.


--
-- Name: sentry_grouphash_id_seq; Type: SEQUENCE SET; Schema: public; Owner: sentry
--

SELECT pg_catalog.setval('sentry_grouphash_id_seq', 1, true);


--
-- Data for Name: sentry_groupmeta; Type: TABLE DATA; Schema: public; Owner: sentry
--

COPY sentry_groupmeta (id, group_id, key, value) FROM stdin;
\.


--
-- Name: sentry_groupmeta_id_seq; Type: SEQUENCE SET; Schema: public; Owner: sentry
--

SELECT pg_catalog.setval('sentry_groupmeta_id_seq', 1, false);


--
-- Data for Name: sentry_groupredirect; Type: TABLE DATA; Schema: public; Owner: sentry
--

COPY sentry_groupredirect (id, group_id, previous_group_id) FROM stdin;
\.


--
-- Name: sentry_groupredirect_id_seq; Type: SEQUENCE SET; Schema: public; Owner: sentry
--

SELECT pg_catalog.setval('sentry_groupredirect_id_seq', 1, false);


--
-- Data for Name: sentry_grouprelease; Type: TABLE DATA; Schema: public; Owner: sentry
--

COPY sentry_grouprelease (id, project_id, group_id, release_id, environment, first_seen, last_seen) FROM stdin;
\.


--
-- Name: sentry_grouprelease_id_seq; Type: SEQUENCE SET; Schema: public; Owner: sentry
--

SELECT pg_catalog.setval('sentry_grouprelease_id_seq', 1, false);


--
-- Data for Name: sentry_groupresolution; Type: TABLE DATA; Schema: public; Owner: sentry
--

COPY sentry_groupresolution (id, group_id, release_id, datetime, status) FROM stdin;
\.


--
-- Name: sentry_groupresolution_id_seq; Type: SEQUENCE SET; Schema: public; Owner: sentry
--

SELECT pg_catalog.setval('sentry_groupresolution_id_seq', 1, false);


--
-- Data for Name: sentry_grouprulestatus; Type: TABLE DATA; Schema: public; Owner: sentry
--

COPY sentry_grouprulestatus (id, project_id, rule_id, group_id, status, date_added, last_active) FROM stdin;
1	2	2	1	0	2017-01-25 12:51:02.193893+00	2017-01-25 12:51:02.19954+00
\.


--
-- Name: sentry_grouprulestatus_id_seq; Type: SEQUENCE SET; Schema: public; Owner: sentry
--

SELECT pg_catalog.setval('sentry_grouprulestatus_id_seq', 1, true);


--
-- Data for Name: sentry_groupseen; Type: TABLE DATA; Schema: public; Owner: sentry
--

COPY sentry_groupseen (id, project_id, group_id, user_id, last_seen) FROM stdin;
\.


--
-- Name: sentry_groupseen_id_seq; Type: SEQUENCE SET; Schema: public; Owner: sentry
--

SELECT pg_catalog.setval('sentry_groupseen_id_seq', 1, false);


--
-- Data for Name: sentry_groupsnooze; Type: TABLE DATA; Schema: public; Owner: sentry
--

COPY sentry_groupsnooze (id, group_id, until) FROM stdin;
\.


--
-- Name: sentry_groupsnooze_id_seq; Type: SEQUENCE SET; Schema: public; Owner: sentry
--

SELECT pg_catalog.setval('sentry_groupsnooze_id_seq', 1, false);


--
-- Data for Name: sentry_groupsubscription; Type: TABLE DATA; Schema: public; Owner: sentry
--

COPY sentry_groupsubscription (id, project_id, group_id, user_id, is_active, reason, date_added) FROM stdin;
\.


--
-- Name: sentry_groupsubscription_id_seq; Type: SEQUENCE SET; Schema: public; Owner: sentry
--

SELECT pg_catalog.setval('sentry_groupsubscription_id_seq', 1, false);


--
-- Data for Name: sentry_grouptagkey; Type: TABLE DATA; Schema: public; Owner: sentry
--

COPY sentry_grouptagkey (id, project_id, group_id, key, values_seen) FROM stdin;
\.


--
-- Name: sentry_grouptagkey_id_seq; Type: SEQUENCE SET; Schema: public; Owner: sentry
--

SELECT pg_catalog.setval('sentry_grouptagkey_id_seq', 1, false);


--
-- Data for Name: sentry_lostpasswordhash; Type: TABLE DATA; Schema: public; Owner: sentry
--

COPY sentry_lostpasswordhash (id, user_id, hash, date_added) FROM stdin;
\.


--
-- Name: sentry_lostpasswordhash_id_seq; Type: SEQUENCE SET; Schema: public; Owner: sentry
--

SELECT pg_catalog.setval('sentry_lostpasswordhash_id_seq', 1, false);


--
-- Data for Name: sentry_message; Type: TABLE DATA; Schema: public; Owner: sentry
--

COPY sentry_message (id, message, datetime, data, group_id, message_id, project_id, time_spent, platform) FROM stdin;
1	This is a test message generated using ``raven test`` __main__ in <module>	2017-01-25 12:51:01+00	eJzTSCkw5ApWz8tPSY3PTFHnKjAC8vwMIkLyI3Iygsoqnf3TIjMMi4MKohxtbYHSxlzFegCg1g+U	1	dcf8c1d1cd284d3fbfeffb43ddb7c0f7	2	\N	python
\.


--
-- Name: sentry_message_id_seq; Type: SEQUENCE SET; Schema: public; Owner: sentry
--

SELECT pg_catalog.setval('sentry_message_id_seq', 1, true);


--
-- Data for Name: sentry_messagefiltervalue; Type: TABLE DATA; Schema: public; Owner: sentry
--

COPY sentry_messagefiltervalue (id, group_id, times_seen, key, value, project_id, last_seen, first_seen) FROM stdin;
\.


--
-- Name: sentry_messagefiltervalue_id_seq; Type: SEQUENCE SET; Schema: public; Owner: sentry
--

SELECT pg_catalog.setval('sentry_messagefiltervalue_id_seq', 1, false);


--
-- Data for Name: sentry_messageindex; Type: TABLE DATA; Schema: public; Owner: sentry
--

COPY sentry_messageindex (id, object_id, "column", value) FROM stdin;
\.


--
-- Name: sentry_messageindex_id_seq; Type: SEQUENCE SET; Schema: public; Owner: sentry
--

SELECT pg_catalog.setval('sentry_messageindex_id_seq', 1, false);


--
-- Data for Name: sentry_option; Type: TABLE DATA; Schema: public; Owner: sentry
--

COPY sentry_option (id, key, value, last_updated) FROM stdin;
1	sentry:version-configured	gAJYBgAAADguMTIuMHEBLg==	2017-01-25 12:40:45.070812+00
\.


--
-- Name: sentry_option_id_seq; Type: SEQUENCE SET; Schema: public; Owner: sentry
--

SELECT pg_catalog.setval('sentry_option_id_seq', 1, true);


--
-- Data for Name: sentry_organization; Type: TABLE DATA; Schema: public; Owner: sentry
--

COPY sentry_organization (id, name, status, date_added, slug, flags, default_role) FROM stdin;
1	Sentry	0	2017-01-25 12:40:41.001161+00	sentry	1	member
2	ACME-Team	0	2017-01-25 12:40:46.243583+00	acme-team	1	member
\.


--
-- Name: sentry_organization_id_seq; Type: SEQUENCE SET; Schema: public; Owner: sentry
--

SELECT pg_catalog.setval('sentry_organization_id_seq', 2, true);


--
-- Data for Name: sentry_organizationaccessrequest; Type: TABLE DATA; Schema: public; Owner: sentry
--

COPY sentry_organizationaccessrequest (id, team_id, member_id) FROM stdin;
\.


--
-- Name: sentry_organizationaccessrequest_id_seq; Type: SEQUENCE SET; Schema: public; Owner: sentry
--

SELECT pg_catalog.setval('sentry_organizationaccessrequest_id_seq', 1, false);


--
-- Data for Name: sentry_organizationmember; Type: TABLE DATA; Schema: public; Owner: sentry
--

COPY sentry_organizationmember (id, organization_id, user_id, type, date_added, email, has_global_access, flags, role, token) FROM stdin;
1	2	1	50	2017-01-25 12:40:46.280595+00	\N	t	0	owner	\N
\.


--
-- Name: sentry_organizationmember_id_seq; Type: SEQUENCE SET; Schema: public; Owner: sentry
--

SELECT pg_catalog.setval('sentry_organizationmember_id_seq', 1, true);


--
-- Data for Name: sentry_organizationmember_teams; Type: TABLE DATA; Schema: public; Owner: sentry
--

COPY sentry_organizationmember_teams (id, organizationmember_id, team_id, is_active) FROM stdin;
1	1	2	t
\.


--
-- Name: sentry_organizationmember_teams_id_seq; Type: SEQUENCE SET; Schema: public; Owner: sentry
--

SELECT pg_catalog.setval('sentry_organizationmember_teams_id_seq', 1, true);


--
-- Data for Name: sentry_organizationonboardingtask; Type: TABLE DATA; Schema: public; Owner: sentry
--

COPY sentry_organizationonboardingtask (id, organization_id, user_id, task, status, date_completed, project_id, data) FROM stdin;
1	2	\N	2	1	2017-01-25 12:51:01+00	2	{"platform": "python"}
\.


--
-- Name: sentry_organizationonboardingtask_id_seq; Type: SEQUENCE SET; Schema: public; Owner: sentry
--

SELECT pg_catalog.setval('sentry_organizationonboardingtask_id_seq', 1, true);


--
-- Data for Name: sentry_organizationoptions; Type: TABLE DATA; Schema: public; Owner: sentry
--

COPY sentry_organizationoptions (id, organization_id, key, value) FROM stdin;
\.


--
-- Name: sentry_organizationoptions_id_seq; Type: SEQUENCE SET; Schema: public; Owner: sentry
--

SELECT pg_catalog.setval('sentry_organizationoptions_id_seq', 1, false);


--
-- Data for Name: sentry_project; Type: TABLE DATA; Schema: public; Owner: sentry
--

COPY sentry_project (id, name, public, date_added, status, slug, team_id, organization_id, first_event, forced_color) FROM stdin;
1	Internal	f	2017-01-25 12:40:41.012938+00	0	internal	1	1	\N	\N
2	ACME	f	2017-01-25 12:40:47.443433+00	0	acme	2	2	2017-01-25 12:51:01+00	\N
\.


--
-- Name: sentry_project_id_seq; Type: SEQUENCE SET; Schema: public; Owner: sentry
--

SELECT pg_catalog.setval('sentry_project_id_seq', 2, true);


--
-- Data for Name: sentry_projectbookmark; Type: TABLE DATA; Schema: public; Owner: sentry
--

COPY sentry_projectbookmark (id, project_id, user_id, date_added) FROM stdin;
\.


--
-- Name: sentry_projectbookmark_id_seq; Type: SEQUENCE SET; Schema: public; Owner: sentry
--

SELECT pg_catalog.setval('sentry_projectbookmark_id_seq', 1, false);


--
-- Data for Name: sentry_projectcounter; Type: TABLE DATA; Schema: public; Owner: sentry
--

COPY sentry_projectcounter (id, project_id, value) FROM stdin;
1	2	1
\.


--
-- Name: sentry_projectcounter_id_seq; Type: SEQUENCE SET; Schema: public; Owner: sentry
--

SELECT pg_catalog.setval('sentry_projectcounter_id_seq', 1, true);


--
-- Data for Name: sentry_projectdsymfile; Type: TABLE DATA; Schema: public; Owner: sentry
--

COPY sentry_projectdsymfile (id, file_id, object_name, cpu_name, project_id, uuid) FROM stdin;
\.


--
-- Name: sentry_projectdsymfile_id_seq; Type: SEQUENCE SET; Schema: public; Owner: sentry
--

SELECT pg_catalog.setval('sentry_projectdsymfile_id_seq', 1, false);


--
-- Data for Name: sentry_projectkey; Type: TABLE DATA; Schema: public; Owner: sentry
--

COPY sentry_projectkey (id, project_id, public_key, secret_key, date_added, roles, label, status) FROM stdin;
1	1	60383c2205b64b4bbbfa57d595578ea3	c01a0aab58164f8f9e4babfed623bab2	2017-01-25 12:40:41.020189+00	1	Default	0
2	2	763a78a695424ed687cf8b7dc26d3161	763a78a695424ed687cf8b7dc26d3161	2017-01-25 12:40:47.44724+00	1	Default	0
\.


--
-- Name: sentry_projectkey_id_seq; Type: SEQUENCE SET; Schema: public; Owner: sentry
--

SELECT pg_catalog.setval('sentry_projectkey_id_seq', 2, true);


--
-- Data for Name: sentry_projectoptions; Type: TABLE DATA; Schema: public; Owner: sentry
--

COPY sentry_projectoptions (id, project_id, key, value) FROM stdin;
1	1	sentry:origins	gAJdcQFVASphLg==
\.


--
-- Name: sentry_projectoptions_id_seq; Type: SEQUENCE SET; Schema: public; Owner: sentry
--

SELECT pg_catalog.setval('sentry_projectoptions_id_seq', 1, true);


--
-- Data for Name: sentry_projectplatform; Type: TABLE DATA; Schema: public; Owner: sentry
--

COPY sentry_projectplatform (id, project_id, platform, date_added, last_seen) FROM stdin;
\.


--
-- Name: sentry_projectplatform_id_seq; Type: SEQUENCE SET; Schema: public; Owner: sentry
--

SELECT pg_catalog.setval('sentry_projectplatform_id_seq', 1, false);


--
-- Data for Name: sentry_release; Type: TABLE DATA; Schema: public; Owner: sentry
--

COPY sentry_release (id, project_id, version, date_added, date_released, ref, url, date_started, data, new_groups, owner_id, organization_id) FROM stdin;
\.


--
-- Name: sentry_release_id_seq; Type: SEQUENCE SET; Schema: public; Owner: sentry
--

SELECT pg_catalog.setval('sentry_release_id_seq', 1, false);


--
-- Data for Name: sentry_release_project; Type: TABLE DATA; Schema: public; Owner: sentry
--

COPY sentry_release_project (id, project_id, release_id) FROM stdin;
\.


--
-- Name: sentry_release_project_id_seq; Type: SEQUENCE SET; Schema: public; Owner: sentry
--

SELECT pg_catalog.setval('sentry_release_project_id_seq', 1, false);


--
-- Data for Name: sentry_releasecommit; Type: TABLE DATA; Schema: public; Owner: sentry
--

COPY sentry_releasecommit (id, project_id, release_id, commit_id, "order", organization_id) FROM stdin;
\.


--
-- Name: sentry_releasecommit_id_seq; Type: SEQUENCE SET; Schema: public; Owner: sentry
--

SELECT pg_catalog.setval('sentry_releasecommit_id_seq', 1, false);


--
-- Data for Name: sentry_releasefile; Type: TABLE DATA; Schema: public; Owner: sentry
--

COPY sentry_releasefile (id, project_id, release_id, file_id, ident, name, organization_id) FROM stdin;
\.


--
-- Name: sentry_releasefile_id_seq; Type: SEQUENCE SET; Schema: public; Owner: sentry
--

SELECT pg_catalog.setval('sentry_releasefile_id_seq', 1, false);


--
-- Data for Name: sentry_repository; Type: TABLE DATA; Schema: public; Owner: sentry
--

COPY sentry_repository (id, organization_id, name, date_added, url, provider, external_id, config, status) FROM stdin;
\.


--
-- Name: sentry_repository_id_seq; Type: SEQUENCE SET; Schema: public; Owner: sentry
--

SELECT pg_catalog.setval('sentry_repository_id_seq', 1, false);


--
-- Data for Name: sentry_rule; Type: TABLE DATA; Schema: public; Owner: sentry
--

COPY sentry_rule (id, project_id, label, data, date_added, status) FROM stdin;
1	1	Send a notification for new events	eJxlj80OAiEMhO99ETgRd/0/GqNHLzzAhgCrJAiEosm+vXRl48FbZzr9JuUmdSCZjsG44mJABqkH7tMauEmbunKmWts6oA0lTyK/vEXxOxCjy1gGtDYM9l0z4kqGrPpC8rwkK2YHqFAypZeqPVUdqOoI97+SlhMhFjdOjX6bxYw+6cbtVl/wUxX9IE0/Ke9p7AHFBx5HTWU=	2017-01-25 12:40:41.026544+00	0
2	2	Send a notification for new events	eJxlj80OAiEMhO99ETgRd/0/GqNHLzzAhgCrJAiEosm+vXRl48FbZzr9JuUmdSCZjsG44mJABqkH7tMauEmbunKmWts6oA0lTyK/vEXxOxCjy1gGtDYM9l0z4kqGrPpC8rwkK2YHqFAypZeqPVUdqOoI97+SlhMhFjdOjX6bxYw+6cbtVl/wUxX9IE0/Ke9p7AHFBx5HTWU=	2017-01-25 12:40:47.448278+00	0
\.


--
-- Name: sentry_rule_id_seq; Type: SEQUENCE SET; Schema: public; Owner: sentry
--

SELECT pg_catalog.setval('sentry_rule_id_seq', 2, true);


--
-- Data for Name: sentry_savedsearch; Type: TABLE DATA; Schema: public; Owner: sentry
--

COPY sentry_savedsearch (id, project_id, name, query, date_added, is_default) FROM stdin;
1	1	Unresolved Issues	is:unresolved	2017-01-25 12:40:41.033156+00	t
2	1	Needs Triage	is:unresolved is:unassigned	2017-01-25 12:40:41.038376+00	f
3	1	Assigned To Me	is:unresolved assigned:me	2017-01-25 12:40:41.043317+00	f
4	1	My Bookmarks	is:unresolved bookmarks:me	2017-01-25 12:40:41.048145+00	f
5	1	New Today	is:unresolved age:-24h	2017-01-25 12:40:41.053069+00	f
6	2	Unresolved Issues	is:unresolved	2017-01-25 12:40:47.449256+00	t
7	2	Needs Triage	is:unresolved is:unassigned	2017-01-25 12:40:47.449852+00	f
8	2	Assigned To Me	is:unresolved assigned:me	2017-01-25 12:40:47.450274+00	f
9	2	My Bookmarks	is:unresolved bookmarks:me	2017-01-25 12:40:47.450655+00	f
10	2	New Today	is:unresolved age:-24h	2017-01-25 12:40:47.45105+00	f
\.


--
-- Name: sentry_savedsearch_id_seq; Type: SEQUENCE SET; Schema: public; Owner: sentry
--

SELECT pg_catalog.setval('sentry_savedsearch_id_seq', 10, true);


--
-- Data for Name: sentry_savedsearch_userdefault; Type: TABLE DATA; Schema: public; Owner: sentry
--

COPY sentry_savedsearch_userdefault (id, savedsearch_id, project_id, user_id) FROM stdin;
\.


--
-- Name: sentry_savedsearch_userdefault_id_seq; Type: SEQUENCE SET; Schema: public; Owner: sentry
--

SELECT pg_catalog.setval('sentry_savedsearch_userdefault_id_seq', 1, false);


--
-- Data for Name: sentry_team; Type: TABLE DATA; Schema: public; Owner: sentry
--

COPY sentry_team (id, slug, name, date_added, status, organization_id) FROM stdin;
1	sentry	Sentry	2017-01-25 12:40:41.00742+00	0	1
2	acme-team	ACME-Team	2017-01-25 12:40:46.287246+00	0	2
\.


--
-- Name: sentry_team_id_seq; Type: SEQUENCE SET; Schema: public; Owner: sentry
--

SELECT pg_catalog.setval('sentry_team_id_seq', 2, true);


--
-- Data for Name: sentry_useravatar; Type: TABLE DATA; Schema: public; Owner: sentry
--

COPY sentry_useravatar (id, user_id, file_id, ident, avatar_type) FROM stdin;
\.


--
-- Name: sentry_useravatar_id_seq; Type: SEQUENCE SET; Schema: public; Owner: sentry
--

SELECT pg_catalog.setval('sentry_useravatar_id_seq', 1, false);


--
-- Data for Name: sentry_useremail; Type: TABLE DATA; Schema: public; Owner: sentry
--

COPY sentry_useremail (id, user_id, email, validation_hash, date_hash_added, is_verified) FROM stdin;
1	1	alexey.diyan@gmail.com	fmyQ8FcxeVc1SwmdV9BEyTwifxtzmKiZ	2017-01-25 12:40:45.065436+00	f
\.


--
-- Name: sentry_useremail_id_seq; Type: SEQUENCE SET; Schema: public; Owner: sentry
--

SELECT pg_catalog.setval('sentry_useremail_id_seq', 1, true);


--
-- Data for Name: sentry_useroption; Type: TABLE DATA; Schema: public; Owner: sentry
--

COPY sentry_useroption (id, user_id, project_id, key, value) FROM stdin;
\.


--
-- Name: sentry_useroption_id_seq; Type: SEQUENCE SET; Schema: public; Owner: sentry
--

SELECT pg_catalog.setval('sentry_useroption_id_seq', 1, false);


--
-- Data for Name: sentry_userreport; Type: TABLE DATA; Schema: public; Owner: sentry
--

COPY sentry_userreport (id, project_id, group_id, event_id, name, email, comments, date_added) FROM stdin;
\.


--
-- Name: sentry_userreport_id_seq; Type: SEQUENCE SET; Schema: public; Owner: sentry
--

SELECT pg_catalog.setval('sentry_userreport_id_seq', 1, false);


--
-- Data for Name: social_auth_usersocialauth; Type: TABLE DATA; Schema: public; Owner: sentry
--

COPY social_auth_usersocialauth (id, user_id, provider, uid, extra_data) FROM stdin;
\.


--
-- Name: social_auth_usersocialauth_id_seq; Type: SEQUENCE SET; Schema: public; Owner: sentry
--

SELECT pg_catalog.setval('social_auth_usersocialauth_id_seq', 1, false);


--
-- Data for Name: south_migrationhistory; Type: TABLE DATA; Schema: public; Owner: sentry
--

COPY south_migrationhistory (id, app_name, migration, applied) FROM stdin;
1	sentry	0001_initial	2017-01-25 12:39:25.769971+00
2	sentry	0002_auto__del_field_groupedmessage_url__chg_field_groupedmessage_view__chg	2017-01-25 12:39:25.805279+00
3	sentry	0003_auto__add_field_message_group__del_field_groupedmessage_server_name	2017-01-25 12:39:25.827033+00
4	sentry	0004_auto__add_filtervalue__add_unique_filtervalue_key_value	2017-01-25 12:39:25.855586+00
5	sentry	0005_auto	2017-01-25 12:39:25.874371+00
6	sentry	0006_auto	2017-01-25 12:39:25.883466+00
7	sentry	0007_auto__add_field_message_site	2017-01-25 12:39:25.911015+00
8	sentry	0008_auto__chg_field_message_view__add_field_groupedmessage_data__chg_field	2017-01-25 12:39:25.958106+00
9	sentry	0009_auto__add_field_message_message_id	2017-01-25 12:39:25.991362+00
10	sentry	0010_auto__add_messageindex__add_unique_messageindex_column_value_object_id	2017-01-25 12:39:26.021281+00
11	sentry	0011_auto__add_field_groupedmessage_score	2017-01-25 12:39:27.35127+00
12	sentry	0012_auto	2017-01-25 12:39:27.371306+00
13	sentry	0013_auto__add_messagecountbyminute__add_unique_messagecountbyminute_group_	2017-01-25 12:39:27.443711+00
14	sentry	0014_auto	2017-01-25 12:39:28.050469+00
15	sentry	0014_auto__add_project__add_projectmember__add_unique_projectmember_project	2017-01-25 12:39:28.166745+00
16	sentry	0015_auto__add_field_message_project__add_field_messagecountbyminute_projec	2017-01-25 12:39:28.298805+00
17	sentry	0016_auto__add_field_projectmember_is_superuser	2017-01-25 12:39:28.359541+00
18	sentry	0017_auto__add_field_projectmember_api_key	2017-01-25 12:39:28.40457+00
19	sentry	0018_auto__chg_field_project_owner	2017-01-25 12:39:28.456314+00
20	sentry	0019_auto__del_field_projectmember_api_key__add_field_projectmember_public_	2017-01-25 12:39:28.535075+00
21	sentry	0020_auto__add_projectdomain__add_unique_projectdomain_project_domain	2017-01-25 12:39:28.595044+00
22	sentry	0021_auto__del_message__del_groupedmessage__del_unique_groupedmessage_proje	2017-01-25 12:39:28.622141+00
23	sentry	0022_auto__del_field_group_class_name__del_field_group_traceback__del_field	2017-01-25 12:39:28.651644+00
24	sentry	0023_auto__add_field_event_time_spent	2017-01-25 12:39:28.677647+00
25	sentry	0024_auto__add_field_group_time_spent_total__add_field_group_time_spent_cou	2017-01-25 12:39:30.839325+00
26	sentry	0025_auto__add_field_messagecountbyminute_time_spent_total__add_field_messa	2017-01-25 12:39:30.949913+00
27	sentry	0026_auto__add_field_project_status	2017-01-25 12:39:31.008152+00
28	sentry	0027_auto__chg_field_event_server_name	2017-01-25 12:39:31.041023+00
29	sentry	0028_auto__add_projectoptions__add_unique_projectoptions_project_key_value	2017-01-25 12:39:31.096831+00
30	sentry	0029_auto__del_field_projectmember_is_superuser__del_field_projectmember_pe	2017-01-25 12:39:31.197657+00
31	sentry	0030_auto__add_view__chg_field_event_group	2017-01-25 12:39:31.287993+00
32	sentry	0031_auto__add_field_view_verbose_name__add_field_view_verbose_name_plural_	2017-01-25 12:39:31.3225+00
33	sentry	0032_auto__add_eventmeta	2017-01-25 12:39:31.400188+00
34	sentry	0033_auto__add_option__add_unique_option_key_value	2017-01-25 12:39:31.558527+00
35	sentry	0034_auto__add_groupbookmark__add_unique_groupbookmark_project_user_group	2017-01-25 12:39:31.634979+00
36	sentry	0034_auto__add_unique_option_key__del_unique_option_value_key__del_unique_g	2017-01-25 12:39:31.707071+00
37	sentry	0036_auto__chg_field_option_value__chg_field_projectoption_value	2017-01-25 12:39:31.802196+00
38	sentry	0037_auto__add_unique_option_key__del_unique_option_value_key__del_unique_g	2017-01-25 12:39:31.851159+00
39	sentry	0038_auto__add_searchtoken__add_unique_searchtoken_document_field_token__ad	2017-01-25 12:39:31.946822+00
40	sentry	0039_auto__add_field_searchdocument_status	2017-01-25 12:39:32.009344+00
41	sentry	0040_auto__del_unique_event_event_id__add_unique_event_project_event_id	2017-01-25 12:39:32.047763+00
42	sentry	0041_auto__add_field_messagefiltervalue_last_seen__add_field_messagefilterv	2017-01-25 12:39:32.125643+00
43	sentry	0042_auto__add_projectcountbyminute__add_unique_projectcountbyminute_projec	2017-01-25 12:39:32.165172+00
44	sentry	0043_auto__chg_field_option_value__chg_field_projectoption_value	2017-01-25 12:39:32.234234+00
45	sentry	0044_auto__add_field_projectmember_is_active	2017-01-25 12:39:32.285751+00
46	sentry	0045_auto__add_pendingprojectmember__add_unique_pendingprojectmember_projec	2017-01-25 12:39:32.3314+00
47	sentry	0046_auto__add_teammember__add_unique_teammember_team_user__add_team__add_p	2017-01-25 12:39:32.454451+00
48	sentry	0047_migrate_project_slugs	2017-01-25 12:39:32.487549+00
49	sentry	0048_migrate_project_keys	2017-01-25 12:39:32.519923+00
50	sentry	0049_create_default_project_keys	2017-01-25 12:39:32.548941+00
51	sentry	0050_remove_project_keys_from_members	2017-01-25 12:39:32.588502+00
52	sentry	0051_auto__del_pendingprojectmember__del_unique_pendingprojectmember_projec	2017-01-25 12:39:32.639839+00
53	sentry	0052_migrate_project_members	2017-01-25 12:39:32.669325+00
54	sentry	0053_auto__del_projectmember__del_unique_projectmember_project_user	2017-01-25 12:39:32.70815+00
55	sentry	0054_fix_project_keys	2017-01-25 12:39:32.736717+00
56	sentry	0055_auto__del_projectdomain__del_unique_projectdomain_project_domain	2017-01-25 12:39:32.770903+00
57	sentry	0056_auto__add_field_group_resolved_at	2017-01-25 12:39:32.805259+00
58	sentry	0057_auto__add_field_group_active_at	2017-01-25 12:39:32.845048+00
59	sentry	0058_auto__add_useroption__add_unique_useroption_user_project_key	2017-01-25 12:39:32.896867+00
60	sentry	0059_auto__add_filterkey__add_unique_filterkey_project_key	2017-01-25 12:39:32.941472+00
61	sentry	0060_fill_filter_key	2017-01-25 12:39:32.975207+00
62	sentry	0061_auto__add_field_group_group_id__add_field_group_is_public	2017-01-25 12:39:34.333368+00
63	sentry	0062_correct_del_index_sentry_groupedmessage_logger__view__checksum	2017-01-25 12:39:34.368405+00
64	sentry	0063_auto	2017-01-25 12:39:34.416689+00
65	sentry	0064_index_checksum	2017-01-25 12:39:34.59494+00
66	sentry	0065_create_default_project_key	2017-01-25 12:39:34.64182+00
67	sentry	0066_auto__del_view	2017-01-25 12:39:34.686373+00
68	sentry	0067_auto__add_field_group_platform__add_field_event_platform	2017-01-25 12:39:34.721618+00
69	sentry	0068_auto__add_field_projectkey_user_added__add_field_projectkey_date_added	2017-01-25 12:39:35.451674+00
70	sentry	0069_auto__add_lostpasswordhash	2017-01-25 12:39:35.510952+00
71	sentry	0070_projectoption_key_length	2017-01-25 12:39:35.557844+00
72	sentry	0071_auto__add_field_group_users_seen	2017-01-25 12:39:37.728643+00
73	sentry	0072_auto__add_affecteduserbygroup__add_unique_affecteduserbygroup_project_	2017-01-25 12:39:38.450819+00
74	sentry	0073_auto__add_field_project_platform	2017-01-25 12:39:38.48788+00
75	sentry	0074_correct_filtervalue_index	2017-01-25 12:39:38.54537+00
76	sentry	0075_add_groupbookmark_index	2017-01-25 12:39:38.595759+00
77	sentry	0076_add_groupmeta_index	2017-01-25 12:39:38.650416+00
78	sentry	0077_auto__add_trackeduser__add_unique_trackeduser_project_ident	2017-01-25 12:39:38.762339+00
79	sentry	0078_auto__add_field_affecteduserbygroup_tuser	2017-01-25 12:39:38.816998+00
80	sentry	0079_auto__del_unique_affecteduserbygroup_project_ident_group__add_unique_a	2017-01-25 12:39:40.208873+00
81	sentry	0080_auto__chg_field_affecteduserbygroup_ident	2017-01-25 12:39:40.263133+00
82	sentry	0081_fill_trackeduser	2017-01-25 12:39:40.306333+00
83	sentry	0082_auto__add_activity__add_field_group_num_comments__add_field_event_num_	2017-01-25 12:39:45.877914+00
84	sentry	0083_migrate_dupe_groups	2017-01-25 12:39:45.92332+00
85	sentry	0084_auto__del_unique_group_project_checksum_logger_culprit__add_unique_gro	2017-01-25 12:39:46.615935+00
86	sentry	0085_auto__del_unique_project_slug__add_unique_project_slug_team	2017-01-25 12:39:46.687569+00
87	sentry	0086_auto__add_field_team_date_added	2017-01-25 12:39:46.781127+00
88	sentry	0087_auto__del_messagefiltervalue__del_unique_messagefiltervalue_project_ke	2017-01-25 12:39:46.831602+00
89	sentry	0088_auto__del_messagecountbyminute__del_unique_messagecountbyminute_projec	2017-01-25 12:39:46.871429+00
90	sentry	0089_auto__add_accessgroup__add_unique_accessgroup_team_name	2017-01-25 12:39:47.070488+00
91	sentry	0090_auto__add_grouptagkey__add_unique_grouptagkey_project_group_key__add_f	2017-01-25 12:39:47.99184+00
92	sentry	0091_auto__add_alert	2017-01-25 12:39:48.90711+00
93	sentry	0092_auto__add_alertrelatedgroup__add_unique_alertrelatedgroup_group_alert	2017-01-25 12:39:49.020183+00
94	sentry	0093_auto__add_field_alert_status	2017-01-25 12:39:50.932571+00
95	sentry	0094_auto__add_eventmapping__add_unique_eventmapping_project_event_id	2017-01-25 12:39:51.029036+00
96	sentry	0095_rebase	2017-01-25 12:39:51.084024+00
97	sentry	0096_auto__add_field_tagvalue_data	2017-01-25 12:39:51.143885+00
98	sentry	0097_auto__del_affecteduserbygroup__del_unique_affecteduserbygroup_project_	2017-01-25 12:39:51.205585+00
99	sentry	0098_auto__add_user__chg_field_team_owner__chg_field_activity_user__chg_fie	2017-01-25 12:39:51.260218+00
100	sentry	0099_auto__del_field_teammember_is_active	2017-01-25 12:39:51.313534+00
101	sentry	0100_auto__add_field_tagkey_label	2017-01-25 12:39:51.367894+00
102	sentry	0101_ensure_teams	2017-01-25 12:39:51.416936+00
103	sentry	0102_ensure_slugs	2017-01-25 12:39:51.467745+00
104	sentry	0103_ensure_non_empty_slugs	2017-01-25 12:39:51.526192+00
105	sentry	0104_auto__add_groupseen__add_unique_groupseen_group_user	2017-01-25 12:39:51.619946+00
106	sentry	0105_auto__chg_field_projectcountbyminute_time_spent_total__chg_field_group	2017-01-25 12:39:54.542053+00
107	sentry	0106_auto__del_searchtoken__del_unique_searchtoken_document_field_token__de	2017-01-25 12:39:54.595395+00
108	sentry	0107_expand_user	2017-01-25 12:39:54.659521+00
109	sentry	0108_fix_user	2017-01-25 12:39:54.712857+00
110	sentry	0109_index_filtervalue_times_seen	2017-01-25 12:39:54.765033+00
111	sentry	0110_index_filtervalue_last_seen	2017-01-25 12:39:54.818098+00
112	sentry	0111_index_filtervalue_first_seen	2017-01-25 12:39:54.869192+00
113	sentry	0112_auto__chg_field_option_value__chg_field_useroption_value__chg_field_pr	2017-01-25 12:39:54.913787+00
114	sentry	0113_auto__add_field_team_status	2017-01-25 12:39:55.609867+00
115	sentry	0114_auto__add_field_projectkey_roles	2017-01-25 12:39:55.732496+00
116	sentry	0115_auto__del_projectcountbyminute__del_unique_projectcountbyminute_projec	2017-01-25 12:39:55.776215+00
117	sentry	0116_auto__del_field_event_server_name__del_field_event_culprit__del_field_	2017-01-25 12:39:55.819515+00
118	sentry	0117_auto__add_rule	2017-01-25 12:39:55.888667+00
119	sentry	0118_create_default_rules	2017-01-25 12:39:56.137457+00
120	sentry	0119_auto__add_field_projectkey_label	2017-01-25 12:39:56.178892+00
121	sentry	0120_auto__add_grouprulestatus	2017-01-25 12:39:56.278444+00
122	sentry	0121_auto__add_unique_grouprulestatus_rule_group	2017-01-25 12:39:56.330455+00
123	sentry	0122_add_event_group_id_datetime_index	2017-01-25 12:39:56.386288+00
124	sentry	0123_auto__add_groupassignee__add_index_event_group_datetime	2017-01-25 12:39:56.469493+00
125	sentry	0124_auto__add_grouphash__add_unique_grouphash_project_hash	2017-01-25 12:39:56.576757+00
126	sentry	0125_auto__add_field_user_is_managed	2017-01-25 12:39:56.66063+00
127	sentry	0126_auto__add_field_option_last_updated	2017-01-25 12:39:57.354592+00
128	sentry	0127_auto__add_release__add_unique_release_project_version	2017-01-25 12:39:57.435379+00
129	sentry	0128_auto__add_broadcast	2017-01-25 12:39:57.510109+00
130	sentry	0129_auto__chg_field_release_id__chg_field_pendingteammember_id__chg_field_	2017-01-25 12:39:57.563307+00
131	sentry	0130_auto__del_field_project_owner	2017-01-25 12:39:57.612897+00
132	sentry	0131_auto__add_organizationmember__add_unique_organizationmember_organizati	2017-01-25 12:39:58.337131+00
133	sentry	0132_add_default_orgs	2017-01-25 12:39:58.393903+00
134	sentry	0133_add_org_members	2017-01-25 12:39:58.451328+00
135	sentry	0134_auto__chg_field_team_organization	2017-01-25 12:39:58.524981+00
136	sentry	0135_auto__chg_field_project_team	2017-01-25 12:39:58.594679+00
137	sentry	0136_auto__add_field_organizationmember_email__chg_field_organizationmember	2017-01-25 12:39:58.674942+00
138	sentry	0137_auto__add_field_organizationmember_has_global_access	2017-01-25 12:39:58.831889+00
139	sentry	0138_migrate_team_members	2017-01-25 12:39:58.890282+00
140	sentry	0139_auto__add_auditlogentry	2017-01-25 12:39:59.012725+00
141	sentry	0140_auto__add_field_organization_slug	2017-01-25 12:39:59.091296+00
142	sentry	0141_fill_org_slugs	2017-01-25 12:39:59.148066+00
143	sentry	0142_auto__add_field_project_organization__add_unique_project_organization_	2017-01-25 12:39:59.234813+00
144	sentry	0143_fill_project_orgs	2017-01-25 12:39:59.295471+00
145	sentry	0144_auto__chg_field_project_organization	2017-01-25 12:39:59.369822+00
146	sentry	0145_auto__chg_field_organization_slug	2017-01-25 12:39:59.446583+00
147	sentry	0146_auto__add_field_auditlogentry_ip_address	2017-01-25 12:39:59.512095+00
148	sentry	0147_auto__del_unique_team_slug__add_unique_team_organization_slug	2017-01-25 12:39:59.586591+00
149	sentry	0148_auto__add_helppage	2017-01-25 12:39:59.686825+00
150	sentry	0149_auto__chg_field_groupseen_project__chg_field_groupseen_user__chg_field	2017-01-25 12:40:00.037507+00
151	sentry	0150_fix_broken_rules	2017-01-25 12:40:00.098355+00
152	sentry	0151_auto__add_file	2017-01-25 12:40:00.193994+00
153	sentry	0152_auto__add_field_file_checksum__chg_field_file_name__add_unique_file_na	2017-01-25 12:40:01.528519+00
154	sentry	0153_auto__add_field_grouprulestatus_last_active	2017-01-25 12:40:01.596116+00
155	sentry	0154_auto__add_field_tagkey_status	2017-01-25 12:40:01.697453+00
156	sentry	0155_auto__add_field_projectkey_status	2017-01-25 12:40:01.850523+00
157	sentry	0156_auto__add_apikey	2017-01-25 12:40:01.965777+00
158	sentry	0157_auto__add_authidentity__add_unique_authidentity_auth_provider_ident__a	2017-01-25 12:40:02.114317+00
159	sentry	0158_auto__add_unique_authidentity_auth_provider_user	2017-01-25 12:40:02.200803+00
160	sentry	0159_auto__add_field_authidentity_last_verified__add_field_organizationmemb	2017-01-25 12:40:03.605904+00
161	sentry	0160_auto__add_field_authprovider_default_global_access	2017-01-25 12:40:04.352723+00
162	sentry	0161_auto__chg_field_authprovider_config	2017-01-25 12:40:04.448794+00
163	sentry	0162_auto__chg_field_authidentity_data	2017-01-25 12:40:04.538975+00
164	sentry	0163_auto__add_field_authidentity_last_synced	2017-01-25 12:40:07.203553+00
165	sentry	0164_auto__add_releasefile__add_unique_releasefile_release_ident__add_field	2017-01-25 12:40:08.608966+00
166	sentry	0165_auto__del_unique_file_name_checksum	2017-01-25 12:40:08.701966+00
167	sentry	0166_auto__chg_field_user_id__add_field_apikey_allowed_origins	2017-01-25 12:40:08.791211+00
168	sentry	0167_auto__add_field_authprovider_flags	2017-01-25 12:40:10.112071+00
169	sentry	0168_unfill_projectkey_user	2017-01-25 12:40:10.192129+00
170	sentry	0169_auto__del_field_projectkey_user	2017-01-25 12:40:10.273717+00
171	sentry	0170_auto__add_organizationmemberteam__add_unique_organizationmemberteam_te	2017-01-25 12:40:10.481411+00
172	sentry	0171_auto__chg_field_team_owner	2017-01-25 12:40:10.58223+00
173	sentry	0172_auto__del_field_team_owner	2017-01-25 12:40:10.661053+00
174	sentry	0173_auto__del_teammember__del_unique_teammember_team_user	2017-01-25 12:40:10.75332+00
175	sentry	0174_auto__del_field_projectkey_user_added	2017-01-25 12:40:10.834312+00
176	sentry	0175_auto__del_pendingteammember__del_unique_pendingteammember_team_email	2017-01-25 12:40:10.923194+00
177	sentry	0176_auto__add_field_organizationmember_counter__add_unique_organizationmem	2017-01-25 12:40:11.019911+00
178	sentry	0177_fill_member_counters	2017-01-25 12:40:11.096453+00
179	sentry	0178_auto__del_unique_organizationmember_organization_counter	2017-01-25 12:40:11.184062+00
180	sentry	0179_auto__add_field_release_date_released	2017-01-25 12:40:11.598891+00
181	sentry	0180_auto__add_field_release_environment__add_field_release_ref__add_field_	2017-01-25 12:40:13.555912+00
182	sentry	0181_auto__del_field_release_environment__del_unique_release_project_versio	2017-01-25 12:40:13.654187+00
183	sentry	0182_auto__add_field_auditlogentry_actor_label__add_field_auditlogentry_act	2017-01-25 12:40:13.768142+00
184	sentry	0183_auto__del_index_grouphash_hash	2017-01-25 12:40:13.853011+00
185	sentry	0184_auto__del_field_group_checksum__del_unique_group_project_checksum__del	2017-01-25 12:40:13.947099+00
186	sentry	0185_auto__add_savedsearch__add_unique_savedsearch_project_name	2017-01-25 12:40:14.076803+00
187	sentry	0186_auto__add_field_group_first_release	2017-01-25 12:40:14.170125+00
188	sentry	0187_auto__add_index_group_project_first_release	2017-01-25 12:40:14.271826+00
189	sentry	0188_auto__add_userreport	2017-01-25 12:40:14.396823+00
190	sentry	0189_auto__add_index_userreport_project_event_id	2017-01-25 12:40:14.494372+00
191	sentry	0190_auto__add_field_release_new_groups	2017-01-25 12:40:15.809749+00
192	sentry	0191_auto__del_alert__del_alertrelatedgroup__del_unique_alertrelatedgroup_g	2017-01-25 12:40:15.90642+00
193	sentry	0192_add_model_groupemailthread	2017-01-25 12:40:16.055107+00
194	sentry	0193_auto__del_unique_groupemailthread_msgid__add_unique_groupemailthread_e	2017-01-25 12:40:16.161407+00
195	sentry	0194_auto__del_field_project_platform	2017-01-25 12:40:16.251957+00
196	sentry	0195_auto__chg_field_organization_owner	2017-01-25 12:40:16.358183+00
197	sentry	0196_auto__del_field_organization_owner	2017-01-25 12:40:16.44842+00
198	sentry	0197_auto__del_accessgroup__del_unique_accessgroup_team_name	2017-01-25 12:40:16.541756+00
199	sentry	0198_auto__add_field_release_primary_owner	2017-01-25 12:40:16.638051+00
200	sentry	0199_auto__add_field_project_first_event	2017-01-25 12:40:16.726178+00
201	sentry	0200_backfill_first_event	2017-01-25 12:40:16.811373+00
202	sentry	0201_auto__add_eventuser__add_unique_eventuser_project_ident__add_index_eve	2017-01-25 12:40:16.967804+00
203	sentry	0202_auto__add_field_eventuser_hash__add_unique_eventuser_project_hash	2017-01-25 12:40:17.06651+00
204	sentry	0203_auto__chg_field_eventuser_username__chg_field_eventuser_ident	2017-01-25 12:40:17.218367+00
205	sentry	0204_backfill_team_membership	2017-01-25 12:40:17.337003+00
206	sentry	0205_auto__add_field_organizationmember_role	2017-01-25 12:40:17.494813+00
207	sentry	0206_backfill_member_role	2017-01-25 12:40:17.606286+00
208	sentry	0207_auto__add_field_organization_default_role	2017-01-25 12:40:17.75028+00
209	sentry	0208_backfill_default_role	2017-01-25 12:40:17.846074+00
210	sentry	0209_auto__add_broadcastseen__add_unique_broadcastseen_broadcast_user	2017-01-25 12:40:18.578565+00
211	sentry	0210_auto__del_field_broadcast_badge	2017-01-25 12:40:19.277918+00
212	sentry	0211_auto__add_field_broadcast_title	2017-01-25 12:40:19.398792+00
213	sentry	0212_auto__add_fileblob__add_field_file_blob	2017-01-25 12:40:19.560174+00
214	sentry	0212_auto__add_organizationoption__add_unique_organizationoption_organizati	2017-01-25 12:40:20.304871+00
215	sentry	0213_migrate_file_blobs	2017-01-25 12:40:20.398765+00
216	sentry	0214_auto__add_field_broadcast_upstream_id	2017-01-25 12:40:20.504942+00
217	sentry	0215_auto__add_field_broadcast_date_expires	2017-01-25 12:40:21.65666+00
218	sentry	0216_auto__add_groupsnooze	2017-01-25 12:40:21.779758+00
219	sentry	0217_auto__add_groupresolution	2017-01-25 12:40:21.915392+00
220	sentry	0218_auto__add_field_groupresolution_status	2017-01-25 12:40:22.070284+00
221	sentry	0219_auto__add_field_groupbookmark_date_added	2017-01-25 12:40:22.232108+00
222	sentry	0220_auto__del_field_fileblob_storage_options__del_field_fileblob_storage__	2017-01-25 12:40:22.329987+00
223	sentry	0221_auto__chg_field_user_first_name	2017-01-25 12:40:22.45587+00
224	sentry	0222_auto__del_field_user_last_name__del_field_user_first_name__add_field_u	2017-01-25 12:40:22.563717+00
225	sentry	0223_delete_old_sentry_docs_options	2017-01-25 12:40:22.664797+00
226	sentry	0224_auto__add_index_userreport_project_date_added	2017-01-25 12:40:22.784158+00
227	sentry	0225_auto__add_fileblobindex__add_unique_fileblobindex_file_blob_offset	2017-01-25 12:40:24.113261+00
228	sentry	0226_backfill_file_size	2017-01-25 12:40:24.21532+00
229	sentry	0227_auto__del_field_activity_event	2017-01-25 12:40:24.315124+00
230	sentry	0228_auto__del_field_event_num_comments	2017-01-25 12:40:24.416724+00
231	sentry	0229_drop_event_constraints	2017-01-25 12:40:24.567216+00
232	sentry	0230_auto__del_field_eventmapping_group__del_field_eventmapping_project__ad	2017-01-25 12:40:24.667522+00
233	sentry	0231_auto__add_field_savedsearch_is_default	2017-01-25 12:40:25.41591+00
234	sentry	0232_default_savedsearch	2017-01-25 12:40:25.531132+00
235	sentry	0233_add_new_savedsearch	2017-01-25 12:40:25.632666+00
236	sentry	0234_auto__add_savedsearchuserdefault__add_unique_savedsearchuserdefault_pr	2017-01-25 12:40:25.79357+00
237	sentry	0235_auto__add_projectbookmark__add_unique_projectbookmark_project_id_user_	2017-01-25 12:40:25.9391+00
238	sentry	0236_auto__add_organizationonboardingtask__add_unique_organizationonboardin	2017-01-25 12:40:26.117929+00
239	sentry	0237_auto__add_eventtag__add_unique_eventtag_event_id_key_id_value_id	2017-01-25 12:40:26.26646+00
240	sentry	0238_fill_org_onboarding_tasks	2017-01-25 12:40:26.388372+00
241	sentry	0239_auto__add_projectdsymfile__add_unique_projectdsymfile_project_uuid__ad	2017-01-25 12:40:26.625441+00
242	sentry	0240_fill_onboarding_option	2017-01-25 12:40:26.76199+00
243	sentry	0241_auto__add_counter__add_unique_counter_project_ident__add_field_group_s	2017-01-25 12:40:26.915582+00
244	sentry	0242_auto__add_field_project_forced_color	2017-01-25 12:40:27.045804+00
245	sentry	0243_remove_inactive_members	2017-01-25 12:40:27.167778+00
246	sentry	0244_auto__add_groupredirect	2017-01-25 12:40:27.323555+00
247	sentry	0245_auto__del_field_project_callsign__del_unique_project_organization_call	2017-01-25 12:40:27.446106+00
248	sentry	0246_auto__add_dsymsymbol__add_unique_dsymsymbol_object_address__add_dsymsd	2017-01-25 12:40:27.754957+00
249	sentry	0247_migrate_file_blobs	2017-01-25 12:40:27.892386+00
250	sentry	0248_auto__add_projectplatform__add_unique_projectplatform_project_id_platf	2017-01-25 12:40:28.053611+00
251	sentry	0249_auto__add_index_eventtag_project_id_key_id_value_id	2017-01-25 12:40:28.201022+00
252	sentry	0250_auto__add_unique_userreport_project_event_id	2017-01-25 12:40:28.351191+00
253	sentry	0251_auto__add_useravatar	2017-01-25 12:40:28.538327+00
254	sentry	0252_default_users_to_gravatar	2017-01-25 12:40:28.682648+00
255	sentry	0253_auto__add_field_eventtag_group_id	2017-01-25 12:40:28.830379+00
256	sentry	0254_auto__add_index_eventtag_group_id_key_id_value_id	2017-01-25 12:40:29.503968+00
257	sentry	0255_auto__add_apitoken	2017-01-25 12:40:29.665732+00
258	sentry	0256_auto__add_authenticator	2017-01-25 12:40:29.824156+00
259	sentry	0257_repair_activity	2017-01-25 12:40:29.969313+00
260	sentry	0258_auto__add_field_user_is_password_expired__add_field_user_last_password	2017-01-25 12:40:31.329978+00
261	sentry	0259_auto__add_useremail__add_unique_useremail_user_email	2017-01-25 12:40:31.518887+00
262	sentry	0260_populate_email_addresses	2017-01-25 12:40:31.690048+00
263	sentry	0261_auto__add_groupsubscription__add_unique_groupsubscription_group_user	2017-01-25 12:40:31.917088+00
264	sentry	0262_fix_tag_indexes	2017-01-25 12:40:32.737657+00
265	sentry	0263_remove_default_regression_rule	2017-01-25 12:40:32.918985+00
266	sentry	0264_drop_grouptagvalue_project_index	2017-01-25 12:40:33.096681+00
267	sentry	0265_auto__add_field_rule_status	2017-01-25 12:40:33.925103+00
268	sentry	0266_auto__add_grouprelease__add_unique_grouprelease_group_id_release_id_en	2017-01-25 12:40:34.154977+00
269	sentry	0267_auto__add_environment__add_unique_environment_project_id_name__add_rel	2017-01-25 12:40:34.422709+00
270	sentry	0268_fill_environment	2017-01-25 12:40:34.610172+00
271	sentry	0269_auto__del_helppage	2017-01-25 12:40:34.791227+00
272	sentry	0270_auto__add_field_organizationmember_token	2017-01-25 12:40:35.582827+00
273	sentry	0271_auto__del_field_organizationmember_counter	2017-01-25 12:40:35.791474+00
274	sentry	0272_auto__add_unique_authenticator_user_type	2017-01-25 12:40:36.001197+00
275	sentry	0273_auto__add_repository__add_unique_repository_organization_id_name__add_	2017-01-25 12:40:36.339176+00
276	sentry	0274_auto__add_index_commit_repository_id_date_added	2017-01-25 12:40:36.552698+00
277	sentry	0275_auto__del_index_grouptagvalue_project_key_value__add_index_grouptagval	2017-01-25 12:40:36.771758+00
278	sentry	0276_auto__add_field_user_session_nonce	2017-01-25 12:40:36.984593+00
279	sentry	0277_auto__add_commitfilechange__add_unique_commitfilechange_commit_filenam	2017-01-25 12:40:39.735511+00
280	sentry	0278_auto__add_releaseproject__add_unique_releaseproject_project_release__a	2017-01-25 12:40:39.973687+00
281	sentry	0279_populate_release_orgs_and_projects	2017-01-25 12:40:40.188264+00
282	sentry	0280_auto__add_field_releasecommit_organization_id	2017-01-25 12:40:40.390908+00
283	sentry	0281_populate_release_commit_organization_id	2017-01-25 12:40:40.580955+00
284	sentry	0282_auto__add_field_releasefile_organization__add_field_releaseenvironment	2017-01-25 12:40:40.797139+00
285	sentry	0283_populate_release_environment_and_release_file_organization	2017-01-25 12:40:40.99385+00
286	nodestore	0001_initial	2017-01-25 12:40:42.306089+00
287	search	0001_initial	2017-01-25 12:40:42.525451+00
288	search	0002_auto__del_searchtoken__del_unique_searchtoken_document_field_token__de	2017-01-25 12:40:42.557345+00
289	social_auth	0001_initial	2017-01-25 12:40:42.842399+00
290	social_auth	0002_auto__add_unique_nonce_timestamp_salt_server_url__add_unique_associati	2017-01-25 12:40:42.968771+00
291	social_auth	0003_auto__del_nonce__del_unique_nonce_server_url_timestamp_salt__del_assoc	2017-01-25 12:40:43.002916+00
292	social_auth	0004_auto__del_unique_usersocialauth_provider_uid__add_unique_usersocialaut	2017-01-25 12:40:43.033069+00
\.


--
-- Name: south_migrationhistory_id_seq; Type: SEQUENCE SET; Schema: public; Owner: sentry
--

SELECT pg_catalog.setval('south_migrationhistory_id_seq', 292, true);


--
-- Name: auth_authenticator auth_authenticator_pkey; Type: CONSTRAINT; Schema: public; Owner: sentry
--

ALTER TABLE ONLY auth_authenticator
    ADD CONSTRAINT auth_authenticator_pkey PRIMARY KEY (id);


--
-- Name: auth_authenticator auth_authenticator_user_id_5774ed51577668d4_uniq; Type: CONSTRAINT; Schema: public; Owner: sentry
--

ALTER TABLE ONLY auth_authenticator
    ADD CONSTRAINT auth_authenticator_user_id_5774ed51577668d4_uniq UNIQUE (user_id, type);


--
-- Name: auth_group auth_group_name_key; Type: CONSTRAINT; Schema: public; Owner: sentry
--

ALTER TABLE ONLY auth_group
    ADD CONSTRAINT auth_group_name_key UNIQUE (name);


--
-- Name: auth_group_permissions auth_group_permissions_group_id_permission_id_key; Type: CONSTRAINT; Schema: public; Owner: sentry
--

ALTER TABLE ONLY auth_group_permissions
    ADD CONSTRAINT auth_group_permissions_group_id_permission_id_key UNIQUE (group_id, permission_id);


--
-- Name: auth_group_permissions auth_group_permissions_pkey; Type: CONSTRAINT; Schema: public; Owner: sentry
--

ALTER TABLE ONLY auth_group_permissions
    ADD CONSTRAINT auth_group_permissions_pkey PRIMARY KEY (id);


--
-- Name: auth_group auth_group_pkey; Type: CONSTRAINT; Schema: public; Owner: sentry
--

ALTER TABLE ONLY auth_group
    ADD CONSTRAINT auth_group_pkey PRIMARY KEY (id);


--
-- Name: auth_permission auth_permission_content_type_id_codename_key; Type: CONSTRAINT; Schema: public; Owner: sentry
--

ALTER TABLE ONLY auth_permission
    ADD CONSTRAINT auth_permission_content_type_id_codename_key UNIQUE (content_type_id, codename);


--
-- Name: auth_permission auth_permission_pkey; Type: CONSTRAINT; Schema: public; Owner: sentry
--

ALTER TABLE ONLY auth_permission
    ADD CONSTRAINT auth_permission_pkey PRIMARY KEY (id);


--
-- Name: auth_user auth_user_pkey; Type: CONSTRAINT; Schema: public; Owner: sentry
--

ALTER TABLE ONLY auth_user
    ADD CONSTRAINT auth_user_pkey PRIMARY KEY (id);


--
-- Name: auth_user auth_user_username_key; Type: CONSTRAINT; Schema: public; Owner: sentry
--

ALTER TABLE ONLY auth_user
    ADD CONSTRAINT auth_user_username_key UNIQUE (username);


--
-- Name: django_admin_log django_admin_log_pkey; Type: CONSTRAINT; Schema: public; Owner: sentry
--

ALTER TABLE ONLY django_admin_log
    ADD CONSTRAINT django_admin_log_pkey PRIMARY KEY (id);


--
-- Name: django_content_type django_content_type_app_label_model_key; Type: CONSTRAINT; Schema: public; Owner: sentry
--

ALTER TABLE ONLY django_content_type
    ADD CONSTRAINT django_content_type_app_label_model_key UNIQUE (app_label, model);


--
-- Name: django_content_type django_content_type_pkey; Type: CONSTRAINT; Schema: public; Owner: sentry
--

ALTER TABLE ONLY django_content_type
    ADD CONSTRAINT django_content_type_pkey PRIMARY KEY (id);


--
-- Name: django_session django_session_pkey; Type: CONSTRAINT; Schema: public; Owner: sentry
--

ALTER TABLE ONLY django_session
    ADD CONSTRAINT django_session_pkey PRIMARY KEY (session_key);


--
-- Name: django_site django_site_pkey; Type: CONSTRAINT; Schema: public; Owner: sentry
--

ALTER TABLE ONLY django_site
    ADD CONSTRAINT django_site_pkey PRIMARY KEY (id);


--
-- Name: nodestore_node nodestore_node_pkey; Type: CONSTRAINT; Schema: public; Owner: sentry
--

ALTER TABLE ONLY nodestore_node
    ADD CONSTRAINT nodestore_node_pkey PRIMARY KEY (id);


--
-- Name: sentry_activity sentry_activity_pkey; Type: CONSTRAINT; Schema: public; Owner: sentry
--

ALTER TABLE ONLY sentry_activity
    ADD CONSTRAINT sentry_activity_pkey PRIMARY KEY (id);


--
-- Name: sentry_apikey sentry_apikey_key_key; Type: CONSTRAINT; Schema: public; Owner: sentry
--

ALTER TABLE ONLY sentry_apikey
    ADD CONSTRAINT sentry_apikey_key_key UNIQUE (key);


--
-- Name: sentry_apikey sentry_apikey_pkey; Type: CONSTRAINT; Schema: public; Owner: sentry
--

ALTER TABLE ONLY sentry_apikey
    ADD CONSTRAINT sentry_apikey_pkey PRIMARY KEY (id);


--
-- Name: sentry_apitoken sentry_apitoken_pkey; Type: CONSTRAINT; Schema: public; Owner: sentry
--

ALTER TABLE ONLY sentry_apitoken
    ADD CONSTRAINT sentry_apitoken_pkey PRIMARY KEY (id);


--
-- Name: sentry_apitoken sentry_apitoken_token_key; Type: CONSTRAINT; Schema: public; Owner: sentry
--

ALTER TABLE ONLY sentry_apitoken
    ADD CONSTRAINT sentry_apitoken_token_key UNIQUE (token);


--
-- Name: sentry_auditlogentry sentry_auditlogentry_pkey; Type: CONSTRAINT; Schema: public; Owner: sentry
--

ALTER TABLE ONLY sentry_auditlogentry
    ADD CONSTRAINT sentry_auditlogentry_pkey PRIMARY KEY (id);


--
-- Name: sentry_authidentity sentry_authidentity_auth_provider_id_2ac89deececdc9d7_uniq; Type: CONSTRAINT; Schema: public; Owner: sentry
--

ALTER TABLE ONLY sentry_authidentity
    ADD CONSTRAINT sentry_authidentity_auth_provider_id_2ac89deececdc9d7_uniq UNIQUE (auth_provider_id, user_id);


--
-- Name: sentry_authidentity sentry_authidentity_auth_provider_id_72ab4375ecd728ba_uniq; Type: CONSTRAINT; Schema: public; Owner: sentry
--

ALTER TABLE ONLY sentry_authidentity
    ADD CONSTRAINT sentry_authidentity_auth_provider_id_72ab4375ecd728ba_uniq UNIQUE (auth_provider_id, ident);


--
-- Name: sentry_authidentity sentry_authidentity_pkey; Type: CONSTRAINT; Schema: public; Owner: sentry
--

ALTER TABLE ONLY sentry_authidentity
    ADD CONSTRAINT sentry_authidentity_pkey PRIMARY KEY (id);


--
-- Name: sentry_authprovider_default_teams sentry_authprovider_defau_authprovider_id_352ee7f2584f4caf_uniq; Type: CONSTRAINT; Schema: public; Owner: sentry
--

ALTER TABLE ONLY sentry_authprovider_default_teams
    ADD CONSTRAINT sentry_authprovider_defau_authprovider_id_352ee7f2584f4caf_uniq UNIQUE (authprovider_id, team_id);


--
-- Name: sentry_authprovider_default_teams sentry_authprovider_default_teams_pkey; Type: CONSTRAINT; Schema: public; Owner: sentry
--

ALTER TABLE ONLY sentry_authprovider_default_teams
    ADD CONSTRAINT sentry_authprovider_default_teams_pkey PRIMARY KEY (id);


--
-- Name: sentry_authprovider sentry_authprovider_organization_id_key; Type: CONSTRAINT; Schema: public; Owner: sentry
--

ALTER TABLE ONLY sentry_authprovider
    ADD CONSTRAINT sentry_authprovider_organization_id_key UNIQUE (organization_id);


--
-- Name: sentry_authprovider sentry_authprovider_pkey; Type: CONSTRAINT; Schema: public; Owner: sentry
--

ALTER TABLE ONLY sentry_authprovider
    ADD CONSTRAINT sentry_authprovider_pkey PRIMARY KEY (id);


--
-- Name: sentry_broadcast sentry_broadcast_pkey; Type: CONSTRAINT; Schema: public; Owner: sentry
--

ALTER TABLE ONLY sentry_broadcast
    ADD CONSTRAINT sentry_broadcast_pkey PRIMARY KEY (id);


--
-- Name: sentry_broadcastseen sentry_broadcastseen_broadcast_id_352c833420c70bd9_uniq; Type: CONSTRAINT; Schema: public; Owner: sentry
--

ALTER TABLE ONLY sentry_broadcastseen
    ADD CONSTRAINT sentry_broadcastseen_broadcast_id_352c833420c70bd9_uniq UNIQUE (broadcast_id, user_id);


--
-- Name: sentry_broadcastseen sentry_broadcastseen_pkey; Type: CONSTRAINT; Schema: public; Owner: sentry
--

ALTER TABLE ONLY sentry_broadcastseen
    ADD CONSTRAINT sentry_broadcastseen_pkey PRIMARY KEY (id);


--
-- Name: sentry_commit sentry_commit_pkey; Type: CONSTRAINT; Schema: public; Owner: sentry
--

ALTER TABLE ONLY sentry_commit
    ADD CONSTRAINT sentry_commit_pkey PRIMARY KEY (id);


--
-- Name: sentry_commit sentry_commit_repository_id_2d25b4d8949fca93_uniq; Type: CONSTRAINT; Schema: public; Owner: sentry
--

ALTER TABLE ONLY sentry_commit
    ADD CONSTRAINT sentry_commit_repository_id_2d25b4d8949fca93_uniq UNIQUE (repository_id, key);


--
-- Name: sentry_commitauthor sentry_commitauthor_organization_id_5656e6a6baa5f6c_uniq; Type: CONSTRAINT; Schema: public; Owner: sentry
--

ALTER TABLE ONLY sentry_commitauthor
    ADD CONSTRAINT sentry_commitauthor_organization_id_5656e6a6baa5f6c_uniq UNIQUE (organization_id, email);


--
-- Name: sentry_commitauthor sentry_commitauthor_pkey; Type: CONSTRAINT; Schema: public; Owner: sentry
--

ALTER TABLE ONLY sentry_commitauthor
    ADD CONSTRAINT sentry_commitauthor_pkey PRIMARY KEY (id);


--
-- Name: sentry_commitfilechange sentry_commitfilechange_commit_id_4c6f7ec25af34227_uniq; Type: CONSTRAINT; Schema: public; Owner: sentry
--

ALTER TABLE ONLY sentry_commitfilechange
    ADD CONSTRAINT sentry_commitfilechange_commit_id_4c6f7ec25af34227_uniq UNIQUE (commit_id, filename);


--
-- Name: sentry_commitfilechange sentry_commitfilechange_pkey; Type: CONSTRAINT; Schema: public; Owner: sentry
--

ALTER TABLE ONLY sentry_commitfilechange
    ADD CONSTRAINT sentry_commitfilechange_pkey PRIMARY KEY (id);


--
-- Name: sentry_dsymbundle sentry_dsymbundle_pkey; Type: CONSTRAINT; Schema: public; Owner: sentry
--

ALTER TABLE ONLY sentry_dsymbundle
    ADD CONSTRAINT sentry_dsymbundle_pkey PRIMARY KEY (id);


--
-- Name: sentry_dsymobject sentry_dsymobject_pkey; Type: CONSTRAINT; Schema: public; Owner: sentry
--

ALTER TABLE ONLY sentry_dsymobject
    ADD CONSTRAINT sentry_dsymobject_pkey PRIMARY KEY (id);


--
-- Name: sentry_dsymsdk sentry_dsymsdk_pkey; Type: CONSTRAINT; Schema: public; Owner: sentry
--

ALTER TABLE ONLY sentry_dsymsdk
    ADD CONSTRAINT sentry_dsymsdk_pkey PRIMARY KEY (id);


--
-- Name: sentry_dsymsymbol sentry_dsymsymbol_object_id_44d3aa1fdc2e5b55_uniq; Type: CONSTRAINT; Schema: public; Owner: sentry
--

ALTER TABLE ONLY sentry_dsymsymbol
    ADD CONSTRAINT sentry_dsymsymbol_object_id_44d3aa1fdc2e5b55_uniq UNIQUE (object_id, address);


--
-- Name: sentry_dsymsymbol sentry_dsymsymbol_pkey; Type: CONSTRAINT; Schema: public; Owner: sentry
--

ALTER TABLE ONLY sentry_dsymsymbol
    ADD CONSTRAINT sentry_dsymsymbol_pkey PRIMARY KEY (id);


--
-- Name: sentry_environment sentry_environment_pkey; Type: CONSTRAINT; Schema: public; Owner: sentry
--

ALTER TABLE ONLY sentry_environment
    ADD CONSTRAINT sentry_environment_pkey PRIMARY KEY (id);


--
-- Name: sentry_environment sentry_environment_project_id_1fbf3bfa87c819df_uniq; Type: CONSTRAINT; Schema: public; Owner: sentry
--

ALTER TABLE ONLY sentry_environment
    ADD CONSTRAINT sentry_environment_project_id_1fbf3bfa87c819df_uniq UNIQUE (project_id, name);


--
-- Name: sentry_environmentrelease sentry_environmentrelease_pkey; Type: CONSTRAINT; Schema: public; Owner: sentry
--

ALTER TABLE ONLY sentry_environmentrelease
    ADD CONSTRAINT sentry_environmentrelease_pkey PRIMARY KEY (id);


--
-- Name: sentry_environmentrelease sentry_environmentrelease_project_id_1762e5bd8b7a0675_uniq; Type: CONSTRAINT; Schema: public; Owner: sentry
--

ALTER TABLE ONLY sentry_environmentrelease
    ADD CONSTRAINT sentry_environmentrelease_project_id_1762e5bd8b7a0675_uniq UNIQUE (project_id, release_id, environment_id);


--
-- Name: sentry_eventmapping sentry_eventmapping_pkey; Type: CONSTRAINT; Schema: public; Owner: sentry
--

ALTER TABLE ONLY sentry_eventmapping
    ADD CONSTRAINT sentry_eventmapping_pkey PRIMARY KEY (id);


--
-- Name: sentry_eventmapping sentry_eventmapping_project_id_eb6c54bf8930ba6_uniq; Type: CONSTRAINT; Schema: public; Owner: sentry
--

ALTER TABLE ONLY sentry_eventmapping
    ADD CONSTRAINT sentry_eventmapping_project_id_eb6c54bf8930ba6_uniq UNIQUE (project_id, event_id);


--
-- Name: sentry_eventtag sentry_eventtag_event_id_430cef8ef4186908_uniq; Type: CONSTRAINT; Schema: public; Owner: sentry
--

ALTER TABLE ONLY sentry_eventtag
    ADD CONSTRAINT sentry_eventtag_event_id_430cef8ef4186908_uniq UNIQUE (event_id, key_id, value_id);


--
-- Name: sentry_eventtag sentry_eventtag_pkey; Type: CONSTRAINT; Schema: public; Owner: sentry
--

ALTER TABLE ONLY sentry_eventtag
    ADD CONSTRAINT sentry_eventtag_pkey PRIMARY KEY (id);


--
-- Name: sentry_eventuser sentry_eventuser_pkey; Type: CONSTRAINT; Schema: public; Owner: sentry
--

ALTER TABLE ONLY sentry_eventuser
    ADD CONSTRAINT sentry_eventuser_pkey PRIMARY KEY (id);


--
-- Name: sentry_eventuser sentry_eventuser_project_id_1a96e3b719e55f9a_uniq; Type: CONSTRAINT; Schema: public; Owner: sentry
--

ALTER TABLE ONLY sentry_eventuser
    ADD CONSTRAINT sentry_eventuser_project_id_1a96e3b719e55f9a_uniq UNIQUE (project_id, hash);


--
-- Name: sentry_eventuser sentry_eventuser_project_id_1dcb94833e2de5cf_uniq; Type: CONSTRAINT; Schema: public; Owner: sentry
--

ALTER TABLE ONLY sentry_eventuser
    ADD CONSTRAINT sentry_eventuser_project_id_1dcb94833e2de5cf_uniq UNIQUE (project_id, ident);


--
-- Name: sentry_file sentry_file_pkey; Type: CONSTRAINT; Schema: public; Owner: sentry
--

ALTER TABLE ONLY sentry_file
    ADD CONSTRAINT sentry_file_pkey PRIMARY KEY (id);


--
-- Name: sentry_fileblob sentry_fileblob_checksum_key; Type: CONSTRAINT; Schema: public; Owner: sentry
--

ALTER TABLE ONLY sentry_fileblob
    ADD CONSTRAINT sentry_fileblob_checksum_key UNIQUE (checksum);


--
-- Name: sentry_fileblob sentry_fileblob_pkey; Type: CONSTRAINT; Schema: public; Owner: sentry
--

ALTER TABLE ONLY sentry_fileblob
    ADD CONSTRAINT sentry_fileblob_pkey PRIMARY KEY (id);


--
-- Name: sentry_fileblobindex sentry_fileblobindex_file_id_56d11844195e33b2_uniq; Type: CONSTRAINT; Schema: public; Owner: sentry
--

ALTER TABLE ONLY sentry_fileblobindex
    ADD CONSTRAINT sentry_fileblobindex_file_id_56d11844195e33b2_uniq UNIQUE (file_id, blob_id, "offset");


--
-- Name: sentry_fileblobindex sentry_fileblobindex_pkey; Type: CONSTRAINT; Schema: public; Owner: sentry
--

ALTER TABLE ONLY sentry_fileblobindex
    ADD CONSTRAINT sentry_fileblobindex_pkey PRIMARY KEY (id);


--
-- Name: sentry_filterkey sentry_filterkey_pkey; Type: CONSTRAINT; Schema: public; Owner: sentry
--

ALTER TABLE ONLY sentry_filterkey
    ADD CONSTRAINT sentry_filterkey_pkey PRIMARY KEY (id);


--
-- Name: sentry_filterkey sentry_filterkey_project_id_67551b8e28dda5a_uniq; Type: CONSTRAINT; Schema: public; Owner: sentry
--

ALTER TABLE ONLY sentry_filterkey
    ADD CONSTRAINT sentry_filterkey_project_id_67551b8e28dda5a_uniq UNIQUE (project_id, key);


--
-- Name: sentry_filtervalue sentry_filtervalue_pkey; Type: CONSTRAINT; Schema: public; Owner: sentry
--

ALTER TABLE ONLY sentry_filtervalue
    ADD CONSTRAINT sentry_filtervalue_pkey PRIMARY KEY (id);


--
-- Name: sentry_filtervalue sentry_filtervalue_project_id_201b156195347397_uniq; Type: CONSTRAINT; Schema: public; Owner: sentry
--

ALTER TABLE ONLY sentry_filtervalue
    ADD CONSTRAINT sentry_filtervalue_project_id_201b156195347397_uniq UNIQUE (project_id, key, value);


--
-- Name: sentry_globaldsymfile sentry_globaldsymfile_pkey; Type: CONSTRAINT; Schema: public; Owner: sentry
--

ALTER TABLE ONLY sentry_globaldsymfile
    ADD CONSTRAINT sentry_globaldsymfile_pkey PRIMARY KEY (id);


--
-- Name: sentry_globaldsymfile sentry_globaldsymfile_uuid_key; Type: CONSTRAINT; Schema: public; Owner: sentry
--

ALTER TABLE ONLY sentry_globaldsymfile
    ADD CONSTRAINT sentry_globaldsymfile_uuid_key UNIQUE (uuid);


--
-- Name: sentry_groupasignee sentry_groupasignee_group_id_key; Type: CONSTRAINT; Schema: public; Owner: sentry
--

ALTER TABLE ONLY sentry_groupasignee
    ADD CONSTRAINT sentry_groupasignee_group_id_key UNIQUE (group_id);


--
-- Name: sentry_groupasignee sentry_groupasignee_pkey; Type: CONSTRAINT; Schema: public; Owner: sentry
--

ALTER TABLE ONLY sentry_groupasignee
    ADD CONSTRAINT sentry_groupasignee_pkey PRIMARY KEY (id);


--
-- Name: sentry_groupbookmark sentry_groupbookmark_pkey; Type: CONSTRAINT; Schema: public; Owner: sentry
--

ALTER TABLE ONLY sentry_groupbookmark
    ADD CONSTRAINT sentry_groupbookmark_pkey PRIMARY KEY (id);


--
-- Name: sentry_groupbookmark sentry_groupbookmark_project_id_6d2bb88ad3832208_uniq; Type: CONSTRAINT; Schema: public; Owner: sentry
--

ALTER TABLE ONLY sentry_groupbookmark
    ADD CONSTRAINT sentry_groupbookmark_project_id_6d2bb88ad3832208_uniq UNIQUE (project_id, user_id, group_id);


--
-- Name: sentry_groupedmessage sentry_groupedmessage_pkey; Type: CONSTRAINT; Schema: public; Owner: sentry
--

ALTER TABLE ONLY sentry_groupedmessage
    ADD CONSTRAINT sentry_groupedmessage_pkey PRIMARY KEY (id);


--
-- Name: sentry_groupedmessage sentry_groupedmessage_project_id_680bfe5607002523_uniq; Type: CONSTRAINT; Schema: public; Owner: sentry
--

ALTER TABLE ONLY sentry_groupedmessage
    ADD CONSTRAINT sentry_groupedmessage_project_id_680bfe5607002523_uniq UNIQUE (project_id, short_id);


--
-- Name: sentry_groupemailthread sentry_groupemailthread_email_456f4d17524b316_uniq; Type: CONSTRAINT; Schema: public; Owner: sentry
--

ALTER TABLE ONLY sentry_groupemailthread
    ADD CONSTRAINT sentry_groupemailthread_email_456f4d17524b316_uniq UNIQUE (email, msgid);


--
-- Name: sentry_groupemailthread sentry_groupemailthread_pkey; Type: CONSTRAINT; Schema: public; Owner: sentry
--

ALTER TABLE ONLY sentry_groupemailthread
    ADD CONSTRAINT sentry_groupemailthread_pkey PRIMARY KEY (id);


--
-- Name: sentry_grouphash sentry_grouphash_pkey; Type: CONSTRAINT; Schema: public; Owner: sentry
--

ALTER TABLE ONLY sentry_grouphash
    ADD CONSTRAINT sentry_grouphash_pkey PRIMARY KEY (id);


--
-- Name: sentry_grouphash sentry_grouphash_project_id_4a293f96a363c9a2_uniq; Type: CONSTRAINT; Schema: public; Owner: sentry
--

ALTER TABLE ONLY sentry_grouphash
    ADD CONSTRAINT sentry_grouphash_project_id_4a293f96a363c9a2_uniq UNIQUE (project_id, hash);


--
-- Name: sentry_groupmeta sentry_groupmeta_key_5d9d7a3c6538b14d_uniq; Type: CONSTRAINT; Schema: public; Owner: sentry
--

ALTER TABLE ONLY sentry_groupmeta
    ADD CONSTRAINT sentry_groupmeta_key_5d9d7a3c6538b14d_uniq UNIQUE (key, group_id);


--
-- Name: sentry_groupmeta sentry_groupmeta_pkey; Type: CONSTRAINT; Schema: public; Owner: sentry
--

ALTER TABLE ONLY sentry_groupmeta
    ADD CONSTRAINT sentry_groupmeta_pkey PRIMARY KEY (id);


--
-- Name: sentry_groupredirect sentry_groupredirect_pkey; Type: CONSTRAINT; Schema: public; Owner: sentry
--

ALTER TABLE ONLY sentry_groupredirect
    ADD CONSTRAINT sentry_groupredirect_pkey PRIMARY KEY (id);


--
-- Name: sentry_groupredirect sentry_groupredirect_previous_group_id_key; Type: CONSTRAINT; Schema: public; Owner: sentry
--

ALTER TABLE ONLY sentry_groupredirect
    ADD CONSTRAINT sentry_groupredirect_previous_group_id_key UNIQUE (previous_group_id);


--
-- Name: sentry_grouprelease sentry_grouprelease_group_id_46ba6e430d088d04_uniq; Type: CONSTRAINT; Schema: public; Owner: sentry
--

ALTER TABLE ONLY sentry_grouprelease
    ADD CONSTRAINT sentry_grouprelease_group_id_46ba6e430d088d04_uniq UNIQUE (group_id, release_id, environment);


--
-- Name: sentry_grouprelease sentry_grouprelease_pkey; Type: CONSTRAINT; Schema: public; Owner: sentry
--

ALTER TABLE ONLY sentry_grouprelease
    ADD CONSTRAINT sentry_grouprelease_pkey PRIMARY KEY (id);


--
-- Name: sentry_groupresolution sentry_groupresolution_group_id_key; Type: CONSTRAINT; Schema: public; Owner: sentry
--

ALTER TABLE ONLY sentry_groupresolution
    ADD CONSTRAINT sentry_groupresolution_group_id_key UNIQUE (group_id);


--
-- Name: sentry_groupresolution sentry_groupresolution_pkey; Type: CONSTRAINT; Schema: public; Owner: sentry
--

ALTER TABLE ONLY sentry_groupresolution
    ADD CONSTRAINT sentry_groupresolution_pkey PRIMARY KEY (id);


--
-- Name: sentry_grouprulestatus sentry_grouprulestatus_pkey; Type: CONSTRAINT; Schema: public; Owner: sentry
--

ALTER TABLE ONLY sentry_grouprulestatus
    ADD CONSTRAINT sentry_grouprulestatus_pkey PRIMARY KEY (id);


--
-- Name: sentry_grouprulestatus sentry_grouprulestatus_rule_id_329bb0edaad3880f_uniq; Type: CONSTRAINT; Schema: public; Owner: sentry
--

ALTER TABLE ONLY sentry_grouprulestatus
    ADD CONSTRAINT sentry_grouprulestatus_rule_id_329bb0edaad3880f_uniq UNIQUE (rule_id, group_id);


--
-- Name: sentry_groupseen sentry_groupseen_pkey; Type: CONSTRAINT; Schema: public; Owner: sentry
--

ALTER TABLE ONLY sentry_groupseen
    ADD CONSTRAINT sentry_groupseen_pkey PRIMARY KEY (id);


--
-- Name: sentry_groupseen sentry_groupseen_user_id_179917bc9974d91b_uniq; Type: CONSTRAINT; Schema: public; Owner: sentry
--

ALTER TABLE ONLY sentry_groupseen
    ADD CONSTRAINT sentry_groupseen_user_id_179917bc9974d91b_uniq UNIQUE (user_id, group_id);


--
-- Name: sentry_groupsnooze sentry_groupsnooze_group_id_key; Type: CONSTRAINT; Schema: public; Owner: sentry
--

ALTER TABLE ONLY sentry_groupsnooze
    ADD CONSTRAINT sentry_groupsnooze_group_id_key UNIQUE (group_id);


--
-- Name: sentry_groupsnooze sentry_groupsnooze_pkey; Type: CONSTRAINT; Schema: public; Owner: sentry
--

ALTER TABLE ONLY sentry_groupsnooze
    ADD CONSTRAINT sentry_groupsnooze_pkey PRIMARY KEY (id);


--
-- Name: sentry_groupsubscription sentry_groupsubscription_group_id_7e18bedd5058ccc3_uniq; Type: CONSTRAINT; Schema: public; Owner: sentry
--

ALTER TABLE ONLY sentry_groupsubscription
    ADD CONSTRAINT sentry_groupsubscription_group_id_7e18bedd5058ccc3_uniq UNIQUE (group_id, user_id);


--
-- Name: sentry_groupsubscription sentry_groupsubscription_pkey; Type: CONSTRAINT; Schema: public; Owner: sentry
--

ALTER TABLE ONLY sentry_groupsubscription
    ADD CONSTRAINT sentry_groupsubscription_pkey PRIMARY KEY (id);


--
-- Name: sentry_grouptagkey sentry_grouptagkey_pkey; Type: CONSTRAINT; Schema: public; Owner: sentry
--

ALTER TABLE ONLY sentry_grouptagkey
    ADD CONSTRAINT sentry_grouptagkey_pkey PRIMARY KEY (id);


--
-- Name: sentry_grouptagkey sentry_grouptagkey_project_id_7b0c8092f47b509f_uniq; Type: CONSTRAINT; Schema: public; Owner: sentry
--

ALTER TABLE ONLY sentry_grouptagkey
    ADD CONSTRAINT sentry_grouptagkey_project_id_7b0c8092f47b509f_uniq UNIQUE (project_id, group_id, key);


--
-- Name: sentry_lostpasswordhash sentry_lostpasswordhash_pkey; Type: CONSTRAINT; Schema: public; Owner: sentry
--

ALTER TABLE ONLY sentry_lostpasswordhash
    ADD CONSTRAINT sentry_lostpasswordhash_pkey PRIMARY KEY (id);


--
-- Name: sentry_lostpasswordhash sentry_lostpasswordhash_user_id_key; Type: CONSTRAINT; Schema: public; Owner: sentry
--

ALTER TABLE ONLY sentry_lostpasswordhash
    ADD CONSTRAINT sentry_lostpasswordhash_user_id_key UNIQUE (user_id);


--
-- Name: sentry_message sentry_message_pkey; Type: CONSTRAINT; Schema: public; Owner: sentry
--

ALTER TABLE ONLY sentry_message
    ADD CONSTRAINT sentry_message_pkey PRIMARY KEY (id);


--
-- Name: sentry_message sentry_message_project_id_b6b4e75e438ca83_uniq; Type: CONSTRAINT; Schema: public; Owner: sentry
--

ALTER TABLE ONLY sentry_message
    ADD CONSTRAINT sentry_message_project_id_b6b4e75e438ca83_uniq UNIQUE (project_id, message_id);


--
-- Name: sentry_messagefiltervalue sentry_messagefiltervalue_pkey; Type: CONSTRAINT; Schema: public; Owner: sentry
--

ALTER TABLE ONLY sentry_messagefiltervalue
    ADD CONSTRAINT sentry_messagefiltervalue_pkey PRIMARY KEY (id);


--
-- Name: sentry_messageindex sentry_messageindex_column_23431fca14e385c1_uniq; Type: CONSTRAINT; Schema: public; Owner: sentry
--

ALTER TABLE ONLY sentry_messageindex
    ADD CONSTRAINT sentry_messageindex_column_23431fca14e385c1_uniq UNIQUE ("column", value, object_id);


--
-- Name: sentry_messageindex sentry_messageindex_pkey; Type: CONSTRAINT; Schema: public; Owner: sentry
--

ALTER TABLE ONLY sentry_messageindex
    ADD CONSTRAINT sentry_messageindex_pkey PRIMARY KEY (id);


--
-- Name: sentry_option sentry_option_key_uniq; Type: CONSTRAINT; Schema: public; Owner: sentry
--

ALTER TABLE ONLY sentry_option
    ADD CONSTRAINT sentry_option_key_uniq UNIQUE (key);


--
-- Name: sentry_option sentry_option_pkey; Type: CONSTRAINT; Schema: public; Owner: sentry
--

ALTER TABLE ONLY sentry_option
    ADD CONSTRAINT sentry_option_pkey PRIMARY KEY (id);


--
-- Name: sentry_organizationmember_teams sentry_organization_organizationmember_id_1634015042409685_uniq; Type: CONSTRAINT; Schema: public; Owner: sentry
--

ALTER TABLE ONLY sentry_organizationmember_teams
    ADD CONSTRAINT sentry_organization_organizationmember_id_1634015042409685_uniq UNIQUE (organizationmember_id, team_id);


--
-- Name: sentry_organization sentry_organization_pkey; Type: CONSTRAINT; Schema: public; Owner: sentry
--

ALTER TABLE ONLY sentry_organization
    ADD CONSTRAINT sentry_organization_pkey PRIMARY KEY (id);


--
-- Name: sentry_organization sentry_organization_slug_key; Type: CONSTRAINT; Schema: public; Owner: sentry
--

ALTER TABLE ONLY sentry_organization
    ADD CONSTRAINT sentry_organization_slug_key UNIQUE (slug);


--
-- Name: sentry_organizationaccessrequest sentry_organizationaccessrequest_pkey; Type: CONSTRAINT; Schema: public; Owner: sentry
--

ALTER TABLE ONLY sentry_organizationaccessrequest
    ADD CONSTRAINT sentry_organizationaccessrequest_pkey PRIMARY KEY (id);


--
-- Name: sentry_organizationaccessrequest sentry_organizationaccessrequest_team_id_2a38219fe738f1d7_uniq; Type: CONSTRAINT; Schema: public; Owner: sentry
--

ALTER TABLE ONLY sentry_organizationaccessrequest
    ADD CONSTRAINT sentry_organizationaccessrequest_team_id_2a38219fe738f1d7_uniq UNIQUE (team_id, member_id);


--
-- Name: sentry_organizationmember sentry_organizationmember_organization_id_404770fc5e3a794_uniq; Type: CONSTRAINT; Schema: public; Owner: sentry
--

ALTER TABLE ONLY sentry_organizationmember
    ADD CONSTRAINT sentry_organizationmember_organization_id_404770fc5e3a794_uniq UNIQUE (organization_id, user_id);


--
-- Name: sentry_organizationmember sentry_organizationmember_organization_id_59ee8d99c683b0e7_uniq; Type: CONSTRAINT; Schema: public; Owner: sentry
--

ALTER TABLE ONLY sentry_organizationmember
    ADD CONSTRAINT sentry_organizationmember_organization_id_59ee8d99c683b0e7_uniq UNIQUE (organization_id, email);


--
-- Name: sentry_organizationmember sentry_organizationmember_pkey; Type: CONSTRAINT; Schema: public; Owner: sentry
--

ALTER TABLE ONLY sentry_organizationmember
    ADD CONSTRAINT sentry_organizationmember_pkey PRIMARY KEY (id);


--
-- Name: sentry_organizationmember_teams sentry_organizationmember_teams_pkey; Type: CONSTRAINT; Schema: public; Owner: sentry
--

ALTER TABLE ONLY sentry_organizationmember_teams
    ADD CONSTRAINT sentry_organizationmember_teams_pkey PRIMARY KEY (id);


--
-- Name: sentry_organizationmember sentry_organizationmember_token_key; Type: CONSTRAINT; Schema: public; Owner: sentry
--

ALTER TABLE ONLY sentry_organizationmember
    ADD CONSTRAINT sentry_organizationmember_token_key UNIQUE (token);


--
-- Name: sentry_organizationonboardingtask sentry_organizationonboar_organization_id_47e98e05cae29cf3_uniq; Type: CONSTRAINT; Schema: public; Owner: sentry
--

ALTER TABLE ONLY sentry_organizationonboardingtask
    ADD CONSTRAINT sentry_organizationonboar_organization_id_47e98e05cae29cf3_uniq UNIQUE (organization_id, task);


--
-- Name: sentry_organizationonboardingtask sentry_organizationonboardingtask_pkey; Type: CONSTRAINT; Schema: public; Owner: sentry
--

ALTER TABLE ONLY sentry_organizationonboardingtask
    ADD CONSTRAINT sentry_organizationonboardingtask_pkey PRIMARY KEY (id);


--
-- Name: sentry_organizationoptions sentry_organizationoption_organization_id_613ac9b501bd6e71_uniq; Type: CONSTRAINT; Schema: public; Owner: sentry
--

ALTER TABLE ONLY sentry_organizationoptions
    ADD CONSTRAINT sentry_organizationoption_organization_id_613ac9b501bd6e71_uniq UNIQUE (organization_id, key);


--
-- Name: sentry_organizationoptions sentry_organizationoptions_pkey; Type: CONSTRAINT; Schema: public; Owner: sentry
--

ALTER TABLE ONLY sentry_organizationoptions
    ADD CONSTRAINT sentry_organizationoptions_pkey PRIMARY KEY (id);


--
-- Name: sentry_project sentry_project_organization_id_3017a54aeb676236_uniq; Type: CONSTRAINT; Schema: public; Owner: sentry
--

ALTER TABLE ONLY sentry_project
    ADD CONSTRAINT sentry_project_organization_id_3017a54aeb676236_uniq UNIQUE (organization_id, slug);


--
-- Name: sentry_project sentry_project_pkey; Type: CONSTRAINT; Schema: public; Owner: sentry
--

ALTER TABLE ONLY sentry_project
    ADD CONSTRAINT sentry_project_pkey PRIMARY KEY (id);


--
-- Name: sentry_project sentry_project_slug_7e0cc0d379eb3e42_uniq; Type: CONSTRAINT; Schema: public; Owner: sentry
--

ALTER TABLE ONLY sentry_project
    ADD CONSTRAINT sentry_project_slug_7e0cc0d379eb3e42_uniq UNIQUE (slug, team_id);


--
-- Name: sentry_projectbookmark sentry_projectbookmark_pkey; Type: CONSTRAINT; Schema: public; Owner: sentry
--

ALTER TABLE ONLY sentry_projectbookmark
    ADD CONSTRAINT sentry_projectbookmark_pkey PRIMARY KEY (id);


--
-- Name: sentry_projectbookmark sentry_projectbookmark_project_id_450321e77adb9106_uniq; Type: CONSTRAINT; Schema: public; Owner: sentry
--

ALTER TABLE ONLY sentry_projectbookmark
    ADD CONSTRAINT sentry_projectbookmark_project_id_450321e77adb9106_uniq UNIQUE (project_id, user_id);


--
-- Name: sentry_projectcounter sentry_projectcounter_pkey; Type: CONSTRAINT; Schema: public; Owner: sentry
--

ALTER TABLE ONLY sentry_projectcounter
    ADD CONSTRAINT sentry_projectcounter_pkey PRIMARY KEY (id);


--
-- Name: sentry_projectcounter sentry_projectcounter_project_id_key; Type: CONSTRAINT; Schema: public; Owner: sentry
--

ALTER TABLE ONLY sentry_projectcounter
    ADD CONSTRAINT sentry_projectcounter_project_id_key UNIQUE (project_id);


--
-- Name: sentry_projectdsymfile sentry_projectdsymfile_pkey; Type: CONSTRAINT; Schema: public; Owner: sentry
--

ALTER TABLE ONLY sentry_projectdsymfile
    ADD CONSTRAINT sentry_projectdsymfile_pkey PRIMARY KEY (id);


--
-- Name: sentry_projectdsymfile sentry_projectdsymfile_project_id_52cf645985146f12_uniq; Type: CONSTRAINT; Schema: public; Owner: sentry
--

ALTER TABLE ONLY sentry_projectdsymfile
    ADD CONSTRAINT sentry_projectdsymfile_project_id_52cf645985146f12_uniq UNIQUE (project_id, uuid);


--
-- Name: sentry_projectkey sentry_projectkey_pkey; Type: CONSTRAINT; Schema: public; Owner: sentry
--

ALTER TABLE ONLY sentry_projectkey
    ADD CONSTRAINT sentry_projectkey_pkey PRIMARY KEY (id);


--
-- Name: sentry_projectkey sentry_projectkey_public_key_key; Type: CONSTRAINT; Schema: public; Owner: sentry
--

ALTER TABLE ONLY sentry_projectkey
    ADD CONSTRAINT sentry_projectkey_public_key_key UNIQUE (public_key);


--
-- Name: sentry_projectkey sentry_projectkey_secret_key_key; Type: CONSTRAINT; Schema: public; Owner: sentry
--

ALTER TABLE ONLY sentry_projectkey
    ADD CONSTRAINT sentry_projectkey_secret_key_key UNIQUE (secret_key);


--
-- Name: sentry_projectoptions sentry_projectoptions_pkey; Type: CONSTRAINT; Schema: public; Owner: sentry
--

ALTER TABLE ONLY sentry_projectoptions
    ADD CONSTRAINT sentry_projectoptions_pkey PRIMARY KEY (id);


--
-- Name: sentry_projectoptions sentry_projectoptions_project_id_2d0b5c5d84cdbe8f_uniq; Type: CONSTRAINT; Schema: public; Owner: sentry
--

ALTER TABLE ONLY sentry_projectoptions
    ADD CONSTRAINT sentry_projectoptions_project_id_2d0b5c5d84cdbe8f_uniq UNIQUE (project_id, key);


--
-- Name: sentry_projectplatform sentry_projectplatform_pkey; Type: CONSTRAINT; Schema: public; Owner: sentry
--

ALTER TABLE ONLY sentry_projectplatform
    ADD CONSTRAINT sentry_projectplatform_pkey PRIMARY KEY (id);


--
-- Name: sentry_projectplatform sentry_projectplatform_project_id_4750cc420a30bf84_uniq; Type: CONSTRAINT; Schema: public; Owner: sentry
--

ALTER TABLE ONLY sentry_projectplatform
    ADD CONSTRAINT sentry_projectplatform_project_id_4750cc420a30bf84_uniq UNIQUE (project_id, platform);


--
-- Name: sentry_release sentry_release_pkey; Type: CONSTRAINT; Schema: public; Owner: sentry
--

ALTER TABLE ONLY sentry_release
    ADD CONSTRAINT sentry_release_pkey PRIMARY KEY (id);


--
-- Name: sentry_release sentry_release_project_id_7309d47e8f37e7ff_uniq; Type: CONSTRAINT; Schema: public; Owner: sentry
--

ALTER TABLE ONLY sentry_release
    ADD CONSTRAINT sentry_release_project_id_7309d47e8f37e7ff_uniq UNIQUE (project_id, version);


--
-- Name: sentry_release_project sentry_release_project_pkey; Type: CONSTRAINT; Schema: public; Owner: sentry
--

ALTER TABLE ONLY sentry_release_project
    ADD CONSTRAINT sentry_release_project_pkey PRIMARY KEY (id);


--
-- Name: sentry_release_project sentry_release_project_project_id_35add08b8e678ec7_uniq; Type: CONSTRAINT; Schema: public; Owner: sentry
--

ALTER TABLE ONLY sentry_release_project
    ADD CONSTRAINT sentry_release_project_project_id_35add08b8e678ec7_uniq UNIQUE (project_id, release_id);


--
-- Name: sentry_releasecommit sentry_releasecommit_pkey; Type: CONSTRAINT; Schema: public; Owner: sentry
--

ALTER TABLE ONLY sentry_releasecommit
    ADD CONSTRAINT sentry_releasecommit_pkey PRIMARY KEY (id);


--
-- Name: sentry_releasecommit sentry_releasecommit_release_id_4394bda1d741e954_uniq; Type: CONSTRAINT; Schema: public; Owner: sentry
--

ALTER TABLE ONLY sentry_releasecommit
    ADD CONSTRAINT sentry_releasecommit_release_id_4394bda1d741e954_uniq UNIQUE (release_id, "order");


--
-- Name: sentry_releasecommit sentry_releasecommit_release_id_4ce87699e8e032b3_uniq; Type: CONSTRAINT; Schema: public; Owner: sentry
--

ALTER TABLE ONLY sentry_releasecommit
    ADD CONSTRAINT sentry_releasecommit_release_id_4ce87699e8e032b3_uniq UNIQUE (release_id, commit_id);


--
-- Name: sentry_releasefile sentry_releasefile_pkey; Type: CONSTRAINT; Schema: public; Owner: sentry
--

ALTER TABLE ONLY sentry_releasefile
    ADD CONSTRAINT sentry_releasefile_pkey PRIMARY KEY (id);


--
-- Name: sentry_releasefile sentry_releasefile_release_id_7809ae7ca24c9589_uniq; Type: CONSTRAINT; Schema: public; Owner: sentry
--

ALTER TABLE ONLY sentry_releasefile
    ADD CONSTRAINT sentry_releasefile_release_id_7809ae7ca24c9589_uniq UNIQUE (release_id, ident);


--
-- Name: sentry_repository sentry_repository_organization_id_2bbb7c67744745b6_uniq; Type: CONSTRAINT; Schema: public; Owner: sentry
--

ALTER TABLE ONLY sentry_repository
    ADD CONSTRAINT sentry_repository_organization_id_2bbb7c67744745b6_uniq UNIQUE (organization_id, name);


--
-- Name: sentry_repository sentry_repository_organization_id_6369691ee795aeaf_uniq; Type: CONSTRAINT; Schema: public; Owner: sentry
--

ALTER TABLE ONLY sentry_repository
    ADD CONSTRAINT sentry_repository_organization_id_6369691ee795aeaf_uniq UNIQUE (organization_id, provider, external_id);


--
-- Name: sentry_repository sentry_repository_pkey; Type: CONSTRAINT; Schema: public; Owner: sentry
--

ALTER TABLE ONLY sentry_repository
    ADD CONSTRAINT sentry_repository_pkey PRIMARY KEY (id);


--
-- Name: sentry_rule sentry_rule_pkey; Type: CONSTRAINT; Schema: public; Owner: sentry
--

ALTER TABLE ONLY sentry_rule
    ADD CONSTRAINT sentry_rule_pkey PRIMARY KEY (id);


--
-- Name: sentry_savedsearch sentry_savedsearch_pkey; Type: CONSTRAINT; Schema: public; Owner: sentry
--

ALTER TABLE ONLY sentry_savedsearch
    ADD CONSTRAINT sentry_savedsearch_pkey PRIMARY KEY (id);


--
-- Name: sentry_savedsearch sentry_savedsearch_project_id_4a2cf58e27d0cc59_uniq; Type: CONSTRAINT; Schema: public; Owner: sentry
--

ALTER TABLE ONLY sentry_savedsearch
    ADD CONSTRAINT sentry_savedsearch_project_id_4a2cf58e27d0cc59_uniq UNIQUE (project_id, name);


--
-- Name: sentry_savedsearch_userdefault sentry_savedsearch_userdefault_pkey; Type: CONSTRAINT; Schema: public; Owner: sentry
--

ALTER TABLE ONLY sentry_savedsearch_userdefault
    ADD CONSTRAINT sentry_savedsearch_userdefault_pkey PRIMARY KEY (id);


--
-- Name: sentry_savedsearch_userdefault sentry_savedsearch_userdefault_project_id_19fbb9813d6a20ef_uniq; Type: CONSTRAINT; Schema: public; Owner: sentry
--

ALTER TABLE ONLY sentry_savedsearch_userdefault
    ADD CONSTRAINT sentry_savedsearch_userdefault_project_id_19fbb9813d6a20ef_uniq UNIQUE (project_id, user_id);


--
-- Name: sentry_team sentry_team_organization_id_1e0ece47434a2ed_uniq; Type: CONSTRAINT; Schema: public; Owner: sentry
--

ALTER TABLE ONLY sentry_team
    ADD CONSTRAINT sentry_team_organization_id_1e0ece47434a2ed_uniq UNIQUE (organization_id, slug);


--
-- Name: sentry_team sentry_team_pkey; Type: CONSTRAINT; Schema: public; Owner: sentry
--

ALTER TABLE ONLY sentry_team
    ADD CONSTRAINT sentry_team_pkey PRIMARY KEY (id);


--
-- Name: sentry_useravatar sentry_useravatar_file_id_key; Type: CONSTRAINT; Schema: public; Owner: sentry
--

ALTER TABLE ONLY sentry_useravatar
    ADD CONSTRAINT sentry_useravatar_file_id_key UNIQUE (file_id);


--
-- Name: sentry_useravatar sentry_useravatar_ident_key; Type: CONSTRAINT; Schema: public; Owner: sentry
--

ALTER TABLE ONLY sentry_useravatar
    ADD CONSTRAINT sentry_useravatar_ident_key UNIQUE (ident);


--
-- Name: sentry_useravatar sentry_useravatar_pkey; Type: CONSTRAINT; Schema: public; Owner: sentry
--

ALTER TABLE ONLY sentry_useravatar
    ADD CONSTRAINT sentry_useravatar_pkey PRIMARY KEY (id);


--
-- Name: sentry_useravatar sentry_useravatar_user_id_key; Type: CONSTRAINT; Schema: public; Owner: sentry
--

ALTER TABLE ONLY sentry_useravatar
    ADD CONSTRAINT sentry_useravatar_user_id_key UNIQUE (user_id);


--
-- Name: sentry_useremail sentry_useremail_pkey; Type: CONSTRAINT; Schema: public; Owner: sentry
--

ALTER TABLE ONLY sentry_useremail
    ADD CONSTRAINT sentry_useremail_pkey PRIMARY KEY (id);


--
-- Name: sentry_useremail sentry_useremail_user_id_469ffbb142507df2_uniq; Type: CONSTRAINT; Schema: public; Owner: sentry
--

ALTER TABLE ONLY sentry_useremail
    ADD CONSTRAINT sentry_useremail_user_id_469ffbb142507df2_uniq UNIQUE (user_id, email);


--
-- Name: sentry_useroption sentry_useroption_pkey; Type: CONSTRAINT; Schema: public; Owner: sentry
--

ALTER TABLE ONLY sentry_useroption
    ADD CONSTRAINT sentry_useroption_pkey PRIMARY KEY (id);


--
-- Name: sentry_useroption sentry_useroption_user_id_4d4ce0b1f7bb578b_uniq; Type: CONSTRAINT; Schema: public; Owner: sentry
--

ALTER TABLE ONLY sentry_useroption
    ADD CONSTRAINT sentry_useroption_user_id_4d4ce0b1f7bb578b_uniq UNIQUE (user_id, project_id, key);


--
-- Name: sentry_userreport sentry_userreport_pkey; Type: CONSTRAINT; Schema: public; Owner: sentry
--

ALTER TABLE ONLY sentry_userreport
    ADD CONSTRAINT sentry_userreport_pkey PRIMARY KEY (id);


--
-- Name: sentry_userreport sentry_userreport_project_id_1ac377e052723c91_uniq; Type: CONSTRAINT; Schema: public; Owner: sentry
--

ALTER TABLE ONLY sentry_userreport
    ADD CONSTRAINT sentry_userreport_project_id_1ac377e052723c91_uniq UNIQUE (project_id, event_id);


--
-- Name: social_auth_usersocialauth social_auth_usersocialauth_pkey; Type: CONSTRAINT; Schema: public; Owner: sentry
--

ALTER TABLE ONLY social_auth_usersocialauth
    ADD CONSTRAINT social_auth_usersocialauth_pkey PRIMARY KEY (id);


--
-- Name: social_auth_usersocialauth social_auth_usersocialauth_provider_69933d2ea493fc8c_uniq; Type: CONSTRAINT; Schema: public; Owner: sentry
--

ALTER TABLE ONLY social_auth_usersocialauth
    ADD CONSTRAINT social_auth_usersocialauth_provider_69933d2ea493fc8c_uniq UNIQUE (provider, uid, user_id);


--
-- Name: south_migrationhistory south_migrationhistory_pkey; Type: CONSTRAINT; Schema: public; Owner: sentry
--

ALTER TABLE ONLY south_migrationhistory
    ADD CONSTRAINT south_migrationhistory_pkey PRIMARY KEY (id);


--
-- Name: auth_authenticator_user_id; Type: INDEX; Schema: public; Owner: sentry
--

CREATE INDEX auth_authenticator_user_id ON auth_authenticator USING btree (user_id);


--
-- Name: auth_group_name_like; Type: INDEX; Schema: public; Owner: sentry
--

CREATE INDEX auth_group_name_like ON auth_group USING btree (name varchar_pattern_ops);


--
-- Name: auth_group_permissions_group_id; Type: INDEX; Schema: public; Owner: sentry
--

CREATE INDEX auth_group_permissions_group_id ON auth_group_permissions USING btree (group_id);


--
-- Name: auth_group_permissions_permission_id; Type: INDEX; Schema: public; Owner: sentry
--

CREATE INDEX auth_group_permissions_permission_id ON auth_group_permissions USING btree (permission_id);


--
-- Name: auth_permission_content_type_id; Type: INDEX; Schema: public; Owner: sentry
--

CREATE INDEX auth_permission_content_type_id ON auth_permission USING btree (content_type_id);


--
-- Name: auth_user_username_like; Type: INDEX; Schema: public; Owner: sentry
--

CREATE INDEX auth_user_username_like ON auth_user USING btree (username varchar_pattern_ops);


--
-- Name: django_admin_log_content_type_id; Type: INDEX; Schema: public; Owner: sentry
--

CREATE INDEX django_admin_log_content_type_id ON django_admin_log USING btree (content_type_id);


--
-- Name: django_admin_log_user_id; Type: INDEX; Schema: public; Owner: sentry
--

CREATE INDEX django_admin_log_user_id ON django_admin_log USING btree (user_id);


--
-- Name: django_session_expire_date; Type: INDEX; Schema: public; Owner: sentry
--

CREATE INDEX django_session_expire_date ON django_session USING btree (expire_date);


--
-- Name: django_session_session_key_like; Type: INDEX; Schema: public; Owner: sentry
--

CREATE INDEX django_session_session_key_like ON django_session USING btree (session_key varchar_pattern_ops);


--
-- Name: nodestore_node_id_like; Type: INDEX; Schema: public; Owner: sentry
--

CREATE INDEX nodestore_node_id_like ON nodestore_node USING btree (id varchar_pattern_ops);


--
-- Name: nodestore_node_timestamp; Type: INDEX; Schema: public; Owner: sentry
--

CREATE INDEX nodestore_node_timestamp ON nodestore_node USING btree ("timestamp");


--
-- Name: sentry_activity_group_id; Type: INDEX; Schema: public; Owner: sentry
--

CREATE INDEX sentry_activity_group_id ON sentry_activity USING btree (group_id);


--
-- Name: sentry_activity_project_id; Type: INDEX; Schema: public; Owner: sentry
--

CREATE INDEX sentry_activity_project_id ON sentry_activity USING btree (project_id);


--
-- Name: sentry_activity_user_id; Type: INDEX; Schema: public; Owner: sentry
--

CREATE INDEX sentry_activity_user_id ON sentry_activity USING btree (user_id);


--
-- Name: sentry_apikey_key_like; Type: INDEX; Schema: public; Owner: sentry
--

CREATE INDEX sentry_apikey_key_like ON sentry_apikey USING btree (key varchar_pattern_ops);


--
-- Name: sentry_apikey_organization_id; Type: INDEX; Schema: public; Owner: sentry
--

CREATE INDEX sentry_apikey_organization_id ON sentry_apikey USING btree (organization_id);


--
-- Name: sentry_apikey_status; Type: INDEX; Schema: public; Owner: sentry
--

CREATE INDEX sentry_apikey_status ON sentry_apikey USING btree (status);


--
-- Name: sentry_apitoken_key_id; Type: INDEX; Schema: public; Owner: sentry
--

CREATE INDEX sentry_apitoken_key_id ON sentry_apitoken USING btree (key_id);


--
-- Name: sentry_apitoken_token_like; Type: INDEX; Schema: public; Owner: sentry
--

CREATE INDEX sentry_apitoken_token_like ON sentry_apitoken USING btree (token varchar_pattern_ops);


--
-- Name: sentry_apitoken_user_id; Type: INDEX; Schema: public; Owner: sentry
--

CREATE INDEX sentry_apitoken_user_id ON sentry_apitoken USING btree (user_id);


--
-- Name: sentry_auditlogentry_actor_id; Type: INDEX; Schema: public; Owner: sentry
--

CREATE INDEX sentry_auditlogentry_actor_id ON sentry_auditlogentry USING btree (actor_id);


--
-- Name: sentry_auditlogentry_actor_key_id; Type: INDEX; Schema: public; Owner: sentry
--

CREATE INDEX sentry_auditlogentry_actor_key_id ON sentry_auditlogentry USING btree (actor_key_id);


--
-- Name: sentry_auditlogentry_organization_id; Type: INDEX; Schema: public; Owner: sentry
--

CREATE INDEX sentry_auditlogentry_organization_id ON sentry_auditlogentry USING btree (organization_id);


--
-- Name: sentry_auditlogentry_target_user_id; Type: INDEX; Schema: public; Owner: sentry
--

CREATE INDEX sentry_auditlogentry_target_user_id ON sentry_auditlogentry USING btree (target_user_id);


--
-- Name: sentry_authidentity_auth_provider_id; Type: INDEX; Schema: public; Owner: sentry
--

CREATE INDEX sentry_authidentity_auth_provider_id ON sentry_authidentity USING btree (auth_provider_id);


--
-- Name: sentry_authidentity_user_id; Type: INDEX; Schema: public; Owner: sentry
--

CREATE INDEX sentry_authidentity_user_id ON sentry_authidentity USING btree (user_id);


--
-- Name: sentry_authprovider_default_teams_authprovider_id; Type: INDEX; Schema: public; Owner: sentry
--

CREATE INDEX sentry_authprovider_default_teams_authprovider_id ON sentry_authprovider_default_teams USING btree (authprovider_id);


--
-- Name: sentry_authprovider_default_teams_team_id; Type: INDEX; Schema: public; Owner: sentry
--

CREATE INDEX sentry_authprovider_default_teams_team_id ON sentry_authprovider_default_teams USING btree (team_id);


--
-- Name: sentry_broadcast_is_active; Type: INDEX; Schema: public; Owner: sentry
--

CREATE INDEX sentry_broadcast_is_active ON sentry_broadcast USING btree (is_active);


--
-- Name: sentry_broadcastseen_broadcast_id; Type: INDEX; Schema: public; Owner: sentry
--

CREATE INDEX sentry_broadcastseen_broadcast_id ON sentry_broadcastseen USING btree (broadcast_id);


--
-- Name: sentry_broadcastseen_user_id; Type: INDEX; Schema: public; Owner: sentry
--

CREATE INDEX sentry_broadcastseen_user_id ON sentry_broadcastseen USING btree (user_id);


--
-- Name: sentry_commit_author_id; Type: INDEX; Schema: public; Owner: sentry
--

CREATE INDEX sentry_commit_author_id ON sentry_commit USING btree (author_id);


--
-- Name: sentry_commit_organization_id; Type: INDEX; Schema: public; Owner: sentry
--

CREATE INDEX sentry_commit_organization_id ON sentry_commit USING btree (organization_id);


--
-- Name: sentry_commit_repository_id_5b0d06238a42bbfc; Type: INDEX; Schema: public; Owner: sentry
--

CREATE INDEX sentry_commit_repository_id_5b0d06238a42bbfc ON sentry_commit USING btree (repository_id, date_added);


--
-- Name: sentry_commitauthor_organization_id; Type: INDEX; Schema: public; Owner: sentry
--

CREATE INDEX sentry_commitauthor_organization_id ON sentry_commitauthor USING btree (organization_id);


--
-- Name: sentry_commitfilechange_commit_id; Type: INDEX; Schema: public; Owner: sentry
--

CREATE INDEX sentry_commitfilechange_commit_id ON sentry_commitfilechange USING btree (commit_id);


--
-- Name: sentry_commitfilechange_organization_id; Type: INDEX; Schema: public; Owner: sentry
--

CREATE INDEX sentry_commitfilechange_organization_id ON sentry_commitfilechange USING btree (organization_id);


--
-- Name: sentry_dsymbundle_object_id; Type: INDEX; Schema: public; Owner: sentry
--

CREATE INDEX sentry_dsymbundle_object_id ON sentry_dsymbundle USING btree (object_id);


--
-- Name: sentry_dsymbundle_sdk_id; Type: INDEX; Schema: public; Owner: sentry
--

CREATE INDEX sentry_dsymbundle_sdk_id ON sentry_dsymbundle USING btree (sdk_id);


--
-- Name: sentry_dsymobject_object_path; Type: INDEX; Schema: public; Owner: sentry
--

CREATE INDEX sentry_dsymobject_object_path ON sentry_dsymobject USING btree (object_path);


--
-- Name: sentry_dsymobject_object_path_like; Type: INDEX; Schema: public; Owner: sentry
--

CREATE INDEX sentry_dsymobject_object_path_like ON sentry_dsymobject USING btree (object_path text_pattern_ops);


--
-- Name: sentry_dsymobject_uuid; Type: INDEX; Schema: public; Owner: sentry
--

CREATE INDEX sentry_dsymobject_uuid ON sentry_dsymobject USING btree (uuid);


--
-- Name: sentry_dsymobject_uuid_like; Type: INDEX; Schema: public; Owner: sentry
--

CREATE INDEX sentry_dsymobject_uuid_like ON sentry_dsymobject USING btree (uuid varchar_pattern_ops);


--
-- Name: sentry_dsymsdk_dsym_type; Type: INDEX; Schema: public; Owner: sentry
--

CREATE INDEX sentry_dsymsdk_dsym_type ON sentry_dsymsdk USING btree (dsym_type);


--
-- Name: sentry_dsymsdk_dsym_type_like; Type: INDEX; Schema: public; Owner: sentry
--

CREATE INDEX sentry_dsymsdk_dsym_type_like ON sentry_dsymsdk USING btree (dsym_type varchar_pattern_ops);


--
-- Name: sentry_dsymsdk_version_major_8d012290987c340; Type: INDEX; Schema: public; Owner: sentry
--

CREATE INDEX sentry_dsymsdk_version_major_8d012290987c340 ON sentry_dsymsdk USING btree (version_major, version_minor, version_patchlevel, version_build);


--
-- Name: sentry_dsymsymbol_address; Type: INDEX; Schema: public; Owner: sentry
--

CREATE INDEX sentry_dsymsymbol_address ON sentry_dsymsymbol USING btree (address);


--
-- Name: sentry_dsymsymbol_object_id; Type: INDEX; Schema: public; Owner: sentry
--

CREATE INDEX sentry_dsymsymbol_object_id ON sentry_dsymsymbol USING btree (object_id);


--
-- Name: sentry_environmentrelease_environment_id; Type: INDEX; Schema: public; Owner: sentry
--

CREATE INDEX sentry_environmentrelease_environment_id ON sentry_environmentrelease USING btree (environment_id);


--
-- Name: sentry_environmentrelease_last_seen; Type: INDEX; Schema: public; Owner: sentry
--

CREATE INDEX sentry_environmentrelease_last_seen ON sentry_environmentrelease USING btree (last_seen);


--
-- Name: sentry_environmentrelease_organization_id; Type: INDEX; Schema: public; Owner: sentry
--

CREATE INDEX sentry_environmentrelease_organization_id ON sentry_environmentrelease USING btree (organization_id);


--
-- Name: sentry_environmentrelease_project_id; Type: INDEX; Schema: public; Owner: sentry
--

CREATE INDEX sentry_environmentrelease_project_id ON sentry_environmentrelease USING btree (project_id);


--
-- Name: sentry_environmentrelease_release_id; Type: INDEX; Schema: public; Owner: sentry
--

CREATE INDEX sentry_environmentrelease_release_id ON sentry_environmentrelease USING btree (release_id);


--
-- Name: sentry_eventmapping_group_id; Type: INDEX; Schema: public; Owner: sentry
--

CREATE INDEX sentry_eventmapping_group_id ON sentry_eventmapping USING btree (group_id);


--
-- Name: sentry_eventmapping_project_id; Type: INDEX; Schema: public; Owner: sentry
--

CREATE INDEX sentry_eventmapping_project_id ON sentry_eventmapping USING btree (project_id);


--
-- Name: sentry_eventtag_group_id_5ad9abfe8e1fa62b; Type: INDEX; Schema: public; Owner: sentry
--

CREATE INDEX sentry_eventtag_group_id_5ad9abfe8e1fa62b ON sentry_eventtag USING btree (group_id, key_id, value_id);


--
-- Name: sentry_eventtag_project_id_42979ba214ba3c43; Type: INDEX; Schema: public; Owner: sentry
--

CREATE INDEX sentry_eventtag_project_id_42979ba214ba3c43 ON sentry_eventtag USING btree (project_id, key_id, value_id);


--
-- Name: sentry_eventuser_date_added; Type: INDEX; Schema: public; Owner: sentry
--

CREATE INDEX sentry_eventuser_date_added ON sentry_eventuser USING btree (date_added);


--
-- Name: sentry_eventuser_project_id; Type: INDEX; Schema: public; Owner: sentry
--

CREATE INDEX sentry_eventuser_project_id ON sentry_eventuser USING btree (project_id);


--
-- Name: sentry_eventuser_project_id_58b4a7f2595290e6; Type: INDEX; Schema: public; Owner: sentry
--

CREATE INDEX sentry_eventuser_project_id_58b4a7f2595290e6 ON sentry_eventuser USING btree (project_id, ip_address);


--
-- Name: sentry_eventuser_project_id_7684267daffc292f; Type: INDEX; Schema: public; Owner: sentry
--

CREATE INDEX sentry_eventuser_project_id_7684267daffc292f ON sentry_eventuser USING btree (project_id, email);


--
-- Name: sentry_eventuser_project_id_8868307f60b6a92; Type: INDEX; Schema: public; Owner: sentry
--

CREATE INDEX sentry_eventuser_project_id_8868307f60b6a92 ON sentry_eventuser USING btree (project_id, username);


--
-- Name: sentry_file_blob_id; Type: INDEX; Schema: public; Owner: sentry
--

CREATE INDEX sentry_file_blob_id ON sentry_file USING btree (blob_id);


--
-- Name: sentry_file_timestamp; Type: INDEX; Schema: public; Owner: sentry
--

CREATE INDEX sentry_file_timestamp ON sentry_file USING btree ("timestamp");


--
-- Name: sentry_fileblob_checksum_like; Type: INDEX; Schema: public; Owner: sentry
--

CREATE INDEX sentry_fileblob_checksum_like ON sentry_fileblob USING btree (checksum varchar_pattern_ops);


--
-- Name: sentry_fileblob_timestamp; Type: INDEX; Schema: public; Owner: sentry
--

CREATE INDEX sentry_fileblob_timestamp ON sentry_fileblob USING btree ("timestamp");


--
-- Name: sentry_fileblobindex_blob_id; Type: INDEX; Schema: public; Owner: sentry
--

CREATE INDEX sentry_fileblobindex_blob_id ON sentry_fileblobindex USING btree (blob_id);


--
-- Name: sentry_fileblobindex_file_id; Type: INDEX; Schema: public; Owner: sentry
--

CREATE INDEX sentry_fileblobindex_file_id ON sentry_fileblobindex USING btree (file_id);


--
-- Name: sentry_filterkey_project_id; Type: INDEX; Schema: public; Owner: sentry
--

CREATE INDEX sentry_filterkey_project_id ON sentry_filterkey USING btree (project_id);


--
-- Name: sentry_filtervalue_first_seen; Type: INDEX; Schema: public; Owner: sentry
--

CREATE INDEX sentry_filtervalue_first_seen ON sentry_filtervalue USING btree (first_seen);


--
-- Name: sentry_filtervalue_last_seen; Type: INDEX; Schema: public; Owner: sentry
--

CREATE INDEX sentry_filtervalue_last_seen ON sentry_filtervalue USING btree (last_seen);


--
-- Name: sentry_filtervalue_project_id; Type: INDEX; Schema: public; Owner: sentry
--

CREATE INDEX sentry_filtervalue_project_id ON sentry_filtervalue USING btree (project_id);


--
-- Name: sentry_filtervalue_project_id_27377f6151fcab56; Type: INDEX; Schema: public; Owner: sentry
--

CREATE INDEX sentry_filtervalue_project_id_27377f6151fcab56 ON sentry_filtervalue USING btree (project_id, value, last_seen);


--
-- Name: sentry_filtervalue_project_id_2b3fdfeac62987c7; Type: INDEX; Schema: public; Owner: sentry
--

CREATE INDEX sentry_filtervalue_project_id_2b3fdfeac62987c7 ON sentry_filtervalue USING btree (project_id, value, first_seen);


--
-- Name: sentry_filtervalue_project_id_737632cad2909511; Type: INDEX; Schema: public; Owner: sentry
--

CREATE INDEX sentry_filtervalue_project_id_737632cad2909511 ON sentry_filtervalue USING btree (project_id, value, times_seen);


--
-- Name: sentry_globaldsymfile_file_id; Type: INDEX; Schema: public; Owner: sentry
--

CREATE INDEX sentry_globaldsymfile_file_id ON sentry_globaldsymfile USING btree (file_id);


--
-- Name: sentry_globaldsymfile_uuid_like; Type: INDEX; Schema: public; Owner: sentry
--

CREATE INDEX sentry_globaldsymfile_uuid_like ON sentry_globaldsymfile USING btree (uuid varchar_pattern_ops);


--
-- Name: sentry_groupasignee_project_id; Type: INDEX; Schema: public; Owner: sentry
--

CREATE INDEX sentry_groupasignee_project_id ON sentry_groupasignee USING btree (project_id);


--
-- Name: sentry_groupasignee_user_id; Type: INDEX; Schema: public; Owner: sentry
--

CREATE INDEX sentry_groupasignee_user_id ON sentry_groupasignee USING btree (user_id);


--
-- Name: sentry_groupbookmark_group_id; Type: INDEX; Schema: public; Owner: sentry
--

CREATE INDEX sentry_groupbookmark_group_id ON sentry_groupbookmark USING btree (group_id);


--
-- Name: sentry_groupbookmark_project_id; Type: INDEX; Schema: public; Owner: sentry
--

CREATE INDEX sentry_groupbookmark_project_id ON sentry_groupbookmark USING btree (project_id);


--
-- Name: sentry_groupbookmark_user_id; Type: INDEX; Schema: public; Owner: sentry
--

CREATE INDEX sentry_groupbookmark_user_id ON sentry_groupbookmark USING btree (user_id);


--
-- Name: sentry_groupbookmark_user_id_5eedb134f529cf58; Type: INDEX; Schema: public; Owner: sentry
--

CREATE INDEX sentry_groupbookmark_user_id_5eedb134f529cf58 ON sentry_groupbookmark USING btree (user_id, group_id);


--
-- Name: sentry_groupedmessage_active_at; Type: INDEX; Schema: public; Owner: sentry
--

CREATE INDEX sentry_groupedmessage_active_at ON sentry_groupedmessage USING btree (active_at);


--
-- Name: sentry_groupedmessage_first_release_id; Type: INDEX; Schema: public; Owner: sentry
--

CREATE INDEX sentry_groupedmessage_first_release_id ON sentry_groupedmessage USING btree (first_release_id);


--
-- Name: sentry_groupedmessage_first_seen; Type: INDEX; Schema: public; Owner: sentry
--

CREATE INDEX sentry_groupedmessage_first_seen ON sentry_groupedmessage USING btree (first_seen);


--
-- Name: sentry_groupedmessage_last_seen; Type: INDEX; Schema: public; Owner: sentry
--

CREATE INDEX sentry_groupedmessage_last_seen ON sentry_groupedmessage USING btree (last_seen);


--
-- Name: sentry_groupedmessage_level; Type: INDEX; Schema: public; Owner: sentry
--

CREATE INDEX sentry_groupedmessage_level ON sentry_groupedmessage USING btree (level);


--
-- Name: sentry_groupedmessage_logger; Type: INDEX; Schema: public; Owner: sentry
--

CREATE INDEX sentry_groupedmessage_logger ON sentry_groupedmessage USING btree (logger);


--
-- Name: sentry_groupedmessage_logger_like; Type: INDEX; Schema: public; Owner: sentry
--

CREATE INDEX sentry_groupedmessage_logger_like ON sentry_groupedmessage USING btree (logger varchar_pattern_ops);


--
-- Name: sentry_groupedmessage_project_id; Type: INDEX; Schema: public; Owner: sentry
--

CREATE INDEX sentry_groupedmessage_project_id ON sentry_groupedmessage USING btree (project_id);


--
-- Name: sentry_groupedmessage_project_id_31335ae34c8ef983; Type: INDEX; Schema: public; Owner: sentry
--

CREATE INDEX sentry_groupedmessage_project_id_31335ae34c8ef983 ON sentry_groupedmessage USING btree (project_id, first_release_id);


--
-- Name: sentry_groupedmessage_resolved_at; Type: INDEX; Schema: public; Owner: sentry
--

CREATE INDEX sentry_groupedmessage_resolved_at ON sentry_groupedmessage USING btree (resolved_at);


--
-- Name: sentry_groupedmessage_status; Type: INDEX; Schema: public; Owner: sentry
--

CREATE INDEX sentry_groupedmessage_status ON sentry_groupedmessage USING btree (status);


--
-- Name: sentry_groupedmessage_times_seen; Type: INDEX; Schema: public; Owner: sentry
--

CREATE INDEX sentry_groupedmessage_times_seen ON sentry_groupedmessage USING btree (times_seen);


--
-- Name: sentry_groupedmessage_view; Type: INDEX; Schema: public; Owner: sentry
--

CREATE INDEX sentry_groupedmessage_view ON sentry_groupedmessage USING btree (view);


--
-- Name: sentry_groupedmessage_view_like; Type: INDEX; Schema: public; Owner: sentry
--

CREATE INDEX sentry_groupedmessage_view_like ON sentry_groupedmessage USING btree (view varchar_pattern_ops);


--
-- Name: sentry_groupemailthread_date; Type: INDEX; Schema: public; Owner: sentry
--

CREATE INDEX sentry_groupemailthread_date ON sentry_groupemailthread USING btree (date);


--
-- Name: sentry_groupemailthread_group_id; Type: INDEX; Schema: public; Owner: sentry
--

CREATE INDEX sentry_groupemailthread_group_id ON sentry_groupemailthread USING btree (group_id);


--
-- Name: sentry_groupemailthread_msgid_like; Type: INDEX; Schema: public; Owner: sentry
--

CREATE INDEX sentry_groupemailthread_msgid_like ON sentry_groupemailthread USING btree (msgid varchar_pattern_ops);


--
-- Name: sentry_groupemailthread_project_id; Type: INDEX; Schema: public; Owner: sentry
--

CREATE INDEX sentry_groupemailthread_project_id ON sentry_groupemailthread USING btree (project_id);


--
-- Name: sentry_grouphash_group_id; Type: INDEX; Schema: public; Owner: sentry
--

CREATE INDEX sentry_grouphash_group_id ON sentry_grouphash USING btree (group_id);


--
-- Name: sentry_grouphash_hash_like; Type: INDEX; Schema: public; Owner: sentry
--

CREATE INDEX sentry_grouphash_hash_like ON sentry_grouphash USING btree (hash varchar_pattern_ops);


--
-- Name: sentry_grouphash_project_id; Type: INDEX; Schema: public; Owner: sentry
--

CREATE INDEX sentry_grouphash_project_id ON sentry_grouphash USING btree (project_id);


--
-- Name: sentry_groupmeta_group_id; Type: INDEX; Schema: public; Owner: sentry
--

CREATE INDEX sentry_groupmeta_group_id ON sentry_groupmeta USING btree (group_id);


--
-- Name: sentry_groupredirect_group_id; Type: INDEX; Schema: public; Owner: sentry
--

CREATE INDEX sentry_groupredirect_group_id ON sentry_groupredirect USING btree (group_id);


--
-- Name: sentry_grouprelease_last_seen; Type: INDEX; Schema: public; Owner: sentry
--

CREATE INDEX sentry_grouprelease_last_seen ON sentry_grouprelease USING btree (last_seen);


--
-- Name: sentry_grouprelease_project_id; Type: INDEX; Schema: public; Owner: sentry
--

CREATE INDEX sentry_grouprelease_project_id ON sentry_grouprelease USING btree (project_id);


--
-- Name: sentry_grouprelease_release_id; Type: INDEX; Schema: public; Owner: sentry
--

CREATE INDEX sentry_grouprelease_release_id ON sentry_grouprelease USING btree (release_id);


--
-- Name: sentry_groupresolution_datetime; Type: INDEX; Schema: public; Owner: sentry
--

CREATE INDEX sentry_groupresolution_datetime ON sentry_groupresolution USING btree (datetime);


--
-- Name: sentry_groupresolution_release_id; Type: INDEX; Schema: public; Owner: sentry
--

CREATE INDEX sentry_groupresolution_release_id ON sentry_groupresolution USING btree (release_id);


--
-- Name: sentry_grouprulestatus_group_id; Type: INDEX; Schema: public; Owner: sentry
--

CREATE INDEX sentry_grouprulestatus_group_id ON sentry_grouprulestatus USING btree (group_id);


--
-- Name: sentry_grouprulestatus_project_id; Type: INDEX; Schema: public; Owner: sentry
--

CREATE INDEX sentry_grouprulestatus_project_id ON sentry_grouprulestatus USING btree (project_id);


--
-- Name: sentry_grouprulestatus_rule_id; Type: INDEX; Schema: public; Owner: sentry
--

CREATE INDEX sentry_grouprulestatus_rule_id ON sentry_grouprulestatus USING btree (rule_id);


--
-- Name: sentry_groupseen_group_id; Type: INDEX; Schema: public; Owner: sentry
--

CREATE INDEX sentry_groupseen_group_id ON sentry_groupseen USING btree (group_id);


--
-- Name: sentry_groupseen_project_id; Type: INDEX; Schema: public; Owner: sentry
--

CREATE INDEX sentry_groupseen_project_id ON sentry_groupseen USING btree (project_id);


--
-- Name: sentry_groupsubscription_group_id; Type: INDEX; Schema: public; Owner: sentry
--

CREATE INDEX sentry_groupsubscription_group_id ON sentry_groupsubscription USING btree (group_id);


--
-- Name: sentry_groupsubscription_project_id; Type: INDEX; Schema: public; Owner: sentry
--

CREATE INDEX sentry_groupsubscription_project_id ON sentry_groupsubscription USING btree (project_id);


--
-- Name: sentry_groupsubscription_user_id; Type: INDEX; Schema: public; Owner: sentry
--

CREATE INDEX sentry_groupsubscription_user_id ON sentry_groupsubscription USING btree (user_id);


--
-- Name: sentry_grouptagkey_group_id; Type: INDEX; Schema: public; Owner: sentry
--

CREATE INDEX sentry_grouptagkey_group_id ON sentry_grouptagkey USING btree (group_id);


--
-- Name: sentry_grouptagkey_project_id; Type: INDEX; Schema: public; Owner: sentry
--

CREATE INDEX sentry_grouptagkey_project_id ON sentry_grouptagkey USING btree (project_id);


--
-- Name: sentry_message_datetime; Type: INDEX; Schema: public; Owner: sentry
--

CREATE INDEX sentry_message_datetime ON sentry_message USING btree (datetime);


--
-- Name: sentry_message_group_id_5f63ffbd9aac1141; Type: INDEX; Schema: public; Owner: sentry
--

CREATE INDEX sentry_message_group_id_5f63ffbd9aac1141 ON sentry_message USING btree (group_id, datetime);


--
-- Name: sentry_message_message_id_like; Type: INDEX; Schema: public; Owner: sentry
--

CREATE INDEX sentry_message_message_id_like ON sentry_message USING btree (message_id varchar_pattern_ops);


--
-- Name: sentry_message_project_id; Type: INDEX; Schema: public; Owner: sentry
--

CREATE INDEX sentry_message_project_id ON sentry_message USING btree (project_id);


--
-- Name: sentry_messagefiltervalue_first_seen; Type: INDEX; Schema: public; Owner: sentry
--

CREATE INDEX sentry_messagefiltervalue_first_seen ON sentry_messagefiltervalue USING btree (first_seen);


--
-- Name: sentry_messagefiltervalue_group_id; Type: INDEX; Schema: public; Owner: sentry
--

CREATE INDEX sentry_messagefiltervalue_group_id ON sentry_messagefiltervalue USING btree (group_id);


--
-- Name: sentry_messagefiltervalue_group_id_59490523e6ee451f; Type: INDEX; Schema: public; Owner: sentry
--

CREATE UNIQUE INDEX sentry_messagefiltervalue_group_id_59490523e6ee451f ON sentry_messagefiltervalue USING btree (group_id, key, value);


--
-- Name: sentry_messagefiltervalue_last_seen; Type: INDEX; Schema: public; Owner: sentry
--

CREATE INDEX sentry_messagefiltervalue_last_seen ON sentry_messagefiltervalue USING btree (last_seen);


--
-- Name: sentry_messagefiltervalue_project_id_6852dd47401b2d7d; Type: INDEX; Schema: public; Owner: sentry
--

CREATE INDEX sentry_messagefiltervalue_project_id_6852dd47401b2d7d ON sentry_messagefiltervalue USING btree (project_id, key, value, last_seen);


--
-- Name: sentry_organization_slug_like; Type: INDEX; Schema: public; Owner: sentry
--

CREATE INDEX sentry_organization_slug_like ON sentry_organization USING btree (slug varchar_pattern_ops);


--
-- Name: sentry_organizationaccessrequest_member_id; Type: INDEX; Schema: public; Owner: sentry
--

CREATE INDEX sentry_organizationaccessrequest_member_id ON sentry_organizationaccessrequest USING btree (member_id);


--
-- Name: sentry_organizationaccessrequest_team_id; Type: INDEX; Schema: public; Owner: sentry
--

CREATE INDEX sentry_organizationaccessrequest_team_id ON sentry_organizationaccessrequest USING btree (team_id);


--
-- Name: sentry_organizationmember_organization_id; Type: INDEX; Schema: public; Owner: sentry
--

CREATE INDEX sentry_organizationmember_organization_id ON sentry_organizationmember USING btree (organization_id);


--
-- Name: sentry_organizationmember_teams_organizationmember_id; Type: INDEX; Schema: public; Owner: sentry
--

CREATE INDEX sentry_organizationmember_teams_organizationmember_id ON sentry_organizationmember_teams USING btree (organizationmember_id);


--
-- Name: sentry_organizationmember_teams_team_id; Type: INDEX; Schema: public; Owner: sentry
--

CREATE INDEX sentry_organizationmember_teams_team_id ON sentry_organizationmember_teams USING btree (team_id);


--
-- Name: sentry_organizationmember_token_like; Type: INDEX; Schema: public; Owner: sentry
--

CREATE INDEX sentry_organizationmember_token_like ON sentry_organizationmember USING btree (token varchar_pattern_ops);


--
-- Name: sentry_organizationmember_user_id; Type: INDEX; Schema: public; Owner: sentry
--

CREATE INDEX sentry_organizationmember_user_id ON sentry_organizationmember USING btree (user_id);


--
-- Name: sentry_organizationonboardingtask_organization_id; Type: INDEX; Schema: public; Owner: sentry
--

CREATE INDEX sentry_organizationonboardingtask_organization_id ON sentry_organizationonboardingtask USING btree (organization_id);


--
-- Name: sentry_organizationonboardingtask_user_id; Type: INDEX; Schema: public; Owner: sentry
--

CREATE INDEX sentry_organizationonboardingtask_user_id ON sentry_organizationonboardingtask USING btree (user_id);


--
-- Name: sentry_organizationoptions_organization_id; Type: INDEX; Schema: public; Owner: sentry
--

CREATE INDEX sentry_organizationoptions_organization_id ON sentry_organizationoptions USING btree (organization_id);


--
-- Name: sentry_project_organization_id; Type: INDEX; Schema: public; Owner: sentry
--

CREATE INDEX sentry_project_organization_id ON sentry_project USING btree (organization_id);


--
-- Name: sentry_project_slug_like; Type: INDEX; Schema: public; Owner: sentry
--

CREATE INDEX sentry_project_slug_like ON sentry_project USING btree (slug varchar_pattern_ops);


--
-- Name: sentry_project_status; Type: INDEX; Schema: public; Owner: sentry
--

CREATE INDEX sentry_project_status ON sentry_project USING btree (status);


--
-- Name: sentry_project_team_id; Type: INDEX; Schema: public; Owner: sentry
--

CREATE INDEX sentry_project_team_id ON sentry_project USING btree (team_id);


--
-- Name: sentry_projectbookmark_user_id; Type: INDEX; Schema: public; Owner: sentry
--

CREATE INDEX sentry_projectbookmark_user_id ON sentry_projectbookmark USING btree (user_id);


--
-- Name: sentry_projectdsymfile_file_id; Type: INDEX; Schema: public; Owner: sentry
--

CREATE INDEX sentry_projectdsymfile_file_id ON sentry_projectdsymfile USING btree (file_id);


--
-- Name: sentry_projectdsymfile_project_id; Type: INDEX; Schema: public; Owner: sentry
--

CREATE INDEX sentry_projectdsymfile_project_id ON sentry_projectdsymfile USING btree (project_id);


--
-- Name: sentry_projectkey_project_id; Type: INDEX; Schema: public; Owner: sentry
--

CREATE INDEX sentry_projectkey_project_id ON sentry_projectkey USING btree (project_id);


--
-- Name: sentry_projectkey_public_key_like; Type: INDEX; Schema: public; Owner: sentry
--

CREATE INDEX sentry_projectkey_public_key_like ON sentry_projectkey USING btree (public_key varchar_pattern_ops);


--
-- Name: sentry_projectkey_secret_key_like; Type: INDEX; Schema: public; Owner: sentry
--

CREATE INDEX sentry_projectkey_secret_key_like ON sentry_projectkey USING btree (secret_key varchar_pattern_ops);


--
-- Name: sentry_projectkey_status; Type: INDEX; Schema: public; Owner: sentry
--

CREATE INDEX sentry_projectkey_status ON sentry_projectkey USING btree (status);


--
-- Name: sentry_projectoptions_project_id; Type: INDEX; Schema: public; Owner: sentry
--

CREATE INDEX sentry_projectoptions_project_id ON sentry_projectoptions USING btree (project_id);


--
-- Name: sentry_release_organization_id; Type: INDEX; Schema: public; Owner: sentry
--

CREATE INDEX sentry_release_organization_id ON sentry_release USING btree (organization_id);


--
-- Name: sentry_release_owner_id; Type: INDEX; Schema: public; Owner: sentry
--

CREATE INDEX sentry_release_owner_id ON sentry_release USING btree (owner_id);


--
-- Name: sentry_release_project_id; Type: INDEX; Schema: public; Owner: sentry
--

CREATE INDEX sentry_release_project_id ON sentry_release USING btree (project_id);


--
-- Name: sentry_release_project_project_id; Type: INDEX; Schema: public; Owner: sentry
--

CREATE INDEX sentry_release_project_project_id ON sentry_release_project USING btree (project_id);


--
-- Name: sentry_release_project_release_id; Type: INDEX; Schema: public; Owner: sentry
--

CREATE INDEX sentry_release_project_release_id ON sentry_release_project USING btree (release_id);


--
-- Name: sentry_releasecommit_commit_id; Type: INDEX; Schema: public; Owner: sentry
--

CREATE INDEX sentry_releasecommit_commit_id ON sentry_releasecommit USING btree (commit_id);


--
-- Name: sentry_releasecommit_organization_id; Type: INDEX; Schema: public; Owner: sentry
--

CREATE INDEX sentry_releasecommit_organization_id ON sentry_releasecommit USING btree (organization_id);


--
-- Name: sentry_releasecommit_project_id; Type: INDEX; Schema: public; Owner: sentry
--

CREATE INDEX sentry_releasecommit_project_id ON sentry_releasecommit USING btree (project_id);


--
-- Name: sentry_releasecommit_release_id; Type: INDEX; Schema: public; Owner: sentry
--

CREATE INDEX sentry_releasecommit_release_id ON sentry_releasecommit USING btree (release_id);


--
-- Name: sentry_releasefile_file_id; Type: INDEX; Schema: public; Owner: sentry
--

CREATE INDEX sentry_releasefile_file_id ON sentry_releasefile USING btree (file_id);


--
-- Name: sentry_releasefile_organization_id; Type: INDEX; Schema: public; Owner: sentry
--

CREATE INDEX sentry_releasefile_organization_id ON sentry_releasefile USING btree (organization_id);


--
-- Name: sentry_releasefile_project_id; Type: INDEX; Schema: public; Owner: sentry
--

CREATE INDEX sentry_releasefile_project_id ON sentry_releasefile USING btree (project_id);


--
-- Name: sentry_releasefile_release_id; Type: INDEX; Schema: public; Owner: sentry
--

CREATE INDEX sentry_releasefile_release_id ON sentry_releasefile USING btree (release_id);


--
-- Name: sentry_repository_organization_id; Type: INDEX; Schema: public; Owner: sentry
--

CREATE INDEX sentry_repository_organization_id ON sentry_repository USING btree (organization_id);


--
-- Name: sentry_repository_status; Type: INDEX; Schema: public; Owner: sentry
--

CREATE INDEX sentry_repository_status ON sentry_repository USING btree (status);


--
-- Name: sentry_rule_project_id; Type: INDEX; Schema: public; Owner: sentry
--

CREATE INDEX sentry_rule_project_id ON sentry_rule USING btree (project_id);


--
-- Name: sentry_rule_status; Type: INDEX; Schema: public; Owner: sentry
--

CREATE INDEX sentry_rule_status ON sentry_rule USING btree (status);


--
-- Name: sentry_savedsearch_project_id; Type: INDEX; Schema: public; Owner: sentry
--

CREATE INDEX sentry_savedsearch_project_id ON sentry_savedsearch USING btree (project_id);


--
-- Name: sentry_savedsearch_userdefault_project_id; Type: INDEX; Schema: public; Owner: sentry
--

CREATE INDEX sentry_savedsearch_userdefault_project_id ON sentry_savedsearch_userdefault USING btree (project_id);


--
-- Name: sentry_savedsearch_userdefault_savedsearch_id; Type: INDEX; Schema: public; Owner: sentry
--

CREATE INDEX sentry_savedsearch_userdefault_savedsearch_id ON sentry_savedsearch_userdefault USING btree (savedsearch_id);


--
-- Name: sentry_savedsearch_userdefault_user_id; Type: INDEX; Schema: public; Owner: sentry
--

CREATE INDEX sentry_savedsearch_userdefault_user_id ON sentry_savedsearch_userdefault USING btree (user_id);


--
-- Name: sentry_team_organization_id; Type: INDEX; Schema: public; Owner: sentry
--

CREATE INDEX sentry_team_organization_id ON sentry_team USING btree (organization_id);


--
-- Name: sentry_team_slug_like; Type: INDEX; Schema: public; Owner: sentry
--

CREATE INDEX sentry_team_slug_like ON sentry_team USING btree (slug varchar_pattern_ops);


--
-- Name: sentry_useravatar_ident_like; Type: INDEX; Schema: public; Owner: sentry
--

CREATE INDEX sentry_useravatar_ident_like ON sentry_useravatar USING btree (ident varchar_pattern_ops);


--
-- Name: sentry_useremail_user_id; Type: INDEX; Schema: public; Owner: sentry
--

CREATE INDEX sentry_useremail_user_id ON sentry_useremail USING btree (user_id);


--
-- Name: sentry_useroption_project_id; Type: INDEX; Schema: public; Owner: sentry
--

CREATE INDEX sentry_useroption_project_id ON sentry_useroption USING btree (project_id);


--
-- Name: sentry_useroption_user_id; Type: INDEX; Schema: public; Owner: sentry
--

CREATE INDEX sentry_useroption_user_id ON sentry_useroption USING btree (user_id);


--
-- Name: sentry_userreport_group_id; Type: INDEX; Schema: public; Owner: sentry
--

CREATE INDEX sentry_userreport_group_id ON sentry_userreport USING btree (group_id);


--
-- Name: sentry_userreport_project_id; Type: INDEX; Schema: public; Owner: sentry
--

CREATE INDEX sentry_userreport_project_id ON sentry_userreport USING btree (project_id);


--
-- Name: sentry_userreport_project_id_1ac377e052723c91; Type: INDEX; Schema: public; Owner: sentry
--

CREATE INDEX sentry_userreport_project_id_1ac377e052723c91 ON sentry_userreport USING btree (project_id, event_id);


--
-- Name: sentry_userreport_project_id_1c06c9ecc190b2e6; Type: INDEX; Schema: public; Owner: sentry
--

CREATE INDEX sentry_userreport_project_id_1c06c9ecc190b2e6 ON sentry_userreport USING btree (project_id, date_added);


--
-- Name: social_auth_usersocialauth_user_id; Type: INDEX; Schema: public; Owner: sentry
--

CREATE INDEX social_auth_usersocialauth_user_id ON social_auth_usersocialauth USING btree (user_id);


--
-- Name: sentry_auditlogentry actor_id_refs_id_cac0f7f5; Type: FK CONSTRAINT; Schema: public; Owner: sentry
--

ALTER TABLE ONLY sentry_auditlogentry
    ADD CONSTRAINT actor_id_refs_id_cac0f7f5 FOREIGN KEY (actor_id) REFERENCES auth_user(id) DEFERRABLE INITIALLY DEFERRED;


--
-- Name: sentry_auditlogentry actor_key_id_refs_id_cc2fc30c; Type: FK CONSTRAINT; Schema: public; Owner: sentry
--

ALTER TABLE ONLY sentry_auditlogentry
    ADD CONSTRAINT actor_key_id_refs_id_cc2fc30c FOREIGN KEY (actor_key_id) REFERENCES sentry_apikey(id) DEFERRABLE INITIALLY DEFERRED;


--
-- Name: auth_group_permissions auth_group_permissions_permission_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: sentry
--

ALTER TABLE ONLY auth_group_permissions
    ADD CONSTRAINT auth_group_permissions_permission_id_fkey FOREIGN KEY (permission_id) REFERENCES auth_permission(id) DEFERRABLE INITIALLY DEFERRED;


--
-- Name: sentry_authidentity auth_provider_id_refs_id_d9990f1d; Type: FK CONSTRAINT; Schema: public; Owner: sentry
--

ALTER TABLE ONLY sentry_authidentity
    ADD CONSTRAINT auth_provider_id_refs_id_d9990f1d FOREIGN KEY (auth_provider_id) REFERENCES sentry_authprovider(id) DEFERRABLE INITIALLY DEFERRED;


--
-- Name: sentry_commit author_id_refs_id_2f962e87; Type: FK CONSTRAINT; Schema: public; Owner: sentry
--

ALTER TABLE ONLY sentry_commit
    ADD CONSTRAINT author_id_refs_id_2f962e87 FOREIGN KEY (author_id) REFERENCES sentry_commitauthor(id) DEFERRABLE INITIALLY DEFERRED;


--
-- Name: sentry_authprovider_default_teams authprovider_id_refs_id_9e7068be; Type: FK CONSTRAINT; Schema: public; Owner: sentry
--

ALTER TABLE ONLY sentry_authprovider_default_teams
    ADD CONSTRAINT authprovider_id_refs_id_9e7068be FOREIGN KEY (authprovider_id) REFERENCES sentry_authprovider(id) DEFERRABLE INITIALLY DEFERRED;


--
-- Name: sentry_fileblobindex blob_id_refs_id_5732bcfb; Type: FK CONSTRAINT; Schema: public; Owner: sentry
--

ALTER TABLE ONLY sentry_fileblobindex
    ADD CONSTRAINT blob_id_refs_id_5732bcfb FOREIGN KEY (blob_id) REFERENCES sentry_fileblob(id) DEFERRABLE INITIALLY DEFERRED;


--
-- Name: sentry_file blob_id_refs_id_912b0028; Type: FK CONSTRAINT; Schema: public; Owner: sentry
--

ALTER TABLE ONLY sentry_file
    ADD CONSTRAINT blob_id_refs_id_912b0028 FOREIGN KEY (blob_id) REFERENCES sentry_fileblob(id) DEFERRABLE INITIALLY DEFERRED;


--
-- Name: sentry_broadcastseen broadcast_id_refs_id_e214087a; Type: FK CONSTRAINT; Schema: public; Owner: sentry
--

ALTER TABLE ONLY sentry_broadcastseen
    ADD CONSTRAINT broadcast_id_refs_id_e214087a FOREIGN KEY (broadcast_id) REFERENCES sentry_broadcast(id) DEFERRABLE INITIALLY DEFERRED;


--
-- Name: sentry_releasecommit commit_id_refs_id_a0857449; Type: FK CONSTRAINT; Schema: public; Owner: sentry
--

ALTER TABLE ONLY sentry_releasecommit
    ADD CONSTRAINT commit_id_refs_id_a0857449 FOREIGN KEY (commit_id) REFERENCES sentry_commit(id) DEFERRABLE INITIALLY DEFERRED;


--
-- Name: sentry_commitfilechange commit_id_refs_id_f9a55f94; Type: FK CONSTRAINT; Schema: public; Owner: sentry
--

ALTER TABLE ONLY sentry_commitfilechange
    ADD CONSTRAINT commit_id_refs_id_f9a55f94 FOREIGN KEY (commit_id) REFERENCES sentry_commit(id) DEFERRABLE INITIALLY DEFERRED;


--
-- Name: django_admin_log content_type_id_refs_id_93d2d1f8; Type: FK CONSTRAINT; Schema: public; Owner: sentry
--

ALTER TABLE ONLY django_admin_log
    ADD CONSTRAINT content_type_id_refs_id_93d2d1f8 FOREIGN KEY (content_type_id) REFERENCES django_content_type(id) DEFERRABLE INITIALLY DEFERRED;


--
-- Name: auth_permission content_type_id_refs_id_d043b34a; Type: FK CONSTRAINT; Schema: public; Owner: sentry
--

ALTER TABLE ONLY auth_permission
    ADD CONSTRAINT content_type_id_refs_id_d043b34a FOREIGN KEY (content_type_id) REFERENCES django_content_type(id) DEFERRABLE INITIALLY DEFERRED;


--
-- Name: sentry_useravatar file_id_refs_id_0c8678bd; Type: FK CONSTRAINT; Schema: public; Owner: sentry
--

ALTER TABLE ONLY sentry_useravatar
    ADD CONSTRAINT file_id_refs_id_0c8678bd FOREIGN KEY (file_id) REFERENCES sentry_file(id) DEFERRABLE INITIALLY DEFERRED;


--
-- Name: sentry_globaldsymfile file_id_refs_id_3efdc5e2; Type: FK CONSTRAINT; Schema: public; Owner: sentry
--

ALTER TABLE ONLY sentry_globaldsymfile
    ADD CONSTRAINT file_id_refs_id_3efdc5e2 FOREIGN KEY (file_id) REFERENCES sentry_file(id) DEFERRABLE INITIALLY DEFERRED;


--
-- Name: sentry_fileblobindex file_id_refs_id_82747ec9; Type: FK CONSTRAINT; Schema: public; Owner: sentry
--

ALTER TABLE ONLY sentry_fileblobindex
    ADD CONSTRAINT file_id_refs_id_82747ec9 FOREIGN KEY (file_id) REFERENCES sentry_file(id) DEFERRABLE INITIALLY DEFERRED;


--
-- Name: sentry_projectdsymfile file_id_refs_id_cc76204b; Type: FK CONSTRAINT; Schema: public; Owner: sentry
--

ALTER TABLE ONLY sentry_projectdsymfile
    ADD CONSTRAINT file_id_refs_id_cc76204b FOREIGN KEY (file_id) REFERENCES sentry_file(id) DEFERRABLE INITIALLY DEFERRED;


--
-- Name: sentry_releasefile file_id_refs_id_fb71e922; Type: FK CONSTRAINT; Schema: public; Owner: sentry
--

ALTER TABLE ONLY sentry_releasefile
    ADD CONSTRAINT file_id_refs_id_fb71e922 FOREIGN KEY (file_id) REFERENCES sentry_file(id) DEFERRABLE INITIALLY DEFERRED;


--
-- Name: sentry_groupedmessage first_release_id_refs_id_d035a570; Type: FK CONSTRAINT; Schema: public; Owner: sentry
--

ALTER TABLE ONLY sentry_groupedmessage
    ADD CONSTRAINT first_release_id_refs_id_d035a570 FOREIGN KEY (first_release_id) REFERENCES sentry_release(id) DEFERRABLE INITIALLY DEFERRED;


--
-- Name: sentry_groupseen group_id_refs_id_09b2694a; Type: FK CONSTRAINT; Schema: public; Owner: sentry
--

ALTER TABLE ONLY sentry_groupseen
    ADD CONSTRAINT group_id_refs_id_09b2694a FOREIGN KEY (group_id) REFERENCES sentry_groupedmessage(id) DEFERRABLE INITIALLY DEFERRED;


--
-- Name: sentry_messagefiltervalue group_id_refs_id_1fb6dc4e; Type: FK CONSTRAINT; Schema: public; Owner: sentry
--

ALTER TABLE ONLY sentry_messagefiltervalue
    ADD CONSTRAINT group_id_refs_id_1fb6dc4e FOREIGN KEY (group_id) REFERENCES sentry_groupedmessage(id) DEFERRABLE INITIALLY DEFERRED;


--
-- Name: sentry_groupbookmark group_id_refs_id_3738447a; Type: FK CONSTRAINT; Schema: public; Owner: sentry
--

ALTER TABLE ONLY sentry_groupbookmark
    ADD CONSTRAINT group_id_refs_id_3738447a FOREIGN KEY (group_id) REFERENCES sentry_groupedmessage(id) DEFERRABLE INITIALLY DEFERRED;


--
-- Name: sentry_groupemailthread group_id_refs_id_3c3dd283; Type: FK CONSTRAINT; Schema: public; Owner: sentry
--

ALTER TABLE ONLY sentry_groupemailthread
    ADD CONSTRAINT group_id_refs_id_3c3dd283 FOREIGN KEY (group_id) REFERENCES sentry_groupedmessage(id) DEFERRABLE INITIALLY DEFERRED;


--
-- Name: sentry_groupasignee group_id_refs_id_47b32b76; Type: FK CONSTRAINT; Schema: public; Owner: sentry
--

ALTER TABLE ONLY sentry_groupasignee
    ADD CONSTRAINT group_id_refs_id_47b32b76 FOREIGN KEY (group_id) REFERENCES sentry_groupedmessage(id) DEFERRABLE INITIALLY DEFERRED;


--
-- Name: sentry_grouprulestatus group_id_refs_id_66981850; Type: FK CONSTRAINT; Schema: public; Owner: sentry
--

ALTER TABLE ONLY sentry_grouprulestatus
    ADD CONSTRAINT group_id_refs_id_66981850 FOREIGN KEY (group_id) REFERENCES sentry_groupedmessage(id) DEFERRABLE INITIALLY DEFERRED;


--
-- Name: sentry_userreport group_id_refs_id_6b3d43d4; Type: FK CONSTRAINT; Schema: public; Owner: sentry
--

ALTER TABLE ONLY sentry_userreport
    ADD CONSTRAINT group_id_refs_id_6b3d43d4 FOREIGN KEY (group_id) REFERENCES sentry_groupedmessage(id) DEFERRABLE INITIALLY DEFERRED;


--
-- Name: sentry_groupmeta group_id_refs_id_6dc57728; Type: FK CONSTRAINT; Schema: public; Owner: sentry
--

ALTER TABLE ONLY sentry_groupmeta
    ADD CONSTRAINT group_id_refs_id_6dc57728 FOREIGN KEY (group_id) REFERENCES sentry_groupedmessage(id) DEFERRABLE INITIALLY DEFERRED;


--
-- Name: sentry_groupsnooze group_id_refs_id_7d70660e; Type: FK CONSTRAINT; Schema: public; Owner: sentry
--

ALTER TABLE ONLY sentry_groupsnooze
    ADD CONSTRAINT group_id_refs_id_7d70660e FOREIGN KEY (group_id) REFERENCES sentry_groupedmessage(id) DEFERRABLE INITIALLY DEFERRED;


--
-- Name: sentry_groupsubscription group_id_refs_id_901a3390; Type: FK CONSTRAINT; Schema: public; Owner: sentry
--

ALTER TABLE ONLY sentry_groupsubscription
    ADD CONSTRAINT group_id_refs_id_901a3390 FOREIGN KEY (group_id) REFERENCES sentry_groupedmessage(id) DEFERRABLE INITIALLY DEFERRED;


--
-- Name: sentry_grouphash group_id_refs_id_9603f6ba; Type: FK CONSTRAINT; Schema: public; Owner: sentry
--

ALTER TABLE ONLY sentry_grouphash
    ADD CONSTRAINT group_id_refs_id_9603f6ba FOREIGN KEY (group_id) REFERENCES sentry_groupedmessage(id) DEFERRABLE INITIALLY DEFERRED;


--
-- Name: sentry_activity group_id_refs_id_b84d67ec; Type: FK CONSTRAINT; Schema: public; Owner: sentry
--

ALTER TABLE ONLY sentry_activity
    ADD CONSTRAINT group_id_refs_id_b84d67ec FOREIGN KEY (group_id) REFERENCES sentry_groupedmessage(id) DEFERRABLE INITIALLY DEFERRED;


--
-- Name: sentry_grouptagkey group_id_refs_id_d78dfb94; Type: FK CONSTRAINT; Schema: public; Owner: sentry
--

ALTER TABLE ONLY sentry_grouptagkey
    ADD CONSTRAINT group_id_refs_id_d78dfb94 FOREIGN KEY (group_id) REFERENCES sentry_groupedmessage(id) DEFERRABLE INITIALLY DEFERRED;


--
-- Name: sentry_groupresolution group_id_refs_id_ed32932f; Type: FK CONSTRAINT; Schema: public; Owner: sentry
--

ALTER TABLE ONLY sentry_groupresolution
    ADD CONSTRAINT group_id_refs_id_ed32932f FOREIGN KEY (group_id) REFERENCES sentry_groupedmessage(id) DEFERRABLE INITIALLY DEFERRED;


--
-- Name: auth_group_permissions group_id_refs_id_f4b32aac; Type: FK CONSTRAINT; Schema: public; Owner: sentry
--

ALTER TABLE ONLY auth_group_permissions
    ADD CONSTRAINT group_id_refs_id_f4b32aac FOREIGN KEY (group_id) REFERENCES auth_group(id) DEFERRABLE INITIALLY DEFERRED;


--
-- Name: sentry_apitoken key_id_refs_id_b3eece75; Type: FK CONSTRAINT; Schema: public; Owner: sentry
--

ALTER TABLE ONLY sentry_apitoken
    ADD CONSTRAINT key_id_refs_id_b3eece75 FOREIGN KEY (key_id) REFERENCES sentry_apikey(id) DEFERRABLE INITIALLY DEFERRED;


--
-- Name: sentry_organizationaccessrequest member_id_refs_id_7c8ccc01; Type: FK CONSTRAINT; Schema: public; Owner: sentry
--

ALTER TABLE ONLY sentry_organizationaccessrequest
    ADD CONSTRAINT member_id_refs_id_7c8ccc01 FOREIGN KEY (member_id) REFERENCES sentry_organizationmember(id) DEFERRABLE INITIALLY DEFERRED;


--
-- Name: sentry_dsymsymbol object_id_refs_id_a9e69860; Type: FK CONSTRAINT; Schema: public; Owner: sentry
--

ALTER TABLE ONLY sentry_dsymsymbol
    ADD CONSTRAINT object_id_refs_id_a9e69860 FOREIGN KEY (object_id) REFERENCES sentry_dsymobject(id) DEFERRABLE INITIALLY DEFERRED;


--
-- Name: sentry_dsymbundle object_id_refs_id_be6e72e2; Type: FK CONSTRAINT; Schema: public; Owner: sentry
--

ALTER TABLE ONLY sentry_dsymbundle
    ADD CONSTRAINT object_id_refs_id_be6e72e2 FOREIGN KEY (object_id) REFERENCES sentry_dsymobject(id) DEFERRABLE INITIALLY DEFERRED;


--
-- Name: sentry_organizationonboardingtask organization_id_refs_id_2203c68b; Type: FK CONSTRAINT; Schema: public; Owner: sentry
--

ALTER TABLE ONLY sentry_organizationonboardingtask
    ADD CONSTRAINT organization_id_refs_id_2203c68b FOREIGN KEY (organization_id) REFERENCES sentry_organization(id) DEFERRABLE INITIALLY DEFERRED;


--
-- Name: sentry_organizationmember organization_id_refs_id_42dc8e8f; Type: FK CONSTRAINT; Schema: public; Owner: sentry
--

ALTER TABLE ONLY sentry_organizationmember
    ADD CONSTRAINT organization_id_refs_id_42dc8e8f FOREIGN KEY (organization_id) REFERENCES sentry_organization(id) DEFERRABLE INITIALLY DEFERRED;


--
-- Name: sentry_team organization_id_refs_id_61038a42; Type: FK CONSTRAINT; Schema: public; Owner: sentry
--

ALTER TABLE ONLY sentry_team
    ADD CONSTRAINT organization_id_refs_id_61038a42 FOREIGN KEY (organization_id) REFERENCES sentry_organization(id) DEFERRABLE INITIALLY DEFERRED;


--
-- Name: sentry_project organization_id_refs_id_6874e5b7; Type: FK CONSTRAINT; Schema: public; Owner: sentry
--

ALTER TABLE ONLY sentry_project
    ADD CONSTRAINT organization_id_refs_id_6874e5b7 FOREIGN KEY (organization_id) REFERENCES sentry_organization(id) DEFERRABLE INITIALLY DEFERRED;


--
-- Name: sentry_authprovider organization_id_refs_id_6a37632f; Type: FK CONSTRAINT; Schema: public; Owner: sentry
--

ALTER TABLE ONLY sentry_authprovider
    ADD CONSTRAINT organization_id_refs_id_6a37632f FOREIGN KEY (organization_id) REFERENCES sentry_organization(id) DEFERRABLE INITIALLY DEFERRED;


--
-- Name: sentry_organizationoptions organization_id_refs_id_83d34346; Type: FK CONSTRAINT; Schema: public; Owner: sentry
--

ALTER TABLE ONLY sentry_organizationoptions
    ADD CONSTRAINT organization_id_refs_id_83d34346 FOREIGN KEY (organization_id) REFERENCES sentry_organization(id) DEFERRABLE INITIALLY DEFERRED;


--
-- Name: sentry_apikey organization_id_refs_id_961ec303; Type: FK CONSTRAINT; Schema: public; Owner: sentry
--

ALTER TABLE ONLY sentry_apikey
    ADD CONSTRAINT organization_id_refs_id_961ec303 FOREIGN KEY (organization_id) REFERENCES sentry_organization(id) DEFERRABLE INITIALLY DEFERRED;


--
-- Name: sentry_release organization_id_refs_id_ba7f8e42; Type: FK CONSTRAINT; Schema: public; Owner: sentry
--

ALTER TABLE ONLY sentry_release
    ADD CONSTRAINT organization_id_refs_id_ba7f8e42 FOREIGN KEY (organization_id) REFERENCES sentry_organization(id) DEFERRABLE INITIALLY DEFERRED;


--
-- Name: sentry_releasefile organization_id_refs_id_ef2843cb; Type: FK CONSTRAINT; Schema: public; Owner: sentry
--

ALTER TABLE ONLY sentry_releasefile
    ADD CONSTRAINT organization_id_refs_id_ef2843cb FOREIGN KEY (organization_id) REFERENCES sentry_organization(id) DEFERRABLE INITIALLY DEFERRED;


--
-- Name: sentry_auditlogentry organization_id_refs_id_f5b1844e; Type: FK CONSTRAINT; Schema: public; Owner: sentry
--

ALTER TABLE ONLY sentry_auditlogentry
    ADD CONSTRAINT organization_id_refs_id_f5b1844e FOREIGN KEY (organization_id) REFERENCES sentry_organization(id) DEFERRABLE INITIALLY DEFERRED;


--
-- Name: sentry_organizationmember_teams organizationmember_id_refs_id_878802f4; Type: FK CONSTRAINT; Schema: public; Owner: sentry
--

ALTER TABLE ONLY sentry_organizationmember_teams
    ADD CONSTRAINT organizationmember_id_refs_id_878802f4 FOREIGN KEY (organizationmember_id) REFERENCES sentry_organizationmember(id) DEFERRABLE INITIALLY DEFERRED;


--
-- Name: sentry_release owner_id_refs_id_65604067; Type: FK CONSTRAINT; Schema: public; Owner: sentry
--

ALTER TABLE ONLY sentry_release
    ADD CONSTRAINT owner_id_refs_id_65604067 FOREIGN KEY (owner_id) REFERENCES auth_user(id) DEFERRABLE INITIALLY DEFERRED;


--
-- Name: sentry_grouprulestatus project_id_refs_id_09c5b95d; Type: FK CONSTRAINT; Schema: public; Owner: sentry
--

ALTER TABLE ONLY sentry_grouprulestatus
    ADD CONSTRAINT project_id_refs_id_09c5b95d FOREIGN KEY (project_id) REFERENCES sentry_project(id) DEFERRABLE INITIALLY DEFERRED;


--
-- Name: sentry_activity project_id_refs_id_0c94d99e; Type: FK CONSTRAINT; Schema: public; Owner: sentry
--

ALTER TABLE ONLY sentry_activity
    ADD CONSTRAINT project_id_refs_id_0c94d99e FOREIGN KEY (project_id) REFERENCES sentry_project(id) DEFERRABLE INITIALLY DEFERRED;


--
-- Name: sentry_groupbookmark project_id_refs_id_18390fbc; Type: FK CONSTRAINT; Schema: public; Owner: sentry
--

ALTER TABLE ONLY sentry_groupbookmark
    ADD CONSTRAINT project_id_refs_id_18390fbc FOREIGN KEY (project_id) REFERENCES sentry_project(id) DEFERRABLE INITIALLY DEFERRED;


--
-- Name: sentry_groupasignee project_id_refs_id_1b5200f8; Type: FK CONSTRAINT; Schema: public; Owner: sentry
--

ALTER TABLE ONLY sentry_groupasignee
    ADD CONSTRAINT project_id_refs_id_1b5200f8 FOREIGN KEY (project_id) REFERENCES sentry_project(id) DEFERRABLE INITIALLY DEFERRED;


--
-- Name: sentry_release project_id_refs_id_21d237e2; Type: FK CONSTRAINT; Schema: public; Owner: sentry
--

ALTER TABLE ONLY sentry_release
    ADD CONSTRAINT project_id_refs_id_21d237e2 FOREIGN KEY (project_id) REFERENCES sentry_project(id) DEFERRABLE INITIALLY DEFERRED;


--
-- Name: sentry_eventuser project_id_refs_id_2e8a33cc; Type: FK CONSTRAINT; Schema: public; Owner: sentry
--

ALTER TABLE ONLY sentry_eventuser
    ADD CONSTRAINT project_id_refs_id_2e8a33cc FOREIGN KEY (project_id) REFERENCES sentry_project(id) DEFERRABLE INITIALLY DEFERRED;


--
-- Name: sentry_savedsearch_userdefault project_id_refs_id_4bc1c005; Type: FK CONSTRAINT; Schema: public; Owner: sentry
--

ALTER TABLE ONLY sentry_savedsearch_userdefault
    ADD CONSTRAINT project_id_refs_id_4bc1c005 FOREIGN KEY (project_id) REFERENCES sentry_project(id) DEFERRABLE INITIALLY DEFERRED;


--
-- Name: sentry_projectcounter project_id_refs_id_58200d0a; Type: FK CONSTRAINT; Schema: public; Owner: sentry
--

ALTER TABLE ONLY sentry_projectcounter
    ADD CONSTRAINT project_id_refs_id_58200d0a FOREIGN KEY (project_id) REFERENCES sentry_project(id) DEFERRABLE INITIALLY DEFERRED;


--
-- Name: sentry_groupseen project_id_refs_id_67db0efd; Type: FK CONSTRAINT; Schema: public; Owner: sentry
--

ALTER TABLE ONLY sentry_groupseen
    ADD CONSTRAINT project_id_refs_id_67db0efd FOREIGN KEY (project_id) REFERENCES sentry_project(id) DEFERRABLE INITIALLY DEFERRED;


--
-- Name: sentry_grouphash project_id_refs_id_6f0a9434; Type: FK CONSTRAINT; Schema: public; Owner: sentry
--

ALTER TABLE ONLY sentry_grouphash
    ADD CONSTRAINT project_id_refs_id_6f0a9434 FOREIGN KEY (project_id) REFERENCES sentry_project(id) DEFERRABLE INITIALLY DEFERRED;


--
-- Name: sentry_userreport project_id_refs_id_723e0b3c; Type: FK CONSTRAINT; Schema: public; Owner: sentry
--

ALTER TABLE ONLY sentry_userreport
    ADD CONSTRAINT project_id_refs_id_723e0b3c FOREIGN KEY (project_id) REFERENCES sentry_project(id) DEFERRABLE INITIALLY DEFERRED;


--
-- Name: sentry_groupedmessage project_id_refs_id_77344b57; Type: FK CONSTRAINT; Schema: public; Owner: sentry
--

ALTER TABLE ONLY sentry_groupedmessage
    ADD CONSTRAINT project_id_refs_id_77344b57 FOREIGN KEY (project_id) REFERENCES sentry_project(id) DEFERRABLE INITIALLY DEFERRED;


--
-- Name: sentry_release_project project_id_refs_id_80894a1c; Type: FK CONSTRAINT; Schema: public; Owner: sentry
--

ALTER TABLE ONLY sentry_release_project
    ADD CONSTRAINT project_id_refs_id_80894a1c FOREIGN KEY (project_id) REFERENCES sentry_project(id) DEFERRABLE INITIALLY DEFERRED;


--
-- Name: sentry_groupemailthread project_id_refs_id_8419ea36; Type: FK CONSTRAINT; Schema: public; Owner: sentry
--

ALTER TABLE ONLY sentry_groupemailthread
    ADD CONSTRAINT project_id_refs_id_8419ea36 FOREIGN KEY (project_id) REFERENCES sentry_project(id) DEFERRABLE INITIALLY DEFERRED;


--
-- Name: sentry_releasefile project_id_refs_id_878696ea; Type: FK CONSTRAINT; Schema: public; Owner: sentry
--

ALTER TABLE ONLY sentry_releasefile
    ADD CONSTRAINT project_id_refs_id_878696ea FOREIGN KEY (project_id) REFERENCES sentry_project(id) DEFERRABLE INITIALLY DEFERRED;


--
-- Name: sentry_projectdsymfile project_id_refs_id_94d40917; Type: FK CONSTRAINT; Schema: public; Owner: sentry
--

ALTER TABLE ONLY sentry_projectdsymfile
    ADD CONSTRAINT project_id_refs_id_94d40917 FOREIGN KEY (project_id) REFERENCES sentry_project(id) DEFERRABLE INITIALLY DEFERRED;


--
-- Name: sentry_projectoptions project_id_refs_id_9b845024; Type: FK CONSTRAINT; Schema: public; Owner: sentry
--

ALTER TABLE ONLY sentry_projectoptions
    ADD CONSTRAINT project_id_refs_id_9b845024 FOREIGN KEY (project_id) REFERENCES sentry_project(id) DEFERRABLE INITIALLY DEFERRED;


--
-- Name: sentry_groupsubscription project_id_refs_id_a564d25b; Type: FK CONSTRAINT; Schema: public; Owner: sentry
--

ALTER TABLE ONLY sentry_groupsubscription
    ADD CONSTRAINT project_id_refs_id_a564d25b FOREIGN KEY (project_id) REFERENCES sentry_project(id) DEFERRABLE INITIALLY DEFERRED;


--
-- Name: sentry_savedsearch project_id_refs_id_b18120e7; Type: FK CONSTRAINT; Schema: public; Owner: sentry
--

ALTER TABLE ONLY sentry_savedsearch
    ADD CONSTRAINT project_id_refs_id_b18120e7 FOREIGN KEY (project_id) REFERENCES sentry_project(id) DEFERRABLE INITIALLY DEFERRED;


--
-- Name: sentry_messagefiltervalue project_id_refs_id_b5eaf112; Type: FK CONSTRAINT; Schema: public; Owner: sentry
--

ALTER TABLE ONLY sentry_messagefiltervalue
    ADD CONSTRAINT project_id_refs_id_b5eaf112 FOREIGN KEY (project_id) REFERENCES sentry_project(id) DEFERRABLE INITIALLY DEFERRED;


--
-- Name: sentry_filterkey project_id_refs_id_c385a797; Type: FK CONSTRAINT; Schema: public; Owner: sentry
--

ALTER TABLE ONLY sentry_filterkey
    ADD CONSTRAINT project_id_refs_id_c385a797 FOREIGN KEY (project_id) REFERENCES sentry_project(id) DEFERRABLE INITIALLY DEFERRED;


--
-- Name: sentry_rule project_id_refs_id_c96b69eb; Type: FK CONSTRAINT; Schema: public; Owner: sentry
--

ALTER TABLE ONLY sentry_rule
    ADD CONSTRAINT project_id_refs_id_c96b69eb FOREIGN KEY (project_id) REFERENCES sentry_project(id) DEFERRABLE INITIALLY DEFERRED;


--
-- Name: sentry_projectkey project_id_refs_id_e4d8a857; Type: FK CONSTRAINT; Schema: public; Owner: sentry
--

ALTER TABLE ONLY sentry_projectkey
    ADD CONSTRAINT project_id_refs_id_e4d8a857 FOREIGN KEY (project_id) REFERENCES sentry_project(id) DEFERRABLE INITIALLY DEFERRED;


--
-- Name: sentry_useroption project_id_refs_id_eb596317; Type: FK CONSTRAINT; Schema: public; Owner: sentry
--

ALTER TABLE ONLY sentry_useroption
    ADD CONSTRAINT project_id_refs_id_eb596317 FOREIGN KEY (project_id) REFERENCES sentry_project(id) DEFERRABLE INITIALLY DEFERRED;


--
-- Name: sentry_filtervalue project_id_refs_id_ee7bf50d; Type: FK CONSTRAINT; Schema: public; Owner: sentry
--

ALTER TABLE ONLY sentry_filtervalue
    ADD CONSTRAINT project_id_refs_id_ee7bf50d FOREIGN KEY (project_id) REFERENCES sentry_project(id) DEFERRABLE INITIALLY DEFERRED;


--
-- Name: sentry_grouptagkey project_id_refs_id_fef15df1; Type: FK CONSTRAINT; Schema: public; Owner: sentry
--

ALTER TABLE ONLY sentry_grouptagkey
    ADD CONSTRAINT project_id_refs_id_fef15df1 FOREIGN KEY (project_id) REFERENCES sentry_project(id) DEFERRABLE INITIALLY DEFERRED;


--
-- Name: sentry_groupresolution release_id_refs_id_0599bf90; Type: FK CONSTRAINT; Schema: public; Owner: sentry
--

ALTER TABLE ONLY sentry_groupresolution
    ADD CONSTRAINT release_id_refs_id_0599bf90 FOREIGN KEY (release_id) REFERENCES sentry_release(id) DEFERRABLE INITIALLY DEFERRED;


--
-- Name: sentry_releasecommit release_id_refs_id_26c8c7a0; Type: FK CONSTRAINT; Schema: public; Owner: sentry
--

ALTER TABLE ONLY sentry_releasecommit
    ADD CONSTRAINT release_id_refs_id_26c8c7a0 FOREIGN KEY (release_id) REFERENCES sentry_release(id) DEFERRABLE INITIALLY DEFERRED;


--
-- Name: sentry_releasefile release_id_refs_id_8c214aaf; Type: FK CONSTRAINT; Schema: public; Owner: sentry
--

ALTER TABLE ONLY sentry_releasefile
    ADD CONSTRAINT release_id_refs_id_8c214aaf FOREIGN KEY (release_id) REFERENCES sentry_release(id) DEFERRABLE INITIALLY DEFERRED;


--
-- Name: sentry_release_project release_id_refs_id_add4a457; Type: FK CONSTRAINT; Schema: public; Owner: sentry
--

ALTER TABLE ONLY sentry_release_project
    ADD CONSTRAINT release_id_refs_id_add4a457 FOREIGN KEY (release_id) REFERENCES sentry_release(id) DEFERRABLE INITIALLY DEFERRED;


--
-- Name: sentry_grouprulestatus rule_id_refs_id_39ff91f8; Type: FK CONSTRAINT; Schema: public; Owner: sentry
--

ALTER TABLE ONLY sentry_grouprulestatus
    ADD CONSTRAINT rule_id_refs_id_39ff91f8 FOREIGN KEY (rule_id) REFERENCES sentry_rule(id) DEFERRABLE INITIALLY DEFERRED;


--
-- Name: sentry_savedsearch_userdefault savedsearch_id_refs_id_8d85995b; Type: FK CONSTRAINT; Schema: public; Owner: sentry
--

ALTER TABLE ONLY sentry_savedsearch_userdefault
    ADD CONSTRAINT savedsearch_id_refs_id_8d85995b FOREIGN KEY (savedsearch_id) REFERENCES sentry_savedsearch(id) DEFERRABLE INITIALLY DEFERRED;


--
-- Name: sentry_dsymbundle sdk_id_refs_id_cf47df8d; Type: FK CONSTRAINT; Schema: public; Owner: sentry
--

ALTER TABLE ONLY sentry_dsymbundle
    ADD CONSTRAINT sdk_id_refs_id_cf47df8d FOREIGN KEY (sdk_id) REFERENCES sentry_dsymsdk(id) DEFERRABLE INITIALLY DEFERRED;


--
-- Name: sentry_auditlogentry target_user_id_refs_id_cac0f7f5; Type: FK CONSTRAINT; Schema: public; Owner: sentry
--

ALTER TABLE ONLY sentry_auditlogentry
    ADD CONSTRAINT target_user_id_refs_id_cac0f7f5 FOREIGN KEY (target_user_id) REFERENCES auth_user(id) DEFERRABLE INITIALLY DEFERRED;


--
-- Name: sentry_authprovider_default_teams team_id_refs_id_10a85f7b; Type: FK CONSTRAINT; Schema: public; Owner: sentry
--

ALTER TABLE ONLY sentry_authprovider_default_teams
    ADD CONSTRAINT team_id_refs_id_10a85f7b FOREIGN KEY (team_id) REFERENCES sentry_team(id) DEFERRABLE INITIALLY DEFERRED;


--
-- Name: sentry_project team_id_refs_id_78750968; Type: FK CONSTRAINT; Schema: public; Owner: sentry
--

ALTER TABLE ONLY sentry_project
    ADD CONSTRAINT team_id_refs_id_78750968 FOREIGN KEY (team_id) REFERENCES sentry_team(id) DEFERRABLE INITIALLY DEFERRED;


--
-- Name: sentry_organizationmember_teams team_id_refs_id_d98f2858; Type: FK CONSTRAINT; Schema: public; Owner: sentry
--

ALTER TABLE ONLY sentry_organizationmember_teams
    ADD CONSTRAINT team_id_refs_id_d98f2858 FOREIGN KEY (team_id) REFERENCES sentry_team(id) DEFERRABLE INITIALLY DEFERRED;


--
-- Name: sentry_organizationaccessrequest team_id_refs_id_ea6e538b; Type: FK CONSTRAINT; Schema: public; Owner: sentry
--

ALTER TABLE ONLY sentry_organizationaccessrequest
    ADD CONSTRAINT team_id_refs_id_ea6e538b FOREIGN KEY (team_id) REFERENCES sentry_team(id) DEFERRABLE INITIALLY DEFERRED;


--
-- Name: sentry_groupbookmark user_id_refs_id_05ac45cc; Type: FK CONSTRAINT; Schema: public; Owner: sentry
--

ALTER TABLE ONLY sentry_groupbookmark
    ADD CONSTRAINT user_id_refs_id_05ac45cc FOREIGN KEY (user_id) REFERENCES auth_user(id) DEFERRABLE INITIALLY DEFERRED;


--
-- Name: sentry_useravatar user_id_refs_id_1a689f2e; Type: FK CONSTRAINT; Schema: public; Owner: sentry
--

ALTER TABLE ONLY sentry_useravatar
    ADD CONSTRAINT user_id_refs_id_1a689f2e FOREIGN KEY (user_id) REFERENCES auth_user(id) DEFERRABLE INITIALLY DEFERRED;


--
-- Name: sentry_organizationonboardingtask user_id_refs_id_22c181a4; Type: FK CONSTRAINT; Schema: public; Owner: sentry
--

ALTER TABLE ONLY sentry_organizationonboardingtask
    ADD CONSTRAINT user_id_refs_id_22c181a4 FOREIGN KEY (user_id) REFERENCES auth_user(id) DEFERRABLE INITIALLY DEFERRED;


--
-- Name: sentry_groupseen user_id_refs_id_270b7315; Type: FK CONSTRAINT; Schema: public; Owner: sentry
--

ALTER TABLE ONLY sentry_groupseen
    ADD CONSTRAINT user_id_refs_id_270b7315 FOREIGN KEY (user_id) REFERENCES auth_user(id) DEFERRABLE INITIALLY DEFERRED;


--
-- Name: sentry_projectbookmark user_id_refs_id_32679665; Type: FK CONSTRAINT; Schema: public; Owner: sentry
--

ALTER TABLE ONLY sentry_projectbookmark
    ADD CONSTRAINT user_id_refs_id_32679665 FOREIGN KEY (user_id) REFERENCES auth_user(id) DEFERRABLE INITIALLY DEFERRED;


--
-- Name: sentry_savedsearch_userdefault user_id_refs_id_3f7101ca; Type: FK CONSTRAINT; Schema: public; Owner: sentry
--

ALTER TABLE ONLY sentry_savedsearch_userdefault
    ADD CONSTRAINT user_id_refs_id_3f7101ca FOREIGN KEY (user_id) REFERENCES auth_user(id) DEFERRABLE INITIALLY DEFERRED;


--
-- Name: sentry_broadcastseen user_id_refs_id_5d9e5ad9; Type: FK CONSTRAINT; Schema: public; Owner: sentry
--

ALTER TABLE ONLY sentry_broadcastseen
    ADD CONSTRAINT user_id_refs_id_5d9e5ad9 FOREIGN KEY (user_id) REFERENCES auth_user(id) DEFERRABLE INITIALLY DEFERRED;


--
-- Name: sentry_activity user_id_refs_id_6caec40e; Type: FK CONSTRAINT; Schema: public; Owner: sentry
--

ALTER TABLE ONLY sentry_activity
    ADD CONSTRAINT user_id_refs_id_6caec40e FOREIGN KEY (user_id) REFERENCES auth_user(id) DEFERRABLE INITIALLY DEFERRED;


--
-- Name: sentry_useroption user_id_refs_id_73734413; Type: FK CONSTRAINT; Schema: public; Owner: sentry
--

ALTER TABLE ONLY sentry_useroption
    ADD CONSTRAINT user_id_refs_id_73734413 FOREIGN KEY (user_id) REFERENCES auth_user(id) DEFERRABLE INITIALLY DEFERRED;


--
-- Name: sentry_authidentity user_id_refs_id_78163ab5; Type: FK CONSTRAINT; Schema: public; Owner: sentry
--

ALTER TABLE ONLY sentry_authidentity
    ADD CONSTRAINT user_id_refs_id_78163ab5 FOREIGN KEY (user_id) REFERENCES auth_user(id) DEFERRABLE INITIALLY DEFERRED;


--
-- Name: sentry_apitoken user_id_refs_id_78c75ee2; Type: FK CONSTRAINT; Schema: public; Owner: sentry
--

ALTER TABLE ONLY sentry_apitoken
    ADD CONSTRAINT user_id_refs_id_78c75ee2 FOREIGN KEY (user_id) REFERENCES auth_user(id) DEFERRABLE INITIALLY DEFERRED;


--
-- Name: auth_authenticator user_id_refs_id_8e85b45f; Type: FK CONSTRAINT; Schema: public; Owner: sentry
--

ALTER TABLE ONLY auth_authenticator
    ADD CONSTRAINT user_id_refs_id_8e85b45f FOREIGN KEY (user_id) REFERENCES auth_user(id) DEFERRABLE INITIALLY DEFERRED;


--
-- Name: sentry_useremail user_id_refs_id_ae956867; Type: FK CONSTRAINT; Schema: public; Owner: sentry
--

ALTER TABLE ONLY sentry_useremail
    ADD CONSTRAINT user_id_refs_id_ae956867 FOREIGN KEY (user_id) REFERENCES auth_user(id) DEFERRABLE INITIALLY DEFERRED;


--
-- Name: sentry_organizationmember user_id_refs_id_be455e60; Type: FK CONSTRAINT; Schema: public; Owner: sentry
--

ALTER TABLE ONLY sentry_organizationmember
    ADD CONSTRAINT user_id_refs_id_be455e60 FOREIGN KEY (user_id) REFERENCES auth_user(id) DEFERRABLE INITIALLY DEFERRED;


--
-- Name: sentry_lostpasswordhash user_id_refs_id_c60bdf9b; Type: FK CONSTRAINT; Schema: public; Owner: sentry
--

ALTER TABLE ONLY sentry_lostpasswordhash
    ADD CONSTRAINT user_id_refs_id_c60bdf9b FOREIGN KEY (user_id) REFERENCES auth_user(id) DEFERRABLE INITIALLY DEFERRED;


--
-- Name: social_auth_usersocialauth user_id_refs_id_e6cbdf29; Type: FK CONSTRAINT; Schema: public; Owner: sentry
--

ALTER TABLE ONLY social_auth_usersocialauth
    ADD CONSTRAINT user_id_refs_id_e6cbdf29 FOREIGN KEY (user_id) REFERENCES auth_user(id) DEFERRABLE INITIALLY DEFERRED;


--
-- Name: sentry_groupsubscription user_id_refs_id_efb4b379; Type: FK CONSTRAINT; Schema: public; Owner: sentry
--

ALTER TABLE ONLY sentry_groupsubscription
    ADD CONSTRAINT user_id_refs_id_efb4b379 FOREIGN KEY (user_id) REFERENCES auth_user(id) DEFERRABLE INITIALLY DEFERRED;


--
-- Name: sentry_groupasignee user_id_refs_id_f4dcb8d1; Type: FK CONSTRAINT; Schema: public; Owner: sentry
--

ALTER TABLE ONLY sentry_groupasignee
    ADD CONSTRAINT user_id_refs_id_f4dcb8d1 FOREIGN KEY (user_id) REFERENCES auth_user(id) DEFERRABLE INITIALLY DEFERRED;


--
-- PostgreSQL database dump complete
--

