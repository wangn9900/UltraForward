<template>
  <div class="max-w-5xl space-y-8 animate-fade-in">
    <!-- 个人名片 -->
    <div class="glass-card p-12 bg-gradient-to-br from-primary/20 via-transparent to-accent/10 flex flex-col md:flex-row items-center gap-12 relative overflow-hidden">
       <!-- 头像 -->
       <div class="relative">
          <div class="w-32 h-32 rounded-full bg-white/5 border-4 border-primary/30 flex items-center justify-center text-5xl shadow-2xl relative z-10">
             👤
          </div>
          <div class="absolute inset-0 bg-primary blur-3xl opacity-20 animate-pulse"></div>
       </div>

       <div class="flex-1 text-center md:text-left relative z-10">
          <div class="flex items-center justify-center md:justify-start gap-4 mb-2">
             <h3 class="text-4xl font-black italic tracking-tighter uppercase">ULTRA_ADMIN</h3>
             <span class="text-[10px] font-black px-3 py-1 bg-white text-black rounded-lg uppercase tracking-widest">VIP Level 4</span>
          </div>
          <p class="text-gray-400 font-medium italic">账户 UID: UF-20251210-001 • 加入于 2025年12月24日</p>
          
          <div class="mt-6 flex flex-wrap justify-center md:justify-start gap-4">
             <div class="bg-white/5 px-4 py-2 rounded-xl border border-white/5 text-xs font-bold uppercase tracking-widest">
                当前套餐: <span class="text-primary italic ml-2">Premium 旗舰版</span>
             </div>
             <div class="bg-white/5 px-4 py-2 rounded-xl border border-white/5 text-xs font-bold uppercase tracking-widest">
                到期时间: <span class="text-white italic ml-2">2026-06-24</span>
             </div>
          </div>
       </div>
    </div>

    <!-- 设置与历史 -->
    <div class="grid grid-cols-1 md:grid-cols-2 gap-8">
       <!-- 安全中心 -->
       <div class="glass-card p-8">
          <h4 class="text-xl font-bold italic tracking-tight uppercase mb-8 flex items-center gap-3">
             <span class="text-primary">🛡️</span> 安全审计与配置
          </h4>
          <div class="space-y-6">
             <div v-for="setting in settings" :key="setting.name" class="flex justify-between items-center p-4 hover:bg-white/5 rounded-2xl transition-all group">
                <div>
                   <div class="text-sm font-bold text-white group-hover:text-primary transition-colors">{{ setting.name }}</div>
                   <div class="text-[10px] text-gray-500 uppercase font-black mt-1">{{ setting.desc }}</div>
                </div>
                <button class="text-xs font-black uppercase text-gray-400 hover:text-white transition-colors">配置</button>
             </div>
          </div>
       </div>

       <!-- 财务历史 -->
       <div class="glass-card p-8 relative overflow-hidden">
          <h4 class="text-xl font-bold italic tracking-tight uppercase mb-8 flex items-center gap-3">
             <span class="text-accent">💳</span> 财务交易流水
          </h4>
          <div class="space-y-4 relative z-10">
             <div v-for="i in 5" :key="i" class="flex justify-between items-center p-3 rounded-xl hover:bg-white/5">
                <div>
                   <div class="text-sm font-bold">{{ i%2 === 0 ? '余额充值' : '套餐续费 (Premium)' }}</div>
                   <div class="text-[10px] text-gray-500 font-bold uppercase">2026-01-0{{ i }} • #TRX-{{ 1254+i }}</div>
                </div>
                <div class="text-sm font-black italic" :class="i%2 === 0 ? 'text-emerald-400' : 'text-red-400'">
                   {{ i%2 === 0 ? '+' : '-' }}￥{{ (49.9 * (i+1)).toFixed(2) }}
                </div>
             </div>
          </div>
          <!-- 装饰水印 -->
          <div class="absolute -right-10 bottom-0 text-[120px] font-black text-white/5 pointer-events-none italic select-none">Cash</div>
       </div>
    </div>
  </div>
</template>

<script setup>
const settings = [
  { name: '登录密码', desc: '最后修改于 15天前' },
  { name: '两步验证 (2FA)', desc: '强烈建议开启，保护资产安全' },
  { name: 'API Access Token', desc: '用于自动化集成与第三方控制' },
  { name: '登录日志库', desc: '查看最近 30 天的异常登录尝试' }
]
</script>
