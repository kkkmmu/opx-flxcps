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
//   This is a auto-generated file, please do not edit!
// _______   __       __________   ___      _______.____    __    ____  __  .___________.  ______  __    __
// |   ____||  |     |   ____\  \ /  /     /       |\   \  /  \  /   / |  | |           | /      ||  |  |  |
// |  |__   |  |     |  |__   \  V  /     |   (----  \   \/    \/   /  |  |  ---|  |---- |  ,---- |  |__|  |
// |   __|  |  |     |   __|   >   <       \   \      \            /   |  |     |  |     |  |     |   __   |
// |  |     |  `----.|  |____ /  .  \  .----)   |      \    /\    /    |  |     |  |     |  `----.|  |  |  |
// |__|     |_______||_______/__/ \__\ |_______/        \__/  \__/     |__|     |__|      \______||__|  |__|
//

#ifndef CPS_H
#define CPS_H
#include <stdio.h>
#include <string.h>
#include <stdlib.h>
#include <stdint.h>
#include <syslog.h>
#include <net/if.h>
#include <arpa/inet.h>
#include <sys/socket.h>
#include <netinet/in.h>
#include <opx/cps_api_key.h>
#include <opx/dell-base-ip.h>
#include <opx/iana-if-type.h>
#include <opx/dell-base-if.h>
#include <opx/iana-if-type.h>
#include <opx/cps_class_map.h>
#include <opx/cps_api_errors.h>
#include <opx/dell-interface.h>
#include <opx/cps_api_events.h>
#include <opx/cps_api_object.h>
#include <opx/ietf-interfaces.h>
#include <opx/dell-base-if-phy.h>
#include <opx/dell-base-routing.h>
#include <opx/dell-base-if-vlan.h>
#include <opx/cps_api_object_key.h>
//#include <opx/dell-base-phy-interface.h>
#include <opx/cps_api_operation_tools.h>

//Struct defs
typedef struct portCfg_s {
	char PortName[IF_NAMESIZE];
} PortCfg_t;

typedef struct portState_s {
	uint64_t IfInOctets; //IF_INTERFACES_STATE_INTERFACE_STATISTICS_IN_OCTETS
	uint64_t IfInUcastPkts; //IF_INTERFACES_STATE_INTERFACE_STATISTICS_IN_UNICAST_PKTS
	uint64_t IfInDiscards; //IF_INTERFACES_STATE_INTERFACE_STATISTICS_IN_DISCARDS
	uint64_t IfInErrors; //IF_INTERFACES_STATE_INTERFACE_STATISTICS_IN_ERRORS
	uint64_t IfInUnknownProtos; //IF_INTERFACES_STATE_INTERFACE_STATISTICS_IN_UNKNOWN_PROTOS
	uint64_t IfOutOctets; //IF_INTERFACES_STATE_INTERFACE_STATISTICS_OUT_OCTETS
	uint64_t IfOutUcastPkts; //IF_INTERFACES_STATE_INTERFACE_STATISTICS_OUT_UNICAST_PKTS
	uint64_t IfOutDiscards; //IF_INTERFACES_STATE_INTERFACE_STATISTICS_OUT_DISCARDS
	uint64_t IfOutErrors; //IF_INTERFACES_STATE_INTERFACE_STATISTICS_OUT_ERRORS
	uint64_t IfEtherUnderSizePktCnt; //DELL_IF_IF_INTERFACES_STATE_INTERFACE_STATISTICS_ETHER_UNDERSIZE_PKTS
	uint64_t IfEtherOverSizePktCnt; //DELL_IF_IF_INTERFACES_STATE_INTERFACE_STATISTICS_ETHER_OVERSIZE_PKTS
	uint64_t IfEtherFragments; //DELL_IF_IF_INTERFACES_STATE_INTERFACE_STATISTICS_ETHER_FRAGMENTS
	uint64_t IfEtherCRCAlignError; //DELL_IF_IF_INTERFACES_STATE_INTERFACE_STATISTICS_ETHER_CRC_ALIGN_ERRORS
	uint64_t IfEtherJabber; //DELL_IF_IF_INTERFACES_STATE_INTERFACE_STATISTICS_ETHER_JABBERS
	uint64_t IfEtherPkts; //DELL_IF_IF_INTERFACES_STATE_INTERFACE_STATISTICS_ETHER_PKTS
	uint64_t IfEtherMCPkts; //DELL_IF_IF_INTERFACES_STATE_INTERFACE_STATISTICS_ETHER_MULTICAST_PKTS
	uint64_t IfEtherBcastPkts; //DELL_IF_IF_INTERFACES_STATE_INTERFACE_STATISTICS_ETHER_BROADCAST_PKTS
	uint64_t IfEtherPkts64OrLessOctets;
	uint64_t IfEtherPkts65To127Octets;
	uint64_t IfEtherPkts128To255Octets;
	uint64_t IfEtherPkts256To511Octets;
	uint64_t IfEtherPkts512To1023Octets;
	uint64_t IfEtherPkts1024To1518Octets;
} PortState_t;

char** MakeCharArray(int size);
void SetArrayString(char **a, char *s, int n);
void FreeCharArray(char **a, int size);

//Config/State related APIs
cps_api_return_code_t CPSCreateVlan(uint32_t vlanId, uint32_t numOfTagPorts, char **tagPorts, uint32_t numOfUntagPorts, char **untagPorts);
cps_api_return_code_t CPSDeleteVlan(uint32_t vlanId);
cps_api_return_code_t CPSCreateIPv4Intf(char *intfRef, char *ipAddr, uint32_t prefix);
cps_api_return_code_t CPSDeleteIPv4Intf(char *intfRef, char *ipAddr, uint32_t prefix);
cps_api_return_code_t CPSCreateIPv4Route(char *destNw, uint32_t prefix, uint32_t numOfNH, char **nhIpList);
cps_api_return_code_t CPSDeleteIPv4Route(char *destNw, uint32_t prefix);
cps_api_return_code_t CPSCreateIPv4Neighbor(char *nbrIp, char *intf, uint8_t macAddr[6]);
cps_api_return_code_t CPSDeleteIPv4Neighbor(char *nbrIp, char *intf, uint8_t macAddr[6]);
cps_api_return_code_t CPSGetAllPortCfg(PortCfg_t *portCfg, uint8_t *count);
cps_api_return_code_t CPSGetPortState(char *ifName, PortState_t *portState);
cps_api_return_code_t CPSSetPortAdminState(char *intfRef, uint8_t val, uint8_t an);

//Notification related APIs
//C functions
cps_api_return_code_t CPSRegisterNotificationHandler();
void CPSUnregisterNotificationHandler();

//Go functions
void HandleLinkNotifications(int ifindex, int oper_state);

#endif /* CPS_H */
