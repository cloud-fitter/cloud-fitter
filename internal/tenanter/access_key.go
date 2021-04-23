package tenanter

type AccessKeyTenant struct {
	id     string
	secret string
}

func NewTenantWithAccessKey(accessKeyId, accessKeySecret string) Tenanter {
	return &AccessKeyTenant{
		id:     accessKeyId,
		secret: accessKeySecret,
	}
}

func (tenant *AccessKeyTenant) GetId() string {
	return tenant.id
}

func (tenant *AccessKeyTenant) GetSecret() string {
	return tenant.secret
}

func (tenant *AccessKeyTenant) Clone() Tenanter {
	return &AccessKeyTenant{
		id:     tenant.id,
		secret: tenant.secret,
	}
}
