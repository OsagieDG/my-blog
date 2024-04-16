document.addEventListener("DOMContentLoaded", function () {
    var isDarkMode = localStorage.getItem("darkMode") === "true";
    var element = document.body;
    var sunIcon = document.querySelector("#darkModeToggle .fa-sun");
    var moonIcon = document.querySelector("#darkModeToggle .fa-moon");

    if (isDarkMode) {
        element.classList.add("dark-mode");
        sunIcon.style.display = "inline";
        moonIcon.style.display = "none";
        sunIcon.style.color = "#fff";
        moonIcon.style.color = "#fff";
    } else {
        sunIcon.style.display = "none";
        moonIcon.style.display = "inline";
        sunIcon.style.color = "#000";
        moonIcon.style.color = "#000";
    }
});

function myFunction() {
    var element = document.body;
    var sunIcon = document.querySelector("#darkModeToggle .fa-sun");
    var moonIcon = document.querySelector("#darkModeToggle .fa-moon");

    element.classList.toggle("dark-mode");
    var isDarkMode = element.classList.contains("dark-mode");

    localStorage.setItem("darkMode", isDarkMode);

    if (isDarkMode) {
        sunIcon.style.display = "inline";
        moonIcon.style.display = "none";
        sunIcon.style.color = "#fff";
        moonIcon.style.color = "#fff";
    } else {
        sunIcon.style.display = "none";
        moonIcon.style.display = "inline";
        sunIcon.style.color = "#000";
        moonIcon.style.color = "#000";
    }
}

