CREATE TABLE `tag` (
  `id` varchar(255) PRIMARY KEY,
  `name` varchar(255)
);

CREATE TABLE `tag_player` (
  `player_id` integer,
  `tag_id` integer,
  PRIMARY KEY (`player_id`, `tag_id`)
);

CREATE TABLE `tag_team` (
  `team_id` integer,
  `tag_id` integer,
  PRIMARY KEY (`team_id`, `tag_id`)
);

CREATE TABLE `stadium` (
  `id` integer PRIMARY KEY,
  `name` varchar(255)
);

CREATE TABLE `player` (
  `id` integer PRIMARY KEY,
  `name` varchar(255),
  `last_name` varchar(255),
  `birth_date` timestamp,
  `team_id` integer,
  `number` integer
);

CREATE TABLE `team` (
  `id` integer PRIMARY KEY,
  `name` varchar(255),
  `birth_date` timestamp,
  `category` varchar(255)
);

CREATE TABLE `league` (
  `id` integer PRIMARY KEY,
  `name` varchar(255),
  `birth_date` timestamp
);

CREATE TABLE `season` (
  `id` integer PRIMARY KEY,
  `league_id` integer,
  `name` varchar(255),
  `stars_at` timestamp,
  `ends_at` timestamp,
  `status` integer
);

CREATE TABLE `season_team` (
  `season_id` integer,
  `team_id` integer,
  PRIMARY KEY (`season_id`, `team_id`)
);

CREATE TABLE `match` (
  `id` integer PRIMARY KEY,
  `home_team_id` integer,
  `away_team_id` integer,
  `season_id` integer,
  `stadium_id` integer,
  `date` timestamp,
  `home_team_score` integer,
  `away_team_score` integer,
  `home_team_points` integer,
  `away_team_points` integer,
  `stage` varchar(255),
  `observation` text
);

CREATE TABLE `match_player` (
  `id` integer PRIMARY KEY,
  `match_id` integer,
  `team_id` integer,
  `player_id` integer,
  `red_car` integer,
  `yellow_card` integer,
  `goals` integer
);

ALTER TABLE `tag_player` ADD FOREIGN KEY (`tag_id`) REFERENCES `tag` (`id`);

ALTER TABLE `tag_player` ADD FOREIGN KEY (`player_id`) REFERENCES `player` (`id`);

ALTER TABLE `tag_team` ADD FOREIGN KEY (`tag_id`) REFERENCES `tag` (`id`);

ALTER TABLE `tag_team` ADD FOREIGN KEY (`team_id`) REFERENCES `team` (`id`);

ALTER TABLE `player` ADD FOREIGN KEY (`team_id`) REFERENCES `team` (`id`);

ALTER TABLE `match` ADD FOREIGN KEY (`home_team_id`) REFERENCES `team` (`id`);

ALTER TABLE `match` ADD FOREIGN KEY (`away_team_id`) REFERENCES `team` (`id`);

ALTER TABLE `match` ADD FOREIGN KEY (`season_id`) REFERENCES `season` (`id`);

ALTER TABLE `match` ADD FOREIGN KEY (`stadium_id`) REFERENCES `stadium` (`id`);

ALTER TABLE `season` ADD FOREIGN KEY (`league_id`) REFERENCES `league` (`id`);

ALTER TABLE `season_team` ADD FOREIGN KEY (`season_id`) REFERENCES `season` (`id`);

ALTER TABLE `season_team` ADD FOREIGN KEY (`team_id`) REFERENCES `team` (`id`);
