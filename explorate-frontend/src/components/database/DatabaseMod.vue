<template>
    <!-- API must be running for this to look right, shouldn't be a problem unless server down -->
    <div class="flex justify-center relative">
        <img src="../../assets/Group_219.png" alt="ngos">
        <div class="absolute">
            <!-- search bar and filters -->
            <div class="relative mt-28 h-10 mx-20 h-[121px]">
                <!-- search bar -->
                <input type="text" class="flex justify-start h-[46px] min-w-[600px] max-w-[946px] py-3.5 px-6 rounded-[50px] bg-white bg-opacity-40 placeholder-[#FEF5DA] text-white" placeholder="Find a Program">
                <!-- Sort -->
                <div class="absolute top-0 right-0 border rounded-[20px] py-[6px] px-[10px] w-[135px] h-[42px] flex justify-center items-center text-sm text-white gap-[10px] border-[#F2DFC1] bg-white bg-opacity-20 ml-1">
                    Sort by: A-Z
                    <svg class="w-3 h-3 text-white" aria-hidden="true" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 14 8">
                        <path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m1 1 5.326 5.7a.909.909 0 0 0 1.348 0L13 1"/>
                    </svg>
                </div>
                <!-- filters -->
                <div class="absolute flex w-[689px] gap-[22px] bottom-0 left-0 justify-center">
                    <!-- SDGs -->
                    <button id="dropdownDefaultButton" data-dropdown-toggle="dropdownSDG" class="w-[246px] h-11 border-2 border-white border-opacity-50 rounded-[36px] text-sm font-medium text-white flex justify-center items-center bg-transparent" type="button">
                        <svg class="w-5 h-5 mr-2 opacity-50" fill="#ffffff" width="64px" height="64px" viewBox="0 0 32 32" version="1.1" xmlns="http://www.w3.org/2000/svg"><g id="SVGRepo_bgCarrier" stroke-width="0"></g><g id="SVGRepo_tracerCarrier" stroke-linecap="round" stroke-linejoin="round"></g><g id="SVGRepo_iconCarrier"> <title>left-bubble</title> <path d="M0 20q0 2.496 1.76 4.256t4.256 1.76l4 5.984 4-5.984h12q2.464 0 4.224-1.76t1.76-4.256v-13.984q0-2.496-1.76-4.256t-4.224-1.76h-20q-2.496 0-4.256 1.76t-1.76 4.256v13.984zM4 20v-13.984q0-0.832 0.576-1.408t1.44-0.608h20q0.8 0 1.408 0.608t0.576 1.408v13.984q0 0.832-0.576 1.44t-1.408 0.576h-13.344l-2.656 4-2.688-4h-1.312q-0.832 0-1.44-0.576t-0.576-1.44z"></path> </g></svg>
                        SDGs
                        <svg class="w-2.5 h-2.5 ml-2.5" aria-hidden="true" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 10 6">
                        <path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m1 1 4 4 4-4"/>
                    </svg></button>
                    <!-- Dropdown menu -->
                    <div id="dropdownSDG" class="z-10 hidden bg-transparent divide-y divide-gray-100 rounded-lg shadow w-[246px]">
                        <ul class="py-2 text-sm text-white bg-[#E7513F] bg-opacity-80" aria-labelledby="dropdownDefaultButton">
                            <li>
                                <a href="#" class="block px-4 py-2 hover:bg-[#E7513F] hover:bg-opacity-100">1. No Poverty</a>
                            </li>
                            <li>
                                <a href="#" class="block px-4 py-2 hover:bg-[#E7513F] hover:bg-opacity-100">2. Zero Hunger</a>
                            </li>
                            <li>
                                <a href="#" class="block px-4 py-2 hover:bg-[#E7513F] hover:bg-opacity-100">3. Good Health and Wellbeing</a>
                            </li>
                            <li>
                                <a href="#" class="block px-4 py-2 hover:bg-[#E7513F] hover:bg-opacity-100">4. Quality Education</a>
                            </li>
                            <li>
                                <a href="#" class="block px-4 py-2 hover:bg-[#E7513F] hover:bg-opacity-100">5. Gender Equality</a>
                            </li>
                            <li>
                                <a href="#" class="block px-4 py-2 hover:bg-[#E7513F] hover:bg-opacity-100">6. Clean Water and Sanitation</a>
                            </li>
                            <li>
                                <a href="#" class="block px-4 py-2 hover:bg-[#E7513F] hover:bg-opacity-100">7. Affordable and Clean Energy</a>
                            </li>
                            <li>
                                <a href="#" class="block px-4 py-2 hover:bg-[#E7513F] hover:bg-opacity-100">8. Decent Work and Economic Growth</a>
                            </li>
                            <li>
                                <a href="#" class="block px-4 py-2 hover:bg-[#E7513F] hover:bg-opacity-100">9. Industry, Innovation, and Infrastructure</a>
                            </li>
                            <li>
                                <a href="#" class="block px-4 py-2 hover:bg-[#E7513F] hover:bg-opacity-100">10. Reduced Inequalities</a>
                            </li>
                            <li>
                                <a href="#" class="block px-4 py-2 hover:bg-[#E7513F] hover:bg-opacity-100">11. Sustainable Cities and Communities</a>
                            </li>
                            <li>
                                <a href="#" class="block px-4 py-2 hover:bg-[#E7513F] hover:bg-opacity-100">12. Responsible Consumption and Production</a>
                            </li>
                            <li>
                                <a href="#" class="block px-4 py-2 hover:bg-[#E7513F] hover:bg-opacity-100">13. Climate Action</a>
                            </li>
                            <li>
                                <a href="#" class="block px-4 py-2 hover:bg-[#E7513F] hover:bg-opacity-100">14. Life Below Water</a>
                            </li>
                            <li>
                                <a href="#" class="block px-4 py-2 hover:bg-[#E7513F] hover:bg-opacity-100">15. Life on Land</a>
                            </li>
                            <li>
                                <a href="#" class="block px-4 py-2 hover:bg-[#E7513F] hover:bg-opacity-100">16. Peace, Justice, and Strong Institutions</a>
                            </li>
                            <li>
                                <a href="#" class="block px-4 py-2 hover:bg-[#E7513F] hover:bg-opacity-100">17. Partnership for the Goals</a>
                            </li>
                        </ul>
                    </div>
                    <!-- Location -->
                    <button id="dropdownDefaultButton" data-dropdown-toggle="dropdownLocation" class="w-[246px] h-11 border-2 border-white border-opacity-50 rounded-[36px] text-sm font-medium text-white flex justify-center items-center bg-transparent" type="button">
                        <svg class="h-6 w-6 mr-2 opacity-50" viewBox="0 0 24 24" fill="none" xmlns="http://www.w3.org/2000/svg" stroke="#ffffff"><g id="SVGRepo_bgCarrier" stroke-width="0"></g><g id="SVGRepo_tracerCarrier" stroke-linecap="round" stroke-linejoin="round"></g><g id="SVGRepo_iconCarrier"> <path d="M12 7V12L10.5 14.5M21 12C21 16.9706 16.9706 21 12 21C7.02944 21 3 16.9706 3 12C3 7.02944 7.02944 3 12 3C16.9706 3 21 7.02944 21 12Z" stroke="#ffffff" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"></path> </g></svg>
                        Location
                        <svg class="w-2.5 h-2.5 ml-2.5" aria-hidden="true" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 10 6">
                        <path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m1 1 4 4 4-4"/>
                    </svg></button>
                    <!-- Dropdown menu -->
                    <div id="dropdownLocation" class="z-10 hidden bg-transparent divide-y divide-gray-100 rounded-lg shadow w-[246px]">
                        <ul class="py-2 text-sm text-white bg-[#E7513F] bg-opacity-80" aria-labelledby="dropdownDefaultButton">
                        <li>
                            <a href="#" class="block px-4 py-2 hover:bg-[#E7513F] hover:bg-opacity-100">Asia</a>
                        </li>
                        <li>
                            <a href="#" class="block px-4 py-2 hover:bg-[#E7513F] hover:bg-opacity-100">Africa</a>
                        </li>
                        <li>
                            <a href="#" class="block px-4 py-2 hover:bg-[#E7513F] hover:bg-opacity-100">North America</a>
                        </li>
                        <li>
                            <a href="#" class="block px-4 py-2 hover:bg-[#E7513F] hover:bg-opacity-100">South America</a>
                        </li>
                        <li>
                            <a href="#" class="block px-4 py-2 hover:bg-[#E7513F] hover:bg-opacity-100">Europe</a>
                        </li>
                        <li>
                            <a href="#" class="block px-4 py-2 hover:bg-[#E7513F] hover:bg-opacity-100">Australia</a>
                        </li>
                        <li>
                            <a href="#" class="block px-4 py-2 hover:bg-[#E7513F] hover:bg-opacity-100">Antarctica</a>
                        </li>
                        </ul>
                    </div>
                    <div class="w-[246px] h-11 border-2 border-white border-opacity-50 rounded-[36px] text-sm font-medium text-white flex justify-center items-center bg-transparent">
                        <div class="h-[32px] flex justify-center items-center">
                            <svg class="h-5 w-5 mr-2 opacity-50" viewBox="0 0 20 20" version="1.1" xmlns="http://www.w3.org/2000/svg" xmlns:xlink="http://www.w3.org/1999/xlink" fill="#000000" stroke="#000000"><g id="SVGRepo_bgCarrier" stroke-width="0"></g><g id="SVGRepo_tracerCarrier" stroke-linecap="round" stroke-linejoin="round"></g><g id="SVGRepo_iconCarrier"> <title>profile_round [#ffffff1342]</title> <desc>Created with Sketch.</desc> <defs> </defs> <g id="Page-1" stroke="none" stroke-width="1" fill="none" fill-rule="evenodd"> <g id="Dribbble-Light-Preview" transform="translate(-140.000000, -2159.000000)" fill="#ffffff"> <g id="icons" transform="translate(56.000000, 160.000000)"> <path d="M100.562548,2016.99998 L87.4381713,2016.99998 C86.7317804,2016.99998 86.2101535,2016.30298 86.4765813,2015.66198 C87.7127655,2012.69798 90.6169306,2010.99998 93.9998492,2010.99998 C97.3837885,2010.99998 100.287954,2012.69798 101.524138,2015.66198 C101.790566,2016.30298 101.268939,2016.99998 100.562548,2016.99998 M89.9166645,2004.99998 C89.9166645,2002.79398 91.7489936,2000.99998 93.9998492,2000.99998 C96.2517256,2000.99998 98.0830339,2002.79398 98.0830339,2004.99998 C98.0830339,2007.20598 96.2517256,2008.99998 93.9998492,2008.99998 C91.7489936,2008.99998 89.9166645,2007.20598 89.9166645,2004.99998 M103.955674,2016.63598 C103.213556,2013.27698 100.892265,2010.79798 97.837022,2009.67298 C99.4560048,2008.39598 100.400241,2006.33098 100.053171,2004.06998 C99.6509769,2001.44698 97.4235996,1999.34798 94.7348224,1999.04198 C91.0232075,1998.61898 87.8750721,2001.44898 87.8750721,2004.99998 C87.8750721,2006.88998 88.7692896,2008.57398 90.1636971,2009.67298 C87.1074334,2010.79798 84.7871636,2013.27698 84.044024,2016.63598 C83.7745338,2017.85698 84.7789973,2018.99998 86.0539717,2018.99998 L101.945727,2018.99998 C103.221722,2018.99998 104.226185,2017.85698 103.955674,2016.63598" id="profile_round-[#ffffff1342]"> </path> </g> </g> </g> </g></svg>
                            Skills
                        </div>
                    </div>
                </div>
            </div>
            <!-- cards -->
            <div class="flex flex-row p-[70px] h-[1617px] overflow-scroll overflow-x-hidden mt-6">
                <div class="grid grid-cols-3 gap-10 inline">
                    <div v-for="opp in oppCollection" :key="opp.ID">
                        <DatabaseCard :name="opp.Org_name.String" :location="opp.Org_location.String"
                            :mission="opp.Impact_space.String" :projects="opp.Projects.String" :link="opp.Website.String" class="w-[379px] h-[530px]"/>
                    </div>
                </div>
            </div>
        </div>
    </div>
    
</template>

<script>
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
        async loadData() {
            let result = await axios.get("https://unry0i0624.execute-api.us-east-2.amazonaws.com/prod/ngo_opps")
            this.oppCollection = result.data
            console.log(this.oppCollection)
        },
        deleteThisOpp(id) {
            axios.delete("https://unry0i0624.execute-api.us-east-2.amazonaws.com/prod/ngo_opps" + id).then(result => {
                console.warn(result)
                if (result.status == 200){
                    this.loadData()
                }
            })
        }
    },
    mounted() {
        this.loadData()
    }
}
</script>
