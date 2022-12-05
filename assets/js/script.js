let menuBtn = document.querySelector('.menuBtn');
let navigation = document.querySelector('nav');
if (menuBtn) {
    menuBtn.onclick = function () {
        navigation.classList.toggle('active');
    }
}