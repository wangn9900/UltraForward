#!/bin/bash
# UltraForward (极前) - 旗舰级隐身转发系统 - 全自动安装程序
# 官方仓库: https://github.com/wangn9900/UltraForward

set -e

COLOR_RED='\033[0;31m'
COLOR_GREEN='\033[0;32m'
COLOR_YELLOW='\033[0;33m'
COLOR_NC='\033[0m'

function install_controller() {
    clear
    echo -e "${COLOR_GREEN}#################################################${COLOR_NC}"
    echo -e "${COLOR_GREEN}#       UltraForward (极前) - 全自动部署模式    #${COLOR_NC}"
    echo -e "${COLOR_GREEN}#################################################${COLOR_NC}"
    
    echo -e "${COLOR_YELLOW}1. 正在检测系统架构...${COLOR_NC}"
    ARCH=$(uname -m)
    case "$ARCH" in
        x86_64) ARCH="amd64" ;;
        aarch64|arm64) ARCH="arm64" ;;
        *) echo -e "${COLOR_RED}不支持的架构: $ARCH${COLOR_NC}"; exit 1 ;;
    esac

    echo -e "${COLOR_YELLOW}2. 正在清理并准备环境...${COLOR_NC}"
    systemctl stop ultra-controller 2>/dev/null || true
    mkdir -p /usr/local/bin
    mkdir -p /etc/ultraforward

    echo -e "${COLOR_YELLOW}3. 正在下载旗舰版内核套件 (v1.1.0)...${COLOR_NC}"
    curl -L -o /usr/local/bin/ultra-controller "https://github.com/wangn9900/UltraForward/releases/latest/download/ultra-controller-linux-$ARCH"
    chmod +x /usr/local/bin/ultra-controller

    echo -e "${COLOR_YELLOW}4. 正在配置系统守护服务 (Systemd)...${COLOR_NC}"
    cat > /etc/systemd/system/ultra-controller.service <<EOF
[Unit]
Description=UltraForward Controller Service
After=network.target

[Service]
Type=simple
User=root
WorkingDirectory=/etc/ultraforward
ExecStart=/usr/local/bin/ultra-controller
Restart=always
RestartSec=5

[Install]
WantedBy=multi-user.target
EOF

    echo -e "${COLOR_YELLOW}5. 正在强制拉起后台服务...${COLOR_NC}"
    systemctl daemon-reload
    systemctl enable ultra-controller
    systemctl start ultra-controller

    echo -e "${COLOR_YELLOW}6. 正在获取公网访问地址...${COLOR_NC}"
    IP=$(curl -s ifconfig.me || curl -s api.ipify.org)
    
    echo -e "${COLOR_GREEN}#################################################${COLOR_NC}"
    echo -e "${COLOR_GREEN}🎊  UltraForward (极前) 部署成功！${COLOR_NC}"
    echo -e "-------------------------------------------------"
    echo -e "🌍  控制面板地址: ${COLOR_YELLOW}http://${IP}:8080${COLOR_NC}"
    echo -e "🔑  初始操作: 请先在页面上注册 Admin 账号"
    echo -e "-------------------------------------------------"
    echo -e "🛠️  管理命令:"
    echo -e "   - 启动: systemctl start ultra-controller"
    echo -e "   - 停止: systemctl stop ultra-controller"
    echo -e "   - 日志: journalctl -u ultra-controller -f"
    echo -e "${COLOR_GREEN}#################################################${COLOR_NC}"
}

# 自动开始安装 (不再等待输入)
install_controller
