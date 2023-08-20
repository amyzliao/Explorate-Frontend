<script>
// import HeaderModVolunteer from './HeaderModVolunteer.vue'
// import HeaderModNgo from './HeaderModNgo.vue'
// import HeaderModSignedOut from './HeaderModSignedOut.vue'
import HeaderMod from './header/HeaderMod.vue'
import FooterMod from './FooterMod.vue'

export default {
    name: 'HomePage',
    components: {
        HeaderMod,
        FooterMod
    }
}
</script>

<template>
    <HeaderMod/>
    <!-- block A:  -->
    <h1 class="mt-[44px] text-[50px] font-bold text-[#323232]">Building A Global Community</h1>
    <p class="text-[25px] font-medium text-[#DE5E56]">Through Empathy, Impact, and Cross-Cultural Impact</p>
    <!-- Counter thing -->
    <div class="inline-flex justify-center w-[784px] p-4 bg-[#CFCFCF] rounded-[30px] h-[137px] bg-opacity-[0.15] translate-y-5 shadow-2xl drop-shadow-md backdrop-blur-[20px]">
        <div class="grid grid-cols-3 gap-24">
            <div class="flex flex-col items-center justify-center">
                <dt class="mb-2 text-5xl font-semibold text-[#5D90CD]">123+</dt>
                <dd class="text-black text-[25px]">Volunteers</dd>
            </div>
            <div class="flex flex-col items-center justify-center">
                <dt class="mb-2 text-5xl font-semibold text-[#5D90CD]">23+</dt>
                <dd class="text-black text-[25px]">NGOs</dd>
            </div>
            <div class="flex flex-col items-center justify-center">
                <dt class="mb-2 text-5xl font-semibold text-[#5D90CD]">12+</dt>
                <dd class="text-black text-[25px]">Countries</dd>
            </div>
        </div>
    </div>

    <!-- Image gallery -->
    <div class="flex justify-center gap-[10px] mx-[77px] cursor-pointer">
        <figure class="group static">
            <img class="group rounded-[20px] justify-self-center group-hover:justify-self-center group-hover:w-[350px] transition-all duration-500 object-none w-[175px] h-[358px]" src="https://flowbite.s3.amazonaws.com/blocks/marketing-ui/content/content-gallery-3.png" alt="default1">
            <div class="absolute group group-hover:block group-hover:-translate-y-[66px] text-white hidden">
                <p class="font-bold text-[20px] flex ml-[22px]">Organization name</p>
                <p class="font-bold text-[15px] flex ml-[22px] underline">read more</p>
            </div>
        </figure>
        <figure class="group static">
            <img class="group rounded-[20px] justify-self-center group-hover:justify-self-center group-hover:w-[350px] transition-all duration-500 object-none w-[175px] h-[358px]" src="https://flowbite.s3.amazonaws.com/docs/gallery/square/image-2.jpg" alt="default2">
            <div class="absolute group group-hover:block group-hover:-translate-y-[66px] text-white hidden">
                <p class="font-bold text-[20px] flex ml-[22px]">Organization name</p>
                <p class="font-bold text-[15px] flex ml-[22px] underline">read more</p>
            </div>
        </figure>
        <div>
            <figure class="group static">
                <img class="group rounded-[20px] justify-self-center group-hover:justify-self-center group-hover:w-[350px] transition-all duration-500 object-none w-[175px] h-[358px]" src="https://flowbite.s3.amazonaws.com/docs/gallery/square/image-3.jpg" alt="default3">
                <div class="absolute group group-hover:block group-hover:-translate-y-[66px] text-white hidden">
                    <p class="font-bold text-[20px] flex ml-[22px]">Organization name</p>
                    <p class="font-bold text-[15px] flex ml-[22px] underline">read more</p>
                </div>
            </figure>
        </div>
        <div>
            <figure class="group static">
                <img class="group rounded-[20px] justify-self-center group-hover:justify-self-center group-hover:w-[350px] transition-all duration-500 object-none w-[175px] h-[358px]" src="https://flowbite.s3.amazonaws.com/docs/gallery/square/image-4.jpg" alt="default4">
                <div class="absolute group group-hover:block group-hover:-translate-y-[66px] text-white hidden">
                    <p class="font-bold text-[20px] flex ml-[22px]">Organization name</p>
                    <p class="font-bold text-[15px] flex ml-[22px] underline">read more</p>
                </div>
            </figure>
        </div>
        <div>
            <figure class="group static">
                <img class="group rounded-[20px] justify-self-center group-hover:justify-self-center group-hover:w-[350px] transition-all duration-500 object-none w-[175px] h-[358px]" src="https://flowbite.s3.amazonaws.com/docs/gallery/square/image-5.jpg" alt="default5">
                <div class="absolute group group-hover:block group-hover:-translate-y-[66px] text-white hidden">
                    <p class="font-bold text-[20px] flex ml-[22px]">Organization name</p>
                    <p class="font-bold text-[15px] flex ml-[22px] underline">read more</p>
                </div>
            </figure>
        </div>
        <div>
            <figure class="group static">
                <img class="group rounded-[20px] justify-self-center group-hover:justify-self-center group-hover:w-[350px] transition-all duration-500 object-none w-[175px] h-[358px]" src="https://flowbite.s3.amazonaws.com/docs/gallery/square/image-6.jpg" alt="default6">
                <div class="absolute group group-hover:block group-hover:-translate-y-[66px] text-white hidden">
                    <p class="font-bold text-[20px] flex ml-[22px]">Organization name</p>
                    <p class="font-bold text-[15px] flex ml-[22px] underline">read more</p>
                </div>
            </figure>
        </div>
        <div>
            <figure class="group static">
                <img class="group rounded-[20px] justify-self-center group-hover:justify-self-center group-hover:w-[350px] transition-all duration-500 object-none w-[175px] h-[358px]" src="https://flowbite.s3.amazonaws.com/docs/gallery/square/image-3.jpg" alt="default7">
                <div class="absolute group group-hover:block group-hover:-translate-y-[66px] text-white hidden">
                    <p class="font-bold text-[20px] flex ml-[22px]">Organization name</p>
                    <p class="font-bold text-[15px] flex ml-[22px] underline">read more</p>
                </div>
            </figure>
        </div>
    </div>
    <!-- block B: two cards, one for volunteers and other for ngos -->
    <div class="flex items-center mx-auto mt-[73px] justify-center gap-5 my-6">
        <div class="p-6 h-[238px] bg-white">
            <a href="#">
                <h5 class="mb-2 text-2xl font-bold text-[#E7513F] text-[45px] h-[56px]">I'm a Volunteer</h5>
            </a>
            <p class="mb-3 text-black text-[25px] h-[86px]">Explore programs that resonate.<br>
                Navigate with a global community.</p>
            <a href="#"
                class="inline-flex w-[420px] items-center justify-center px-[18px] py-3.5 text-[20px] font-bold text-center text-[#E7513F] bg-white border border-[#E7513F] rounded-[15px]">
                Find a Program
                <svg class="w-3.5 h-3.5 ml-2" aria-hidden="true" xmlns="http://www.w3.org/2000/svg" fill="none"
                    viewBox="0 0 14 10">
                    <path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"
                        d="M1 5h12m0 0L9 1m4 4L9 9" />
                </svg>
            </a>
        </div>
        <div class="p-6 h-[238px] bg-white">
            <a href="#">
                <h5 class="mb-2 text-2xl font-bold text-[#E7513F] text-[45px] h-[56px]">I'm a NGO</h5>
            </a>
            <p class="mb-3 text-black text-[25px] h-[86px]">Volunteer Coordination Streamlined.</p>
            <a href="#"
                class="inline-flex w-[420px] items-center justify-center px-[18px] py-3.5 text-[20px] font-bold text-center text-[#E7513F] bg-white border border-[#E7513F] rounded-[15px]">
                Partner with Us
                <svg class="w-3.5 h-3.5 ml-2" aria-hidden="true" xmlns="http://www.w3.org/2000/svg" fill="none"
                    viewBox="0 0 14 10">
                    <path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"
                        d="M1 5h12m0 0L9 1m4 4L9 9" />
                </svg>
            </a>
        </div>
    </div>
    <!-- block C: image -->
    <img class="relative inline-flex" src="../assets/WhyExplorate.png" alt="Why Explorate">
    <!-- Sponsors -->
    <h1 class="font-medium text-4xl mt-12">Proudly Supported by</h1>
    <div class="flex justify-center mt-5 mb-12">
        <img src="../assets/sponsors.png" alt="sponsors">
    </div>
    <!-- Interested -->
    <div class="relative inline-flex justify-center items-center">
        <img class="inline" src="../assets/high-angle-people-world-map.png" alt="background interested">
        <!-- Middle section -->
        <div class="p-14 absolute bg-[#CFCFCF] drop-shadow-md bg-opacity-[0.15] grid grid-cols-1 justify-items-center w-[784px] h-[510px] rounded-[30px] backdrop-blur-[20px]">
            <h1 class="text-[32px] text-center text-black">Interested?
                <img src="../assets/pen_under_digital.png" alt="pen under digital">
            </h1>
            <h1 class="font-bold text-[45px] text-black w-[431px]">Join Explorate's Community Now!</h1>
            <div class="flex gap-[55px]">
                <a href="#" class="bg-[#E7513F] h-[56px] text-white border-2 border-[#E7513F] rounded-[15px] p-4 flex items-center">
                    <p class="font-medium text-xl px-4">Sign Up for Free</p>
                </a>
                <a href="#" class="bg-[#E7513F] h-[56px] text-white border-2 border-[#E7513F] rounded-[15px] p-4 flex items-center">
                    <p class="font-medium text-xl px-4">Log in</p>
                    </a>
            </div>
        </div>
    </div>
    <FooterMod />
</template>

