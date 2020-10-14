package encryptionhttp

//Request request json to get the flying status
type Request struct {
	Data string `json:"data"`
}

//RequestByte request json to get the flying status
type RequestByte struct {
	Data []byte `json:"data"`
}
