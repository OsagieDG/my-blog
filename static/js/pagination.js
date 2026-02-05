// @ts-check

document.addEventListener("DOMContentLoaded", function() {
  const itemsPerPage = 2;
  const items = /** @type {NodeListOf<HTMLElement>} */ (document.querySelectorAll('.blog-list ul li'));
  const numPages = Math.ceil(items.length / itemsPerPage);
  let currentPage = parseInt(localStorage.getItem('currentPage') || '1') || 1;

  /**
   * @param {number} page
   */
  function showPage(page) {
    items.forEach((item, index) => {
      item.style.display = index >= (page - 1) * itemsPerPage &&
        index < page * itemsPerPage ? 'block' : 'none';
    });
  }

  function updatePagination() {
    const paginationContainer = document.querySelector('.pagination');
    if (!paginationContainer) {
      console.error("Pagination container not found!");
      return;
    }
    paginationContainer.innerHTML = '';

    /**
     * @param {string} text
     * @param {number} page
     * @returns {HTMLButtonElement}
     */
    const createButton = (text, page) => {
      const button = document.createElement('button');
      button.textContent = text;
      button.disabled = page === currentPage;
      button.addEventListener('click', () => {
        currentPage = page;
        localStorage.setItem('currentPage', String(currentPage));
        showPage(currentPage);
        updatePagination();
      });
      return button;
    };

    paginationContainer.appendChild(createButton('Latest', 1));

    paginationContainer.appendChild(createButton('Back',
      currentPage > 1 ? currentPage - 1 : 1));

    const windowSize = 2;
    const windowStart = Math.max(1, Math.min(currentPage, numPages - windowSize + 1));
    const windowEnd = Math.min(windowStart + windowSize - 1, numPages);
    for (let i = windowStart; i <= windowEnd; i++) {
      paginationContainer.appendChild(createButton(String(i), i));
    }

    paginationContainer.appendChild(createButton('Next',
      currentPage < numPages ? currentPage + 1 : numPages));

    paginationContainer.appendChild(createButton('First', numPages));

  }

  showPage(currentPage);
  updatePagination();
});

