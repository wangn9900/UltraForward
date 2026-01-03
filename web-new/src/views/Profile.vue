<template>
  <div class="max-w-4xl space-y-8">
    <div class="glass-card flex items-center gap-8">
       <div class="w-24 h-24 rounded-full bg-gradient-to-br from-primary-500 to-accent flex items-center justify-center text-4xl shadow-xl">
          👤
       </div>
       <div>
          <h3 class="text-3xl font-black italic">ULTRA_CLIENT_001</h3>
          <p class="text-gray-500 mt-1">注册于 2025年12月24日 | 商用旗舰计划 (Premium)</p>
       </div>
    </div>

    <div class="grid grid-cols-1 md:grid-cols-2 gap-8">
       <div class="glass-card">
          <h4 class="font-bold mb-6">安全设置</h4>
          <div class="space-y-6 text-sm">
             <div class="flex justify-between items-center border-b border-white/5 pb-4">
                <span class="text-gray-400">登入密码</span>
                <button class="text-primary-400 font-bold">修改</button>
             </div>
             <div class="flex justify-between items-center border-b border-white/5 pb-4">
                <span class="text-gray-400">两步验证 (2FA)</span>
                <span class="px-2 py-0.5 bg-red-500/20 text-red-400 rounded-full text-[10px] font-bold">未开启</span>
             </div>
             <div class="flex justify-between items-center">
                <span class="text-gray-400">API 访问令牌</span>
                <code class="text-xs bg-black/40 px-2 py-1 rounded">uf_live_******</code>
             </div>
          </div>
       </div>

       <div class="glass-card">
          <h4 class="font-bold mb-6">财务历史</h4>
          <div class="space-y-4">
             <div v-for="i in 3" :key="i" class="flex justify-between items-center text-sm">
                <div>
                   <div class="font-bold">Premium 计划续费</div>
                   <div class="text-[10px] text-gray-500">2026-01-01 12:30</div>
                </div>
                <div class="text-emerald-400 font-bold">-￥49.90</div>
             </div>
          </div>
       </div>
    </div>
  </div>
</template>
