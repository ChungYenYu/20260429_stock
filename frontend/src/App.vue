<script setup>
import AppleNavbar from './components/AppleNavbar.vue'
import StockCard from './components/StockCard.vue'
import { ref } from 'vue'

const stocks = ref([
  { code: '0056', name: 'Yuanta High Div', price: 38.5 },
  { code: '00919', name: 'Capital High Div', price: 24.2 },
  { code: '2330', name: 'TSMC', price: 1045.0 }
])
</script>

<template>
  <div class="min-h-screen bg-white">
    <AppleNavbar />
    
    <main>
      <!-- Hero Section -->
      <section class="bg-black py-32 px-4">
        <div class="max-w-4xl mx-auto text-center">
          <h1 class="text-white text-5xl md:text-7xl font-bold tracking-tight mb-6">
            Invest in what matters.
          </h1>
          <p class="text-gray-400 text-xl md:text-2xl font-medium max-w-2xl mx-auto">
            Experience the future of portfolio management with precision and elegance.
          </p>
          <div class="mt-10 flex justify-center gap-4">
            <button class="apple-button-primary text-lg px-8 py-3">Get Started</button>
            <a href="#" class="apple-pill-link text-lg px-8 py-3 bg-transparent text-apple-blue border-none hover:underline">
              Learn more >
            </a>
          </div>
        </div>
      </section>

      <!-- Stock Grid Section -->
      <section class="py-24 px-4 bg-white">
        <div class="max-w-7xl mx-auto">
          <div class="flex justify-between items-end mb-12">
            <div>
              <h2 class="text-4xl font-bold text-near-black">Markets</h2>
              <p class="text-gray-500 mt-2 text-lg">Real-time performance of your favorite picks.</p>
            </div>
            <a href="#" class="text-apple-blue font-medium hover:underline">View all</a>
          </div>
          
          <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-8">
            <StockCard v-for="stock in stocks" :key="stock.code" :stock="stock" />
          </div>
        </div>
      </section>
    </main>

    <!-- Footer -->
    <footer class="bg-apple-gray py-12 px-4 border-t border-gray-200">
      <div class="max-w-7xl mx-auto text-sm text-gray-500">
        <p>© 2026 Stock Station. All rights reserved.</p>
      </div>
    </footer>
  </div>
</template>

<style>
/* Global styles are in style.css */
</style>
