let loginBtn = document.querySelector('#btn');
if (loginBtn) {
    loginBtn.onclick = function (event) {
        let login = document.querySelector('#login').value;
        let password = document.querySelector('#password').value;

        if (login && password) {
            const loginData = {
                Login: document.querySelector('#login').value,
                Password: document.querySelector('#login').value,
            }
        } else {
            let loginAlert = document.querySelector('#login').closest('label').querySelector('p');
            let passwordAlert = document.querySelector('#password').closest('label').querySelector('p');
            if (!login) {
                loginAlert.textContent = "Необходимо заполнить поле «Логин»";
            } else {
                loginAlert.textContent = "";
            }
            if (!password) {
                passwordAlert.textContent = "Необходимо заполнить поле «Пароль»";
            } else {
                passwordAlert.textContent = "";
            }
        }
    }
}