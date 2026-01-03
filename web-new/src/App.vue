<script setup>
import { ref, onMounted } from 'vue'
import Dashboard from './views/Dashboard.vue'
import MyRules from './views/MyRules.vue'
import Mall from './views/Mall.vue'
import Nodes from './views/Nodes.vue'
import Profile from './views/Profile.vue'

const activeTab = ref('dashboard')
const user = ref({ balance: 0, used: 0, total: 0 })

const menuItems = [
  { id: 'dashboard', label: '中心控制台', icon: '🏠' },
  { id: 'rules', label: '转发链路', icon: '🔗' },
  { id: 'mall', label: '流量集市', icon: '🛍️' },
  { id: 'nodes', label: '资源调度', icon: '🖥️' },
  { id: 'profile', label: '账户中心', icon: '👤' }
]

onMounted(() => {
  // Fetch initial data
})
</script>

<template>
  <div class="app-container">
    <aside class="sidebar">
      <div class="logo-area mb-12">
        <h1 class="text-2xl font-black tracking-tighter italic">ULTRA<span class="text-primary-500">FORWARD</span></h1>
        <p class="text-[10px] text-gray-500 font-bold uppercase tracking-widest mt-1">High-Speed Stealth Relay</p>
      </div>

      <nav class="flex-1">
        <div 
          v-for="item in menuItems" 
          :key="item.id"
          class="nav-item"
          :class="{ active: activeTab === item.id }"
          @click="activeTab = item.id"
        >
          <span class="text-lg">{{ item.icon }}</span>
          <span>{{ item.label }}</span>
        </div>
      </nav>

      <div class="user-preview glass-card p-4 mt-auto">
        <div class="text-[10px] uppercase font-bold text-gray-500 mb-2">我的账号</div>
        <div class="text-xl font-bold">￥{{ user.balance.toFixed(2) }}</div>
        <div class="mt-2 h-1 bg-white/10 rounded-full overflow-hidden">
          <div class="h-full bg-primary-500" :style="{ width: (user.used/user.total*100) + '%' }"></div>
        </div>
      </div>
    </aside>

    <main class="content">
      <header class="flex justify-between items-center mb-12 animate-fade">
        <div>
          <h2 class="text-3xl font-bold gradient-text">{{ menuItems.find(i => i.id === activeTab).label }}</h2>
          <p class="text-gray-500 text-sm mt-1">控制、观测与扩展您的隐身分流网络</p>
        </div>
        <div class="flex gap-4">
           <button class="btn-primary">新建链路</button>
        </div>
      </header>

      <div class="view-container animate-fade">
        <Dashboard v-if="activeTab === 'dashboard'" />
        <MyRules v-if="activeTab === 'rules'" />
        <Mall v-if="activeTab === 'mall'" />
        <Nodes v-if="activeTab === 'nodes'" />
        <Profile v-if="activeTab === 'profile'" />
      </div>
    </main>
  </div>
</template>

<style scoped>
.app-container {
  display: flex;
}
.logo-area h1 {
  font-family: 'JetBrains Mono', monospace;
}
</style>
