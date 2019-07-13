#ifndef POOL_H_
#define POOL_H_

#include <stdbool.h>

typedef struct _Port {
	int difficulty;
	char host[20];
	int port;
	int maxconn;
} Port;

struct Stratum {
	char timeout[20];
	Port ports;
};

typedef struct _Config {
	char address[20];
	struct Stratum stratum;
	bool bypassAddressValidation;
	bool bypassShareValidation;
	char blockRefreshInterval[20];
	char upstreamCheckInterval[20];
	char estimationWindow[20];
	char luckWindow[20];
	char largeLuckWindow[20];
	int threads;
	char newrelicKey[20];
	char newrelicName[20];
	bool newrelicEnabled;
	bool newrelicVerbose;
} Config;

#endif








