import Vue from 'vue'
import App from './App.vue'

import VueRouter from "vue-router";
import axios  from "axios"

import {router} from "./router/router"
import {store} from "./store/store"

Vue.use(VueRouter)
axios.defaults.baseURL = "http://localhost:3000/"

Vue.config.productionTip = false

new Vue({
    render: h => h(App), router: router, store: store
}).$mount('#app')
