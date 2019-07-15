#ifndef ENDPOINT_H
#define ENDPOINT_H


#include "pool.h"


namespace minercave {


class EndPoint {
	public:
		EndPoint(Port cfg);
		void listen();
		
	private:
		unsigned int jobSequence;
		Pool::Port config;
		short instanceId[];
		unsigned int extraNonce;
		std::string targetHex;
		int difficulty;
};


}

#endif
