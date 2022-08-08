#include "api.h"

cl_int cl12ExtTerminateContextKHR(void *fn, cl_context context)
{
    return ((clTerminateContextKHR_fn)(fn))(context);
}
