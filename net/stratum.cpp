#include <iostream>
#include "stratum.h"
#include "pool.h"
#include "rpc.h"


void StratumServer::newStratum(Config cfg) {
	config = cfg;
	//blockStats = BlockEntry;
			
	for(int i=0; i < (sizeof(cfg.upstream)/sizeof(cfg.upstream[0])); i++) {
		Client client = minercave::Client::getClient(&cfg.upstream[i]);
		upstreams[i] = client;
		printf("Upstream: %s => %s", client.name(), client.url());
	}
			
	printf("Default upstream: %s => %s", rpc.name(), rpc.url());
	return *this;
}


void StratumServer::listen() {
	
}

void StratumServer::rpc() {
	
}


void EndPoint::&newEndpoint(Pool cfg) {
	this->config = cfg;
	this->instanceId = 4;
	return *this;
}


void EndPoint::listen() {
	
}


