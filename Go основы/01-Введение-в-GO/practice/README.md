# Задания

1. Измените программу `echo` так, чтобы она выводила так-же имя выполняемой программы
2. Измените программу `echo` так, чтобы она выводила индекс и значение аргумента по одному аргументу в строке
3. Измените программу `dub2` так, чтобы она выводила имена всех файлов, в которых найдены повторяющиеся строки.
4. Измените программу `gif` таким образом, чтобы изображение осцилографо было зеленого цвета, на черном фоне. Чтобы создать цвет `#RRGGBB` надо использовать следующий вызов: `color.RGBA(0xRR,0xGG,0xBB,0xff)`, каждая из которых является 16-ричной цифрой индекса компоненты цвета.
5. Измените программу `gif` так, чтобы она генерировала изображения различных цветов, добавляя в палитру `palette` больше значений, а затем выводя их путем изменения третьего аргумента функции `SetColorIndex`
6. Модифицируйте программу `get_request_body`: вызов функции `io.Copy(dst,src)` выполняет чтение `src` и запись `dst`. Используйте ее вместо `ioutill.ReadAll` для копирование тела ответа в поток `os.Stdout` без необходимости выделения большого буфера для хранения ответа.
7. Измените программу `get_request_body` так, чтобы каждому аргументу URL автоматически добавлялся префикс `http://` в случае его отсутствия в запросе. Можно использовать функцию `strings.HasPrefix`
8. Измените программу `get_request_body` так, чтобы она выводила код состояния ответа HTTP, содержащийся в `resp.Status`
9. Найлите сайт, который содрежит большое количество данных. Исследуйте работу кеширования, путем повторного запуска программы `get_many_requests`. Как изменяется время выполнения запросов. Получаете ли вы одно и то же содержимое. Измените программу так, чтобы вести запись в отдельные файлы.
10. Попробуйте обратиться к сайту который не отвечает. Как ведет себя программа?
11. Измените сервер с фигурами Лиссанжу так, чтобы значения параметров считывались из URL. Например, URL вида `http.localhost:8000/cycles=20`. Используйте функцию `strconv.Atioi` для преобразования строкого параметра в целое число. Посмотреть документацию можно с помошью команды `go doc strconv.Atoi`
