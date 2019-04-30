import Vue from "vue";
import VueRouter from "vue-router";
import Vuex from "vuex";
import VueI18n from "vue-i18n";
import axios from "axios";

import "../static/main.scss";
import NewStore from "./store/store.js";
import * as pages from "./pages";

Vue.use(VueRouter);
Vue.use(Vuex);
Vue.use(VueI18n);
Vue.config.productionTip = false;
Vue.config.devtools = false;

// reponse error handler
axios.interceptors.response.use(undefined, err => {
	if(err.response.data.error){
		console.log(err.response.data.error);
	}
	
	return Promise.reject(err);
});

// token interceptor for every request
axios.interceptors.request.use((config) => {
	const token = window.localStorage.getItem("token");

	if(token){
		config.headers.Authorization = `Bearer ${token}`;
	}

	return config;
}, (err) => {
	return Promise.reject(err);
});

// router
const routes = [
	{path: "/", component: pages.Overview, meta: {protected: true}},
	{path: "/login", component: pages.Login},
	{path: "/server", component: pages.Server, meta: {protected: true}},
	{path: "/logs", component: pages.Logs, meta: {protected: true}},
	{path: "*", component: pages.Error404}
];

let router = new VueRouter({routes, mode: "history"});

// router interceptor to check token for protected pages
router.beforeEach((to, from, next) => {
	if(to.meta.protected){
		axios.get("/api/token")
		.then(() => {
			next();
		})
		.catch(() => {
			next("/login");
		});
	}
	else{
		next();
	}
});

// i18n
const i18n = new VueI18n({
	locale: "en",
	fallbackLocale: "en"
});

// main component
new Vue({
	el: "#app",
	store: NewStore(),
	router,
	i18n
});
