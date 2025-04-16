# testTask
#### Запрос для пагинации:
GET http://localhost:8080/posts/list?limit=100&offset=0
#### Зпрос для получения всех постов:
GET http://localhost:8080/posts/list
#### Запрос для регистрации:
Post http://localhost:8080/register
{
    "username": "vasya",
    "password": "1234"
}
#### Запрос для входа: Post http://localhost:8080/login
{
    "username": "vasya",
    "password": "1234"
}
#### Запрос для создания поста: Post http://localhost:8080/posts
{
  "title": "Пример поста",
  "content": "Это тестовый пост.",
  "comments_enabled": true
}
Поле "comments_enabled" - отвечает за разрешение пистать комментарии под постом
#### Запрос на комментарий: POST http://localhost:8080/comment?post_id=1
{
    "content": "Это мой комментарий!",
    "parent_id": null
}
####  Ответ на комментарий: POST http://localhost:8080/comment?post_id=1
{
    "content": "Это ответ на комментарий!",
    "parent_id": 1
}
#### UPD Для создания поста, написания комметрия требуется jwt токен который генериуется при входе, его необходимо вставлять в ручную. 
#### Схема бд:
──────────────
│   users    │
──────────────
│ id (PK)    │
│ username   │
│ password   │
│ created_at │
──────────────
      │
      |
──────────────
│   posts    │
──────────────
│ id (PK)    │
│ user_id ─────>(users.id)
│ title      │
│ content    │
│ comments_enabled 
│ created_at │
──────────────
      │
      |
──────────────────────────────
│         comments           │
──────────────────────────────
│ id (PK)                    │
│ user_id ─────────────────────>(users.id)
│ post_id ─────────────────────>(posts.id)
│ parent_id ───────────────────>(comments.id)
│ content                    │
│ created_at                 │
──────────────────────────────
