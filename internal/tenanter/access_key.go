package tenanter

type AccessKeyTenant struct {
	name   string
	id     string
	secret string
}

func NewTenantWithAccessKey(name, accessKeyId, accessKeySecret string) Tenanter {
	return &AccessKeyTenant{
		name:   name,
		id:     accessKeyId,
		secret: accessKeySecret,
	}
}

func (tenant *AccessKeyTenant) AccountName() string {
	return tenant.name
}

func (tenant *AccessKeyTenant) Clone() Tenanter {
	return &AccessKeyTenant{
		id:     tenant.id,
		secret: tenant.secret,
		name:   tenant.name,
	}
}

func (tenant *AccessKeyTenant) GetId() string {
	return tenant.id
}

func (tenant *AccessKeyTenant) GetSecret() string {
	return tenant.secret
}
