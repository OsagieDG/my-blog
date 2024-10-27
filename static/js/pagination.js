document.addEventListener("DOMContentLoaded", function() {
    const itemsPerPage = 1;
    const items = document.querySelectorAll('.blog-list ul li');
    const numPages = Math.ceil(items.length / itemsPerPage);
    let currentPage = parseInt(localStorage.getItem('currentPage')) || 1;

    function showPage(page) {
        items.forEach((item, index) => item.style.display = index >= (page - 1) * itemsPerPage && index < page * itemsPerPage ? 'block' : 'none');
    }

    function updatePagination() {
        const paginationContainer = document.querySelector('.pagination');
        paginationContainer.innerHTML = '';

        const createButton = (text, page) => {
            const button = document.createElement('button');
            button.textContent = text;
            button.disabled = page === currentPage;
            button.addEventListener('click', () => {
                currentPage = page;
                localStorage.setItem('currentPage', currentPage);
                showPage(currentPage);
                updatePagination();
            });
            return button;
        };

        paginationContainer.appendChild(createButton('Previous', currentPage > 1 ? currentPage - 1 : 1));
        for (let i = 1; i <= numPages; i++) paginationContainer.appendChild(createButton(i, i));
        paginationContainer.appendChild(createButton('Next', currentPage < numPages ? currentPage + 1 : numPages));
    }

    showPage(currentPage);
    updatePagination();
});
