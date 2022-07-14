#include <stdio.h>
#include <stdlib.h>
#include <unistd.h>
#include <pthread.h>


void *forking_thread(void* arg)
{
    int pid = fork();
    if (pid == 0) execl("/bin/echo", NULL);
    return NULL;
}

int main()
{
    pthread_t thread_id;
    pthread_create(&thread_id, NULL, forking_thread, NULL);
    usleep(20 * 1000);
    exit(0);
}
