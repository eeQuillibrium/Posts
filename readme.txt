Go application for post, comments and user storaging with GraphQL, postgresql

Если нужно запустить с INMEMORY хранилищем, то изменить в docker-compose.yml STORAGE_MODE=INMEMORY
Если нужно запустить с POSTGRES хранилищем, то изменить в docker-compose.yml STORAGE_MODE=POSTGRES

Для запуска:
  Через compose:
    1. изменить host на postgres_db в config/config.yaml
    2. docker-compose up --build
  Через cmd/main.go:
    1. изменить host на localhost в config/config.yaml
    2. cmd/main.go из папки posts
GraphQL playground на localhost:8080

Mutations:  
  1. createComments: создает комментарии в бд и записывает в канал уведомление, 
  которое приходит в Subscriptons notification(postId: Int!), из которого читает клиент, на чтение из канала дается 5 секунд
  2. createPost: создает пост  
  3. createUser: создает юзера  

Queries:
  1. posts: возвращает посты без подгрузки комментариев с offset-limit пагинацией  
  2. post: подгружает(кеширует) в in-memory комментарии по postID и возвращает limit комментариев на каждом уровне  
  3. paginationComment: возвращает limit комментариев верхнего уровня с любой вложенностью  

Subscriptions:
  notification: читает по postId: Int! из канала все приходящие  


!для nested комментариев сделал запись в gqlgen.yml, чтобы он сгерерировал отдельный резолвер и читал из бд комменты 
с parent_id comment-а, который запрашивается
