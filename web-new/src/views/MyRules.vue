<template>
  <div class="space-y-8 animate-fade-in">
    <!-- 快捷操作栏 -->
    <div class="flex justify-between items-center glass-card p-6 bg-gradient-to-r from-primary/5 to-transparent">
       <div>
          <h3 class="text-xl font-bold italic tracking-tight uppercase">转发规则集群</h3>
          <p class="text-gray-500 text-xs mt-1 font-medium italic">ACTIVE RELAY RULES • CLUSTER_SECURE_01</p>
       </div>
       <div class="flex gap-4">
          <input type="text" placeholder="搜索链路名称..." class="bg-black/40 border border-white/5 rounded-xl px-5 py-2 text-sm focus:outline-none focus:border-primary transition-all w-64" />
       </div>
    </div>

    <!-- 链路规则阵列 -->
    <div class="grid grid-cols-1 gap-6">
       <div v-for="rule in rules" :key="rule.id" class="glass-card p-0 overflow-hidden group">
          <div class="flex flex-col md:flex-row">
             <!-- 状态标识区 -->
             <div class="w-2 md:w-3 bg-gradient-to-b from-primary to-accent" :class="{ 'opacity-20': !rule.status }"></div>
             
             <!-- 内容区 -->
             <div class="flex-1 p-8 flex flex-col md:flex-row items-center gap-8">
                <!-- 线路逻辑视觉 -->
                <div class="flex items-center gap-6">
                   <div class="text-center">
                      <div class="w-14 h-14 rounded-2xl bg-white/5 flex items-center justify-center text-2xl border border-white/10 group-hover:border-primary/40 transition-all font-bold italic">CN</div>
                      <div class="text-[10px] text-gray-500 mt-2 font-black uppercase">Entry</div>
                   </div>
                   
                   <div class="flex flex-col items-center gap-1 group-hover:scale-x-110 transition-transform origin-left">
                      <div class="text-[10px] text-primary font-black uppercase">Stealth-Pass</div>
                      <div class="flex gap-1">
                         <div v-for="i in 5" :key="i" class="w-4 h-1 bg-primary/20 rounded-full overflow-hidden">
                            <div class="h-full bg-primary animate-pulse" :style="{ animationDelay: i*0.2 + 's' }"></div>
                         </div>
                      </div>
                   </div>

                   <div class="text-center">
                      <div class="w-14 h-14 rounded-2xl bg-white/5 flex items-center justify-center text-2xl border border-white/10 group-hover:border-accent/40 transition-all font-bold italic">HK</div>
                      <div class="text-[10px] text-gray-500 mt-2 font-black uppercase">Exit</div>
                   </div>
                </div>

                <!-- 详细信息 -->
                <div class="flex-1">
                   <div class="flex items-center gap-3 mb-2">
                      <h4 class="text-xl font-bold italic tracking-tight">{{ rule.name }}</h4>
                      <span class="text-[10px] font-black px-2 py-0.5 bg-primary/10 text-primary border border-primary/20 rounded-lg">x{{ rule.multiplier }}</span>
                   </div>
                   <div class="flex items-center gap-4 text-xs font-mono text-gray-500">
                      <div class="flex items-center gap-1">📍 {{ rule.listen_port }}</div>
                      <div class="text-gray-700">|</div>
                      <div class="flex items-center gap-1 truncate max-w-[200px]">🎯 {{ rule.target }}</div>
                   </div>
                </div>

                <!-- 实时流量管路 (能量管视觉) -->
                <div class="w-full md:w-64 space-y-4">
                   <div class="space-y-1">
                      <div class="flex justify-between text-[10px] font-black text-gray-500 uppercase">
                         <span>Upload</span>
                         <span class="text-white italic">{{ formatBytes(rule.upload) }}</span>
                      </div>
                      <div class="h-1 bg-white/5 rounded-full overflow-hidden">
                         <div class="h-full bg-primary opacity-60" :style="{ width: '45%' }"></div>
                      </div>
                   </div>
                   <div class="space-y-1">
                      <div class="flex justify-between text-[10px] font-black text-gray-500 uppercase">
                         <span>Download</span>
                         <span class="text-white italic">{{ formatBytes(rule.download) }}</span>
                      </div>
                      <div class="h-1 bg-white/5 rounded-full overflow-hidden">
                         <div class="h-full bg-accent opacity-60" :style="{ width: '75%' }"></div>
                      </div>
                   </div>
                </div>

                <!-- 动作 -->
                <div class="flex gap-3">
                   <button class="w-12 h-12 rounded-2xl bg-white/5 hover:bg-white/10 flex items-center justify-center transition-all border border-white/5">
                      <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M10.325 4.317c.426-1.756 2.924-1.756 3.35 0a1.724 1.724 0 002.573 1.066c1.543-.94 3.31.826 2.37 2.37a1.724 1.724 0 001.065 2.572c1.756.426 1.756 2.924 0 3.35a1.724 1.724 0 00-1.066 2.573c.94 1.543-.826 3.31-2.37 2.37a1.724 1.724 0 00-2.572 1.065c-.426 1.756-2.924 1.756-3.35 0a1.724 1.724 0 00-2.573-1.066c-1.543.94-3.31-.826-2.37-2.37a1.724 1.724 0 00-1.065-2.572c-1.756-.426-1.756-2.924 0-3.35a1.724 1.724 0 001.066-2.573c-.94-1.543.826-3.31 2.37-2.37.996.608 2.296.07 2.572-1.065z"></path><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 12a3 3 0 11-6 0 3 3 0 016 0z"></path></svg>
                   </button>
                   <button class="w-12 h-12 rounded-2xl bg-red-500/10 hover:bg-red-500/20 text-red-500 flex items-center justify-center transition-all border border-red-500/10">
                      <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 7l-.867 12.142A2 2 0 0116.138 21H7.862a2 2 0 01-1.995-1.858L5 7m5 4v6m4-6v6m1-10V4a1 1 0 00-1-1h-4a1 1 0 00-1 1v3M4 7h16"></path></svg>
                   </button>
                </div>
             </div>
          </div>
       </div>
    </div>
  </div>
</template>

<script setup>
const rules = [
   { id: 1, name: '深港专线 - BGP 高防', listen_port: 8443, target: '1.2.3.4:443', multiplier: 3.5, upload: 1024*1024*890, download: 1024*1024*1024*5.2, status: true },
   { id: 2, name: '日本精品 - 独享家宽', listen_port: 8444, target: 'netflix.com:443', multiplier: 1.2, upload: 1024*1024*12, download: 1024*1024*45, status: true },
   { id: 3, name: '中德直连 - 德国数据中心', listen_port: 8445, target: 'fr-node.net:443', multiplier: 1.0, upload: 1024*1024*5, download: 1024*1024*2, status: false }
]

const formatBytes = (bytes) => {
  if (bytes === 0) return '0 B'
  const k = 1024
  const sizes = ['B', 'KB', 'MB', 'GB', 'TB']
  const i = Math.floor(Math.log(bytes) / Math.log(k))
  return parseFloat((bytes / Math.pow(k, i)).toFixed(2)) + ' ' + sizes[i]
}
</script>
