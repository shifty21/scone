package cas

//FailRequest in case of anything else than 200 req status
type FailRequest struct {
	Message []string `json:"msg,omitempty"`
}

//SessionResponse gets session data
type SessionResponse struct {
	Session string `json:"session,omitempty"`
}

//PostResponse gets session data
type PostResponse struct {
	Hash string `json:"hash,omitempty"`
}

//Service struct for session
type Service struct {
	Name        string            `yaml:"name,omitempty"`
	ImageName   string            `yaml:"image_name,omitempty"`
	MREnclaves  []string          `yaml:"mrenclaves,omitempty"`
	Command     string            `yaml:"command,omitempty"`
	PWD         string            `yaml:"pwd,omitempty"`
	Environment map[string]string `yaml:"environment,omitempty"`
}

//Security struct for session
type Security struct {
	Attestation struct {
		Tolerate         []string `yaml:"tolerate,omitempty"`
		IgnoreAdvisories string   `yaml:"ignore_advisories,omitempty"`
		Mode             string   `yaml:"mode,omitempty"`
	} `yaml:"attestation,omitempty"`
}

//InjectionFile for file injection
type InjectionFile struct {
	Path    string `yaml:"path,omitempty"`
	Content string `yaml:"content,omitempty"`
}

//Image struct for session
type Image struct {
	Name           string          `yaml:"name,omitempty"`
	InjectionFiles []InjectionFile `yaml:"injection_files,omitempty"`
}

//ExportTo struct specifies session to which secrets are exposed
type ExportTo struct {
	Session     string `yaml:"session,omitempty"`
	SessionHash string `yaml:"session_hash,omitempty"`
}

//Import struct specifies secrets imported from
type Import struct {
	Session string `yaml:"session,omitempty"`
	Secret  string `yaml:"secret,omitempty"`
}

//Secret for secrets
type Secret struct {
	Name         string     `yaml:"name,omitempty"`
	Kind         string     `yaml:"kind,omitempty"`
	ExportPublic bool       `yaml:"export_public,omitempty"`
	Value        string     `yaml:"value,omitempty"`
	Export       []ExportTo `yaml:"export,omitempty"`
	Import       Import     `yaml:"import,omitempty"`
	PrivateKey   string     `yaml:"private_key,omitempty"`
	Issuer       string     `yaml:"issuer,omitempty"`
	CommonName   string     `yaml:"common_name,omitempty"`
	Endpoint     string     `yaml:"endpoint,omitempty"`
	DNS          []string   `yaml:"dns,omitempty"`
	ValidFor     string     `yaml:"valid_for,omitempty"`
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
	Services     []Service    `yaml:"services,omitempty"`
	Images       []Image      `yaml:"images,omitempty"`
	Security     Security     `yaml:"security,omitempty"`
	Secrets      []Secret     `yaml:"secrets,omitempty"`
	AccessPolicy AccessPolicy `yaml:"access_policy,omitempty"`
}

//PredecessorHash for reading predecessor hash of client session
type PredecessorHash struct {
	PredecessorHash string `yaml:"predecessor_hash,omitempty"`
}
