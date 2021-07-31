package pool

import (
	"context"
	"encoding/xml"
	go_f5_soap "github.com/wule61/go-f5-soap"
)

const tns = "urn:iControl:GlobalLB/Pool"

type GetAlternateLbMethodBody struct {
	GetAlternateLbMethod GetAlternateLbMethod `xml:"tns:get_alternate_lb_method"`
}

type GetAlternateLbMethod struct {
	PoolNames PoolNames `xml:"pool_names"`
}

type PoolNames struct {
	Item []string `xml:"item"`
}

type Pool struct {
	c *go_f5_soap.Client
}

func NewPool(c *go_f5_soap.Client) *Pool {
	return &Pool{
		c: c,
	}
}

type AlternateLbMethodByPoolNamesResp struct {
	XMLName xml.Name `xml:"Envelope"`
	Body    struct {
		GetAlternateLbMethodResponse struct {
			Return struct {
				Item []string `xml:"item"`
			} `xml:"return"`
		} `xml:"get_alternate_lb_methodResponse"`
	} `xml:"Body"`
}

func (p *Pool) GetAlternateLBMethod(poolName string) (string, error) {

	arr, err := p.GetAlternateLBMethodByPoolNames([]string{poolName})
	if err != nil {
		return "", err
	}
	if len(arr) > 0 {
		return arr[0], nil
	}

	return "LB_METHOD_ROUND_ROBIN", nil
}

func (p *Pool) GetAlternateLBMethodByPoolNames(poolNames []string) ([]string, error) {

	type req struct {
		go_f5_soap.BaseEnvEnvelope
		Body GetAlternateLbMethodBody `xml:"env:Body"`
	}

	bt, err := p.c.Call(context.Background(), req{
		BaseEnvEnvelope: go_f5_soap.NewBaseEnvEnvelope(tns),
		Body:            GetAlternateLbMethodBody{GetAlternateLbMethod{PoolNames{Item: poolNames}}},
	})

	if err != nil {
		return nil, err
	}

	var resp AlternateLbMethodByPoolNamesResp
	if err := xml.Unmarshal(bt, &resp); err != nil {
		return nil, err
	}

	return resp.Body.GetAlternateLbMethodResponse.Return.Item, nil
}

type GetPreferredLBMethodBody struct {
	GetPreferredLBMethod GetPreferredLBMethod `xml:"tns:get_preferred_lb_method"`
}

type GetPreferredLBMethod struct {
	PoolNames PoolNames `xml:"pool_names"`
}

type PreferredLBMethodByPoolNamesResp struct {
	XMLName xml.Name `xml:"Envelope"`
	Body    struct {
		GetPreferredLbMethodResponse struct {
			Return struct {
				Item []string `xml:"item"`
			} `xml:"return"`
		} `xml:"get_preferred_lb_methodResponse"`
	} `xml:"Body"`
}

func (p *Pool) GetPreferredLBMethod(poolName string) (string, error) {
	arr, err := p.GetPreferredLBMethodByPoolNames([]string{poolName})
	if err != nil {
		return "", err
	}
	if len(arr) > 0 {
		return arr[0], nil
	}

	return "LB_METHOD_ROUND_ROBIN", nil
}

func (p *Pool) GetPreferredLBMethodByPoolNames(poolNames []string) ([]string, error) {

	type req struct {
		go_f5_soap.BaseEnvEnvelope
		Body GetPreferredLBMethodBody `xml:"env:Body"`
	}

	bt, err := p.c.Call(context.Background(), req{
		BaseEnvEnvelope: go_f5_soap.NewBaseEnvEnvelope(tns),
		Body:            GetPreferredLBMethodBody{GetPreferredLBMethod{PoolNames{Item: poolNames}}},
	})

	if err != nil {
		return nil, err
	}

	var resp PreferredLBMethodByPoolNamesResp
	if err := xml.Unmarshal(bt, &resp); err != nil {
		return nil, err
	}

	return resp.Body.GetPreferredLbMethodResponse.Return.Item, nil
}

type GetTTLBody struct {
	GetTTL GetTTL `xml:"tns:get_ttl"`
}

type GetTTL struct {
	PoolNames PoolNames `xml:"pool_names"`
}

type TTLResp struct {
	XMLName xml.Name `xml:"Envelope"`
	Body    struct {
		GetTtlResponse struct {
			Return struct {
				Item []int64 `xml:"item"`
			} `xml:"return"`
		} `xml:"get_ttlResponse"`
	} `xml:"Body"`
}

func (p *Pool) GetTTL(poolName string) (int64, error) {

	arr, err := p.GetTTLByPoolNames([]string{poolName})
	if err != nil {
		return 0, nil
	}

	if len(arr) > 0 {
		return arr[0], nil
	}

	return 0, nil
}

func (p *Pool) GetTTLByPoolNames(poolNames []string) ([]int64, error) {

	type req struct {
		go_f5_soap.BaseEnvEnvelope
		Body GetTTLBody `xml:"env:Body"`
	}

	bt, err := p.c.Call(context.Background(), req{
		BaseEnvEnvelope: go_f5_soap.NewBaseEnvEnvelope(tns),
		Body:            GetTTLBody{GetTTL{PoolNames{Item: poolNames}}},
	})
	if err != nil {
		return nil, err
	}

	var resp TTLResp
	if err := xml.Unmarshal(bt, &resp); err != nil {
		return nil, err
	}

	return resp.Body.GetTtlResponse.Return.Item, err
}

type GetVerifyMemberAvailabilityStateBody struct {
	GetVerifyMemberAvailabilityState GetVerifyMemberAvailabilityState `xml:"tns:get_verify_member_availability_state"`
}

type GetVerifyMemberAvailabilityState struct {
	PoolNames PoolNames `xml:"pool_names"`
}

type VerifyMemberAvailabilityStateResp struct {
	XMLName xml.Name `xml:"Envelope"`
	Body    struct {
		GetVerifyMemberAvailabilityStateResponse struct {
			Return struct {
				Item []string `xml:"item"`
			} `xml:"return"`
		} `xml:"get_verify_member_availability_stateResponse"`
	} `xml:"Body"`
}

func (p *Pool) GetVerifyMemberAvailabilityState(poolName string) (string, error) {

	arr, err := p.GetVerifyMemberAvailabilityStateByPoolNames([]string{poolName})
	if err != nil {
		return "", err
	}

	if len(arr) > 0 {
		return arr[0], nil
	}

	return "", err
}

func (p *Pool) GetVerifyMemberAvailabilityStateByPoolNames(poolNames []string) ([]string, error) {

	type req struct {
		go_f5_soap.BaseEnvEnvelope
		Body GetVerifyMemberAvailabilityStateBody `xml:"env:Body"`
	}

	bt, err := p.c.Call(context.Background(), req{
		BaseEnvEnvelope: go_f5_soap.NewBaseEnvEnvelope(tns),
		Body:            GetVerifyMemberAvailabilityStateBody{GetVerifyMemberAvailabilityState{PoolNames{Item: poolNames}}},
	})
	if err != nil {
		return nil, err
	}

	var resp VerifyMemberAvailabilityStateResp
	if err := xml.Unmarshal(bt, &resp); err != nil {
		return nil, err
	}

	return resp.Body.GetVerifyMemberAvailabilityStateResponse.Return.Item, err
}

type GetAnswersToReturnBody struct {
	GetAnswersToReturn GetAnswersToReturn `xml:"tns:get_answers_to_return"`
}

type GetAnswersToReturn struct {
	PoolNames PoolNames `xml:"pool_names"`
}

type AnswersToReturnResp struct {
	XMLName xml.Name `xml:"Envelope"`
	Body    struct {
		GetAnswersToReturnResponse struct {
			Return struct {
				Item []int64 `xml:"item"`
			} `xml:"return"`
		} `xml:"get_answers_to_returnResponse"`
	} `xml:"Body"`
}

func (p *Pool) GetAnswersToReturn(poolName string) (int64, error) {

	arr, err := p.GetAnswersToReturnByPoolNames([]string{poolName})
	if err != nil {
		return 0, err
	}

	if len(arr) > 0 {
		return arr[0], nil
	}

	return 0, err
}

func (p *Pool) GetAnswersToReturnByPoolNames(poolNames []string) ([]int64, error) {

	type req struct {
		go_f5_soap.BaseEnvEnvelope
		Body GetAnswersToReturnBody `xml:"env:Body"`
	}

	bt, err := p.c.Call(context.Background(), req{
		BaseEnvEnvelope: go_f5_soap.NewBaseEnvEnvelope(tns),
		Body:            GetAnswersToReturnBody{GetAnswersToReturn{PoolNames{Item: poolNames}}},
	})
	if err != nil {
		return nil, err
	}

	var resp AnswersToReturnResp
	if err := xml.Unmarshal(bt, &resp); err != nil {
		return nil, err
	}

	return resp.Body.GetAnswersToReturnResponse.Return.Item, err
}

type GetObjectStatusBody struct {
	GetObjectStatus GetObjectStatus `xml:"tns:get_object_status"`
}

type GetObjectStatus struct {
	PoolNames PoolNames `xml:"pool_names"`
}

type ObjectStatus struct {
	AvailabilityStatus string
	EnabledStatus      string
	StatusDescription  string
}

type ObjectStatusResp struct {
	XMLName xml.Name `xml:"Envelope"`
	Body    struct {
		GetObjectStatusResponse struct {
			Return struct {
				Item []struct {
					AvailabilityStatus struct {
						Text string `xml:",chardata"`
					} `xml:"availability_status"`
					EnabledStatus struct {
						Text string `xml:",chardata"`
					} `xml:"enabled_status"`
					StatusDescription struct {
						Text string `xml:",chardata"`
					} `xml:"status_description"`
				} `xml:"item"`
			} `xml:"return"`
		} `xml:"get_object_statusResponse"`
	} `xml:"Body"`
}

func (p *Pool) GetObjectStatus(poolName string) (ObjectStatus, error) {
	arr, err := p.GetObjectStatusByPoolNames([]string{poolName})
	if err != nil {
		return ObjectStatus{}, err
	}
	if len(arr) > 0 {
		return arr[0], nil
	}

	return ObjectStatus{}, nil
}

func (p *Pool) GetObjectStatusByPoolNames(poolNames []string) ([]ObjectStatus, error) {

	type req struct {
		go_f5_soap.BaseEnvEnvelope
		Body GetObjectStatusBody `xml:"env:Body"`
	}

	bt, err := p.c.Call(context.Background(), req{
		BaseEnvEnvelope: go_f5_soap.NewBaseEnvEnvelope(tns),
		Body:            GetObjectStatusBody{GetObjectStatus{PoolNames{Item: poolNames}}},
	})
	if err != nil {
		return nil, err
	}

	var resp ObjectStatusResp
	if err := xml.Unmarshal(bt, &resp); err != nil {
		return nil, err
	}

	var res []ObjectStatus
	for _, v := range resp.Body.GetObjectStatusResponse.Return.Item {
		res = append(res, ObjectStatus{
			AvailabilityStatus: v.AvailabilityStatus.Text,
			EnabledStatus:      v.EnabledStatus.Text,
			StatusDescription:  v.StatusDescription.Text,
		})
	}

	return res, nil
}

type GetEnabledStateBody struct {
	GetEnabledState GetEnabledState `xml:"tns:get_enabled_state"`
}

type GetEnabledState struct {
	PoolNames PoolNames `xml:"pool_names"`
}

type EnabledStateResp struct {
	XMLName xml.Name `xml:"Envelope"`
	Body    struct {
		GetEnabledStateResponse struct {
			Return struct {
				Item []string `xml:"item"`
			} `xml:"return"`
		} `xml:"get_enabled_stateResponse"`
	} `xml:"Body"`
}

func (p *Pool) GetEnabledState(poolName string) (string, error) {
	arr, err := p.GetEnabledStateByNames([]string{poolName})
	if err != nil {
		return "", err
	}

	if len(arr) > 0 {
		return arr[0], err
	}

	return "", nil
}

func (p *Pool) GetEnabledStateByNames(poolNames []string) ([]string, error) {

	type req struct {
		go_f5_soap.BaseEnvEnvelope
		Body GetEnabledStateBody `xml:"env:Body"`
	}

	bt, err := p.c.Call(context.Background(), req{
		BaseEnvEnvelope: go_f5_soap.NewBaseEnvEnvelope(tns),
		Body:            GetEnabledStateBody{GetEnabledState{PoolNames{Item: poolNames}}},
	})
	if err != nil {
		return nil, err
	}

	var resp EnabledStateResp
	if err := xml.Unmarshal(bt, &resp); err != nil {
		return nil, err
	}

	return resp.Body.GetEnabledStateResponse.Return.Item, nil
}
