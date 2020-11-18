package cas

//FailRequest in case of anything else than 200 req status
type FailRequest struct {
	Message []string `json:"msg,omitempty"`
}

//GetRequest gets session data
type GetRequest struct {
	Session string `json:"session,omitempty"`
}

//PostRequest gets session data
type PostRequest struct {
}

//Services struct for session
type Services struct {
	Name        string   `yaml:"name,omitempty"`
	ImageName   string   `yaml:"image_name,omitempty"`
	MREnclaves  []string `yaml:"mrenclaves,omitempty"`
	Command     string   `yaml:"command,omitempty"`
	PWD         string   `yaml:"pwd,omitempty"`
	Environment struct {
		SconeMode string `yaml:"SCONE_MODE,omitempty"`
	} `yaml:"environment,omitempty"`
}

//Security struct for session
type Security struct {
	Attestation struct {
		Tolerate         []string `yaml:"tolerate,omitempty"`
		IgnoreAdvisories string   `yaml:"ignore_advisories,omitempty"`
		Mode             string   `yaml:"mode,omitempty"`
	} `yaml:"attestation,omitempty"`
}

//Images struct for session
type Images struct {
	Name           string `yaml:"name,omitempty"`
	InjectionFiles []struct {
		Path    string `yaml:"path,omitempty"`
		Content string `yaml:"content,omitempty"`
	} `yaml:"injection_files,omitempty"`
}

//Secrets for secrets
type Secrets struct {
	Name         string `yaml:"name,omitempty"`
	Kind         string `yaml:"kind,omitempty"`
	ExportPublic bool   `yaml:"export_public,omitempty"`
	Value        string `yaml:"value,omitempty"`
	Export       []struct {
		Session     string `yaml:"value,omitempty"`
		SessionHash string `yaml:"session_hash,omitempty"`
	} `yaml:"export,omitempty"`
}

//AccessPolicy policies
type AccessPolicy struct {
	Read   []string `yaml:"read,omitempty"`
	Update []string `yaml:"update,omitempty"`
}

//SessionYAML content
type SessionYAML struct {
	Version      string       `yaml:"version,omitempty"`
	Name         string       `yaml:"name,omitempty"`
	Predecessor  string       `yaml:"predecessor,omitempty"`
	Services     []Services   `yaml:"services,omitempty"`
	Images       []Images     `yaml:"images,omitempty"`
	Security     Security     `yaml:"security,omitempty"`
	Secrets      []Secrets    `yaml:"secrets,omitempty"`
	Creator      string       `yaml:"creator,omitempty"`
	AccessPolicy AccessPolicy `yaml:"access_policy,omitempty"`
}
