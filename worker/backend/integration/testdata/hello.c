// hello - a minimal "hello" that just writes `hello` to stdout.
//
// usage:
//   
//   gcc -static ./hello.c -o hello
//   ./hello
//
// 
// ps.: this is intended to be built as a statically linked binary in order to
//      not need to load anything else.
//

#include <stdio.h>

int
main(int argc, char **argv)
{
	printf("hello\n");
}
