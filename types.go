package strawpoll

// Poll is a structure that contains the generic responses that
// are returned by StrawPoll's server
type Poll struct {
	ID       int      `json:"id"`
	Title    string   `json:"title"`
	Options  []string `json:"options"`
	Votes    []int    `json:"votes,omitempty"`
	Multi    bool     `json:"multi"`
	Dupcheck string   `json:"dupcheck"`
	Captcha  bool     `json:"captcha"`
}
