package models

import (
	"time"

	"gorm.io/gorm"
)

// User 代表商用系统的用户
type User struct {
	gorm.Model
	Username     string    `json:"username" gorm:"unique"`
	Password     string    `json:"-"`
	Balance      float64   `json:"balance"`       // 余额 (CNY)
	TotalTraffic int64     `json:"total_traffic"` // 总额度 (Bytes)
	UsedTraffic  int64     `json:"used_traffic"`  // 已用计费流量
	MaxRules     int       `json:"max_rules"`     // 最大规则数
	ExpiredAt    time.Time `json:"expired_at"`    // 到期时间
	AutoRenewal  bool      `json:"auto_renewal"`  // 自动续费
	PlanID       uint      `json:"plan_id"`
}

// Plan 套餐定义
type Plan struct {
	gorm.Model
	Name         string  `json:"name"`
	Price        float64 `json:"price"`
	Traffic      int64   `json:"traffic"`
	Rules        int     `json:"rules"`
	DurationDays int     `json:"duration_days"`
	Description  string  `json:"description"`
}

// UltraNode 物理服务器资源 (中转或落地)
type UltraNode struct {
	gorm.Model
	Name       string `json:"name"`
	PublicAddr string `json:"public_addr"`
	SSHPort    int    `json:"ssh_port"`
	SSHUser    string `json:"ssh_user"`
	SSHPass    string `json:"ssh_pass"` // 或密钥 ID
	IsActive   bool   `json:"is_active"`
}

// Line 预定义的转发线路 (商用抽象)
type Line struct {
	gorm.Model
	Name          string  `json:"name"`
	TransitNodeID uint    `json:"transit_node_id"`
	ExitNodeID    uint    `json:"exit_node_id"`
	Price         float64 `json:"price"` // 流量倍率
	IsPublic      bool    `json:"is_public"`
}

// Rule 用户实际开启的转发实例
type Rule struct {
	gorm.Model
	Name       string `json:"name"`
	UserID     uint   `json:"user_id"`
	LineID     uint   `json:"line_id"`
	ListenPort int    `json:"listen_port"` // 入口监听
	TunnelPort int    `json:"tunnel_port"` // 隧道内部通讯
	LocalDest  string `json:"local_dest"`  // 落地最终目标
	Key        string `json:"key"`         // 隧道加密 Key
	Upload     int64  `json:"upload"`
	Download   int64  `json:"download"`
	Status     bool   `json:"status"`
	Line       Line   `gorm:"foreignKey:LineID"`
}

// Order 财务订单
type Order struct {
	gorm.Model
	UserID        uint    `json:"user_id"`
	PlanID        uint    `json:"plan_id"`
	Amount        float64 `json:"amount"`
	Status        string  `json:"status"` // pending, completed, failed
	PaymentMethod string  `json:"payment_method"`
	TradeNo       string  `json:"trade_no"`
}
