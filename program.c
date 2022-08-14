#include "api.h"

extern void cl12GoProgramBuildCallback(cl_program, uintptr_t *);

static CL_CALLBACK void cl12CProgramBuildCallback(cl_program program, void *userData)
{
    cl12GoProgramBuildCallback(program, (uintptr_t *)(userData));
}

cl_int cl12BuildProgram(cl_program program,
    cl_uint numDevices, cl_device_id *devices,
    char *options, uintptr_t *userData)
{
    return clBuildProgram(program, numDevices, devices, options,
        (userData != NULL) ? cl12CProgramBuildCallback : NULL, userData);
}

extern void cl12GoProgramCompileCallback(cl_program, uintptr_t *);

static CL_CALLBACK void cl12CProgramCompileCallback(cl_program program, void *userData)
{
    cl12GoProgramCompileCallback(program, (uintptr_t *)(userData));
}

cl_int cl12CompileProgram(cl_program program,
    cl_uint numDevices, cl_device_id *devices,
    char *options,
    cl_uint numInputHeaders, cl_program *headers, char const **includeNames,
    uintptr_t *userData)
{
    return clCompileProgram(program, numDevices, devices, options,
        numInputHeaders, headers, includeNames,
        (userData != NULL) ? cl12CProgramCompileCallback : NULL, userData);
}

extern void cl12GoProgramLinkCallback(cl_program, uintptr_t *);

static CL_CALLBACK void cl12CProgramLinkCallback(cl_program program, void *userData)
{
    cl12GoProgramLinkCallback(program, (uintptr_t *)(userData));
}

cl_program cl12LinkProgram(cl_context context,
    cl_uint numDevices, cl_device_id *devices,
    char *options,
    cl_uint numInputPrograms, cl_program *programs,
    uintptr_t *userData,
    cl_int *errReturn)
{
    return clLinkProgram(context, numDevices, devices, options,
        numInputPrograms, programs,
        (userData != NULL) ? cl12CProgramLinkCallback : NULL, userData,
        errReturn);
}
