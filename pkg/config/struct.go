package config

type SideBarItem struct {
	Name string
	Path string
}

type Config struct {
	SiteName     string
	CodeTheme    string
	SidebarItems []SideBarItem
	SiteLogo     string
	Footer       string
}
