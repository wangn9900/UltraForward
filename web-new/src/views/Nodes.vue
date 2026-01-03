<template>
  <div class="space-y-8 animate-fade-in">
    <div class="flex justify-between items-end border-b border-white/5 pb-8">
       <div>
          <h3 class="text-3xl font-black italic tracking-tighter uppercase">资源调度矩阵</h3>
          <p class="text-gray-500 text-xs mt-1 font-bold uppercase tracking-widest">Global Infrastructure Cluster</p>
       </div>
       <button class="btn-aura">
          + 注册新服务器
       </button>
    </div>

    <!-- 服务器矩阵 -->
    <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
       <div v-for="node in nodes" :key="node.id" class="glass-card group p-6 flex gap-6 relative overflow-hidden">
          <!-- 地区浮窗 -->
          <div class="absolute -right-8 -top-8 w-24 h-24 bg-primary/5 rounded-full blur-2xl group-hover:bg-primary/20 transition-all"></div>
          
          <div class="w-20 h-20 rounded-3xl bg-white/5 flex flex-col items-center justify-center border border-white/10 group-hover:border-primary/50 transition-all font-black text-3xl italic">
             {{ node.region }}
          </div>

          <div class="flex-1 space-y-4">
             <div class="flex justify-between items-start">
                <div>
                   <h4 class="text-xl font-bold tracking-tight text-white">{{ node.name }}</h4>
                   <div class="text-[10px] font-mono text-gray-500 mt-1">{{ node.ip }}</div>
                </div>
                <div class="flex items-center gap-2">
                   <div class="w-2 h-2 rounded-full bg-emerald-500 shadow-[0_0_10px_#10b981]"></div>
                   <span class="text-[10px] font-black uppercase text-emerald-500">Active</span>
                </div>
             </div>

             <div class="grid grid-cols-3 gap-4">
                <div v-for="metric in ['CPU', 'MEM', 'NET']" :key="metric" class="space-y-1">
                   <div class="text-[10px] font-black text-gray-500 uppercase">{{ metric }}</div>
                   <div class="h-1 bg-white/5 rounded-full overflow-hidden">
                      <div class="h-full bg-primary" :style="{ width: Math.random()*60+20 + '%' }"></div>
                   </div>
                </div>
             </div>

             <div class="flex justify-between items-center pt-2">
                <div class="flex gap-4">
                   <div class="text-[10px] font-bold text-gray-400">持续运行: <span class="text-white">12d 14h</span></div>
                   <div class="text-[10px] font-bold text-gray-400">负载: <span class="text-white">0.45</span></div>
                </div>
                <div class="flex gap-2">
                   <button class="p-2 bg-white/5 rounded-lg hover:bg-white/10 transition-all">
                      <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8 9l3 3-3 3m5 0h3"></path></svg>
                   </button>
                   <button class="p-2 bg-red-500/10 text-red-500 rounded-lg hover:bg-red-500/20 transition-all">
                      <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13 7a4 4 0 11-8 0 4 4 0 018 0zM9 14a6 6 0 00-6 6v1h12v-1a6 6 0 00-6-6zM21 12h-6"></path></svg>
                   </button>
                </div>
             </div>
          </div>
       </div>
    </div>
  </div>
</template>

<script setup>
const nodes = [
  { id: 1, name: 'CN-SH-ALIYUN-02', ip: '120.55.x.x', region: 'SH' },
  { id: 2, name: 'HK-AZURE-B1MS', ip: '20.124.x.x', region: 'HK' },
  { id: 3, name: 'JP-LINODE-TYO', ip: '139.162.x.x', region: 'JP' },
  { id: 4, name: 'US-LAX-GIA-01', ip: '154.21.x.x', region: 'LA' }
]
</script>
