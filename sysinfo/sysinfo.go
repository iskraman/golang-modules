package sysinfo

import (
	"strconv"

	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/host"
	"github.com/shirou/gopsutil/mem"
)

type SysInfo struct {
	Hostname string `json:"hostname,omitempty"`
	Platform string `json:"platform,omitempty"`
	Uuid     string `json:"uuid,omitempty"`
	Uptime   string `json:"uptime,omitempty"`

	Cpu CpuInfo `json:"cpu,omitempty"`
	Mem MemInfo `json:"mem,omitempty"`
	Hdd HddInfo `json:"hdd,omitempty"`
}

type CpuInfo struct {
	Model   string `json:"model,omitempty"`
	Speed   string `json:"speed,omitempty"`
	Core    string `json:"core,omitempty"`
	Percent string `json:"percent,omitempty"`
}

type MemInfo struct {
	Total   string `json:"total,omitempty"`
	Free    string `json:"free,omitempty"`
	Percent string `json:"percent,omitempty"`
}

type HddInfo struct {
	Total   string `json:"total,omitempty"`
	Used    string `json:"used,omitempty"`
	Free    string `json:"free,omitempty"`
	Percent string `json:"percent,omitempty"`
}

func Cpu() CpuInfo {
	cpuStat, _ := cpu.Info()
	percentage, _ := cpu.Percent(0, true)

	cpuinfo := CpuInfo{}
	cpuinfo.Model = cpuStat[0].ModelName
	cpuinfo.Speed = strconv.FormatFloat(cpuStat[0].Mhz, 'f', 2, 64)
	cpuinfo.Core = strconv.FormatInt(int64(len(percentage)), 10)

	percent := 0.0
	for _, cpupercent := range percentage {
		percent += cpupercent
	}
	cpuinfo.Percent = strconv.FormatFloat(percent, 'f', 2, 64)
	//cpuinfo.Percent = cpuinfo.Percent / float64(cpuinfo.Core)
	return cpuinfo
}

func Mem() MemInfo {
	// memory
	vmStat, _ := mem.VirtualMemory()

	meminfo := MemInfo{}
	meminfo.Total = strconv.FormatUint(vmStat.Total, 10)
	meminfo.Free = strconv.FormatUint(vmStat.Free, 10)
	meminfo.Percent = strconv.FormatFloat(vmStat.UsedPercent, 'f', 2, 64)
	return meminfo
}

func Hdd() HddInfo {
	diskStat, _ := disk.Usage("/")

	hddinfo := HddInfo{}
	hddinfo.Total = strconv.FormatUint(diskStat.Total, 10)
	hddinfo.Used = strconv.FormatUint(diskStat.Used, 10)
	hddinfo.Free = strconv.FormatUint(diskStat.Free, 10)
	hddinfo.Percent = strconv.FormatFloat(diskStat.UsedPercent, 'f', 2, 64)

	return hddinfo
}

func System() SysInfo {
	hostStat, _ := host.Info()

	sysinfo := SysInfo{}
	sysinfo.Platform = hostStat.Platform
	sysinfo.Hostname = hostStat.Hostname
	sysinfo.Uuid = hostStat.HostID
	sysinfo.Uptime = strconv.FormatUint(hostStat.Uptime, 10)

	sysinfo.Cpu = Cpu()
	sysinfo.Mem = Mem()
	sysinfo.Hdd = Hdd()

	return sysinfo
}

/*
func main() {
	for {
		sysinfo := System()
		json, _ := jsonlib.EncodingIndent(sysinfo)
		syslog.STDLN(string(json))
		time.Sleep(time.Second * 2)
	}
}
*/
