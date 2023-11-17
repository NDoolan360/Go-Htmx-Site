// Update copyright year
const copyright = document.getElementById('copyright')!;
copyright.innerHTML = copyright?.innerHTML.replace('{current year}', new Date().getFullYear().toString());
