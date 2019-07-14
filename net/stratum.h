#ifndef STRATUM_H
#define STRATUM_H
#include <stdio.h>
#include <memory>
#include "pool.h"
#include "client.h"

#define MaxReqSize 10 * 1024


typedef struct _BlockEntry {
	int height;
	float variance;
	char hash[20];
} BlockEntry;


namespace minercave {
	
	
class Stratum {
	public:
		Stratum registerStratum(Config cfg) {
			config = cfg;
			//blockStats = BlockEntry;
			
			for(int i=0; i < (sizeof(cfg.upstream)/sizeof(cfg.upstream[0])); i++) {
				Client client; 
				client.getClient(&cfg.upstream[i]);
				upstreams[i] = client;
				printf("Upstream: %s => %s", client.name(), client.url());
			}
			
			printf("Default upstream: %s => %s", rpc.name(), rpc.url());
			return *this;
		};
		
		Client rpc();
		void listen();
		void checkUpstreams();
		void currentBlock();
		void setDeadline();
		void removeMiner();
		void registerMiner();
		bool isActive();
		void removeSession();
		void registerSession();
		void handleClient();
		void handleLoginRPC();
		void handleGetJobRPC();
		void handleSubmitRPC();
		void handleUnknownRPC();
		void broadcastNewJobs();
		void refreshBlockTemplate();
	
	private:
		int luckWindow;
		int luckLargeWindow;
		int roundShares;
		int upstream;
		Client upstreams[];
		Config config;
		BlockEntry blockStats;
};
	
}


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
	public:
		void sendError();
		void pushMessage();
		void sendResult();
		void handleMessage();
	private:
		int lastBlockHeight;
		char ip[20];
		EndPoint endpoint;
};


#endif

