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
	"errors"
	"utils/clntUtils/clntIntfs"
	"utils/logging"
)

type Port struct {
	IntfRef     string
	IfIndex     int32
	Description string
	AdminState  string
	OperState	string
	MacAddr     string
	Speed       int32
	Duplex      string
}

type Vlan struct {
	VlanName string
	OperState string
	IfIndex int32
	AdminState string
	IntfList map[int32]bool
	UntagIntfList map[int32]bool
}

type CPSAsicdClntMgr struct {
	//NCtrlCh   chan bool
	//ClientHdl *asicdServices.ASICDServicesClient
	PortDB []Port
	IfIdxToIfIdMap map[int32]int32
	IfIdxToIfTypeMap map[int32]int32
	VlanDB map[int32]Vlan
	VlanList []int32
}

var Logger logging.LoggerIntf

func NewAsicdClntInit(clntInitParams *clntIntfs.BaseClntInitParams) (*CPSAsicdClntMgr, error) {
	//var err error
	cpsAsicdClntMgr := new(CPSAsicdClntMgr)
	Logger = clntInitParams.Logger
	cpsAsicdClntMgr.IfIdxToIfIdMap = make(map[int32]int32)
	cpsAsicdClntMgr.IfIdxToIfTypeMap = make(map[int32]int32)
	cpsAsicdClntMgr.VlanDB = make(map[int32]Vlan)
	err := cpsAsicdClntMgr.GetAllPortConfig()
	if err != nil {
		return nil, errors.New("Failed GetAllPortConfig()")
	}
	/*
		err = cpsAsicdClntMgr.GetAsicdThriftClientHdl(clntInitParams)
		if cpsAsicdClntMgr.ClientHdl == nil || err != nil {
			Logger.Err("Unable Initialize Asicd Client", err)
			return nil, errors.New(fmt.Sprintln("Unable Initialize Asicd Client", err))
		}
		if clntInitParams.NHdl != nil {
			err = cpsAsicdClntMgr.InitCPSAsicdSubscriber(clntInitParams)
			if err != nil {
				Logger.Err("Unable Initialize Asicd Client", err)
				cpsAsicdClntMgr.DeinitAsicdThriftClientHdl()
				return nil, errors.New(fmt.Sprintln("Unable Initialize Asicd Client", err))
			}
		}
	*/
	return cpsAsicdClntMgr, nil
}

func (asicdClientMgr *CPSAsicdClntMgr) AsicdClntDeinit() {
	/*
		if asicdClientMgr.ClientHdl != nil {
			asicdClientMgr.DeinitAsicdThriftClientHdl()
			asicdClientMgr.ClientHdl = nil
		}
		if asicdClientMgr.NCtrlCh != nil {
			asicdClientMgr.DeinitCPSAsicdSubscriber()
			asicdClientMgr.NCtrlCh = nil
		}
	*/
	asicdClientMgr.IfIdxToIfIdMap = nil
	asicdClientMgr.IfIdxToIfTypeMap = nil
}
