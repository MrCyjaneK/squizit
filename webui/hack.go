package webui

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

func realHackOk(gamecode string, key string) {
	answers = Response{}
	url := Host + "/api5.php?code=" + gamecode + "&version=" + version + "&key=" + key
	resp, err := httpClient.Get(url)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()
	out, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	err = json.Unmarshal(out, &answers)
	if err != nil {
		fmt.Println(err)
		return
	}
}
