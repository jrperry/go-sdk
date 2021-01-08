package iland

import "fmt"

type VCCFailoverPlan struct {
	ID              string                          `json:"uuid"`
	UID             string                          `json:"uid"`
	Name            string                          `json:"name"`
	Description     string                          `json:"description"`
	VCCTenantUID    string                          `json:"vcc_tenant_uid"`
	VCCTenantName   string                          `json:"vcc_tenant_name"`
	Status          string                          `json:"status"`
	OrgID           string                          `json:"org_uuid"`
	CompanyID       string                          `json:"company_id"`
	UpdatedDate     int                             `json:"updated_date"`
	LastTest        int                             `json:"last_test"`
	VirtualMachines []VCCFailoverPlanVirtualMachine `json:"vms"`
}

type VCCFailoverPlanVirtualMachine struct {
	Name string `json:"name"`
	ID   string `json:"id"`
}

type vccFailoverPlanService struct {
	client *client
}

func (s *vccFailoverPlanService) Get(id string) (VCCFailoverPlan, error) {
	obj := VCCFailoverPlan{}
	err := s.client.getObject(fmt.Sprintf("/v1/vcc-failover-plan/%s", id), &obj)
	if err != nil {
		return VCCFailoverPlan{}, err
	}
	return obj, nil
}
