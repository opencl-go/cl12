#include "api.h"

extern void cl12GoEventCallback(cl_event, cl_int, void*);

static CL_CALLBACK void cl12CEventCallback(cl_event event, cl_int commandStatus, void *userData)
{
    cl12GoEventCallback(event, commandStatus, (uintptr_t *)(userData));
}

cl_int cl12SetEventCallback(cl_event event, cl_int callbackType, uintptr_t *userData)
{
    return clSetEventCallback(event, callbackType, cl12CEventCallback, userData);
}
