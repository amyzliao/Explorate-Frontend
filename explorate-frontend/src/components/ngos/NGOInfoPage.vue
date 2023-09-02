<script>
import HeaderMod from '../header/HeaderMod.vue'
import FooterMod from '../FooterMod.vue'
import axios from 'axios'

export default {
    name: 'NGOInfoPage',
    components: {
        HeaderMod,
        FooterMod
    },
    data() {
        return {
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
    async mounted() {
        this.oppId = this.$route.params.id
        // let result = this.opp
        await axios.get("https://unry0i0624.execute-api.us-east-2.amazonaws.com/prod/ngo_opps/" + this.oppId)
            .then(result => {
                console.log("result", result)
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
                
            }).catch((error) => {
                console.warn(error)
            })
        // const data = result.data


        // console.log(this.opp)
        // console.log(result.data)
    }
}
</script>


<template>
    <HeaderMod />
    <h1
        class="mt-5 mb-4 text-4xl font-extrabold leading-none tracking-tight text-gray-900 md:text-5xl lg:text-6xl dark:text-white">
        {{ opp.Org_name }}
    </h1>
    <div class="w-1/2 mx-auto mb-20">
        <h4> {{ opp.ID }}</h4>
        <h4> {{ opp.Org_location }}</h4>
        <h4> {{ opp.Mission_statement }}</h4>
        <h4> {{ opp.Impact_space }}</h4>
        <h4> {{ opp.Un_goals }}</h4>
        <h4> {{ opp.Projects }}</h4>
        <h4> {{ opp.Values }}</h4>
        <h4> {{ opp.Benefits }}</h4>
        <h4> {{ opp.Ethical_criteria }}</h4>
        <h4> {{ opp.Ethical_criteria_details }}</h4>
        <h4> {{ opp.Volunteer_location }}</h4>
        <h4> {{ opp.Volunteer_skills }}</h4>
        <h4> {{ opp.Skill_projects }}</h4>
        <h4> {{ opp.Language_requirements }}</h4>
        <h4> {{ opp.Volunteer_qualities }}</h4>
        <h4> {{ opp.Work_env }}</h4>
        <h4> {{ opp.Accomodations }}</h4>
        <h4> {{ opp.Cost }}</h4>
        <h4> {{ opp.Website }}</h4>
        <h4> {{ opp.Contact }}</h4>
        <h4> {{ opp.Photos }}</h4>
        <h4> {{ opp.Refer_others }}</h4>
        <h4> {{ opp.Additional_info }}</h4>
        <h4> {{ opp.Finding_volunteers }}</h4>
        <h4> {{ opp.Org_logo }}</h4>
        <h4> {{ opp.Program_duration }}</h4>
        <h4> {{ opp.Project_timeframe }}</h4>
    </div>
    <FooterMod />
</template>

