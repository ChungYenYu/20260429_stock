<script setup>
import { ref, onMounted } from 'vue'
import axios from 'axios'

const news = ref([])
const loading = ref(true)
const error = ref(null)

const fetchNews = async () => {
  try {
    loading.value = true
    // API endpoint as per SPEC
    const response = await axios.get('/api/news')
    news.value = response.data.map(item => ({
      ...item,
      impactedStocksList: JSON.parse(item.impacted_stocks || '[]')
    }))
  } catch (err) {
    console.error('Failed to fetch news:', err)
    error.value = 'Failed to load news insights. Please ensure the backend is running.'
  } finally {
    loading.value = false
  }
}

onMounted(() => {
  fetchNews()
})

const getImpactColor = (impact) => {
  switch (impact?.toLowerCase()) {
    case 'positive': return 'bg-green-100 text-green-700 border-green-200'
    case 'negative': return 'bg-red-100 text-red-700 border-red-200'
    case 'neutral': return 'bg-gray-100 text-gray-700 border-gray-200'
    default: return 'bg-gray-100 text-gray-700 border-gray-200'
  }
}

const formatDate = (dateString) => {
  const date = new Date(dateString)
  return date.toLocaleDateString('en-US', {
    month: 'short',
    day: 'numeric',
    year: 'numeric',
    hour: '2-digit',
    minute: '2-digit'
  })
}
</script>

<template>
  <div class="news-board">
    <div v-if="loading" class="flex flex-col items-center justify-center py-20">
      <div class="w-10 h-10 border-4 border-apple-blue border-t-transparent rounded-full animate-spin"></div>
      <p class="mt-4 text-gray-500 font-medium">Curating latest market insights...</p>
    </div>

    <div v-else-if="error" class="p-6 bg-red-50 rounded-apple-lg border border-red-100 text-red-600 text-center">
      <p class="font-semibold">{{ error }}</p>
      <button @click="fetchNews" class="mt-4 text-sm text-red-700 underline font-medium">Try again</button>
    </div>

    <div v-else class="space-y-8">
      <div v-for="(item, index) in news" :key="index" 
           class="group bg-apple-gray rounded-apple-lg p-8 transition-all duration-300 hover:shadow-xl border border-transparent hover:border-gray-200">
        <div class="flex flex-col lg:flex-row gap-8">
          <!-- Content Section -->
          <div class="flex-1">
            <div class="flex items-center gap-3 mb-4">
              <span class="text-xs font-bold text-gray-400 uppercase tracking-widest">{{ item.source }}</span>
              <span class="w-1 h-1 bg-gray-300 rounded-full"></span>
              <span class="text-xs text-gray-400">{{ formatDate(item.published_at) }}</span>
            </div>
            
            <a :href="item.url" target="_blank" class="block">
              <h3 class="text-2xl font-bold text-near-black mb-4 group-hover:text-apple-blue transition-colors leading-tight">
                {{ item.title }}
              </h3>
            </a>
            
            <p class="text-gray-600 text-lg leading-relaxed mb-6">
              {{ item.summary }}
            </p>
          </div>

          <!-- Impact Analysis Section -->
          <div class="lg:w-80 flex flex-col bg-white rounded-apple-lg p-6 shadow-sm border border-gray-100">
            <div class="flex items-center justify-between mb-4">
              <h4 class="text-xs font-black uppercase tracking-tighter text-gray-400">Impact Analysis</h4>
              <div :class="['px-3 py-1 rounded-full text-xs font-bold border uppercase', getImpactColor(item.impact_analysis)]">
                {{ item.impact_analysis }}
              </div>
            </div>
            
            <div class="mb-6">
              <p class="text-sm text-near-black leading-snug">
                {{ item.impact_analysis === 'Positive' ? 'AI detects strong bullish signals for related sectors based on this report.' : 
                   item.impact_analysis === 'Negative' ? 'AI suggests caution as this news might introduce market volatility.' :
                   'Market consensus remains neutral following this announcement.' }}
              </p>
            </div>

            <div v-if="item.impactedStocksList.length > 0">
              <span class="block text-[10px] font-bold text-gray-400 uppercase mb-3 tracking-widest">Related Stocks</span>
              <div class="flex flex-wrap gap-2">
                <button v-for="stock in item.impactedStocksList" :key="stock"
                   class="px-3 py-1.5 rounded-apple-md bg-apple-blue/5 text-apple-blue text-xs font-bold hover:bg-apple-blue hover:text-white transition-all">
                  ${{ stock }}
                </button>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<style scoped>
.news-board {
  max-width: 100%;
}
</style>
