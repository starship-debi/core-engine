#include <stdio.h>
#include <stdlib.h>
#include <string.h>

#define MAX_LENGTH 1024

void print_help() {
    printf("Usage: core-engine [OPTION]...\n");
    printf("Options:\n");
    printf("  -h, --help          show this help message and exit\n");
    printf("  -v, --version       show the version number and exit\n");
    printf("  -c, --config        specify the configuration file\n");
}

void print_version() {
    printf("Core Engine v1.0.0\n");
}

int main(int argc, char *argv[]) {
    int opt;
    char *config_file = NULL;

    while ((opt = getopt(argc, argv, "hvc:"))!= -1) {
        switch (opt) {
            case 'h':
                print_help();
                return 0;
            case 'v':
                print_version();
                return 0;
            case 'c':
                config_file = optarg;
                break;
            default:
                fprintf(stderr, "Unknown option: %c\n", optopt);
                print_help();
                return 1;
        }
    }

    if (config_file == NULL) {
        printf("Error: configuration file not specified\n");
        print_help();
        return 1;
    }

    FILE *file = fopen(config_file, "r");
    if (file == NULL) {
        fprintf(stderr, "Error opening configuration file: %s\n", config_file);
        return 1;
    }

    char buffer[MAX_LENGTH];
    while (fgets(buffer, MAX_LENGTH, file)!= NULL) {
        printf("%s", buffer);
    }

    fclose(file);
    return 0;
}