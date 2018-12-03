package gitlab

import (
	"fmt"
	"net/url"
)

type ApprovalsService struct {
	client *Client
}

type Approval struct {
}

type ApprovalOptions struct {
	Approvals_before_merge                         *int  `url:"approvals_before_merge" json:"approvals_before_merge"`
	Disable_overriding_approvers_per_merge_request *bool `url:"disable_overriding_approvers_per_merge_request" json:"disable_overriding_approvers_per_merge_request"`
	Reset_approvals_on_push                        *bool `url:"reset_approvals_on_push" json:"reset_approvals_on_push"`
}

func (s *ApprovalsService) EditApprovals(pid interface{}, opt *ApprovalOptions, options ...OptionFunc) (*Approval, *Response, error) {
	project, err := parseID(pid)
	if err != nil {
		return nil, nil, err
	}
	u := fmt.Sprintf("projects/%s/approvals", url.QueryEscape(project))

	req, err := s.client.NewRequest("POST", u, opt, options)
	if err != nil {
		return nil, nil, err
	}

	p := new(Approval)
	resp, err := s.client.Do(req, p)
	if err != nil {
		return nil, resp, err
	}

	return p, resp, err
}
