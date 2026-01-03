<template>
  <div class="space-y-8">
    <!-- Quick Stats -->
    <div class="grid grid-cols-1 md:grid-cols-4 gap-6">
      <div v-for="s in stats" :key="s.label" class="glass-card">
        <div class="text-[10px] uppercase font-bold text-gray-400 mb-1">{{ s.label }}</div>
        <div class="text-2xl font-bold">{{ s.value }}</div>
        <div class="text-[10px] mt-2 text-primary-400">{{ s.trend }}</div>
      </div>
    </div>

    <!-- Visual Map / Relay Graph Placeholder -->
    <div class="glass-card h-80 flex items-center justify-center relative overflow-hidden group">
      <div class="absolute inset-0 bg-gradient-to-br from-primary-500/10 to-transparent"></div>
      <div class="text-center relative z-10">
        <div class="text-4xl mb-4">🌍</div>
        <h3 class="text-xl font-bold mb-2">全球链路流向图</h3>
        <p class="text-gray-500 text-sm max-w-xs mx-auto">实时观测流量从入口节点通过隐身隧道到达落地的流动轨迹</p>
      </div>
      <!-- Decorative lines -->
      <div class="absolute w-full h-px bg-gradient-to-r from-transparent via-primary-500/20 to-transparent top-1/2 -translate-y-1/2"></div>
      <div class="absolute h-full w-px bg-gradient-to-b from-transparent via-primary-500/20 to-transparent left-1/2 -translate-x-1/2"></div>
    </div>

    <!-- Recent Activity -->
    <div class="glass-card">
       <h3 class="font-bold mb-6">近期系统事件</h3>
       <div class="space-y-4">
          <div v-for="i in 3" :key="i" class="flex items-center gap-4 text-sm border-b border-white/5 pb-4 last:border-0 last:pb-0">
             <div class="w-2 h-2 rounded-full bg-emerald-500"></div>
             <div class="flex-1 text-gray-300">成功在 <span class="text-white">SH-CN-02</span> 部署了新的 Stealth-Pass 实例</div>
             <div class="text-gray-500 text-xs">2 分钟前</div>
          </div>
       </div>
    </div>
  </div>
</template>

<script setup>
const stats = [
  { label: '活跃链路', value: '12', trend: '+2 较昨日' },
  { label: '今日流量', value: '156.4 GB', trend: '实时结算中' },
  { label: '平均延迟', value: '42ms', trend: '-5ms 优化中' },
  { label: '在线节点', value: '8', trend: '所有节点健康' }
]
</script>
