package collector

import (
	"github.com/D9ni3l/werc/internal/collector/ad"
	"github.com/D9ni3l/werc/internal/collector/adcs"
	"github.com/D9ni3l/werc/internal/collector/adfs"
	"github.com/D9ni3l/werc/internal/collector/cache"
	"github.com/D9ni3l/werc/internal/collector/container"
	"github.com/D9ni3l/werc/internal/collector/cpu"
	"github.com/D9ni3l/werc/internal/collector/cpu_info"
	"github.com/D9ni3l/werc/internal/collector/cs"
	"github.com/D9ni3l/werc/internal/collector/dfsr"
	"github.com/D9ni3l/werc/internal/collector/dhcp"
	"github.com/D9ni3l/werc/internal/collector/diskdrive"
	"github.com/D9ni3l/werc/internal/collector/dns"
	"github.com/D9ni3l/werc/internal/collector/exchange"
	"github.com/D9ni3l/werc/internal/collector/filetime"
	"github.com/D9ni3l/werc/internal/collector/fsrmquota"
	"github.com/D9ni3l/werc/internal/collector/hyperv"
	"github.com/D9ni3l/werc/internal/collector/iis"
	"github.com/D9ni3l/werc/internal/collector/license"
	"github.com/D9ni3l/werc/internal/collector/logical_disk"
	"github.com/D9ni3l/werc/internal/collector/logon"
	"github.com/D9ni3l/werc/internal/collector/memory"
	"github.com/D9ni3l/werc/internal/collector/mscluster"
	"github.com/D9ni3l/werc/internal/collector/msmq"
	"github.com/D9ni3l/werc/internal/collector/mssql"
	"github.com/D9ni3l/werc/internal/collector/net"
	"github.com/D9ni3l/werc/internal/collector/netframework"
	"github.com/D9ni3l/werc/internal/collector/nps"
	"github.com/D9ni3l/werc/internal/collector/os"
	"github.com/D9ni3l/werc/internal/collector/pagefile"
	"github.com/D9ni3l/werc/internal/collector/perfdata"
	"github.com/D9ni3l/werc/internal/collector/physical_disk"
	"github.com/D9ni3l/werc/internal/collector/printer"
	"github.com/D9ni3l/werc/internal/collector/process"
	"github.com/D9ni3l/werc/internal/collector/remote_fx"
	"github.com/D9ni3l/werc/internal/collector/scheduled_task"
	"github.com/D9ni3l/werc/internal/collector/service"
	"github.com/D9ni3l/werc/internal/collector/smb"
	"github.com/D9ni3l/werc/internal/collector/smbclient"
	"github.com/D9ni3l/werc/internal/collector/smtp"
	"github.com/D9ni3l/werc/internal/collector/system"
	"github.com/D9ni3l/werc/internal/collector/tcp"
	"github.com/D9ni3l/werc/internal/collector/terminal_services"
	"github.com/D9ni3l/werc/internal/collector/textfile"
	"github.com/D9ni3l/werc/internal/collector/thermalzone"
	"github.com/D9ni3l/werc/internal/collector/time"
	"github.com/D9ni3l/werc/internal/collector/udp"
	"github.com/D9ni3l/werc/internal/collector/update"
	"github.com/D9ni3l/werc/internal/collector/vmware"
        "github.com/d9ni3l/werc/internal/collector/rdp_client"
)

type Config struct {
	AD               ad.Config                `yaml:"ad"`
	ADCS             adcs.Config              `yaml:"adcs"`
	ADFS             adfs.Config              `yaml:"adfs"`
	Cache            cache.Config             `yaml:"cache"`
	Container        container.Config         `yaml:"container"`
	CPU              cpu.Config               `yaml:"cpu"`
	CPUInfo          cpu_info.Config          `yaml:"cpu_info"`
	Cs               cs.Config                `yaml:"cs"`
	DFSR             dfsr.Config              `yaml:"dfsr"`
	Dhcp             dhcp.Config              `yaml:"dhcp"`
	DiskDrive        diskdrive.Config         `yaml:"disk_drive"`
	DNS              dns.Config               `yaml:"dns"`
	Exchange         exchange.Config          `yaml:"exchange"`
	Filetime         filetime.Config          `yaml:"filetime"`
	Fsrmquota        fsrmquota.Config         `yaml:"fsrmquota"`
	HyperV           hyperv.Config            `yaml:"hyper_v"`
	IIS              iis.Config               `yaml:"iis"`
	License          license.Config           `yaml:"license"`
	LogicalDisk      logical_disk.Config      `yaml:"logical_disk"`
	Logon            logon.Config             `yaml:"logon"`
	Memory           memory.Config            `yaml:"memory"`
	MSCluster        mscluster.Config         `yaml:"ms_cluster"`
	Msmq             msmq.Config              `yaml:"msmq"`
	Mssql            mssql.Config             `yaml:"mssql"`
	Net              net.Config               `yaml:"net"`
	NetFramework     netframework.Config      `yaml:"net_framework"`
	Nps              nps.Config               `yaml:"nps"`
	OS               os.Config                `yaml:"os"`
	Paging           pagefile.Config          `yaml:"paging"`
	PerfData         perfdata.Config          `yaml:"perf_data"`
	PhysicalDisk     physical_disk.Config     `yaml:"physical_disk"`
	Printer          printer.Config           `yaml:"printer"`
	Process          process.Config           `yaml:"process"`
	RemoteFx         remote_fx.Config         `yaml:"remote_fx"`
	ScheduledTask    scheduled_task.Config    `yaml:"scheduled_task"`
	Service          service.Config           `yaml:"service"`
	SMB              smb.Config               `yaml:"smb"`
	SMBClient        smbclient.Config         `yaml:"smb_client"`
	SMTP             smtp.Config              `yaml:"smtp"`
	System           system.Config            `yaml:"system"`
	TCP              tcp.Config               `yaml:"tcp"`
	TerminalServices terminal_services.Config `yaml:"terminal_services"`
	Textfile         textfile.Config          `yaml:"textfile"`
	ThermalZone      thermalzone.Config       `yaml:"thermal_zone"`
	Time             time.Config              `yaml:"time"`
	UDP              udp.Config               `yaml:"udp"`
	Update           update.Config            `yaml:"update"`
	Vmware           vmware.Config            `yaml:"vmware"`
        RDPClient        rdp_client.Config        `yaml:"rdp_client"`
}

// ConfigDefaults Is an interface to be used by the external libraries. It holds all ConfigDefaults form all collectors
//
//goland:noinspection GoUnusedGlobalVariable
var ConfigDefaults = Config{
	AD:               ad.ConfigDefaults,
	ADCS:             adcs.ConfigDefaults,
	ADFS:             adfs.ConfigDefaults,
	Cache:            cache.ConfigDefaults,
	Container:        container.ConfigDefaults,
	CPU:              cpu.ConfigDefaults,
	CPUInfo:          cpu_info.ConfigDefaults,
	Cs:               cs.ConfigDefaults,
	DFSR:             dfsr.ConfigDefaults,
	Dhcp:             dhcp.ConfigDefaults,
	DiskDrive:        diskdrive.ConfigDefaults,
	DNS:              dns.ConfigDefaults,
	Exchange:         exchange.ConfigDefaults,
	Filetime:         filetime.ConfigDefaults,
	Fsrmquota:        fsrmquota.ConfigDefaults,
	HyperV:           hyperv.ConfigDefaults,
	IIS:              iis.ConfigDefaults,
	License:          license.ConfigDefaults,
	LogicalDisk:      logical_disk.ConfigDefaults,
	Logon:            logon.ConfigDefaults,
	Memory:           memory.ConfigDefaults,
	MSCluster:        mscluster.ConfigDefaults,
	Msmq:             msmq.ConfigDefaults,
	Mssql:            mssql.ConfigDefaults,
	Net:              net.ConfigDefaults,
	NetFramework:     netframework.ConfigDefaults,
	Nps:              nps.ConfigDefaults,
	OS:               os.ConfigDefaults,
	Paging:           pagefile.ConfigDefaults,
	PerfData:         perfdata.ConfigDefaults,
	PhysicalDisk:     physical_disk.ConfigDefaults,
	Printer:          printer.ConfigDefaults,
	Process:          process.ConfigDefaults,
	RemoteFx:         remote_fx.ConfigDefaults,
	ScheduledTask:    scheduled_task.ConfigDefaults,
	Service:          service.ConfigDefaults,
	SMB:              smb.ConfigDefaults,
	SMBClient:        smbclient.ConfigDefaults,
	SMTP:             smtp.ConfigDefaults,
	System:           system.ConfigDefaults,
	TCP:              tcp.ConfigDefaults,
	TerminalServices: terminal_services.ConfigDefaults,
	Textfile:         textfile.ConfigDefaults,
	ThermalZone:      thermalzone.ConfigDefaults,
	Time:             time.ConfigDefaults,
	UDP:              udp.ConfigDefaults,
	Update:           update.ConfigDefaults,
	Vmware:           vmware.ConfigDefaults,
        RDPClient:        rdp_client.ConfigDefaults,
}
