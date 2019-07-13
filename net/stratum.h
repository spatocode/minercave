#ifndef STRATUM_H
#define STRATUM_H
#include <time.h>
#include "pool.h"
#include "rpc.h"

#define MaxReqSize 10 * 1024


typedef struct _BlockEntry {
	int height;
	float variance;
	char hash[20];
} BlockEntry;


class StratumServer {
	public:
		StratumServer &newStratum(Config cfg);
		void listen();
	
	private:
		int luckWindow;
		int luckLargeWindow;
		int roundShares;
		int upstream;
		RPCClient upstreams;
		Config config;
		BlockEntry blockStats;
};


class EndPoint {
	public:
		EndPoint &newEndpoint(Port cfg);
		void listen();
		
	private:
		unsigned int jobSequence;
		Port config;
		short instanceId[];
		unsigned int extraNonce;
		char targetHex[20];
};

class Session {
	private:
		int lastBlockHeight;
		char ip[20];
		EndPoint endpoint;
};


#endif

