import {createRouter, createWebHistory} from 'vue-router'
// pages
import HomePage from './components/HomePage.vue'
import AboutPage from './components/AboutPage.vue'
import DatabasePage from './components/DatabasePage.vue'
import ResourcePage from './components/ResourcePage.vue'
// import TestPage from './components/TestPage.vue'

const routes=[
    {
        name:'HomePage',
        component: HomePage,
        path: '/'
    },
    {
        name:'AboutPage',
        component: AboutPage,
        path: '/about'
    },
    {
        name: 'DatabasePage',
        component: DatabasePage,
        path: '/database'
    },
    {
        name: 'ResourcePage',
        component: ResourcePage,
        path: '/resources'
    }
];

const router = createRouter({
    history:createWebHistory(),
    routes
});

export default router