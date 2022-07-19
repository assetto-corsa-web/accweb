package helper

import (
	"fmt"
	"os/exec"
)

func SetCoreAffinity(pid int, value uint) error {
	args := []string{fmt.Sprintf("$Process = Get-Process -Id %d; $Process.ProcessorAffinity=%d", pid, value)}
	cmd := exec.Command("PowerShell", args...)
	return cmd.Start()
}

func SetCpuPriority(pid int, p uint) error {
	args := []string{"process", "where", fmt.Sprintf("ProcessId=%d", pid), "call", "setpriority", fmt.Sprintf("%d", p)}
	cmd := exec.Command("wmic", args...)
	return cmd.Start()
}

func AddFirewallRules(pid, tcp, udp int) error {
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

func DelFirewallRules(pid int) error {
	args := []string{"advfirewall", "firewall", "del", "rule", fmt.Sprintf("name=\"ACCSERVER_%d\"", pid)}
	cmd := exec.Command("netsh.exe", args...)
	return cmd.Start()
}
