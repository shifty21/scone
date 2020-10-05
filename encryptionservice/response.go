package encryptionservice

//Response json for session
type Response struct {
	//string bytes for data
	Data string `json:"encrypted_data"`
}
