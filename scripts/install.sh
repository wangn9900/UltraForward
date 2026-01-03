#!/bin/bash
# UltraForward (极前) - 旗舰级隐身转发系统
# 官方仓库: https://github.com/wangn9900/UltraForward

COLOR_RED='\033[0;31m'
COLOR_GREEN='\033[0;32m'
COLOR_YELLOW='\033[0;33m'
COLOR_NC='\033[0m'

function show_menu() {
    echo -e "${COLOR_GREEN}UltraForward (极前) V1.0.0${COLOR_NC}"
    echo "1) 安装主控端 (Controller)"
    echo "2) 安装受控端 (Agent)"
    echo "3) 卸载系统"
    echo "q) 退出"
    read -p "请选择操作: " choice
}

function install_controller() {
    echo -e "${COLOR_YELLOW}正在下载最新主控内核...${COLOR_NC}"
    curl -L -o /usr/local/bin/ultra-controller https://github.com/wangn9900/UltraForward/releases/latest/download/ultra-controller-linux-amd64
    chmod +x /usr/local/bin/ultra-controller
    echo "主控端安装成功，请使用 systemctl start ultra-controller 启动。"
}

show_menu
case $choice in
    1) install_controller ;;
    *) exit 0 ;;
esac
