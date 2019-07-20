#ifndef WIN32_LEAN_AND_MEAN
#define WIN32_LEAN_AND_MEAN
#endif

#include <stdio.h>
#include <cassert>
#include <windows.h>
#include <winsock2.h>
#include <ws2tcpip.h>
#include "endpoint.h"
#include "pool.h"

#pragma comment(lib, "Ws2_32.lib")

#define DEFAULT_BUFLEN 512


minercave::EndPoint::EndPoint(Port cfg) : 
	m_config(cfg),
	m_difficulty(m_config.difficulty),
	m_instanceId({0,0,0,0})
{
}


void minercave::EndPoint::listen(Stratum st) {	
	WSADATA wsaData;
    int iResult;

    SOCKET ListenSocket = INVALID_SOCKET;
    SOCKET ClientSocket = INVALID_SOCKET;

    struct addrinfo *result = NULL;
    struct addrinfo hints;
    struct sockaddr_in client;

    int iSendResult;
    char recvbuf[DEFAULT_BUFLEN];
    int recvbuflen = DEFAULT_BUFLEN;
    
    // Initialize Winsock
    if (WSAStartup(MAKEWORD(2,2), &wsaData) != 0) {
        printf("WSAStartup failed with error: %d\n", WSAGetLastError());
        return;
    }

    ZeroMemory(&hints, sizeof(hints));
    hints.ai_family = AF_INET;
    hints.ai_socktype = SOCK_STREAM;
    hints.ai_protocol = IPPROTO_TCP;
    hints.ai_flags = AI_PASSIVE;

    iResult = getaddrinfo( m_config.host, m_config.port, &hints, &result);
    if ( iResult != 0 ) {
        printf("Resolve address failed with error: %d\n", iResult);
        WSACleanup();
        return;
    }

    ListenSocket = socket(result->ai_family, result->ai_socktype, result->ai_protocol);
    if (ListenSocket == INVALID_SOCKET) {
        printf("Creating socket failed with error: %ld\n", WSAGetLastError());
        freeaddrinfo(result);
        WSACleanup();
        return;
    }

    iResult = bind( ListenSocket, result->ai_addr, (int)result->ai_addrlen);
    if (iResult == SOCKET_ERROR) {
        printf("Bind failed with error: %d\n", WSAGetLastError());
        freeaddrinfo(result);
        closesocket(ListenSocket);
        WSACleanup();
        return;
    }

    freeaddrinfo(result);

    iResult = listen(ListenSocket, m_config.maxconn);
    if (iResult == SOCKET_ERROR) {
        printf("Listen failed with error: %d\n", WSAGetLastError());
        closesocket(ListenSocket);
        WSACleanup();
        return;
    }
    else {
    	printf("Stratum listening on %s:%s", m_config.host, m_config.port);
	}

    do {
    	ClientSocket = accept(ListenSocket, (struct sockaddr *)&client, (int)result->ai_addrlen);
    	if (ClientSocket == INVALID_SOCKET) {
        	printf("Accepting connection failed with error: %d\n", WSAGetLastError());
        	closesocket(ListenSocket);
        	WSACleanup();
        	return;
    	}
    	
    	char *ip = inet_ntoa(client.sin_addr);
    	int port = ntohs(client.sin_port);
		printf("Connection accepted: %s:%d\n", *ip, port);
		
		EndPoint ep;
		minercave::Session session(ip, ep);
		
		st.handleClient(session, ep);

    } while (iResult > 0);

    if (shutdown(ClientSocket, SD_SEND) == SOCKET_ERROR) {
        printf("Shutdown failed with error: %d\n", WSAGetLastError());
        closesocket(ClientSocket);
        WSACleanup();
        return;
    }

    closesocket(ClientSocket);
    WSACleanup();
}









