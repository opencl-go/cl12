package cl12

// #include "api.h"
import "C"
import (
	"fmt"
	"unsafe"
)

// Sampler objects describe how color information from an image is being sampled.
type Sampler uintptr

func (sampler Sampler) handle() C.cl_sampler {
	return *(*C.cl_sampler)(unsafe.Pointer(&sampler))
}

// String provides a readable presentation of the sampler identifier.
// It is based on the numerical value of the underlying pointer.
func (sampler Sampler) String() string {
	return fmt.Sprintf("0x%X", uintptr(sampler))
}

// SamplerAddressingMode specifies how out-of-range image coordinates are handled when reading from an image.
type SamplerAddressingMode C.cl_addressing_mode

const (
	// AddressNoneMode specifies that behavior is undefined for out-of-range image coordinates.
	AddressNoneMode SamplerAddressingMode = C.CL_ADDRESS_NONE
	// AddressClampToEdgeMode specifies that out-of-range image coordinates are clamped to the edge of the image.
	AddressClampToEdgeMode SamplerAddressingMode = C.CL_ADDRESS_CLAMP_TO_EDGE
	// AddressClampMode specifies that out-of-range image coordinates are assigned a border color value.
	AddressClampMode SamplerAddressingMode = C.CL_ADDRESS_CLAMP
	// AddressRepeatMode specifies that out-of-range image coordinates read from the image as-if the image data were
	// replicated in all dimensions.
	AddressRepeatMode SamplerAddressingMode = C.CL_ADDRESS_REPEAT
	// AddressMirroredRepeatMode specifies that out-of-range image coordinates read from the image as-if the image data
	// were replicated in all dimensions, mirroring the image contents at the edge of each replication.
	//
	// Since: 1.1
	AddressMirroredRepeatMode SamplerAddressingMode = C.CL_ADDRESS_MIRRORED_REPEAT
)

// SamplerFilterMode specifies the type of filter that is applied when reading an image.
type SamplerFilterMode C.cl_filter_mode

const (
	// FilterNearestMode returns the image element nearest to the image coordinate.
	FilterNearestMode SamplerFilterMode = C.CL_FILTER_NEAREST
	// FilterLinearMode returns a weighted average of the four image elements nearest to the image coordinate.
	FilterLinearMode SamplerFilterMode = C.CL_FILTER_LINEAR
)

// CreateSampler creates a sampler object.
//
// See also: https://registry.khronos.org/OpenCL/sdk/1.2/docs/man/xhtml/clCreateSampler.html
func CreateSampler(context Context, normalizedCoords bool, addressingMode SamplerAddressingMode, filterMode SamplerFilterMode) (Sampler, error) {
	var status C.cl_int
	sampler := C.clCreateSampler(
		context.handle(),
		C.cl_bool(BoolFrom(normalizedCoords)),
		C.cl_addressing_mode(addressingMode),
		C.cl_filter_mode(filterMode),
		&status)
	if status != C.CL_SUCCESS {
		return 0, StatusError(status)
	}
	return Sampler(*((*uintptr)(unsafe.Pointer(&sampler)))), nil
}

// RetainSampler increments the sampler reference count.
//
// CreateSamplerWithProperties() and CreateSampler() perform an implicit retain.
//
// See also: https://registry.khronos.org/OpenCL/sdk/1.2/docs/man/xhtml/clRetainSampler.html
func RetainSampler(sampler Sampler) error {
	status := C.clRetainSampler(sampler.handle())
	if status != C.CL_SUCCESS {
		return StatusError(status)
	}
	return nil
}

// ReleaseSampler decrements the sampler reference count.
//
// The sampler object is deleted after the reference count becomes zero and commands queued for execution on a
// command-queue(s) that use sampler have finished.
//
// See also: https://registry.khronos.org/OpenCL/sdk/1.2/docs/man/xhtml/clReleaseSampler.html
func ReleaseSampler(sampler Sampler) error {
	status := C.clReleaseSampler(sampler.handle())
	if status != C.CL_SUCCESS {
		return StatusError(status)
	}
	return nil
}

// SamplerInfoName identifies properties of a sampler, which can be queried with SamplerInfo().
type SamplerInfoName C.cl_sampler_info

const (
	// SamplerReferenceCountInfo returns the sampler reference count.
	//
	// Note: The reference count returned should be considered immediately stale. It is unsuitable for general use
	// in applications. This feature is provided for identifying memory leaks.
	//
	// Returned type: Uint
	SamplerReferenceCountInfo SamplerInfoName = C.CL_SAMPLER_REFERENCE_COUNT
	// SamplerContextInfo returns the context specified when the sampler is created.
	//
	// Returned type: Context
	SamplerContextInfo SamplerInfoName = C.CL_SAMPLER_CONTEXT
	// SamplerNormalizedCoordsInfo return the normalized coords value associated with sampler.
	//
	// Returned type: Bool
	SamplerNormalizedCoordsInfo SamplerInfoName = C.CL_SAMPLER_NORMALIZED_COORDS
	// SamplerAddressingModeInfo returns the addressing mode value associated with sampler.
	//
	// Returned type: SamplerAddressingMode
	SamplerAddressingModeInfo SamplerInfoName = C.CL_SAMPLER_ADDRESSING_MODE
	// SamplerFilterModeInfo returns the filter mode value associated with sampler.
	//
	// Returned type: SamplerFilterMode
	SamplerFilterModeInfo SamplerInfoName = C.CL_SAMPLER_FILTER_MODE
)

// SamplerInfo queries information about a sampler.
//
// The provided size need to specify the size of the available space pointed to the provided value in bytes.
//
// The returned number is the required size, in bytes, for the queried information.
// Call the function with a zero size and nil value to request the required size. This helps in determining
// the necessary space for dynamic information, such as arrays.
//
// Raw strings are with a terminating NUL character.
//
// See also: https://registry.khronos.org/OpenCL/sdk/1.2/docs/man/xhtml/clGetSamplerInfo.html
func SamplerInfo(sampler Sampler, paramName ContextInfoName, paramSize uintptr, paramValue unsafe.Pointer) (uintptr, error) {
	sizeReturn := C.size_t(0)
	status := C.clGetSamplerInfo(
		sampler.handle(),
		C.cl_sampler_info(paramName),
		C.size_t(paramSize),
		paramValue,
		&sizeReturn)
	if status != C.CL_SUCCESS {
		return 0, StatusError(status)
	}
	return uintptr(sizeReturn), nil
}
