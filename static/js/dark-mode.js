document.addEventListener("DOMContentLoaded", function () {
    toggleDarkMode(localStorage.getItem("darkMode") === "true");
});

function toggleDarkMode(isDarkMode) {
    document.body.classList.toggle("dark-mode", isDarkMode);
    var sunIcon = document.querySelector("#darkModeToggle .fa-sun");
    var moonIcon = document.querySelector("#darkModeToggle .fa-moon");

    sunIcon.style.display = isDarkMode ? "inline" : "none";
    moonIcon.style.display = isDarkMode ? "none" : "inline";
    sunIcon.style.color = moonIcon.style.color = isDarkMode ? "#fff" : "#000";
}

function myFunction() {
    var isDarkMode = !document.body.classList.contains("dark-mode");
    localStorage.setItem("darkMode", isDarkMode);
    toggleDarkMode(isDarkMode);
}

