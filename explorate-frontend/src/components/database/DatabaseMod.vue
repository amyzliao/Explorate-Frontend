<script>
import { dbUseLoad, dbDelete } from '@/firebase'
import DatabaseCard from './DatabaseCard.vue'
import axios from 'axios'

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
        },
        async loadLocal() {
            let result = await axios.get('http://localhost:3000/ngo_opps')
            console.log(result)
            this.oppCollection = result.data
            console.log(this.oppCollection)
            // console.log(result[0])
        }
    },
    async mounted() {
        // this.loadData()
        this.loadLocal()
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
                <div v-for="opp in oppCollection" :key="opp.Org_ID">
                    <DatabaseCard :name="opp.Org_name.String" :location="opp.Org_location.String" :mission="opp.Mission_statement.String" :projects="opp.Projects.String" :link="opp.Website.String" />
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
            <td>Functions</td>
        </tr>
        <tr v-for="opp in oppCollection" :key="opp.Org_ID">
            <td>{{ opp.Org_name.String }}</td>
            <td>{{ opp.Contact.String }}</td>
            <td>{{ opp.Impact_space.String }}</td>
            <td>{{ opp.Mission_statement.String }}</td>
            <td>{{ opp.Org_location.String }}</td>
            <td>{{ opp.Website.String }}</td>
            <td>
                <router-link :to="'/database/edit/' + opp.id">Edit</router-link>
                <button v-on:click="deleteThisOpp(opp.id)">Delete</button>
            </td>
        </tr>
    </table>
</template>



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
</style>