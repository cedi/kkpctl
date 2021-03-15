package model

// Version is the type that the KKP API endpoint /api/v1/upgrades/cluster returns as an array
type Version struct {
	Version string `json:"version"`
	Default bool   `json:"default,omitempty"`
}

// VersionList is a slice of Version that implements the sort interface
type VersionList []Version

// Len returns the length of a VersionList
func (vl VersionList) Len() int {
	return len(vl)
}

// Less returns true if the left Version is bigger than the right Version object
//	Note: This is inversed to what is considered "normal", since we always want to display the highest Version first
func (vl VersionList) Less(i, j int) bool {
	return vl[i].Version > vl[j].Version
}

// Swap swaps two Items in the array
func (vl VersionList) Swap(i, j int) {
	vl[i], vl[j] = vl[j], vl[i]
}
