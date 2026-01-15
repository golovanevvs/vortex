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

WITH
    first_activity AS (
        SELECT
            player_id,
            MIN(event_date) AS first_date
        FROM activity
        GROUP BY player_id
    )
SELECT player_id
FROM
    activity a JOIN
    first_activity fa ON
    a.player_id = fa.player_id
WHERE
    a.event_date - fa.first_date = 1

-- Задание 5 (Средний/Сложный уровень)
-- Цель: Найти всех игроков, которые вернулись в игру на следующий день после своей первой активности (т.е. зашли и в первый, и во второй день). Под "следующим днем" подразумевается дата, следующая за first_event_date, независимо от того, были ли пропуски в дальнейшем.
-- Ожидаемый вывод: Таблица с колонкой player_id, содержащая только ID таких игроков.

SELECT player_id
FROM (
    SELECT
        player_id,
        event_date,
        LEAD(event_date) OVER (PARTITION BY player_id ORDER BY event_date) AS lead_date,
        ROW_NUMBER() OVER (PARTITION BY player_id ORDER BY event_date) AS rn
    FROM activity
) t
WHERE
    rn = 1 AND
    lead_date - event_date = 1

-- Задание 6 (Сложный уровень)
-- Цель: Рассчитать "долю удержания" игроков на 7-й день. Доля удержания — это процент игроков, которые были активны (имели хотя бы одну сессию) ровно через 7 дней после своей первой активности (т.е. на first_event_date + 7).
-- Уточнение: Считайте, что если у игрока есть запись с event_date, в точности равной дате его первого входа + 7 дней, то он "удержан".
-- Ожидаемый вывод: Одно число retention_rate, округленное до 2 знаков после запятой (в процентах или долях, уточните в формулировке задачи).

WITH t_diff AS (
    SELECT
        player_id,
        event_date,
        MIN(event_date) OVER (PARTITION BY player_id) AS first_date,
        event_date - first_date AS diff
    FROM activity
    )
SELECT ROUND(COUNT(*)*100.0/(SELECT COUNT(DISTINCT player_id) FROM activity), 2) AS retention_rate
FROM t_diff
WHERE diff = 7

-- Задание 7 (Сложный уровень, аналитика)
-- Цель: Для каждого календарного дня (даты из колонки event_date по всем игрокам) рассчитать:
-- Количество активных игроков в этот день (DAU - Daily Active Users).
-- Количество новых игроков в этот день (тех, для кого эта дата является first_event_date).
-- Ожидаемый вывод: Таблица с колонками event_date, dau, new_players, отсортированная по event_date.
-- Эти задания охватывают ключевые концепции SQL: агрегацию (COUNT, SUM, AVG), группировку (GROUP BY), фильтрацию (HAVING, подзапросы в WHERE), работу с датами, использование оконных функций (ROW_NUMBER(), FIRST_VALUE()) и сложную логику для аналитических расчетов.

WITH temp AS (
    SELECT
        player_id,
        event_date,
        MIN(event_date) OVER (PARTITION BY player_id) AS first_date,
        event_date - first_date AS diff
    FROM activity
)
SELECT
    event_date,
    COUNT(DISTINCT player_id) AS dau,
    COUNT(CASE WHEN diff = 0 THEN 1 END) AS new_players
FROM temp
GROUP BY event_date
ORDER BY event_date

-- Задание 7 (Сложный уровень, аналитика)
-- Цель: Для каждого календарного дня (даты из колонки event_date по всем игрокам) рассчитать:
-- Количество активных игроков в этот день (DAU - Daily Active Users).
-- Количество новых игроков в этот день (тех, для кого эта дата является first_event_date).
-- Ожидаемый вывод: Таблица с колонками event_date, dau, new_players, отсортированная по event_date.
-- Эти задания охватывают ключевые концепции SQL: агрегацию (COUNT, SUM, AVG), группировку (GROUP BY), фильтрацию (HAVING, подзапросы в WHERE), работу с датами, использование оконных функций (ROW_NUMBER(), FIRST_VALUE()) и сложную логику для аналитических расчетов.

SELECT
    event_date,
    COUNT(DISTINCT player_id) AS dau,
    SUM(CASE WHEN event_date = MIN(event_date) OVER (PARTITION BY player_id) THEN 1 ELSE 0 END) AS new_players
FROM activity
GROUP BY event_date
ORDER BY event_date

-- Задание 8 (Сложный уровень, оконные функции)
-- Цель: Для каждой записи активности (player_id, event_date) найти "максимальное накопленное" количество игр. Это означает: для каждой даты игрока нужно знать рекордное (максимальное) общее количество игр, которое он когда-либо набрал ко всем предыдущим датам включительно.
-- Пример для понимания:
-- 2016-03-01, сыграно 5 → накопленная сумма = 5, максимум = 5
-- 2016-03-02, сыграно 0 → накопленная сумма = 5, максимум = 5
-- 2016-03-03, сыграно 3 → накопленная сумма = 8, максимум = 8
-- 2016-03-04, сыграно 1 → накопленная сумма = 9, максимум = 8 (рекорд всё еще 8)
-- Ожидаемый вывод: Таблица с исходными колонками и новой колонкой running_max_games.

SELECT
    player_id,
    device_id,
    event_date,
    games_played,
    MAX(running_sum) OVER (PARTITION BY player_id ORDER BY event_date) AS running_max_games
FROM
    (
        SELECT
            *,
            SUM(games_played) OVER (PARTITION BY player_id ORDER BY event_date)AS running_sum
            FROM activity
    ) t

-- Задание 9 (Сложный уровень, анализ сессий)
-- Цель: Определить "время до следующей активности" для каждого игрока. Для каждой даты входа игрока нужно найти, через сколько дней произошел его следующий вход. Для последней даты активности игрока значение должно быть NULL.
-- Уточнение: Если игрок заходил несколько раз в один день (по условию задачи не может, т.к. (player_id, event_date) — PK), но в реальности такое возможно, считайте, что даты уникальны.
-- Ожидаемый вывод: Таблица с колонками player_id, event_date, days_until_next_activity.

SELECT
    player_id,
    event_date,
    LEAD(event_date) OVER (PARTITION BY player_id ORDER BY event_date) - event_date AS days_until_next_activity
FROM activity
