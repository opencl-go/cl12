package cl12

// #include "api.h"
import "C"
import (
	"fmt"
	"unsafe"
)

// DeviceID references an OpenCL device of the system.
type DeviceID uintptr

func (id DeviceID) handle() C.cl_device_id {
	return *(*C.cl_device_id)(unsafe.Pointer(&id))
}

// String provides a readable presentation of the device identifier.
// It is based on the numerical value of the underlying pointer.
func (id DeviceID) String() string {
	return fmt.Sprintf("0x%X", uintptr(id))
}

// DeviceTypeFlags is a bitfield that identifies the type of OpenCL device.
// It can be used for DeviceIDs() to filter for the requested devices.
type DeviceTypeFlags C.cl_device_type

const (
	// DeviceTypeCPU is an OpenCL device similar to a traditional CPU (Central Processing Unit).
	// The host processor that executes OpenCL host code may also be considered a CPU OpenCL device.
	DeviceTypeCPU DeviceTypeFlags = C.CL_DEVICE_TYPE_CPU
	// DeviceTypeDefault is the default OpenCL device in the platform.
	// The default OpenCL device must not be a DeviceTypeCustom device.
	DeviceTypeDefault DeviceTypeFlags = C.CL_DEVICE_TYPE_DEFAULT
	// DeviceTypeGpu is an OpenCL device similar to a GPU (Graphics Processing Unit).
	// Many systems include a dedicated processor for graphics or rendering that may be considered a GPU OpenCL device.
	DeviceTypeGpu DeviceTypeFlags = C.CL_DEVICE_TYPE_GPU
	// DeviceTypeAccelerator are dedicated devices that may accelerate OpenCL programs, such as FPGAs
	// (Field Programmable Gate Arrays), DSPs (Digital Signal Processors), or AI (Artificial Intelligence) processors.
	DeviceTypeAccelerator DeviceTypeFlags = C.CL_DEVICE_TYPE_ACCELERATOR
	// DeviceTypeCustom are specialized devices that implement some OpenCL runtime APIs but do not support
	// all required OpenCL functionality.
	//
	// Since: 1.2
	DeviceTypeCustom DeviceTypeFlags = C.CL_DEVICE_TYPE_CUSTOM
	// DeviceTypeAll identifies all OpenCL devices available in the platform, except for DeviceTypeCustom devices.
	DeviceTypeAll DeviceTypeFlags = C.CL_DEVICE_TYPE_ALL
)

// DeviceIDs queries devices available on a platform.
//
// The deviceType is a bitfield that identifies the type of OpenCL device. The deviceType can be used to query
// specific OpenCL devices or all OpenCL devices available.
//
// See also: https://registry.khronos.org/OpenCL/sdk/1.2/docs/man/xhtml/clGetDeviceIDs.html
func DeviceIDs(platformID PlatformID, deviceType DeviceTypeFlags) ([]DeviceID, error) {
	count := C.cl_uint(0)
	status := C.clGetDeviceIDs(platformID.handle(), C.cl_device_type(deviceType), 0, nil, &count)
	if status != C.CL_SUCCESS {
		return nil, StatusError(status)
	}
	if count == 0 {
		return nil, nil
	}
	ids := make([]DeviceID, count)
	status = C.clGetDeviceIDs(platformID.handle(), C.cl_device_type(deviceType), count, (*C.cl_device_id)(unsafe.Pointer(&ids[0])), &count)
	if status != C.CL_SUCCESS {
		return nil, StatusError(status)
	}
	return ids[:count], nil
}

// DeviceInfoName identifies properties of a device, which can be queried with DeviceInfo().
type DeviceInfoName C.cl_device_info

const (
	// DeviceAddressBitsInfo specify the default compute device address space size of the global address space specified
	// as an unsigned integer value in bits. Currently supported values are 32 or 64 bits.
	//
	// Returned type: uint
	DeviceAddressBitsInfo DeviceInfoName = C.CL_DEVICE_ADDRESS_BITS
	// DeviceAvailableInfo is true if the device is available and false otherwise. A device is considered to be
	// available if the device can be expected to successfully execute commands enqueued to the device.
	//
	// Returned type: Bool
	DeviceAvailableInfo DeviceInfoName = C.CL_DEVICE_AVAILABLE
	// DeviceBuiltInKernelsInfo represents a semicolon separated list of built-in kernels supported by the device.
	// An empty string is returned if no built-in kernels are supported by the device.
	//
	// Returned type: string
	// Since: 1.2
	DeviceBuiltInKernelsInfo DeviceInfoName = C.CL_DEVICE_BUILT_IN_KERNELS
	// DeviceCompilerAvailableInfo is False if the implementation does not have a compiler available to compile the
	// program source. It is True if the compiler is available.
	//
	// This can be False for the embedded platform profile only.
	//
	// Returned type: Bool
	DeviceCompilerAvailableInfo DeviceInfoName = C.CL_DEVICE_COMPILER_AVAILABLE
	// DeviceDoubleFpConfigInfo describes double precision floating-point capability of the OpenCL device.
	// This is a bit-field.
	//
	// Returned type: DeviceFpConfigFlags
	// Since: 1.2
	DeviceDoubleFpConfigInfo DeviceInfoName = C.CL_DEVICE_DOUBLE_FP_CONFIG
	// DeviceEndianLittleInfo is True if the OpenCL device is a little endian device and False otherwise.
	//
	// Returned type: Bool
	DeviceEndianLittleInfo DeviceInfoName = C.CL_DEVICE_ENDIAN_LITTLE
	// DeviceErrorCorrectionSupportInfo is True if the device implements error correction for all accesses to compute
	// device memory (global and constant). Is False if the device does not implement such error correction.
	//
	// Returned type: Bool
	DeviceErrorCorrectionSupportInfo DeviceInfoName = C.CL_DEVICE_ERROR_CORRECTION_SUPPORT
	// DeviceExecutionCapabilitiesInfo describes the execution capabilities of the device. This is a bit-field.
	//
	// The mandated minimum capability is ExecKernel.
	//
	// Returned type: DeviceExecCapabilitiesFlags
	DeviceExecutionCapabilitiesInfo DeviceInfoName = C.CL_DEVICE_EXECUTION_CAPABILITIES
	// DeviceExtensionsInfo returns a space separated list of extension names (the extension names themselves do not
	// contain any spaces) supported by the device. The list of extension names may include Khronos approved
	// extension names and vendor specified extension names.
	//
	// Returned type: string
	DeviceExtensionsInfo DeviceInfoName = C.CL_DEVICE_EXTENSIONS
	// DeviceGlobalMemCacheSizeInfo returns the size of global memory cache in bytes.
	//
	// Returned type: uint64
	DeviceGlobalMemCacheSizeInfo DeviceInfoName = C.CL_DEVICE_GLOBAL_MEM_CACHE_SIZE
	// DeviceGlobalMemCacheTypeInfo represents the type of global memory cache supported.
	//
	// Returned type: DeviceMemCacheTypeEnum
	DeviceGlobalMemCacheTypeInfo DeviceInfoName = C.CL_DEVICE_GLOBAL_MEM_CACHE_TYPE
	// DeviceGlobalMemCachelineSizeInfo is the size of global memory cache line in bytes.
	//
	// Returned type: Uint
	DeviceGlobalMemCachelineSizeInfo DeviceInfoName = C.CL_DEVICE_GLOBAL_MEM_CACHELINE_SIZE
	// DeviceGlobalMemSizeInfo is the size of global device memory in bytes.
	//
	// Returned type: uint64
	DeviceGlobalMemSizeInfo DeviceInfoName = C.CL_DEVICE_GLOBAL_MEM_SIZE
	// DeviceHostUnifiedMemoryInfo is True if the device and the host have a unified memory subsystem and is
	// False otherwise.
	//
	// Returned type: Bool
	DeviceHostUnifiedMemoryInfo DeviceInfoName = C.CL_DEVICE_HOST_UNIFIED_MEMORY
	// DeviceImage2dMaxHeightInfo is the maximum height of 2D image in pixels.
	// The minimum value is 16384 if DeviceImageSupportInfo is True, the value is 0 otherwise.
	//
	// Returned type: uintptr
	DeviceImage2dMaxHeightInfo DeviceInfoName = C.CL_DEVICE_IMAGE2D_MAX_HEIGHT
	// DeviceImage2dMaxWidthInfo is the maximum width of 2D image or 1D image not created from a buffer object in pixels.
	// The minimum value is 16384 if DeviceImageSupportInfo is True, the value is 0 otherwise.
	//
	// Returned type: uintptr
	DeviceImage2dMaxWidthInfo DeviceInfoName = C.CL_DEVICE_IMAGE2D_MAX_WIDTH
	// DeviceImage3dMaxDepthInfo is the maximum depth of 3D image in pixels.
	// The minimum value is 2048 if DeviceImageSupportInfo is True, the value is 0 otherwise.
	//
	// Returned type: uintptr
	DeviceImage3dMaxDepthInfo DeviceInfoName = C.CL_DEVICE_IMAGE3D_MAX_DEPTH
	// DeviceImage3dMaxHeightInfo is the maximum height of 3D image in pixels.
	// The minimum value is 2048 if DeviceImageSupportInfo is True, the value is 0 otherwise.
	//
	// Returned type: uintptr
	DeviceImage3dMaxHeightInfo DeviceInfoName = C.CL_DEVICE_IMAGE3D_MAX_HEIGHT
	// DeviceImage3dMaxWidthInfo is the maximum width of 3D image in pixels.
	// The minimum value is 2048 if DeviceImageSupportInfo is True, the value is 0 otherwise.
	//
	// Returned type: uintptr
	DeviceImage3dMaxWidthInfo DeviceInfoName = C.CL_DEVICE_IMAGE3D_MAX_WIDTH
	// DeviceImageMaxArraySizeInfo is the maximum number of images in a 1D or 2D image array.
	// The minimum value is 2048 if DeviceImageSupportInfo is True, the value is 0 otherwise.
	//
	// Returned type: uintptr
	// Since: 1.2
	DeviceImageMaxArraySizeInfo DeviceInfoName = C.CL_DEVICE_IMAGE_MAX_ARRAY_SIZE
	// DeviceImageMaxBufferSizeInfo is the maximum number of pixels for a 1D image created from a buffer object.
	// The minimum value is 65536 if DeviceImageSupportInfo is True, the value is 0 otherwise.
	//
	// Returned type: uintptr
	// Since: 1.2
	DeviceImageMaxBufferSizeInfo DeviceInfoName = C.CL_DEVICE_IMAGE_MAX_BUFFER_SIZE
	// DeviceImageSupportInfo is True if images are supported by the OpenCL device and False otherwise.
	//
	// Returned type: Bool
	DeviceImageSupportInfo DeviceInfoName = C.CL_DEVICE_IMAGE_SUPPORT
	// DeviceLinkerAvailableInfo is False if the implementation does not have a linker available.
	// It is True if the linker is available.
	//
	// This can be False for the embedded platform profile only.
	// This must be True if DeviceCompilerAvailableInfo is True.
	//
	// Returned type: Bool
	// Since: 1.2
	DeviceLinkerAvailableInfo DeviceInfoName = C.CL_DEVICE_LINKER_AVAILABLE
	// DeviceLocalMemSizeInfo is the size of local memory region in bytes. The minimum value is 32 KB for devices
	// that are not of type DeviceTypeCustom.
	//
	// Returned type: uint64
	DeviceLocalMemSizeInfo DeviceInfoName = C.CL_DEVICE_LOCAL_MEM_SIZE
	// DeviceLocalMemTypeInfo is the type of local memory supported.
	// This can be set to DeviceLocalMemTypeLocal implying dedicated local memory storage such as SRAM, or
	// DeviceLocalMemTypeGlobal.
	//
	// For custom devices, DeviceLocalMemTypeNone can also be returned indicating no local memory support.
	//
	// Returned type: DeviceLocalMemTypeEnum
	DeviceLocalMemTypeInfo DeviceInfoName = C.CL_DEVICE_LOCAL_MEM_TYPE
	// DeviceMaxClockFrequencyInfo is the clock frequency of the device in MHz. The meaning of this value is
	// implementation-defined. For devices with multiple clock domains, the clock frequency for any of the clock
	// domains may be returned. For devices that dynamically change frequency for power or thermal reasons, the
	// returned clock frequency may be any valid frequency.
	//
	// Note: This definition is missing before version 2.2.
	//
	// Returned type: Uint
	DeviceMaxClockFrequencyInfo DeviceInfoName = C.CL_DEVICE_MAX_CLOCK_FREQUENCY
	// DeviceMaxComputeUnitsInfo refers to the number of parallel compute units on the OpenCL device.
	// A work-group executes on a single compute unit. The minimum value is 1.
	//
	// Returned type: Uint
	DeviceMaxComputeUnitsInfo DeviceInfoName = C.CL_DEVICE_MAX_COMPUTE_UNITS
	// DeviceMaxConstantArgsInfo is the maximum number of arguments declared with the __constant qualifier in a kernel.
	// The minimum value is 8 for devices that are not of type DeviceTypeCustom.
	//
	// Returned type: Uint
	DeviceMaxConstantArgsInfo DeviceInfoName = C.CL_DEVICE_MAX_CONSTANT_ARGS
	// DeviceMaxConstantBufferSizeInfo is the maximum size in bytes of a constant buffer allocation. The minimum value
	// is 64 KB for devices that are not of type DeviceTypeCustom.
	//
	// Returned type: uint64
	DeviceMaxConstantBufferSizeInfo DeviceInfoName = C.CL_DEVICE_MAX_CONSTANT_BUFFER_SIZE
	// DeviceMaxMemAllocSizeInfo is the maximum size of memory object allocation in bytes. The minimum value is
	// max(min(1024 * 1024 * 1024, 1/4th of DeviceGlobalMemSizeInfo), 32 * 1024 * 1024)
	// for devices that are not of type DeviceTypeCustom.
	//
	// Returned type: uint64
	DeviceMaxMemAllocSizeInfo DeviceInfoName = C.CL_DEVICE_MAX_MEM_ALLOC_SIZE
	// DeviceMaxParameterSizeInfo is the maximum size in bytes of all arguments that can be passed to a kernel.
	//
	// The minimum value is 1024 for devices that are not of type DeviceTypeCustom. For this minimum value,
	// only a maximum of 128 arguments can be passed to a kernel
	//
	// Returned type: uintptr
	DeviceMaxParameterSizeInfo DeviceInfoName = C.CL_DEVICE_MAX_PARAMETER_SIZE
	// DeviceMaxReadImageArgsInfo is the maximum number of image objects arguments of a kernel declared with the
	// read_only qualifier. The minimum value is 128 if DeviceImageSupportInfo is True, the value is 0 otherwise.
	//
	// Returned type: Uint
	DeviceMaxReadImageArgsInfo DeviceInfoName = C.CL_DEVICE_MAX_READ_IMAGE_ARGS
	// DeviceMaxSamplersInfo is the maximum number of samplers that can be used in a kernel.
	// The minimum value is 16 if DeviceImageSupportInfo is True, the value is 0 otherwise.
	//
	// Returned type: Uint
	DeviceMaxSamplersInfo DeviceInfoName = C.CL_DEVICE_MAX_SAMPLERS
	// DeviceMaxWorkGroupSizeInfo is the maximum number of work-items in a work-group that a device is capable of
	// executing on a single compute unit, for any given kernel-instance running on the device.
	// The minimum value is 1. The returned value is an upper limit and will not necessarily maximize performance.
	// This maximum may be larger than supported by a specific kernel.
	//
	// Returned type: uintptr
	DeviceMaxWorkGroupSizeInfo DeviceInfoName = C.CL_DEVICE_MAX_WORK_GROUP_SIZE
	// DeviceMaxWorkItemDimensionsInfo is the maximum dimensions that specify the global and local work-item IDs used by
	// the data parallel execution model. The minimum value is 3 for devices that are not of type DeviceTypeCustom.
	//
	// Return type: Uint
	DeviceMaxWorkItemDimensionsInfo DeviceInfoName = C.CL_DEVICE_MAX_WORK_ITEM_DIMENSIONS
	// DeviceMaxWorkItemSizesInfo is the maximum number of work-items that can be specified in each dimension of the
	// work-group to EnqueueNDRangeKernel().
	// Returns N uintptr entries, where N is the value returned by the query for DeviceMaxWorkItemDimensionsInfo.
	// The minimum value is (1, 1, 1) for devices that are not of type DeviceTypeCustom.
	//
	// Returned type: []uintptr
	DeviceMaxWorkItemSizesInfo DeviceInfoName = C.CL_DEVICE_MAX_WORK_ITEM_SIZES
	// DeviceMaxWriteImageArgsInfo is the maximum number of image objects arguments of a kernel declared with the
	// write_only qualifier. The minimum value is 64 if DeviceImageSupportInfo is True, the value is 0 otherwise.
	DeviceMaxWriteImageArgsInfo DeviceInfoName = C.CL_DEVICE_MAX_WRITE_IMAGE_ARGS
	// DeviceMemBaseAddrAlignInfo is the alignment requirement (in bits) for sub-buffer offsets. The minimum value is
	// the size (in bits) of the largest OpenCL built-in data type supported by the device
	// (long16 in FULL profile, long16 or int16 in EMBEDDED profile) for devices that are not of type DeviceTypeCustom.
	//
	// Returned type: Uint
	DeviceMemBaseAddrAlignInfo DeviceInfoName = C.CL_DEVICE_MEM_BASE_ADDR_ALIGN
	// DeviceNameInfo refers to a human-readable string that identifies the device.
	//
	// Returned type: string
	DeviceNameInfo DeviceInfoName = C.CL_DEVICE_NAME
	// DeviceNativeVectorWidthCharInfo returns the native ISA vector width.
	// The vector width is defined as the number of scalar elements that can be stored in the vector.
	//
	// Returned type: Uint
	// Since: 1.1
	DeviceNativeVectorWidthCharInfo DeviceInfoName = C.CL_DEVICE_NATIVE_VECTOR_WIDTH_CHAR
	// DeviceNativeVectorWidthDoubleInfo returns the native ISA vector width.
	// The vector width is defined as the number of scalar elements that can be stored in the vector.
	//
	// If double precision is not supported, this value must be 0.
	//
	// Returned type: Uint
	// Since: 1.1
	DeviceNativeVectorWidthDoubleInfo DeviceInfoName = C.CL_DEVICE_NATIVE_VECTOR_WIDTH_DOUBLE
	// DeviceNativeVectorWidthFloatInfo returns the native ISA vector width.
	// The vector width is defined as the number of scalar elements that can be stored in the vector.
	//
	// Returned type: Uint
	// Since: 1.1
	DeviceNativeVectorWidthFloatInfo DeviceInfoName = C.CL_DEVICE_NATIVE_VECTOR_WIDTH_FLOAT
	// DeviceNativeVectorWidthHalfInfo returns the native ISA vector width.
	// The vector width is defined as the number of scalar elements that can be stored in the vector.
	//
	// If the cl_khr_fp16 extension is not supported, this value must be 0.
	//
	// Returned type: Uint
	// Since: 1.1
	// Extension: cl_khr_fp16
	DeviceNativeVectorWidthHalfInfo DeviceInfoName = C.CL_DEVICE_NATIVE_VECTOR_WIDTH_HALF
	// DeviceNativeVectorWidthIntInfo returns the native ISA vector width.
	// The vector width is defined as the number of scalar elements that can be stored in the vector.
	//
	// Returned type: Uint
	// Since: 1.1
	DeviceNativeVectorWidthIntInfo DeviceInfoName = C.CL_DEVICE_NATIVE_VECTOR_WIDTH_INT
	// DeviceNativeVectorWidthLongInfo returns the native ISA vector width.
	// The vector width is defined as the number of scalar elements that can be stored in the vector.
	//
	// Returned type: Uint
	// Since: 1.1
	DeviceNativeVectorWidthLongInfo DeviceInfoName = C.CL_DEVICE_NATIVE_VECTOR_WIDTH_LONG
	// DeviceNativeVectorWidthShortInfo returns the native ISA vector width.
	// The vector width is defined as the number of scalar elements that can be stored in the vector.
	//
	// Returned type: Uint
	// Since: 1.1
	DeviceNativeVectorWidthShortInfo DeviceInfoName = C.CL_DEVICE_NATIVE_VECTOR_WIDTH_SHORT
	// DeviceOpenClCVersionInfo returns the highest fully backwards compatible OpenCL C version supported by the
	// compiler for the device. For devices supporting compilation from OpenCL C source, this will return
	// a version string with the following format:
	//
	// OpenCL<space>C<space><major_version.minor_version><space><vendor-specific information>
	//
	// Returned type: string
	// Since: 1.1
	DeviceOpenClCVersionInfo DeviceInfoName = C.CL_DEVICE_OPENCL_C_VERSION
	// DeviceParentDeviceInfo returns the DeviceID of the parent device to which this sub-device belongs.
	// If device is a root-level device, a zero value is returned.
	//
	// Returned type: DeviceID
	// Since: 1.2
	DeviceParentDeviceInfo DeviceInfoName = C.CL_DEVICE_PARENT_DEVICE
	// DevicePartitionAffinityDomainInfo returns the list of supported affinity domains for partitioning the device
	// using DevicePartitionByAffinityDomainProperty. This is a bit-field.
	// If the device does not support any affinity domains, a value of 0 will be returned.
	//
	// Returned type: DeviceAffinityDomainFlags
	// Since: 1.2
	DevicePartitionAffinityDomainInfo DeviceInfoName = C.CL_DEVICE_PARTITION_AFFINITY_DOMAIN
	// DevicePartitionMaxSubDevicesInfo returns the maximum number of sub-devices that can be created when
	// a device is partitioned.
	// The value returned cannot exceed DeviceMaxComputeUnitsInfo.
	//
	// Returned type: Uint
	// Since: 1.2
	DevicePartitionMaxSubDevicesInfo DeviceInfoName = C.CL_DEVICE_PARTITION_MAX_SUB_DEVICES
	// DevicePartitionPropertiesInfo returns the list of partition types supported by device.
	// This is an array of uintptr values drawn from the list of DevicePartitionEquallyProperty, DevicePartitionByCountsProperty,
	// and DevicePartitionByAffinityDomainProperty.
	// If the device cannot be partitioned (i.e. there is no partitioning scheme supported by the device that will
	// return at least two subdevices), a value of 0 will be returned.
	//
	// Returned type: []uintptr
	// Since: 1.2
	DevicePartitionPropertiesInfo DeviceInfoName = C.CL_DEVICE_PARTITION_PROPERTIES
	// DevicePartitionTypeInfo returns the properties argument specified in CreateSubDevices() if device is a sub-device.
	// In the case where the properties argument to CreateSubDevices() is DevicePartitionByAffinityDomainProperty,
	// DeviceAffinityDomainNextPartitionable, the affinity domain used to perform the partition will be returned.
	// This can be one of the following values:
	//
	// DeviceAffinityDomainNuma
	// DeviceAffinityDomainL4Cache
	// DeviceAffinityDomainL3Cache
	// DeviceAffinityDomainL2Cache
	// DeviceAffinityDomainL1Cache
	//
	// Otherwise the implementation may either return a parameter size of 0 i.e. there is no partition type
	// associated with device or can return a property value of 0 (where 0 is used to terminate the
	// partition property list) in the memory that the result value points to.
	//
	// Returned type: []uintptr
	// Since: 1.2
	DevicePartitionTypeInfo DeviceInfoName = C.CL_DEVICE_PARTITION_TYPE
	// DevicePlatformInfo returns the platform associated with this device.
	//
	// Returned type: PlatformID
	DevicePlatformInfo DeviceInfoName = C.CL_DEVICE_PLATFORM
	// DevicePreferredInteropUserSyncInfo is True if the devices preference is for the user to be responsible for
	// synchronization, when sharing memory objects between OpenCL and other APIs such as DirectX,
	// False if the device / implementation has a performant path for performing synchronization of memory object
	// shared between OpenCL and other APIs such as DirectX.
	//
	// Returned type: Bool
	// Since: 1.2
	DevicePreferredInteropUserSyncInfo DeviceInfoName = C.CL_DEVICE_PREFERRED_INTEROP_USER_SYNC
	// DevicePreferredVectorWidthCharInfo is the preferred native vector width size for built-in scalar types that
	// can be put into vectors. The vector width is defined as the number of scalar elements that can be stored
	// in the vector.
	//
	// Returned type: Uint
	DevicePreferredVectorWidthCharInfo DeviceInfoName = C.CL_DEVICE_PREFERRED_VECTOR_WIDTH_CHAR
	// DevicePreferredVectorWidthDoubleInfo is the preferred native vector width size for built-in scalar types that
	// can be put into vectors. The vector width is defined as the number of scalar elements that can be stored
	// in the vector.
	// If double precision is not supported, this value must be 0.
	//
	// Returned type: Uint
	DevicePreferredVectorWidthDoubleInfo DeviceInfoName = C.CL_DEVICE_PREFERRED_VECTOR_WIDTH_DOUBLE
	// DevicePreferredVectorWidthFloatInfo is the preferred native vector width size for built-in scalar types that
	// can be put into vectors. The vector width is defined as the number of scalar elements that can be stored
	// in the vector.
	//
	// Returned type: Uint
	DevicePreferredVectorWidthFloatInfo DeviceInfoName = C.CL_DEVICE_PREFERRED_VECTOR_WIDTH_FLOAT
	// DevicePreferredVectorWidthHalfInfo is the preferred native vector width size for built-in scalar types that
	// can be put into vectors. The vector width is defined as the number of scalar elements that can be stored
	// in the vector.
	// If the cl_khr_fp16 extension is not supported, this value must be 0.
	//
	// Returned type: Uint
	// Since: 1.1
	// Extension: cl_khr_fp16
	DevicePreferredVectorWidthHalfInfo DeviceInfoName = C.CL_DEVICE_PREFERRED_VECTOR_WIDTH_HALF
	// DevicePreferredVectorWidthIntInfo is the preferred native vector width size for built-in scalar types that
	// can be put into vectors. The vector width is defined as the number of scalar elements that can be stored
	// in the vector.
	//
	// Returned type: Uint
	DevicePreferredVectorWidthIntInfo DeviceInfoName = C.CL_DEVICE_PREFERRED_VECTOR_WIDTH_INT
	// DevicePreferredVectorWidthLongInfo is the preferred native vector width size for built-in scalar types that
	// can be put into vectors. The vector width is defined as the number of scalar elements that can be stored
	// in the vector.
	//
	// Returned type: Uint
	DevicePreferredVectorWidthLongInfo DeviceInfoName = C.CL_DEVICE_PREFERRED_VECTOR_WIDTH_LONG
	// DevicePreferredVectorWidthShortInfo is the preferred native vector width size for built-in scalar types that
	// can be put into vectors. The vector width is defined as the number of scalar elements that can be stored
	// in the vector.
	//
	// Returned type: Uint
	DevicePreferredVectorWidthShortInfo DeviceInfoName = C.CL_DEVICE_PREFERRED_VECTOR_WIDTH_SHORT
	// DevicePrintfBufferSizeInfo is the maximum size in bytes of the internal buffer that holds the output of printf
	// calls from a kernel. The minimum value for the FULL profile is 1 MB.
	//
	// Returned type: uintptr
	// Since: 1.2
	DevicePrintfBufferSizeInfo DeviceInfoName = C.CL_DEVICE_PRINTF_BUFFER_SIZE
	// DeviceProfileInfo is the OpenCL profile string. Returns the profile name supported by the device.
	// The profile name returned can be one of the following strings:
	//
	// "FULL_PROFILE" - if the device supports the OpenCL specification (functionality defined as part of the core
	// specification and does not require any extensions to be supported).
	//
	// "EMBEDDED_PROFILE" - if the device supports the OpenCL embedded profile.
	//
	// Returned type: string
	DeviceProfileInfo DeviceInfoName = C.CL_DEVICE_PROFILE
	// DeviceProfilingTimerResolutionInfo describes the resolution of device timer. This is measured in nanoseconds.
	//
	// Returned type: uintptr
	DeviceProfilingTimerResolutionInfo DeviceInfoName = C.CL_DEVICE_PROFILING_TIMER_RESOLUTION
	// DeviceQueuePropertiesInfo describes the command-queue properties supported by the device. This is a bit-field
	// that describes one or more of the following values: QueueOutOfOrderExecModeEnable, QueueProfilingEnable.
	//
	// Returned type: CommandQueuePropertiesFlags
	DeviceQueuePropertiesInfo DeviceInfoName = C.CL_DEVICE_QUEUE_PROPERTIES
	// DeviceReferenceCountInfo returns the device reference count. If the device is a root-level device,
	// a reference count of one is returned.
	//
	// Note: The reference count returned should be considered immediately stale. It is unsuitable for general
	// use in applications. This feature is provided for identifying memory leaks.
	//
	// Returned type: Uint
	// Since: 1.2
	DeviceReferenceCountInfo DeviceInfoName = C.CL_DEVICE_REFERENCE_COUNT
	// DeviceSingleFpConfigInfo describes single precision floating-point capability of the device. This is a bit-field.
	//
	// Returned type: DeviceFpConfigFlags
	DeviceSingleFpConfigInfo DeviceInfoName = C.CL_DEVICE_SINGLE_FP_CONFIG
	// DeviceTypeInfo is the type or types of the OpenCL device.
	//
	// Returned type: DeviceTypeFlags
	DeviceTypeInfo DeviceInfoName = C.CL_DEVICE_TYPE
	// DeviceVendorInfo refers to a human-readable string that identifies the vendor of the device.
	//
	// Returned type: string
	DeviceVendorInfo DeviceInfoName = C.CL_DEVICE_VENDOR
	// DeviceVendorIDInfo is a unique device vendor identifier.
	//
	// Returned type: Uint
	DeviceVendorIDInfo DeviceInfoName = C.CL_DEVICE_VENDOR_ID
	// DeviceVersionInfo is an OpenCL version string. Returns the OpenCL version supported by the device.
	// This version string has the following format:
	//
	// OpenCL<space><major_version.minor_version><space><vendor-specific information>
	//
	// Returned type: string
	DeviceVersionInfo DeviceInfoName = C.CL_DEVICE_VERSION
	// DriverVersionInfo specifies the OpenCL software driver version string. Follows a vendor-specific format.
	//
	// Returned type: string
	DriverVersionInfo DeviceInfoName = C.CL_DRIVER_VERSION
)

// DeviceFpConfigFlags are used to determine the DeviceSingleFpConfigInfo and DeviceDoubleFpConfigInfo with DeviceInfo().
type DeviceFpConfigFlags C.cl_device_fp_config

const (
	// FpDenorm identifies denorms are supported.
	FpDenorm DeviceFpConfigFlags = C.CL_FP_DENORM
	// FpInfNan identifies INF and quiet NaNs are supported.
	FpInfNan DeviceFpConfigFlags = C.CL_FP_INF_NAN
	// FpRoundToNearest identifies round to nearest even rounding mode supported.
	FpRoundToNearest DeviceFpConfigFlags = C.CL_FP_ROUND_TO_NEAREST
	// FpRoundToZero identifies round to zero rounding mode supported.
	FpRoundToZero DeviceFpConfigFlags = C.CL_FP_ROUND_TO_ZERO
	// FpRoundToInf identifies round to positive and negative infinity rounding modes supported.
	FpRoundToInf DeviceFpConfigFlags = C.CL_FP_ROUND_TO_INF
	// FpFma identifies IEEE754-2008 fused multiply-add is supported.
	FpFma DeviceFpConfigFlags = C.CL_FP_FMA
	// FpSoftFloat identifies basic floating-point operations (such as addition, subtraction, multiplication)
	// are implemented in software.
	//
	// Since: 1.1
	FpSoftFloat DeviceFpConfigFlags = C.CL_FP_SOFT_FLOAT
	// FpCorrectlyRoundedDivideSqrt identifies divide and sqrt are correctly rounded as defined by the IEEE754
	// specification.
	//
	// Since: 1.2
	FpCorrectlyRoundedDivideSqrt DeviceFpConfigFlags = C.CL_FP_CORRECTLY_ROUNDED_DIVIDE_SQRT
)

// DeviceExecCapabilitiesFlags are used to determine the DeviceExecutionCapabilitiesInfo with DeviceInfo().
type DeviceExecCapabilitiesFlags C.cl_device_exec_capabilities

const (
	// ExecKernel identifies that the OpenCL device can execute OpenCL kernels.
	ExecKernel DeviceExecCapabilitiesFlags = C.CL_EXEC_KERNEL
	// ExecNativeKernel identifies that the OpenCL device can execute native kernels.
	ExecNativeKernel DeviceExecCapabilitiesFlags = C.CL_EXEC_NATIVE_KERNEL
)

// DeviceMemCacheTypeEnum is used to determine the DeviceGlobalMemCacheTypeInfo with DeviceInfo().
type DeviceMemCacheTypeEnum C.cl_device_mem_cache_type

// These are the possible values for DeviceMemCacheTypeEnum. They are slightly reworded compared to the original
// OpenCL API to avoid potential name/type clashes in the future.
const (
	DeviceMemCacheNone      DeviceMemCacheTypeEnum = C.CL_NONE
	DeviceMemCacheReadOnly  DeviceMemCacheTypeEnum = C.CL_READ_ONLY_CACHE
	DeviceMemCacheReadWrite DeviceMemCacheTypeEnum = C.CL_READ_WRITE_CACHE
)

// DeviceLocalMemTypeEnum is used to determine the DeviceLocalMemTypeInfo with DeviceInfo().
type DeviceLocalMemTypeEnum C.cl_device_local_mem_type

// These are the possible values for DeviceLocalMemTypeEnum. They are slightly reworded compared to the original
// OpenCL API to avoid potential name/type clashes in the future.
const (
	DeviceLocalMemTypeNone   DeviceLocalMemTypeEnum = C.CL_NONE
	DeviceLocalMemTypeLocal  DeviceLocalMemTypeEnum = C.CL_LOCAL
	DeviceLocalMemTypeGlobal DeviceLocalMemTypeEnum = C.CL_GLOBAL
)

// DeviceInfo queries specific information about a device.
//
// The provided size need to specify the size of the available space pointed to the provided value in bytes.
//
// The returned number is the required size, in bytes, for the queried information.
// Call the function with a zero size and nil value to request the required size. This helps in determining
// the necessary space for dynamic information, such as arrays.
//
// Raw strings are with a terminating NUL character. For convenience, use DeviceInfoString().
//
// See also: https://registry.khronos.org/OpenCL/sdk/1.2/docs/man/xhtml/clGetDeviceInfo.html
func DeviceInfo(id DeviceID, paramName DeviceInfoName, paramSize uintptr, paramValue unsafe.Pointer) (uintptr, error) {
	sizeReturn := C.size_t(0)
	status := C.clGetDeviceInfo(
		id.handle(),
		C.cl_device_info(paramName),
		C.size_t(paramSize),
		paramValue,
		&sizeReturn)
	if status != C.CL_SUCCESS {
		return 0, StatusError(status)
	}
	return uintptr(sizeReturn), nil
}

// DeviceInfoString is a convenience method for DeviceInfo() to query information values that are string-based.
//
// This function does not verify the queried information is indeed of type string. It assumes the information is
// a NUL terminated raw string and will extract the bytes as characters before that.
func DeviceInfoString(id DeviceID, paramName DeviceInfoName) (string, error) {
	return queryString(func(paramSize uintptr, paramValue unsafe.Pointer) (uintptr, error) {
		return DeviceInfo(id, paramName, paramSize, paramValue)
	})
}

const (
	// DevicePartitionEquallyProperty requests to split the aggregate device into as many smaller aggregate devices as
	// can be created, each containing N compute units. The value N is passed as the value accompanying this property.
	// If N does not divide evenly into DeviceMaxComputeUnitsInfo, then the remaining compute units are not used.
	//
	// Use PartitionedEqually() for convenience.
	//
	// Property value type: Uint
	// Since: 1.2
	DevicePartitionEquallyProperty uintptr = C.CL_DEVICE_PARTITION_EQUALLY
	// DevicePartitionByCountsProperty is followed by a list of compute unit counts terminated with 0 or
	// DevicePartitionByCountsListEndProperty. For each non-zero count M in the list, a sub-device is created with M compute
	// units in it.
	//
	// The number of non-zero count entries in the list may not exceed DevicePartitionMaxSubDevicesInfo.
	//
	// The total number of compute units specified may not exceed DeviceMaxComputeUnitsInfo.
	//
	// Use PartitionedByCounts() for convenience.
	//
	// Property value type: Uint
	// Since: 1.2
	DevicePartitionByCountsProperty uintptr = C.CL_DEVICE_PARTITION_BY_COUNTS
	// DevicePartitionByCountsListEndProperty terminates the property value list started by DevicePartitionByCountsProperty.
	//
	// Since: 1.2
	DevicePartitionByCountsListEndProperty uintptr = C.CL_DEVICE_PARTITION_BY_COUNTS_LIST_END
	// DevicePartitionByAffinityDomainProperty splits the device into smaller aggregate devices containing one or more
	// compute units that all share part of a cache hierarchy. The value accompanying this property may be drawn
	// from the constants of DeviceAffinityDomainFlags.
	//
	// Use PartitionedByAffinityDomain() for convenience.
	//
	// Property value type: DeviceAffinityDomainFlags
	// Since: 1.2
	DevicePartitionByAffinityDomainProperty uintptr = C.CL_DEVICE_PARTITION_BY_AFFINITY_DOMAIN
)

// DevicePartitionProperty is one entry of properties which are taken into account when creating sub-devices
// with CreateSubDevices().
type DevicePartitionProperty []uintptr

// PartitionedEqually is a convenience function to create a valid DevicePartitionEquallyProperty.
// Use it in combination with CreateSubDevices().
func PartitionedEqually(units uint) DevicePartitionProperty {
	return DevicePartitionProperty{DevicePartitionEquallyProperty, uintptr(units)}
}

// PartitionedByCounts is a convenience function to create a valid DevicePartitionByCountsProperty.
// Use it in combination with CreateSubDevices().
func PartitionedByCounts(units []uint) DevicePartitionProperty {
	values := make(DevicePartitionProperty, 0, len(units)+2)
	values = append(values, DevicePartitionByCountsProperty)
	for _, unit := range units {
		values = append(values, uintptr(unit))
	}
	values = append(values, DevicePartitionByCountsListEndProperty)
	return values
}

// DeviceAffinityDomainFlags describe how sub-devices are partitioned according to their cache hierarchy.
type DeviceAffinityDomainFlags C.cl_device_affinity_domain

const (
	// DeviceAffinityDomainNuma splits the device into sub-devices comprised of compute units that share a NUMA node.
	//
	// Since: 1.2
	DeviceAffinityDomainNuma DeviceAffinityDomainFlags = C.CL_DEVICE_AFFINITY_DOMAIN_NUMA
	// DeviceAffinityDomainL4Cache splits the device into sub-devices comprised of compute units that share a
	// level 4 data cache.
	//
	// Since: 1.2
	DeviceAffinityDomainL4Cache DeviceAffinityDomainFlags = C.CL_DEVICE_AFFINITY_DOMAIN_L4_CACHE
	// DeviceAffinityDomainL3Cache splits the device into sub-devices comprised of compute units that share a
	// level 3 data cache.
	//
	// Since: 1.2
	DeviceAffinityDomainL3Cache DeviceAffinityDomainFlags = C.CL_DEVICE_AFFINITY_DOMAIN_L3_CACHE
	// DeviceAffinityDomainL2Cache splits the device into sub-devices comprised of compute units that share a
	// level 2 data cache.
	//
	// Since: 1.2
	DeviceAffinityDomainL2Cache DeviceAffinityDomainFlags = C.CL_DEVICE_AFFINITY_DOMAIN_L2_CACHE
	// DeviceAffinityDomainL1Cache splits the device into sub-devices comprised of compute units that share a
	// level 1 data cache.
	//
	// Since: 1.2
	DeviceAffinityDomainL1Cache DeviceAffinityDomainFlags = C.CL_DEVICE_AFFINITY_DOMAIN_L1_CACHE
	// DeviceAffinityDomainNextPartitionable splits the device along the next partitionable affinity domain.
	// The implementation shall find the first level along which the device or sub-device may be further subdivided
	// in the order NUMA, L4, L3, L2, L1, and partition the device into sub-devices comprised of compute units that
	// share memory subsystems at this level.
	//
	// Determine what happened by calling DeviceInfo() with DevicePartitionTypeInfo on the sub-devices.
	//
	// Since: 1.2
	DeviceAffinityDomainNextPartitionable DeviceAffinityDomainFlags = C.CL_DEVICE_AFFINITY_DOMAIN_NEXT_PARTITIONABLE
)

// PartitionedByAffinityDomain is a convenience function to create a valid DevicePartitionByAffinityDomainProperty.
// Use it in combination with CreateSubDevices().
func PartitionedByAffinityDomain(domain DeviceAffinityDomainFlags) DevicePartitionProperty {
	return DevicePartitionProperty{DevicePartitionByAffinityDomainProperty, uintptr(domain)}
}

// CreateSubDevices creates an array of sub-devices that each reference a non-intersecting set of compute units within
// the device identified by id, according to the partition scheme given by properties.
// Only one of the available partitioning schemes can be specified in properties.
//
// The output sub-devices may be used in every way that the root (or parent) device can be used, including
// creating contexts, building programs, further calls to CreateSubDevices(), and creating command-queues.
// When a command-queue is created against a sub-device, the commands enqueued on the queue are executed only
// on the sub-device.
//
// Since: 1.2
// See also: https://registry.khronos.org/OpenCL/sdk/1.2/docs/man/xhtml/clCreateSubDevices.html
func CreateSubDevices(id DeviceID, properties ...DevicePartitionProperty) ([]DeviceID, error) {
	var rawPropertyList []uintptr
	for _, property := range properties {
		rawPropertyList = append(rawPropertyList, property...)
	}
	var rawProperties unsafe.Pointer
	if len(properties) > 0 {
		rawPropertyList = append(rawPropertyList, 0)
		rawProperties = unsafe.Pointer(&rawPropertyList[0])
	}

	requiredCount := C.cl_uint(0)
	status := C.clCreateSubDevices(
		id.handle(),
		(*C.cl_device_partition_property)(rawProperties),
		0, nil,
		&requiredCount)
	if status != C.CL_SUCCESS {
		return nil, StatusError(status)
	}
	ids := make([]DeviceID, requiredCount)
	reportedCount := C.cl_uint(0)
	status = C.clCreateSubDevices(
		id.handle(),
		(*C.cl_device_partition_property)(rawProperties),
		requiredCount,
		(*C.cl_device_id)(unsafe.Pointer(&ids[0])),
		&reportedCount)
	if status != C.CL_SUCCESS {
		return nil, StatusError(status)
	}
	return ids[:reportedCount], nil
}

// RetainDevice increments the device reference count if device is a valid sub-device created by a call to
// CreateSubDevices(). If id refers to a root level device, meaning a DeviceID returned by DeviceIDs(), the device
// reference count remains unchanged.
//
// Since: 1.2
// See also: https://registry.khronos.org/OpenCL/sdk/1.2/docs/man/xhtml/clRetainDevice.html
func RetainDevice(id DeviceID) error {
	status := C.clRetainDevice(id.handle())
	if status != C.CL_SUCCESS {
		return StatusError(status)
	}
	return nil
}

// ReleaseDevice decrements the device reference count if device is a valid sub-device created by a call to
// CreateSubDevices(). If id refers to a root level device, meaning a DeviceID returned by DeviceIDs(), the device
// reference count remains unchanged.
//
// Since: 1.2
// See also: https://registry.khronos.org/OpenCL/sdk/1.2/docs/man/xhtml/clReleaseDevice.html
func ReleaseDevice(id DeviceID) error {
	status := C.clReleaseDevice(id.handle())
	if status != C.CL_SUCCESS {
		return StatusError(status)
	}
	return nil
}
