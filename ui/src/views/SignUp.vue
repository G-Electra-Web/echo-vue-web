<script setup lang="ts">
import NavBar from '../components/NavBar.vue'
import UserAuthForm from '../components/UserSignup.vue'
</script>

<template>
    <nav>
        <NavBar />
    </nav>
    <div class="flex flex-col items-center h-screen p-[15px] mx-auto justify-center w-full ">
        <UserAuthForm/>
    </div>

    <router-view />
</template>