// Generated by text/template; DO NOT EDIT
package database

import (
	"subLease/src/server/domain"
)



type TenantUpdate struct {
    Id *int
    FirstName *string
    LastName *string
    SocialSecurityNumber *socialSecurityNumber.SocialSecurityNumber
}

func (t *TenantUpdate) UpdateTenantWithValuesFrom(e *domain.Tenant, database Database) {
    if t.Id != nil {
		e.Id = *t.Id
    }
    if t.FirstName != nil {
		e.FirstName = *t.FirstName
    }
    if t.LastName != nil {
		e.LastName = *t.LastName
    }
    if t.SocialSecurityNumber != nil {
		e.SocialSecurityNumber = *t.SocialSecurityNumber
    }
}

