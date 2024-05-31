DROP TABLE  IF EXISTS boardgames;

CREATE TABLE boardgames (
    id SERIAL PRIMARY KEY,
    title VARCHAR(200) NOT NULL,
    short_desc VARCHAR(200),
    long_desc TEXT,
    release_year INTEGER,
    min_player INTEGER,
    max_player INTEGER,
    min_time INTEGER,
    max_time INTEGER,
    age INTEGER
);

INSERT INTO boardgames (title, short_desc, long_desc, release_year, min_player, max_player, min_time, max_time, age) VALUES
('Calico', 'Sew a quilt, collect buttons, attract cats!', 'Calico is a puzzly tile-laying game of quilts and cats.

In Calico, players compete to sew the coziest quilt as they collect and place patches of different colors and patterns. Each quilt has a particular pattern that must be followed, and players are also trying to create color and pattern combinations that are not only aesthetically pleasing, but also able to attract the cuddliest cats!

Turns are simple. Select a single patch tile from your hand and sew it into your quilt, then draw another patch into your hand from the three available. If you are able to create a color group, you may sew a button onto your quilt. If you are able to create a pattern combination that is attractive to any of the cats, it will come over and curl up on your quilt! At the end of the game, you score points for buttons, cats, and how well you were able to complete your unique quilt pattern.

—description from the publisher', 2020, 1, 4, 30, 45, 10);