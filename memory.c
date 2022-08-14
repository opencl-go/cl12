#include "api.h"

extern void cl12GoMemObjectDestructorCallback(cl_mem, uintptr_t *);

static CL_CALLBACK void cl12CMemObjectDestructorCallback(cl_mem mem, void *userData)
{
    cl12GoMemObjectDestructorCallback(mem, (uintptr_t *)(userData));
}

cl_int cl12SetMemObjectDestructorCallback(cl_mem mem, uintptr_t *userData)
{
    return clSetMemObjectDestructorCallback(mem, cl12CMemObjectDestructorCallback, userData);
}
