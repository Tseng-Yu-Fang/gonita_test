package gonita

import "github.com/go-resty/resty/v2"

const (
	URI_BPM = "API/bpm/"
)

// TODO: Need chanage to Variables of Environment
var (
	user_ppassword = "12345"
	server_ip_port = "54.169.182.165:8080"
)

// sources := fmt.Sprintf(server_addr,
//  // os.Getenv("BPM_SERVER_ADDR"),
//  os.Getenv("b.server"),
// )

//
//  FormInput
//  @Description: CRUD必要的 JSON 外層結構
//
type FormInput struct {
	ModelInput *interface{} `json:"modelInput,omitempty"`
}

type BPMClient struct {
	serverUri  string
	apiUri     string
	username   string
	password   string
	request    *resty.Request
	token      string
	jSessionId string // JSESSIONID
}
