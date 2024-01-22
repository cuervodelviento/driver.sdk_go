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
