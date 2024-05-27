# Write your MySQL query statement below
SELECT 
    x, y, z,
    IF(x+y>z AND y+z>x and z+x>y, 'Yes', 'No') as triangle
FROM
    triangle
