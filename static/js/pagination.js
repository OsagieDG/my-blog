document.addEventListener("DOMContentLoaded", function() {
    const itemsPerPage = 1; // Number of items per page
    const items = document.querySelectorAll('.blog-list ul li');
    const numPages = Math.ceil(items.length / itemsPerPage); // Calculate the number of pages
    let currentPage = 1;

    function showPage(page) {
        const startIndex = (page - 1) * itemsPerPage;
        const endIndex = page * itemsPerPage;

        items.forEach((item, index) => {
            if (index >= startIndex && index < endIndex) {
                item.style.display = 'block';
            } else {
                item.style.display = 'none';
            }
        });
    }

    showPage(currentPage);

    function updatePagination() {
        const paginationContainer = document.querySelector('.pagination');
        paginationContainer.innerHTML = '';

        const prevButton = document.createElement('button');
        prevButton.textContent = 'Previous';
        prevButton.addEventListener('click', function() {
            if (currentPage > 1) {
                currentPage--;
                showPage(currentPage);
                updatePagination();
            }
        });
        paginationContainer.appendChild(prevButton);

        for (let i = 1; i <= numPages; i++) {
            const button = document.createElement('button');
            button.textContent = i;
            button.addEventListener('click', function() {
                currentPage = i;
                showPage(currentPage);
                updatePagination();
            });
            if (i === currentPage) {
                button.disabled = true;
            }
            paginationContainer.appendChild(button);
        }

        const nextButton = document.createElement('button');
        nextButton.textContent = 'Next';
        nextButton.addEventListener('click', function() {
            if (currentPage < numPages) {
                currentPage++;
                showPage(currentPage);
                updatePagination();
            }
        });
        paginationContainer.appendChild(nextButton);
    }

    updatePagination();
});

