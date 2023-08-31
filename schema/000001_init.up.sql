CREATE TABLE users
(
	id 	 SERIAL  PRIMARY KEY,
	username VARCHAR ( 50 ) UNIQUE NOT NULL,
	password VARCHAR NOT NULL
);

CREATE TABLE segments
(
    id   	SERIAL  PRIMARY KEY,
    segmentname VARCHAR ( 50 ) NOT NULL
);

CREATE TABLE user_segment
(
    id        	 SERIAL  PRIMARY KEY,
    user_id   	 INT REFERENCES users (id) ON DELETE CASCADE NOT NULL,
    segment_id   INT REFERENCES segments (id) ON DELETE CASCADE NOT NULL,
    UNIQUE (user_id, segment_id)
);
