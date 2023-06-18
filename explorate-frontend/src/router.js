import {createRouter, createWebHistory} from 'vue-router'
// pages
import HomePage from './components/HomePage.vue'
import AboutPage from './components/AboutPage.vue'
import DatabasePage from './components/DatabasePage.vue'
import PartnerPage from './components/PartnerPage.vue'
import AddOppPage from './components/AddOppPage.vue'
import EditOppPage from './components/EditOppPage.vue'
import SignUp from './components/SignUp.vue'
import SignIn from './components/SignIn.vue'
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
        name: 'PartnerPage',
        component: PartnerPage,
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
    },
    {
        name: 'SignInPage',
        component: SignIn,
        path: '/signin'
    }
];

const router = createRouter({
    history:createWebHistory(),
    routes
});

export default router