package types

const (
	// ModuleName defines the module name
	ModuleName = "deweb"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// RouterKey is the message route for slashing
	RouterKey = ModuleName

	// QuerierRoute defines the module's query routing key
	QuerierRoute = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_deweb"

	// RecordsKey defines prefix for users storage. Here we get record by UUID
	RecordsKey = "records-"

	// UsersRecords defines prefix for storage where we store mapping between user and list of records uuids
	UsersRecords = "users-"
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}
