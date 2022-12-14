package cl12

// #cgo !darwin LDFLAGS: -lOpenCL
// #cgo darwin LDFLAGS: -framework OpenCL
// #include "api.h"
import "C"
import (
	"fmt"
	"unsafe"
)

// PlatformID references one of the available OpenCL platforms of the system.
// It allows applications to query OpenCL devices, device configuration information, and to create OpenCL contexts
// using one or more devices.
// Retrieve a list of available platforms with the function PlatformIDs().
type PlatformID uintptr

func (id PlatformID) handle() C.cl_platform_id {
	return *(*C.cl_platform_id)(unsafe.Pointer(&id))
}

// String provides a readable presentation of the platform identifier.
// It is based on the numerical value of the underlying pointer.
func (id PlatformID) String() string {
	return fmt.Sprintf("0x%X", uintptr(id))
}

// PlatformIDs returns the list of available platforms on the system.
//
// See also: https://registry.khronos.org/OpenCL/sdk/1.2/docs/man/xhtml/clGetPlatformIDs.html
func PlatformIDs() ([]PlatformID, error) {
	count := C.cl_uint(0)
	status := C.clGetPlatformIDs(0, nil, &count)
	if status != C.CL_SUCCESS {
		return nil, StatusError(status)
	}
	if count == 0 {
		return nil, nil
	}
	ids := make([]PlatformID, count)
	status = C.clGetPlatformIDs(count, (*C.cl_platform_id)(unsafe.Pointer(&ids[0])), &count)
	if status != C.CL_SUCCESS {
		return nil, StatusError(status)
	}
	return ids[:count], nil
}

// PlatformInfoName identifies properties of a platform, which can be queried with PlatformInfo().
type PlatformInfoName C.cl_platform_info

const (
	// PlatformNameInfo refers to a human-readable string that identifies the platform.
	//
	// Returned type: string
	PlatformNameInfo PlatformInfoName = C.CL_PLATFORM_NAME
	// PlatformVendorInfo refers to a human-readable string that identifies the vendor of the platform.
	//
	// Returned type: string
	PlatformVendorInfo PlatformInfoName = C.CL_PLATFORM_VENDOR
	// PlatformProfileInfo refers to the profile name supported by the implementation.
	// The profile name returned can be one of the following strings:
	//
	// "FULL_PROFILE" - if the implementation supports the OpenCL specification (functionality defined as part of the
	// core specification and does not require any extensions to be supported).
	//
	// "EMBEDDED_PROFILE" - if the implementation supports the OpenCL embedded profile. The embedded profile is defined
	// to be a subset for each version of OpenCL.
	//
	// Returned type: string
	PlatformProfileInfo PlatformInfoName = C.CL_PLATFORM_PROFILE
	// PlatformVersionInfo refers to the OpenCL version supported by the implementation.
	// This version string has the following format:
	//
	// OpenCL<space><major_version.minor_version><space><platform-specific information>
	//
	// Returned type: string
	PlatformVersionInfo PlatformInfoName = C.CL_PLATFORM_VERSION
	// PlatformExtensionsInfo refers to a space separated list of extension names (the extension names themselves do not
	// contain any spaces) supported by the platform. Each extension that is supported by all devices associated with
	// this platform must be reported here.
	//
	// Returned type: string
	PlatformExtensionsInfo PlatformInfoName = C.CL_PLATFORM_EXTENSIONS
)

// PlatformInfo queries information about an OpenCL platform.
//
// The provided size need to specify the size of the available space pointed to the provided value in bytes.
//
// The returned number is the required size, in bytes, for the queried information.
// Call the function with a zero size and nil value to request the required size. This helps in determining
// the necessary space for dynamic information, such as arrays.
//
// Raw strings are with a terminating NUL character. For convenience, use PlatformInfoString().
//
// See also: https://registry.khronos.org/OpenCL/sdk/1.2/docs/man/xhtml/clGetPlatformInfo.html
func PlatformInfo(id PlatformID, paramName PlatformInfoName, paramSize uintptr, paramValue unsafe.Pointer) (uintptr, error) {
	sizeReturn := C.size_t(0)
	status := C.clGetPlatformInfo(
		id.handle(),
		C.cl_platform_info(paramName),
		C.size_t(paramSize),
		paramValue,
		&sizeReturn)
	if status != C.CL_SUCCESS {
		return 0, StatusError(status)
	}
	return uintptr(sizeReturn), nil
}

// PlatformInfoString is a convenience method for PlatformInfo() to query information values that are string-based.
//
// This function does not verify the queried information is indeed of type string. It assumes the information is
// a NUL terminated raw string and will extract the bytes as characters before that.
func PlatformInfoString(id PlatformID, paramName PlatformInfoName) (string, error) {
	return queryString(func(paramSize uintptr, paramValue unsafe.Pointer) (uintptr, error) {
		return PlatformInfo(id, paramName, paramSize, paramValue)
	})
}

// ExtensionFunctionAddressForPlatform returns the address of the extension function named by functionName
// for a given platform.
//
// The pointer returned should be cast to a C-function pointer type matching the extension function's definition
// defined in the appropriate extension specification and header file.
//
// A return value of nil indicates that the specified function does not exist for the implementation or
// platform is not a valid platform.
// A non-nil return value for ExtensionFunctionAddressForPlatform() does not guarantee that an extension function
// is actually supported by the platform. The application must also make a corresponding query using
// PlatformInfo(platform, PlatformExtensionsInfo, ...) or DeviceInfo(device, DeviceExtensionsInfo, ...) to determine
// if an extension is supported by the OpenCL implementation.
//
// See also: https://registry.khronos.org/OpenCL/sdk/1.2/docs/man/xhtml/clGetExtensionFunctionAddressForPlatform.html
func ExtensionFunctionAddressForPlatform(id PlatformID, functionName string) unsafe.Pointer {
	rawName := C.CString(functionName)
	defer C.free(unsafe.Pointer(rawName))
	return C.clGetExtensionFunctionAddressForPlatform(id.handle(), rawName)
}

// UnloadPlatformCompiler allows the implementation to release the resources allocated by the OpenCL compiler for
// a platform.
//
// This function allows the implementation to release the resources allocated by the OpenCL compiler for platform.
// This is a hint from the application and does not guarantee that the compiler will not be used in the future or
// that the compiler will actually be unloaded by the implementation.
// Calls to BuildProgram(), CompileProgram(), or LinkProgram() after UnloadPlatformCompiler() will reload the compiler,
// if necessary, to build the appropriate program executable.
//
// Since: 1.2
// See also: https://registry.khronos.org/OpenCL/sdk/1.2/docs/man/xhtml/clUnloadPlatformCompiler.html
func UnloadPlatformCompiler(id PlatformID) error {
	status := C.clUnloadPlatformCompiler(id.handle())
	if status != C.CL_SUCCESS {
		return StatusError(status)
	}
	return nil
}
