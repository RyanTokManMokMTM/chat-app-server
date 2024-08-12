package types

type NotificationMessage struct {
	To      string
	From    string
	Content string
}

func (nm *NotificationMessage) GetToID() string {
	return nm.To
}

func (nm *NotificationMessage) GetFormID() string {
	return nm.From
}

func (nm *NotificationMessage) GetContent() string {
	return nm.Content
}
