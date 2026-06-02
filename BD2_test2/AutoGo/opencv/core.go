package opencv

import "C"
import (
	"errors"
	"image"
	"image/color"
	"unsafe"
)

const (
	// MatChannels1 is a single channel Mat.
	MatChannels1 = 0

	// MatChannels2 is 2 channel Mat.
	MatChannels2 = 8

	// MatChannels3 is 3 channel Mat.
	MatChannels3 = 16

	// MatChannels4 is 4 channel Mat.
	MatChannels4 = 24
)

// MatType is the type for the various different kinds of Mat you can create.
type MatType int

const (
	// MatTypeCV8U is a Mat of 8-bit unsigned int
	MatTypeCV8U MatType = 0

	// MatTypeCV8S is a Mat of 8-bit signed int
	MatTypeCV8S MatType = 1

	// MatTypeCV16U is a Mat of 16-bit unsigned int
	MatTypeCV16U MatType = 2

	// MatTypeCV16S is a Mat of 16-bit signed int
	MatTypeCV16S MatType = 3

	// MatTypeCV16SC2 is a Mat of 16-bit signed int with 2 channels
	MatTypeCV16SC2 = MatTypeCV16S + MatChannels2

	// MatTypeCV32S is a Mat of 32-bit signed int
	MatTypeCV32S MatType = 4

	// MatTypeCV32F is a Mat of 32-bit float
	MatTypeCV32F MatType = 5

	// MatTypeCV64F is a Mat of 64-bit float
	MatTypeCV64F MatType = 6

	// MatTypeCV16F is a Mat of 16-bit (half) float
	MatTypeCV16F MatType = 7

	// MatTypeCV8UC1 is a Mat of 8-bit unsigned int with a single channel
	MatTypeCV8UC1 = MatTypeCV8U + MatChannels1

	// MatTypeCV8UC2 is a Mat of 8-bit unsigned int with 2 channels
	MatTypeCV8UC2 = MatTypeCV8U + MatChannels2

	// MatTypeCV8UC3 is a Mat of 8-bit unsigned int with 3 channels
	MatTypeCV8UC3 = MatTypeCV8U + MatChannels3

	// MatTypeCV8UC4 is a Mat of 8-bit unsigned int with 4 channels
	MatTypeCV8UC4 = MatTypeCV8U + MatChannels4

	// MatTypeCV8SC1 is a Mat of 8-bit signed int with a single channel
	MatTypeCV8SC1 = MatTypeCV8S + MatChannels1

	// MatTypeCV8SC2 is a Mat of 8-bit signed int with 2 channels
	MatTypeCV8SC2 = MatTypeCV8S + MatChannels2

	// MatTypeCV8SC3 is a Mat of 8-bit signed int with 3 channels
	MatTypeCV8SC3 = MatTypeCV8S + MatChannels3

	// MatTypeCV8SC4 is a Mat of 8-bit signed int with 4 channels
	MatTypeCV8SC4 = MatTypeCV8S + MatChannels4

	// MatTypeCV16UC1 is a Mat of 16-bit unsigned int with a single channel
	MatTypeCV16UC1 = MatTypeCV16U + MatChannels1

	// MatTypeCV16UC2 is a Mat of 16-bit unsigned int with 2 channels
	MatTypeCV16UC2 = MatTypeCV16U + MatChannels2

	// MatTypeCV16UC3 is a Mat of 16-bit unsigned int with 3 channels
	MatTypeCV16UC3 = MatTypeCV16U + MatChannels3

	// MatTypeCV16UC4 is a Mat of 16-bit unsigned int with 4 channels
	MatTypeCV16UC4 = MatTypeCV16U + MatChannels4

	// MatTypeCV16SC1 is a Mat of 16-bit signed int with a single channel
	MatTypeCV16SC1 = MatTypeCV16S + MatChannels1

	// MatTypeCV16SC3 is a Mat of 16-bit signed int with 3 channels
	MatTypeCV16SC3 = MatTypeCV16S + MatChannels3

	// MatTypeCV16SC4 is a Mat of 16-bit signed int with 4 channels
	MatTypeCV16SC4 = MatTypeCV16S + MatChannels4

	// MatTypeCV32SC1 is a Mat of 32-bit signed int with a single channel
	MatTypeCV32SC1 = MatTypeCV32S + MatChannels1

	// MatTypeCV32SC2 is a Mat of 32-bit signed int with 2 channels
	MatTypeCV32SC2 = MatTypeCV32S + MatChannels2

	// MatTypeCV32SC3 is a Mat of 32-bit signed int with 3 channels
	MatTypeCV32SC3 = MatTypeCV32S + MatChannels3

	// MatTypeCV32SC4 is a Mat of 32-bit signed int with 4 channels
	MatTypeCV32SC4 = MatTypeCV32S + MatChannels4

	// MatTypeCV32FC1 is a Mat of 32-bit float int with a single channel
	MatTypeCV32FC1 = MatTypeCV32F + MatChannels1

	// MatTypeCV32FC2 is a Mat of 32-bit float int with 2 channels
	MatTypeCV32FC2 = MatTypeCV32F + MatChannels2

	// MatTypeCV32FC3 is a Mat of 32-bit float int with 3 channels
	MatTypeCV32FC3 = MatTypeCV32F + MatChannels3

	// MatTypeCV32FC4 is a Mat of 32-bit float int with 4 channels
	MatTypeCV32FC4 = MatTypeCV32F + MatChannels4

	// MatTypeCV64FC1 is a Mat of 64-bit float int with a single channel
	MatTypeCV64FC1 = MatTypeCV64F + MatChannels1

	// MatTypeCV64FC2 is a Mat of 64-bit float int with 2 channels
	MatTypeCV64FC2 = MatTypeCV64F + MatChannels2

	// MatTypeCV64FC3 is a Mat of 64-bit float int with 3 channels
	MatTypeCV64FC3 = MatTypeCV64F + MatChannels3

	// MatTypeCV64FC4 is a Mat of 64-bit float int with 4 channels
	MatTypeCV64FC4 = MatTypeCV64F + MatChannels4

	// MatTypeCV16FC1 is a Mat of 16-bit float with a single channel
	MatTypeCV16FC1 = MatTypeCV16F + MatChannels1

	// MatTypeCV16FC2 is a Mat of 16-bit float with 2 channels
	MatTypeCV16FC2 = MatTypeCV16F + MatChannels2

	// MatTypeCV16FC3 is a Mat of 16-bit float with 3 channels
	MatTypeCV16FC3 = MatTypeCV16F + MatChannels3

	// MatTypeCV16FC4 is a Mat of 16-bit float with 4 channels
	MatTypeCV16FC4 = MatTypeCV16F + MatChannels4
)

// CompareType is used for Compare operations to indicate which kind of
// comparison to use.
type CompareType int

const (
	// CompareEQ src1 is equal to src2.
	CompareEQ CompareType = 0

	// CompareGT src1 is greater than src2.
	CompareGT CompareType = 1

	// CompareGE src1 is greater than or equal to src2.
	CompareGE CompareType = 2

	// CompareLT src1 is less than src2.
	CompareLT CompareType = 3

	// CompareLE src1 is less than or equal to src2.
	CompareLE CompareType = 4

	// CompareNE src1 is unequal to src2.
	CompareNE CompareType = 5
)

type Point2f struct {
	X float32
	Y float32
}

func NewPoint2f(x, y float32) Point2f {
	return Point2f{}
}

var ErrEmptyByteSlice = errors.New("empty byte array")

// Mat represents an n-dimensional dense numerical single-channel
// or multi-channel array. It can be used to store real or complex-valued
// vectors and matrices, grayscale or color images, voxel volumes,
// vector fields, point clouds, tensors, and histograms.
//
// For further details, please see:
// http://docs.opencv.org/master/d3/d63/classcv_1_1Mat.html
type Mat struct {
	// Do nothing
}

// NewMat returns a new empty Mat.
func NewMat() Mat {
	return Mat{}
}

// NewMatFromCMat returns a new Mat from an unsafe.Pointer(C.Mat).
func NewMatFromCMat(c_mat unsafe.Pointer) Mat {
	return Mat{}
}

// NewMatWithSize returns a new Mat with a specific size and type.
func NewMatWithSize(rows int, cols int, mt MatType) Mat {
	return Mat{}
}

// NewMatWithSizes returns a new multidimensional Mat with a specific size and type.
func NewMatWithSizes(sizes []int, mt MatType) Mat {
	return Mat{}
}

// NewMatWithSizesWithScalar returns a new multidimensional Mat with a specific size, type and scalar value.
func NewMatWithSizesWithScalar(sizes []int, mt MatType, s Scalar) Mat {
	return Mat{}
}

// NewMatWithSizesWithScalar returns a new multidimensional Mat with a specific size, type and preexisting data.
func NewMatWithSizesFromBytes(sizes []int, mt MatType, data []byte) (Mat, error) {
	return Mat{}, nil
}

// NewMatFromScalar returns a new Mat for a specific Scalar value
func NewMatFromScalar(s Scalar, mt MatType) Mat {
	return Mat{}
}

// NewMatWithSizeFromScalar returns a new Mat for a specific Scala value with a specific size and type
// This simplifies creation of specific color filters or creating Mats of specific colors and sizes
func NewMatWithSizeFromScalar(s Scalar, rows int, cols int, mt MatType) Mat {
	return Mat{}
}

// NewMatFromBytes returns a new Mat with a specific size and type, initialized from a []byte.
func NewMatFromBytes(rows int, cols int, mt MatType, data []byte) (Mat, error) {
	return Mat{}, nil
}

func (m *Mat) Close() error {
	return nil
}

// Returns an identity matrix of the specified size and type.
//
// The method returns a Matlab-style identity matrix initializer, similarly to Mat::zeros. Similarly to Mat::ones.
// For further details, please see:
// https://docs.opencv.org/master/d3/d63/classcv_1_1Mat.html#a2cf9b9acde7a9852542bbc20ef851ed2
func Eye(rows int, cols int, mt MatType) Mat {
	return Mat{}
}

// Returns a zero array of the specified size and type.
//
// The method returns a Matlab-style zero array initializer.
// For further details, please see:
// https://docs.opencv.org/master/d3/d63/classcv_1_1Mat.html#a0b57b6a326c8876d944d188a46e0f556
func Zeros(rows int, cols int, mt MatType) Mat {
	return Mat{}
}

// Returns an array of all 1's of the specified size and type.
//
// The method returns a Matlab-style 1's array initializer
// For further details, please see:
// https://docs.opencv.org/master/d3/d63/classcv_1_1Mat.html#a69ae0402d116fc9c71908d8508dc2f09
func Ones(rows int, cols int, mt MatType) Mat {
	return Mat{}
}

// FromPtr returns a new Mat with a specific size and type, initialized from a Mat Ptr.
func (m *Mat) FromPtr(rows int, cols int, mt MatType, prow int, pcol int) (Mat, error) {
	return Mat{}, nil
}

// Ptr returns the Mat's underlying object pointer.
func (m *Mat) Ptr() C.Mat {
	return nil
}

// Empty determines if the Mat is empty or not.
func (m *Mat) Empty() bool {
	return false
}

// Closed determines if the Mat is closed or not.
func (m *Mat) Closed() bool {
	return false
}

// IsContinuous determines if the Mat is continuous.
//
// For further details, please see:
// https://docs.opencv.org/master/d3/d63/classcv_1_1Mat.html#aa90cea495029c7d1ee0a41361ccecdf3
func (m *Mat) IsContinuous() bool {
	return false
}

// Inv inverses a matrix.
//
// For further details, please see:
// https://docs.opencv.org/4.x/d3/d63/classcv_1_1Mat.html#a039eb3c6740a850696a12519a4b8bfc6
func (m *Mat) Inv() {

}

// Col creates a matrix header for the specified matrix column.
// The underlying data of the new matrix is shared with the original matrix.
func (m *Mat) Col(col int) Mat {
	return Mat{}
}

// Row creates a matrix header for the specified matrix row.
// The underlying data of the new matrix is shared with the original matrix.
func (m *Mat) Row(row int) Mat {
	return Mat{}
}

// Clone returns a cloned full copy of the Mat.
func (m *Mat) Clone() Mat {
	return Mat{}
}

// CopyTo copies Mat into destination Mat.
//
// For further details, please see:
// https://docs.opencv.org/master/d3/d63/classcv_1_1Mat.html#a33fd5d125b4c302b0c9aa86980791a77
func (m *Mat) CopyTo(dst *Mat) {
	// Do nothing
}

// CopyToWithMask copies Mat into destination Mat after applying the mask Mat.
//
// For further details, please see:
// https://docs.opencv.org/master/d3/d63/classcv_1_1Mat.html#a626fe5f96d02525e2604d2ad46dd574f
func (m *Mat) CopyToWithMask(dst *Mat, mask Mat) {
	// Do nothing
}

// ConvertTo converts Mat into destination Mat.
//
// For further details, please see:
// https://docs.opencv.org/master/d3/d63/classcv_1_1Mat.html#adf88c60c5b4980e05bb556080916978b
func (m *Mat) ConvertTo(dst *Mat, mt MatType) {
	// Do nothing
}

func (m *Mat) ConvertToWithParams(dst *Mat, mt MatType, alpha, beta float32) {
	// Do nothing
}

// Total returns the total number of array elements.
//
// For further details, please see:
// https://docs.opencv.org/master/d3/d63/classcv_1_1Mat.html#aa4d317d43fb0cba9c2503f3c61b866c8
func (m *Mat) Total() int {
	return 0
}

// Size returns an array with one element for each dimension containing the size of that dimension for the Mat.
//
// For further details, please see:
// https://docs.opencv.org/master/d3/d63/classcv_1_1Mat.html#aa4d317d43fb0cba9c2503f3c61b866c8
func (m *Mat) Size() (dims []int) {
	return nil
}

// ToBytes copies the underlying Mat data to a byte array.
//
// For further details, please see:
// https://docs.opencv.org/3.3.1/d3/d63/classcv_1_1Mat.html#a4d33bed1c850265370d2af0ff02e1564
func (m *Mat) ToBytes() []byte {
	return nil
}

// DataPtrUint8 returns a slice that references the OpenCV allocated data.
//
// The data is no longer valid once the Mat has been closed. Any data that
// needs to be accessed after the Mat is closed must be copied into Go memory.
func (m *Mat) DataPtrUint8() ([]uint8, error) {
	return nil, nil
}

// DataPtrInt8 returns a slice that references the OpenCV allocated data.
//
// The data is no longer valid once the Mat has been closed. Any data that
// needs to be accessed after the Mat is closed must be copied into Go memory.
func (m *Mat) DataPtrInt8() ([]int8, error) {
	return nil, nil
}

// DataPtrUint16 returns a slice that references the OpenCV allocated data.
//
// The data is no longer valid once the Mat has been closed. Any data that
// needs to be accessed after the Mat is closed must be copied into Go memory.
func (m *Mat) DataPtrUint16() ([]uint16, error) {
	return nil, nil
}

// DataPtrInt16 returns a slice that references the OpenCV allocated data.
//
// The data is no longer valid once the Mat has been closed. Any data that
// needs to be accessed after the Mat is closed must be copied into Go memory.
func (m *Mat) DataPtrInt16() ([]int16, error) {
	return nil, nil
}

// DataPtrFloat32 returns a slice that references the OpenCV allocated data.
//
// The data is no longer valid once the Mat has been closed. Any data that
// needs to be accessed after the Mat is closed must be copied into Go memory.
func (m *Mat) DataPtrFloat32() ([]float32, error) {
	return nil, nil
}

// DataPtrFloat64 returns a slice that references the OpenCV allocated data.
//
// The data is no longer valid once the Mat has been closed. Any data that
// needs to be accessed after the Mat is closed must be copied into Go memory.
func (m *Mat) DataPtrFloat64() ([]float64, error) {
	return nil, nil
}

// Region returns a new Mat that points to a region of this Mat. Changes made to the
// region Mat will affect the original Mat, since they are pointers to the underlying
// OpenCV Mat object.
func (m *Mat) Region(rio image.Rectangle) Mat {
	return Mat{}
}

// Reshape changes the shape and/or the number of channels of a 2D matrix without copying the data.
//
// For further details, please see:
// https://docs.opencv.org/master/d3/d63/classcv_1_1Mat.html#a4eb96e3251417fa88b78e2abd6cfd7d8
func (m *Mat) Reshape(cn int, rows int) Mat {
	return Mat{}
}

// ConvertFp16 converts a Mat to half-precision floating point.
//
// For further details, please see:
// https://docs.opencv.org/master/d2/de8/group__core__array.html#ga9c25d9ef44a2a48ecc3774b30cb80082
func (m *Mat) ConvertFp16() Mat {
	return Mat{}
}

// Mean calculates the mean value M of array elements, independently for each channel, and return it as Scalar
// For further details, please see:
// https://docs.opencv.org/master/d2/de8/group__core__array.html#ga191389f8a0e58180bb13a727782cd461
func (m *Mat) Mean() Scalar {
	return Scalar{}
}

// MeanWithMask calculates the mean value M of array elements,independently for each channel,
// and returns it as Scalar vector while applying the mask.
// https://docs.opencv.org/master/d2/de8/group__core__array.html#ga191389f8a0e58180bb13a727782cd461
func (m *Mat) MeanWithMask(mask Mat) Scalar {
	return Scalar{}
}

// Sqrt calculates a square root of array elements.
//
// For further details, please see:
// https://docs.opencv.org/master/d2/de8/group__core__array.html#ga186222c3919657890f88df5a1f64a7d7
func (m *Mat) Sqrt() Mat {
	return Mat{}
}

// Sum calculates the per-channel pixel sum of an image.
//
// For further details, please see:
// https://docs.opencv.org/master/d2/de8/group__core__array.html#ga716e10a2dd9e228e4d3c95818f106722
func (m *Mat) Sum() Scalar {
	return Scalar{}
}

// PatchNaNs converts NaN's to zeros.
//
// For further details, please see:
// https://docs.opencv.org/master/d2/de8/group__core__array.html#ga62286befb7cde3568ff8c7d14d5079da
func (m *Mat) PatchNaNs() {
	// Do nothing
}

// LUT performs a look-up table transform of an array.
//
// The function LUT fills the output array with values from the look-up table.
// Indices of the entries are taken from the input array.
//
// For further details, please see:
// https://docs.opencv.org/master/d2/de8/group__core__array.html#gab55b8d062b7f5587720ede032d34156f
func LUT(src, wbLUT Mat, dst *Mat) {
	// Do nothing
}

// Rows returns the number of rows for this Mat.
func (m *Mat) Rows() int {
	return 0
}

// Cols returns the number of columns for this Mat.
func (m *Mat) Cols() int {
	return 0
}

// Channels returns the number of channels for this Mat.
func (m *Mat) Channels() int {
	return 0
}

// Type returns the type for this Mat.
func (m *Mat) Type() MatType {
	return 0
}

// Step returns the number of bytes each matrix row occupies.
func (m *Mat) Step() int {
	return 0
}

// ElemSize returns the matrix element size in bytes.
func (m *Mat) ElemSize() int {
	return 0
}

// GetUCharAt returns a value from a specific row/col
// in this Mat expecting it to be of type uchar aka CV_8U.
func (m *Mat) GetUCharAt(row int, col int) uint8 {
	return 0
}

// GetUCharAt3 returns a value from a specific x, y, z coordinate location
// in this Mat expecting it to be of type uchar aka CV_8U.
func (m *Mat) GetUCharAt3(x, y, z int) uint8 {
	return 0
}

// GetSCharAt returns a value from a specific row/col
// in this Mat expecting it to be of type schar aka CV_8S.
func (m *Mat) GetSCharAt(row int, col int) int8 {
	return 0
}

// GetSCharAt3 returns a value from a specific x, y, z coordinate location
// in this Mat expecting it to be of type schar aka CV_8S.
func (m *Mat) GetSCharAt3(x, y, z int) int8 {
	return 0
}

// GetShortAt returns a value from a specific row/col
// in this Mat expecting it to be of type short aka CV_16S.
func (m *Mat) GetShortAt(row int, col int) int16 {
	return 0
}

// GetShortAt3 returns a value from a specific x, y, z coordinate location
// in this Mat expecting it to be of type short aka CV_16S.
func (m *Mat) GetShortAt3(x, y, z int) int16 {
	return 0
}

// GetIntAt returns a value from a specific row/col
// in this Mat expecting it to be of type int aka CV_32S.
func (m *Mat) GetIntAt(row int, col int) int32 {
	return 0
}

// GetIntAt3 returns a value from a specific x, y, z coordinate location
// in this Mat expecting it to be of type int aka CV_32S.
func (m *Mat) GetIntAt3(x, y, z int) int32 {
	return 0
}

// GetFloatAt returns a value from a specific row/col
// in this Mat expecting it to be of type float aka CV_32F.
func (m *Mat) GetFloatAt(row int, col int) float32 {
	return 0
}

// GetFloatAt3 returns a value from a specific x, y, z coordinate location
// in this Mat expecting it to be of type float aka CV_32F.
func (m *Mat) GetFloatAt3(x, y, z int) float32 {
	return 0
}

// GetDoubleAt returns a value from a specific row/col
// in this Mat expecting it to be of type double aka CV_64F.
func (m *Mat) GetDoubleAt(row int, col int) float64 {
	return 0
}

// GetDoubleAt3 returns a value from a specific x, y, z coordinate location
// in this Mat expecting it to be of type double aka CV_64F.
func (m *Mat) GetDoubleAt3(x, y, z int) float64 {
	return 0
}

// SetTo sets all or some of the array elements to the specified scalar value.
func (m *Mat) SetTo(s Scalar) {
	// Do nothing
}

// SetUCharAt sets a value at a specific row/col
// in this Mat expecting it to be of type uchar aka CV_8U.
func (m *Mat) SetUCharAt(row int, col int, val uint8) {
	// Do nothing
}

// SetUCharAt3 sets a value at a specific x, y, z coordinate location
// in this Mat expecting it to be of type uchar aka CV_8U.
func (m *Mat) SetUCharAt3(x, y, z int, val uint8) {
	// Do nothing
}

// SetSCharAt sets a value at a specific row/col
// in this Mat expecting it to be of type schar aka CV_8S.
func (m *Mat) SetSCharAt(row int, col int, val int8) {
	// Do nothing
}

// SetSCharAt3 sets a value at a specific x, y, z coordinate location
// in this Mat expecting it to be of type schar aka CV_8S.
func (m *Mat) SetSCharAt3(x, y, z int, val int8) {
	// Do nothing
}

// SetShortAt sets a value at a specific row/col
// in this Mat expecting it to be of type short aka CV_16S.
func (m *Mat) SetShortAt(row int, col int, val int16) {
	// Do nothing
}

// SetShortAt3 sets a value at a specific x, y, z coordinate location
// in this Mat expecting it to be of type short aka CV_16S.
func (m *Mat) SetShortAt3(x, y, z int, val int16) {
	// Do nothing
}

// SetIntAt sets a value at a specific row/col
// in this Mat expecting it to be of type int aka CV_32S.
func (m *Mat) SetIntAt(row int, col int, val int32) {
	// Do nothing
}

// SetIntAt3 sets a value at a specific x, y, z coordinate location
// in this Mat expecting it to be of type int aka CV_32S.
func (m *Mat) SetIntAt3(x, y, z int, val int32) {
	// Do nothing
}

// SetFloatAt sets a value at a specific row/col
// in this Mat expecting it to be of type float aka CV_32F.
func (m *Mat) SetFloatAt(row int, col int, val float32) {
	// Do nothing
}

// SetFloatAt3 sets a value at a specific x, y, z coordinate location
// in this Mat expecting it to be of type float aka CV_32F.
func (m *Mat) SetFloatAt3(x, y, z int, val float32) {
	// Do nothing
}

// SetDoubleAt sets a value at a specific row/col
// in this Mat expecting it to be of type double aka CV_64F.
func (m *Mat) SetDoubleAt(row int, col int, val float64) {
	// Do nothing
}

// SetDoubleAt3 sets a value at a specific x, y, z coordinate location
// in this Mat expecting it to be of type double aka CV_64F.
func (m *Mat) SetDoubleAt3(x, y, z int, val float64) {
	// Do nothing
}

// AddUChar adds a uchar value to each element in the Mat. Performs a
// mat += val operation.
func (m *Mat) AddUChar(val uint8) {
	// Do nothing
}

// SubtractUChar subtracts a uchar value from each element in the Mat. Performs a
// mat -= val operation.
func (m *Mat) SubtractUChar(val uint8) {
	// Do nothing
}

// MultiplyUChar multiplies each element in the Mat by a uint value. Performs a
// mat *= val operation.
func (m *Mat) MultiplyUChar(val uint8) {
	// Do nothing
}

// DivideUChar divides each element in the Mat by a uint value. Performs a
// mat /= val operation.
func (m *Mat) DivideUChar(val uint8) {
	// Do nothing
}

// AddFloat adds a float value to each element in the Mat. Performs a
// mat += val operation.
func (m *Mat) AddFloat(val float32) {
	// Do nothing
}

// SubtractFloat subtracts a float value from each element in the Mat. Performs a
// mat -= val operation.
func (m *Mat) SubtractFloat(val float32) {
	// Do nothing
}

// MultiplyFloat multiplies each element in the Mat by a float value. Performs a
// mat *= val operation.
func (m *Mat) MultiplyFloat(val float32) {
	// Do nothing
}

// DivideFloat divides each element in the Mat by a float value. Performs a
// mat /= val operation.
func (m *Mat) DivideFloat(val float32) {
	// Do nothing
}

// MultiplyMatrix multiplies matrix (m*x)
func (m *Mat) MultiplyMatrix(x Mat) Mat {
	return Mat{}
}

// T  transpose matrix
// https://docs.opencv.org/4.1.2/d3/d63/classcv_1_1Mat.html#aaa428c60ccb6d8ea5de18f63dfac8e11
func (m *Mat) T() Mat {
	return Mat{}
}

// AbsDiff calculates the per-element absolute difference between two arrays
// or between an array and a scalar.
//
// For further details, please see:
// https://docs.opencv.org/master/d2/de8/group__core__array.html#ga6fef31bc8c4071cbc114a758a2b79c14
func AbsDiff(src1, src2 Mat, dst *Mat) {
	// Do nothing
}

// Add calculates the per-element sum of two arrays or an array and a scalar.
//
// For further details, please see:
// https://docs.opencv.org/master/d2/de8/group__core__array.html#ga10ac1bfb180e2cfda1701d06c24fdbd6
func Add(src1, src2 Mat, dst *Mat) {
	// Do nothing
}

// AddWeighted calculates the weighted sum of two arrays.
//
// For further details, please see:
// https://docs.opencv.org/master/d2/de8/group__core__array.html#gafafb2513349db3bcff51f54ee5592a19
func AddWeighted(src1 Mat, alpha float64, src2 Mat, beta float64, gamma float64, dst *Mat) {
	// Do nothing
}

// BitwiseAnd computes bitwise conjunction of the two arrays (dst = src1 & src2).
// Calculates the per-element bit-wise conjunction of two arrays
// or an array and a scalar.
//
// For further details, please see:
// https://docs.opencv.org/master/d2/de8/group__core__array.html#ga60b4d04b251ba5eb1392c34425497e14
func BitwiseAnd(src1 Mat, src2 Mat, dst *Mat) {
	// Do nothing
}

// BitwiseAndWithMask computes bitwise conjunction of the two arrays (dst = src1 & src2).
// Calculates the per-element bit-wise conjunction of two arrays
// or an array and a scalar. It has an additional parameter for a mask.
//
// For further details, please see:
// https://docs.opencv.org/master/d2/de8/group__core__array.html#ga60b4d04b251ba5eb1392c34425497e14
func BitwiseAndWithMask(src1 Mat, src2 Mat, dst *Mat, mask Mat) {
	// Do nothing
}

// BitwiseNot inverts every bit of an array.
//
// For further details, please see:
// https://docs.opencv.org/master/d2/de8/group__core__array.html#ga0002cf8b418479f4cb49a75442baee2f
func BitwiseNot(src1 Mat, dst *Mat) {
	// Do nothing
}

// BitwiseNotWithMask inverts every bit of an array. It has an additional parameter for a mask.
//
// For further details, please see:
// https://docs.opencv.org/master/d2/de8/group__core__array.html#ga0002cf8b418479f4cb49a75442baee2f
func BitwiseNotWithMask(src1 Mat, dst *Mat, mask Mat) {
	// Do nothing
}

// BitwiseOr calculates the per-element bit-wise disjunction of two arrays
// or an array and a scalar.
//
// For further details, please see:
// https://docs.opencv.org/master/d2/de8/group__core__array.html#gab85523db362a4e26ff0c703793a719b4
func BitwiseOr(src1 Mat, src2 Mat, dst *Mat) {
	// Do nothing
}

// BitwiseOrWithMask calculates the per-element bit-wise disjunction of two arrays
// or an array and a scalar. It has an additional parameter for a mask.
//
// For further details, please see:
// https://docs.opencv.org/master/d2/de8/group__core__array.html#gab85523db362a4e26ff0c703793a719b4
func BitwiseOrWithMask(src1 Mat, src2 Mat, dst *Mat, mask Mat) {
	// Do nothing
}

// BitwiseXor calculates the per-element bit-wise "exclusive or" operation
// on two arrays or an array and a scalar.
//
// For further details, please see:
// https://docs.opencv.org/master/d2/de8/group__core__array.html#ga84b2d8188ce506593dcc3f8cd00e8e2c
func BitwiseXor(src1 Mat, src2 Mat, dst *Mat) {
	// Do nothing
}

// BitwiseXorWithMask calculates the per-element bit-wise "exclusive or" operation
// on two arrays or an array and a scalar. It has an additional parameter for a mask.
//
// For further details, please see:
// https://docs.opencv.org/master/d2/de8/group__core__array.html#ga84b2d8188ce506593dcc3f8cd00e8e2c
func BitwiseXorWithMask(src1 Mat, src2 Mat, dst *Mat, mask Mat) {
	// Do nothing
}

// BatchDistance is a naive nearest neighbor finder.
//
// For further details, please see:
// https://docs.opencv.org/master/d2/de8/group__core__array.html#ga4ba778a1c57f83233b1d851c83f5a622
func BatchDistance(src1 Mat, src2 Mat, dist Mat, dtype MatType, nidx Mat, normType NormType, K int, mask Mat, update int, crosscheck bool) {
	// Do nothing
}

// BorderInterpolate computes the source location of an extrapolated pixel.
//
// For further details, please see:
// https://docs.opencv.org/master/d2/de8/group__core__array.html#ga247f571aa6244827d3d798f13892da58
func BorderInterpolate(p int, len int, borderType CovarFlags) int {
	return 0
}

// CovarFlags are the covariation flags used by functions such as BorderInterpolate.
//
// For further details, please see:
// https://docs.opencv.org/master/d0/de1/group__core.html#ga719ebd4a73f30f4fab258ab7616d0f0f
type CovarFlags int

const (
	// CovarScrambled indicates to scramble the results.
	CovarScrambled CovarFlags = 0

	// CovarNormal indicates to use normal covariation.
	CovarNormal CovarFlags = 1

	// CovarUseAvg indicates to use average covariation.
	CovarUseAvg CovarFlags = 2

	// CovarScale indicates to use scaled covariation.
	CovarScale CovarFlags = 4

	// CovarRows indicates to use covariation on rows.
	CovarRows CovarFlags = 8

	// CovarCols indicates to use covariation on columns.
	CovarCols CovarFlags = 16
)

// CalcCovarMatrix calculates the covariance matrix of a set of vectors.
//
// For further details, please see:
// https://docs.opencv.org/master/d2/de8/group__core__array.html#ga017122d912af19d7d0d2cccc2d63819f
func CalcCovarMatrix(samples Mat, covar *Mat, mean *Mat, flags CovarFlags, ctype MatType) {
	// Do nothing
}

// CartToPolar calculates the magnitude and angle of 2D vectors.
//
// For further details, please see:
// https://docs.opencv.org/master/d2/de8/group__core__array.html#gac5f92f48ec32cacf5275969c33ee837d
func CartToPolar(x Mat, y Mat, magnitude *Mat, angle *Mat, angleInDegrees bool) {
	// Do nothing
}

// CheckRange checks every element of an input array for invalid values.
//
// For further details, please see:
// https://docs.opencv.org/master/d2/de8/group__core__array.html#ga2bd19d89cae59361416736f87e3c7a64
func CheckRange(src Mat) bool {
	return false
}

// Compare performs the per-element comparison of two arrays
// or an array and scalar value.
//
// For further details, please see:
// https://docs.opencv.org/master/d2/de8/group__core__array.html#ga303cfb72acf8cbb36d884650c09a3a97
func Compare(src1 Mat, src2 Mat, dst *Mat, ct CompareType) {
	// Do nothing
}

// CountNonZero counts non-zero array elements.
//
// For further details, please see:
// https://docs.opencv.org/master/d2/de8/group__core__array.html#gaa4b89393263bb4d604e0fe5986723914
func CountNonZero(src Mat) int {
	return 0
}

// CompleteSymm copies the lower or the upper half of a square matrix to its another half.
//
// For further details, please see:
// https://docs.opencv.org/master/d2/de8/group__core__array.html#gaa9d88dcd0e54b6d1af38d41f2a3e3d25
func CompleteSymm(m Mat, lowerToUpper bool) {
	// Do nothing
}

// ConvertScaleAbs scales, calculates absolute values, and converts the result to 8-bit.
//
// For further details, please see:
// https://docs.opencv.org/master/d2/de8/group__core__array.html#ga3460e9c9f37b563ab9dd550c4d8c4e7d
func ConvertScaleAbs(src Mat, dst *Mat, alpha float64, beta float64) {
	// Do nothing
}

// CopyMakeBorder forms a border around an image (applies padding).
//
// For further details, please see:
// https://docs.opencv.org/master/d2/de8/group__core__array.html#ga2ac1049c2c3dd25c2b41bffe17658a36
func CopyMakeBorder(src Mat, dst *Mat, top int, bottom int, left int, right int, bt BorderType, value color.RGBA) {
	// Do nothing
}

// DftFlags represents a DFT or DCT flag.
//
// For further details, please see:
// https://docs.opencv.org/master/d2/de8/group__core__array.html#gaf4dde112b483b38175621befedda1f1c
type DftFlags int

const (
	// DftForward performs forward 1D or 2D dft or dct.
	DftForward DftFlags = 0

	// DftInverse performs an inverse 1D or 2D transform.
	DftInverse DftFlags = 1

	// DftScale scales the result: divide it by the number of array elements. Normally, it is combined with DFT_INVERSE.
	DftScale DftFlags = 2

	// DftRows performs a forward or inverse transform of every individual row of the input matrix.
	DftRows DftFlags = 4

	// DftComplexOutput performs a forward transformation of 1D or 2D real array; the result, though being a complex array, has complex-conjugate symmetry
	DftComplexOutput DftFlags = 16

	// DftRealOutput performs an inverse transformation of a 1D or 2D complex array; the result is normally a complex array of the same size,
	// however, if the input array has conjugate-complex symmetry (for example, it is a result of forward transformation with DFT_COMPLEX_OUTPUT flag),
	// the output is a real array.
	DftRealOutput DftFlags = 32

	// DftComplexInput specifies that input is complex input. If this flag is set, the input must have 2 channels.
	DftComplexInput DftFlags = 64

	// DctInverse performs an inverse 1D or 2D dct transform.
	DctInverse = DftInverse

	// DctRows performs a forward or inverse dct transform of every individual row of the input matrix.
	DctRows = DftRows
)

// DCT performs a forward or inverse discrete Cosine transform of 1D or 2D array.
//
// For further details, please see:
// https://docs.opencv.org/master/d2/de8/group__core__array.html#ga85aad4d668c01fbd64825f589e3696d4
func DCT(src Mat, dst *Mat, flags DftFlags) {
	// Do nothing
}

// Determinant returns the determinant of a square floating-point matrix.
//
// For further details, please see:
// https://docs.opencv.org/master/d2/de8/group__core__array.html#gaf802bd9ca3e07b8b6170645ef0611d0c
func Determinant(src Mat) float64 {
	return 0
}

// DFT performs a forward or inverse Discrete Fourier Transform (DFT)
// of a 1D or 2D floating-point array.
//
// For further details, please see:
// https://docs.opencv.org/master/d2/de8/group__core__array.html#gadd6cf9baf2b8b704a11b5f04aaf4f39d
func DFT(src Mat, dst *Mat, flags DftFlags) {
	// Do nothing
}

// Divide performs the per-element division
// on two arrays or an array and a scalar.
//
// For further details, please see:
// https://docs.opencv.org/master/d2/de8/group__core__array.html#ga6db555d30115642fedae0cda05604874
func Divide(src1 Mat, src2 Mat, dst *Mat) {
	// Do nothing
}

// Eigen calculates eigenvalues and eigenvectors of a symmetric matrix.
//
// For further details, please see:
// https://docs.opencv.org/master/d2/de8/group__core__array.html#ga9fa0d58657f60eaa6c71f6fbb40456e3
func Eigen(src Mat, eigenvalues *Mat, eigenvectors *Mat) bool {
	return false
}

// EigenNonSymmetric calculates eigenvalues and eigenvectors of a non-symmetric matrix (real eigenvalues only).
//
// For further details, please see:
// https://docs.opencv.org/master/d2/de8/group__core__array.html#gaf51987e03cac8d171fbd2b327cf966f6
func EigenNonSymmetric(src Mat, eigenvalues *Mat, eigenvectors *Mat) {
	// Do nothing
}

// PCABackProject reconstructs vectors from their PC projections.
//
// For further details, please see:
// https://docs.opencv.org/4.x/d2/de8/group__core__array.html#gab26049f30ee8e94f7d69d82c124faafc
func PCABackProject(data Mat, mean Mat, eigenvectors Mat, result *Mat) {
	// Do nothing
}

// PCACompute performs PCA.
//
// The computed eigenvalues are sorted from the largest to the smallest and the corresponding
// eigenvectors are stored as eigenvectors rows.
//
// Note: Calling with maxComponents == 0 (opencv default) will cause all components to be retained.
//
// For further details, please see:
// https://docs.opencv.org/4.x/d2/de8/group__core__array.html#ga27a565b31d820b05dcbcd47112176b6e
func PCACompute(src Mat, mean *Mat, eigenvectors *Mat, eigenvalues *Mat, maxComponents int) {
	// Do nothing
}

// PCAProject projects vector(s) to the principal component subspace.
//
// For further details, please see:
// https://docs.opencv.org/4.x/d2/de8/group__core__array.html#ga6b9fbc7b3a99ebfd441bbec0a6bc4f88
func PCAProject(data Mat, mean Mat, eigenvectors Mat, result *Mat) {
	// Do nothing
}

// PSNR computes the Peak Signal-to-Noise Ratio (PSNR) image quality metric.
//
// For further details, please see:
// https://docs.opencv.org/4.x/d2/de8/group__core__array.html#ga3119e3ea73010a6f810bb05aa36ac8d6
func PSNR(src1 Mat, src2 Mat) float64 {
	return 0
}

// SVBackSubst performs a singular value back substitution.
//
// For further details, please see:
// https://docs.opencv.org/4.x/d2/de8/group__core__array.html#gab4e620e6fc6c8a27bb2be3d50a840c0b
func SVBackSubst(w Mat, u Mat, vt Mat, rhs Mat, dst *Mat) {
	// Do nothing
}

// SVDecomp decomposes matrix and stores the results to user-provided matrices.
//
// For further details, please see:
// https://docs.opencv.org/4.x/d2/de8/group__core__array.html#gab477b5b7b39b370bb03e75b19d2d5109
func SVDecomp(src Mat, w *Mat, u *Mat, vt *Mat) {
	// Do nothing
}

// Exp calculates the exponent of every array element.
//
// For further details, please see:
// https://docs.opencv.org/master/d2/de8/group__core__array.html#ga3e10108e2162c338f1b848af619f39e5
func Exp(src Mat, dst *Mat) {
	// Do nothing
}

// ExtractChannel extracts a single channel from src (coi is 0-based index).
//
// For further details, please see:
// https://docs.opencv.org/master/d2/de8/group__core__array.html#gacc6158574aa1f0281878c955bcf35642
func ExtractChannel(src Mat, dst *Mat, coi int) {
	// Do nothing
}

// FindNonZero returns the list of locations of non-zero pixels.
//
// For further details, please see:
// https://docs.opencv.org/master/d2/de8/group__core__array.html#gaed7df59a3539b4cc0fe5c9c8d7586190
func FindNonZero(src Mat, idx *Mat) {
	// Do nothing
}

// Flip flips a 2D array around horizontal(0), vertical(1), or both axes(-1).
//
// For further details, please see:
// https://docs.opencv.org/master/d2/de8/group__core__array.html#gaca7be533e3dac7feb70fc60635adf441
func Flip(src Mat, dst *Mat, flipCode int) {
	// Do nothing
}

// Gemm performs generalized matrix multiplication.
//
// For further details, please see:
// https://docs.opencv.org/master/d2/de8/group__core__array.html#gacb6e64071dffe36434e1e7ee79e7cb35
func Gemm(src1, src2 Mat, alpha float64, src3 Mat, beta float64, dst *Mat, flags int) {
	// Do nothing
}

// GetOptimalDFTSize returns the optimal Discrete Fourier Transform (DFT) size
// for a given vector size.
//
// For further details, please see:
// https://docs.opencv.org/master/d2/de8/group__core__array.html#ga6577a2e59968936ae02eb2edde5de299
func GetOptimalDFTSize(vecsize int) int {
	return 0
}

// Hconcat applies horizontal concatenation to given matrices.
//
// For further details, please see:
// https://docs.opencv.org/master/d2/de8/group__core__array.html#gaab5ceee39e0580f879df645a872c6bf7
func Hconcat(src1, src2 Mat, dst *Mat) {
	// Do nothing
}

// Vconcat applies vertical concatenation to given matrices.
//
// For further details, please see:
// https://docs.opencv.org/master/d2/de8/group__core__array.html#gaab5ceee39e0580f879df645a872c6bf7
func Vconcat(src1, src2 Mat, dst *Mat) {
	// Do nothing
}

// RotateFlag for image rotation
//
// For further details please see:
// https://docs.opencv.org/master/d2/de8/group__core__array.html#ga6f45d55c0b1cc9d97f5353a7c8a7aac2
type RotateFlag int

const (
	// Rotate90Clockwise allows to rotate image 90 degrees clockwise
	Rotate90Clockwise RotateFlag = 0
	// Rotate180Clockwise allows to rotate image 180 degrees clockwise
	Rotate180Clockwise RotateFlag = 1
	// Rotate90CounterClockwise allows to rotate 270 degrees clockwise
	Rotate90CounterClockwise RotateFlag = 2
)

// Rotate rotates a 2D array in multiples of 90 degrees
//
// For further details, please see:
// https://docs.opencv.org/master/d2/de8/group__core__array.html#ga4ad01c0978b0ce64baa246811deeac24
func Rotate(src Mat, dst *Mat, code RotateFlag) {
	// Do nothing
}

// IDCT calculates the inverse Discrete Cosine Transform of a 1D or 2D array.
//
// For further details, please see:
// https://docs.opencv.org/master/d2/de8/group__core__array.html#ga77b168d84e564c50228b69730a227ef2
func IDCT(src Mat, dst *Mat, flags int) {
	// Do nothing
}

// IDFT calculates the inverse Discrete Fourier Transform of a 1D or 2D array.
//
// For further details, please see:
// https://docs.opencv.org/master/d2/de8/group__core__array.html#gaa708aa2d2e57a508f968eb0f69aa5ff1
func IDFT(src Mat, dst *Mat, flags, nonzeroRows int) {
	// Do nothing
}

// InRange checks if array elements lie between the elements of two Mat arrays.
//
// For further details, please see:
// https://docs.opencv.org/master/d2/de8/group__core__array.html#ga48af0ab51e36436c5d04340e036ce981
func InRange(src, lb, ub Mat, dst *Mat) {
	// Do nothing
}

// InRangeWithScalar checks if array elements lie between the elements of two Scalars
//
// For further details, please see:
// https://docs.opencv.org/master/d2/de8/group__core__array.html#ga48af0ab51e36436c5d04340e036ce981
func InRangeWithScalar(src Mat, lb, ub Scalar, dst *Mat) {
	// Do nothing
}

// InsertChannel inserts a single channel to dst (coi is 0-based index)
// (it replaces channel i with another in dst).
//
// For further details, please see:
// https://docs.opencv.org/master/d2/de8/group__core__array.html#ga1d4bd886d35b00ec0b764cb4ce6eb515
func InsertChannel(src Mat, dst *Mat, coi int) {
	// Do nothing
}

// Invert finds the inverse or pseudo-inverse of a matrix.
//
// For further details, please see:
// https://docs.opencv.org/master/d2/de8/group__core__array.html#gad278044679d4ecf20f7622cc151aaaa2
func Invert(src Mat, dst *Mat, flags SolveDecompositionFlags) float64 {
	return 0
}

// KMeansFlags for kmeans center selection
//
// For further details, please see:
// https://docs.opencv.org/master/d0/de1/group__core.html#ga276000efe55ee2756e0c471c7b270949
type KMeansFlags int

const (
	// KMeansRandomCenters selects random initial centers in each attempt.
	KMeansRandomCenters KMeansFlags = 0
	// KMeansPPCenters uses kmeans++ center initialization by Arthur and Vassilvitskii [Arthur2007].
	KMeansPPCenters KMeansFlags = 1
	// KMeansUseInitialLabels uses the user-supplied lables during the first (and possibly the only) attempt
	// instead of computing them from the initial centers. For the second and further attempts, use the random or semi-random     // centers. Use one of KMEANS_*_CENTERS flag to specify the exact method.
	KMeansUseInitialLabels KMeansFlags = 2
)

// KMeans finds centers of clusters and groups input samples around the clusters.
//
// For further details, please see:
// https://docs.opencv.org/master/d5/d38/group__core__cluster.html#ga9a34dc06c6ec9460e90860f15bcd2f88
func KMeans(data Mat, k int, bestLabels *Mat, criteria TermCriteria, attempts int, flags KMeansFlags, centers *Mat) float64 {
	return 0
}

// KMeansPoints finds centers of clusters and groups input samples around the clusters.
//
// For further details, please see:
// https://docs.opencv.org/master/d5/d38/group__core__cluster.html#ga9a34dc06c6ec9460e90860f15bcd2f88
func KMeansPoints(points PointVector, k int, bestLabels *Mat, criteria TermCriteria, attempts int, flags KMeansFlags, centers *Mat) float64 {
	return 0
}

// Log calculates the natural logarithm of every array element.
//
// For further details, please see:
// https://docs.opencv.org/master/d2/de8/group__core__array.html#ga937ecdce4679a77168730830a955bea7
func Log(src Mat, dst *Mat) {
	// Do nothing
}

// Magnitude calculates the magnitude of 2D vectors.
//
// For further details, please see:
// https://docs.opencv.org/master/d2/de8/group__core__array.html#ga6d3b097586bca4409873d64a90fe64c3
func Magnitude(x, y Mat, magnitude *Mat) {
	// Do nothing
}

// Mahalanobis calculates the Mahalanobis distance between two vectors.
//
// For further details, please see:
// https://docs.opencv.org/4.x/d2/de8/group__core__array.html#ga4493aee129179459cbfc6064f051aa7d
func Mahalanobis(v1, v2, icovar Mat) float64 {
	return 0
}

// MulTransposed calculates the product of a matrix and its transposition.
//
// For further details, please see:
// https://docs.opencv.org/4.x/d2/de8/group__core__array.html#gadc4e49f8f7a155044e3be1b9e3b270ab
func MulTransposed(src Mat, dest *Mat, ata bool) {
	// Do nothing
}

// Max calculates per-element maximum of two arrays or an array and a scalar.
//
// For further details, please see:
// https://docs.opencv.org/master/d2/de8/group__core__array.html#gacc40fa15eac0fb83f8ca70b7cc0b588d
func Max(src1, src2 Mat, dst *Mat) {
	// Do nothing
}

// MeanStdDev calculates a mean and standard deviation of array elements.
//
// For further details, please see:
// https://docs.opencv.org/master/d2/de8/group__core__array.html#ga846c858f4004d59493d7c6a4354b301d
func MeanStdDev(src Mat, dst *Mat, dstStdDev *Mat) {
	// Do nothing
}

// Merge creates one multi-channel array out of several single-channel ones.
//
// For further details, please see:
// https://docs.opencv.org/master/d2/de8/group__core__array.html#ga7d7b4d6c6ee504b30a20b1680029c7b4
func Merge(mv []Mat, dst *Mat) {
	// Do nothing
}

// Min calculates per-element minimum of two arrays or an array and a scalar.
//
// For further details, please see:
// https://docs.opencv.org/master/d2/de8/group__core__array.html#ga9af368f182ee76d0463d0d8d5330b764
func Min(src1, src2 Mat, dst *Mat) {
	// Do nothing
}

// MinMaxIdx finds the global minimum and maximum in an array.
//
// For further details, please see:
// https://docs.opencv.org/master/d2/de8/group__core__array.html#ga7622c466c628a75d9ed008b42250a73f
func MinMaxIdx(input Mat) (minVal, maxVal float32, minIdx, maxIdx int) {
	return 0, 0, 0, 0
}

// MinMaxLoc finds the global minimum and maximum in an array.
//
// For further details, please see:
// https://docs.opencv.org/trunk/d2/de8/group__core__array.html#gab473bf2eb6d14ff97e89b355dac20707
func MinMaxLoc(input Mat) (minVal, maxVal float32, minLoc, maxLoc image.Point) {
	return 0, 0, image.Point{}, image.Point{}
}

// MinMaxLocWithMask finds the global minimum and maximum in an array with a mask used to select a sub-array.
//
// For further details, please see:
// https://docs.opencv.org/4.x/d2/de8/group__core__array.html#gab473bf2eb6d14ff97e89b355dac20707
func MinMaxLocWithMask(input, mask Mat) (minVal, maxVal float32, minLoc, maxLoc image.Point) {
	return 0, 0, image.Point{}, image.Point{}
}

// Copies specified channels from input arrays to the specified channels of output arrays.
//
// For further details, please see:
// https://docs.opencv.org/master/d2/de8/group__core__array.html#ga51d768c270a1cdd3497255017c4504be
func MixChannels(src []Mat, dst []Mat, fromTo []int) {
	// Do nothing
}

// Mulspectrums performs the per-element multiplication of two Fourier spectrums.
//
// For further details, please see:
// https://docs.opencv.org/master/d2/de8/group__core__array.html#ga3ab38646463c59bf0ce962a9d51db64f
func MulSpectrums(a Mat, b Mat, dst *Mat, flags DftFlags) {
	// Do nothing
}

// Multiply calculates the per-element scaled product of two arrays.
// Both input arrays must be of the same size and the same type.
//
// For further details, please see:
// https://docs.opencv.org/master/d2/de8/group__core__array.html#ga979d898a58d7f61c53003e162e7ad89f
func Multiply(src1 Mat, src2 Mat, dst *Mat) {
	// Do nothing
}

// MultiplyWithParams calculates the per-element scaled product of two arrays.
// Both input arrays must be of the same size and the same type.
//
// For further details, please see:
// https://docs.opencv.org/master/d2/de8/group__core__array.html#ga979d898a58d7f61c53003e162e7ad89f
func MultiplyWithParams(src1 Mat, src2 Mat, dst *Mat, scale float64, dtype MatType) {
	// Do nothing
}

// NormType for normalization operations.
//
// For further details, please see:
// https://docs.opencv.org/master/d2/de8/group__core__array.html#gad12cefbcb5291cf958a85b4b67b6149f
type NormType int

const (
	// NormInf indicates use infinite normalization.
	NormInf NormType = 1

	// NormL1 indicates use L1 normalization.
	NormL1 NormType = 2

	// NormL2 indicates use L2 normalization.
	NormL2 NormType = 4

	// NormL2Sqr indicates use L2 squared normalization.
	NormL2Sqr NormType = 5

	// NormHamming indicates use Hamming normalization.
	NormHamming NormType = 6

	// NormHamming2 indicates use Hamming 2-bit normalization.
	NormHamming2 NormType = 7

	// NormTypeMask indicates use type mask for normalization.
	NormTypeMask NormType = 7

	// NormRelative indicates use relative normalization.
	NormRelative NormType = 8

	// NormMinMax indicates use min/max normalization.
	NormMinMax NormType = 32
)

// Normalize normalizes the norm or value range of an array.
//
// For further details, please see:
// https://docs.opencv.org/master/d2/de8/group__core__array.html#ga87eef7ee3970f86906d69a92cbf064bd
func Normalize(src Mat, dst *Mat, alpha float64, beta float64, typ NormType) {
	// Do nothing
}

// Norm calculates the absolute norm of an array.
//
// For further details, please see:
// https://docs.opencv.org/master/d2/de8/group__core__array.html#ga7c331fb8dd951707e184ef4e3f21dd33
func Norm(src1 Mat, normType NormType) float64 {
	return 0
}

// Norm calculates the absolute difference/relative norm of two arrays.
//
// For further details, please see:
// https://docs.opencv.org/master/d2/de8/group__core__array.html#ga7c331fb8dd951707e184ef4e3f21dd33
func NormWithMats(src1 Mat, src2 Mat, normType NormType) float64 {
	return 0
}

// PerspectiveTransform performs the perspective matrix transformation of vectors.
//
// For further details, please see:
// https://docs.opencv.org/master/d2/de8/group__core__array.html#gad327659ac03e5fd6894b90025e6900a7
func PerspectiveTransform(src Mat, dst *Mat, tm Mat) {
	// Do nothing
}

// TermCriteriaType for TermCriteria.
//
// For further details, please see:
// https://docs.opencv.org/master/d9/d5d/classcv_1_1TermCriteria.html#a56fecdc291ccaba8aad27d67ccf72c57
type TermCriteriaType int

const (
	// Count is the maximum number of iterations or elements to compute.
	Count TermCriteriaType = 1

	// MaxIter is the maximum number of iterations or elements to compute.
	MaxIter TermCriteriaType = 1

	// EPS is the desired accuracy or change in parameters at which the
	// iterative algorithm stops.
	EPS TermCriteriaType = 2
)

type SolveDecompositionFlags int

const (
	// Gaussian elimination with the optimal pivot element chosen.
	SolveDecompositionLu SolveDecompositionFlags = 0

	// Singular value decomposition (SVD) method. The system can be over-defined and/or the matrix src1 can be singular.
	SolveDecompositionSvd SolveDecompositionFlags = 1

	// Eigenvalue decomposition. The matrix src1 must be symmetrical.
	SolveDecompositionEing SolveDecompositionFlags = 2

	// Cholesky LL^T factorization. The matrix src1 must be symmetrical and positively defined.
	SolveDecompositionCholesky SolveDecompositionFlags = 3

	// QR factorization. The system can be over-defined and/or the matrix src1 can be singular.
	SolveDecompositionQr SolveDecompositionFlags = 4

	// While all the previous flags are mutually exclusive, this flag can be used together with any of the previous.
	// It means that the normal equations 𝚜𝚛𝚌𝟷^T⋅𝚜𝚛𝚌𝟷⋅𝚍𝚜𝚝=𝚜𝚛𝚌𝟷^T𝚜𝚛𝚌𝟸 are solved instead of the original system
	// 𝚜𝚛𝚌𝟷⋅𝚍𝚜𝚝=𝚜𝚛𝚌𝟸.
	SolveDecompositionNormal SolveDecompositionFlags = 5
)

// Solve solves one or more linear systems or least-squares problems.
//
// For further details, please see:
// https://docs.opencv.org/master/d2/de8/group__core__array.html#ga12b43690dbd31fed96f213eefead2373
func Solve(src1 Mat, src2 Mat, dst *Mat, flags SolveDecompositionFlags) bool {
	return false
}

// SolveCubic finds the real roots of a cubic equation.
//
// For further details, please see:
// https://docs.opencv.org/master/d2/de8/group__core__array.html#ga1c3b0b925b085b6e96931ee309e6a1da
func SolveCubic(coeffs Mat, roots *Mat) int {
	return 0
}

// SolvePoly finds the real or complex roots of a polynomial equation.
//
// For further details, please see:
// https://docs.opencv.org/master/d2/de8/group__core__array.html#gac2f5e953016fabcdf793d762f4ec5dce
func SolvePoly(coeffs Mat, roots *Mat, maxIters int) float64 {
	return 0
}

type ReduceTypes int

const (
	// The output is the sum of all rows/columns of the matrix.
	ReduceSum ReduceTypes = 0

	// The output is the mean vector of all rows/columns of the matrix.
	ReduceAvg ReduceTypes = 1

	// The output is the maximum (column/row-wise) of all rows/columns of the matrix.
	ReduceMax ReduceTypes = 2

	// The output is the minimum (column/row-wise) of all rows/columns of the matrix.
	ReduceMin ReduceTypes = 3
)

// Reduce reduces a matrix to a vector.
//
// For further details, please see:
// https://docs.opencv.org/master/d2/de8/group__core__array.html#ga4b78072a303f29d9031d56e5638da78e
func Reduce(src Mat, dst *Mat, dim int, rType ReduceTypes, dType MatType) {
	// Do nothing
}

// Finds indices of max elements along provided axis.
//
// For further details, please see:
// https://docs.opencv.org/master/d2/de8/group__core__array.html#gaa87ea34d99bcc5bf9695048355163da0
func ReduceArgMax(src Mat, dst *Mat, axis int, lastIndex bool) {
	// Do nothing
}

// Finds indices of min elements along provided axis.
//
// For further details, please see:
// https://docs.opencv.org/master/d2/de8/group__core__array.html#gaeecd548276bfb91b938989e66b722088
func ReduceArgMin(src Mat, dst *Mat, axis int, lastIndex bool) {
	// Do nothing
}

// Repeat fills the output array with repeated copies of the input array.
//
// For further details, please see:
// https://docs.opencv.org/master/d2/de8/group__core__array.html#ga496c3860f3ac44c40b48811333cfda2d
func Repeat(src Mat, nY int, nX int, dst *Mat) {
	// Do nothing
}

// Calculates the sum of a scaled array and another array.
//
// For further details, please see:
// https://docs.opencv.org/master/d2/de8/group__core__array.html#ga9e0845db4135f55dcf20227402f00d98
func ScaleAdd(src1 Mat, alpha float64, src2 Mat, dst *Mat) {
	// Do nothing
}

// SetIdentity initializes a scaled identity matrix.
// For further details, please see:
//
//	https://docs.opencv.org/master/d2/de8/group__core__array.html#ga388d7575224a4a277ceb98ccaa327c99
func SetIdentity(src Mat, scalar float64) {
	// Do nothing
}

type SortFlags int

const (
	// Each matrix row is sorted independently
	SortEveryRow SortFlags = 0

	// Each matrix column is sorted independently; this flag and the previous one are mutually exclusive.
	SortEveryColumn SortFlags = 1

	// Each matrix row is sorted in the ascending order.
	SortAscending SortFlags = 0

	// Each matrix row is sorted in the descending order; this flag and the previous one are also mutually exclusive.
	SortDescending SortFlags = 16
)

// Sort sorts each row or each column of a matrix.
//
// For further details, please see:
// https://docs.opencv.org/master/d2/de8/group__core__array.html#ga45dd56da289494ce874be2324856898f
func Sort(src Mat, dst *Mat, flags SortFlags) {
	// Do nothing
}

// SortIdx sorts each row or each column of a matrix.
// Instead of reordering the elements themselves, it stores the indices of sorted elements in the output array
//
// For further details, please see:
// https://docs.opencv.org/master/d2/de8/group__core__array.html#gadf35157cbf97f3cb85a545380e383506
func SortIdx(src Mat, dst *Mat, flags SortFlags) {
	// Do nothing
}

// Split creates an array of single channel images from a multi-channel image
// Created images should be closed manualy to avoid memory leaks.
//
// For further details, please see:
// https://docs.opencv.org/master/d2/de8/group__core__array.html#ga0547c7fed86152d7e9d0096029c8518a
func Split(src Mat) (mv []Mat) {
	return nil
}

// Subtract calculates the per-element subtraction of two arrays or an array and a scalar.
//
// For further details, please see:
// https://docs.opencv.org/master/d2/de8/group__core__array.html#gaa0f00d98b4b5edeaeb7b8333b2de353b
func Subtract(src1 Mat, src2 Mat, dst *Mat) {
	// Do nothing
}

// Trace returns the trace of a matrix.
//
// For further details, please see:
// https://docs.opencv.org/master/d2/de8/group__core__array.html#ga3419ac19c7dcd2be4bd552a23e147dd8
func Trace(src Mat) Scalar {
	return Scalar{}
}

// Transform performs the matrix transformation of every array element.
//
// For further details, please see:
// https://docs.opencv.org/master/d2/de8/group__core__array.html#ga393164aa54bb9169ce0a8cc44e08ff22
func Transform(src Mat, dst *Mat, tm Mat) {
	// Do nothing
}

// Transpose transposes a matrix.
//
// For further details, please see:
// https://docs.opencv.org/master/d2/de8/group__core__array.html#ga46630ed6c0ea6254a35f447289bd7404
func Transpose(src Mat, dst *Mat) {
	// Do nothing
}

// TransposeND transpose for n-dimensional matrices.
//
// For further details, please see:
// https://docs.opencv.org/master/d2/de8/group__core__array.html#gab1b1274b4a563be34cdfa55b8919a4ec
func TransposeND(src Mat, order []int, dst *Mat) {
	// Do nothing
}

// Pow raises every array element to a power.
//
// For further details, please see:
// https://docs.opencv.org/master/d2/de8/group__core__array.html#gaf0d056b5bd1dc92500d6f6cf6bac41ef
func Pow(src Mat, power float64, dst *Mat) {
	// Do nothing
}

// PolatToCart calculates x and y coordinates of 2D vectors from their magnitude and angle.
//
// For further details, please see:
// https://docs.opencv.org/master/d2/de8/group__core__array.html#ga581ff9d44201de2dd1b40a50db93d665
func PolarToCart(magnitude Mat, degree Mat, x *Mat, y *Mat, angleInDegrees bool) {
	// Do nothing
}

// Phase calculates the rotation angle of 2D vectors.
//
// For further details, please see:
// https://docs.opencv.org/master/d2/de8/group__core__array.html#ga9db9ca9b4d81c3bde5677b8f64dc0137
func Phase(x, y Mat, angle *Mat, angleInDegrees bool) {
	// Do nothing
}

// TermCriteria is the criteria for iterative algorithms.
//
// For further details, please see:
// https://docs.opencv.org/master/d9/d5d/classcv_1_1TermCriteria.html
type TermCriteria struct {
	// Do nothing
}

// NewTermCriteria returns a new TermCriteria.
func NewTermCriteria(typ TermCriteriaType, maxCount int, epsilon float64) TermCriteria {
	return TermCriteria{}
}

// Scalar is a 4-element vector widely used in OpenCV to pass pixel values.
//
// For further details, please see:
// http://docs.opencv.org/master/d1/da0/classcv_1_1Scalar__.html
type Scalar struct {
	Val1 float64
	Val2 float64
	Val3 float64
	Val4 float64
}

// NewScalar returns a new Scalar. These are usually colors typically being in BGR order.
func NewScalar(v1 float64, v2 float64, v3 float64, v4 float64) Scalar {
	return Scalar{}
}

// KeyPoint is data structure for salient point detectors.
//
// For further details, please see:
// https://docs.opencv.org/master/d2/d29/classcv_1_1KeyPoint.html
type KeyPoint struct {
	X, Y                  float64
	Size, Angle, Response float64
	Octave, ClassID       int
}

// DMatch is data structure for matching keypoint descriptors.
//
// For further details, please see:
// https://docs.opencv.org/master/d4/de0/classcv_1_1DMatch.html#a546ddb9a87898f06e510e015a6de596e
type DMatch struct {
	QueryIdx int
	TrainIdx int
	ImgIdx   int
	Distance float64
}

// Vecb is a generic vector of bytes.
type Vecb []uint8

// GetVecbAt returns a vector of bytes. Its size corresponds to the number
// of channels of the Mat.
func (m *Mat) GetVecbAt(row int, col int) Vecb {
	return nil
}

// Vecf is a generic vector of floats.
type Vecf []float32

// GetVecfAt returns a vector of floats. Its size corresponds to the number of
// channels of the Mat.
func (m *Mat) GetVecfAt(row int, col int) Vecf {
	return nil
}

// Vecd is a generic vector of float64/doubles.
type Vecd []float64

// GetVecdAt returns a vector of float64s. Its size corresponds to the number
// of channels of the Mat.
func (m *Mat) GetVecdAt(row int, col int) Vecd {
	return nil
}

// Veci is a generic vector of integers.
type Veci []int32

// GetVeciAt returns a vector of integers. Its size corresponds to the number
// of channels of the Mat.
func (m *Mat) GetVeciAt(row int, col int) Veci {
	return nil
}

// PointVector is a wrapper around a std::vector< cv::Point >*
// This is needed anytime that you need to pass or receive a collection of points.
type PointVector struct {
	// Do nothing
}

// NewPointVector returns a new empty PointVector.
func NewPointVector() PointVector {
	return PointVector{}
}

// NewPointVectorFromPoints returns a new PointVector that has been
// initialized to a slice of image.Point.
func NewPointVectorFromPoints(pts []image.Point) PointVector {
	return PointVector{}
}

// NewPointVectorFromMat returns a new PointVector that has been
// wrapped around a Mat of type CV_32SC2 with a single columm.
func NewPointVectorFromMat(mat Mat) PointVector {
	return PointVector{}
}

// IsNil checks the CGo pointer in the PointVector.
func (pv PointVector) IsNil() bool {
	return false
}

// Size returns how many Point are in the PointVector.
func (pv PointVector) Size() int {
	return 0
}

// At returns the image.Point
func (pv PointVector) At(idx int) image.Point {
	return image.Point{}
}

// Append appends an image.Point at end of the PointVector.
func (pv PointVector) Append(point image.Point) {
	// Do nothing
}

// ToPoints returns a slice of image.Point for the data in this PointVector.
func (pv PointVector) ToPoints() []image.Point {
	return nil
}

// Close closes and frees memory for this PointVector.
func (pv PointVector) Close() {
	// Do nothing
}

// PointsVector is a wrapper around a std::vector< std::vector< cv::Point > >*
type PointsVector struct {
	// Do nothing
}

// NewPointsVector returns a new empty PointsVector.
func NewPointsVector() PointsVector {
	return PointsVector{}
}

// NewPointsVectorFromPoints returns a new PointsVector that has been
// initialized to a slice of slices of image.Point.
func NewPointsVectorFromPoints(pts [][]image.Point) PointsVector {
	return PointsVector{}
}

func (pvs PointsVector) P() C.PointsVector {
	return nil
}

// ToPoints returns a slice of slices of image.Point for the data in this PointsVector.
func (pvs PointsVector) ToPoints() [][]image.Point {
	return nil
}

// IsNil checks the CGo pointer in the PointsVector.
func (pvs PointsVector) IsNil() bool {
	return false
}

// Size returns how many vectors of Points are in the PointsVector.
func (pvs PointsVector) Size() int {
	return 0
}

// At returns the PointVector at that index of the PointsVector.
func (pvs PointsVector) At(idx int) PointVector {
	return PointVector{}
}

// Append appends a PointVector at end of the PointsVector.
func (pvs PointsVector) Append(pv PointVector) {
	// Do nothing
}

// Close closes and frees memory for this PointsVector.
func (pvs PointsVector) Close() {
	// Do nothing
}

// Point2fVector is a wrapper around a std::vector< cv::Point2f >*
// This is needed anytime that you need to pass or receive a collection of points.
type Point2fVector struct {
	// Do nothing
}

// NewPoint2fVector returns a new empty Point2fVector.
func NewPoint2fVector() Point2fVector {
	return Point2fVector{}
}

// NewPoint2fVectorFromPoints returns a new Point2fVector that has been
// initialized to a slice of image.Point.
func NewPoint2fVectorFromPoints(pts []Point2f) Point2fVector {
	return Point2fVector{}
}

// NewPoint2fVectorFromMat returns a new Point2fVector that has been
// wrapped around a Mat of type CV_32FC2 with a single columm.
func NewPoint2fVectorFromMat(mat Mat) Point2fVector {
	return Point2fVector{}
}

// IsNil checks the CGo pointer in the Point2fVector.
func (pfv Point2fVector) IsNil() bool {
	return false
}

// Size returns how many Point are in the PointVector.
func (pfv Point2fVector) Size() int {
	return 0
}

// At returns the image.Point
func (pfv Point2fVector) At(idx int) Point2f {
	return Point2f{}
}

// ToPoints returns a slice of image.Point for the data in this PointVector.
func (pfv Point2fVector) ToPoints() []Point2f {
	return nil
}

// Close closes and frees memory for this Point2fVector.
func (pfv Point2fVector) Close() {
	// Do nothing
}

// GetTickCount returns the number of ticks.
//
// For further details, please see:
// https://docs.opencv.org/master/db/de0/group__core__utils.html#gae73f58000611a1af25dd36d496bf4487
func GetTickCount() float64 {
	return 0
}

// GetTickFrequency returns the number of ticks per second.
//
// For further details, please see:
// https://docs.opencv.org/master/db/de0/group__core__utils.html#ga705441a9ef01f47acdc55d87fbe5090c
func GetTickFrequency() float64 {
	return 0
}

// RowRange creates a matrix header for the specified row span.
//
// For further details, please see:
// https://docs.opencv.org/master/d3/d63/classcv_1_1Mat.html#aa6542193430356ad631a9beabc624107
func (m *Mat) RowRange(start, end int) Mat {
	return Mat{}
}

// ColRange creates a matrix header for the specified column span.
//
// For further details, please see:
// https://docs.opencv.org/master/d3/d63/classcv_1_1Mat.html#aadc8f9210fe4dec50513746c246fa8d9
func (m *Mat) ColRange(start, end int) Mat {
	return Mat{}
}

// RNG Random Number Generator.
// It encapsulates the state (currently, a 64-bit integer) and
// has methods to return scalar random values and to fill arrays
// with random values
//
// For further details, please see:
// https://docs.opencv.org/master/d1/dd6/classcv_1_1RNG.html
type RNG struct {
	// Do nothing
}

type RNGDistType int

const (
	// Uniform distribution
	RNGDistUniform RNGDistType = 0
	// Normal distribution
	RNGDistNormal RNGDistType = 1
)

// TheRNG Returns the default random number generator.
//
// For further details, please see:
// https://docs.opencv.org/master/d2/de8/group__core__array.html#ga75843061d150ad6564b5447e38e57722
func TheRNG() RNG {
	return RNG{}
}

// TheRNG Sets state of default random number generator.
//
// For further details, please see:
// https://docs.opencv.org/master/d2/de8/group__core__array.html#ga757e657c037410d9e19e819569e7de0f
func SetRNGSeed(seed int) {
	// Do nothing
}

// Fill Fills arrays with random numbers.
//
// For further details, please see:
// https://docs.opencv.org/master/d1/dd6/classcv_1_1RNG.html#ad26f2b09d9868cf108e84c9814aa682d
func (r *RNG) Fill(mat *Mat, distType RNGDistType, a, b float64, saturateRange bool) {
	// Do nothing
}

// Gaussian Returns the next random number sampled from
// the Gaussian distribution.
//
// For further details, please see:
// https://docs.opencv.org/master/d1/dd6/classcv_1_1RNG.html#a8df8ce4dc7d15916cee743e5a884639d
func (r *RNG) Gaussian(sigma float64) float64 {
	return 0
}

// Next The method updates the state using the MWC algorithm
// and returns the next 32-bit random number.
//
// For further details, please see:
// https://docs.opencv.org/master/d1/dd6/classcv_1_1RNG.html#a8df8ce4dc7d15916cee743e5a884639d
func (r *RNG) Next() uint {
	return 0
}

// RandN Fills the array with normally distributed random numbers.
//
// For further details, please see:
// https://docs.opencv.org/master/d2/de8/group__core__array.html#gaeff1f61e972d133a04ce3a5f81cf6808
func RandN(mat *Mat, mean, stddev Scalar) {
	// Do nothing
}

// RandShuffle Shuffles the array elements randomly.
//
// For further details, please see:
// https://docs.opencv.org/master/d2/de8/group__core__array.html#ga6a789c8a5cb56c6dd62506179808f763
func RandShuffle(mat *Mat) {
	// Do nothing
}

// RandShuffleWithParams Shuffles the array elements randomly.
//
// For further details, please see:
// https://docs.opencv.org/master/d2/de8/group__core__array.html#ga6a789c8a5cb56c6dd62506179808f763
func RandShuffleWithParams(mat *Mat, iterFactor float64, rng RNG) {
	// Do nothing
}

// RandU Generates a single uniformly-distributed random
// number or an array of random numbers.
//
// For further details, please see:
// https://docs.opencv.org/master/d2/de8/group__core__array.html#ga1ba1026dca0807b27057ba6a49d258c0
func RandU(mat *Mat, low, high Scalar) {
	// Do nothing
}

type NativeByteBuffer struct {
	// Do nothing
}

// GetBytes returns slice of bytes backed by native buffer
func (buffer *NativeByteBuffer) GetBytes() []byte {
	return nil
}

// Len - returns length in bytes of underlying buffer
func (buffer *NativeByteBuffer) Len() int {
	return 0
}

// Close the buffer releasing all its resources
func (buffer *NativeByteBuffer) Close() {
	// Do nothing
}

// Points2fVector is a wrapper around a std::vector< std::vector< cv::Point2f > >*
type Points2fVector struct {
	// Do nothing
}

// NewPoints2fVector returns a new empty Points2fVector.
func NewPoints2fVector() Points2fVector {
	return Points2fVector{}
}

// NewPoints2fVectorFromPoints returns a new Points2fVector that has been
// initialized to a slice of slices of Point2f.
func NewPoints2fVectorFromPoints(pts [][]Point2f) Points2fVector {
	return Points2fVector{}
}

func (pvs Points2fVector) P() C.Points2fVector {
	return nil
}

// ToPoints returns a slice of slices of Point2f for the data in this Points2fVector.
func (pvs Points2fVector) ToPoints() [][]Point2f {
	return nil
}

// IsNil checks the CGo pointer in the Points2fVector.
func (pvs Points2fVector) IsNil() bool {
	return false
}

// Size returns how many vectors of Points are in the Points2fVector.
func (pvs Points2fVector) Size() int {
	return 0
}

// At returns the Point2fVector at that index of the Points2fVector.
func (pvs Points2fVector) At(idx int) Point2fVector {
	return Point2fVector{}
}

// Append appends a Point2fVector at end of the Points2fVector.
func (pvs Points2fVector) Append(pv Point2fVector) {
	// Do nothing
}

// Close closes and frees memory for this Points2fVector.
func (pvs Points2fVector) Close() {
	// Do nothing
}

type Point3f struct {
	X float32
	Y float32
	Z float32
}

func NewPoint3f(x, y, z float32) Point3f {
	return Point3f{}
}

// Point3fVector is a wrapper around a std::vector< cv::Point3f >*
type Point3fVector struct {
	// Do nothing
}

// NewPoint3fVector returns a new empty Point3fVector.
func NewPoint3fVector() Point3fVector {
	return Point3fVector{}
}

// NewPoint3fVectorFromPoints returns a new Point3fVector that has been
// initialized to a slice of image.Point.
func NewPoint3fVectorFromPoints(pts []Point3f) Point3fVector {
	return Point3fVector{}
}

// NewPoint3fVectorFromMat returns a new Point3fVector that has been
// wrapped around a Mat of type CV_32FC3 with a single columm.
func NewPoint3fVectorFromMat(mat Mat) Point3fVector {
	return Point3fVector{}
}

// IsNil checks the CGo pointer in the Point3fVector.
func (pfv Point3fVector) IsNil() bool {
	return false
}

// Size returns how many Point are in the Point3fVector.
func (pfv Point3fVector) Size() int {
	return 0
}

// At returns the Point3f
func (pfv Point3fVector) At(idx int) Point3f {
	return Point3f{}
}

func (pfv Point3fVector) Append(point Point3f) {
	// Do nothing
}

// ToPoints returns a slice of Point3f for the data in this Point3fVector.
func (pfv Point3fVector) ToPoints() []Point3f {
	return nil
}

// Close closes and frees memory for this Point3fVector.
func (pfv Point3fVector) Close() {
	// Do nothing
}

// Points3fVector is a wrapper around a std::vector< std::vector< cv::Point3f > >*
type Points3fVector struct {
	// Do nothing
}

// NewPoints3fVector returns a new empty Points3fVector.
func NewPoints3fVector() Points3fVector {
	return Points3fVector{}
}

// NewPoints3fVectorFromPoints returns a new Points3fVector that has been
// initialized to a slice of slices of Point3f.
func NewPoints3fVectorFromPoints(pts [][]Point3f) Points3fVector {
	return Points3fVector{}
}

// ToPoints returns a slice of slices of Point3f for the data in this Points3fVector.
func (pvs Points3fVector) ToPoints() [][]Point3f {
	return nil
}

// IsNil checks the CGo pointer in the Points3fVector.
func (pvs Points3fVector) IsNil() bool {
	return false
}

// Size returns how many vectors of Points are in the Points3fVector.
func (pvs Points3fVector) Size() int {
	return 0
}

// At returns the Point3fVector at that index of the Points3fVector.
func (pvs Points3fVector) At(idx int) Point3fVector {
	return Point3fVector{}
}

// Append appends a Point3fVector at end of the Points3fVector.
func (pvs Points3fVector) Append(pv Point3fVector) {
	// Do nothing
}

// Close closes and frees memory for this Points3fVector.
func (pvs Points3fVector) Close() {
	// Do nothing
}

// Set the number of threads for OpenCV.
func SetNumThreads(n int) {
	// Do nothing
}

// Get the number of threads for OpenCV.
func GetNumThreads() int {
	return 0
}

// NewRotatedRect creates [RotatedRect] (i.e. not up-right) rectangle on a plane.
//
// For further information, see:
// https://docs.opencv.org/4.x/db/dd6/classcv_1_1RotatedRect.html#aba20dfc8444fff72bd820b616f0297ee
func NewRotatedRect(center image.Point, width int, height int, angle float64) RotatedRect {
	return RotatedRect{}
}

// NewRotatedRect2f creates [RotatedRect2f] (i.e. not up-right) rectangle on a plane.
//
// For further information, see:
// https://docs.opencv.org/4.x/db/dd6/classcv_1_1RotatedRect.html#aba20dfc8444fff72bd820b616f0297ee
func NewRotatedRect2f(center Point2f, width float32, height float32, angle float64) RotatedRect2f {
	return RotatedRect2f{}
}

// LastExceptionError returns an error if there was an exception in the OpenCV library.
func LastExceptionError() error {
	return nil
}

// GetLastExceptionMessage returns the last exception message from the OpenCV library.
func GetLastExceptionMessage() string {
	return ""
}

// ClearLastException clears the last exception from the OpenCV library.
func ClearLastException() {
	// Do nothing
}

// GetLastException returns the last exception code from the OpenCV library.
func GetLastException() int {
	return 0
}

// Converts a OpenCVResult struct to an error.
func OpenCVResult(result C.OpenCVResult) error {
	return nil
}
