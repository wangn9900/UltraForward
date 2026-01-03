<template>
  <div class="space-y-8 animate-fade-in">
    <!-- 核心指标 -->
    <div class="grid grid-cols-1 md:grid-cols-4 gap-6">
      <div v-for="s in stats" :key="s.label" class="glass-card group overflow-hidden relative">
        <div class="relative z-10">
          <div class="text-[10px] uppercase font-black text-primary tracking-widest mb-2">{{ s.label }}</div>
          <div class="text-3xl font-black text-white italic tracking-tighter">{{ s.value }}</div>
          <div class="mt-4 flex items-center gap-2">
             <span class="text-[10px] bg-emerald-500/10 text-emerald-400 px-2 py-0.5 rounded-md font-bold">{{ s.trend }}</span>
             <span class="text-[10px] text-gray-500">vs 昨日</span>
          </div>
        </div>
        <!-- 装饰背景 -->
        <div class="absolute -right-2 -top-2 text-4xl opacity-5 group-hover:opacity-10 transition-opacity rotate-12">{{ s.icon }}</div>
      </div>
    </div>

    <!-- 全球链路动态监测 -->
    <div class="glass-card min-h-[400px] flex flex-col items-center justify-center relative overflow-hidden">
      <!-- 动态波纹底图 -->
      <div class="absolute inset-0 opacity-20">
         <div class="absolute top-1/2 left-1/2 -translate-x-1/2 -translate-y-1/2 w-[800px] h-[800px] border border-primary/20 rounded-full animate-ping-slow"></div>
         <div class="absolute top-1/2 left-1/2 -translate-x-1/2 -translate-y-1/2 w-[500px] h-[500px] border border-primary/10 rounded-full animate-ping-slow-delay"></div>
      </div>

      <div class="relative z-10 text-center">
         <div class="mb-6 relative inline-block">
            <div class="text-6xl filter drop-shadow(0 0 20px var(--primary-glow))">🌐</div>
            <!-- 雷达扫描线 -->
            <div class="absolute top-1/2 left-1/2 -translate-x-1/2 -translate-y-1/2 w-48 h-48 border-t-2 border-primary/40 rounded-full animate-spin-slow"></div>
         </div>
         <h3 class="text-2xl font-black italic tracking-tight mb-2">全球隐身中转骨干网</h3>
         <p class="text-gray-400 text-sm max-w-md">当前有 <span class="text-white font-bold">12</span> 条专线正在运行，Stealth-Pass 加密协议已全面覆盖。</p>
         
         <div class="mt-10 flex gap-4 text-xs font-bold uppercase tracking-widest text-gray-500">
            <div class="flex items-center gap-2"><span class="w-2 h-2 rounded-full bg-primary animate-pulse"></span> CN-SHENZHEN</div>
            <div class="text-white">→</div>
            <div class="flex items-center gap-2"><span class="w-2 h-2 rounded-full bg-accent animate-pulse"></span> HK-AZURE</div>
            <div class="text-white">→</div>
            <div class="flex items-center gap-2"><span class="w-2 h-2 rounded-full bg-blue-500 animate-pulse"></span> US-LAX</div>
         </div>
      </div>
    </div>

    <!-- 实时系统日志 -->
    <div class="grid grid-cols-1 md:grid-cols-2 gap-8">
       <div class="glass-card">
          <div class="flex justify-between items-center mb-6">
             <h4 class="font-bold text-lg italic tracking-tight uppercase">近期安全事件</h4>
             <span class="text-[10px] text-primary">查看全部</span>
          </div>
          <div class="space-y-4">
             <div v-for="i in 4" :key="i" class="flex items-start gap-4 p-3 hover:bg-white/5 rounded-2xl transition-all">
                <div class="mt-1 w-2 h-2 rounded-full" :class="i===1 ? 'bg-primary' : 'bg-gray-700'"></div>
                <div class="flex-1">
                   <div class="text-sm font-bold text-gray-200">成功对核心节点 <span class="text-white">HK-01</span> 下发了安全补丁</div>
                   <div class="text-[10px] text-gray-500 mt-1 uppercase font-bold tracking-widest">System Engine • 12:45</div>
                </div>
             </div>
          </div>
       </div>

       <div class="glass-card">
          <div class="flex justify-between items-center mb-6">
             <h4 class="font-bold text-lg italic tracking-tight uppercase">流量排行 (24h)</h4>
          </div>
          <div class="space-y-6">
             <div v-for="n in ['深港专线', '日本精品', '美国专线']" :key="n" class="space-y-2">
                <div class="flex justify-between text-xs font-bold">
                   <span class="text-gray-300">{{ n }}</span>
                   <span class="text-primary">{{ (Math.random()*100).toFixed(1) }} GB</span>
                </div>
                <div class="h-1 bg-white/5 rounded-full overflow-hidden">
                   <div class="h-full bg-gradient-to-r from-primary to-accent" :style="{ width: Math.random()*80+20 + '%' }"></div>
                </div>
             </div>
          </div>
       </div>
    </div>
  </div>
</template>

<script setup>
const stats = [
  { label: '累计流量', value: '4.2 TB', trend: '+12%', icon: '📊' },
  { label: '活跃用户', value: '1,256', trend: '+5%', icon: '👥' },
  { label: '在线节点', value: '18', trend: '稳定', icon: '📡' },
  { label: '平均可用率', value: '99.9%', trend: '极优', icon: '🛡️' }
]
</script>

<style scoped>
.animate-ping-slow { animation: ping 3s cubic-bezier(0, 0, 0.2, 1) infinite; }
.animate-ping-slow-delay { animation: ping 3s cubic-bezier(0, 0, 0.2, 1) infinite; animation-delay: 1.5s; }
.animate-spin-slow { animation: spin 4s linear infinite; }

@keyframes ping {
  75%, 100% { transform: translate(-50%, -50%) scale(1.5); opacity: 0; }
}
@keyframes spin {
  from { transform: translate(-50%, -50%) rotate(0deg); }
  to { transform: translate(-50%, -50%) rotate(360deg); }
}
.animate-fade-in {
  animation: fadeIn 0.8s cubic-bezier(0.16, 1, 0.3, 1);
}
@keyframes fadeIn {
  from { opacity: 0; transform: translateY(20px); }
  to { opacity: 1; transform: translateY(0); }
}
</style>
