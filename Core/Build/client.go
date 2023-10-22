package build

import (
	client "github.com/Aran404/Discord-Build-Number/Core/Client"
	"fmt"
	"strings"
)

func New(proxy string) (*Instance, error) {
	c := client.NewOptions()
	c.SetNewCookieJar()
	c.SetTimeout(60)

	httpClient, err := c.NewClient()

	if err != nil {
		return nil, err
	}

	if strings.Contains(proxy, "@") {
		subs := strings.Split(proxy, "@")
		proxy = fmt.Sprintf("%v:%v", subs[1], subs[0])
	}

	client := &client.Client{HttpClient: *httpClient}

	return &Instance{
		Client: client,
		Proxy:  proxy,
	}, nil
}
