import Vuex from "vuex";
import {AuthStore} from "./auth.js";

export default function NewStore() {
	return new Vuex.Store({
		modules: {
			auth: AuthStore
		}
	});
}
