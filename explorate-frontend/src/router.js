import {createRouter, createWebHistory} from 'vue-router'
// pages
import HomePage from './components/HomePage.vue'
import AboutPage from './components/AboutPage.vue'
import DatabasePage from './components/DatabasePage.vue'
import ResourcePage from './components/ResourcePage.vue'
import AddOppPage from './components/AddOppPage.vue'
import EditOppPage from './components/EditOppPage.vue'
import SignUp from './components/SignUp.vue'
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
    },
    {
        name: 'AddOppPage',
        component: AddOppPage,
        path: '/add'
    },
    {
        name: 'EditOppPage',
        component: EditOppPage,
        path: '/edit/:id'
    },
    {
        name: 'SignUpPage',
        component: SignUp,
        path: '/signup'
    }
];

const router = createRouter({
    history:createWebHistory(),
    routes
});

export default router