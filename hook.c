#include <dlfcn.h>
#include <stddef.h>
#include <stdio.h>
#include <time.h>
#include "_cgo_export.h"

#define RTLD_NEXT	((void *) -1l)

#define HOOK_SYS_FUNC(name) if( !orig_##name##_func ) { orig_##name##_func = (name##_pfn_t)dlsym(RTLD_NEXT,#name); }

typedef int (*clock_gettime_pfn_t)(clockid_t clk_id, struct timespec *tp);
static clock_gettime_pfn_t orig_clock_gettime_func;

int clock_gettime(clockid_t clk_id, struct timespec *tp) {
    HOOK_SYS_FUNC( clock_gettime );
    int rval = orig_clock_gettime_func(clk_id, tp);
    on_clock_gettime(clk_id, tp);
    return rval;
}