#include <stdio.h>
#include "pool.h"
#define MaxReqSize 10 * 1024

struct BlockEntry {
	int height;
	float variance;
	char hash[20];
};

struct StratumServer {
	int luckWindow;
	int luckLargeWindow;
	int roundShares;
	int upstream;
	struct RPCClient upstreams;
	struct Config config;
	struct BlockEntry blockStats;
	struct MinersMap miners;
};

struct Endpoint {
	uint jobSequence;
	struct Port config;
	byte instanceId[];
	uint extraNonce;
	char targetHex[20];
};

struct Session {
	int lastBlockHeight;
	char[20] ip;
	struct Job validJobs;
	struct Endpoint endpoint;
};






