const links = document.querySelectorAll("a");
links.forEach(link => {
    link.innerHTML += " (link)";
});