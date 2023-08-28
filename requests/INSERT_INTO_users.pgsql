INSERT INTO users (username, password)
VALUES ('nikita', '14614634'),
       ('denis', '345645312'),
       ('ivan', '36346513'),
       ('sergey', '75756')
RETURNING id;

INSERT INTO segments (segmentname)
VALUES ('AVITO_VOICE_MESSAGES'),
       ('AVITO_PERFORMANCE_VAS'),
       ('AVITO_DISCOUNT_30'),
       ('AVITO_DISCOUNT_50')
RETURNING id;

INSERT INTO user_segment (user_id, segment_id)
VALUES (1, 1),
       (1, 3),
       (2, 2),
       (2, 3),
       (2, 4),
       (3, 1),
       (3, 4)
RETURNING id;