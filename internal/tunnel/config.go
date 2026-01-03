package tunnel

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
)

type Task struct {
	ID         uint   `json:"id"`
	Mode       string `json:"mode"` // transit, exit
	ListenAddr string `json:"listen_addr"`
	TargetAddr string `json:"target_addr"` // 只有 transit 需要
	Key        string `json:"key"`
}

type TunnelConfig struct {
	Tasks []Task `json:"tasks"`
}

func RunStandalone(configPath string) error {
	data, err := os.ReadFile(configPath)
	if err != nil {
		return err
	}

	var cfg TunnelConfig
	if err := json.Unmarshal(data, &cfg); err != nil {
		return err
	}

	ctx := context.Background()
	for _, t := range cfg.Tasks {
		if t.Mode == "transit" {
			srv := &TransitServer{
				RuleID:     t.ID,
				ListenAddr: t.ListenAddr,
				TargetAddr: t.TargetAddr,
				Key:        t.Key,
				Stats:      &TrafficCounter{},
			}
			go srv.Start(ctx)
		} else {
			srv := &ExitServer{
				RuleID:     t.ID,
				ListenAddr: t.ListenAddr,
				LocalDest:  t.TargetAddr,
				Key:        t.Key,
			}
			go srv.Start(ctx)
		}
	}

	fmt.Printf("[UltraForward] Started %d relay tasks. Running...\n", len(cfg.Tasks))
	select {} // 永久阻塞
}
