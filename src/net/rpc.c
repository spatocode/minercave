#include <stdio.h>
#include <stdbool.h>

struct RPCClient {
	int sickRate;
	int successRate;
	int accepts;
	int rejects;
	int lastSubmitAt;
	int failsCount;
	char login[20];
	char password[20];
	char name[20];
	bool sick;
};
