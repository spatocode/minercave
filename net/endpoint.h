#ifndef ENDPOINT_H
#define ENDPOINT_H


#include "pool.h"


namespace minercave {


class EndPoint {
	public:
		EndPoint(Pool::Port cfg);
		~EndPoint();
		void listen();
		
	private:
		unsigned int m_jobSequence;
		Pool::Port m_config;
		short m_instanceId[4];
		unsigned int m_extraNonce;
		std::string m_targetHex;
		int m_difficulty;
};


}

#endif
