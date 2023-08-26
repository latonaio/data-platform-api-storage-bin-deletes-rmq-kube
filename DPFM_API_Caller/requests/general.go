package requests

type General struct {
	BusinessPartner     int    `json:"BusinessPartner"`
	Plant               string `json:"Plant"`
	StorageLocation     string `json:"StorageLocation"`
	StorageBin          string `json:"StorageBin"`
	IsMarkedForDeletion *bool  `json:"IsMarkedForDeletion"`
}
