import  Vue from  "vue"
import VueRouter from "vue-router";

import Home from "@/components/pages/Home";
import productAdd from "@/components/product/productAdd";

Vue.use(VueRouter)



const routes = [
    {
        path:"/",
        component:Home
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