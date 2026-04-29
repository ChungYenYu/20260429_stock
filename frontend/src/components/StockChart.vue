<script setup>
import { ref, onMounted, onUnmounted, watch } from 'vue';
import { createChart } from 'lightweight-charts';
import axios from 'axios';

const props = defineProps({
  code: {
    type: String,
    required: true
  }
});

const chartContainer = ref(null);
let chart = null;
let candlestickSeries = null;

const initChart = () => {
  if (!chartContainer.value) return;

  chart = createChart(chartContainer.value, {
    layout: {
      background: { color: 'transparent' },
      textColor: '#FFFFFF',
    },
    grid: {
      vertLines: { visible: false },
      horzLines: { visible: false },
    },
    rightPriceScale: {
      borderVisible: false,
    },
    timeScale: {
      borderVisible: false,
    },
    handleScroll: {
      mouseWheel: true,
      pressedMouseMove: true,
    },
    handleScale: {
      axisPressedMouseMove: true,
      mouseWheel: true,
      pinch: true,
    },
  });

  candlestickSeries = chart.addCandlestickSeries({
    upColor: '#26a69a',
    downColor: '#ef5350',
    borderVisible: false,
    wickUpColor: '#26a69a',
    wickDownColor: '#ef5350',
  });

  window.addEventListener('resize', handleResize);
};

const handleResize = () => {
  if (chart && chartContainer.value) {
    chart.applyOptions({
      width: chartContainer.value.clientWidth,
      height: chartContainer.value.clientHeight,
    });
  }
};

const fetchData = async () => {
  try {
    const response = await axios.get(`/api/stock/${props.code}/history`);
    const data = response.data.map(item => ({
      time: item.time,
      open: parseFloat(item.open),
      high: parseFloat(item.high),
      low: parseFloat(item.low),
      close: parseFloat(item.close),
    }));
    
    // Sort data by time to ensure it's in order for lightweight-charts
    data.sort((a, b) => a.time - b.time);
    
    if (candlestickSeries) {
      candlestickSeries.setData(data);
    }
  } catch (error) {
    console.error('Error fetching stock history:', error);
  }
};

onMounted(() => {
  initChart();
  fetchData();
  handleResize();
});

onUnmounted(() => {
  if (chart) {
    chart.remove();
  }
  window.removeEventListener('resize', handleResize);
});

watch(() => props.code, () => {
  fetchData();
});
</script>

<template>
  <div class="w-full h-[400px] bg-gray-900 rounded-xl p-4">
    <div ref="chartContainer" class="w-full h-full"></div>
  </div>
</template>

<style scoped>
</style>
