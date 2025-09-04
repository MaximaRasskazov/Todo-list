function initCalendar() {
    const calendarDays = document.getElementById('calendar-days');
    const currentMonthYear = document.getElementById('current-month-year');
    const prevMonthBtn = document.getElementById('prev-month');
    const nextMonthBtn = document.getElementById('next-month');
    
    let currentDate = new Date();
    
    function renderCalendar() {
        // Очищаем календарь
        calendarDays.innerHTML = '';
        
        // Получаем первый день месяца и последний день месяца
        const firstDay = new Date(currentDate.getFullYear(), currentDate.getMonth(), 1);
        const lastDay = new Date(currentDate.getFullYear(), currentDate.getMonth() + 1, 0);
        
        // Обновляем заголовок
        const months = ['Январь', 'Февраль', 'Март', 'Апрель', 'Май', 'Июнь', 
                        'Июль', 'Август', 'Сентябрь', 'Октябрь', 'Ноябрь', 'Декабрь'];
        currentMonthYear.textContent = `${months[currentDate.getMonth()]} ${currentDate.getFullYear()}`;
        
        // Добавляем пустые ячейки для дней перед первым днем месяца
        for (let i = 0; i < (firstDay.getDay() === 0 ? 6 : firstDay.getDay() - 1); i++) {
            const emptyDay = document.createElement('div');
            emptyDay.classList.add('empty');
            calendarDays.appendChild(emptyDay);
        }
        
        // Добавляем дни месяца
        for (let i = 1; i <= lastDay.getDate(); i++) {
            const day = document.createElement('div');
            day.textContent = i;
            
            // Проверяем, является ли день текущим
            const today = new Date();
            if (i === today.getDate() && 
                currentDate.getMonth() === today.getMonth() && 
                currentDate.getFullYear() === today.getFullYear()) {
                day.classList.add('today');
            }
            
            calendarDays.appendChild(day);
        }
    }
    
    // Обработчики для кнопок навигации
    prevMonthBtn.addEventListener('click', () => {
        currentDate.setMonth(currentDate.getMonth() - 1);
        renderCalendar();
    });
    
    nextMonthBtn.addEventListener('click', () => {
        currentDate.setMonth(currentDate.getMonth() + 1);
        renderCalendar();
    });
    
    // Первоначальная отрисовка календаря
    renderCalendar();
}

// Инициализация календаря после загрузки DOM
document.addEventListener('DOMContentLoaded', initCalendar);