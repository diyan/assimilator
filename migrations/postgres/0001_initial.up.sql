-- todo consider bigserial, serial instead of sequence
--   see http://stackoverflow.com/questions/20781111/postgresql-9-1-primary-key-autoincrement

create table sentry_project (
    id integer not null,
    name character varying(200) not null,
    public boolean not null,
    date_added timestamp with time zone not null,
    status integer not null,
    slug character varying(50),
    team_id integer not null,
    organization_id integer not null,
    first_event timestamp with time zone,
    forced_color character varying(6),
    constraint ck_status_pstv_3af8360b8a37db73 check ((status >= 0)),
    constraint sentry_project_status_check check ((status >= 0))
);
create sequence sentry_project_id_seq
    start with 1
    increment by 1
    no minvalue
    no maxvalue
    cache 1;
alter sequence sentry_project_id_seq owned by sentry_project.id;



create table sentry_filterkey (
    id integer not null,
    project_id integer not null,
    key varchar(32) not null,
    values_seen integer not null,
    label varchar(64),
    status integer not null,
    constraint ck_status_pstv_56aaa5973127b013 check ((status >= 0)),
    constraint ck_values_seen_pstv_12eab0d3ff94a35c check ((values_seen >= 0)),
    constraint sentry_filterkey_status_check check ((status >= 0)),
    constraint sentry_filterkey_values_seen_check check ((values_seen >= 0))
);
create sequence sentry_filterkey_id_seq
    start with 1
    increment by 1
    no minvalue
    no maxvalue
    cache 1;
alter sequence sentry_filterkey_id_seq owned by sentry_filterkey.id;



create table sentry_organization (
    id integer not null,
    name character varying(64) not null,
    status integer not null,
    date_added timestamp with time zone not null,
    slug character varying(50) not null,
    flags bigint not null,
    default_role character varying(32) not null,
    constraint sentry_organization_status_check check ((status >= 0))
);
create sequence sentry_organization_id_seq
    start with 1
    increment by 1
    no minvalue
    no maxvalue
    cache 1;
alter sequence sentry_organization_id_seq owned by sentry_organization.id;
