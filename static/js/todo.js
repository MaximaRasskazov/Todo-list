// Функционал Todo List
function initTodo() {
    const todoInput = document.querySelector('.input-area input');
    const addButton = document.getElementById('add-btn');
    const todoList = document.querySelector('.todo-list');
    const filterButtons = document.querySelectorAll('.filter-btn');
    
    // Добавление задачи
    addButton.addEventListener('click', addTodo);
    todoInput.addEventListener('keypress', function(e) {
        if (e.key === 'Enter') addTodo();
    });
    
    function addTodo() {
        const text = todoInput.value.trim();
        if (text) {
            const li = document.createElement('li');
            li.innerHTML = `
                <label>
                    <input type="checkbox">
                    <span class="checkmark"></span>
                    <span class="todo-text">${text}</span>
                </label>
                <button class="delete-btn">×</button>
            `;
            
            // Обработчики для новой задачи
            li.querySelector('input[type="checkbox"]').addEventListener('change', toggleComplete);
            li.querySelector('.delete-btn').addEventListener('click', function() {
                li.remove();
            });
            
            todoList.appendChild(li);
            todoInput.value = '';
            
            // Применяем текущий фильтр
            applyFilter(document.querySelector('.filter-btn.active').dataset.filter);
        }
    }
    
    // Переключение статуса задачи
    function toggleComplete() {
        this.closest('li').classList.toggle('completed');
    }
    
    // Назначение обработчиков для существующих задач
    document.querySelectorAll('ul.todo-list li input[type="checkbox"]').forEach(checkbox => {
        checkbox.addEventListener('change', toggleComplete);
    });
    
    document.querySelectorAll('.delete-btn').forEach(button => {
        button.addEventListener('click', function() {
            this.closest('li').remove();
        });
    });
    
    // Фильтрация задач
    filterButtons.forEach(button => {
        button.addEventListener('click', function() {
            // Убираем активный класс у всех кнопок
            filterButtons.forEach(btn => btn.classList.remove('active'));
            // Добавляем активный класс текущей кнопке
            this.classList.add('active');
            
            const filter = this.dataset.filter;
            applyFilter(filter);
        });
    });
    
    function applyFilter(filter) {
        const items = todoList.querySelectorAll('li');
        
        items.forEach(item => {
            switch(filter) {
                case 'active':
                    item.style.display = item.classList.contains('completed') ? 'none' : 'flex';
                    break;
                case 'completed':
                    item.style.display = item.classList.contains('completed') ? 'flex' : 'none';
                    break;
                default:
                    item.style.display = 'flex';
            }
        });
    }
}

// Инициализация Todo после загрузки DOM
document.addEventListener('DOMContentLoaded', initTodo);