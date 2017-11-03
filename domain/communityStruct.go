package domain

type Community struct {
	payload int
	content []byte
}
type CommunityOperations interface {
	ToByte() []byte
}

func (community Community) ToByte() []byte {
	payload := make([]byte, 1)
	payload[0] = byte(community.payload)
	return append(payload, community.content...)
}
