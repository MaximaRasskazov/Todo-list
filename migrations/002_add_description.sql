-- Добавляем поле description к таблице todos
ALTER TABLE todos ADD COLUMN description TEXT;

-- Создаем индекс для поиска по описанию
CREATE INDEX IF NOT EXISTS idx_todos_description ON todos(description);