package dominio

import (
	"index/enron/models"
	"index/enron/service"
	"io/ioutil"
	"log"
	"strings"
	//"sync"
)

var global = 1
var path = "./maildir/"
var count = 0
var writing = 50000

func Search() {
	//defer wg.Done()
	folders, err := ioutil.ReadDir(path)
	if err != nil {
		log.Fatal(err)
	}
	for i, folder := range folders {

		i++
		if folder.IsDir() {
			path = path + folder.Name() + "/"
			global++
			Search()
		} else {
			count++
			if count > 378707 {
				result := ReadFile(path, folder, count)
				if result.Message_ID != "Error" {
					data := convertJsonToString(result)
					service.Post(data, path, folder.Name(), count)
				}
			}
		}
		if i == len(folders) {
			pathArr := strings.Split(path, "/")
			path = ""
			for i, pathC := range pathArr {
				i++
				if i <= global {
					path += pathC + "/"
				}
			}
			global--

		}

	}
}

func convertJsonToString(jDta models.Email) string {
	result := "{ \"index\" : { \"_index\" : \"enron_mail\" } } \n"
	result += `{"message_Id":"` + jDta.Message_ID + `","date":"` + jDta.Date + `","from":"` + jDta.From +
		`","to":"` + jDta.To + `","subject":"` + jDta.Subject + `","mime_version":"` + jDta.Mime_Version +
		`","content_type":"` + jDta.Content_Type + `","content_transfer_encoding":"` + jDta.Content_Transfer_Encoding +
		`","x_from":"` + jDta.X_From + `","x_to":"` + jDta.X_To + `","x_cc":"` + jDta.X_cc + `","x_bcc":"` + jDta.X_bcc +
		`","x_folder":"` + jDta.X_Folder + `","x_origin":"` + jDta.X_Origin + `","x_filename":"` + jDta.X_FileName +
		`","message":"` + jDta.Message + `"}`

	return result
}
