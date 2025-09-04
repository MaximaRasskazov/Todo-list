// Обновление времени в шапке
function updateTime() {
    const now = new Date();
    const hours = now.getHours().toString().padStart(2, '0');
    const minutes = now.getMinutes().toString().padStart(2, '0');
    const timeString = `${hours}:${minutes}`;
    document.getElementById('current-time').textContent = timeString;
}

// Инициализация времени и установка интервала обновления
updateTime();
setInterval(updateTime, 1000);