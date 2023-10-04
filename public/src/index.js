import Vue from "vue";
import VueRouter from "vue-router";
import Vuex from "vuex";
import VueI18n from "vue-i18n";
import axios from "axios";

import "./main.scss";
import NewStore from "./store/store.js";
import * as pages from "./pages";
import {getLocale} from "./util/locale.js";

Vue.use(VueRouter);
Vue.use(Vuex);
Vue.use(VueI18n);
Vue.config.productionTip = false;
Vue.config.devtools = false;


// router
const routes = [
	{path: "/", component: pages.Overview, meta: {protected: true}},
	{path: "/login", component: pages.Login},
	{path: "/server", component: pages.Server, meta: {protected: true}},
	{path: "/logs", component: pages.Logs, meta: {protected: true}},
	{path: "/live", component: pages.Live, meta: {protected: true}},
	{path: "/status", component: pages.Status},
	{path: "*", component: pages.Error404}
];

let router = new VueRouter({routes, mode: "history"});

// router interceptor to check token for protected pages
router.beforeEach((to, from, next) => {
	if (to.meta.protected) {
		axios.get("/api/token")
		.then(() => {
			next();
		})
		.catch(() => {
			next("/login");
		});
	}
	else {
		next();
	}
});

// token interceptor for every request
axios.interceptors.request.use((config) => {
	const token = window.localStorage.getItem("token");

	if (token) {
		config.headers.Authorization = `Bearer ${token}`;
	}

	return config;
}, (err) => {
	return Promise.reject(err);
});

// response error handler
axios.interceptors.response.use(function (response) {
    return response;
  }, err => {
	if (err.response.data.message) {
		console.log(err.response.data);
	}

	if (err.response.status === 401) {
		router.push("/login");
	}

	return Promise.reject(err);
});


// i18n
const i18n = new VueI18n({
	locale: getLocale(),
	fallbackLocale: "en"
});

Vue.mixin({
	computed: {
		is_admin() {
			return this.$store.state.auth.admin;
		},
		is_mod() {
			return this.$store.state.auth.mod;
		},
		is_ro() {
			return this.$store.state.auth.read_only;
		}
	}
});

// main component
new Vue({
	el: "#app",
	store: NewStore(),
	router,
	i18n
});
