package provider

// JailedStatus enum that represents jailed status
type JailedStatus int

const (
	// Jailed status is when a node has been jailed due to missing a determined amount of blocks and/or byzantine behavior and thus cannot serve relays nor participate in consensus
	Jailed JailedStatus = iota + 1
	// Unjailed status is when a node is not jailed and thus can serve relays
	Unjailed
)

// StakingStatus enum that represents staking status
type StakingStatus int

const (
	// Unstaked represents unstaked status
	Unstaked StakingStatus = iota
	// Unstaking represents unstaking status
	Unstaking
	// Staked represents staked status
	Staked
)