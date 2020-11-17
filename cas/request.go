package cas

//FailRequest in case of anything else than 200 req status
type FailRequest struct {
	Message []string `json:"msg"`
}

//GetRequest gets session data
type GetRequest struct {
	Session string `json:"session"`
}

//PostRequest gets session data
type PostRequest struct {
}
