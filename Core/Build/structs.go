package build

import client "github.com/Aran404/Discord-Build-Number/Core/Client"

type Instance struct {
	*client.Client
	Proxy       string
	UserAgent   string
	BuildNumber string
}
