package util

import (
	"net"
	"strings"
)

// isVirtualMachineIP 判断IP是否属于常见虚拟机使用的IP段
func isVirtualMachineIP(ip net.IP) bool {
	// 可以根据实际情况添加更多虚拟机环境的特定IP前缀
	vmIPRanges := []string{
		"10.0.2.", // 常见于VirtualBox
		"192.168.146.",
		"192.168.152.",
	}
	for _, prefix := range vmIPRanges {
		if strings.HasPrefix(ip.String(), prefix) {
			return true
		}
	}
	return false
}

// GetLocalIP 获取过滤掉虚拟机IP后的内网IPv4地址
func GetLocalIP() string {
	ifaces, err := net.Interfaces()
	if err != nil {
		return ""
	}

	for _, iface := range ifaces {
		if iface.Flags&net.FlagUp == 0 || iface.Flags&net.FlagLoopback != 0 {
			continue
		}

		addrs, err := iface.Addrs()
		if err != nil {
			continue
		}

		for _, addr := range addrs {
			var ip net.IP
			switch v := addr.(type) {
			case *net.IPNet:
				ip = v.IP
			case *net.IPAddr:
				ip = v.IP
			default:
				continue
			}

			if ip.IsPrivate() && ip.To4() != nil && !isVirtualMachineIP(ip) {
				return ip.String()
			}
		}
	}

	return ""
}
