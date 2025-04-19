# Posts and Comments Service

Система для добавления и чтения постов и комментариев с использованием GraphQL, созданная с применением современных технологий для обеспечения масштабируемости и удобства.

## Описание проекта

Реализована система управления постами и комментариями, аналогичная комментариям на популярных платформах, таких как Хабр или Reddit.

### Функциональные возможности

#### Посты:
- Просмотр списка постов с поддержкой пагинации.
- Просмотр деталей поста, включая связанные комментарии.
- Возможность для автора поста запретить комментарии к своему посту.

#### Комментарии:
- Иерархическая структура с неограниченной вложенностью.
- Ограничение длины комментария (до 2000 символов).
- Поддержка пагинации для получения комментариев.

#### Subscriptions (GraphQL):
- Асинхронная доставка новых комментариев к посту через подписки.

---

## Технологии

### Backend:
- **Язык:** Go
- **GraphQL:** Библиотека [gqlgen](https://gqlgen.com/) для реализации GraphQL API.
- **Хранилище данных:**
    - PostgreSQL для постоянного хранения данных.
    - In-memory хранилище для быстрого тестирования и разработки.
- **Docker:** Контейнеризация сервиса для удобного развёртывания.
- **PostgreSQL Migrations:** Управление схемой базы данных через миграции.
- **Logging:** Логирование ошибок и событий с помощью кастомного логгера.
- **Подписки:** Реализация через `CommentObserver` для управления слушателями.

---

## Запуск

### Локальный запуск
2. Запустите базу данных:
   ```bash
   make docker.run.db
   ```
3. Примените миграции:
   ```bash
   make docker.run.migrate
   ```
   *Или используйте локальную утилиту `golang-migrate` для применения миграций.*
4. Запустите сервер:
   ```bash
   make local.run
   ```

### Запуск в Docker
1. Запустите все контейнеры:
   ```bash
   make docker.run
   ```
   *Миграции применяются автоматически.*

### Запуск тестов
Для запуска тестов используйте:
```bash
make tests.run
```

### Остановка контейнеров
Остановить и удалить контейнеры можно командой:
```bash
make docker.down
```

---

## API

Описание API реализовано в `.graphqls` файлах в директории `api/graphql`.
Для тестирования можно использовать [Postman](https://www.postman.com/) или аналогичные инструменты. Примеры запросов:

### Примеры запросов

#### Создание поста:
```graphql
mutation CreatePost {
    CreatePost(
        post: {
            name: "Interesting Post"
            content: "This is the content."
            author: "Author Name"
            commentsAllowed: true
        }
    ) {
        id
        createdAt
        name
        author
        content
    }
}
```

#### Получение списка постов:
```graphql
query GetAllPosts {
    GetAllPosts(page: 1, pageSize: 5) {
        id
        createdAt
        name
        author
        content
    }
}
```

#### Получение деталей поста с комментариями:
```graphql
query GetPostById {
    GetPostById(id: 1) {
        id
        createdAt
        name
        author
        content
        commentsAllowed
        comments(page: 1, pageSize: 2) {
            id
            createdAt
            author
            content
            post
            replies {
                id
                createdAt
                author
                content
                post
                replyTo
            }
        }
    }
}
```

#### Создание комментария:
```graphql
mutation CreateComment {
    CreateComment(input: { author: "Commenter", content: "Great post!", post: "1" }) {
        id
        createdAt
        author
        content
        post
        replyTo
    }
}
```

#### Подписка на новые комментарии:
```graphql
subscription CommentsSubscription {
    CommentsSubscription(postId: "1") {
        id
        createdAt
        author
        content
        post
        replyTo
    }
}
```

#### Тестирование через GraphQL Playground

##### 1. Подписка на комментарии
Откройте вкладку и выполните следующий запрос:
```graphql
subscription OnCommentAdded {
  commentAdded(postId: "1") {
    id
    author
    text
    createdAt
  }
}
```

##### 2. Создание комментария
Откройте вторую вкладку и выполните запрос:
```graphql
mutation CreateTestComment {
  createComment(input: {
    author: "Тестовый пользователь",
    text: "Проверка подписки 123",
    postId: "1"
  }) {
    id
  }
}
```

##### 3. Ожидаемый результат
Во вкладке с подпиской появится:
```json
{
  "data": {
    "commentAdded": {
      "id": "3",
      "author": "Тестовый пользователь",
      "text": "Проверка подписки 123",
      "createdAt": "2025-04-20T12:34:56Z"
    }
  }
}
```

---

### Полный тестовый сценарий

#### 1. Создание поста
```graphql
mutation CreateFirstPost {
    createPost(input: {
        title: "Тестовый пост 1",
        content: "Содержание первого тестового поста",
        author: "Администратор",
        allowComments: true
    }) {
        id
    }
}
```

#### 2. Добавление комментария
```graphql
mutation AddFirstComment {
  createComment(input: {
    author: "Пользователь 1",
    text: "Первый комментарий!",
    postId: "1"
  }) {
    id
  }
}
```

#### 3. Получение поста по ID с комментариями
```graphql
query CheckPost {
  post(id: "1") {
    id
    title
    comments(first: 1) {
      id
      author
      text
    }
  }
}
```

#### 4. Получение списка всех постов
В каждом списке по 10 постов
```graphql
query CheckAllPosts {
  posts(first: 1) { # указывает какая страница
    edges {
      node {
        id
        title
      }
    }
  }
}
```

#### 5. Добавление ответа к комментарию
```graphql
mutation ReplyToComment {
  createComment(input: {
    author: "Пользователь 2",
    text: "Это ответ на первый комментарий",
    postId: "1",  # Должен совпадать с постом
    parentId: "1"  # ID родительского комментария
  }) {
    id
  }
}
```

#### 6. Получение комментариев поста с вложенными ответами
```graphql
query GetPostWithComments {
    post(id: "1") {
        id
        title
        comments(first: 1) {
            id
            author
            text
            replies(first: 1) {
                id
                author
                text
            }
        }
    }
}
```

---


## Выбор хранилища данных
Поддерживается два варианта хранения данных:
1. **PostgreSQL:** Используется по умолчанию для постоянного хранения.
2. **In-memory:** Опционально, включается через переменную окружения:
    - В `.env` файле: `USE_IN_MEMORY=true`
    - В `docker-compose.yml` для Docker.

---

## Заключение
Система построена с учётом современных практик разработки. Она обеспечивает масштабируемость, высокую производительность и удобство работы с GraphQL API. Возможность выбора хранилища данных и поддержка подписок делают её гибкой и подходящей для различных сценариев использования.

