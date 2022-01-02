import axios from "axios";

const state = {
    products: [],
};

const getters = {
    getAllProducts(state) {
        return state.products
    }
};

const mutations = {
    addProductToSatate(state, product) {
        state.products.push(product)
        console.log(state.products)
    }
};

const actions = {
    getProductsFromApi() {

        console.log("getProductsFromApi..")
        axios.get("").then(response => {
            console.log(response)
        })
    },
    addProduct({commit}, product) {
        commit("addProductToSatate", product)
    }
}

export default {
    state,
    getters,
    mutations,
    actions
}