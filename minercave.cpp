#include <iostream>
#include <fstream>
#include "base/app.h"
#include "net/stratum.h"
#include "net/pool.h"
#define JSON "./config.json"


void loadConfig();
void execStratum();


int main(int argc, char** argv) {
	App app;
	
	app.init();
	loadConfig();
	execStratum();
	
	return 0;
}


void execStratum() {
	Config config;
	if (config.threads > 0) {
		std::cout<<"Running with "<<config.threads<<" threads"<<std::endl;
	}
	else{
		std::cout<<"Running with default "<<config.threads<<" threads"<<std::endl;;
	}
	
	StratumServer ss;
	ss.newStratum(&config).listen();
}


void loadConfig() {
	char buff[409];
	std::ifstream fp;
	
	fp.open(JSON);
	fp >> buff;
	fp.close();
}
