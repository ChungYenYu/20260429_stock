<script setup>
import { ref, onMounted, onUnmounted } from 'vue';
import axios from 'axios';

const searchQuery = ref('');
const results = ref([]);
const isDropdownVisible = ref(false);
const searchContainer = ref(null);
let debounceTimer = null;

const emit = defineEmits(['show-portfolio', 'show-home']);

const handleSearch = () => {
  if (debounceTimer) clearTimeout(debounceTimer);
  
  if (!searchQuery.value.trim()) {
    results.value = [];
    isDropdownVisible.value = false;
    return;
  }

  debounceTimer = setTimeout(async () => {
    try {
      const response = await axios.get(`http://localhost:8080/api/stocks/search?q=${searchQuery.value}`);
      results.value = response.data;
      isDropdownVisible.value = results.value.length > 0;
    } catch (error) {
      console.error('Search error:', error);
      results.value = [];
    }
  }, 300);
};

const selectStock = (stock) => {
  console.log('Stock info:', stock);
  isDropdownVisible.value = false;
  searchQuery.value = '';
};

const handleClickOutside = (event) => {
  if (searchContainer.value && !searchContainer.value.contains(event.target)) {
    isDropdownVisible.value = false;
  }
};

onMounted(() => {
  document.addEventListener('click', handleClickOutside);
});

onUnmounted(() => {
  document.removeEventListener('click', handleClickOutside);
});
</script>

<template>
  <nav class="glass-nav">
    <div class="max-w-7xl mx-auto w-full flex justify-between items-center h-full">
      <div class="flex items-center flex-shrink-0">
        <span class="text-white font-bold text-xl tracking-tight">Stock Station</span>
      </div>
      
      <div class="hidden lg:flex items-center space-x-8 mx-6">
        <a href="#" @click.prevent="emit('show-home')" class="text-white/80 hover:text-white text-sm font-medium transition-colors">Home</a>
        <a href="#" @click.prevent="emit('show-portfolio')" class="text-white/80 hover:text-white text-sm font-medium transition-colors">Portfolio</a>
        <a href="#" class="text-white/80 hover:text-white text-sm font-medium transition-colors">Markets</a>
      </div>

      <div class="flex-1 flex justify-end items-center space-x-4">
        <div ref="searchContainer" class="relative w-full max-w-[240px]">
          <div class="relative group">
            <input 
              v-model="searchQuery"
              @input="handleSearch"
              @focus="isDropdownVisible = results.length > 0"
              type="text" 
              placeholder="Search" 
              class="w-full bg-white/10 text-white placeholder-white/40 border-none rounded-apple-pill px-4 py-1 text-sm focus:ring-1 focus:ring-apple-blue focus:bg-white/20 transition-all outline-none"
            />
            <div 
              v-if="isDropdownVisible && results.length > 0" 
              class="absolute top-full mt-2 w-full bg-black/90 backdrop-blur-md rounded-apple-lg border border-white/10 shadow-2xl overflow-hidden z-50 py-1"
            >
              <ul>
                <li 
                  v-for="stock in results" 
                  :key="stock.code"
                  @click="selectStock(stock)"
                  class="px-4 py-2 hover:bg-white/10 cursor-pointer flex flex-col"
                >
                  <span class="text-white text-sm font-medium">{{ stock.name }}</span>
                  <span class="text-white/40 text-xs font-mono uppercase">{{ stock.code }}</span>
                </li>
              </ul>
            </div>
          </div>
        </div>

        <div class="flex items-center flex-shrink-0">
          <button class="apple-button-primary text-sm whitespace-nowrap">Sign In</button>
        </div>
      </div>
    </div>
  </nav>
</template>

<style scoped>
/* Any component-specific styles if needed */
</style>
