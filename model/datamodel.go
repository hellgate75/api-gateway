package model

type HttpResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type Result struct {
	Process string `json:"name"`
	Content string `json:"content"`
}

type Error struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

type JSonResult struct {
	Name string                 `json:"name"`
	Body map[string]interface{} `json:"content"`
}

type GenericModel struct {
	Content string
}

type Site struct {
	Name      string `json:"site"`
	Address   string `json:"address"`
	Port      int    `json:"port"`
	Scheme    string `json:"scheme"`
	Type      string `json:"type"`
	Override  bool   `json:"override"`
	APIUri    string `json:"apiuri"`
	Protocol  string `json:"protocol"`
	Concat    bool   `json:"concatenatepath"`
	BeforeApi bool   `json:"concatenatebeforeapi"`
}

type IndexConfig struct {
	Enabled        bool   `json:"enabled"`
	Address        string `json:"address"`
	ServiceAddress string `json:"serviceaddress"`
	Port           int64  `json:"port"`
	Protocol       string `json:"protocol"`
	UseToken       bool   `json:"usetokenprotection"`
	SecurityToken  string `json:"securitytoken"`
	UseTLS         bool   `json:"usetls"`
	CACertFile   string `json:"cacertificatefile"`
	X509CertFile   string `json:"tlsx509certificatefile"`
	X509KeyFile    string `json:"tlsx509certificatekeyfile"`
}

type Configuration struct {
	Address       string `json:"ipaddress"`
	Port          int64  `json:"port"`
	APIUrl        string `json:"apiurl"`
	Concat        bool   `json:"concatenate"`
	BeforeApi     bool   `json:"beforeapi"`
	File          string `json:"servicefile"`
	Protocol      string `json:"protocol"`
	User          string `json:"user"`
	Password      string `json:"password"`
	UseToken      bool   `json:"usetokenprotection"`
	SecurityToken string `json:"securitytoken"`
	UseTLS        bool   `json:"usetls"`
	CACertFile   string `json:"cacertificatefile"`
	X509CertFile  string `json:"tlsx509certificatefile"`
	X509KeyFile   string `json:"tlsx509certificatekeyfile"`
}

type Response struct {
	Name    string `json:"name"`
	Type    string `json:"type"`
	SiteObj Site   `json:"-"`
}

type Balancer struct {
	Enabled   bool
	Valid     bool
	Sites     []Response
	QueryText string
	Diff      int
}
