#include <avm.hpp>

uint64 avm_main() {
    if(1) {
        avm_log("if");
    }

    if(1) {
        avm_log("if");
    } else if (2) {
        avm_log("else if");
    }

    if(1) {
        avm_log("if");
    } else {
        avm_log("else");
    }
    
    if(1) {
        avm_log("if");
    } else if (2) {
        avm_log("else if");
    } else {
        avm_log("else");
    }
}