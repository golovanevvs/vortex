-- Таблица: Activity
-- +--------------+---------+
-- | Column Name  | Type    |
-- +--------------+---------+
-- | player_id    | int     |
-- | device_id    | int     |
-- | event_date   | date    |
-- | games_played | int     |
-- +--------------+---------+
 
-- (player_id, event_date) является первичным ключом (комбинацией столбцов со уникальными значениями) этой таблицы.  
-- Эта таблица отображает активность игроков в некоторых играх.  
-- Каждая строка — это запись игрока, который вошёл в систему и сыграл определённое количество игр (возможно, 0), а затем вышел в какой-то день, используя какое-то устройство.
 
-- Задача:
-- - Найдите среднее количество игр (games_played), сыгранных игроками в их первый день входа.
-- - Округлите результат до 2 знаков после запятой.
-- 
-- 
-- Пример:
-- 
-- +-----------+------------+--------------+
-- | player_id | event_date | games_played |
-- +-----------+------------+--------------+
-- | 1         | 2016-03-01 | 5            |  ← первый день игрока 1
-- | 1         | 2016-03-02 | 6            |
-- | 2         | 2017-06-25 | 1            |  ← первый день игрока 2
-- | 3         | 2016-03-02 | 0            |  ← первый день игрока 3
-- | 3         | 2018-07-03 | 5            |
-- +-----------+------------+--------------+
-- 
-- 
-- Результат:
-- 
-- +--------+
-- | avg    |
-- +--------+
-- | 2.00   |  -- (5 + 1 + 0) / 3 = 2.00
-- +--------+

WITH
    first_date AS (
        SELECT
            player_id,
            MIN(event_date) AS fed
        FROM activity
        GROUP BY player_id
    )
SELECT
    ROUND(AVG(games_played), 2) AS avg
FROM
    activity
    JOIN first_date ON
        activity.player_id = first_date.player_id AND
        activity.event_date = first_date.fed

SELECT ROUND(AVG(fgp),2) AS avg
FROM (
    SELECT 
        DISTINCT player_id,
        FIRST_VALUE(games_played) OVER (PARTITION BY player_id ORDER BY event_date) AS fgp
    FROM activity
) temp