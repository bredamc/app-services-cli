package api

import (
	"net/http"
	"net/url"

	"github.com/redhat-developer/app-services-cli/pkg/api/generic"
	"github.com/redhat-developer/app-services-cli/pkg/api/rbac"
	"github.com/redhat-developer/app-services-cli/pkg/core/logging"
	connectormgmtclient "github.com/redhat-developer/app-services-sdk-core/app-services-sdk-go/connectormgmt/apiv1/client"

	amsclient "github.com/redhat-developer/app-services-sdk-core/app-services-sdk-go/accountmgmt/apiv1/client"
	kafkainstanceclient "github.com/redhat-developer/app-services-sdk-core/app-services-sdk-go/kafkainstance/apiv1/client"
	kafkamgmtclient "github.com/redhat-developer/app-services-sdk-core/app-services-sdk-go/kafkamgmt/apiv1/client"
	registryinstanceclient "github.com/redhat-developer/app-services-sdk-core/app-services-sdk-go/registryinstance/apiv1internal/client"
	registrymgmtclient "github.com/redhat-developer/app-services-sdk-core/app-services-sdk-go/registrymgmt/apiv1/client"
	svcacctmgmtclient "github.com/redhat-developer/app-services-sdk-core/app-services-sdk-go/serviceaccountmgmt/apiv1/client"
)

type API interface {
	KafkaMgmt() kafkamgmtclient.DefaultApi
	ServiceRegistryMgmt() registrymgmtclient.RegistriesApi
	ConnectorsMgmt() connectormgmtclient.APIClient
	ServiceAccountMgmt() svcacctmgmtclient.ServiceAccountsApi
	KafkaAdmin(instanceID string) (*kafkainstanceclient.APIClient, *kafkamgmtclient.KafkaRequest, error)
	ServiceRegistryInstance(instanceID string) (*registryinstanceclient.APIClient, *registrymgmtclient.Registry, error)
	AccountMgmt() amsclient.AppServicesApi
	RBAC() rbac.RbacAPI
	GenericAPI() generic.GenericAPI
	GetConfig() Config
}

type Config struct {
	AccessToken string
	ApiURL      *url.URL
	AuthURL     *url.URL
	ConsoleURL  *url.URL
	UserAgent   string
	HTTPClient  *http.Client
	Logger      logging.Logger
}
