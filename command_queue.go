package cl12

// #include "api.h"
import "C"
import (
	"fmt"
	"unsafe"
)

// CommandQueue describes a sequence of events for OpenCL operations.
// Create a new command-queue with CreateCommandQueue().
type CommandQueue uintptr

func (cq CommandQueue) handle() C.cl_command_queue {
	return *(*C.cl_command_queue)(unsafe.Pointer(&cq))
}

// String provides a readable presentation of the command-queue identifier.
// It is based on the numerical value of the underlying pointer.
func (cq CommandQueue) String() string {
	return fmt.Sprintf("0x%X", uintptr(cq))
}

// CommandQueuePropertiesFlags is used to determine DeviceQueuePropertiesInfo with DeviceInfo(), as well as
// properties for CreateCommandQueue().
type CommandQueuePropertiesFlags C.cl_command_queue_properties

const (
	// QueueOutOfOrderExecModeEnable determines whether the commands queued in the command-queue are executed
	// in-order or out-of-order. If set, the commands in the command-queue are executed out-of-order.
	// Otherwise, commands are executed in-order.
	QueueOutOfOrderExecModeEnable CommandQueuePropertiesFlags = C.CL_QUEUE_OUT_OF_ORDER_EXEC_MODE_ENABLE
	// QueueProfilingEnable enables or disables profiling of commands in the command-queue. If set,
	// the profiling of commands is enabled. Otherwise, profiling of commands is disabled.
	QueueProfilingEnable CommandQueuePropertiesFlags = C.CL_QUEUE_PROFILING_ENABLE
)

// CreateCommandQueue creates a command-queue on a specific device.
//
// See also: https://registry.khronos.org/OpenCL/sdk/1.2/docs/man/xhtml/clCreateCommandQueue.html
func CreateCommandQueue(context Context, deviceID DeviceID, properties CommandQueuePropertiesFlags) (CommandQueue, error) {
	var status C.cl_int
	commandQueue := C.clCreateCommandQueue(
		context.handle(),
		deviceID.handle(),
		C.cl_command_queue_properties(properties),
		&status)
	if status != C.CL_SUCCESS {
		return 0, StatusError(status)
	}
	return CommandQueue(*((*uintptr)(unsafe.Pointer(&commandQueue)))), nil
}

// RetainCommandQueue increments the commandQueue reference count.
//
// CreateCommandQueue() performs an implicit retain.
// This is very helpful for 3rd party libraries, which typically get a command-queue passed to them by the application.
// However, it is possible that the application may delete the command-queue without informing the library.
// Allowing functions to attach to (i.e. retain) and release a command-queue solves the problem of a command-queue
// being used by a library no longer being valid.
//
// See also: https://registry.khronos.org/OpenCL/sdk/1.2/docs/man/xhtml/clRetainCommandQueue.html
func RetainCommandQueue(commandQueue CommandQueue) error {
	status := C.clRetainCommandQueue(commandQueue.handle())
	if status != C.CL_SUCCESS {
		return StatusError(status)
	}
	return nil
}

// ReleaseCommandQueue decrements the commandQueue reference count.
//
// After the commandQueue reference count becomes zero and all commands queued to commandQueue have finished
// (eg. kernel-instances, memory object updates etc.), the command-queue is deleted.
//
// ReleaseCommandQueue() performs an implicit flush to issue any previously queued OpenCL commands in commandQueue.
//
// See also: https://registry.khronos.org/OpenCL/sdk/1.2/docs/man/xhtml/clReleaseCommandQueue.html
func ReleaseCommandQueue(commandQueue CommandQueue) error {
	status := C.clReleaseCommandQueue(commandQueue.handle())
	if status != C.CL_SUCCESS {
		return StatusError(status)
	}
	return nil
}

// CommandQueueInfoName identifies properties of a command-queue, which can be queried with CommandQueueInfo().
type CommandQueueInfoName C.cl_command_queue_info

const (
	// QueueContextInfo returns the context specified when the command-queue is created.
	//
	// Returned type: Context
	QueueContextInfo CommandQueueInfoName = C.CL_QUEUE_CONTEXT
	// QueueDeviceInfo returns the device specified when the command-queue is created
	//
	// Returned type: DeviceID
	QueueDeviceInfo CommandQueueInfoName = C.CL_QUEUE_DEVICE
	// QueueReferenceCountInfo returns the command-queue reference count.
	//
	// Note: The reference count returned should be considered immediately stale. It is unsuitable for
	// general use in applications. This feature is provided for identifying memory leaks.
	//
	// Returned type: Uint
	QueueReferenceCountInfo CommandQueueInfoName = C.CL_QUEUE_REFERENCE_COUNT
	// QueuePropertiesInfo returns the currently specified properties for the command-queue.
	// These properties are specified by the value associated with the QueuePropertiesProperty passed in as
	// the value of the properties argument in CreateCommandQueue().
	//
	// Returned type: uint64
	QueuePropertiesInfo CommandQueueInfoName = C.CL_QUEUE_PROPERTIES
)

// CommandQueueInfo queries information about a command-queue.
//
// The provided size need to specify the size of the available space pointed to the provided value in bytes.
//
// The returned number is the required size, in bytes, for the queried information.
// Call the function with a zero size and nil value to request the required size. This helps in determining
// the necessary space for dynamic information, such as arrays.
//
// Raw strings are with a terminating NUL character. For convenience, use CommandQueueInfoString().
//
// See also: https://registry.khronos.org/OpenCL/sdk/1.2/docs/man/xhtml/clGetCommandQueueInfo.html
func CommandQueueInfo(commandQueue CommandQueue, paramName CommandQueueInfoName, paramSize uintptr, paramValue unsafe.Pointer) (uintptr, error) {
	sizeReturn := C.size_t(0)
	status := C.clGetCommandQueueInfo(
		commandQueue.handle(),
		C.cl_command_queue_info(paramName),
		C.size_t(paramSize),
		paramValue,
		&sizeReturn)
	if status != C.CL_SUCCESS {
		return 0, StatusError(status)
	}
	return uintptr(sizeReturn), nil
}

// Flush issues all previously queued OpenCL commands in a command-queue to the device associated with the
// command-queue.
//
// All previously queued OpenCL commands in commandQueue are issued to the device associated with commandQueue.
// Flush() only guarantees that all queued commands to commandQueue will eventually be submitted to the appropriate
// device. There is no guarantee that they will be complete after Flush() returns.
//
// See also: https://registry.khronos.org/OpenCL/sdk/1.2/docs/man/xhtml/clFlush.html
func Flush(commandQueue CommandQueue) error {
	status := C.clFlush(commandQueue.handle())
	if status != C.CL_SUCCESS {
		return StatusError(status)
	}
	return nil
}

// Finish blocks until all previously queued OpenCL commands in a command-queue are issued to the associated device
// and have completed.
//
// All previously queued OpenCL commands in commandQueue are issued to the associated device, and the function blocks
// until all previously queued commands have completed. Finish() does not return until all previously queued commands
// in commandQueue have been processed and completed. Finish() is also a synchronization point.
//
// See also: https://registry.khronos.org/OpenCL/sdk/1.2/docs/man/xhtml/clFinish.html
func Finish(commandQueue CommandQueue) error {
	status := C.clFinish(commandQueue.handle())
	if status != C.CL_SUCCESS {
		return StatusError(status)
	}
	return nil
}
