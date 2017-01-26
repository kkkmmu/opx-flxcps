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
	"utils/clntUtils/clntIntfs"
	"utils/logging"
)

type CPSAsicdClntMgr struct {
	//NCtrlCh   chan bool
	//ClientHdl *asicdServices.ASICDServicesClient
}

var Logger logging.LoggerIntf

func NewAsicdClntInit(clntInitParams *clntIntfs.BaseClntInitParams) (*CPSAsicdClntMgr, error) {
	//var err error
	cpsAsicdClntMgr := new(CPSAsicdClntMgr)
	Logger = clntInitParams.Logger
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
}
