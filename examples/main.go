package main

import (
	"encoding/json"
	"fmt"
	teleflow "github.com/teleflow/go-client"
	"io/ioutil"
	"net/http"
)

func main() {

	client, err := teleflow.NewClient(&teleflow.Config{
		Server: "https://teleflow-server",
		//Username: "username-here",
		//Password: "password-here",
		// or
		AccessToken: "...",
	})
	if err != nil {
		panic(err)
	}

	hasErr := func(apiErr *teleflow.ApiError, err error) bool {
		if err != nil {
			fmt.Printf("Error: %s\n", err)
			return true
		}
		if apiErr != nil {
			fmt.Printf("API Error: %d - %s\n", apiErr.Code, apiErr.Message)
			if apiErr.Code == 401 {
				println("Check token..")
			}
			return true
		}
		return false
	}

	println("Start..")

	tasks, apiErr, err := client.GetAutocallTasks()
	if hasErr(apiErr, err) {
		return
	}
	for _, z := range tasks {
		println("Z: ", z.Phone, " - ", z.Result)
	}

	//// Get campaigns
	//campaigns, err := client.GetCampaigns()
	//
	//if len(camaigns) > 0 {
	//	// see examples/web-hook-server
	//	err := client.SetWebHook(campaigns[0], "http://localhost:8090/web-hook-for-teleflow")
	//}
	//

	//records, apiErr, err := client.FindRecords(&teleflow.RecordOptions{
	//	//Src: "1234567890",
	//	//Dst: "111",
	//	UniqId: "1606179193.1457761",
	//	//TimeStart: time.Now().Add(-time.Hour * 24 * 360),
	//	//TimeEnd:   time.Now().Add(time.Minute * 5), //.Format(time.RFC3339),
	//})
	//if hasErr(apiErr, err) {
	//	return
	//}
	//for _, z := range records {
	//	println("Record: ", z.Src, z.Dst, z.RecordingFile, z.UniqId, z.CreatedAt.String())
	//}
	//
	//if len(records) > 0 {
	//
	//	filePath := "./file1.wav"
	//	apiErr, err = client.DownloadAudioFile(&teleflow.RecordOptions{
	//		UniqId: "1606179193.1457761",
	//	}, filePath)
	//	if hasErr(apiErr, err) {
	//		return
	//	}
	//	if _, err := os.Stat(filePath); err == nil {
	//	}
	//}

	// Set a webhook to receive an event when the status of a task changes
	_, apiErr, err = client.SetWebhook(1, "http://127.0.0.1:8899/web-hook-for-teleflow")
	if hasErr(apiErr, err) {
		return
	}

	task, apiErr, err := client.AddAutocallTask(&teleflow.AutocallTask{
		Phone:      "+199988888888",
		CampaignId: 1,
		Params: map[string]string{
			"name":   "Jon Doe",
			"amount": "8900",
		}})
	if hasErr(apiErr, err) {
		return
	}
	println("Created task: ", task.Id, task.Phone, task.IsCompleted)

	runWebHook()
}

func webhook(w http.ResponseWriter, req *http.Request) {

	fmt.Fprintf(w, "hello\n")
	content, err := ioutil.ReadAll(req.Body)
	if err != nil {
		println("Error: ", err)
		return
	}
	var task *teleflow.AutocallTask
	err = json.Unmarshal(content, task)
	println("Task: ", task)

}

func runWebHook() {

	println("Server running...")
	http.HandleFunc("/web-hook-for-teleflow", webhook)

	http.ListenAndServe(":8899", nil)

}
