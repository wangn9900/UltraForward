<script setup>
import { ref, onMounted } from 'vue'

const isAdmin = ref(true) // 实际应从 User 数据判断
const plans = ref([])
const newPlan = ref({ name: '', price: 0, traffic: 0, rules: 5, duration_days: 30, description: '' })

const fetchPlans = async () => {
  try {
    const res = await fetch('/api/v1/plans')
    plans.value = await res.json()
  } catch (err) {
    console.error('Fetch plans failed', err)
  }
}

const addPlan = async () => {
  await fetch('/api/v1/admin/plans', {
    method: 'POST',
    headers: { 'Content-Type': 'application/json' },
    body: JSON.stringify(newPlan.value)
  })
  fetchPlans()
}

const deletePlan = async (id) => {
  await fetch(`/api/v1/admin/plans/${id}`, { method: 'DELETE' })
  fetchPlans()
}

onMounted(fetchPlans)
</script>

<template>
  <div class="space-y-12 animate-fade-in">
    <!-- 管理员控制台 -->
    <div v-if="isAdmin" class="glass-card border-primary/20 bg-primary/5 p-8">
       <h3 class="text-xl font-black italic uppercase mb-6 text-primary">商城内务管理 (ADMIN ONLY)</h3>
       <div class="grid grid-cols-1 md:grid-cols-4 gap-4">
          <input v-model="newPlan.name" placeholder="套餐名称" class="bg-black/40 border border-white/10 rounded-xl px-4 py-2 text-sm" />
          <input v-model.number="newPlan.price" type="number" placeholder="价格 (CNY)" class="bg-black/40 border border-white/10 rounded-xl px-4 py-2 text-sm" />
          <input v-model.number="newPlan.traffic" type="number" placeholder="流量 (GB)" class="bg-black/40 border border-white/10 rounded-xl px-4 py-2 text-sm" />
          <button @click="addPlan" class="btn-aura py-2 text-xs">发布新套餐</button>
       </div>
    </div>
    
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
       <div v-for="plan in plans" :key="plan.ID" 
            class="glass-card p-10 flex flex-col relative group"
            :class="{ 'ring-2 ring-primary/40': plan.price > 40 }">
          
          <button v-if="isAdmin" @click="deletePlan(plan.ID)" class="absolute top-4 right-4 text-red-500 text-xs hover:font-bold">删除</button>

          <div class="mb-8">
             <div class="text-[10px] font-black text-gray-500 uppercase tracking-widest mb-4">商业计划</div>
             <h3 class="text-3xl font-black text-white italic tracking-tighter mb-2">{{ plan.name }}</h3>
             <div class="flex items-baseline gap-1">
                <span class="text-4xl font-black text-white">￥{{ plan.price }}</span>
                <span class="text-xs text-gray-500 font-bold uppercase">/ 月</span>
             </div>
          </div>

          <div class="space-y-6 flex-1">
             <div class="flex items-center gap-3">
                <div class="w-5 h-5 rounded-full bg-primary/10 flex items-center justify-center">
                   <svg class="w-3 h-3 text-primary" fill="none" stroke="currentColor" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="3" d="M5 13l4 4L19 7"></path></svg>
                </div>
                <span class="text-sm font-medium text-gray-300">{{ plan.traffic }}GB 极速流量</span>
             </div>
             <div class="flex items-center gap-3">
                <div class="w-5 h-5 rounded-full bg-primary/10 flex items-center justify-center">
                   <svg class="w-3 h-3 text-primary" fill="none" stroke="currentColor" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="3" d="M5 13l4 4L19 7"></path></svg>
                </div>
                <span class="text-sm font-medium text-gray-300">{{ plan.rules }} 条活跃链路</span>
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
