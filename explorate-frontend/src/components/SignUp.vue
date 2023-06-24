<template>
    <HeaderModSignedOut />
        <div class="signup">
            <img class="logo" src="../assets/-Insert_image_here-.svg.png"/>
            <h1>Sign Up</h1>
            <button class="google">Sign Up with Google</button>
            <input type="text" placeholder="Full Name" v-model="name"/>
            <input type="text" placeholder="Email" v-model="email"/>
            <input type="password" placeholder="Password" v-model="password"/>
            <button class="signuppagebutton" v-on:click="signUp">Sign Up</button>
            <button class="signuppagesigninbutton">Sign In</button>
        </div>
    <FooterMod/>
</template>

<script>
import { dbCreate, dbGetUser } from '@/firebase';
import FooterMod from './FooterMod.vue';
import HeaderModSignedOut from './HeaderModSignedOut.vue';
export default {
    name: 'SignUp',
    data() {
        return {
            name: '',
            email: '',
            account: 'volunteer',
            password: ''
        }
    },
    components: {
        HeaderModSignedOut,
        FooterMod
    },
    methods: {
        async signUp() {
            let result = dbCreate('users', 
            {name: this.name,
            contact: this.email,
            password: this.password,
            account: this.account})

            if (result == 201) {
                dbGetUser(this.email, this.password).then((res) => {
                    let userId = res.docs[0].id
                    // console.log(userId)
                    dbCreate('volunteers', 
                    {uid: userId,
                    name: this.name,
                    contact: this.email})
                })
                this.$router.push({name:'HomePage'});
            }
        }
    }
}
</script>

