package proto

func StringToLogType(val string) LogType {
	out := LogType_UNKNOWN

	switch val {
	case "PurchasedPaper":
		out = LogType_PURCHASED_WORK
	case "ReviewerRewardsClaimed":
		out = LogType_CLAIM_REWARDS
	}

	return out
}
