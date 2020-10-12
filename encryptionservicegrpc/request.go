package encryptionservicegrpc

//Request request json to get the flying status
type Request struct {
	Data string `json:"data"`
}

//Request request json to get the flying status
type RequestByte struct {
	Data []byte `json:"data"`
}

//Response json for session
type Response struct {
	//string bytes for data
	Data string `json:"data"`
}
