create table auth_user (
    id serial not null primary key,
    password character varying(128) not null,
    last_login timestamp with time zone not null,
    username character varying(128) not null,
    first_name character varying(200) not null,
    email character varying(75) not null,
    is_staff boolean not null,
    is_active boolean not null,
    is_superuser boolean not null,
    date_joined timestamp with time zone not null,
    is_managed boolean not null,
    is_password_expired boolean not null,
    last_password_change timestamp with time zone,
    session_nonce character varying(12)
);


create table sentry_organization (
    id serial not null primary key,
    name character varying(64) not null,
    status integer not null,
    date_added timestamp with time zone not null,
    slug character varying(50) not null,
    flags bigint not null,
    default_role character varying(32) not null,
    constraint sentry_organization_status_check check ((status >= 0))
);


create table sentry_organizationmember (
    id serial not null primary key,
    organization_id integer not null references sentry_organization(id),
    user_id integer references auth_user(id),
    type integer not null,
    date_added timestamp with time zone not null,
    email character varying(75),
    has_global_access boolean not null,
    flags bigint not null,
    role character varying(32) not null,
    token character varying(64),
    constraint sentry_organizationmember_type_check check ((type >= 0))
);


create table sentry_organizationoptions (
    id serial not null primary key,
    organization_id integer not null references sentry_organization(id),
    key character varying(64) not null,
    value text not null
);


create table sentry_team (
    id serial not null primary key,
    slug character varying(50) not null,
    name character varying(64) not null,
    date_added timestamp with time zone,
    status integer not null,
    organization_id integer not null references sentry_organization(id),
    constraint ck_status_pstv_1772e42d30eba7ba check ((status >= 0)),
    constraint sentry_team_status_check check ((status >= 0))
);


create table sentry_organizationmember_teams (
    id serial not null primary key,
    organizationmember_id integer not null,
    team_id integer not null references sentry_team(id),
    is_active boolean not null
);


create table sentry_organizationaccessrequest (
    id serial not null primary key,
    team_id integer not null references sentry_team(id),
    member_id integer not null
);


create table sentry_project (
    id serial not null primary key,
    name character varying(200) not null,
    public boolean not null,
    date_added timestamp with time zone not null,
    status integer not null,
    slug character varying(50),
    team_id integer not null references sentry_team(id),
    organization_id integer not null references sentry_organization(id),
    first_event timestamp with time zone,
    forced_color character varying(6),
    constraint ck_status_pstv_3af8360b8a37db73 check ((status >= 0)),
    constraint sentry_project_status_check check ((status >= 0))
);


create table sentry_environment (
    id serial not null primary key,
    project_id integer not null references sentry_project(id),
    name character varying(64) not null,
    date_added timestamp with time zone not null,
    constraint sentry_environment_project_id_check check ((project_id >= 0))
);


create table sentry_savedsearch (
    id serial not null primary key,
    project_id integer not null references sentry_project(id),
    name character varying(128) not null,
    query text not null,
    date_added timestamp with time zone not null,
    is_default boolean not null
);


create table sentry_filterkey (
    id serial not null primary key,
    project_id integer not null references sentry_project(id),
    key varchar(32) not null,
    values_seen integer not null,
    label varchar(64),
    status integer not null,
    constraint ck_status_pstv_56aaa5973127b013 check ((status >= 0)),
    constraint ck_values_seen_pstv_12eab0d3ff94a35c check ((values_seen >= 0)),
    constraint sentry_filterkey_status_check check ((status >= 0)),
    constraint sentry_filterkey_values_seen_check check ((values_seen >= 0))
);


create table sentry_groupedmessage (
    id serial not null primary key,
    logger character varying(64) not null,
    level integer not null,
    message text not null,
    view character varying(200),
    status integer not null,
    times_seen integer not null,
    last_seen timestamp with time zone not null,
    first_seen timestamp with time zone not null,
    data text,
    score integer not null,
    project_id integer references sentry_project(id),
    time_spent_total integer not null,
    time_spent_count integer not null,
    resolved_at timestamp with time zone,
    active_at timestamp with time zone,
    is_public boolean,
    platform character varying(64),
    num_comments integer,
    first_release_id integer,
    short_id integer,
    constraint ck_num_comments_pstv_44851d4d5d739eab check ((num_comments >= 0)),
    constraint sentry_groupedmessage_level_check check ((level >= 0)),
    constraint sentry_groupedmessage_num_comments_check check ((num_comments >= 0)),
    constraint sentry_groupedmessage_status_check check ((status >= 0)),
    constraint sentry_groupedmessage_times_seen_check check ((times_seen >= 0))
);


create table sentry_message (
    id serial not null primary key,
    message text not null,
    datetime timestamp with time zone not null,
    data text,
    group_id integer references sentry_groupedmessage(id),
    message_id character varying(32),
    project_id integer references sentry_project(id),
    time_spent integer,
    platform character varying(64)
);


create table nodestore_node (
    id character varying(40) not null primary key,
    data text not null,
    "timestamp" timestamp with time zone not null
);
