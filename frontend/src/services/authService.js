class AuthService {
    constructor(api) {
        this.api = api;
    }

    login (payload) {
        return this.api.post('auth/login', payload)
            .then((response) => {
                if (response.data.token) {
                    localStorage.setItem('user', JSON.stringify(response.data));
                    console.log('User logged in!', localStorage.getItem('user'));
                }
            })
            .catch((e) => {
                alert(e.response.data.error);
            });
    }

    logout() {
        localStorage.removeItem('user');
        console.log('User logged out!', localStorage.getItem('user'));
    }
}

export default AuthService;