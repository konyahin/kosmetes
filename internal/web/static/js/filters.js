function addTextToSearch(element) {
    const searchInput = document.getElementById('search-input');
    let textToAdd = element.textContent || element.innerText;

    if (element.classList.contains('project')) {
	textToAdd = 'project:' + textToAdd;
    }

    if (searchInput.value.length > 0) {
	searchInput.value += ' ' + textToAdd;
    } else {
	searchInput.value += textToAdd;
    }

    searchInput.dispatchEvent(new Event('input', { bubbles: true }));
}
