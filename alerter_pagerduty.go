package gotel

import (
	"fmt"

	"github.com/stvp/pager"
)

type pagerDutyAlerter struct {
	Cfg config
}

func (s *pagerDutyAlerter) Bootstrap() {

}

func (s *pagerDutyAlerter) Name() string {
	return "PagerDuty"
}

func (s *pagerDutyAlerter) Alert(res reservation) bool {

	l.info("PagerDuty API key [%s]", s.Cfg.Pagerduty.Servicekey)

	ip, err := externalIP()
	if err != nil {
		ip = "N/A"
	}

	pager.ServiceKey = s.Cfg.Pagerduty.Servicekey
	incidentKey, err := pager.Trigger(
		fmt.Sprintf("App [%s] Component: [%s] failed to checkin on ip [%s]. Contact owner [%s]", res.App, res.Component, ip, res.Owner))

	if err != nil {
		l.err("[ERROR] Unable to create pagerduty alert for job [%s] component [%s] error [%v]\n", res.App, res.Component, err)
		return false
	} else {
		l.info("PagerDuty incident key created %s\n", incidentKey)
		return true
	}
}
