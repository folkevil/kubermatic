package ecs

//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//http://www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.
//
// Code generated by Alibaba Cloud SDK Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

// AutoProvisioningGroup is a nested struct in ecs response
type AutoProvisioningGroup struct {
	AutoProvisioningGroupId          string                                                `json:"AutoProvisioningGroupId" xml:"AutoProvisioningGroupId"`
	AutoProvisioningGroupName        string                                                `json:"AutoProvisioningGroupName" xml:"AutoProvisioningGroupName"`
	AutoProvisioningGroupType        string                                                `json:"AutoProvisioningGroupType" xml:"AutoProvisioningGroupType"`
	Status                           string                                                `json:"Status" xml:"Status"`
	State                            string                                                `json:"State" xml:"State"`
	RegionId                         string                                                `json:"RegionId" xml:"RegionId"`
	ValidFrom                        string                                                `json:"ValidFrom" xml:"ValidFrom"`
	ValidUntil                       string                                                `json:"ValidUntil" xml:"ValidUntil"`
	ExcessCapacityTerminationPolicy  string                                                `json:"ExcessCapacityTerminationPolicy" xml:"ExcessCapacityTerminationPolicy"`
	MaxSpotPrice                     float64                                               `json:"MaxSpotPrice" xml:"MaxSpotPrice"`
	LaunchTemplateId                 string                                                `json:"LaunchTemplateId" xml:"LaunchTemplateId"`
	LaunchTemplateVersion            string                                                `json:"LaunchTemplateVersion" xml:"LaunchTemplateVersion"`
	TerminateInstances               bool                                                  `json:"TerminateInstances" xml:"TerminateInstances"`
	TerminateInstancesWithExpiration bool                                                  `json:"TerminateInstancesWithExpiration" xml:"TerminateInstancesWithExpiration"`
	CreationTime                     string                                                `json:"CreationTime" xml:"CreationTime"`
	SpotOptions                      SpotOptions                                           `json:"SpotOptions" xml:"SpotOptions"`
	PayAsYouGoOptions                PayAsYouGoOptions                                     `json:"PayAsYouGoOptions" xml:"PayAsYouGoOptions"`
	TargetCapacitySpecification      TargetCapacitySpecification                           `json:"TargetCapacitySpecification" xml:"TargetCapacitySpecification"`
	LaunchTemplateConfigs            LaunchTemplateConfigsInDescribeAutoProvisioningGroups `json:"LaunchTemplateConfigs" xml:"LaunchTemplateConfigs"`
}
