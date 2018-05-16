package stores

type Store interface {
	SetUp()
	Close()
	QueryServices(name string) ([]QService, error)
	QueryAuthMethods(service string, host string, typeName string) ([]map[string]interface{}, error)
	QueryBindingsByDN(dn string, host string) ([]QBinding, error)
	QueryBindings(service string, host string) ([]QBinding, error)
	InsertService(name string, hosts []string, authTypes []string, authMethod string, retrievalField string, createdOn string) (QService, error)
	InsertBinding(name string, service string, host string, dn string, oidcToken string, uniqueKey string) (QBinding, error)
	UpdateBinding(original QBinding, updated QBinding) (QBinding, error)
}
