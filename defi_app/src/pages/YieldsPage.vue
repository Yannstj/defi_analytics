<script setup lang="ts">
import { onMounted, ref } from 'vue'
// On importe les sous-composants de la table Shadcn
import {
  Table,
  TableBody,
  TableCaption,
  TableCell,
  TableHead,
  TableHeader,
  TableRow,
} from '@/components/ui/table'

// Type simple pour tes tokens (TypeScript)
interface Token {
  id: string
  name: string
  symbol: string
}

const tokens = ref<Token[]>([])
const error = ref<string | null>(null)

const fetchTokens = async () => {
  try {
    const response = await fetch('http://localhost:8080/tokens')
    if (!response.ok) throw new Error(`Erreur HTTP : ${response.status}`)
    const data = await response.json()
    // Si ton API renvoie un objet direct, assure-toi que c'est bien un tableau
    tokens.value = Array.isArray(data) ? data : [data]
  } catch (err) {
    console.log(err)
    error.value = 'Impossible de joindre le backend.'
  }
}

onMounted(() => {
  fetchTokens()
})
</script>

<template>
  <main class="p-8">
    <h1 class="text-3xl font-bold mb-6 tracking-tight">DeFi Dashboard</h1>

    <div v-if="error" class="p-4 mb-4 text-red-700 bg-red-50 rounded-lg border border-red-200">
      {{ error }}
    </div>

    <div v-else class="rounded-md border">
      <Table>
        <TableCaption>Liste des tokens récupérés depuis l'API Go.</TableCaption>
        <TableHeader>
          <TableRow>
            <TableHead class="w-[100px]">Symbole</TableHead>
            <TableHead>Nom du Token</TableHead>
            <TableHead class="text-right">ID</TableHead>
          </TableRow>
        </TableHeader>
        <TableBody>
          <TableRow v-for="token in tokens" :key="token.id">
            <TableCell class="font-bold text-blue-600">{{ token.symbol }}</TableCell>
            <TableCell>{{ token.name }}</TableCell>
            <TableCell class="text-right font-mono text-xs text-gray-500 italic">
              {{ token.id }}
            </TableCell>
          </TableRow>
        </TableBody>
      </Table>

      <div v-if="tokens.length === 0" class="p-8 text-center text-gray-500">
        Aucun token trouvé...
      </div>
    </div>
  </main>
</template>
