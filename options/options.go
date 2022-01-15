package options

type Options struct {
	Command   string   `json:"command"`
	Arguments []string `json:"arguments"`
	Interval  int      `json:"interval"`
	Current   string   `json:"current"`
}
