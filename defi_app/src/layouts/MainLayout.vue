<script setup lang="ts">
import AppHeader from '@/components/AppHeader.vue'
</script>
<template>
  <div class="flex h-screen">
    <div class="flex-1 flex flex-col">
      <AppHeader />
      <main class="p-8">
        <router-view />
      </main>
    </div>
  </div>
</template>
