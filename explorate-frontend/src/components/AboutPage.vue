<template>
    <component :is="view"></component>
    <h1>About Page</h1>

    <FooterMod />
</template>

<script>
import HeaderModVolunteer from './HeaderModVolunteer.vue'
import HeaderModNgo from './HeaderModNgo.vue'
import HeaderModSignedOut from './HeaderModSignedOut'
import FooterMod from './FooterMod.vue'

export default {
    name: 'AboutPage',
    components: {
        HeaderModVolunteer,
        FooterMod,
        HeaderModNgo,
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