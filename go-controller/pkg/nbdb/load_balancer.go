// Code generated by "libovsdb.modelgen"
// DO NOT EDIT.

package nbdb

type (
	LoadBalancerProtocol        = string
	LoadBalancerSelectionFields = string
)

const (
	LoadBalancerProtocolTCP           LoadBalancerProtocol        = "tcp"
	LoadBalancerProtocolUDP           LoadBalancerProtocol        = "udp"
	LoadBalancerProtocolSCTP          LoadBalancerProtocol        = "sctp"
	LoadBalancerSelectionFieldsEthSrc LoadBalancerSelectionFields = "eth_src"
	LoadBalancerSelectionFieldsEthDst LoadBalancerSelectionFields = "eth_dst"
	LoadBalancerSelectionFieldsIPSrc  LoadBalancerSelectionFields = "ip_src"
	LoadBalancerSelectionFieldsIPDst  LoadBalancerSelectionFields = "ip_dst"
	LoadBalancerSelectionFieldsTpSrc  LoadBalancerSelectionFields = "tp_src"
	LoadBalancerSelectionFieldsTpDst  LoadBalancerSelectionFields = "tp_dst"
)

// LoadBalancer defines an object in Load_Balancer table
type LoadBalancer struct {
	UUID            string                        `ovsdb:"_uuid"`
	ExternalIDs     map[string]string             `ovsdb:"external_ids"`
	HealthCheck     []string                      `ovsdb:"health_check"`
	IPPortMappings  map[string]string             `ovsdb:"ip_port_mappings"`
	Name            string                        `ovsdb:"name"`
	Options         map[string]string             `ovsdb:"options"`
	Protocol        []LoadBalancerProtocol        `ovsdb:"protocol"`
	SelectionFields []LoadBalancerSelectionFields `ovsdb:"selection_fields"`
	Vips            map[string]string             `ovsdb:"vips"`
}
