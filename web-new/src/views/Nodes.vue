<template>
  <div class="space-y-6">
    <div class="flex justify-between items-center bg-white/5 p-6 rounded-3xl border border-white/10">
       <div>
          <h3 class="font-bold">资源管理</h3>
          <p class="text-xs text-gray-500 mt-1">管理您的物理服务器与中转集群</p>
       </div>
       <button class="btn-primary">添加服务器</button>
    </div>

    <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
       <div v-for="node in nodes" :key="node.id" class="glass-card flex gap-6">
          <div class="w-16 h-16 rounded-2xl bg-white/5 flex items-center justify-center text-2xl">
             {{ node.region_icon }}
          </div>
          <div class="flex-1">
             <div class="flex justify-between items-start">
                <h4 class="font-bold">{{ node.name }}</h4>
                <span :class="node.status ? 'text-emerald-500' : 'text-red-500'" class="text-[10px] font-bold uppercase">Online</span>
             </div>
             <div class="text-xs text-gray-500 mt-1 font-mono">{{ node.ip }}</div>
             <div class="mt-4 flex gap-4 text-[10px] font-bold text-gray-400">
                <span>CPU: {{ node.cpu }}%</span>
                <span>MEM: {{ node.mem }}%</span>
                <span>UP: {{ node.uptime }}</span>
             </div>
          </div>
       </div>
    </div>
  </div>
</template>

<script setup>
const nodes = [
   { id: 1, name: 'CN-Shenzhen-Aliyun-01', ip: '120.24.x.x', region_icon: '🇨🇳', cpu: 12, mem: 45, uptime: '14d', status: true },
   { id: 2, name: 'HK-Azure-B1ms-01', ip: '20.124.x.x', region_icon: '🇭🇰', cpu: 5, mem: 30, uptime: '32d', status: true }
]
</script>
