<template>
    <HeaderMod />
    <h1
        class="mt-44 mb-4 text-4xl font-extrabold leading-none tracking-tight text-gray-900 md:text-5xl lg:text-6xl dark:text-white">
        Edit an NGO in the database</h1>
    <div className="big-form">
        <!-- <input type="text" name="name" v-model="opp.name" placeholder="NGO Name" />
        <input type="text" name="contact" v-model="opp.contact" placeholder="NGO Contact" />
        <input type="text" name="title" v-model="opp.title" placeholder="Program Title" />
        <input type="text" name="description" v-model="opp.description" placeholder="Program Description" />
        <input type="text" name="location" v-model="opp.location" placeholder="Program Location" />
        <input type="text" name="website" v-model="opp.website" placeholder="Link to website" /> -->
        <input placeholder="org_name" v-model="opp.Org_name">
        <input placeholder="org_location" v-model="opp.Org_location">
        <input placeholder="mission statement" v-model="opp.Mission_statement">
        <input placeholder="impact space" v-model="opp.Impact_space">
        <input placeholder="un_goals" v-model="opp.Un_goals">
        <input placeholder="projects" v-model="opp.Projects">
        <input placeholder="values" v-model="opp.Values">
        <input placeholder="benefits" v-model="opp.Benefits">
        <input placeholder="ethical criteria" v-model="opp.Ethical_criteria">
        <input placeholder="ethical details" v-model="opp.Ethical_criteria_details">
        <input placeholder="volunteer_location" v-model="opp.Volunteer_location">
        <input placeholder="volunteer skills" v-model="opp.Volunteer_skills">
        <input placeholder="skill projects" v-model="opp.Skill_projects">
        <input placeholder="language" v-model="opp.Language_requirements">
        <input placeholder="volunteer qualities" v-model="opp.Volunteer_qualities">
        <input placeholder="work_env" v-model="opp.Work_env">
        <input placeholder="accomodations" v-model="opp.Accomodations">
        <input placeholder="cost" v-model="opp.Cost">
        <input placeholder="website" v-model="opp.Website">
        <input placeholder="contact" v-model="opp.Contact">
        <input placeholder="photos" v-model="opp.Photos">
        <input placeholder="refer_others" v-model="opp.Refer_others">
        <input placeholder="additional info" v-model="opp.Additional_info">
        <input placeholder="finding_volunteers" v-model="opp.Finding_volunteers">
        <input placeholder="logo" v-model="opp.Org_logo">
        <input placeholder="duration" v-model="opp.Program_duration">
        <input placeholder="timeframe" v-model="opp.Project_timeframe">
        <button v-on:click="editOpp" type="button"
            class="focus:outline-none text-white bg-red-700 hover:bg-red-800 focus:ring-4 focus:ring-red-300 font-medium rounded-lg text-sm px-5 py-2.5 mr-2 mb-2 dark:bg-red-600 dark:hover:bg-red-700 dark:focus:ring-red-900">
            Save Edits
        </button>
        <button v-on:click="deleteOpp" type="button"
            class="focus:outline-none text-white bg-red-700 hover:bg-red-800 focus:ring-4 focus:ring-red-300 font-medium rounded-lg text-sm px-5 py-2.5 mr-2 mb-2 dark:bg-red-600 dark:hover:bg-red-700 dark:focus:ring-red-900">
            Delete This Posting
        </button>
    </div>
    <FooterMod />
</template>

<script>
import HeaderMod from '../header/HeaderMod.vue'
import FooterMod from '../FooterMod.vue'
import axios from 'axios'

export default {
    name: 'EditOppPage',
    components: {
        HeaderMod,
        FooterMod
    },
    data() {
        return {
            // opp: {
            //     name:'',
            //     contact:'',
            //     title:'',
            //     description:'',
            //     location:'',
            //     website:''
            // },\
            oppId: '',
            opp: {
                ID: '',
                Org_name: '',
                Org_location: '',
                Mission_statement: '',
                Impact_space: '',
                Un_goals: '',
                Projects: '',
                Values: '',
                Benefits: '',
                Ethical_criteria: '',
                Ethical_criteria_details: '',
                Volunteer_location: '',
                Volunteer_skills: '',
                Skill_projects: '',
                Language_requirements: '',
                Volunteer_qualities: '',
                Work_env: '',
                Accomodations: '',
                Cost: '',
                Website: '',
                Contact: '',
                Photos: '',
                Refer_others: '',
                Additional_info: '',
                Finding_volunteers: '',
                Org_logo: '',
                Program_duration: '',
                Project_timeframe: '',
            }
        }
    },
    methods: {
        async deleteOpp() {
            await axios.delete("https://unry0i0624.execute-api.us-east-2.amazonaws.com/prod/ngo_opps/" + this.oppId)
                .then(result => {
                    console.warn(result)
                    if (result.status == 200) {
                        this.$router.push({ name: 'DatabasePage' });
                    }
                }).catch((error) => {
                    console.warn(error)
                })
        },
        async editOpp() {
            // let result = dbUpdate('opportunities', this.oppId, {
            //     name: this.opp.name,
            //     contact: this.opp.contact,
            //     title: this.opp.title,
            //     description: this.opp.description,
            //     location: this.opp.location,
            //     website: this.opp.website
            // });
            console.log("https://unry0i0624.execute-api.us-east-2.amazonaws.com/prod/ngo_opps/" + this.oppId)
            await axios.patch("https://unry0i0624.execute-api.us-east-2.amazonaws.com/prod/ngo_opps/" + this.oppId, {
                Org_name: this.opp.Org_name,
                Org_location: this.opp.Org_location,
                Mission_statement: this.opp.Mission_statement,
                Impact_space: this.opp.Impact_space,
                Un_goals: this.opp.Un_goals,
                Projects: this.opp.Projects,
                Values: this.opp.Values,
                Benefits: this.opp.Benefits,
                Ethical_criteria: this.opp.Ethical_criteria,
                Ethical_criteria_details: this.opp.Ethical_criteria_details,
                Volunteer_location: this.opp.Volunteer_location,
                Volunteer_skills: this.opp.Volunteer_skills,
                Skill_projects: this.opp.Skill_projects,
                Language_requirements: this.opp.Language_requirements,
                Volunteer_qualities: this.opp.Volunteer_qualities,
                Work_env: this.opp.Work_env,
                Accomodations: this.opp.Accomodations,
                Cost: this.opp.Cost,
                Website: this.opp.Website,
                Contact: this.opp.Contact,
                Photos: this.opp.Photos,
                Refer_others: this.opp.Refer_others,
                Additional_info: this.opp.Additional_info,
                Finding_volunteers: this.opp.Finding_volunteers,
                Org_logo: this.opp.Org_logo,
                Program_duration: this.opp.Program_duration,
                Project_timeframe: this.opp.Project_timeframe
            }).then(result => {
                console.warn("result", result)
                if (result.status == 200) {
                    this.$router.push({ name: 'DatabasePage' });
                }
            }).catch((error) => {
                console.warn(error)
            })
            // console.log("http://localhost:3000/ngo_opps/"+this.oppId)
        }
    },
    async mounted() {
        this.oppId = this.$route.params.id
        let result = {
            Org_name: "",
            Org_location: "",
            Mission_statement: "",
            Impact_space: "",
            Un_goals: "",
            Projects: "",
            Values: "",
            Benefits: "",
            Ethical_criteria: "",
            Ethical_criteria_details: "",
            Volunteer_location: "",
            Volunteer_skills: "",
            Skill_projects: "",
            Language_requirements: "",
            Volunteer_qualities: "",
            Work_env: "",
            Accomodations: "",
            Cost: "",
            Website: "",
            Contact: "",
            Photos: "",
            Refer_others: "",
            Additional_info: "",
            Finding_volunteers: "",
            Org_logo: "",
            Program_duration: "",
            Project_timeframe: ""
        }
        await axios.get("https://unry0i0624.execute-api.us-east-2.amazonaws.com/prod/ngo_opps/" + this.oppId)
            .then(res => {
                result = res
                console.log("result", result)
            }).catch((error) => {
                console.warn(error)
            })
        const data = result.data

        this.opp.Org_name = data.Org_name.String
        this.opp.Org_location = data.Org_location.String
        this.opp.Mission_statement = data.Mission_statement.String
        this.opp.Impact_space = data.Impact_space.String
        this.opp.Un_goals = data.Un_goals.String
        this.opp.Projects = data.Projects.String
        this.opp.Values = data.Values.String
        this.opp.Benefits = data.Benefits.String
        this.opp.Ethical_criteria = data.Ethical_criteria.String
        this.opp.Ethical_criteria_details = data.Ethical_criteria_details.String
        this.opp.Volunteer_location = data.Volunteer_location.String
        this.opp.Volunteer_skills = data.Volunteer_skills.String
        this.opp.Skill_projects = data.Skill_projects.String
        this.opp.Language_requirements = data.Language_requirements.String
        this.opp.Volunteer_qualities = data.Volunteer_qualities.String
        this.opp.Work_env = data.Work_env.String
        this.opp.Accomodations = data.Accomodations.String
        this.opp.Cost = data.Cost.String
        this.opp.Website = data.Website.String
        this.opp.Contact = data.Contact.String
        this.opp.Photos = data.Photos.String
        this.opp.Refer_others = data.Refer_others.String
        this.opp.Additional_info = data.Additional_info.String
        this.opp.Finding_volunteers = data.Finding_volunteers.String
        this.opp.Org_logo = data.Org_logo.String
        this.opp.Program_duration = data.Program_duration.String
        this.opp.Project_timeframe = data.Project_timeframe.String
        console.log(this.opp)
        console.log(result.data)
    }
}
</script>