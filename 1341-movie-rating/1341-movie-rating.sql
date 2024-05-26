# Write your MySQL query statement below
(
    SELECT u.name AS results
    FROM MovieRating mr JOIN Users u ON mr.user_id = u.user_id
    GROUP BY mr.user_id
    ORDER BY COUNT(u.user_id) DESC, u.name ASC
    LIMIT 1
)
UNION ALL
(
    SELECT m.title AS results
    FROM MovieRating mr JOIN Movies m ON mr.movie_id = m.movie_id
    WHERE mr.created_at LIKE '2020-02%'
    GROUP BY mr.movie_id
    ORDER BY AVG(rating) DESC, m.title
    LIMIT 1
)
