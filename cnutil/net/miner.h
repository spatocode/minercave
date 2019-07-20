#include <stdbool.h>

struct Job {
	int height;
	char id[20];
	uint extraNonce;
};

struct Miner {
	int lastBeat;
	int startedAt;
	int validShares;
	int invalidShares;
	int staleShares;
	int accepts;
	int rejects;
	char id[20];
	char ip[20];
};
