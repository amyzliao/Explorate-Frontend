<script>
import { dbUseLoad, dbDelete } from '@/firebase'
import DatabaseCard from './DatabaseCard.vue'

export default {
    name: 'DatabaseMod',
    components: {
        DatabaseCard
    },
    data() {
        return {
            oppCollection: []
        }
    },
    methods: {
        loadData() {
            let result = dbUseLoad('opportunities')
            this.oppCollection = result
            console.log('load data')
            console.warn(this.oppCollection.ownkeys)
        },
        deleteThisOpp(id) {
            let result = dbDelete('opportunities', id)
            console.warn(result)
            if (result == 200) {
                this.loadData()
            }
        }
    },
    async mounted() {
        this.loadData()
    }
}
</script>

<template>
    <div class="mx-10 mt-6">
        <div class="flex flex-row gap-5">
            <!-- column for filtering -->
            <div class="basis-1/4 bg-gray-200 flex-col">
                <div>filter a</div>
                <div>filter a</div>
                <div>filter a</div>
                <div>filter a</div>
                <div>filter a</div>
                <div>filter a</div>
                <div>filter a</div>
            </div>
            <!-- column for displaying cards -->
            <div class="basis-3/4 bg-gray-50 grid grid-cols-3 gap-5">
                <div v-for="opp in oppCollection" :key="opp.id">
                    <DatabaseCard :name="opp.name" :location="opp.location" :mission="opp.description" :projects="opp.contact" :link="opp.website" />
                </div>
            </div>
        </div>
    </div>

    <!-- old -->
    <table border="1">
        <tr class="col-title">
            <td>NGO</td>
            <td>Contact</td>
            <td>Title</td>
            <td>Description</td>
            <td>Location</td>
            <td>Website</td>
            <td v-if="loggedIn()">Functions</td>
        </tr>
<<<<<<< HEAD:explorate-frontend/src/components/DatabaseMod.vue
        <tr v-for="opp in oppCollection" :key="opp.Org_ID">
            <td>{{ opp.Org_name.String }}</td>
            <td>{{ opp.Contact.String }}</td>
            <td>{{ opp.Impact_space.String }}</td>
            <td>{{ opp.Mission_statement.String }}</td>
            <td>{{ opp.Org_location.String }}</td>
            <td>{{ opp.Website.String }}</td>
            <td v-if="loggedIn()">
                <button v-on:click="goToEdit(opp.Org_ID)">Edit</button>
                <br/>
=======
        <tr v-for="opp in oppCollection" :key="opp.id">
            <td>{{ opp.name }}</td>
            <td>{{ opp.contact }}</td>
            <td>{{ opp.title }}</td>
            <td>{{ opp.description }}</td>
            <td>{{ opp.location }}</td>
            <td>{{ opp.website }}</td>
            <td>
                <router-link :to="'/database/edit/' + opp.id">Edit</router-link>
>>>>>>> 21d6eb17ad94e01ea88264cc27a97ddec23f71d4:explorate-frontend/src/components/database/DatabaseMod.vue
                <button v-on:click="deleteThisOpp(opp.id)">Delete</button>
            </td>
            <!-- <h1>oppcollection</h1>
                <p>{{ oppCollection }}</p> -->
        </tr>
    </table>
</template>

<<<<<<< HEAD:explorate-frontend/src/components/DatabaseMod.vue
<script>
import { dbDelete } from '@/firebase';
import axios from 'axios';

export default {
    name: 'DatabaseMod',
    data() {
        return {
            oppCollection: []
        }
    },
    methods: {
        async loadData() {
            // let result = dbUseLoad('opportunities')
            let result = await axios.get("http://localhost:3000/ngo_opps")
            this.oppCollection = result.data
            console.warn(this.oppCollection)
        },
        deleteThisOpp(id) {
            let result = dbDelete('opportunities', id)
            console.warn(result)
            if (result == 200) {
                this.loadData()
            }
        },
        goToEdit(id) {
            const path = "edit/" + id
            this.$router.push(path)
        },
        loggedIn() {
            const user = localStorage.getItem("user-info")
            return user != null
        }
    },
    async mounted() {
        this.loadData()
    }
}
</script>
=======

>>>>>>> 21d6eb17ad94e01ea88264cc27a97ddec23f71d4:explorate-frontend/src/components/database/DatabaseMod.vue

<style>
table {
    margin: 10px;
}

td {
    width: 160px;
    height: 40px;
}

.col-title {
    color: blue;
}

td button {
    width: 60px;
    text-align: center;
}
</style>