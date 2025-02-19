// Code generated by "libovsdb.modelgen"
// DO NOT EDIT.

package sbdb

// PortBinding defines an object in Port_Binding table
type PortBinding struct {
	UUID           string            `ovsdb:"_uuid"`
	Chassis        []string          `ovsdb:"chassis"`
	Datapath       string            `ovsdb:"datapath"`
	Encap          []string          `ovsdb:"encap"`
	ExternalIDs    map[string]string `ovsdb:"external_ids"`
	GatewayChassis []string          `ovsdb:"gateway_chassis"`
	HaChassisGroup []string          `ovsdb:"ha_chassis_group"`
	LogicalPort    string            `ovsdb:"logical_port"`
	MAC            []string          `ovsdb:"mac"`
	NatAddresses   []string          `ovsdb:"nat_addresses"`
	Options        map[string]string `ovsdb:"options"`
	ParentPort     []string          `ovsdb:"parent_port"`
	Tag            []int             `ovsdb:"tag"`
	TunnelKey      int               `ovsdb:"tunnel_key"`
	Type           string            `ovsdb:"type"`
	Up             []bool            `ovsdb:"up"`
	VirtualParent  []string          `ovsdb:"virtual_parent"`
}
