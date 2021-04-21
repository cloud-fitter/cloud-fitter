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

func (tennat *AccessKeyTenant) GetId() string {
	return tennat.id
}

func (tennat *AccessKeyTenant) GetSecret() string {
	return tennat.secret
}
