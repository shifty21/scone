package encryptionservice

//Request request json to get the flying status
type Request struct {
	Data string `json:"data"`
}

//Request request json to get the flying status
type RequestByte struct {
	Data []byte `json:"data"`
}
