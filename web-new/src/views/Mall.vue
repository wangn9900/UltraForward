<template>
  <div class="space-y-12">
    <!-- Recharge Section -->
    <div class="glass-card bg-gradient-to-r from-primary-500/10 to-transparent flex justify-between items-center">
       <div>
          <h3 class="text-xl font-bold mb-1">账户余额充值</h3>
          <p class="text-gray-500 text-sm">支持 支付宝 / 微信 / USDT 实时到账</p>
       </div>
       <div class="flex gap-4">
          <input type="number" placeholder="输入金额" class="bg-black/40 border border-white/10 rounded-xl px-4 py-2 w-32 focus:outline-none focus:border-primary-500" />
          <button class="btn-primary">立即充值</button>
       </div>
    </div>

    <!-- Plans Grid -->
    <div class="grid grid-cols-1 md:grid-cols-3 gap-8">
       <div v-for="plan in plans" :key="plan.name" class="glass-card hover:border-primary-500/50 group">
          <div class="text-xs font-bold text-gray-500 uppercase tracking-widest mb-4">{{ plan.tag }}</div>
          <h3 class="text-2xl font-bold mb-2">{{ plan.name }}</h3>
          <div class="text-3xl font-black mb-6">￥{{ plan.price }} <span class="text-sm font-normal text-gray-500">/ 月</span></div>
          
          <ul class="space-y-4 mb-8 text-sm text-gray-300">
             <li class="flex items-center gap-2">✨ {{ plan.traffic }} 计费流量</li>
             <li class="flex items-center gap-2">✨ {{ plan.nodes }} 节点全解锁</li>
             <li class="flex items-center gap-2">✨ {{ plan.rules }} 条转发链路</li>
             <li class="flex items-center gap-2">✨ 企业级 SLA 保证</li>
          </ul>

          <button class="w-full py-3 rounded-xl border border-primary-500/20 group-hover:bg-primary-500 group-hover:text-white transition-all font-bold">立即订购</button>
       </div>
    </div>
  </div>
</template>

<script setup>
const plans = [
   { tag: '入门级', name: 'Starter', price: 19.9, traffic: '100GB', nodes: '基础', rules: 5 },
   { tag: '推荐', name: 'Premium', price: 49.9, traffic: '500GB', nodes: '全部', rules: 20 },
   { tag: '商用级', name: 'Ultimate', price: 99.9, traffic: '2TB', nodes: '全部', rules: '无限' }
]
</script>
