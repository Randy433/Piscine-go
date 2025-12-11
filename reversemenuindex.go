package piscine

func ReverseMenuIndex(menu []string) []string {
	menuIndex := make([]string, len(menu))
	for i, v := range menu {
		menuIndex[(len(menu)-1)-i] = v
	}
	return menuIndex
}
