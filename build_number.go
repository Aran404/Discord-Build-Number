package discord_build_number

import (
	"encoding/base64"
	"fmt"
	"log"
	"regexp"
	"strings"
	"sync"
)

func (in *Instance) GetBuildNumber() (string, error) {
	resp := Request("GET", "https://discord.com/login", nil, NoHeaders, true, in.Client)
	if resp.Error != nil {
		return "", resp.Error
	}

	if !resp.Ok {
		return "", fmt.Errorf("could not get sesssion, status code: %v", resp.StatusCode)
	}

	jsFile := regexp.MustCompile(`<script defer src="\/assets\/([a-zA-Z0-9]+\.)?([a-zA-Z0-9]+)\.js`).FindAll(resp.Body, -1)

	var wg sync.WaitGroup
	respCh := make(chan string, 1)
	for _, v := range jsFile {
		wg.Add(1)
		go func(js string) {
			defer wg.Done()
			jsFileParsed := strings.Split(string(js), "assets/")[1]

			in._GetBuildNumber(jsFileParsed, respCh)
		}(string(v))
	}

	go func() {
		wg.Wait()
		close(respCh)
	}()

	in.BuildNumber = <-respCh
	return in.BuildNumber, nil
}

func (in *Instance) _GetBuildNumber(js string, r chan string) {
	request := Request("GET", "https://discord.com/assets/"+js, nil, NoHeaders, true, in.Client)
	if request.Error != nil {
		return
	}

	if !request.Ok {
		return
	}

	request.Body = []byte(strings.ReplaceAll(string(request.Body), " ", ""))
	if !strings.Contains(string(request.Body), `buildNumber:"`) {
		return
	}

	buildNumber := strings.Split(strings.Split(string(request.Body), `buildNumber:"`)[1], `"`)[0]
	r <- buildNumber
}

func (in *Instance) GetSuperProperties() string {
	if in.BuildNumber == "" {
		log.Fatal("Build number not initialized")
	}

	parsed := `{"os":"Windows","browser":"Chrome","device":"","system_locale":"en-CA","browser_user_agent":"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/117.0.0.0 Safari/537.36","browser_version":"117.0.0.0","os_version":"10","referrer":"","referring_domain":"","referrer_current":"","referring_domain_current":"","release_channel":"stable","client_build_number":%s,"client_event_source":null,"design_id":0}`
	return base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf(parsed, in.BuildNumber)))
}
