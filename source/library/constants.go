package library

const (
	Version     string = "v0.0.1-alpha"
	ProjectName string = "yathahi"

	EntrySocketPath string = "/tmp/yathahi.sock"

	BinaryURL string = "https://raw.githubusercontent.com/yathahi/yios/main/packages/"

	BinaryFolder string = "/opt/yathahi/bin/"

	CoreServiceName string = "yathahi-core.service"

	EntryLogFile string = "/var/log/yathahi/entry.log"

	CorePackageName          string = "yathahi-core"
	WebPackageName           string = "web-engine"
	AdministratorPackageName string = "administrator-engine"
	ClientPackageName        string = "client-engine"
	DnsPackageName           string = "dns-engine"
)
