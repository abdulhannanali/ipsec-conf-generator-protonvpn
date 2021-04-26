package generator

import (
	"fmt"
	"strings"
)

/**
 * Labels the connections
 */
func generateConnectionLabel(domain string, tier int, serverIdx int) string {
	domainName := convertProtonVpnDomainToName(domain)

	return fmt.Sprintf(
		"%s_tier_%d_server_idx_%d",
		domainName,
		tier,
		serverIdx,
	)
}

/**
 * Converts the domain into a label
 */
func convertProtonVpnDomainToName(domain string) string {
	x := strings.Replace(domain, "-", "_", -1)
	y := strings.Replace(x, ".", "_", -1)
	return y
}
