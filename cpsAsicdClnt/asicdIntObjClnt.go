//
//Copyright [2016] [SnapRoute Inc]
//
//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
//       Unless required by applicable law or agreed to in writing, software
//       distributed under the License is distributed on an "AS IS" BASIS,
//       WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//       See the License for the specific language governing permissions and
//       limitations under the License.
//
// _______  __       __________   ___      _______.____    __    ____  __  .___________.  ______  __    __
// |   ____||  |     |   ____\  \ /  /     /       |\   \  /  \  /   / |  | |           | /      ||  |  |  |
// |  |__   |  |     |  |__   \  V  /     |   (----` \   \/    \/   /  |  | `---|  |----`|  ,----'|  |__|  |
// |   __|  |  |     |   __|   >   <       \   \      \            /   |  |     |  |     |  |     |   __   |
// |  |     |  `----.|  |____ /  .  \  .----)   |      \    /\    /    |  |     |  |     |  `----.|  |  |  |
// |__|     |_______||_______/__/ \__\ |_______/        \__/  \__/     |__|     |__|      \______||__|  |__|
//

package cpsAsicdClnt

import (
	"utils/clntUtils/clntDefs/asicdClntDefs"
)

func (asicdClientMgr *CPSAsicdClntMgr) CreateIPv4Neighbor(ipAddr, macAddr string, vlanId, ifIndex int32) (val int32, err error) {
	return val, err
}

func (asicdClientMgr *CPSAsicdClntMgr) UpdateIPv4Neighbor(ipAddr, macAddr string, vlanId, ifIndex int32) (val int32, err error) {
	return val, err
}

func (asicdClientMgr *CPSAsicdClntMgr) DeleteIPv4Neighbor(ipAddr, macAddr string, vlanId, ifIndex int32) (val int32, err error) {
	return val, err
}

func (asicdClientMgr *CPSAsicdClntMgr) CreateIPv6Neighbor(ipAddr, macAddr string, vlanId, ifIndex int32) (val int32, err error) {
	return val, err
}

func (asicdClientMgr *CPSAsicdClntMgr) UpdateIPv6Neighbor(ipAddr, macAddr string, vlanId, ifIndex int32) (val int32, err error) {
	return val, err
}

func (asicdClientMgr *CPSAsicdClntMgr) DeleteIPv6Neighbor(ipAddr, macAddr string, vlanId, ifIndex int32) (val int32, err error) {
	return val, err
}

func (asicdClientMgr *CPSAsicdClntMgr) GetBulkVlan(curMark, count int) (*asicdClntDefs.VlanGetInfo, error) {
	var retObj asicdClntDefs.VlanGetInfo
	return &retObj, nil
}

func (asicdClientMgr *CPSAsicdClntMgr) CreateLag(ifName string, hashType int32, ifIndexList string) (val int32, err error) {
	return val, err
}

func (asicdClientMgr *CPSAsicdClntMgr) DeleteLag(ifIndex int32) (val int32, err error) {
	return val, err
}

func (asicdClientMgr *CPSAsicdClntMgr) UpdateLag(ifIndex, hashType int32, ifIndexList string) (val int32, err error) {
	return val, err
}

func (asicdClientMgr *CPSAsicdClntMgr) GetBulkLag(fromIndex, count int) (*asicdClntDefs.LagGetInfo, error) {
	var retObj asicdClntDefs.LagGetInfo
	return &retObj, nil
}

func (asicdClientMgr *CPSAsicdClntMgr) CreateLagCfgIntfList(ifName string, ifIndexList []int32) (rv bool, err error) {
	return rv, err
}

func (asicdClientMgr *CPSAsicdClntMgr) UpdateLagCfgIntfList(ifName string, ifIndexList []int32) (rv bool, err error) {
	return rv, err
}

func (asicdClientMgr *CPSAsicdClntMgr) DeleteLagCfgIntfList(ifName string, ifIndexList []int32) (rv bool, err error) {
	return rv, err
}

func (asicdClientMgr *CPSAsicdClntMgr) IsLinuxOnlyPlugin() (bool, error) {
	return true, nil
}

func (asicdClientMgr *CPSAsicdClntMgr) GetAllPortsWithDirtyCache() ([]*asicdClntDefs.Port, error) {
	return nil, nil
}

func (asicdClientMgr *CPSAsicdClntMgr) OnewayCreateIPv4Route(ipv4RouteList []*asicdClntDefs.IPv4Route) {
}
func (asicdClientMgr *CPSAsicdClntMgr) OnewayDeleteIPv4Route(ipv4RouteList []*asicdClntDefs.IPv4Route) {
}
func (asicdClientMgr *CPSAsicdClntMgr) OnewayCreateIPv6Route(ipv6RouteList []*asicdClntDefs.IPv6Route) {
}
func (asicdClientMgr *CPSAsicdClntMgr) OnewayDeleteIPv6Route(ipv6RouteList []*asicdClntDefs.IPv6Route) {
}
