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
	"utils/clntUtils/clntIntfs"
)

type processMsg func(uint8, []byte) (clntIntfs.NotifyMsg, error)

var AsicdMsgMap map[uint8]processMsg = map[uint8]processMsg{
	asicdClntDefs.NOTIFY_L2INTF_STATE_CHANGE:       processL2IntfStateNotifyMsg,
	asicdClntDefs.NOTIFY_IPV4_L3INTF_STATE_CHANGE:  processIPv4L3IntfStateNotifyMsg,
	asicdClntDefs.NOTIFY_IPV6_L3INTF_STATE_CHANGE:  processIPv6L3IntfStateNotifyMsg,
	asicdClntDefs.NOTIFY_VLAN_CREATE:               processVlanNotifyMsg,
	asicdClntDefs.NOTIFY_VLAN_DELETE:               processVlanNotifyMsg,
	asicdClntDefs.NOTIFY_VLAN_UPDATE:               processVlanNotifyMsg,
	asicdClntDefs.NOTIFY_LOGICAL_INTF_CREATE:       processLogicalIntfNotifyMsg,
	asicdClntDefs.NOTIFY_LOGICAL_INTF_DELETE:       processLogicalIntfNotifyMsg,
	asicdClntDefs.NOTIFY_LOGICAL_INTF_UPDATE:       processLogicalIntfNotifyMsg,
	asicdClntDefs.NOTIFY_IPV4INTF_CREATE:           processIPv4IntfNotifyMsg,
	asicdClntDefs.NOTIFY_IPV4INTF_DELETE:           processIPv4IntfNotifyMsg,
	asicdClntDefs.NOTIFY_IPV6INTF_CREATE:           processIPv6IntfNotifyMsg,
	asicdClntDefs.NOTIFY_IPV6INTF_DELETE:           processIPv6IntfNotifyMsg,
	asicdClntDefs.NOTIFY_LAG_CREATE:                processLagNotifyMsg,
	asicdClntDefs.NOTIFY_LAG_DELETE:                processLagNotifyMsg,
	asicdClntDefs.NOTIFY_LAG_UPDATE:                processLagNotifyMsg,
	asicdClntDefs.NOTIFY_IPV4NBR_MAC_MOVE:          processIPv4NbrMacMoveNotifyMsg,
	asicdClntDefs.NOTIFY_IPV6NBR_MAC_MOVE:          processIPv6NbrMacMoveNotifyMsg,
	asicdClntDefs.NOTIFY_IPV4_ROUTE_CREATE_FAILURE: processIPv4RouteAddDelFailNotifyMsg,
	asicdClntDefs.NOTIFY_IPV4_ROUTE_DELETE_FAILURE: processIPv4RouteAddDelFailNotifyMsg,
	asicdClntDefs.NOTIFY_IPV6_ROUTE_CREATE_FAILURE: processIPv6RouteAddDelFailNotifyMsg,
	asicdClntDefs.NOTIFY_IPV6_ROUTE_DELETE_FAILURE: processIPv6RouteAddDelFailNotifyMsg,
	asicdClntDefs.NOTIFY_VTEP_CREATE:               processVtepNotifyMsg,
	asicdClntDefs.NOTIFY_VTEP_DELETE:               processVtepNotifyMsg,
	//asicdClntDefs.NOTIFY_MPLSINTF_STATE_CHANGE:         processMplsIntfStateChangeNotifyMsg,
	//asicdClntDefs.NOTIFY_MPLSINTF_CREATE:               processMplsIntfNotifyMsg,
	//asicdClntDefs.NOTIFY_MPLSINTF_DELETE:               processMplsIntfNotifyMsg,
	asicdClntDefs.NOTIFY_PORT_CONFIG_MODE_CHANGE:       processPortConfigModeChgNotifyMsg,
	asicdClntDefs.NOTIFY_PORT_ATTR_CHANGE:              processPortAttrChangeNotifyMsg,
	asicdClntDefs.NOTIFY_IPV4VIRTUAL_INTF_CREATE:       processIPv4VirutalIntfNotifyMsg,
	asicdClntDefs.NOTIFY_IPV4VIRTUAL_INTF_DELETE:       processIPv4VirutalIntfNotifyMsg,
	asicdClntDefs.NOTIFY_IPV6VIRTUAL_INTF_CREATE:       processIPv6VirutalIntfNotifyMsg,
	asicdClntDefs.NOTIFY_IPV6VIRTUAL_INTF_DELETE:       processIPv6VirutalIntfNotifyMsg,
	asicdClntDefs.NOTIFY_IPV4_VIRTUALINTF_STATE_CHANGE: processIPv4VirtualIntfStateNotifyMsg,
	asicdClntDefs.NOTIFY_IPV6_VIRTUALINTF_STATE_CHANGE: processIPv6VirtualIntfStateNotifyMsg,
}

func processL2IntfStateNotifyMsg(rxMsgType uint8, rxMsg []byte) (clntIntfs.NotifyMsg, error) {
	var msg clntIntfs.NotifyMsg
	return msg, nil
}

func processIPv4L3IntfStateNotifyMsg(rxMsgType uint8, rxMsg []byte) (clntIntfs.NotifyMsg, error) {
	var msg clntIntfs.NotifyMsg
	return msg, nil
}

func processIPv6L3IntfStateNotifyMsg(rxMsgType uint8, rxMsg []byte) (clntIntfs.NotifyMsg, error) {
	var msg clntIntfs.NotifyMsg
	return msg, nil
}

func processIPv4VirtualIntfStateNotifyMsg(rxMsgType uint8, rxMsg []byte) (clntIntfs.NotifyMsg, error) {
	var msg clntIntfs.NotifyMsg
	return msg, nil
}

func processIPv6VirtualIntfStateNotifyMsg(rxMsgType uint8, rxMsg []byte) (clntIntfs.NotifyMsg, error) {
	var msg clntIntfs.NotifyMsg
	return msg, nil
}

func processVlanNotifyMsg(rxMsgType uint8, rxMsg []byte) (clntIntfs.NotifyMsg, error) {
	var msg clntIntfs.NotifyMsg
	return msg, nil
}

func processLogicalIntfNotifyMsg(rxMsgType uint8, rxMsg []byte) (clntIntfs.NotifyMsg, error) {
	var msg clntIntfs.NotifyMsg
	return msg, nil

}

func processIPv4IntfNotifyMsg(rxMsgType uint8, rxMsg []byte) (clntIntfs.NotifyMsg, error) {
	var msg clntIntfs.NotifyMsg
	return msg, nil
}

func processIPv6IntfNotifyMsg(rxMsgType uint8, rxMsg []byte) (clntIntfs.NotifyMsg, error) {
	var msg clntIntfs.NotifyMsg
	return msg, nil
}

func processIPv4VirutalIntfNotifyMsg(rxMsgType uint8, rxMsg []byte) (clntIntfs.NotifyMsg, error) {
	var msg clntIntfs.NotifyMsg
	return msg, nil
}

func processIPv6VirutalIntfNotifyMsg(rxMsgType uint8, rxMsg []byte) (clntIntfs.NotifyMsg, error) {
	var msg clntIntfs.NotifyMsg
	return msg, nil
}

func processLagNotifyMsg(rxMsgType uint8, rxMsg []byte) (clntIntfs.NotifyMsg, error) {
	var msg clntIntfs.NotifyMsg
	return msg, nil
}

func processIPv4NbrMacMoveNotifyMsg(rxMsgType uint8, rxMsg []byte) (clntIntfs.NotifyMsg, error) {
	var msg clntIntfs.NotifyMsg
	return msg, nil
}

func processIPv6NbrMacMoveNotifyMsg(rxMsgType uint8, rxMsg []byte) (clntIntfs.NotifyMsg, error) {
	var msg clntIntfs.NotifyMsg
	return msg, nil
}

func processIPv4RouteAddDelFailNotifyMsg(rxMsgType uint8, rxMsg []byte) (clntIntfs.NotifyMsg, error) {
	var msg clntIntfs.NotifyMsg
	//TODO: Discuss this with Madhavi
	return msg, nil
}

func processIPv6RouteAddDelFailNotifyMsg(rxMsgType uint8, rxMsg []byte) (clntIntfs.NotifyMsg, error) {
	var msg clntIntfs.NotifyMsg
	//TODO: Discuss this with Madhavi
	return msg, nil
}

func processVtepNotifyMsg(rxMsgType uint8, rxMsg []byte) (clntIntfs.NotifyMsg, error) {
	var msg clntIntfs.NotifyMsg
	return msg, nil

}

func processMplsIntfStateChangeNotifyMsg(rxMsgType uint8, rxMsg []byte) (clntIntfs.NotifyMsg, error) {
	var msg clntIntfs.NotifyMsg
	//TODO: Need to be done along with MPLS changes
	return msg, nil
}

func processMplsIntfNotifyMsg(rxMsgType uint8, rxMsg []byte) (clntIntfs.NotifyMsg, error) {
	var msg clntIntfs.NotifyMsg
	//TODO: Need to be done along with MPLS changes
	return msg, nil
}

func processPortConfigModeChgNotifyMsg(rxMsgType uint8, rxMsg []byte) (clntIntfs.NotifyMsg, error) {
	var msg clntIntfs.NotifyMsg
	return msg, nil
}

func processPortAttrChangeNotifyMsg(rxMsgType uint8, rxMsg []byte) (clntIntfs.NotifyMsg, error) {
	var msg clntIntfs.NotifyMsg
	return msg, nil
}

func (asicdClientMgr *CPSAsicdClntMgr) InitFSAsicdSubscriber(clntInitParams *clntIntfs.BaseClntInitParams) error {
	return nil
}

func (asicdClientMgr *CPSAsicdClntMgr) DeinitFSAsicdSubscriber() {
}
