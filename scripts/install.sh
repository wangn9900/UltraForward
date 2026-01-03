#!/bin/bash
# UltraForward (极前) - 旗舰级隐身转发系统
# 官方仓库: https://github.com/wangn9900/UltraForward

set -e

COLOR_RED='\033[0;31m'
COLOR_GREEN='\033[0;32m'
COLOR_YELLOW='\033[0;33m'
COLOR_NC='\033[0m'

function show_menu() {
    clear
    echo -e "${COLOR_GREEN}#################################################${COLOR_NC}"
    echo -e "${COLOR_GREEN}#       UltraForward (极前) V1.0.4              #${COLOR_NC}"
    echo -e "${COLOR_GREEN}#################################################${COLOR_NC}"
    echo " 1) 安装主控端 (Controller)"
    echo " 2) 安装受控端 (Agent)"
    echo " 3) 卸载系统"
    echo " q) 退出"
    echo -e "${COLOR_GREEN}#################################################${COLOR_NC}"
    printf "请选择操作 [1-3/q]: "
    # 修复管道环境下的输入读取问题
    read choice < /dev/tty
}

function install_controller() {
    echo -e "${COLOR_YELLOW}正在检测系统架构...${COLOR_NC}"
    ARCH=$(uname -m)
    case "$ARCH" in
        x86_64) ARCH="amd64" ;;
        aarch64|arm64) ARCH="arm64" ;;
        *) echo "不支持的架构: $ARCH"; exit 1 ;;
    esac

    echo -e "${COLOR_YELLOW}正在下载最新主控内核 ($ARCH)...${COLOR_NC}"
    mkdir -p /usr/local/bin
    
    # 下载地址指向最新的 release
    curl -L -o /usr/local/bin/ultra-controller "https://github.com/wangn9900/UltraForward/releases/latest/download/ultra-controller-linux-$ARCH"
    chmod +x /usr/local/bin/ultra-controller
    
    echo -e "${COLOR_GREEN}主控端安装成功！${COLOR_NC}"
    echo -e "你可以通过执行 /usr/local/bin/ultra-controller 尝试手动启动。"
    echo -e "或者稍后我会为你补全 systemd 服务配置。"
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
    q|Q)
        echo "退出脚本"
        exit 0
        ;;
    *)
        echo "无效选项: $choice"
        exit 1
        ;;
esac
