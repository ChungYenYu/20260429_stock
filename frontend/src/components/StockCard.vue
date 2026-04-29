<script setup>
defineProps({
  stock: {
    type: Object,
    required: true
  }
})
</script>

<template>
  <div class="bg-apple-gray rounded-apple-lg p-6 shadow-sm hover:shadow-md transition-shadow">
    <div class="flex flex-col gap-2">
      <div class="flex justify-between items-start">
        <span class="text-sm font-medium text-gray-500 uppercase tracking-wider">{{ stock.code }}</span>
      </div>
      <h3 class="text-xl font-semibold text-near-black">{{ stock.name }}</h3>
      <div class="mt-4">
        <span class="text-3xl font-bold text-near-black">
          ${{ stock.price.toLocaleString() }}
        </span>
      </div>
    </div>
  </div>
</template>
