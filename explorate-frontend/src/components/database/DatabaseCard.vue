<script>
export default {
    name: 'DatabaseCard',
    props: ['name', 'location', 'mission', 'projects', 'link']
}
</script>

<template>
    <div class="bg-white border border-gray-200 rounded-xl shadow">
        <!-- Title section -->
        <a href="#">
            <div class="h-[72px] flex items-center gap-4 py-3 pl-4 pr-1">
                <!-- profile pic circle, just placeholder with first initial -->
                <div class="h-10 w-10 bg-blue-400 rounded-full flex justify-center items-center text-white font-bold text-2xl">
                    {{ name[0] }}
                </div>
                <!-- NGO name -->
                <h1 class="font-bold text-[#1D1B20]">{{ name }}</h1>
            </div>
        </a>
        <!-- image -->
        <a href="#">
            <img class="rounded-t-lg" src="../../assets/picking_up_trash.jpg" alt="" />
        </a>
        <!-- info section -->
        <div class="p-2">
            <p class="mb-3 text-start line-clamp-3 h-[60px] text-sm">{{ mission }}</p>
            <p class="text-start text-[14px] font-semibold h-6 my-1">{{ location }}</p>
            <h2 class="text-start font-medium underline h-7">Projects</h2>
            <p class="mb-3 font-normal text-start line-clamp-2 text-sm">{{ projects }}</p>
            <div class="flex justify-end gap-2 m-4">
                <div class="border rounded-[100px] py-2.5 px-6 font-medium text-[#5D90CD] border-[#79747E] text-sm">Save</div>
                <div class="rounded-[100px] py-2.5 px-6 font-medium text-white bg-[#5D90CD] text-sm">Learn more</div>
            </div>
        </div>
    </div>
</template>