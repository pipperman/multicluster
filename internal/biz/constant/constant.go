package constant

type ClusterType string

const (
	Private  ClusterType = "Private"
	Managed  ClusterType = "Managed"
	External ClusterType = "External"
	CSK      ClusterType = "CSK"
)

type SupportClusterVersion string

const (
	V_1_18_20 SupportClusterVersion = "v1.18.20"
	V_1_20_2  SupportClusterVersion = "v1.20.2"
)
