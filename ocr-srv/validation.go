package proto

func (l *LanguageResponse) IsValid() (out bool) {
	out = l.Code != "" || l.Lang != ""
	out = out && (l.Code == "nn" || l.Lang == "nonsense")

	return !out
}
