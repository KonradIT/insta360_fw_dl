package cmd

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"path"

	"github.com/fatih/color"
	"github.com/konradit/insta360_fw_dl/pkg/insta360"
)

func RunDownloader(camera insta360.Camera) error {
	client := &http.Client{}
	req, err := http.NewRequest("GET", fmt.Sprintf("https://openapi.insta360.com/website/appDownload/getGroupApp?group=%s&X-Language=en-us", camera), nil)
	if err != nil {
		log.Fatal(err)
	}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	var response = &insta360.FirmwareDownloadList{}
	err = json.NewDecoder(resp.Body).Decode(response)
	if err != nil {
		log.Fatal(err)
	}
	rolling := 0
	var itemList []insta360.Item

	for _, firmware := range response.Data.Apps {
		color.Cyan(">>> %s", firmware.Name)
		for _, item := range firmware.Items {
			rolling++
			itemList = append(itemList, item)
			color.Yellow("\t>>> [%d] - %s (Platform: %s UpdateTime: %s)", rolling, item.Version, item.Platform, item.UpdateTime)
		}
	}

	fmt.Printf(">>> ")
	var choice int
	_, err = fmt.Scanln(&choice)
	if err != nil {
		return err
	}

	_, file := path.Split(itemList[choice-1].Channels[0].DownloadURL)
	insta360.DownloadFile(file, itemList[choice-1].Channels[0].DownloadURL)
	return nil
}
