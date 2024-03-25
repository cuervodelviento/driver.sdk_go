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
	ChannelNumber int     `json:"channelNumber"`
	UTCStart      string  `json:"utcStart"`
	UTCEnd        *string `json:"utcEnd"`
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

type GetAvailableOutputsRequest struct{}

type GetAvailableOutputsResponse []*GetAvailableOutputsResponseItem
type GetAvailableOutputsResponseItem struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	// 0 = relay, 1 = led
	Type int `json:"type"`
	// 0 = normalmente abierto, 1 = normalmente cerrado, 2 = encendido, 3 = apagado
	State int `json:"state"`
}

type GetAvailableSpeakersRequest struct{}
type GetAvailableSpeakersResponse []*GetAvailableSpeakersResponseItem
type GetAvailableSpeakersResponseItem struct {
	ID       string `json:"id"`
	Level    int    `json:"level"`
	MinLevel int    `json:"minLevel"`
	MaxLevel int    `json:"maxLevel"`
	Enabled  bool   `json:"enabled"`
	Name     string `json:"name"`
}

type GetStorageRequest struct{}
type GetStorageResponse []*GetStorageResponseItem
type GetStorageResponseItem struct {
	ID         string `json:"id"`
	DeviceName string `json:"deviceName"`
	Recording  bool   `json:"recording"`
	FreeSize   int    `json:"freeSize"`
	TotalSize  int    `json:"totalSize"`
	Status     int    `json:"status"`
}

type SetAdminUserRequest struct {
	Username     string `json:"username"`
	PrevPassword string `json:"prevPassword"`
	NewPassword  string `json:"newPassword"`
}

type SetAdminUserResponse error

type GetAdminUsersRequest struct{}

type GetAdminUsersResponse []*GetAdminUsersResponseItem
type GetAdminUsersResponseItem struct {
	Username string `json:"username"`
	// 0: Online, 1: Offline, 2: Unknown
	Status int `json:"status"`
	// Created at in format ISO 8601
	CreatedAt *string `json:"created_at"`
	// Last login in format ISO 8601
	LastLogin *string `json:"last_login"`
	// Last logout in format ISO 8601
	LastLogout *string `json:"last_logout"`
	LastIP     *string `json:"last_ip"`
}

type GetCurrentVideoResolutionByChannelRequest struct {
	ChannelNumber int `json:"channelNumber"`
}
type GetCurrentVideoResolutionByChannelResponse struct {
	Resolution string `json:"resolution"`
}
type GetAvailableVideoResolutionsRequest struct {
	ChannelNumber int `json:"channelNumber"`
}

type GetAvailableVideoResolutionsResponse struct {
	Resolutions []string `json:"resolutions"`
}

type SetVideoResolutionRequest struct {
	ChannelNumber int    `json:"channelNumber"`
	Resolution    string `json:"resolution"`
}

type SetVideoResolutionResponse error

type SetVideoInBlackAndWhiteRequest struct {
	ChannelNumber int  `json:"channelNumber"`
	Enabled       bool `json:"enabled"`
}

type SetVideoInBlackAndWhiteResponse error

type GetVideoInBlackAndWhiteStatusRequest struct {
	ChannelNumber int `json:"channelNumber"`
}

type GetVideoInBlackAndWhiteStatusResponse struct {
	Enabled bool `json:"enabled"`
}

type SetMirrorVideoRequest struct {
	ChannelNumber int  `json:"channelNumber"`
	Enabled       bool `json:"enabled"`
}

type SetMirrorVideoResponse error

type GetMirrorVideoStatusRequest struct {
	ChannelNumber int `json:"channelNumber"`
}

type GetMirrorVideoStatusResponse struct {
	Enabled bool `json:"enabled"`
}

type SetFlipVideoRequest struct {
	ChannelNumber int  `json:"channelNumber"`
	Enabled       bool `json:"enabled"`
}

type SetFlipVideoResponse error

type GetFlipVideoStatusRequest struct {
	ChannelNumber int `json:"channelNumber"`
}

type GetFlipVideoStatusResponse struct {
	Enabled bool `json:"enabled"`
}

type ActionPlayAudioClipRequest struct {
	ChannelNumber int    `json:"channelNumber"`
	SpeakerID     string `json:"speakerId"`
}

type GetHeatmapImageRequest struct {
	ChannelNumber int `json:"channelNumber"`
	// Start in format ISO 8601
	Start string `json:"start"`
	// End in format ISO 8601
	End string `json:"end"`
}

type GetHeatmapImageResponse struct {
	Filename string `json:"filename"`
}

type SetBlockPersonToACRequest struct {
	PersonID string `json:"personId"`
}

type SetBlockPersonToACResponse error

type DeleteAllPeopleACRequest struct{}

type DeleteAllPeopleACResponse error
type GetInputsResponseItem struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Enabled     bool   `json:"enabled"`
	Type        int    `json:"type"`
}

// https://github.com/Netsocs-Team/DevDocs/blob/main/markdown/drivers/config-schemas/getChannels.md
type GetInputsResponse []*GetInputsResponseItem

type GetMicrophoneResponseItem struct {
	Name    string `json:"name"`
	Enabled bool   `json:"enabled"`
	ID      string `json:"id"`
}

type GetMicrophoneResponse []*GetMicrophoneResponseItem

type GetFtpInfoRequest struct{}

type GetFtpInfoResponseItem struct {
	ID         string `json:"id"`
	Host       string `json:"host"`
	Port       int    `json:"port"`
	Username   string `json:"username"`
	Directory  string `json:"directory"`
	Encryption string `json:"encryption"`
	Mode       string `json:"mode"`
}
type GetFtpInfoResponse []*GetFtpInfoResponseItem

type SetFtpInfoRequest struct {
	ID         string `json:"id"`
	Host       string `json:"host"`
	Port       int    `json:"port"`
	Username   string `json:"username"`
	Password   string `json:"password"`
	Directory  string `json:"directory"`
	Encryption string `json:"encryption"`
	Mode       string `json:"mode"`
}

type SetFtpInfoResponse error
