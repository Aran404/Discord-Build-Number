package build

import (
	client "github.com/Aran404/Discord-Build-Number/Core/Client"
	"encoding/base64"
	"fmt"
	"log"
	"regexp"
	"strings"
)

func (in *Instance) GetBuildNumber() (string, error) {
	resp := client.Request("GET", "https://discord.com/login", nil, client.NoHeaders, true, in.Client)

	if resp.Error != nil {
		return "", resp.Error
	}

	if !resp.Ok {
		return "", fmt.Errorf("could not get sesssion, status code: %v", resp.StatusCode)
	}

	jsFile := regexp.MustCompile(`<script src="\/assets\/([a-zA-z0-9]+)\.js`).FindAll(resp.Body, -1)
	jsFileParsed := strings.Split(string(jsFile[3]), "assets/")[1]
	buildNumber, err := in._GetBuildNumber(jsFileParsed)

	in.BuildNumber = buildNumber

	if err != nil {
		return "", err
	}

	return in.BuildNumber, nil
}

func (in *Instance) _GetBuildNumber(js string) (string, error) {
	request := client.Request("GET", "https://discord.com/assets/"+js, nil, client.NoHeaders, true, in.Client)

	if request.Error != nil {
		return "", request.Error
	}

	if !request.Ok {
		return "", fmt.Errorf("could not get build number, Status Code: %v", request.StatusCode)
	}

	buildNumber := strings.Split(strings.Split(string(request.Body), `Build Number: ").concat("`)[1], `"`)[0]

	return buildNumber, nil
}

func (in *Instance) GetSuperProperties() string {
	if in.BuildNumber == "" {
		log.Fatal("Build number not initialized")
	}

	parsed := `{"os":"Windows","browser":"Chrome","device":"","system_locale":"en-US","browser_user_agent":"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/112.0.0.0 Safari/537.36","browser_version":"112.0.0.0","os_version":"10","referrer":"","referring_domain":"","referrer_current":"","referring_domain_current":"","release_channel":"stable","client_build_number":%S,"client_event_source":null,"design_id":0}`
	return base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf(parsed, in.BuildNumber)))
}
