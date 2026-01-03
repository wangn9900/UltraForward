# UltraForward (极前) 🚀

**The Next-Generation Stealth Relay Platform.**

UltraForward 是从 StealthForward 项目中剥离出来的商用旗舰版。它去除了所有传统的 VLESS 转发包袱，专注于通过 **Stealth-Pass** 协议实现的高性能、高隐身性的专线中转服务。

## ✨ 核心特性

- **Stealth-Pass 协议**：基于 Chacha20-Poly1305 的全加密隧道，内置长度混淆与流量分片，完美绕过边界探测。
- **旗舰级 UI/UX**：全新的侧边栏布局，玻璃拟态（Glassmorphism）设计语言，带来极致的操控体验。
- **商业级计费系统**：内置流量倍率结算、用户余额管理、套餐商城与订单追踪。
- **自动化运维**：集成 SSH 编排引擎，一键下发中转规则并自动配置 Linux Systemd 服务。

## 🛠️ 快速开始 (开发环境)

### 1. 启动后端控制器
```bash
go mod tidy
go run ./cmd/controller
```

### 2. 启动前端界面
```bash
cd web-new
npm install
npm run dev
```

## 📜 许可证

UltraForward License. (Modified commercial use only)
