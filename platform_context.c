#include "api.h"

extern void cl12GoContextErrorCallback(char *, uint8_t *, size_t, uintptr_t *);

static CL_CALLBACK void cl12CContextErrorCallback(char const *errorInfo,
    void const *privateInfoPtr, size_t privateInfoLen,
    void *userData)
{
    cl12GoContextErrorCallback((char *)(errorInfo), (uint8_t *)(privateInfoPtr), privateInfoLen, (uintptr_t *)(userData));
}

cl_context cl12CreateContext(cl_context_properties *properties,
    cl_uint numDevices, cl_device_id *devices,
    uintptr_t *userData,
    cl_int *errcodeReturn)
{
    return clCreateContext(properties, numDevices, devices,
        (userData != NULL) ? cl12CContextErrorCallback : NULL, userData,
        errcodeReturn);
}

cl_context cl12CreateContextFromType(cl_context_properties *properties,
    cl_device_type deviceType,
    uintptr_t *userData,
    cl_int *errcodeReturn)
{
    return clCreateContextFromType(properties, deviceType,
        (userData != NULL) ? cl12CContextErrorCallback : NULL, userData,
        errcodeReturn);
}
