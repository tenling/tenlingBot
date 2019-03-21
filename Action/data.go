package Action

var (
	conanBank         = "bank.conan.site"
	addKeywordSuccess = "我學會了喵！"
	addKeywordFail    = "這個字有人教過我了喵～"
	replyAirplanFail  = "喵喵喵，他什麼時候出門我不知道喔～～"
	unknowMessage     = "喵喵喵～"
	tenling           = "tenling"

	autoRelyMap = map[string]string{
		"bank":      conanBank,
		"conanBank": conanBank,
		"帳本":        conanBank,
		"喵":         "",
	}

	airplanMap = map[string][]string{}
)
