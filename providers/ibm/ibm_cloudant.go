// Copyright 2019 The Terraformer Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package ibm

import (
	"os"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
	bluemix "github.com/IBM-Cloud/bluemix-go"
	"github.com/IBM-Cloud/bluemix-go/api/resource/resourcev1/catalog"
	"github.com/IBM-Cloud/bluemix-go/api/resource/resourcev2/controllerv2"
	"github.com/IBM-Cloud/bluemix-go/session"
)

// CloudantGenerator ...
type CloudantGenerator struct {
	IBMService
}

// loadMongoDB ...
func (g CloudantGenerator) loadCloudant(dbID string, dbName string) terraformutils.Resource {
	resources := terraformutils.NewSimpleResource(
		dbID,
		normalizeResourceName(dbName, false),
		"ibm_cloudant",
		"ibm",
		[]string{})
	return resources
}

// InitResources ...
func (g *CloudantGenerator) InitResources() error {
	region := g.Args["region"].(string)

	bmxConfig := &bluemix.Config{
		BluemixAPIKey: os.Getenv("IC_API_KEY"),
		Region:        region,
	}
	sess, err := session.New(bmxConfig)
	if err != nil {
		return err
	}

	catalogClient, err := catalog.New(sess)
	if err != nil {
		return err
	}
	controllerClient, err := controllerv2.New(sess)
	if err != nil {
		return err
	}
	serviceID, err := catalogClient.ResourceCatalog().FindByName("cloudantnosqldb", true)
	if err != nil {
		return err
	}
	for _, service := range serviceID {
		query := controllerv2.ServiceInstanceQuery{
			ServiceID: service.ID,
		}
		cloudantInstances, err := controllerClient.ResourceServiceInstanceV2().ListInstances(query)
		if err != nil {
			return err
		}
		for _, cloudantInstance := range cloudantInstances {
			if cloudantInstance.RegionID == region {
				// load Cloudant DBs for each Instance
				g.Resources = append(g.Resources, g.loadCloudant(cloudantInstance.ID, cloudantInstance.Name))
			}
		}
	}
	return nil
}
