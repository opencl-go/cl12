#include "api.h"

cl_int cl12ExtTerminateContextKHR(void *fn, cl_context context)
{
#ifdef cl_khr_terminate_context
    return ((clTerminateContextKHR_fn)(fn))(context);
#else
    return CL_INVALID_PLATFORM;
#endif
}
