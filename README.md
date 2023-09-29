# notes-service

## Сборка и запуск

make build && make run

## Запуск тестов

make test

## Создание документации

make swag

## Формат JSON файла

{

  "header": "Тестовое задание",

  "description": "проверить",

  "date": "2023-09-27 12:00:00",

  "status": "uncompleted"

}

## Входные данные

1. create - JSON файл с данными.
2. get - id задачи.
3. update - id задачи.
4. delete - id задачи.
5. list - status задачи, limit и offset, sort - по какому столбцу будут сортироваться данные.

## Выходные данные

1. create - id созданной задачиб status code о выполнение запроса.
2. get - JSON файл с данными по задаче, status code о выполнение запроса.
3. update - status code о выполнение запроса, status code о выполнение запроса.
4. delete- status code о выполнение запроса, status code о выполнение запроса.
5. list - список задач, которые удовлетворяют параметрам запроса, status code о выполнение запроса.
