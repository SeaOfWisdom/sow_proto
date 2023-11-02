package proto

func (l *DetectLanguageResponse) IsValid() (out bool) {
	out = l.Code != "" && l.Lang != ""
	out = out && l.Code != "none" && l.Lang != "nonsense"

	return out
}
