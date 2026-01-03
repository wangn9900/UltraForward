<script setup>
import { ref, defineEmits } from 'vue'

const emit = defineEmits(['auth-success'])
const isLogin = ref(true)
const username = ref('')
const password = ref('')
const confirmPassword = ref('')
const loading = ref(false)
const errorMsg = ref('')

const handleAuth = async () => {
  if (!isLogin.value && password.value !== confirmPassword.value) {
    errorMsg.value = '两次输入的密码不一致'
    return
  }

  loading.value = true
  errorMsg.value = ''
  
  const endpoint = isLogin.value ? '/api/v1/auth/login' : '/api/v1/auth/register'
  
  try {
    const res = await fetch(endpoint, {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify({
        username: username.value,
        password: password.value
      })
    })
    
    const data = await res.json()
    
    if (res.ok) {
      if (isLogin.value) {
        localStorage.setItem('uf_token', data.token)
        emit('auth-success', data.user)
      } else {
        isLogin.value = true
        errorMsg.value = '注册成功，请登录'
      }
    } else {
      errorMsg.value = data.error || '认证失败'
    }
  } catch (err) {
    errorMsg.value = '无法连接到服务器'
  } finally {
    loading.value = false
  }
}
</script>

<template>
  <div class="auth-wrapper">
    <!-- 动态背景 -->
    <div class="bg-blobs">
      <div class="blob blob-1"></div>
      <div class="blob blob-2"></div>
      <div class="blob blob-3"></div>
    </div>

    <div class="auth-card">
      <div class="auth-header">
        <div class="logo-aura">
          <span>U</span>
        </div>
        <h1>ULTRA<span>FORWARD</span></h1>
        <p>{{ isLogin ? '旗舰级转发控制台' : '申请加入极前计划' }}</p>
      </div>

      <div class="auth-body">
        <div class="input-group">
          <label>账号标识 / USERNAME</label>
          <input v-model="username" type="text" placeholder="请输入您的账号" @keyup.enter="handleAuth" />
        </div>

        <div class="input-group">
          <label>安全密钥 / PASSWORD</label>
          <input v-model="password" type="password" placeholder="请输入密码" @keyup.enter="handleAuth" />
        </div>

        <div v-if="!isLogin" class="input-group animate-slide-down">
          <label>确认密钥 / CONFIRM</label>
          <input v-model="confirmPassword" type="password" placeholder="请再次输入密码" @keyup.enter="handleAuth" />
        </div>

        <div v-if="errorMsg" class="error-banner">
          {{ errorMsg }}
        </div>

        <button class="auth-btn" :class="{ 'loading': loading }" @click="handleAuth">
          <span>{{ isLogin ? '立即验证身份' : '创建新账号' }}</span>
          <div v-if="loading" class="spinner"></div>
        </button>
      </div>

      <div class="auth-footer">
        <span @click="isLogin = !isLogin">
          {{ isLogin ? '还没有账号? 现在申请加入' : '已有授权账号? 立即返回登录' }}
        </span>
      </div>
    </div>
  </div>
</template>

<style scoped>
.auth-wrapper {
  position: fixed;
  inset: 0;
  display: flex;
  align-items: center;
  justify-content: center;
  background: #050505;
  font-family: 'Outfit', sans-serif;
  overflow: hidden;
}

/* 动态背景球 */
.bg-blobs {
  position: absolute;
  inset: 0;
  overflow: hidden;
  filter: blur(80px);
}
.blob {
  position: absolute;
  width: 500px;
  height: 500px;
  border-radius: 50%;
  opacity: 0.2;
  animation: move 20s infinite alternate;
}
.blob-1 { background: var(--primary); top: -100px; left: -100px; }
.blob-2 { background: var(--accent); bottom: -100px; right: -100px; animation-delay: -5s; }
.blob-3 { background: #6366f1; top: 30%; left: 40%; animation-duration: 30s; }

@keyframes move {
  from { transform: translate(0, 0) scale(1); }
  to { transform: translate(100px, 100px) scale(1.2); }
}

.auth-card {
  position: relative;
  z-index: 10;
  width: 100%;
  max-width: 420px;
  background: rgba(255, 255, 255, 0.03);
  backdrop-filter: blur(20px);
  border: 1px solid rgba(255, 255, 255, 0.08);
  border-radius: 32px;
  padding: 48px;
  box-shadow: 0 25px 50px -12px rgba(0, 0, 0, 0.5);
}

.auth-header {
  text-align: center;
  margin-bottom: 40px;
}

.logo-aura {
  width: 64px;
  height: 64px;
  background: linear-gradient(135deg, var(--primary), var(--accent));
  border-radius: 18px;
  display: flex;
  align-items: center;
  justify-content: center;
  margin: 0 auto 20px;
  font-size: 32px;
  font-weight: 900;
  color: white;
  box-shadow: 0 10px 20px rgba(var(--primary-rgb), 0.3);
}

.auth-header h1 {
  font-size: 24px;
  font-weight: 900;
  letter-spacing: -1px;
  color: white;
}
.auth-header h1 span { color: var(--primary); }
.auth-header p {
  font-size: 13px;
  color: #666;
  margin-top: 8px;
  font-weight: 500;
}

.input-group {
  margin-bottom: 24px;
}
.input-group label {
  display: block;
  font-size: 10px;
  font-weight: 800;
  color: #555;
  text-transform: uppercase;
  margin-bottom: 8px;
  letter-spacing: 1px;
}
.input-group input {
  width: 100%;
  background: rgba(255, 255, 255, 0.05);
  border: 1px solid rgba(255, 255, 255, 0.1);
  border-radius: 14px;
  padding: 14px 18px;
  color: white;
  font-size: 15px;
  transition: all 0.3s;
}
.input-group input:focus {
  outline: none;
  background: rgba(255, 255, 255, 0.08);
  border-color: var(--primary);
  box-shadow: 0 0 0 4px rgba(var(--primary-rgb), 0.1);
}

.error-banner {
  background: rgba(239, 68, 68, 0.1);
  border-left: 3px solid #ef4444;
  padding: 10px 15px;
  border-radius: 8px;
  font-size: 13px;
  color: #ef4444;
  margin-bottom: 24px;
}

.auth-btn {
  width: 100%;
  background: white;
  color: black;
  border: none;
  border-radius: 14px;
  padding: 16px;
  font-weight: 900;
  font-size: 15px;
  cursor: pointer;
  transition: all 0.3s;
  display: flex;
  align-items: center;
  justify-content: center;
  gap: 10px;
}
.auth-btn:hover {
  transform: translateY(-2px);
  box-shadow: 0 10px 20px rgba(255, 255, 255, 0.1);
}
.auth-btn.loading { opacity: 0.7; pointer-events: none; }

.auth-footer {
  margin-top: 32px;
  text-align: center;
}
.auth-footer span {
  font-size: 13px;
  color: #666;
  font-weight: 600;
  cursor: pointer;
  transition: color 0.3s;
}
.auth-footer span:hover { color: white; }

.spinner {
  width: 18px;
  height: 18px;
  border: 3px solid rgba(0, 0, 0, 0.1);
  border-top-color: black;
  border-radius: 50%;
  animation: spin 0.8s linear infinite;
}
@keyframes spin { to { transform: rotate(360deg); } }
</style>
