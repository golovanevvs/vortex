-- Таблица: Activity 
-- +--------------+---------+ 
-- | Column Name | Type | 
-- +--------------+---------+ 
-- | player_id | int | 
-- | device_id | int | 
-- | event_date | date | 
-- | games_played | int | 
-- +--------------+---------+ 
-- (player_id, event_date) является первичным ключом (комбинацией столбцов со уникальными значениями) этой таблицы. 
-- Эта таблица отображает активность игроков в некоторых играх. 
-- Каждая строка — это запись игрока, который вошёл в систему и сыграл определённое количество игр (возможно, 0), а затем вышел в какой-то день, используя какое-то устройство. 

-- Задача 1: 
-- Найти общее количество уникальных игроков, которые были активны в 2016 году.
-- Ожидаемый вывод: Одно число — количество игроков.

SELECT COUNT(DISTINCT player_id) AS cnt
FROM activity
WHERE YEAR(event_date) = 2016

-- Задача 2: 
-- Для каждого игрока (player_id) определить, сколько всего игр (games_played) он сыграл за всё время.
-- Ожидаемый вывод: Таблица с колонками player_id и total_games_played, отсортированная по player_id.

SELECT
    player_id,
    SUM(games_played) AS total_games_played
FROM activity
GROUP BY player_id

-- Задание 3 (Средний уровень)
-- Цель: Определить дату первой активности каждого игрока и устройство (device_id), которое он использовал в этот день.
-- Ожидаемый вывод: Таблица с колонками player_id, first_event_date, first_device_id.

WITH first_date AS (
    SELECT
        player_id,
        MIN(event_date) AS fed
    FROM activity
    GROUP BY player_id
)
SELECT
    activity.player_id,
    first_date.fed,
    device_id AS first_device_id
FROM activity
JOIN first_date ON
    activity.player_id = first_date.player_id AND
    activity.event_date = first_date.fed

-- Задание 3 (Средний уровень)
-- Цель: Определить дату первой активности каждого игрока и устройство (device_id), которое он использовал в этот день.
-- Ожидаемый вывод: Таблица с колонками player_id, first_event_date, first_device_id.

SELECT
    player_id,
    fd AS first_event_date,
    fdi AS first_device_id
FROM (
    SELECT
        player_id,
        event_date AS fd,
        device_id AS fdi,
        ROW_NUMBER() OVER (PARTITION BY player_id ORDER BY event_date) AS rn
    FROM activity
) temp
WHERE rn = 1

-- Задание 3 (Средний уровень)
-- Цель: Определить дату первой активности каждого игрока и устройство (device_id), которое он использовал в этот день.
-- Ожидаемый вывод: Таблица с колонками player_id, first_event_date, first_device_id.

SELECT
    DISTINCT player_id,
    FIRST_VALUE(event_date) OVER (PARTITION BY player_id ORDER BY event_date) AS first_event_date,
    FIRST_VALUE(device_id) OVER (PARTITION BY player_id ORDER BY event_date) AS first_device_id
FROM activity

-- Задание 4 (Средний уровень)
-- Цель: Найти количество сессий (записей в таблице) и среднее количество сыгранных игр (games_played) для каждого игрока, но только для тех, у кого общее количество сессий больше 1.
-- Ожидаемый вывод: Таблица с колонками player_id, session_count, avg_games_per_session.

SELECT
    player_id,
    COUNT(*) AS session_count,
    AVG(games_played) AS avg_games_per_session
FROM activity
GROUP BY player_id
HAVING COUNT(*) > 1

-- Задание 5 (Средний/Сложный уровень)
-- Цель: Найти всех игроков, которые вернулись в игру на следующий день после своей первой активности (т.е. зашли и в первый, и во второй день). Под "следующим днем" подразумевается дата, следующая за first_event_date, независимо от того, были ли пропуски в дальнейшем.
-- Ожидаемый вывод: Таблица с колонкой player_id, содержащая только ID таких игроков.

SELECT