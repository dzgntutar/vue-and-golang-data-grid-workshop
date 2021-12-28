const state = {
    products: [],
};

const getters = {};

const mutations = {
    addProductToSatate(state, product) {
        state.products.push(product)
        console.log(state.products)
    }
};

const actions = {
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