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
	"github.com/vishvananda/netlink"
	"strings"
)

func GetAllPortCfg() (ret int, count uint8, portCfg []PortCfgGo) {
	linkList, err := netlink.LinkList()
	if err != nil {
		return -1, 0, portCfg
	}
	for _, link := range linkList {
		linkAttr := link.Attrs()
		if strings.HasPrefix(linkAttr.Name, "e101-") {
			var pCfg PortCfgGo
			pCfg.Name = linkAttr.Name
			portCfg = append(portCfg, pCfg)
			count++
		}
	}
	return 0, count, portCfg
}
