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
)

var LocationIDs = []string{
	"ams01.ilandcloud.com",
	"ams04.ilandcloud.com",
	"dal02.ilandcloud.com",
	"dal06.ilandcloud.com",
	"dal22.ilandcloud.com",
	"dal23.ilandcloud.com",
	"dal25.ilandcloud.com",
	"lax01.ilandcloud.com",
	"lon02.ilandcloud.com",
	"man01.ilandcloud.com",
	"mel02.ilandcloud.com",
	"res01.ilandcloud.com",
	"sin01.ilandcloud.com",
	"str02.ilandcloud.com",
	"str03.ilandcloud.com",
	"str05.ilandcloud.com",
	"syd02.ilandcloud.com",
}
