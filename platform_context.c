#include "api.h"

extern void cl12GoContextErrorCallback(char *, uint8_t *, size_t, intptr_t);

CL_CALLBACK void cl12CContextErrorCallback(char const *errorInfo,
    void const *privateInfoPtr, size_t privateInfoLen,
    void *userData)
{
    cl12GoContextErrorCallback((char *)(errorInfo), (uint8_t *)(privateInfoPtr), (size_t)(privateInfoLen), (intptr_t)(userData));
}

cl_context cl12CreateContext(cl_context_properties *properties,
    cl_uint numDevices, cl_device_id *devices,
    void *notify, intptr_t userData,
    cl_int *errcodeReturn)
{
    return clCreateContext(properties, numDevices, devices,
        (void (CL_CALLBACK *)(char const *, void const *, size_t, void *))(notify), (void *)(userData),
        errcodeReturn);
}

cl_context cl12CreateContextFromType(cl_context_properties *properties,
    cl_device_type deviceType,
    void *notify, intptr_t userData,
    cl_int *errcodeReturn)
{
    return clCreateContextFromType(properties, deviceType,
        (void (CL_CALLBACK *)(char const *, void const *, size_t, void *))(notify), (void *)(userData),
        errcodeReturn);
}

