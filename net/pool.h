#ifndef POOL_H_
#define POOL_H_


typedef struct _Port {
	int difficulty;
	std::string host;
	int port;
	int maxconn;
} Port;


typedef struct _Stratum {
	std::string timeout;
	Port ports;
} Stratum;


typedef struct _Upstream {
	std::string name;
	std::string host;
	int port;
	std::string timeout;
} Upstream;


typedef struct _Config {
	std::string address;
	Stratum stratum;
	Upstream upstream[];
	bool bypassAddressValidation;
	bool bypassShareValidation;
	std::string blockRefreshInterval;
	std::string upstreamCheckInterval;
	std::string estimationWindow;
	std::string luckWindow;
	std::string largeLuckWindow;
	int threads;
	std::string newrelicKey;
	char newrelicName;
	bool newrelicEnabled;
	bool newrelicVerbose;
} Config;

#endif








