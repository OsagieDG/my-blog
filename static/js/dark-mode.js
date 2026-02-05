// @ts-check

document.addEventListener("DOMContentLoaded", function() {
  toggleDarkMode(localStorage.getItem("darkMode") === "true");
});

/**
 * @param {boolean} isDarkMode
 */
function toggleDarkMode(isDarkMode) {
  document.body.classList.toggle("dark-mode", isDarkMode);
  var sunIcon = /** @type {HTMLElement | null} */ (document.querySelector("#darkModeToggle .fa-sun"));
  var moonIcon = /** @type {HTMLElement | null} */ (document.querySelector("#darkModeToggle .fa-moon"));

  if (sunIcon && moonIcon) {
    sunIcon.style.display = isDarkMode ? "inline" : "none";
    moonIcon.style.display = isDarkMode ? "none" : "inline";
    sunIcon.style.color = moonIcon.style.color = isDarkMode ? "#fff" : "#000";
  }
}

function myFunction() {
  var isDarkMode = !document.body.classList.contains("dark-mode");
  localStorage.setItem("darkMode", String(isDarkMode));
  toggleDarkMode(isDarkMode);
}

