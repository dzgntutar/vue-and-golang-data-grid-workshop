import Vue from "vue";
import Vuex from "vuex";
import product from "./module/product";


Vue.use(Vuex);

export const store = new Vuex.Store({
    modules: {product},
});