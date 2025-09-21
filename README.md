# To-Do List Application

Простое веб-приложение для управления задачами, состоящее из бэкенда на Go и фронтенда на HTML/CSS/JavaScript.

## 🚀 Функциональность

- ✅ Создание, просмотр, обновление и удаление задач
- ✅ Отметка задач как выполненных
- ❌ RESTful API с поддержкой CORS (In progress)
- ✅ Хранение данных в PostgreSQL
- ❌ Контейнеризация с Docker (In progress)

## 🛠 Технологии

### Бэкенд
- **Go** (Golang) - язык программирования
- **PostgreSQL** - система управления базами данных
- **pgx** - драйвер PostgreSQL для Go
- **Docker** - контейнеризация приложения

### Фронтенд
- **HTML5** - структура веб-страницы
- **CSS3** - стилизация
- **JavaScript** - взаимодействие с пользователем

## 📦 Установка и запуск

### Предварительные требования
- Установленный Go (версия 1.21 или выше)
- --Установленный Docker и Docker Compose--
- Установленный PostgreSQL (для локальной разработки)

### 1. Клонирование репозитория
```bash
git clone <URL вашего репозитория>
cd to-do-list
```

### 2. Настройка базы данных
```bash
# Создание базы данных
createdb todo_db

# Подключение к БД и создание таблицы
psql todo_db -c "CREATE TABLE todos (
    id SERIAL PRIMARY KEY,
    title TEXT NOT NULL,
    completed BOOLEAN DEFAULT FALSE,
    created_at TIMESTAMP DEFAULT NOW()
);"
```

### 3. Настройка окружения
Создайте файл `.env` в корневой директории проекта:
```env
DATABASE_URL=postgresql://username:password@localhost:5432/todo_db?sslmode=disable
PORT=3000
```

### 4. Запуск приложения
```bash
# Установка зависимостей
go mod download

# Запуск сервера
go run .
```

### 5. Запуск с помощью Docker
```bash
# Сборка и запуск контейнеров
docker-compose up --build

# Остановка контейнеров
docker-compose down
```

Приложение будет доступно по адресу: http://localhost:3000

## 📡 API Endpoints

### Получить все задачи
```
GET /api/todos
```
Пример ответа:
```json
[
  {
    "id": 1,
    "title": "Изучить Go",
    "completed": false,
    "createdAt": "2023-05-15T10:30:00Z"
  }
]
```

### Создать новую задачу
```
POST /api/todos
Content-Type: application/json

{
  "title": "Новая задача"
}
```

### Обновить задачу
```
PUT /api/todos/:id
Content-Type: application/json

{
  "title": "Обновленная задача",
  "completed": true
}
```

### Удалить задачу
```
DELETE /api/todos/:id
```

## 🗂 Структура проекта

```
to-do-list/
├── frontend/           # Фронтенд приложения
│   ├── index.html
│   ├── styles.css
│   └── script.js
├── cmd/
│   └── server/         # Основной исполняемый пакет
│       └── main.go
├── internal/           # Внутренние пакеты приложения
│   ├── handlers/       # HTTP обработчики
│   ├── models/         # Модели данных
│   └── database/       # Работа с базой данных
├── docker-compose.yml  # Конфигурация Docker Compose
├── Dockerfile         # Конфигурация Docker
├── .env.example       # Пример файла окружения
├── go.mod            # Модули Go
└── README.md         # Документация
```

## 🔧 Разработка

### Добавление новых зависимостей
```bash
go get <package-name>
```

### Форматирование кода
```bash
go fmt ./...
```

### Тестирование
```bash
go test ./...
```

## 🚀 Деплой

### Сборка для продакшена
```bash
go build -o todo-app ./cmd/server
```

### Запуск в продакшене
```bash
./todo-app
```

## 📝 Планы по развитию

- [ ] Добавление аутентификации пользователей
- [ ] Категории и теги для задач
- [ ] Написание unit-тестов
- [ ] Добавление пагинации для API
- [ ] Реализация полнотекстового поиска
- [ ] Добавление дедлайнов для задач
- [ ] Уведомления о предстоящих задачах

## 🤝 Участие в разработке

1. Форкните репозиторий
2. Создайте ветку для вашей функции (`git checkout -b feature/amazing-feature`)
3. Закоммитьте изменения (`git commit -m 'Add some amazing feature'`)
4. Запушьте в ветку (`git push origin feature/amazing-feature`)
5. Откройте Pull Request

## 📄 Лицензия

Этот проект распространяется под лицензией MIT. Подробнее см. в файле LICENSE.

## 📞 Контакты

Ваше имя - [your.email@example.com](mailto:your.email@example.com)

Ссылка на проект: [https://github.com/your-username/to-do-list](https://github.com/your-username/to-do-list)

---

⭐ Не забудьте поставить звезду репозиторию, если проект вам понравился!
