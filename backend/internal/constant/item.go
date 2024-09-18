package constant

type ItemStatus string

const (
	ItemPendingStatus ItemStatus = "PENDING"
	ItemAPPROVEDStatus ItemStatus = "APPROVED"
	ItemREJECTEDStatus ItemStatus = "REJECTED"
)