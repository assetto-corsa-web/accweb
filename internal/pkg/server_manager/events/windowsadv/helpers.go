package windowsadv

import (
	"fmt"
	"os/exec"
	"runtime"

	"github.com/assetto-corsa-web/accweb/internal/pkg/instance"
	"github.com/sirupsen/logrus"
)

func setCoreAffinity(pid int, value uint) error {
	args := []string{fmt.Sprintf("$Process = Get-Process -Id %d; $Process.ProcessorAffinity=%d", pid, value)}
	cmd := exec.Command("PowerShell", args...)
	return cmd.Start()
}

func setCpuPriority(pid int, p uint) error {
	args := []string{"process", "where", fmt.Sprintf("ProcessId=%d", pid), "call", "setpriority", fmt.Sprintf("%d", p)}
	cmd := exec.Command("wmic", args...)
	return cmd.Start()
}

func addFirewallRules(pid, tcp, udp int) error {
	if err := addFWRule(pid, tcp, "TCP"); err != nil {
		return err
	}

	return addFWRule(pid, udp, "UDP")
}

func addFWRule(pid, port int, t string) error {
	args := []string{
		"advfirewall", "firewall", "add", "rule", fmt.Sprintf("name=\"ACCSERVER_%d\"", pid),
		"dir=in", "action=allow", fmt.Sprintf("protocol=%s", t), fmt.Sprintf("localport=%d", port),
	}

	cmd := exec.Command("netsh.exe", args...)
	return cmd.Start()
}

func delFirewallRules(pid int) error {
	args := []string{"advfirewall", "firewall", "del", "rule", fmt.Sprintf("name=\"ACCSERVER_%d\"", pid)}
	cmd := exec.Command("netsh.exe", args...)
	return cmd.Start()
}

func startWithAdvWindows(s *instance.Instance) {
	cfg := s.Cfg.Settings.AdvWindowsCfg
	l := logrus.WithField("server_id", s.GetID()).WithField("PID", s.GetProcessID())

	l.Infof("Defining core affinity to %d", cfg.CoreAffinity)
	if err := setCoreAffinity(s.GetProcessID(), cfg.CoreAffinity); err != nil {
		l.Errorf("failed to define affinity with value: %d. ERROR: %s", cfg.CoreAffinity, err.Error())
	}

	l.Infof("Defining cpu priority to %d", cfg.CpuPriority)
	if err := setCpuPriority(s.GetProcessID(), cfg.CpuPriority); err != nil {
		l.Errorf("failed to define cpu priority with value: %d. ERROR: %s", cfg.CpuPriority, err.Error())
	}

	if cfg.EnableWinFW {
		l.Info("Add Firewall Rules")
		if err := addFirewallRules(s.GetProcessID(), s.AccCfg.Configuration.TcpPort, s.AccCfg.Configuration.UdpPort); err != nil {
			l.Errorf("Failed to add accserver firewall rule. ERROR: %s", err.Error())
		}
	}
}

func stopWithAdvWindows(s *instance.Instance) {
	if !s.Cfg.Settings.AdvWindowsCfg.EnableWinFW {
		return
	}

	logrus.Info("Removing Firewall Rules")
	if err := delFirewallRules(s.GetProcessID()); err != nil {
		logrus.Errorf("Failed to add accserver firewall rule for TCP. ERROR: %s", err.Error())
	}
}

func hasAdvancedWindowsConfig(s *instance.Instance) bool {
	return runtime.GOOS == "windows" && s.Cfg.Settings.EnableAdvWinCfg
}
