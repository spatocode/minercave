#include <iostream>
#include "stratum.h"


minercave::Stratum::Stratum(Config cfg) {
	m_config = cfg;
	//m_blockStats = BlockEntry;
			
	for(int i=0; i < (sizeof(cfg.upstream)/sizeof(cfg.upstream[0])); i++) {
		Client client(&cfg.upstream[i]);
		m_upstreams[i] = client.get();
		printf("Upstream: %s => %s", client.name(), client.url());
	}
			
	printf("Default upstream: %s => %s", rpc.name(), rpc.url());
}


void minercave::Stratum::listen() {
	for(int i=0; i < m_config.stratum.ports; i++) {
		EndPoint endpoint(m_config.stratum.ports[i]);
		Stratum ss;
		endpoint.listen(ss);
	}
}

void minercave::StratumServer::rpc() {
	
}

