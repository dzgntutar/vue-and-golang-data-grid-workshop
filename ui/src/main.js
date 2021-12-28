import Vue from 'vue'
import App from './App.vue'

import VueRouter from "vue-router";

import {router} from "./router/router"
import {store} from "./store/store"

Vue.use(VueRouter)

Vue.config.productionTip = false

new Vue({
    render: h => h(App), router: router, store: store
}).$mount('#app')
