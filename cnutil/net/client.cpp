#include <stdio.h>
#include <cassert>
#include <ctime>
#include "client.h"
#include "pool.h"


minercave::Client::Client(Upstream* cfg) {
	const size_t size = cfg.host.size() + 8;
	assert(size > 8);
	
	m_url = new[size]();
	snprintf(m_url, size - 1, "http://%s:%d/client", cfg.host, cfg.port);
	
	m_name = cfg.name;
}


minercave::Client::~Client() {
	delete m_url;
};


