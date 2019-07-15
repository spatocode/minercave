#include <stdio.h>
#include <cassert>
#include "endpoint.h"
#include "pool.h"


minercave::EndPoint::EndPoint(Port cfg) : 
	config(cfg),
	difficulty(config.difficulty),
	instanceId({0,0,0,0})
{
}


minercave::EndPoint::listen() {
	const size_t size = config.host.size() + 8;
	assert(size > 8);
	
	char* bindAddr = new[size]();
	snprintf(bindAddr, size - 1, "%s:%s", config.host, config.port);
}









