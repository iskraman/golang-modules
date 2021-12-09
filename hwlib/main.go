package main

import (
	"strconv"
	"time"

	"github.com/iskraman/golang-modules/jsonlib"
	"github.com/iskraman/golang-modules/utils/syslog"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/host"
	"github.com/shirou/gopsutil/mem"
)

type HostInfo struct {
	Hostname string `json:"hostname"`
	Platform string `json:"platform"`
	Uuid     string `json:"uuid"`
	Uptime   string `json:"uptime"`

	Cpu CpuInfo `json:"cpu"`
	Mem MemInfo `json:"mem"`
	Hdd HddInfo `json:"hdd"`
}

type CpuInfo struct {
	Model   string `json:"model"`
	Speed   string `json:"speed"`
	Core    string `json:"core"`
	Percent string `json:"percent"`
}

type MemInfo struct {
	Total   string `json:"total"`
	Free    string `json:"free"`
	Percent string `json:"percent"`
}

type HddInfo struct {
	Total   string `json:"total"`
	Used    string `json:"used"`
	Free    string `json:"free"`
	Percent string `json:"percent"`
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

func SystemInfo() HostInfo {
	hostStat, _ := host.Info()

	sysinfo := HostInfo{}
	sysinfo.Platform = hostStat.Platform
	sysinfo.Hostname = hostStat.Hostname
	sysinfo.Uuid = hostStat.HostID
	sysinfo.Uptime = strconv.FormatUint(hostStat.Uptime, 10)

	sysinfo.Cpu = Cpu()
	sysinfo.Mem = Mem()
	sysinfo.Hdd = Hdd()

	return sysinfo
}

func main() {
	for {
		sysinfo := SystemInfo()
		json, _ := jsonlib.EncodingIndent(sysinfo)
		syslog.STDLN(string(json))
		time.Sleep(time.Second * 2)
	}
}
