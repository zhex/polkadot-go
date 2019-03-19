package types

type Health struct {
	Peers           int  `mapstructure:"peers"`
	IsSyncing       bool `mapstructure:"isSyncing"`
	ShouldHavePeers bool `mapstructure:"shouldHavePeers"`
}
