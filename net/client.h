#ifndef CLIENT_H_
#define CLIENT_H_

#include <iostream>
#include "pool.h"


typedef struct _GetBlockReply {
	int difficulty;
	int height;
	std::string prevHash;
	std::string blob;
	int reservedOffset;
} GetBlockReply;


typedef struct _GetInfoReply {
	int height;
	char status[20];
	int poolSize;
	int incomingConnections;
	int outgoingConnections;
} GetInfoReply;


namespace minercave {
	
	
class Client {
	public:
		std::string name() { return m_name; };
		std::string url() { return m_url; };
		Client getClient(Upstream* cfg);
		GetBlockReply getBlock();
		GetInfoReply getInfo();
		GetInfoReply info();
		GetInfoReply updateInfo();
		void markAlive();
		void markSick();
		bool check(int size, std::string address);
		bool sick();
		void submit(std::string);
		void doPost(std::string url, std::string method, int params);
		void close();
		void startTimeout();
		
	private:
		int sickRate;
		int successRate;
		int accepts;
		int rejects;
		int lastSubmitAt;
		int failsCount;
		std::string login;
		std::string m_url;
		std::string m_password;
		std::string m_name;
		std::string m_client;
		bool m_sick;
};
	
	
}


#endif








