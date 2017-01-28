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

//#include "cps.h"
import "C"

var nHdl clntIntfs.NotificationHdl

func (asicdClientMgr *CPSAsicdClntMgr) InitFSAsicdSubscriber(clntInitParams *clntIntfs.BaseClntInitParams) error {
	//Init notification hdl
	nHdl = clntInitParams.NHdl
	//Register with CPS to receive notifications
	ret := int(C.CPSRegisterNotificationHandler())
	if ret < 0 {
		Logger.Err("InitFSAsicdSubscriber failed")
	}
	return nil
}

func (asicdClientMgr *CPSAsicdClntMgr) DeinitFSAsicdSubscriber() {
	//Deregister notification handler
	C.CPSUnregisterNotificationHandler()
}

//export HandleLinkNotifications
func HandleLinkNotifications(ifIndex, operState C.int) {
	var ifState uint8

	switch operState {
	case C.IF_INTERFACES_STATE_INTERFACE_ADMIN_STATUS_UP:
		ifState = asicdClntDefs.INTF_STATE_UP

	case C.IF_INTERFACES_STATE_INTERFACE_ADMIN_STATUS_DOWN:
		ifState = asicdClntDefs.INTF_STATE_DOWN

	default:
		Logger.Err("Unrecognized operstate value received")
		return
	}
	msg := asicdClntDefs.L2IntfStateNotifyMsg{
		MsgType: asicdClntDefs.NOTIFY_L2INTF_STATE_CHANGE,
		IfIndex: int32(ifIndex),
		IfState: ifState,
	}
	Logger.Debug("Sending port state change notification for : ", ifIndex, ifState)
	if nHdl == nil {
		nHdl.ProcessNotification(msg)
	}
}
