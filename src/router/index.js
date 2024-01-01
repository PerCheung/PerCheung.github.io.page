import Vue from 'vue'
import VueRouter from 'vue-router'
import HomeView from "@/views/HomeView.vue";
import TOC from "@/views/TOC.vue";


Vue.use(VueRouter)

const routes = [{
    path: '/', name: 'home', component: HomeView
}, {
    path: '/toc', name: 'TOC', component: TOC
}]

const router = new VueRouter({
    routes, mode: 'history'
})

export default router
