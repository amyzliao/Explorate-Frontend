<script>
// import HeaderModVolunteer from '../HeaderModVolunteer.vue'
// import HeaderModNgo from '../HeaderModNgo.vue'
// import HeaderModSignedOut from './HeaderModSignedOut.vue'
import HeaderMod from '../header/HeaderMod.vue'
import FooterMod from '../FooterMod.vue'

export default {
    name: 'VolunteerProfilePage',
    components: {
        FooterMod,
        HeaderMod
    },
    computed: {
        view() {
            let user = localStorage.getItem("user-info");
            if (user) {
                const name = JSON.parse(user).account;
                return name == 'volunteer' ? 'HeaderModVolunteer' : 'HeaderModNgo'
            } else {
                return 'HeaderMod'
            }
        }
    }
}
</script>

<template>
    <HeaderMod/>
    <!-- block A:  -->
    <h1
        class="mt-44 mb-4 text-4xl font-extrabold leading-none tracking-tight text-gray-900 md:text-5xl lg:text-6xl dark:text-white">
        Volunteer Profile Page</h1>
    <p class="mb-6 text-lg font-normal text-gray-500 lg:text-xl sm:px-16 xl:px-48 dark:text-gray-400">Here at Flowbite we
        focus on markets where technology, innovation, and capital can unlock long-term value and drive economic growth.</p>
    <a href="#"
        class="inline-flex items-center justify-center px-5 py-3 text-base font-medium text-center text-white bg-blue-700 rounded-lg hover:bg-blue-800 focus:ring-4 focus:ring-blue-300 dark:focus:ring-blue-900">
        Learn more
        <svg class="w-3.5 h-3.5 ml-2" aria-hidden="true" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 14 10">
            <path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"
                d="M1 5h12m0 0L9 1m4 4L9 9" />
        </svg>
    </a>
    <!-- block B: two cards, one for volunteers and other for ngos -->
    <div class="flex items-center mx-auto w-[50%] mt-10">
        <div class="max-w-sm p-6 bg-white border border-gray-200 rounded-lg shadow dark:bg-gray-800 dark:border-gray-700">
            <a href="#">
                <h5 class="mb-2 text-2xl font-bold tracking-tight text-gray-900 dark:text-white">I'm a volunteer</h5>
            </a>
            <p class="mb-3 font-normal text-gray-700 dark:text-gray-400">Here are the biggest enterprise technology
                acquisitions of 2021 so far, in reverse chronological order.</p>
            <a href="#"
                class="inline-flex items-center px-3 py-2 text-sm font-medium text-center text-white bg-blue-700 rounded-lg hover:bg-blue-800 focus:ring-4 focus:outline-none focus:ring-blue-300 dark:bg-blue-600 dark:hover:bg-blue-700 dark:focus:ring-blue-800">
                Read more
                <svg class="w-3.5 h-3.5 ml-2" aria-hidden="true" xmlns="http://www.w3.org/2000/svg" fill="none"
                    viewBox="0 0 14 10">
                    <path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"
                        d="M1 5h12m0 0L9 1m4 4L9 9" />
                </svg>
            </a>
        </div>
        <div class="max-w-sm p-6 bg-white border border-gray-200 rounded-lg shadow dark:bg-gray-800 dark:border-gray-700">
            <a href="#">
                <h5 class="mb-2 text-2xl font-bold tracking-tight text-gray-900 dark:text-white">I'm a NGO</h5>
            </a>
            <p class="mb-3 font-normal text-gray-700 dark:text-gray-400">Here are the biggest enterprise technology
                acquisitions of 2021 so far, in reverse chronological order.</p>
            <a href="#"
                class="inline-flex items-center px-3 py-2 text-sm font-medium text-center text-white bg-blue-700 rounded-lg hover:bg-blue-800 focus:ring-4 focus:outline-none focus:ring-blue-300 dark:bg-blue-600 dark:hover:bg-blue-700 dark:focus:ring-blue-800">
                Read more
                <svg class="w-3.5 h-3.5 ml-2" aria-hidden="true" xmlns="http://www.w3.org/2000/svg" fill="none"
                    viewBox="0 0 14 10">
                    <path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"
                        d="M1 5h12m0 0L9 1m4 4L9 9" />
                </svg>
            </a>
        </div>
    </div>
    <!-- block C: image -->
    <img class="h-auto max-w-lg rounded-lg mx-auto mt-16" src="../../assets/picking_up_trash.jpg" alt="image description">
    <!-- block D: quote -->
    <figure class="max-w-screen-md mx-auto text-center mt-16 mb-16">
        <svg class="w-10 h-10 mx-auto mb-3 text-gray-400 dark:text-gray-600" aria-hidden="true"
            xmlns="http://www.w3.org/2000/svg" fill="currentColor" viewBox="0 0 18 14">
            <path
                d="M6 0H2a2 2 0 0 0-2 2v4a2 2 0 0 0 2 2h4v1a3 3 0 0 1-3 3H2a1 1 0 0 0 0 2h1a5.006 5.006 0 0 0 5-5V2a2 2 0 0 0-2-2Zm10 0h-4a2 2 0 0 0-2 2v4a2 2 0 0 0 2 2h4v1a3 3 0 0 1-3 3h-1a1 1 0 0 0 0 2h1a5.006 5.006 0 0 0 5-5V2a2 2 0 0 0-2-2Z" />
        </svg>
        <blockquote>
            <p class="text-2xl italic font-medium text-gray-900 dark:text-white">"Flowbite is just awesome. It contains tons
                of predesigned components and pages starting from login screen to complex dashboard. Perfect choice for your
                next SaaS application."</p>
        </blockquote>
        <figcaption class="flex items-center justify-center mt-6 space-x-3">
            <img class="w-6 h-6 rounded-full"
                src="https://flowbite.s3.amazonaws.com/blocks/marketing-ui/avatars/michael-gouch.png" alt="profile picture">
            <div class="flex items-center divide-x-2 divide-gray-500 dark:divide-gray-700">
                <cite class="pr-3 font-medium text-gray-900 dark:text-white">Micheal Gough</cite>
                <cite class="pl-3 text-sm text-gray-500 dark:text-gray-400">CEO at Google</cite>
            </div>
        </figcaption>
    </figure>

    <FooterMod />
</template>
