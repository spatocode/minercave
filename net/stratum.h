#ifndef STRATUM_H
#define STRATUM_H
#include <stdio.h>
#include <memory>
#include "pool.h"
#include "client.h"
#include "endpoint.h"

#define MaxReqSize 10 * 1024


typedef struct _BlockEntry {
	int height;
	float variance;
	char hash[20];
} BlockEntry;


namespace minercave {
	
	
class Stratum {
	public:
		Stratum(Pool::Config cfg);
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
		int m_luck;
		int m_large;
		int m_roundShares;
		int m_upstream;
		Client m_upstreams[];
		Pool::Config m_config;
		BlockEntry m_blockStats;
};
	
}


class Session {
	public:
		void sendError();
		void pushMessage();
		void sendResult();
		void handleMessage();
	private:
		int m_lastBlockHeight;
		char m_ip[20];
		minercave::EndPoint m_endpoint;
};


#endif

