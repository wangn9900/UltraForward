<script setup>
import { ref } from 'vue'

const isLogin = ref(true)
const emit = defineEmits(['auth-success'])

const form = ref({
  username: '',
  password: '',
  confirmPassword: ''
})

const loading = ref(false)
const error = ref('')

const handleSubmit = async () => {
  error.value = ''
  if (!isLogin.value && form.value.password !== form.value.confirmPassword) {
    error.value = '两次输入的密码不一致'
    return
  }

  loading.value = true
  try {
    const endpoint = isLogin.value ? '/api/v1/auth/login' : '/api/v1/auth/register'
    const res = await fetch(endpoint, {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify({
        username: form.value.username,
        password: form.value.password
      })
    })
    
    const data = await res.json()
    if (!res.ok) throw new Error(data.error || '操作失败')
    
    if (isLogin.value) {
      localStorage.setItem('uf_token', data.token)
      emit('auth-success', data.user)
    } else {
      // 注册后自动切换到登录
      isLogin.value = true
      alert('注册成功，请登录')
    }
  } catch (err) {
    error.value = err.message
  } finally {
    loading.value = false
  }
}
</script>

<template>
  <div class="fixed inset-0 z-[100] flex items-center justify-center bg-[#050505] overflow-hidden">
    <!-- 动态背景 -->
    <div class="absolute inset-0">
       <div class="absolute top-[-10%] left-[-10%] w-[40%] h-[40%] bg-primary/20 blur-[120px] rounded-full animate-pulse"></div>
       <div class="absolute bottom-[-10%] right-[-10%] w-[40%] h-[40%] bg-accent/10 blur-[120px] rounded-full animate-pulse" style="animation-delay: 1s"></div>
    </div>

    <div class="glass-card w-full max-w-[420px] p-10 relative z-10 animate-fade-in">
      <div class="text-center mb-10">
        <div class="w-16 h-16 bg-gradient-to-br from-primary to-accent rounded-2xl flex items-center justify-center shadow-2xl shadow-primary/20 mx-auto mb-6">
           <span class="text-white font-black text-3xl italic">U</span>
        </div>
        <h2 class="text-3xl font-black italic tracking-tighter text-white uppercase">{{ isLogin ? '验证访问' : '加入极前' }}</h2>
        <p class="text-gray-500 text-xs font-bold tracking-widest mt-2">ULTRAFORWARD SECURITY GATE</p>
      </div>

      <form @submit.prevent="handleSubmit" class="space-y-6">
        <div v-if="error" class="bg-red-500/10 border border-red-500/20 text-red-400 text-[10px] font-bold p-3 rounded-xl text-center uppercase tracking-widest">
           {{ error }}
        </div>

        <div class="space-y-4">
          <div class="space-y-1">
             <label class="text-[10px] font-black text-gray-500 uppercase tracking-widest px-1">账号标识</label>
             <input v-model="form.username" type="text" placeholder="Username" required
                    class="w-full bg-white/5 border border-white/5 rounded-xl px-5 py-3 text-white focus:outline-none focus:border-primary transition-all font-medium" />
          </div>
          <div class="space-y-1">
             <label class="text-[10px] font-black text-gray-500 uppercase tracking-widest px-1">安全密钥</label>
             <input v-model="form.password" type="password" placeholder="Password" required
                    class="w-full bg-white/5 border border-white/5 rounded-xl px-5 py-3 text-white focus:outline-none focus:border-primary transition-all font-medium" />
          </div>
          <div v-if="!isLogin" class="space-y-1 animate-fade-in">
             <label class="text-[10px] font-black text-gray-500 uppercase tracking-widest px-1">重复校验</label>
             <input v-model="form.confirmPassword" type="password" placeholder="Confirm Password" required
                    class="w-full bg-white/5 border border-white/5 rounded-xl px-5 py-3 text-white focus:outline-none focus:border-primary transition-all font-medium" />
          </div>
        </div>

        <button type="submit" :disabled="loading" class="btn-aura w-full py-4 text-sm tracking-[4px] uppercase italic">
          {{ loading ? 'PROCESSING...' : (isLogin ? 'LOG IN' : 'REGISTER') }}
        </button>

        <div class="text-center pt-4">
          <button type="button" @click="isLogin = !isLogin" class="text-[10px] font-bold text-gray-500 hover:text-white transition-colors uppercase tracking-[2px]">
            {{ isLogin ? '还没有账号？现在申请加入' : '已有授权？立即返回登录' }}
          </button>
        </div>
      </form>

      <!-- 装饰 -->
      <div class="mt-12 text-center text-[8px] text-gray-700 font-black uppercase tracking-[5px]">
         Secured by Stealth-Pass Protocol
      </div>
    </div>
  </div>
</template>

<style scoped>
.animate-fade-in {
  animation: fadeIn 0.6s cubic-bezier(0.16, 1, 0.3, 1);
}
@keyframes fadeIn {
  from { opacity: 0; transform: translateY(20px); }
  to { opacity: 1; transform: translateY(0); }
}
</style>
