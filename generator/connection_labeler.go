package generator

import (
	"strings"
)

/**
 * Labels the connections
 */
func generateConnectionLabel(domain string, tier int, serverIdx int) string {
	return convertProtonVpnDomainToName(domain)
}

/**
 * Converts the domain into a label
 */
func convertProtonVpnDomainToName(domain string) string {
	x := strings.Replace(domain, "-", "_", -1)
	y := strings.Replace(x, ".", "_", -1)
	return y
}
