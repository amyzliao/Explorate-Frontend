<template>
    <component :is="view" />
    <h1>Database Page</h1>
    <router-link to="/add">Add a new opportunity</router-link>
    <DatabaseMod />
    <FooterMod />
</template>

<script>
import HeaderModVolunteer from './HeaderModVolunteer.vue'
import HeaderModNgo from './HeaderModNgo.vue';
import HeaderModSignedOut from './HeaderModSignedOut.vue';
import FooterMod from './FooterMod.vue'
import DatabaseMod from './DatabaseMod.vue';

export default {
    name: 'DatabasePage',
    components: {
        HeaderModVolunteer,
        FooterMod,
        HeaderModNgo,
        DatabaseMod,
        HeaderModSignedOut
    },
    computed: {
        view() {
            let user = localStorage.getItem("user-info");
            if (user) {
                const name = JSON.parse(user).account;
                return name == 'volunteer' ? 'HeaderModVolunteer' : 'HeaderModNgo'
            } else {
                return 'HeaderModSignedOut'
            }
        }
    }
}
</script>