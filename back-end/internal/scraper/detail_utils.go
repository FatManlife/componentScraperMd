package scraper

import "strings"

func CpuGeneralization(cpu string) string {
	if strings.Contains(cpu, "AMD") {
		return strings.TrimSpace(strings.ReplaceAll(cpu, "AMD", ""))
	} else if strings.Contains(cpu, "Intel") {
		return strings.TrimSpace(strings.ReplaceAll(cpu, "Intel", ""))
	} else {
		return strings.TrimSpace(cpu)
	}
}

func GpuGeneralization(gpu string) string {
	if strings.Contains(gpu, "AMD") {
		return strings.TrimSpace(strings.ReplaceAll(gpu, "AMD", ""))
	} else if gpu == "" {
		return "Integrated"
	} else {
		return strings.TrimSpace(gpu)
	}
}