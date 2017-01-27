#ifndef CPS_H
#define CPS_H
#include <stdio.h>
#include <string.h>
#include <stdlib.h>
#include <stdint.h>
#include <syslog.h>
#include <net/if.h>
#include <sys/socket.h>
#include <netinet/in.h>
#include <arpa/inet.h>
#include <sonic/cps_api_key.h>
#include <sonic/dell-interface.h>
#include <sonic/dell-base-vlan.h>
#include <sonic/dell-base-if-vlan.h>
#include <sonic/iana-if-type.h>
#include <sonic/dell-base-if.h>
#include <sonic/cps_class_map.h>
#include <sonic/cps_api_object_key.h>
#include <sonic/dell-base-ip.h>
#include <sonic/dell-base-routing.h>

char** MakeCharArray(int size);
void SetArrayString(char **a, char *s, int n);
void FreeCharArray(char **a, int size);

cps_api_return_code_t CPSCreateVlan(uint32_t vlanId, uint32_t numOfTagPorts, char **tagPorts, uint32_t numOfUntagPorts, char **untagPorts);
cps_api_return_code_t CPSDeleteVlan(uint32_t vlanId);
cps_api_return_code_t CPSCreateIPv4Intf(char *intfRef, char *ipAddr, uint32_t prefix);
cps_api_return_code_t CPSDeleteIPv4Intf(char *intfRef, char *ipAddr, uint32_t prefix);
#endif /* CPS_H */
