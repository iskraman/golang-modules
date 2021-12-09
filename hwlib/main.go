// https://github.com/shirou/gopsutil
// https://socketloop.com/tutorials/golang-get-hardware-information-such-as-disk-memory-and-cpu-usage
// https://minwook-shin.github.io/go-process-system-utilization-gopsutil/

package main

import (
	"fmt"
	"runtime"
	"strconv"

	"github.com/iskraman/golang-modules/utils/syslog"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/host"
	"github.com/shirou/gopsutil/mem"
	"github.com/shirou/gopsutil/net"
)

func dealwithErr(err error) {
	if err != nil {
		fmt.Println(err)
		//os.Exit(-1)
	}
}

func Cpu() {
	// cpu - get CPU number of cores and speed
	cpuStat, err := cpu.Info()
	dealwithErr(err)
	percentage, err := cpu.Percent(0, true)
	dealwithErr(err)

	// since my machine has one CPU, I'll use the 0 index
	// if your machine has more than 1 CPU, use the correct index
	// to get the proper data
	syslog.STDLN("CPU index number: " + strconv.FormatInt(int64(cpuStat[0].CPU), 10) + "<br>")
	syslog.STDLN("VendorID: " + cpuStat[0].VendorID + "<br>")
	syslog.STDLN("Family: " + cpuStat[0].Family + "<br>")
	syslog.STDLN("Number of cores: " + strconv.FormatInt(int64(cpuStat[0].Cores), 10) + "<br>")
	syslog.STDLN("Model Name: " + cpuStat[0].ModelName + "<br>")
	syslog.STDLN("Speed: " + strconv.FormatFloat(cpuStat[0].Mhz, 'f', 2, 64) + " MHz <br>")

	for idx, cpupercent := range percentage {
		syslog.STDLN("Current CPU utilization: [" + strconv.Itoa(idx) + "] " + strconv.FormatFloat(cpupercent, 'f', 2, 64) + "%<br>")
	}

}

func Mem() {
	runtimeOS := runtime.GOOS
	// memory
	vmStat, err := mem.VirtualMemory()
	dealwithErr(err)

	syslog.STDLN("<html>OS : " + runtimeOS + "<br>")
	syslog.STDLN("Total memory: " + strconv.FormatUint(vmStat.Total, 10) + " bytes <br>")
	syslog.STDLN("Free memory: " + strconv.FormatUint(vmStat.Free, 10) + " bytes<br>")
	syslog.STDLN("Percentage used memory: " + strconv.FormatFloat(vmStat.UsedPercent, 'f', 2, 64) + "%<br>")

}

func Hdd() {
	diskStat, err := disk.Usage("/")
	dealwithErr(err)

	syslog.STDLN("Total disk space: " + strconv.FormatUint(diskStat.Total, 10) + " bytes <br>")
	syslog.STDLN("Used disk space: " + strconv.FormatUint(diskStat.Used, 10) + " bytes<br>")
	syslog.STDLN("Free disk space: " + strconv.FormatUint(diskStat.Free, 10) + " bytes<br>")
	syslog.STDLN("Percentage disk space usage: " + strconv.FormatFloat(diskStat.UsedPercent, 'f', 2, 64) + "%<br>")
}

func Net() {
	// get interfaces MAC/hardware address
	interfStat, err := net.Interfaces()
	dealwithErr(err)
	for _, interf := range interfStat {
		syslog.STDLN("------------------------------------------------------<br>")
		syslog.STDLN("Interface Name: " + interf.Name + "<br>")

		if interf.HardwareAddr != "" {
			syslog.STDLN("Hardware(MAC) Address: " + interf.HardwareAddr + "<br>")
		}

		for _, flag := range interf.Flags {
			syslog.STDLN("Interface behavior or flags: " + flag + "<br>")
		}

		for _, addr := range interf.Addrs {
			syslog.STDLN("IPv6 or IPv4 addresses: " + addr.String() + "<br>")

		}
	}
}

func Host() {
	runtimeOS := runtime.GOOS
	// host or machine kernel, uptime, platform Info
	hostStat, err := host.Info()
	dealwithErr(err)

	syslog.STDLN("<html>OS : " + runtimeOS + "<br>")
	syslog.STDLN("Hostname: " + hostStat.Hostname + "<br>")
	syslog.STDLN("Uptime: " + strconv.FormatUint(hostStat.Uptime, 10) + "<br>")
	syslog.STDLN("Number of processes running: " + strconv.FormatUint(hostStat.Procs, 10) + "<br>")

	// another way to get the operating system name
	// both darwin for Mac OSX, For Linux, can be ubuntu as platform
	// and linux for OS

	syslog.STDLN("OS: " + hostStat.OS + "<br>")
	syslog.STDLN("Platform: " + hostStat.Platform + "<br>")

	// the unique hardware id for this machine
	syslog.STDLN("Host ID(uuid): " + hostStat.HostID + "<br>")
}

func main() {
	Cpu()
	Mem()
	Hdd()
	Net()
	Host()
}
