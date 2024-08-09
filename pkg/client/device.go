package client

type Device struct {
	Username        string                 `json:"username"`
	Password        string                 `json:"password"`
	IpAddressPublic string                 `json:"ipAddressPublic"`
	Port            int                    `json:"port"`
	ID              string                 `json:"id"`
	IDModel         int                    `json:"idModel"`
	ChildID         string                 `json:"childId"`
	IDBrand         int                    `json:"idBrand"`
	IDManufacturer  int                    `json:"idManufacturer"`
	IDDeviceGroup   int                    `json:"idDeviceGroup"`
	IDSubSystem     int                    `json:"idSubSystem"`
	Params          map[string]interface{} `json:"params"`
}
