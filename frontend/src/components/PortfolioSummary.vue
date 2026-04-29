<script setup>
import { ref, onMounted, computed } from 'vue';
import axios from 'axios';

const holdings = ref([
  { code: '0056', name: 'Yuanta High Div', sheets: 20, cost: 31.5, currentPrice: 0 },
  { code: '00919', name: 'Capital High Div', sheets: 30, cost: 20.0, currentPrice: 0 }
]);

const fetchPrices = async () => {
  for (const holding of holdings.value) {
    try {
      const response = await axios.get(`http://localhost:8080/api/stock/${holding.code}`);
      holding.currentPrice = response.data.price;
    } catch (error) {
      console.error(`Error fetching price for ${holding.code}:`, error);
      // Fallback prices for demo/safety
      if (holding.code === '0056') holding.currentPrice = 38.5;
      if (holding.code === '00919') holding.currentPrice = 24.2;
    }
  }
};

const totalCost = computed(() => {
  return holdings.value.reduce((acc, h) => acc + (h.sheets * 1000 * h.cost), 0);
});

const totalValue = computed(() => {
  return holdings.value.reduce((acc, h) => {
    const price = h.currentPrice || h.cost; // Use cost as fallback for calculation if not loaded
    return acc + (h.sheets * 1000 * price);
  }, 0);
});

const totalPnL = computed(() => totalValue.value - totalCost.value);
const totalPnLPercent = computed(() => totalCost.value === 0 ? 0 : (totalPnL.value / totalCost.value) * 100);

const formatCurrency = (val) => {
  return new Intl.NumberFormat('zh-TW', { style: 'currency', currency: 'TWD', minimumFractionDigits: 0 }).format(val);
};

const formatPercent = (val) => {
  return val.toFixed(2) + '%';
};

onMounted(() => {
  fetchPrices();
});
</script>

<template>
  <div class="max-w-7xl mx-auto py-12 px-4">
    <div class="grid grid-cols-1 lg:grid-cols-3 gap-8">
      <!-- Summary Card -->
      <div class="lg:col-span-1 glass-card p-8 flex flex-col justify-between overflow-hidden relative">
        <div class="absolute top-0 right-0 w-32 h-32 bg-apple-blue/5 rounded-full -mr-16 -mt-16 blur-3xl"></div>
        <div class="relative z-10">
          <h3 class="text-gray-500 text-sm font-semibold uppercase tracking-wider mb-2">Net Worth</h3>
          <div class="text-5xl font-bold text-near-black mb-6 tracking-tight">{{ formatCurrency(totalValue) }}</div>
          
          <div class="inline-flex items-center px-4 py-2 rounded-full" :class="totalPnL >= 0 ? 'bg-green-50 text-green-600' : 'bg-red-50 text-red-600'">
            <span class="text-2xl font-bold mr-2">{{ totalPnL >= 0 ? '↑' : '↓' }} {{ formatPercent(Math.abs(totalPnLPercent)) }}</span>
            <span class="font-medium">{{ formatCurrency(Math.abs(totalPnL)) }}</span>
          </div>
        </div>
        
        <div class="mt-12 space-y-4 relative z-10">
          <div class="flex justify-between items-center text-sm">
            <span class="text-gray-400 font-medium uppercase tracking-tight">Invested</span>
            <span class="text-near-black font-bold">{{ formatCurrency(totalCost) }}</span>
          </div>
          <div class="w-full bg-gray-100 h-1.5 rounded-full overflow-hidden">
            <div 
              class="h-full bg-apple-blue transition-all duration-1000"
              :style="{ width: Math.min(100, (totalCost / totalValue) * 100) + '%' }"
            ></div>
          </div>
        </div>
      </div>

      <!-- Moat Visualization / Holdings -->
      <div class="lg:col-span-2 space-y-6">
        <div v-for="holding in holdings" :key="holding.code" class="glass-card p-8 hover:shadow-xl transition-shadow duration-500">
          <div class="flex justify-between items-start mb-8">
            <div>
              <div class="flex items-center gap-3 mb-1">
                <span class="bg-near-black text-white text-[10px] font-bold px-2 py-0.5 rounded tracking-widest">{{ holding.code }}</span>
                <h4 class="text-2xl font-bold text-near-black tracking-tight">{{ holding.name }}</h4>
              </div>
              <p class="text-gray-400 font-medium">{{ holding.sheets }} Sheets • Cost {{ holding.cost }}</p>
            </div>
            <div class="text-right">
              <div class="text-2xl font-bold text-near-black tracking-tight">{{ formatCurrency(holding.currentPrice) }}</div>
              <div :class="['text-sm font-bold', holding.currentPrice >= holding.cost ? 'text-green-500' : 'text-red-500']">
                {{ holding.currentPrice >= holding.cost ? '+' : '' }}{{ ((holding.currentPrice - holding.cost) / holding.cost * 100).toFixed(2) }}%
              </div>
            </div>
          </div>
          
          <!-- Moat Visualization -->
          <div class="space-y-2">
            <div class="flex justify-between text-[10px] font-bold text-gray-400 uppercase tracking-widest">
              <span>Cost Moat</span>
              <span>{{ ((holding.currentPrice / holding.cost - 1) * 100).toFixed(1) }}% Safety</span>
            </div>
            <div class="relative h-12 bg-gray-50 rounded-2xl p-2 flex items-center">
              <!-- Background bar -->
              <div class="absolute inset-x-2 h-2 bg-gray-200 rounded-full"></div>
              
              <!-- Growth bar -->
              <div 
                v-if="holding.currentPrice > holding.cost"
                class="absolute h-2 bg-green-400 rounded-full transition-all duration-1000"
                :style="{ 
                  left: 'calc(8px + ' + (holding.cost / Math.max(holding.cost, holding.currentPrice) * 80) + '%)',
                  width: ((holding.currentPrice - holding.cost) / Math.max(holding.cost, holding.currentPrice) * 80) + '%'
                }"
              ></div>

              <!-- Cost point -->
              <div 
                class="absolute w-4 h-4 bg-white border-4 border-near-black rounded-full shadow-lg z-20 transition-all duration-1000"
                :style="{ left: 'calc(8px + ' + (holding.cost / Math.max(holding.cost, holding.currentPrice) * 80) + '%)' }"
              >
                <div class="absolute -bottom-6 left-1/2 -translate-x-1/2 text-[9px] font-bold text-near-black whitespace-nowrap">COST {{ holding.cost }}</div>
              </div>

              <!-- Current point -->
              <div 
                class="absolute w-6 h-6 bg-apple-blue rounded-full shadow-xl border-4 border-white z-30 transition-all duration-1000 flex items-center justify-center"
                :style="{ left: 'calc(8px + ' + (holding.currentPrice / Math.max(holding.cost, holding.currentPrice) * 80) + '%)' }"
              >
                <div class="absolute -top-8 left-1/2 -translate-x-1/2 bg-near-black text-white text-[9px] px-2 py-1 rounded-md whitespace-nowrap shadow-xl">
                  NOW {{ holding.currentPrice }}
                  <div class="absolute top-full left-1/2 -translate-x-1/2 border-4 border-transparent border-t-near-black"></div>
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<style scoped>
.glass-card {
  background-color: white;
  border: 1px solid #f3f4f6;
  border-radius: 2rem;
  box-shadow: 0 8px 30px rgba(0, 0, 0, 0.04);
  backdrop-filter: blur(8px);
}
</style>
