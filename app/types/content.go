package types

type SocialLink struct {
	Label 	string
	Href  	string
  Icon  	string
	Handler string
}

type MenuItem struct {
	Label   string
	Href    string
	SubMenu []MenuItem
	Icon    string
}

type ContactInfo struct {
	Name 		string
  Email		string
  Phone		string
}