package brute

type Result struct {
	Subdomain string   `json:"subdomain"`
	IPs       []string `json:"ips"`
}
