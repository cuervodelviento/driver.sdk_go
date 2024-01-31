package config

// https://github.com/Netsocs-Team/DevDocs/blob/main/markdown/drivers/config-schemas/getChannels.md
type GetChannelResponseItem struct {
	ChannelNumber int    `json:"channelNumber"`
	ChannelName   string `json:"name"`
	Source        string `json:"rtspSource"`
}

// https://github.com/Netsocs-Team/DevDocs/blob/main/markdown/drivers/config-schemas/getChannels.md
type GetChannelResponse []*GetChannelResponseItem

// https://github.com/Netsocs-Team/DevDocs/blob/main/markdown/drivers/config-schemas/setUsers.md
type SetUserRequestItem struct {
	ID       int    `json:"id"`
	Used     bool   `json:"used"`
	UserName string `json:"userName"`
	Password string `json:"password"`
}

// https://github.com/Netsocs-Team/DevDocs/blob/main/markdown/drivers/config-schemas/setUsers.md
type SetUserRequest []*SetUserRequestItem

// https://github.com/Netsocs-Team/DevDocs/blob/main/markdown/drivers/config-schemas/setUsers.md
type SetUserResponse error

// https://github.com/Netsocs-Team/DevDocs/blob/main/markdown/drivers/config-schemas/deleteUser.md
type DeleteUserRequest struct {
	ID int `json:"id"`
}

// https://github.com/Netsocs-Team/DevDocs/blob/main/markdown/drivers/config-schemas/deleteUser.md
type DeleteUserResponse error

// https://github.com/Netsocs-Team/DevDocs/blob/main/markdown/drivers/config-schemas/getUsers.md
type GetUserRequest struct{}

// https://github.com/Netsocs-Team/DevDocs/blob/main/markdown/drivers/config-schemas/getUsers.md
type GetUserResponse struct {
	ID       int    `json:"id"`
	Enabled  bool   `json:"enabled"`
	UserName string `json:"userName"`
}

// https://github.com/Netsocs-Team/DevDocs/blob/main/markdown/drivers/config-schemas/setAddPersonToAC.md
type SetAddPersonToACRequest struct {
	PersonID   string `json:"personId"`
	PersonName string `json:"personName"`
}

type SetAddPersonToACResponse error

// https://github.com/Netsocs-Team/DevDocs/blob/main/markdown/drivers/config-schemas/setCardToPersonAC.md
type SetCardToPersonACRequest struct {
	PersonID string   `json:"personId"`
	Cards    []string `json:"cards"`
}

type SetCardToPersonACResponse error

// https://github.com/Netsocs-Team/DevDocs/blob/main/markdown/drivers/config-schemas/setFaceToPersonAC.md
type SetFaceToPersonACRequest struct {
	PersonID string   `json:"personId"`
	Faces    []string `json:"faces"`
}

type SetFaceToPersonACResponse error

// https://github.com/Netsocs-Team/DevDocs/blob/main/markdown/drivers/config-schemas/setDelPersonToAC.md

type SetDelToPersonToACRequest struct {
	PersonID string `json:"personId"`
}

type SetDelToPersonToACResponse error

// https://github.com/Netsocs-Team/DevDocs/blob/main/markdown/drivers/config-schemas/getAdminUser.md

type GetAdminUserRequest struct{}

type GetAdminUserResponse []*GetAdminUserResponseItem

type GetAdminUserResponseItem struct {
	Username string `json:"username"`
}

type GetRecordingSourceRequest struct {
	UTCStart string  `json:"utcStart"`
	UTCEnd   *string `json:"utcEnd"`
}

type GetRecordingSourceResponse struct {
	Source string `json:"source"`
}

type ZoomDirection int

const (
	ZoomIn  ZoomDirection = 0
	ZoomOut ZoomDirection = 1
)

type ActionZoomRequest struct {
	ChannelNumber int           `json:"channelNumber"`
	ZoomDirection ZoomDirection `json:"zoomDirection"` // 0: zoom in, 1: zoom out
}

type ActionZoomResponse error

type ActionMoveRequestDirection int

const (
	MoveUp    ActionMoveRequestDirection = 0
	MoveDown  ActionMoveRequestDirection = 1
	MoveLeft  ActionMoveRequestDirection = 2
	MoveRight ActionMoveRequestDirection = 3
)

type ActionMoveRequest struct {
	ChannelNumber int                        `json:"channelNumber"`
	Direction     ActionMoveRequestDirection `json:"direction"`
}

type ActionMoveResponse error
