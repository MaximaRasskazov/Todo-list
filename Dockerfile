# Используем официальный образ Go для сборки
FROM golang:1.24-alpine AS builder

# Устанавливаем рабочую директорию
WORKDIR /app

# Копируем файлы модулей и загружаем зависимости
COPY go.mod go.sum ./
RUN go mod download

# Копируем весь исходный код
COPY . .

# Собираем приложение
RUN go build -o todo-app ./cmd/server

# Финальный образ для запуска
FROM alpine:latest

# Устанавливаем необходимые пакеты
RUN apk --no-cache add ca-certificates

# Создаем пользователя для безопасности
RUN addgroup -S appgroup && adduser -S appuser -G appgroup
USER appuser

# Устанавливаем рабочую директорию
WORKDIR /root/

# Копируем собранное приложение из стадии builder
COPY --from=builder /app/todo-app .
COPY --from=builder /app/migrations ./migrations

# Открываем порт
EXPOSE 3000

# Команда для запуска приложения
CMD ["./todo-app"]