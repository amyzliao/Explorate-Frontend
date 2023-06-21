<template>
    <HeaderMod/>
    <h1>Edit opportunity</h1>
    <div class="form">
        <input type="text" name="name" v-model="opp.name" placeholder="NGO Name" />
        <input type="text" name="contact" v-model="opp.contact" placeholder="NGO Contact" />
        <input type="text" name="title" v-model="opp.title" placeholder="Program Title" />
        <input type="text" name="description" v-model="opp.description" placeholder="Program Description" />
        <input type="text" name="location" v-model="opp.location" placeholder="Program Location" />
        <input type="text" name="website" v-model="opp.website" placeholder="Link to website" />
        <button v-on:click="editOpp">Save changes</button>
    </div>
    <FooterMod/>
</template>

<script>
import { dbUpdate, dbGet } from '@/firebase'
import HeaderMod from './HeaderMod.vue'
import FooterMod from './FooterMod.vue'

export default {
    name: 'EditOppPage',
    components: {
        HeaderMod,
        FooterMod
    },
    data() {
        return {
            opp: {
                name:'',
                contact:'',
                title:'',
                description:'',
                location:'',
                website:''
            },
            oppId:''
        }
    },
    methods: {
        async editOpp() {
            console.log("editOpp called")
            console.log(this.opp)
            let result = await dbUpdate('opportunities', this.oppId, {
                name: this.opp.name,
                contact: this.opp.contact,
                title: this.opp.title,
                description: this.opp.description,
                location: this.opp.location,
                website: this.opp.website
            });
            console.warn(result)
            if (result == 200) {
                this.$router.push({name:'DatabasePage'});
            }
            // else popup?
        }
    },
    async mounted() {
        this.oppId = this.$route.params.id
        let result = await dbGet('opportunities', this.oppId)
        this.opp = result
        console.log(this.opp)
    }
}
</script>