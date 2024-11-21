package collector

import (
	"maps"
	"slices"

	"github.com/alecthomas/kingpin/v2"
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

func NewBuilderWithFlags[C Collector](fn BuilderWithFlags[C]) BuilderWithFlags[Collector] {
	return func(app *kingpin.Application) Collector {
		return fn(app)
	}
}

var BuildersWithFlags = map[string]BuilderWithFlags[Collector]{
	ad.Name:                NewBuilderWithFlags(ad.NewWithFlags),
	adcs.Name:              NewBuilderWithFlags(adcs.NewWithFlags),
	adfs.Name:              NewBuilderWithFlags(adfs.NewWithFlags),
	cache.Name:             NewBuilderWithFlags(cache.NewWithFlags),
	container.Name:         NewBuilderWithFlags(container.NewWithFlags),
	cpu.Name:               NewBuilderWithFlags(cpu.NewWithFlags),
	cpu_info.Name:          NewBuilderWithFlags(cpu_info.NewWithFlags),
	cs.Name:                NewBuilderWithFlags(cs.NewWithFlags),
	dfsr.Name:              NewBuilderWithFlags(dfsr.NewWithFlags),
	dhcp.Name:              NewBuilderWithFlags(dhcp.NewWithFlags),
	diskdrive.Name:         NewBuilderWithFlags(diskdrive.NewWithFlags),
	dns.Name:               NewBuilderWithFlags(dns.NewWithFlags),
	exchange.Name:          NewBuilderWithFlags(exchange.NewWithFlags),
	filetime.Name:          NewBuilderWithFlags(filetime.NewWithFlags),
	fsrmquota.Name:         NewBuilderWithFlags(fsrmquota.NewWithFlags),
	hyperv.Name:            NewBuilderWithFlags(hyperv.NewWithFlags),
	iis.Name:               NewBuilderWithFlags(iis.NewWithFlags),
	license.Name:           NewBuilderWithFlags(license.NewWithFlags),
	logical_disk.Name:      NewBuilderWithFlags(logical_disk.NewWithFlags),
	logon.Name:             NewBuilderWithFlags(logon.NewWithFlags),
	memory.Name:            NewBuilderWithFlags(memory.NewWithFlags),
	mscluster.Name:         NewBuilderWithFlags(mscluster.NewWithFlags),
	msmq.Name:              NewBuilderWithFlags(msmq.NewWithFlags),
	mssql.Name:             NewBuilderWithFlags(mssql.NewWithFlags),
	net.Name:               NewBuilderWithFlags(net.NewWithFlags),
	netframework.Name:      NewBuilderWithFlags(netframework.NewWithFlags),
	nps.Name:               NewBuilderWithFlags(nps.NewWithFlags),
	os.Name:                NewBuilderWithFlags(os.NewWithFlags),
	pagefile.Name:          NewBuilderWithFlags(pagefile.NewWithFlags),
	perfdata.Name:          NewBuilderWithFlags(perfdata.NewWithFlags),
	physical_disk.Name:     NewBuilderWithFlags(physical_disk.NewWithFlags),
	printer.Name:           NewBuilderWithFlags(printer.NewWithFlags),
	process.Name:           NewBuilderWithFlags(process.NewWithFlags),
	remote_fx.Name:         NewBuilderWithFlags(remote_fx.NewWithFlags),
	scheduled_task.Name:    NewBuilderWithFlags(scheduled_task.NewWithFlags),
	service.Name:           NewBuilderWithFlags(service.NewWithFlags),
	smb.Name:               NewBuilderWithFlags(smb.NewWithFlags),
	smbclient.Name:         NewBuilderWithFlags(smbclient.NewWithFlags),
	smtp.Name:              NewBuilderWithFlags(smtp.NewWithFlags),
	system.Name:            NewBuilderWithFlags(system.NewWithFlags),
	tcp.Name:               NewBuilderWithFlags(tcp.NewWithFlags),
	terminal_services.Name: NewBuilderWithFlags(terminal_services.NewWithFlags),
	textfile.Name:          NewBuilderWithFlags(textfile.NewWithFlags),
	thermalzone.Name:       NewBuilderWithFlags(thermalzone.NewWithFlags),
	time.Name:              NewBuilderWithFlags(time.NewWithFlags),
	udp.Name:               NewBuilderWithFlags(udp.NewWithFlags),
	update.Name:            NewBuilderWithFlags(update.NewWithFlags),
	vmware.Name:            NewBuilderWithFlags(vmware.NewWithFlags),
        rdp_client.Name:        NewBuilderWithFlags(rdp_client.NewWithFlags),
}

func Available() []string {
	return slices.Sorted(maps.Keys(BuildersWithFlags))
}
