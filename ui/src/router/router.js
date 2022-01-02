import  Vue from  "vue"
import VueRouter from "vue-router";

import home from "@/components/pages/home";
import productAdd from "@/components/product/productAdd";

Vue.use(VueRouter)



const routes = [
    {
        path:"/",
        component:home
    },
    {
        path: "/addProduct",
        component: productAdd
    }
]

export  const router  = new VueRouter({
    mode:"history",
    routes:routes
})