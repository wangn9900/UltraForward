<script setup>
import { ref, onMounted } from 'vue'
import Dashboard from './views/Dashboard.vue'
import MyRules from './views/MyRules.vue'
import Mall from './views/Mall.vue'
import Nodes from './views/Nodes.vue'
import Profile from './views/Profile.vue'

import Auth from './views/Auth.vue'

const isLoggedIn = ref(false)
const activeTab = ref('dashboard')
const user = ref({ 
  id: 0,
  username: '',
  balance: 0, 
  used: 0, 
  total: 100,
  planName: '免费体验'
})

const handleAuthSuccess = (userData) => {
  isLoggedIn.value = true
  user.value = { ...user.value, ...userData }
}

const menuItems = [
  { id: 'dashboard', label: '中心控制台', icon: 'M3 12l2-2m0 0l7-7 7 7M5 10v10a1 1 0 001 1h3m10-11l2 2m-2-2v10a1 1 0 01-1 1h3m-6 0a1 1 0 001-1v-4a1 1 0 011-1h2a1 1 0 011 1v4a1 1 0 001 1m-6 0h6' },
  { id: 'rules', label: '转发链路', icon: 'M13 10V3L4 14h7v7l9-11h-7z' },
  { id: 'mall', label: '流量集市', icon: 'M16 11V7a4 4 0 00-8 0v4M5 9h14l1 12H4L5 9z' },
  { id: 'nodes', label: '资源调度', icon: 'M5 12h14M5 12a2 2 0 01-2-2V6a2 2 0 012-2h14a2 2 0 012 2v4a2 2 0 01-2 2M5 12a2 2 0 00-2 2v4a2 2 0 002 2h14a2 2 0 002-2v-4a2 2 0 00-2-2' },
  { id: 'profile', label: '账户中心', icon: 'M16 7a4 4 0 11-8 0 4 4 0 018 0zM12 14a7 7 0 00-7 7h14a7 7 0 00-7-7z' }
]
</script>

<template>
  <div class="app-container">
    <Auth v-if="!isLoggedIn" @auth-success="handleAuthSuccess" />
    
    <div v-else class="flex min-h-screen">
      <!-- 侧边栏 -->
      <aside class="w-[280px] bg-[rgba(10,10,12,0.8)] backdrop-blur-3xl border-r border-white/5 fixed inset-y-0 z-50 flex flex-col p-6">
        <div class="mb-12 px-2">
          <div class="flex items-center gap-3">
             <div class="w-10 h-10 bg-gradient-to-br from-primary to-accent rounded-xl flex items-center justify-center shadow-lg shadow-primary/20">
                <span class="text-white font-black text-xl italic">U</span>
             </div>
             <div>
                <h1 class="text-xl font-bold tracking-tight text-white leading-none">ULTRA<span class="text-primary">FORWARD</span></h1>
                <span class="text-[10px] text-gray-500 font-bold tracking-widest uppercase">Stealth Relay Pro</span>
             </div>
          </div>
        </div>

        <nav class="flex-1 space-y-2">
          <div 
            v-for="item in menuItems" 
            :key="item.id"
            class="nav-item group"
            :class="{ active: activeTab === item.id }"
            @click="activeTab = item.id"
          >
            <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
               <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" :d="item.icon"></path>
            </svg>
            <span>{{ item.label }}</span>
          </div>
        </nav>

        <!-- 底部钱包卡片 -->
        <div class="mt-auto glass-card p-5 border-white/5 relative overflow-hidden">
          <div class="relative z-10">
            <div class="flex justify-between items-center mb-3">
              <span class="text-[10px] font-bold text-gray-500 uppercase">我的资产</span>
              <span class="text-[10px] px-2 py-0.5 bg-primary/20 text-primary-300 rounded-full">正常</span>
            </div>
            <div class="text-2xl font-black italic">￥{{ user.balance.toFixed(2) }}</div>
            <div class="mt-4">
               <div class="flex justify-between text-[10px] text-gray-400 mb-1">
                  <span>套餐用量</span>
                  <span>{{ user.used }}/{{ user.total }} GB</span>
               </div>
               <div class="h-1.5 bg-white/5 rounded-full overflow-hidden">
                  <div class="h-full bg-gradient-to-r from-primary to-accent" :style="{ width: (user.used/user.total*100) + '%' }"></div>
               </div>
            </div>
          </div>
          <!-- 背景装饰 -->
          <div class="absolute -right-4 -bottom-4 w-24 h-24 bg-primary/10 blur-2xl rounded-full"></div>
        </div>
      </aside>

      <!-- 主体内容 -->
      <main class="flex-1 ml-[280px] p-10 relative">
        <header class="flex justify-between items-end mb-12">
          <div>
            <h2 class="text-4xl font-extrabold gradient-text tracking-tighter">{{ menuItems.find(i => i.id === activeTab).label }}</h2>
            <p class="text-gray-400 mt-2 font-medium">欢迎回来，{{ user.username }}。正在管理您的极前全球中转网络。</p>
          </div>
          <div class="flex items-center gap-4">
             <div class="flex -space-x-2">
                <div v-for="i in 3" :key="i" class="w-8 h-8 rounded-full border-2 border-[var(--bg-main)] bg-gray-800 flex items-center justify-center text-[10px] font-bold">
                   {{ ['H', 'S', 'J'][i-1] }}
                </div>
             </div>
             <button class="btn-aura">
               🚀 快速新建链路
             </button>
          </div>
        </header>

        <section class="page-view">
          <component :is="{
            dashboard: Dashboard,
            rules: MyRules,
            mall: Mall,
            nodes: Nodes,
            profile: Profile
          }[activeTab]" />
        </section>
      </main>
    </div>
  </div>
</template>

<style>
@import './index.css';
</style>
