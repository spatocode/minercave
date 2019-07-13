#include <iostream>
#include "stratum.h"
#include "pool.h"
#include "rpc.h"


void StratumServer::&newStratum(Config cfg) {
	this->config = cfg;
	this->blockStats = BlockEntry;
	this.upstreams = RPCClient;
	return *this;
}


void StratumServer::listen() {
	
}


void EndPoint::&newEndpoint(Pool cfg) {
	this->config = cfg;
	this->instanceId = 4;
	return *this;
}


void EndPoint::listen() {
	
}


