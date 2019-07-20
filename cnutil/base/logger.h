#ifndef LOG_H
#define LOG_H

#define RED_BOLD(x)     "\x1B[1;31m" x "\x1B[0m"
#define RED(x)          "\x1B[0;31m" x "\x1B[0m"
#define GREEN_BOLD(x)   "\x1B[1;32m" x "\x1B[0m"
#define GREEN(x)        "\x1B[0;32m" x "\x1B[0m"
#define YELLOW(x)       "\x1B[0;33m" x "\x1B[0m"
#define YELLOW_BOLD(x)  "\x1B[1;33m" x "\x1B[0m"
#define MAGENTA_BOLD(x) "\x1B[1;35m" x "\x1B[0m"
#define MAGENTA(x)      "\x1B[0;35m" x "\x1B[0m"
#define CYAN_BOLD(x)    "\x1B[1;36m" x "\x1B[0m"
#define CYAN(x)         "\x1B[0;36m" x "\x1B[0m"
#define WHITE_BOLD(x)   "\x1B[1;37m" x "\x1B[0m"
#define WHITE(x)        "\x1B[0;37m" x "\x1B[0m"
#define GRAY(x)         "\x1B[1;30m" x "\x1B[0m"


class Logger {
	public:
		void print();
		void print_cpuinfo();
		void print_memoryinfo();
		void print_threadusage();
};


#define LOG_WARN()		Logger::print()
#define LOG_ERR()		Logger::print()
#define LOG_INFO()		Logger::print()
#define LOG_NOTICE()	Logger::print()

#endif
