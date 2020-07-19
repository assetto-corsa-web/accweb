export const AuthStore = {
	state: {
		token: window.localStorage.getItem("token"),
		admin: window.localStorage.getItem("admin") === "true",
		mod: window.localStorage.getItem("mod") === "true",
		read_only: window.localStorage.getItem("read_only") === "true",
	},
	mutations: {
		login(state, data) {
			state.token = data.token;
			state.admin = data.admin;
			state.mod = data.mod;
			state.read_only = data.read_only;
			window.localStorage.setItem("token", data.token);
			window.localStorage.setItem("admin", data.admin);
			window.localStorage.setItem("mod", data.mod);
			window.localStorage.setItem("read_only", data.read_only);
		},
		logout(state) {
			window.localStorage.removeItem("token");
			window.localStorage.removeItem("admin");
			window.localStorage.removeItem("mod");
			window.localStorage.removeItem("read_only");
		}
	}
};
