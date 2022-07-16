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

create index title on movies (title);
create index genre_name on genres (name);
create index production_name on genres (name);

create view movie_genre_view as
select movie_id as id,
       genre_id,
       name,
       adult,
       budget,
       homepage,
       imdb_id,
       overview,
       popularity,
       poster,
       release_date,
       revenue,
       runtime,
       status,
       title,
       vote_average,
       vote_count
from movie_genres
         join genres g on movie_genres.genre_id = g.id
         join movies m on movie_genres.movie_id = m.id;


create view movie_production_view as
select movie_id as id,
       production_id,
       adult,
       budget,
       homepage,
       imdb_id,
       overview,
       popularity,
       poster,
       release_date,
       revenue,
       runtime,
       status,
       title,
       vote_average,
       vote_count,
       name
from movie_productions
         join movies m on movie_productions.movie_id = m.id
         join productions p on movie_productions.production_id = p.id;
