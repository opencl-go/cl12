package cl12

// #include "api.h"
// extern cl_int cl12SetMemObjectDestructorCallback(cl_mem mem, uintptr_t *userData);
import "C"
import (
	"fmt"
	"unsafe"
)

// MemObject represents a reference counted region of global memory.
type MemObject uintptr

func (mem MemObject) handle() C.cl_mem {
	return *(*C.cl_mem)(unsafe.Pointer(&mem))
}

// String provides a readable presentation of the memory identifier.
// It is based on the numerical value of the underlying pointer.
func (mem MemObject) String() string {
	return fmt.Sprintf("0x%X", uintptr(mem))
}

// RetainMemObject increments the memory object reference count.
//
// Function that create a memory object perform an implicit retain.
//
// See also: https://registry.khronos.org/OpenCL/sdk/1.2/docs/man/xhtml/clRetainMemObject.html
func RetainMemObject(mem MemObject) error {
	status := C.clRetainMemObject(mem.handle())
	if status != C.CL_SUCCESS {
		return StatusError(status)
	}
	return nil
}

// ReleaseMemObject decrements the memory object reference count.
//
// After the reference count becomes zero and commands queued for execution on a command-queue(s) that use mem have
// finished, the memory object is deleted. If mem is a buffer object, mem cannot be deleted until all sub-buffer
// objects associated with mem are deleted.
//
// See also: https://registry.khronos.org/OpenCL/sdk/1.2/docs/man/xhtml/clReleaseMemObject.html
func ReleaseMemObject(mem MemObject) error {
	status := C.clReleaseMemObject(mem.handle())
	if status != C.CL_SUCCESS {
		return StatusError(status)
	}
	return nil
}

// SetMemObjectDestructorCallback registers a destructor callback function with a memory object.
//
// Each call to SetMemObjectDestructorCallback() registers the specified callback function on a destructor callback
// stack associated with mem.
// The registered callback functions are called in the reverse order in which they were registered.
//
// See also: https://registry.khronos.org/OpenCL/sdk/1.2/docs/man/xhtml/clSetMemObjectDestructorCallback.html
func SetMemObjectDestructorCallback(mem MemObject, callback func()) error {
	callbackUserData, err := userDataFor(callback)
	if err != nil {
		return err
	}
	status := C.cl12SetMemObjectDestructorCallback(mem.handle(), callbackUserData.ptr)
	if status != C.CL_SUCCESS {
		callbackUserData.Delete()
		return StatusError(status)
	}
	return nil
}

//export cl12GoMemObjectDestructorCallback
func cl12GoMemObjectDestructorCallback(_ MemObject, userData *C.uintptr_t) {
	callbackUserData := userDataFrom(userData)
	callback := callbackUserData.Value().(func())
	callbackUserData.Delete()
	callback()
}

// MemObjectInfoName identifies properties of a memory object, which can be queried with MemObjectInfo().
type MemObjectInfoName C.cl_mem_info

const (
	// MemTypeInfo returns the type of the memory object.
	//
	// Returned type: MemObjectType
	MemTypeInfo MemObjectInfoName = C.CL_MEM_TYPE
	// MemFlagsInfo returns the flags argument value specified when the memory object was created.
	// If the memory object is a sub-buffer the memory access qualifiers inherited from parent buffer are also returned.
	//
	// Returned type: MemFlags
	MemFlagsInfo MemObjectInfoName = C.CL_MEM_FLAGS
	// MemSizeInfo returns the actual size of the data store associated with the memory object in bytes.
	//
	// Returned type: uintptr
	MemSizeInfo MemObjectInfoName = C.CL_MEM_SIZE
	// MemHostPtrInfo returns the underlying host pointer for a MemObject if it (or its source buffer) was
	// created with the MemUseHostPtrFlag. It returns nil otherwise.
	//
	// Returned type: unsafe.Pointer
	MemHostPtrInfo MemObjectInfoName = C.CL_MEM_HOST_PTR
	// MemContextInfo returns the context specified when memory object is created.
	//
	// Returned type: Context
	MemContextInfo MemObjectInfoName = C.CL_MEM_CONTEXT
	// MemOffsetInfo returns the offset if memory object is a sub-buffer object created using CreateSubBuffer().
	// It returns 0 if memory object is not a sub-buffer object.
	//
	// Returned type: uintptr
	// Since: 1.1
	MemOffsetInfo MemObjectInfoName = C.CL_MEM_OFFSET
	// MemMapCountInfo returns the current map count.
	//
	// Note: The map count returned should be considered immediately stale. It is unsuitable for
	// general use in applications. This feature is provided for debugging.
	//
	// Returned type: uint32
	MemMapCountInfo MemObjectInfoName = C.CL_MEM_MAP_COUNT
	// MemReferenceCountInfo returns the memory reference count.
	//
	// Note: The reference count returned should be considered immediately stale. It is unsuitable for
	// general use in applications. This feature is provided for identifying memory leaks.
	//
	// Returned type: uint32
	MemReferenceCountInfo MemObjectInfoName = C.CL_MEM_REFERENCE_COUNT
	// MemAssociatedMemObjectInfo returns the memory object from which the queried memory object is created.
	//
	// Returned type: MemObject
	// Since: 1.1
	MemAssociatedMemObjectInfo MemObjectInfoName = C.CL_MEM_ASSOCIATED_MEMOBJECT
)

// MemObjectType identifies the specific type of MemObject.
type MemObjectType C.cl_mem_object_type

// These constants represent specific type identifier.
const (
	MemObjectBufferType  MemObjectType = C.CL_MEM_OBJECT_BUFFER
	MemObjectImage2DType MemObjectType = C.CL_MEM_OBJECT_IMAGE2D
	MemObjectImage3DType MemObjectType = C.CL_MEM_OBJECT_IMAGE3D

	MemObjectImage2DArrayType  MemObjectType = C.CL_MEM_OBJECT_IMAGE2D_ARRAY
	MemObjectImage1DType       MemObjectType = C.CL_MEM_OBJECT_IMAGE1D
	MemObjectImage1DArrayType  MemObjectType = C.CL_MEM_OBJECT_IMAGE1D_ARRAY
	MemObjectImage1DBufferType MemObjectType = C.CL_MEM_OBJECT_IMAGE1D_BUFFER
)

// MemFlags describe properties of a MemObject.
type MemFlags C.cl_mem_flags

// These constants identify possible properties of a MemObject.
const (
	MemReadWriteFlag    = C.CL_MEM_READ_WRITE
	MemWriteOnlyFlag    = C.CL_MEM_WRITE_ONLY
	MemReadOnlyFlag     = C.CL_MEM_READ_ONLY
	MemUseHostPtrFlag   = C.CL_MEM_USE_HOST_PTR
	MemAllocHostPtrFlag = C.CL_MEM_ALLOC_HOST_PTR
	MemCopyHostPtrFlag  = C.CL_MEM_COPY_HOST_PTR

	MemHostWriteOnlyFlag = C.CL_MEM_HOST_WRITE_ONLY
	MemHostReadOnlyFlag  = C.CL_MEM_HOST_READ_ONLY
	MemHostNoAccessFlag  = C.CL_MEM_HOST_NO_ACCESS
)

// MemObjectInfo queries information about a memory object.
//
// The provided size need to specify the size of the available space pointed to the provided value in bytes.
//
// The returned number is the required size, in bytes, for the queried information.
// Call the function with a zero size and nil value to request the required size. This helps in determining
// the necessary space for dynamic information, such as arrays.
//
// Raw strings are with a terminating NUL character.
//
// See also: https://registry.khronos.org/OpenCL/sdk/1.2/docs/man/xhtml/clGetMemObjectInfo.html
func MemObjectInfo(mem MemObject, paramName MemObjectInfoName, paramSize uintptr, paramValue unsafe.Pointer) (uintptr, error) {
	sizeReturn := C.size_t(0)
	status := C.clGetMemObjectInfo(
		mem.handle(),
		C.cl_mem_info(paramName),
		C.size_t(paramSize),
		paramValue,
		&sizeReturn)
	if status != C.CL_SUCCESS {
		return 0, StatusError(status)
	}
	return uintptr(sizeReturn), nil
}

// MapFlags describe how a memory object shall be mapped into host memory.
type MapFlags C.cl_map_flags

const (
	// MapRead specifies that the region being mapped in the memory object is being mapped for reading.
	MapRead MapFlags = C.CL_MAP_READ
	// MapWrite specifies that the region being mapped in the memory object is being mapped for writing.
	MapWrite MapFlags = C.CL_MAP_WRITE
	// MapWriteInvalidateRegion specifies that the region being mapped in the memory object is being mapped for writing.
	//
	// The contents of the region being mapped are to be discarded. This is typically the case when the region
	// being mapped is overwritten by the host. This flag allows the implementation to no longer guarantee that the
	// pointer returned by EnqueueMapBuffer() (EnqueueMapImage()) contains the latest bits in the region being mapped
	// which can be a significant performance enhancement.
	//
	// Since: 1.2
	MapWriteInvalidateRegion MapFlags = C.CL_MAP_WRITE_INVALIDATE_REGION
)

// EnqueueUnmapMemObject enqueues a command to unmap a previously mapped region of a memory object.
//
// Reads or writes from the host using the pointer returned by the mapping functions are considered to be complete.
//
// See also: https://registry.khronos.org/OpenCL/sdk/1.2/docs/man/xhtml/clEnqueueUnmapMemObject.html
func EnqueueUnmapMemObject(commandQueue CommandQueue, mem MemObject, mappedPtr unsafe.Pointer, waitList []Event, event *Event) error {
	var rawWaitList unsafe.Pointer
	if len(waitList) > 0 {
		rawWaitList = unsafe.Pointer(&waitList[0])
	}
	status := C.clEnqueueUnmapMemObject(
		commandQueue.handle(),
		mem.handle(),
		mappedPtr,
		C.cl_uint(len(waitList)),
		(*C.cl_event)(rawWaitList),
		(*C.cl_event)(unsafe.Pointer(event)))
	if status != C.CL_SUCCESS {
		return StatusError(status)
	}
	return nil
}

// MemMigrationFlags determine the migration options of memory objects.
type MemMigrationFlags C.cl_mem_migration_flags

const (
	// MigrateMemObjectHost indicates that the specified set of memory objects are to be migrated to the host,
	// regardless of the target command-queue.
	//
	// Since: 1.2
	MigrateMemObjectHost MemMigrationFlags = C.CL_MIGRATE_MEM_OBJECT_HOST
	// MigrateMemObjectContentUndefined indicates that the contents of the set of memory objects are undefined after
	// migration. The specified set of memory objects are migrated to the device associated with the command-queue
	// without incurring the overhead of migrating their contents.
	//
	// Since: 1.2
	MigrateMemObjectContentUndefined MemMigrationFlags = C.CL_MIGRATE_MEM_OBJECT_CONTENT_UNDEFINED
)

// EnqueueMigrateMemObjects enqueues a command to indicate which device a set of memory objects should be associated
// with.
//
// Typically, memory objects are implicitly migrated to a device for which enqueued commands, using the memory object,
// are targeted. EnqueueMigrateMemObjects() allows this migration to be explicitly performed ahead of the dependent
// commands. This allows a user to preemptively change the association of a memory object, through regular command
// queue scheduling, in order to prepare for another upcoming command. This also permits an application to overlap
// the placement of memory objects with other unrelated operations before these memory objects are needed potentially
// hiding transfer latencies.
//
// Since: 1.2
// See also: https://registry.khronos.org/OpenCL/sdk/1.2/docs/man/xhtml/clEnqueueMigrateMemObjects.html
func EnqueueMigrateMemObjects(commandQueue CommandQueue, memObjects []MemObject, migrationFlags MemMigrationFlags, waitList []Event, event *Event) error {
	var rawMemObjects unsafe.Pointer
	if len(memObjects) > 0 {
		rawMemObjects = unsafe.Pointer(&memObjects[0])
	}
	var rawWaitList unsafe.Pointer
	if len(waitList) > 0 {
		rawWaitList = unsafe.Pointer(&waitList[0])
	}
	status := C.clEnqueueMigrateMemObjects(
		commandQueue.handle(),
		C.cl_uint(len(memObjects)),
		(*C.cl_mem)(rawMemObjects),
		C.cl_mem_migration_flags(migrationFlags),
		C.cl_uint(len(waitList)),
		(*C.cl_event)(rawWaitList),
		(*C.cl_event)(unsafe.Pointer(event)))
	if status != C.CL_SUCCESS {
		return StatusError(status)
	}
	return nil
}
