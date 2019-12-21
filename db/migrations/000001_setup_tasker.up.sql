BEGIN;

CREATE TABLE plugins (
    id SERIAL PRIMARY KEY,
    url text NOT NULL UNIQUE,
    type varchar(256) NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMPTZ
);

CREATE TABLE plugin_versions (
    id SERIAL PRIMARY KEY,
    version text NOT NULL,
    plugin_id integer,
    FOREIGN KEY(plugin_id) REFERENCES plugins (id),
    created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMPTZ
);

CREATE TABLE scheduled_crons (
    id SERIAL PRIMARY KEY,
    cron_id numeric UNIQUE,
    created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMPTZ
);

CREATE TABLE tasks (
    id SERIAL PRIMARY KEY,
    name VARCHAR(256) NOT NULL,
    description text,
    args json NOT NULL,
    schedule text NOT NULL,
    executor text NOT NULL,
    cron_id integer,
    FOREIGN KEY(cron_id) REFERENCES scheduled_crons (id),
    plugin_version_id integer NOT NULL,
    FOREIGN KEY(plugin_version_id) REFERENCES plugin_versions (id) ON DELETE CASCADE,
    repeatable boolean NOT NULL DEFAULT false,
    enabled boolean NOT NULL DEFAULT true,
    complete boolean NOT NULL DEFAULT false,
    completed_at TIMESTAMPTZ,
    created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMPTZ
);

CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    email text NOT NULL UNIQUE,
    username text NOT NULL,
    access_token text NOT NULL,
    bio text,
    github_url text NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMPTZ
);

CREATE TABLE teams (
    id SERIAL PRIMARY KEY,
    name VARCHAR(256) NOT NULL UNIQUE,
    created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMPTZ
);

CREATE TABLE users_teams (
    user_id integer NOT NULL,
    team_id integer NOT NULL,
    PRIMARY KEY (user_id, team_id),
    FOREIGN KEY (user_id) REFERENCES users(id) ON UPDATE CASCADE,
    FOREIGN KEY (team_id) REFERENCES teams(id) ON UPDATE CASCADE
);

CREATE TABLE task_access (
    task_id integer NOT NULL,
    team_id integer NOT NULL,
    PRIMARY KEY (task_id, team_id),
    FOREIGN KEY (task_id) REFERENCES tasks(id) ON UPDATE CASCADE,
    FOREIGN KEY (team_id) REFERENCES teams(id) ON UPDATE CASCADE
);

COMMIT;