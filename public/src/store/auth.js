export const AuthStore = {
	state: {
		token: window.localStorage.getItem("token")
	},
	mutations: {
		login(state, token) {
			window.localStorage.setItem("token", token);
		},
		logout(state) {
			window.localStorage.removeItem("token");
		}
	}
};
