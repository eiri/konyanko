<script setup>
import Navigation from './components/Navigation.vue'
import Main from './components/Main.vue'
</script>

<template>
  <header class="flex justify-center px-4 py-2 h-12 w-full bg-gray-900 text-gray-300">
    <Navigation />
  </header>
  <main class="container mx-auto">
    <Main />
  </main>
</template>

<style scoped>
</style>
