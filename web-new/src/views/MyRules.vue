<template>
  <div class="space-y-6">
    <div v-if="rules.length === 0" class="glass-card py-20 text-center">
       <div class="text-4xl mb-4">🛰️</div>
       <p class="text-gray-400">目前没有任何活跃链路，点击右上角新建</p>
    </div>

    <div v-for="rule in rules" :key="rule.id" class="glass-card flex flex-col md:flex-row gap-6 items-center">
      <!-- Status Icon -->
      <div class="w-12 h-12 rounded-2xl bg-primary-500/10 flex items-center justify-center text-xl">
        {{ rule.status ? '✅' : '⏸️' }}
      </div>

      <!-- Info -->
      <div class="flex-1">
        <div class="flex items-center gap-2 mb-1">
           <h3 class="font-bold text-lg">{{ rule.name }}</h3>
           <span class="text-[10px] px-2 py-0.5 bg-primary-500/20 text-primary-300 rounded-full font-bold">x{{ rule.line_multiplier }} 计费</span>
        </div>
        <div class="text-xs text-gray-500 font-mono">
           {{ rule.listen_port }} (入口) -> {{ rule.exit_name }} -> {{ rule.target }}
        </div>
      </div>

      <!-- Traffic Stats -->
      <div class="flex gap-8 px-8 border-l border-r border-white/10 text-center">
         <div>
            <div class="text-[10px] uppercase text-gray-500 mb-1">上行</div>
            <div class="font-mono text-sm">{{ formatBytes(rule.upload) }}</div>
         </div>
         <div>
            <div class="text-[10px] uppercase text-gray-500 mb-1">下行</div>
            <div class="font-mono text-sm">{{ formatBytes(rule.download) }}</div>
         </div>
      </div>

      <!-- Actions -->
      <div class="flex gap-2">
         <button class="p-2 hover:bg-white/5 rounded-xl transition-colors">⚙️</button>
         <button class="p-2 hover:bg-white/5 rounded-xl transition-colors text-red-500">🗑️</button>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref } from 'vue'

const rules = ref([
   { id: 1, name: '深港专线 - SSR/VLESS', listen_port: 44021, exit_name: 'HK-Azure-01', target: '1.2.3.4:443', line_multiplier: 3.5, upload: 1024*1024*450, download: 1024*1024*1024*3.2, status: true },
   { id: 2, name: '日本精品 - 企业独享', listen_port: 44022, exit_name: 'JP-Linode-02', target: '5.6.7.8:443', line_multiplier: 1.0, upload: 1024*1024*12, download: 1024*1024*45, status: true }
])

const formatBytes = (bytes) => {
  if (bytes === 0) return '0 B'
  const k = 1024
  const sizes = ['B', 'KB', 'MB', 'GB', 'TB']
  const i = Math.floor(Math.log(bytes) / Math.log(k))
  return parseFloat((bytes / Math.pow(k, i)).toFixed(2)) + ' ' + sizes[i]
}
</script>
