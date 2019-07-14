#include <stdio.h>
#include <cassert>
#include <ctime>
#include "client.h"
#include "pool.h"


Client* minercave::Client::getClient(Upstream* cfg) {
	const size_t size = cfg.host.size() + 8;
	assert(size > 8);
	
	char url[size];
	snprintf(url, size - 1, "http://%s:%d/json_client", cfg.host, cfg.port);
	
	m_name = cfg.name;
	m_url = url;
	return *this;
}
