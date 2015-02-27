#include "out.h"
#include "in.h"

int main(int argc, char **argvs){
    out("this is a test program");
    char c = in();
    out(&c);
    return 0;
}
