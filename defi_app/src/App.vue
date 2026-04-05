<script setup lang="ts">
import { onMounted, ref } from 'vue'

// On crée une variable réactive pour stocker nos tokens
const tokens = ref([])
const error = ref<string | null>(null)

// La fonction qui va interroger ton backend Go
const fetchTokens = async () => {
  try {
    const response = await fetch('http://localhost:8080/tokens')

    if (!response.ok) {
      throw new Error(`Erreur HTTP : ${response.status}`)
    }

    const data = await response.json()
    tokens.value = data
    console.log('Données reçues du backend :', data)
  } catch (err) {
    error.value = 'Impossible de joindre le backend.'
    console.error('Erreur lors du fetch :', err)
  }
}

// Appelé automatiquement quand le composant est affiché
onMounted(() => {
  fetchTokens()
})
</script>

<template>
  <main class="p-8">
    <h1 class="text-2xl font-bold mb-4">Dashboard DeFi</h1>

    <div v-if="error" class="text-red-500 bg-red-100 p-4 rounded">
      {{ error }}
    </div>

    <div v-else class="bg-gray-100 p-4 rounded shadow-inner">
      <h2 class="font-semibold mb-2">Tokens (API Go) :</h2>
      <pre v-if="tokens.length">{{ tokens }}</pre>
      <p v-else>Chargement des données...</p>
    </div>
  </main>
</template>
