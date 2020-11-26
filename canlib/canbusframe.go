package canlib

// CanFrame - Structure of canbus frame
type CanFrame struct {
	// CAN averts message/data collisions by using the arbitration ID of the node,
	// i.e. the message with the highest priority (= lowest arbitration ID) will gain access to the bus,
	// while all other nodes (with lower priority arbitration IDs) switch to a “listening” mode.
	ArbitrationId uint32
	Dlc           byte
	// Data bytes
	Data []byte
	// Standard - 2048 different messages (2^11)
	// Extended - 536+ different messages (2^29)
	Extended bool
}
