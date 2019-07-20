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
		Client(Pool::Upstream* cfg);
		~Client();
		inline const std::string name() const { return m_name; };
		inline const char *url() const { return m_url; };
		Client get() { return *this; };
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
		int m_sickRate;
		int m_successRate;
		int m_accepts;
		int m_rejects;
		int m_lastSubmitAt;
		int m_failsCount;
		char m_login[20];
		char m_url[46];
		char m_password[20];
		std::string m_name;
		std::string m_client;
		bool m_sick;
};
	
	
}


#endif








