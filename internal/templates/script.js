document.addEventListener("DOMContentLoaded", function () {
    // Загрузка данных с сервера
    fetch('http://localhost:8081/v1/day-schedule')
        .then(response => response.json())
        .then(data => {
            // Заполнение таблицы
            const tableBody = document.querySelector('#scheduleTable tbody');
            data.task.forEach(task => {
                const row = tableBody.insertRow();
                const deadlineCell = row.insertCell(0);
                const descriptionCell = row.insertCell(1);

                deadlineCell.textContent = task.deadline;
                descriptionCell.textContent = task.description;
            });
        })
        .catch(error => console.error('Error fetching data:', error));
});
