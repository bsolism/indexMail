package models

type Email struct {
	Message_ID                string `json:"message_id"`
	Date                      string `json:"Date"`
	From                      string `json:"From"`
	To                        string `json:"To"`
	Subject                   string `json:"Subject"`
	Mime_Version              string `json:"Mime_Version"`
	Content_Type              string `json:"Content_Type"`
	Content_Transfer_Encoding string `json:"Content_Transfer_Encoding"`
	X_From                    string `json:"X-From"`
	X_To                      string `json:"X-To"`
	X_cc                      string `json:"X-cc"`
	X_bcc                     string `json:"X-bcc"`
	X_Folder                  string `json:"X-Folder"`
	X_Origin                  string `json:"X-Origin"`
	X_FileName                string `json:"X-FileName"`
	Message                   string `json:"Message"`
}
