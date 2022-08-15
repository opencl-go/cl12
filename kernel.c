#include "api.h"

extern void cl12GoKernelNativeCallback(void *);

static CL_CALLBACK void cl12CKernelNativeCallback(void *args)
{
    cl12GoKernelNativeCallback(args);
}

cl_int cl12EnqueueNativeKernel(cl_command_queue commandQueue,
    void *args, size_t argsSize,
    cl_uint numMemObjects, cl_mem *memList, void const *argsMemLoc,
    cl_uint waitListCount, cl_event const *waitList,
    cl_event *event)
{
    return clEnqueueNativeKernel(
        commandQueue,
        cl12CKernelNativeCallback, args, argsSize,
        numMemObjects, memList, (void const **)(argsMemLoc),
        waitListCount, waitList,
        event);
}
