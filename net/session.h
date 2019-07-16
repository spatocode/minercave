#ifndef SESSION_H
#define SESSION_H

#include "endpoint.h"


namespace minercave {
	

class Session {
	public:
		Session(char ip, EndPoint ep);
		void sendError();
		void pushMessage();
		void sendResult();
		void handleMessage();
	private:
		int m_lastBlockHeight;
		char m_ip[20];
		minercave::EndPoint *m_endpoint;
};


}

#endif
