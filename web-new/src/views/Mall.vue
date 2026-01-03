<template>
  <div class="space-y-12 animate-fade-in">
    <!-- 顶部钱包余额 -->
    <div class="glass-card bg-gradient-to-br from-primary/10 via-transparent to-accent/5 p-8 flex flex-col md:flex-row justify-between items-center relative overflow-hidden">
       <div class="relative z-10">
          <h3 class="text-3xl font-black italic tracking-tighter mb-2 italic uppercase">资产钱包充值</h3>
          <p class="text-gray-500 text-sm font-medium">账户余额：<span class="text-white">￥1258.50</span> • 支持全自动实时结算</p>
       </div>
       <div class="flex gap-4 relative z-10 w-full md:w-auto">
          <div class="flex-1 md:w-48 bg-black/40 border border-white/5 rounded-2xl flex items-center px-4">
             <span class="text-gray-500 mr-2">￥</span>
             <input type="number" placeholder="0.00" class="bg-transparent border-none outline-none text-white font-bold w-full" />
          </div>
          <button class="btn-aura px-8">立即充值</button>
       </div>
       <!-- 装饰 -->
       <div class="absolute -right-10 -bottom-10 w-40 h-40 bg-accent/5 blur-[80px] rounded-full"></div>
    </div>

    <!-- 套餐矩阵 -->
    <div class="grid grid-cols-1 md:grid-cols-3 gap-8">
       <div v-for="plan in plans" :key="plan.name" 
            class="glass-card p-10 flex flex-col relative group"
            :class="{ 'ring-2 ring-primary/40': plan.featured }">
          
          <!-- 推荐勋章 -->
          <div v-if="plan.featured" class="absolute -top-4 left-1/2 -translate-x-1/2 bg-primary text-white text-[10px] font-black uppercase px-4 py-1.5 rounded-full shadow-lg shadow-primary/40">
             Best Value • 最受欢迎
          </div>

          <div class="mb-8">
             <div class="text-[10px] font-black text-gray-500 uppercase tracking-widest mb-4">{{ plan.tag }}</div>
             <h3 class="text-3xl font-black text-white italic tracking-tighter mb-2">{{ plan.name }}</h3>
             <div class="flex items-baseline gap-1">
                <span class="text-4xl font-black text-white">￥{{ plan.price }}</span>
                <span class="text-xs text-gray-500 font-bold uppercase">/ 月</span>
             </div>
          </div>

          <div class="space-y-6 flex-1">
             <div v-for="feat in plan.features" :key="feat" class="flex items-center gap-3">
                <div class="w-5 h-5 rounded-full bg-primary/10 flex items-center justify-center">
                   <svg class="w-3 h-3 text-primary" fill="none" stroke="currentColor" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="3" d="M5 13l4 4L19 7"></path></svg>
                </div>
                <span class="text-sm font-medium text-gray-300">{{ feat }}</span>
             </div>
          </div>

          <div class="mt-12">
             <button class="w-full py-4 rounded-2xl font-black uppercase tracking-widest text-xs transition-all border border-white/10 group-hover:bg-white group-hover:text-black hover:scale-[1.02]">
                立即开通计划
             </button>
          </div>
       </div>
    </div>

    <!-- 补充说明 -->
    <div class="text-center text-gray-500 text-[10px] font-bold uppercase tracking-[3px]">
       ULTRAFORWARD COMMERCIAL PLATFORM • ENTERPRISE SLA GUARANTEED
    </div>
  </div>
</template>

<script setup>
const plans = [
   { 
      tag: '入门版', 
      name: 'Starter', 
      price: 19.9, 
      featured: false,
      features: ['100GB 极速流量', '5 条活跃链路', '隐身协议全支持', '基础 BGP 线路']
   },
   { 
      tag: '进阶版', 
      name: 'Premium', 
      price: 49.9, 
      featured: true,
      features: ['500GB 极速流量', '20 条活跃链路', '独立 IP 解析', '中港/中日专线解锁', '24/7 优先支持']
   },
   { 
      tag: '商业版', 
      name: 'Ultimate', 
      price: 99.9, 
      featured: false,
      features: ['2TB 极速流量', '无限活跃链路', '独享出口节点', '全球全线路解锁', '专属运维服务']
   }
]
</script>
