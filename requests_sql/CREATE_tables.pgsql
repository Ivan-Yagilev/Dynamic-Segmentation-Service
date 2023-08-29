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
    user_id   	 INTEGER NOT NULL REFERENCES users ON DELETE CASCADE,
    segment_id   INTEGER NOT NULL REFERENCES segments ON DELETE CASCADE,
    UNIQUE (user_id, segment_id)
);
