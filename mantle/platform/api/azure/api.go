// Copyright 2023 Red Hat
// Copyright 2016 CoreOS, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package azure

import (
	"fmt"
	"math/rand"
	"os"
	"os/user"
	"path/filepath"
	"strings"
	"time"

	"github.com/Azure/azure-sdk-for-go/arm/compute"
	"github.com/Azure/azure-sdk-for-go/arm/network"
	"github.com/Azure/azure-sdk-for-go/arm/resources/resources"
	armStorage "github.com/Azure/azure-sdk-for-go/arm/storage"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/storage"
	"github.com/coreos/pkg/capnslog"

	"github.com/coreos/coreos-assembler/mantle/auth"
)

var (
	plog = capnslog.NewPackageLogger("github.com/coreos/coreos-assembler/mantle", "platform/api/azure")
)

type API struct {
	azIdCred   *azidentity.DefaultAzureCredential
	rgClient   resources.GroupsClient
	imgClient  compute.ImagesClient
	compClient compute.VirtualMachinesClient
	netClient  network.VirtualNetworksClient
	subClient  network.SubnetsClient
	ipClient   network.PublicIPAddressesClient
	intClient  network.InterfacesClient
	accClient  armStorage.AccountsClient
	opts       *Options
}

// New creates a new Azure client. If no publish settings file is provided or
// can't be parsed, an anonymous client is created.
func New(opts *Options) (*API, error) {
	azCreds, err := auth.ReadAzureCredentials(opts.AzureCredentials)
	if err != nil {
		return nil, fmt.Errorf("couldn't read Azure Credentials file: %v", err)
	}

	opts.SubscriptionID = azCreds.SubscriptionID
	os.Setenv("AZURE_CLIENT_ID", azCreds.ClientID)
	os.Setenv("AZURE_TENANT_ID", azCreds.TenantID)
	os.Setenv("AZURE_CLIENT_SECRET", azCreds.ClientSecret)

	api := &API{
		opts: opts,
	}

	if opts.Sku != "" && opts.DiskURI == "" && opts.Version == "" {
		return nil, fmt.Errorf("SKU set to %q but Disk URI and version not set; can't resolve", opts.Sku)
	}

	return api, nil
}

func (a *API) SetupClients() error {
	var err error
	a.azIdCred, err = azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		return err
	}
	a.imgClient = compute.NewImagesClientWithBaseURI(auther.BaseURI, auther.SubscriptionID)
	a.imgClient.Authorizer = auther
	a.compClient = compute.NewVirtualMachinesClientWithBaseURI(auther.BaseURI, auther.SubscriptionID)
	a.compClient.Authorizer = auther

	auther, err = auth.GetClientSetup(network.DefaultBaseURI)
	if err != nil {
		return err
	}
	a.netClient = network.NewVirtualNetworksClientWithBaseURI(auther.BaseURI, auther.SubscriptionID)
	a.netClient.Authorizer = auther
	a.subClient = network.NewSubnetsClientWithBaseURI(auther.BaseURI, auther.SubscriptionID)
	a.subClient.Authorizer = auther
	a.ipClient = network.NewPublicIPAddressesClientWithBaseURI(auther.BaseURI, auther.SubscriptionID)
	a.ipClient.Authorizer = auther
	a.intClient = network.NewInterfacesClientWithBaseURI(auther.BaseURI, auther.SubscriptionID)
	a.intClient.Authorizer = auther

	auther, err = auth.GetClientSetup(armStorage.DefaultBaseURI)
	if err != nil {
		return err
	}
	a.accClient = armStorage.NewAccountsClientWithBaseURI(auther.BaseURI, auther.SubscriptionID)
	a.accClient.Authorizer = auther

	return nil
}

func randomName(prefix string) string {
	b := make([]byte, 5)
	rand.Read(b)
	return fmt.Sprintf("%s-%x", prefix, b)
}

func (a *API) GC(gracePeriod time.Duration) error {
	durationAgo := time.Now().Add(-1 * gracePeriod)

	listGroups, err := a.ListResourceGroups("")
	if err != nil {
		return fmt.Errorf("listing resource groups: %v", err)
	}

	for _, l := range *listGroups.Value {
		if strings.HasPrefix(*l.Name, "kola-cluster") {
			terminate := false
			if l.Tags == nil || (*l.Tags)["createdAt"] == nil {
				// If the group name starts with kola-cluster and has
				// no tags OR no createdAt then it failed to properly
				// get created and we should clean it up.
				// https://github.com/coreos/coreos-assembler/issues/3057
				terminate = true
			} else {
				createdAt := *(*l.Tags)["createdAt"]
				timeCreated, err := time.Parse(time.RFC3339, createdAt)
				if err != nil {
					return fmt.Errorf("error parsing time: %v", err)
				}
				if !timeCreated.After(durationAgo) {
					// If the group is older than specified time then gc
					terminate = true
				}
			}
			if terminate {
				if err = a.TerminateResourceGroup(*l.Name); err != nil {
					return err
				}
			}
		}
	}

	return nil
}
