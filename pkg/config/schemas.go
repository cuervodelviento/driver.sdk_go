package config

// https://github.com/Netsocs-Team/DevDocs/blob/main/markdown/drivers/config-schemas/getChannels.md
type GetChannelResponse []*struct {
	ChannelNumber int    `json:"channelNumber"`
	ChannelName   string `json:"name"`
	Source        string `json:"rtspSource"`
}

// https://github.com/Netsocs-Team/DevDocs/blob/main/markdown/drivers/config-schemas/setUsers.md
type SetUserRequest []*struct {
	ID       int    `json:"id"`
	Used     bool   `json:"used"`
	UserName string `json:"userName"`
	Password string `json:"password"`
}

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
