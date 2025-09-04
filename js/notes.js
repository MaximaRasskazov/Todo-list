function initNotes() {
    const notesList = document.getElementById('notes-list');
    const noteEditor = document.getElementById('note-editor');
    const addNoteBtn = document.getElementById('add-note-btn');
    const saveNoteBtn = document.getElementById('save-note');
    const cancelNoteBtn = document.getElementById('cancel-note');
    const deleteNoteBtn = document.getElementById('delete-note');
    const noteTitle = document.getElementById('note-title');
    const noteContent = document.getElementById('note-content');
    
    let notes = JSON.parse(localStorage.getItem('notes')) || [];
    let currentNoteId = null;
    
    // Отображение списка заметок
    function renderNotes() {
        notesList.innerHTML = '';
        
        notes.forEach((note, index) => {
            const noteElement = document.createElement('div');
            noteElement.classList.add('note-item');
            noteElement.innerHTML = `
                <h3>${note.title || 'Без названия'}</h3>
                <p>${note.content.substring(0, 50)}${note.content.length > 50 ? '...' : ''}</p>
            `;
            
            noteElement.addEventListener('click', () => {
                openNote(index);
            });
            
            notesList.appendChild(noteElement);
        });
    }
    
    // Открытие заметки для редактирования
    function openNote(id) {
        currentNoteId = id;
        const note = notes[id];
        
        noteTitle.value = note.title || '';
        noteContent.value = note.content || '';
        
        noteEditor.style.display = 'block';
        notesList.style.display = 'none';
        
        if (id === null) {
            deleteNoteBtn.style.display = 'none';
        } else {
            deleteNoteBtn.style.display = 'block';
        }
    }
    
    // Закрытие редактора заметок
    function closeEditor() {
        noteEditor.style.display = 'none';
        notesList.style.display = 'block';
        currentNoteId = null;
        noteTitle.value = '';
        noteContent.value = '';
    }
    
    // Сохранение заметки
    function saveNote() {
        const title = noteTitle.value.trim();
        const content = noteContent.value.trim();
        
        if (content) {
            if (currentNoteId !== null) {
                // Редактирование существующей заметки
                notes[currentNoteId] = { title, content };
            } else {
                // Создание новой заметки
                notes.push({ title, content });
            }
            
            localStorage.setItem('notes', JSON.stringify(notes));
            renderNotes();
            closeEditor();
        }
    }
    
    // Удаление заметки
    function deleteNote() {
        if (currentNoteId !== null) {
            notes.splice(currentNoteId, 1);
            localStorage.setItem('notes', JSON.stringify(notes));
            renderNotes();
            closeEditor();
        }
    }
    
    // Обработчики событий
    addNoteBtn.addEventListener('click', () => {
        currentNoteId = null;
        openNote(null);
    });
    
    saveNoteBtn.addEventListener('click', saveNote);
    cancelNoteBtn.addEventListener('click', closeEditor);
    deleteNoteBtn.addEventListener('click', deleteNote);
    
    // Первоначальная отрисовка заметок
    renderNotes();
}

// Инициализация заметок после загрузки DOM
document.addEventListener('DOMContentLoaded', initNotes);