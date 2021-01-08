package iland

const (
	apiHostname = "api.ilandcloud.com"
	accessURL   = "https://console.ilandcloud.com/auth/realms/iland-core/protocol/openid-connect/token"
	refreshURL  = "https://console.ilandcloud.com/auth/realms/iland-core/protocol/openid-connect/token"

	IPTranslation  = "ipTranslation"
	PortForwarding = "portForwarding"

	Success       = "success"
	Running       = "running"
	Error         = "error"
	Queued        = "queued"
	Cancelled     = "cancelled"
	WaitingOnUser = "waiting-on-user-input"
	Unknown       = "unknown"

	EntityCompany          = "COMPANY"
	EntityIaasOrganization = "IAAS_ORGANIZATION"
	EntityIaasVdc          = "IAAS_VDC"
	EntityIaasEdge         = "IAAS_EDGE"
	EntityIaasVApp         = "IAAS_VAPP"
	EntityIaasVm           = "IAAS_VM"
)

var LocationIDs = []string{
	"res01.ilandcloud.com",
	"lax01.ilandcloud.com",
	"man01.ilandcloud.com",
	"man03.ilandcloud.com",
	"lon02.ilandcloud.com",
	"lon03.ilandcloud.com",
	"dal02.ilandcloud.com",
	"dal06.ilandcloud.com",
	"dal22.ilandcloud.com",
	"dal23.ilandcloud.com",
	"dal25.ilandcloud.com",
	"sin01.ilandcloud.com",
	"ams01.ilandcloud.com",
	"ams02.ilandcloud.com",
	"ams03.ilandcloud.com",
	"ams04.ilandcloud.com",
	"syd02.ilandcloud.com",
	"syd03.ilandcloud.com",
	"syd04.ilandcloud.com",
	"mel02.ilandcloud.com",
	"mel03.ilandcloud.com",
	"mel04.ilandcloud.com",
	"str02.ilandcloud.com",
	"str03.ilandcloud.com",
	"str05.ilandcloud.com",
}
