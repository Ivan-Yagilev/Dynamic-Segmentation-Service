SELECT u.username AS username, s.segmentname AS segmentname
FROM users u
LEFT JOIN user_segment us ON us.user_id = u.id
LEFT JOIN segments s ON s.id = us.segment_id;