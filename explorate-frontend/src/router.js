import {createRouter, createWebHistory} from 'vue-router'
// pages
import HomePage from './components/HomePage.vue'
import DatabasePage from './components/database/DatabasePage.vue'
import VolunteerResourcesPage from './components/volunteers/VolunteerResourcesPage.vue'
import VolunteerProfilePage from './components/volunteers/VolunteerProfilePage.vue'
import PartnerPage from './components/ngos/PartnerPage.vue'
import AboutPage from './components/AboutPage.vue'
import PickSignUpPage from './components/sessions/PickSignUpPage.vue'
import VolunteerSignUpPage from './components/volunteers/VolunteerSignUpPage.vue'
import PartnerSignUpPage from './components/ngos/PartnerSignUpPage.vue'
import SignInPage from './components/sessions/SignInPage.vue'
import SignOutPage from './components/sessions/SignOutPage.vue'

import AddOppPage from './components/database/AddOppPage.vue'
import EditOppPage from './components/database/EditOppPage.vue'

import TestT from './components/TestT.vue'
// import TestPage from './components/TestPage.vue'
// import HeaderMod from './components/HeaderMod.vue'

const routes=[
    {
        name:'HomePage',
        component: HomePage,
        path: '/'
    },
    {
        name: 'DatabasePage',
        component: DatabasePage,
        path: '/database'
    },
    {
        name:'VolunteerResourcesPage',
        component: VolunteerResourcesPage,
        path: '/volunteers/resources'
    },
    {
        name: 'VolunteerProfilePage',
        component: VolunteerProfilePage,
        path: '/volunteers/profile'
    },
    {
        name: 'PartnerPage',
        component: PartnerPage,
        path: '/partners'
    },
    {
        name:'AboutPage',
        component: AboutPage,
        path: '/about'
    },
    {
        name: 'SignInPage',
        component: SignInPage,
        path: '/signin'
    },
    {
        name: 'PickSignUpPage',
        component: PickSignUpPage,
        path: '/signup'
    },
    {
        name: 'VolunteerSignUpPage',
        component: VolunteerSignUpPage,
        path: '/signup/volunteers'
    },
    {
        name: 'PartnerSignUpPage',
        component: PartnerSignUpPage,
        path: '/signup/partners'
    },
    {
        name: 'SignOutPage',
        component: SignOutPage,
        path: '/signout'
    },
    {
        name: 'AddOppPage',
        component: AddOppPage,
        path: '/database/add'
    },
    {
        name: 'EditOppPage',
        component: EditOppPage,
        path: '/database/edit/:id'
    },
    {
        name: 'TestT',
        component: TestT,
        path: '/test'
    }
];

const router = createRouter({
    history:createWebHistory(),
    routes
});

export default router