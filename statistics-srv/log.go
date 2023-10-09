package proto

func StringToLogType(val string) LogType {
	out := LogType_UNKNOWN

	switch val {
	case "PurchasePaper":
		out = LogType_PURCHASED_WORK
	case "ClaimReviewRewards":
		out = LogType_CLAIM_REWARDS
	}

	return out
}
