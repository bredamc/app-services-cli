package ams

import (
	"context"
	"errors"

	"github.com/redhat-developer/app-services-cli/pkg/api/ams/amsclient"
	"github.com/redhat-developer/app-services-cli/pkg/connection"
	"github.com/redhat-developer/app-services-cli/pkg/remote"
)

func CheckTermsAccepted(ctx context.Context, spec remote.AmsConfig, conn connection.Connection) (accepted bool, redirectURI string, err error) {
	termsReview, _, err := conn.API().AccountMgmt().
		ApiAuthorizationsV1SelfTermsReviewPost(ctx).
		SelfTermsReview(amsclient.SelfTermsReview{
			EventCode: &spec.TermsAndConditionsEventCode,
			SiteCode:  &spec.TermsAndConditionsSiteCode,
		}).Execute()
	if err != nil {
		return false, "", err
	}

	if !termsReview.GetTermsAvailable() && !termsReview.GetTermsRequired() {
		return true, "", nil
	}

	if !termsReview.HasRedirectUrl() {
		return false, "", errors.New("terms must be signed, but there is no terms URL")
	}

	return false, termsReview.GetRedirectUrl(), nil
}

func GetUserSupportedInstanceTypes(ctx context.Context, spec remote.AmsConfig, conn connection.Connection) (err error, quota []string) {
	err, orgId := GetOrganizationId(ctx, conn)
	if err != nil {
		return err, nil
	}
	quotaCostGet, _, err := conn.API().AccountMgmt().
		ApiAccountsMgmtV1OrganizationsOrgIdQuotaCostGet(ctx, orgId).
		Execute()
	if err != nil {
		return err, nil
	}

	var quotas []string
	for _, quota := range quotaCostGet.GetItems() {
		if quota.Id == &spec.TrialQuotaID {
			quotas = append(quotas, QuotaTrialType)
		}
		if quota.Id == &spec.InstanceQuotaID {
			quotas = append(quotas, QuotaStandardType)
		}
	}
	if len(quotas) == 0 {
		return errors.New("Your account missing quota for creating instance of specified type"), nil
	}
	return nil, quotas
}

func GetOrganizationId(ctx context.Context, conn connection.Connection) (err error, accountId string) {
	account, _, err := conn.API().AccountMgmt().ApiAccountsMgmtV1CurrentAccountGet(ctx).
		Execute()
	if err != nil {
		return err, ""
	}

	if account.GetBanned() {
		return errors.New("Your account has been banned from using the App Services. If you believe this is an error, please contact our support team."), ""
	}

	return nil, account.GetId()
}
