<script setup lang="ts">
import { Avatar, AvatarFallback, AvatarImage } from '@/components/ui/avatar'
import { Button } from '@/components/ui/button'
import { Input } from '@/components/ui/input'
</script>

<template>
  <header
    class="h-16 border-b bg-white/80 backdrop-blur-md sticky top-0 z-10 px-8 flex items-center justify-between"
  >
    <div class="flex items-center w-full max-w-sm relative">
      <span class="absolute left-3 text-gray-400 text-sm">🔍</span>
      <Input
        type="search"
        placeholder="Rechercher un protocole ou un token..."
        class="pl-10 rounded-full bg-gray-50 border-none focus-visible:ring-1 focus-visible:ring-blue-500"
      />
    </div>

    <div class="flex items-center gap-4">
      <Button
        variant="default"
        class="rounded-full bg-blue-600 hover:bg-blue-700 px-6 shadow-md shadow-blue-100"
      >
        Connect Wallet
      </Button>

      <div class="h-8 w-[1px] bg-gray-200 mx-2"></div>
      <Avatar
        class="h-9 w-9 border-2 border-white shadow-sm cursor-pointer hover:opacity-80 transition"
      >
        <AvatarImage src="https://github.com/shadcn.png" alt="@shadcn" />
        <AvatarFallback>DF</AvatarFallback>
      </Avatar>
    </div>
  </header>
</template>
