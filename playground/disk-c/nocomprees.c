/*
* Read no compressed data block generated from go
* mmap code based on http://beej.us/guide/bgipc/output/html/multipage/mmap.html
* 
* And there are arguments about mmap and read http://stackoverflow.com/questions/45972/mmap-vs-reading-blocks
* I guess it might perform better when you mmap file in many processes
*/
#include <stdio.h>
#include <stdlib.h>
#include <fcntl.h>
#include <unistd.h>
#include <sys/types.h>
#include <sys/mman.h>
#include <sys/stat.h>
#include <errno.h>

int main(int argc, char *argv[])
{
    int fd, offset;
    char *data;
    struct stat sbuf;

    char *file = "../fixture/xephon-no-compress";

    if ((fd = open(file, O_RDONLY)) == -1)
    {
        perror("open");
        exit(1);
    }

    if (stat(file, &sbuf) == -1)
    {
        perror("stat");
        exit(1);
    }

    // http://stackoverflow.com/questions/38561/what-is-the-argument-for-printf-that-formats-a-long
    printf("file size %ld\n", sbuf.st_size);

    data = mmap((caddr_t)0, sbuf.st_size, PROT_READ, MAP_SHARED, fd, 0);

    if (data == (caddr_t)(-1))
    {
        perror("mmap");
        exit(1);
    }

    // it seems I don't need to care about byte order for single byte (since there is no order problem for single byte?)
    printf("version is %d\n", data[0]);
    printf("time compression is %d\n", data[1]);
    printf("value compression is %d\n", data[2]);

    munmap(data, sbuf.st_size);
    close(fd);

    return 0;
}