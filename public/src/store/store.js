import Vuex from "vuex";
import {AuthStore} from "./auth.js";
import {ErrorStore} from "./error";

export default function NewStore() {
	return new Vuex.Store({
		modules: {
			auth: AuthStore,
			error: ErrorStore
		}
	});
}
