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
	"models/objects"
	"utils/clntUtils/clntDefs/asicdClntDefs"
	"errors"
	"fmt"
	"net"
	"sync"
)

/*
#include <stdint.h>
#include <stdlib.h>
#include <cps.h>
#cgo CFLAGS: -I. -I/usr/include/
#cgo LDFLAGS: -L/usr/lib/x86_64-linux-gnu/ -lsonic_object_library
*/
import "C"

var cpsAsicdMutex *sync.Mutex = &sync.Mutex{}

func (asicdClientMgr *CPSAsicdClntMgr) GetBulkAsicGlobalState(fromIndex int, count int) (*asicdClntDefs.AsicGlobalStateGetInfo, error) {
	var retObj asicdClntDefs.AsicGlobalStateGetInfo
	return &retObj, nil
}

func (asicdClientMgr *CPSAsicdClntMgr) GetAsicGlobalState(ModuleId uint8) (*objects.AsicGlobalState, error) {
	return nil, nil
}

func (asicdClientMgr *CPSAsicdClntMgr) CreateAsicGlobalPM(cfg *objects.AsicGlobalPM) (bool, error) {
	return false, nil
}

func (asicdClientMgr *CPSAsicdClntMgr) UpdateAsicGlobalPM(origCfg, newCfg *objects.AsicGlobalPM, attrset []bool, op []*objects.PatchOpInfo) (bool, error) {
	return true, nil
}

func (asicdClientMgr *CPSAsicdClntMgr) DeleteAsicGlobalPM(cfg *objects.AsicGlobalPM) (bool, error) {
	return true, nil
}

func (asicdClientMgr *CPSAsicdClntMgr) GetBulkAsicGlobalPM(fromIndex int, count int) (*asicdClntDefs.AsicGlobalPMGetInfo, error) {
	var retObj asicdClntDefs.AsicGlobalPMGetInfo
	return &retObj, nil
}

func (asicdClientMgr *CPSAsicdClntMgr) GetAsicGlobalPM(ModuleId uint8, Resource string) (*objects.AsicGlobalPM, error) {
	return nil, nil
}

func (asicdClientMgr *CPSAsicdClntMgr) GetBulkAsicGlobalPMState(fromIndex int, count int) (*asicdClntDefs.AsicGlobalPMStateGetInfo, error) {
	var retObj asicdClntDefs.AsicGlobalPMStateGetInfo
	return &retObj, nil
}

func (asicdClientMgr *CPSAsicdClntMgr) GetAsicGlobalPMState(ModuleId uint8, Resource string) (*objects.AsicGlobalPMState, error) {
	return nil, nil
}

func (asicdClientMgr *CPSAsicdClntMgr) CreateEthernetPM(cfg *objects.EthernetPM) (bool, error) {
	return true, nil
}

func (asicdClientMgr *CPSAsicdClntMgr) UpdateEthernetPM(origCfg, newCfg *objects.EthernetPM, attrset []bool, op []*objects.PatchOpInfo) (bool, error) {
	return true, nil
}

func (asicdClientMgr *CPSAsicdClntMgr) DeleteEthernetPM(cfg *objects.EthernetPM) (bool, error) {
	return true, nil
}

func (asicdClientMgr *CPSAsicdClntMgr) GetBulkEthernetPM(fromIndex int, count int) (*asicdClntDefs.EthernetPMGetInfo, error) {
	var retObj asicdClntDefs.EthernetPMGetInfo
	return &retObj, nil
}

func (asicdClientMgr *CPSAsicdClntMgr) GetEthernetPM(IntfRef string, Resource string) (*objects.EthernetPM, error) {
	return nil, nil
}

func (asicdClientMgr *CPSAsicdClntMgr) GetBulkEthernetPMState(fromIndex int, count int) (*asicdClntDefs.EthernetPMStateGetInfo, error) {
	var retObj asicdClntDefs.EthernetPMStateGetInfo
	return &retObj, nil
}

func (asicdClientMgr *CPSAsicdClntMgr) GetEthernetPMState(IntfRef string, Resource string) (*objects.EthernetPMState, error) {
	return nil, nil
}

func (asicdClientMgr *CPSAsicdClntMgr) GetBulkAsicSummaryState(fromIndex int, count int) (*asicdClntDefs.AsicSummaryStateGetInfo, error) {
	var retObj asicdClntDefs.AsicSummaryStateGetInfo
	return &retObj, nil
}

func (asicdClientMgr *CPSAsicdClntMgr) GetAsicSummaryState(ModuleId uint8) (*objects.AsicSummaryState, error) {
	return nil, nil
}

func (asicdClientMgr *CPSAsicdClntMgr) CreateVlan(cfg *objects.Vlan) (bool, error) {
	cpsAsicdMutex.Lock()
	defer cpsAsicdMutex.Unlock()
	Logger.Info("Calling CPS CreateVlan:", cfg)
	tagPorts := C.MakeCharArray(C.int(len(cfg.IntfList)))
	defer C.FreeCharArray(tagPorts, C.int(len(cfg.IntfList)))
	for idx, intf := range cfg.IntfList {
		C.SetArrayString(tagPorts, C.CString(intf), C.int(idx))
	}
	untagPorts := C.MakeCharArray(C.int(len(cfg.UntagIntfList)))
	defer C.FreeCharArray(untagPorts, C.int(len(cfg.UntagIntfList)))
	for idx, intf := range cfg.UntagIntfList {
		C.SetArrayString(untagPorts, C.CString(intf), C.int(idx))
	}
	rv := int(C.CPSCreateVlan(C.uint32_t(cfg.VlanId), C.uint32_t(len(cfg.IntfList)), tagPorts, C.uint32_t(len(cfg.UntagIntfList)), untagPorts))
	if rv != 0 {
		return false, errors.New("Error Creating Vlan")
	}
	return true, nil
}

func (asicdClientMgr *CPSAsicdClntMgr) DeleteVlan(cfg *objects.Vlan) (bool, error) {
	cpsAsicdMutex.Lock()
	defer cpsAsicdMutex.Unlock()
	Logger.Info("Calling CPS DeleteVlan:", cfg)
	rv := int(C.CPSDeleteVlan(C.uint32_t(cfg.VlanId)))
	if rv != 0 {
		return false, errors.New("Error Delete Vlan")
	}
	return true, nil
}

func (asicdClientMgr *CPSAsicdClntMgr) UpdateVlan(origCfg, newCfg *objects.Vlan, attrset []bool, op []*objects.PatchOpInfo) (bool, error) {
	return true, nil
}

func (asicdClientMgr *CPSAsicdClntMgr) GetBulkVlanState(fromIdx, count int) (*asicdClntDefs.VlanStateGetInfo, error) {
	var retObj asicdClntDefs.VlanStateGetInfo
	return &retObj, nil
}

func (asicdClientMgr *CPSAsicdClntMgr) GetVlanState(VlanId int32) (*objects.VlanState, error) {
	return nil, nil
}

func (asicdClientMgr *CPSAsicdClntMgr) CreateIPv4Intf(cfg *objects.IPv4Intf) (bool, error) {
	cpsAsicdMutex.Lock()
	defer cpsAsicdMutex.Unlock()
	Logger.Info("Calling CPS CreateIPv4Intf:", cfg)
	ipAddr, ipNet, err := net.ParseCIDR(cfg.IpAddr)
	if err != nil {
		return false, errors.New(fmt.Sprintln("Invalid IP Address", err))
	}
	prefixLen, _ := ipNet.Mask.Size()
	if err != nil {
		return false, errors.New(fmt.Sprintln("Invalid IP Address", err))
	}

	rv := int(C.CPSCreateIPv4Intf(C.CString(cfg.IntfRef), C.CString(ipAddr.String()), C.uint32_t(prefixLen)))
	if rv != 0 {
		return false, errors.New("Error Creating IPv4 Intf")
	}
	return true, nil
}

func (asicdClientMgr *CPSAsicdClntMgr) DeleteIPv4Intf(cfg *objects.IPv4Intf) (bool, error) {
	cpsAsicdMutex.Lock()
	defer cpsAsicdMutex.Unlock()
	Logger.Info("Calling CPS DeleteIPv4Intf:", cfg)
	ipAddr, ipNet, err := net.ParseCIDR(cfg.IpAddr)
	if err != nil {
		return false, errors.New(fmt.Sprintln("Invalid IP Address", err))
	}
	prefixLen, _ := ipNet.Mask.Size()
	if err != nil {
		return false, errors.New(fmt.Sprintln("Invalid IP Address", err))
	}

	rv := int(C.CPSDeleteIPv4Intf(C.CString(cfg.IntfRef), C.CString(ipAddr.String()), C.uint32_t(prefixLen)))
	if rv != 0 {
		return false, errors.New("Error Creating IPv4 Intf")
	}
	return true, nil
}

func (asicdClientMgr *CPSAsicdClntMgr) UpdateIPv4Intf(origCfg, newCfg *objects.IPv4Intf, attrset []bool, op []*objects.PatchOpInfo) (bool, error) {
	return true, nil
}

func (asicdClientMgr *CPSAsicdClntMgr) GetBulkIPv4IntfState(fromIdx, count int) (*asicdClntDefs.IPv4IntfStateGetInfo, error) {
	var retObj asicdClntDefs.IPv4IntfStateGetInfo
	return &retObj, nil
}

func (asicdClientMgr *CPSAsicdClntMgr) GetIPv4IntfState(IntfRef string) (*objects.IPv4IntfState, error) {
	return nil, nil
}

func (asicdClientMgr *CPSAsicdClntMgr) CreatePort(cfg *objects.Port) (bool, error) {
	return true, nil
}

func (asicdClientMgr *CPSAsicdClntMgr) UpdatePort(origCfg, newCfg *objects.Port, attrset []bool, op []*objects.PatchOpInfo) (bool, error) {
	return true, nil
}

func (asicdClientMgr *CPSAsicdClntMgr) DeletePort(cfg *objects.Port) (bool, error) {
	return true, nil
}

func (asicdClientMgr *CPSAsicdClntMgr) GetBulkPort(fromIdx, count int) (*asicdClntDefs.PortGetInfo, error) {
	var retObj asicdClntDefs.PortGetInfo
	return &retObj, nil
}

func (asicdClientMgr *CPSAsicdClntMgr) GetPort(IntfRef string) (*objects.Port, error) {
	return nil, nil
}

func (asicdClientMgr *CPSAsicdClntMgr) GetBulkPortState(fromIdx, count int) (*asicdClntDefs.PortStateGetInfo, error) {
	var retObj asicdClntDefs.PortStateGetInfo
	return &retObj, nil
}

func (asicdClientMgr *CPSAsicdClntMgr) GetPortState(IntfRef string) (*objects.PortState, error) {
	return nil, nil
}

func (asicdClientMgr *CPSAsicdClntMgr) GetBulkMacTableEntryState(fromIndex int, count int) (*asicdClntDefs.MacTableEntryStateGetInfo, error) {
	var retObj asicdClntDefs.MacTableEntryStateGetInfo
	return &retObj, nil
}

func (asicdClientMgr *CPSAsicdClntMgr) GetMacTableEntryState(MacAddr string) (*objects.MacTableEntryState, error) {
	return nil, nil
}

func (asicdClientMgr *CPSAsicdClntMgr) GetBulkIPv4RouteHwState(fromIndex int, count int) (*asicdClntDefs.IPv4RouteHwStateGetInfo, error) {
	var retObj asicdClntDefs.IPv4RouteHwStateGetInfo
	return &retObj, nil
}

func (asicdClientMgr *CPSAsicdClntMgr) GetIPv4RouteHwState(DestinationNw string) (*objects.IPv4RouteHwState, error) {
	return nil, nil
}

func (asicdClientMgr *CPSAsicdClntMgr) GetBulkIPv6RouteHwState(fromIndex int, count int) (*asicdClntDefs.IPv6RouteHwStateGetInfo, error) {
	var retObj asicdClntDefs.IPv6RouteHwStateGetInfo
	return &retObj, nil
}

func (asicdClientMgr *CPSAsicdClntMgr) GetIPv6RouteHwState(DestinationNw string) (*objects.IPv6RouteHwState, error) {
	return nil, nil
}

func (asicdClientMgr *CPSAsicdClntMgr) GetBulkArpEntryHwState(fromIndex int, count int) (*asicdClntDefs.ArpEntryHwStateGetInfo, error) {
	var retObj asicdClntDefs.ArpEntryHwStateGetInfo
	return &retObj, nil
}

func (asicdClientMgr *CPSAsicdClntMgr) GetArpEntryHwState(IpAddr string) (*objects.ArpEntryHwState, error) {
	return nil, nil
}

func (asicdClientMgr *CPSAsicdClntMgr) GetBulkNdpEntryHwState(fromIndex int, count int) (*asicdClntDefs.NdpEntryHwStateGetInfo, error) {
	var retObj asicdClntDefs.NdpEntryHwStateGetInfo
	return &retObj, nil
}

func (asicdClientMgr *CPSAsicdClntMgr) GetNdpEntryHwState(IpAddr string) (*objects.NdpEntryHwState, error) {
	return nil, nil
}

func (asicdClientMgr *CPSAsicdClntMgr) CreateLogicalIntf(cfg *objects.LogicalIntf) (bool, error) {
	return true, nil
}

func (asicdClientMgr *CPSAsicdClntMgr) UpdateLogicalIntf(origCfg, newCfg *objects.LogicalIntf, attrset []bool, op []*objects.PatchOpInfo) (bool, error) {
	return true, nil
}

func (asicdClientMgr *CPSAsicdClntMgr) DeleteLogicalIntf(cfg *objects.LogicalIntf) (bool, error) {
	return true, nil
}

func (asicdClientMgr *CPSAsicdClntMgr) GetBulkLogicalIntfState(fromIndex int, count int) (*asicdClntDefs.LogicalIntfStateGetInfo, error) {
	var retObj asicdClntDefs.LogicalIntfStateGetInfo
	return &retObj, nil
}

func (asicdClientMgr *CPSAsicdClntMgr) GetLogicalIntfState(Name string) (*objects.LogicalIntfState, error) {
	return nil, nil
}

func (asicdClientMgr *CPSAsicdClntMgr) CreateSubIPv4Intf(cfg *objects.SubIPv4Intf) (bool, error) {
	return true, nil
}

func (asicdClientMgr *CPSAsicdClntMgr) UpdateSubIPv4Intf(origCfg, newCfg *objects.SubIPv4Intf, attrset []bool, op []*objects.PatchOpInfo) (bool, error) {
	return true, nil
}

func (asicdClientMgr *CPSAsicdClntMgr) DeleteSubIPv4Intf(cfg *objects.SubIPv4Intf) (bool, error) {
	return true, nil
}

func (asicdClientMgr *CPSAsicdClntMgr) GetBulkSubIPv4IntfState(fromIndex int, count int) (*asicdClntDefs.SubIPv4IntfStateGetInfo, error) {
	var retObj asicdClntDefs.SubIPv4IntfStateGetInfo
	return &retObj, nil
}

func (asicdClientMgr *CPSAsicdClntMgr) GetSubIPv4IntfState(IntfRef string, Type string) (*objects.SubIPv4IntfState, error) {
	return nil, nil
}

func (asicdClientMgr *CPSAsicdClntMgr) CreateIPv6Intf(cfg *objects.IPv6Intf) (bool, error) {
	return true, nil
}

func (asicdClientMgr *CPSAsicdClntMgr) UpdateIPv6Intf(origCfg, newCfg *objects.IPv6Intf, attrset []bool, op []*objects.PatchOpInfo) (bool, error) {
	return true, nil
}

func (asicdClientMgr *CPSAsicdClntMgr) DeleteIPv6Intf(cfg *objects.IPv6Intf) (bool, error) {
	return true, nil
}

func (asicdClientMgr *CPSAsicdClntMgr) GetBulkIPv6IntfState(fromIndex int, count int) (*asicdClntDefs.IPv6IntfStateGetInfo, error) {
	var retObj asicdClntDefs.IPv6IntfStateGetInfo
	return &retObj, nil
}

func (asicdClientMgr *CPSAsicdClntMgr) GetIPv6IntfState(IntfRef string) (*objects.IPv6IntfState, error) {
	return nil, nil
}

func (asicdClientMgr *CPSAsicdClntMgr) CreateSubIPv6Intf(cfg *objects.SubIPv6Intf) (bool, error) {
	return true, nil
}

func (asicdClientMgr *CPSAsicdClntMgr) UpdateSubIPv6Intf(origCfg, newCfg *objects.SubIPv6Intf, attrset []bool, op []*objects.PatchOpInfo) (bool, error) {
	return true, nil
}

func (asicdClientMgr *CPSAsicdClntMgr) DeleteSubIPv6Intf(cfg *objects.SubIPv6Intf) (bool, error) {
	return true, nil
}

func (asicdClientMgr *CPSAsicdClntMgr) GetBulkSubIPv6IntfState(fromIndex int, count int) (*asicdClntDefs.SubIPv6IntfStateGetInfo, error) {
	var retObj asicdClntDefs.SubIPv6IntfStateGetInfo
	return &retObj, nil
}

func (asicdClientMgr *CPSAsicdClntMgr) GetSubIPv6IntfState(IntfRef string, Type string) (*objects.SubIPv6IntfState, error) {
	return nil, nil
}

func (asicdClientMgr *CPSAsicdClntMgr) GetBulkBufferPortStatState(fromIndex int, count int) (*asicdClntDefs.BufferPortStatStateGetInfo, error) {
	var retObj asicdClntDefs.BufferPortStatStateGetInfo
	return &retObj, nil
}

func (asicdClientMgr *CPSAsicdClntMgr) GetBufferPortStatState(IntfRef string) (*objects.BufferPortStatState, error) {
	return nil, nil
}

func (asicdClientMgr *CPSAsicdClntMgr) GetBulkBufferGlobalStatState(fromIndex int, count int) (*asicdClntDefs.BufferGlobalStatStateGetInfo, error) {
	var retObj asicdClntDefs.BufferGlobalStatStateGetInfo
	return &retObj, nil
}

func (asicdClientMgr *CPSAsicdClntMgr) GetBufferGlobalStatState(DeviceId uint32) (*objects.BufferGlobalStatState, error) {
	return nil, nil
}

func (asicdClientMgr *CPSAsicdClntMgr) CreateAclGlobal(cfg *objects.AclGlobal) (bool, error) {
	return true, nil
}

func (asicdClientMgr *CPSAsicdClntMgr) UpdateAclGlobal(origCfg, newCfg *objects.AclGlobal, attrset []bool, op []*objects.PatchOpInfo) (bool, error) {
	return true, nil
}

func (asicdClientMgr *CPSAsicdClntMgr) DeleteAclGlobal(cfg *objects.AclGlobal) (bool, error) {
	return true, nil
}

func (asicdClientMgr *CPSAsicdClntMgr) CreateAcl(cfg *objects.Acl) (bool, error) {
	return true, nil
}

func (asicdClientMgr *CPSAsicdClntMgr) UpdateAcl(origCfg, newCfg *objects.Acl, attrset []bool, op []*objects.PatchOpInfo) (bool, error) {
	return true, nil
}

func (asicdClientMgr *CPSAsicdClntMgr) DeleteAcl(cfg *objects.Acl) (bool, error) {
	return true, nil
}

func (asicdClientMgr *CPSAsicdClntMgr) CreateAclIpv4Filter(cfg *objects.AclIpv4Filter) (bool, error) {
	return true, nil
}

func (asicdClientMgr *CPSAsicdClntMgr) UpdateAclIpv4Filter(origCfg, newCfg *objects.AclIpv4Filter, attrset []bool, op []*objects.PatchOpInfo) (bool, error) {
	return true, nil
}

func (asicdClientMgr *CPSAsicdClntMgr) DeleteAclIpv4Filter(cfg *objects.AclIpv4Filter) (bool, error) {
	return true, nil
}

func (asicdClientMgr *CPSAsicdClntMgr) CreateAclMacFilter(cfg *objects.AclMacFilter) (bool, error) {
	return true, nil
}

func (asicdClientMgr *CPSAsicdClntMgr) UpdateAclMacFilter(origCfg, newCfg *objects.AclMacFilter, attrset []bool, op []*objects.PatchOpInfo) (bool, error) {
	return true, nil
}

func (asicdClientMgr *CPSAsicdClntMgr) DeleteAclMacFilter(cfg *objects.AclMacFilter) (bool, error) {
	return true, nil
}

func (asicdClientMgr *CPSAsicdClntMgr) CreateAclIpv6Filter(cfg *objects.AclIpv6Filter) (bool, error) {
	return true, nil
}

func (asicdClientMgr *CPSAsicdClntMgr) UpdateAclIpv6Filter(origCfg, newCfg *objects.AclIpv6Filter, attrset []bool, op []*objects.PatchOpInfo) (bool, error) {
	return true, nil
}

func (asicdClientMgr *CPSAsicdClntMgr) DeleteAclIpv6Filter(cfg *objects.AclIpv6Filter) (bool, error) {
	return true, nil
}

func (asicdClientMgr *CPSAsicdClntMgr) GetBulkAclState(fromIndex int, count int) (*asicdClntDefs.AclStateGetInfo, error) {
	var retObj asicdClntDefs.AclStateGetInfo
	return &retObj, nil
}

func (asicdClientMgr *CPSAsicdClntMgr) GetAclState(AclName string) (*objects.AclState, error) {
	return nil, nil
}

func (asicdClientMgr *CPSAsicdClntMgr) GetBulkLinkScopeIpState(fromIndex int, count int) (*asicdClntDefs.LinkScopeIpStateGetInfo, error) {
	var retObj asicdClntDefs.LinkScopeIpStateGetInfo
	return &retObj, nil
}

func (asicdClientMgr *CPSAsicdClntMgr) GetLinkScopeIpState(LinkScopeIp string) (*objects.LinkScopeIpState, error) {
	return nil, nil
}

func (asicdClientMgr *CPSAsicdClntMgr) GetBulkCoppStatState(fromIndex int, count int) (*asicdClntDefs.CoppStatStateGetInfo, error) {
	var retObj asicdClntDefs.CoppStatStateGetInfo
	return &retObj, nil
}

func (asicdClientMgr *CPSAsicdClntMgr) GetCoppStatState(Protocol string) (*objects.CoppStatState, error) {
	return nil, nil
}
