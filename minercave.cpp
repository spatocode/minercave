#include <iostream>
#include <fstream>
#include "base/app.h"
#include "net/stratum.h"
#include "net/pool.h"


void execStratum();


int main(int argc, char** argv) {	
	App app;
	
	app.init();
	execStratum();
	
	return 0;
}


void execStratum() {
	Pool::Config config;
	if (config.threads > 0) {
		std::cout<<"Running with "<<config.threads<<" threads"<<std::endl;
	}
	else{
		std::cout<<"Running with default "<<config.threads<<" threads"<<std::endl;;
	}
	
	minercave::Stratum stratum(&config);
	stratum.listen();
}
