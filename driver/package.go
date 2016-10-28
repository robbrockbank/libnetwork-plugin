package driver

const (
	IPv4 = "IPv4"
	IPv6 = "IPv6"
	// Calico IPAM module does not allow selection of pools from which to allocate
	// IP addresses.  The pool ID, which has to be supplied in the libnetwork IPAM
	// API is therefore fixed.  We use different values for IPv4 and IPv6 so that
	// during allocation we know which IP version to use.
	PoolIDV4 = "CalicoPoolIPv4"
	PoolIDV6 = "CalicoPoolIPv6"

	// Fix pool and gateway CIDRs.  As per comment above, Calico IPAM does not allow
	// assignment from a specific pool, so we choose a dummy value that will not be
	// used in practise.  A 0/0 value is used for both IPv4 and IPv6.  This value is
	// also used by the Network Driver to indicate that the Calico IPAM driver was
	// used rather than the default libnetwork IPAM driver - this is useful because
	// Calico Network Driver behavior depends on whether our IPAM driver was used or
	// not.
	PoolCIDRV4    = "0.0.0.0/0"
	PoolCIDRV6    = "::/0"
	GatewayCIDRV4 = "0.0.0.0/0"
	GatewayCIDRV6 = "::/0"

	CalicoGlobalAddressSpace = "CalicoGlobalAddressSpace"

	IFPrefix = "cali"
)
