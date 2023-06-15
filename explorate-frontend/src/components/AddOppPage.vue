<template>
    <HeaderMod/>
    <h1>Add a new program to the database</h1>
    <div class="form">
        <input type="text" v-model="name" placeholder="NGO Name" />
        <input type="text" v-model="contact" placeholder="NGO Contact" />
        <input type="text" v-model="title" placeholder="Program Title" />
        <input type="text" v-model="description" placeholder="Program Description" />
        <input type="text" v-model="location" placeholder="Program Location" />
        <input type="text" v-model="website" placeholder="Link to website" />
        <button v-on:click="addOpp">Add new program</button>
    </div>
    <FooterMod/>
</template>

<script>
import { createOpp } from '@/firebase'
// import { reactive } from 'vue'
import HeaderMod from './HeaderMod.vue'
import FooterMod from './FooterMod.vue'

export default {
    name: 'AddOppPage',
    components: {
        HeaderMod,
        FooterMod
    },
    data() {
        return {
            name:'',
            contact:'',
            title:'',
            description:'',
            location:'',
            website:''
        }
    },
    methods: {
        async addOpp() {
            // console.log("addOpp called")
            // const form = reactive({})
            let result = await createOpp({
                name: this.name,
                contact: this.contact,
                title: this.title,
                description: this.description,
                location: this.location,
                website: this.website
            });
            console.warn(result)
            if (result == 201) {
                this.$router.push({name:'DatabasePage'});
            }
            // console.log("wat")
        }
    }
}
</script>