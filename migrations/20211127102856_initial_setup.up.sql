create table genres
(
    id   int     not null primary key,
    name varchar not null
);

create table productions
(
    id   int     not null primary key,
    name varchar not null
);

create table movies
(
    id           serial primary key,
    adult        bool not null,
    budget       bigint,
    homepage     varchar,
    imdb_id      varchar,
    overview     varchar,
    popularity   float,
    poster       varchar,
    release_date timestamp,
    revenue      bigint,
    runtime      float,
    status       varchar,
    title        varchar,
    vote_average float,
    vote_count   int
);

create table movie_genres
(
    id       serial primary key,
    movie_id int references movies (id) not null,
    genre_id int references genres (id) not null
);

create table movie_productions
(
    id            serial primary key,
    movie_id      int references movies (id)      not null,
    production_id int references productions (id) not null
);

create index title on movies(title);
