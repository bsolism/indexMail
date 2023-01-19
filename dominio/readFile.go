package dominio

import (
	"fmt"
	"index/enron/models"
	"os"
	"strings"
)

func ReadFile(path string, archivo os.FileInfo, count int) models.Email {
	var email models.Email
	contents, err := os.ReadFile(path + archivo.Name())
	if err != nil {
		fmt.Println("Error en directorio")
		fmt.Println(archivo.Name())
		fmt.Println(path)
		fmt.Println(err.Error())
		email.Message_ID = "Error"
		return email
	}
	content := strings.ReplaceAll(string(contents), `\`, "/")
	content = strings.ReplaceAll(content, "\"", "'")
	contentSplit := strings.Split(content, "\n")
	message := ""
	for i, dta := range contentSplit {
		if i == 0 {
			email.Message_ID = assignProp(dta, "Message-ID:")
		} else if i == 1 {
			email.Date = assignProp(dta, "Date:")
		} else if i == 2 {
			email.From = assignProp(dta, "From:")
		} else if i == 3 {
			email.To = assignProp(dta, "To:")
		} else if i == 4 {
			email.Subject = assignProp(dta, "Subject:")
		} else if i == 5 {
			email.Mime_Version = assignProp(dta, "Mime-Version:")
		} else if i == 6 {
			email.Content_Type = assignProp(dta, "Content-Type:")
		} else if i == 7 {
			email.Content_Transfer_Encoding = assignProp(dta, "Content-Transfer-Encoding:")
		} else if i == 8 {
			email.X_From = assignProp(dta, "X-From:")
		} else if i == 9 {
			email.X_To = assignProp(dta, "X-To:")
		} else if i == 10 {
			email.X_cc = assignProp(dta, "X-cc:")
		} else if i == 11 {
			email.X_bcc = assignProp(dta, "X-bcc:")
		} else if i == 12 {
			email.X_Folder = assignProp(dta, "X-Folder:")
		} else if i == 13 {
			email.X_Origin = assignProp(dta, "X-Origin:")
		} else if i == 14 {
			email.X_FileName = assignProp(dta, "X-FileName:")
		} else if i > 14 {
			message += dta
		}

	}
	message = strings.ReplaceAll(message, "\r", "")
	email.Message = message
	return email

}
func assignProp(dta string, nameProp string) string {
	data := ""
	if strings.Contains(dta, nameProp) {
		props := strings.Split(dta, nameProp)
		dataprop := ""
		for i, propDta := range props {

			if i > 0 {
				dataprop += propDta
			}
		}
		data = strings.ReplaceAll(dataprop, "\r", "")

	}
	return data
}
