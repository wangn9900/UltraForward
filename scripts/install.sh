#!/bin/bash
# UltraForward (极前) - 旗舰级隐身转发系统
# 官方仓库: https://github.com/wangn9900/UltraForward

# 强制转换换行符环境适配
set -e

COLOR_RED='\033[0;31m'
COLOR_GREEN='\033[0;32m'
COLOR_YELLOW='\033[0;33m'
COLOR_NC='\033[0m'

function show_menu() {
    clear
    echo -e "${COLOR_GREEN}#################################################${COLOR_NC}"
    echo -e "${COLOR_GREEN}#       UltraForward (极前) V1.0.0              #${COLOR_NC}"
    echo -e "${COLOR_GREEN}#################################################${COLOR_NC}"
    echo " 1) 安装主控端 (Controller)"
    echo " 2) 安装受控端 (Agent)"
    echo " 3) 卸载系统"
    echo " q) 退出"
    echo -e "${COLOR_GREEN}#################################################${COLOR_NC}"
    printf "请选择操作 [1-3/q]: "
    read choice
}

function install_controller() {
    echo -e "${COLOR_YELLOW}正在下载最新主控内核...${COLOR_NC}"
    # 确保目录存在
    mkdir -p /usr/local/bin
    mkdir -p /etc/ultraforward
    
    # 获取架构
    ARCH=$(uname -m)
    if [ "$ARCH" == "x86_64" ]; then ARCH="amd64"; fi
    if [ "$ARCH" == "aarch64" ]; then ARCH="arm64"; fi

    curl -L -o /usr/local/bin/ultra-controller "https://github.com/wangn9900/UltraForward/releases/latest/download/ultra-controller-linux-$ARCH"
    chmod +x /usr/local/bin/ultra-controller
    echo -e "${COLOR_GREEN}主控端安装成功！${COLOR_NC}"
    echo -e "配置文件目录: /etc/ultraforward"
    echo -e "请使用 systemctl start ultra-controller 启动服务 (需先配置 Service 文件)"
}

show_menu

case "$choice" in
    1)
        install_controller
        ;;
    2)
        echo "受控端安装逻辑开发中..."
        ;;
    3)
        echo "卸载逻辑开发中..."
        ;;
    *)
        echo "退出脚本"
        exit 0
        ;;
esac
