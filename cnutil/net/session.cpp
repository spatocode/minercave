#include "session.h"


minercave::Session::Session(char ip, EndPoint ep) :
	m_ip(ip),
	*m_endpoint(ep)
{
}


void minercave::Session::handleMessage() {
	
}


void minercave::Session::pushMessage() {
	
}


void minercave::Session::sendError() {
	
}


void minercave::Session::sendResult() {
	
}
