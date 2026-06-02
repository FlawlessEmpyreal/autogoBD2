package opencv

import (
	"image"
	"image/color"
)

// ArcLength calculates a contour perimeter or a curve length.
//
// For further details, please see:
//
// https://docs.opencv.org/master/d3/dc0/group__imgproc__shape.html#ga8d26483c636be6b35c3ec6335798a47c
func ArcLength(curve PointVector, isClosed bool) float64 {
	return 0
}

// ApproxPolyDP approximates a polygonal curve(s) with the specified precision.
//
// For further details, please see:
//
// https://docs.opencv.org/master/d3/dc0/group__imgproc__shape.html#ga0012a5fdaea70b8a9970165d98722b4c
func ApproxPolyDP(curve PointVector, epsilon float64, closed bool) PointVector {
	return PointVector{}
}

// ConvexHull finds the convex hull of a point set.
//
// For further details, please see:
// https://docs.opencv.org/master/d3/dc0/group__imgproc__shape.html#ga014b28e56cb8854c0de4a211cb2be656
func ConvexHull(points PointVector, hull *Mat, clockwise bool, returnPoints bool) {
	// Do nothing
}

// ConvexityDefects finds the convexity defects of a contour.
//
// For further details, please see:
// https://docs.opencv.org/master/d3/dc0/group__imgproc__shape.html#gada4437098113fd8683c932e0567f47ba
func ConvexityDefects(contour PointVector, hull Mat, result *Mat) {
	// Do nothing
}

// CvtColor converts an image from one color space to another.
// It converts the src Mat image to the dst Mat using the
// code param containing the desired ColorConversionCode color space.
//
// For further details, please see:
// http://docs.opencv.org/master/d7/d1b/group__imgproc__misc.html#ga4e0972be5de079fed4e3a10e24ef5ef0
func CvtColor(src Mat, dst *Mat, code ColorConversionCode) {
	// Do nothing
}

// Demosaicing converts an image from Bayer pattern to RGB or grayscale.
// It converts the src Mat image to the dst Mat using the
// code param containing the desired ColorConversionCode color space.
//
// For further details, please see:
// https://docs.opencv.org/master/d7/d1b/group__imgproc__color__conversions.html#ga57261f12fccf872a2b2d66daf29d5bd0
func Demosaicing(src Mat, dst *Mat, code ColorConversionCode) {
	// Do nothing
}

// EqualizeHist normalizes the brightness and increases the contrast of the image.
//
// For further details, please see:
// https://docs.opencv.org/master/d6/dc7/group__imgproc__hist.html#ga7e54091f0c937d49bf84152a16f76d6e
func EqualizeHist(src Mat, dst *Mat) {
	// Do nothing
}

// CalcHist Calculates a histogram of a set of images
//
// For futher details, please see:
// https://docs.opencv.org/master/d6/dc7/group__imgproc__hist.html#ga6ca1876785483836f72a77ced8ea759a
func CalcHist(src []Mat, channels []int, mask Mat, hist *Mat, size []int, ranges []float64, acc bool) {
	// Do nothing
}

// CalcBackProject calculates the back projection of a histogram.
//
// For futher details, please see:
// https://docs.opencv.org/3.4/d6/dc7/group__imgproc__hist.html#ga3a0af640716b456c3d14af8aee12e3ca
func CalcBackProject(src []Mat, channels []int, hist Mat, backProject *Mat, ranges []float64, uniform bool) {
	// Do nothing
}

// HistCompMethod is the method for Histogram comparison
// For more information, see https://docs.opencv.org/master/d6/dc7/group__imgproc__hist.html#ga994f53817d621e2e4228fc646342d386
type HistCompMethod int

const (
	// HistCmpCorrel calculates the Correlation
	HistCmpCorrel HistCompMethod = 0

	// HistCmpChiSqr calculates the Chi-Square
	HistCmpChiSqr HistCompMethod = 1

	// HistCmpIntersect calculates the Intersection
	HistCmpIntersect HistCompMethod = 2

	// HistCmpBhattacharya applies the HistCmpBhattacharya by calculating the Bhattacharya distance.
	HistCmpBhattacharya HistCompMethod = 3

	// HistCmpHellinger applies the HistCmpBhattacharya comparison. It is a synonym to HistCmpBhattacharya.
	HistCmpHellinger = HistCmpBhattacharya

	// HistCmpChiSqrAlt applies the Alternative Chi-Square (regularly used for texture comparsion).
	HistCmpChiSqrAlt HistCompMethod = 4

	// HistCmpKlDiv applies the Kullback-Liebler divergence comparison.
	HistCmpKlDiv HistCompMethod = 5
)

// CompareHist Compares two histograms.
//
// For further details, please see:
// https://docs.opencv.org/master/d6/dc7/group__imgproc__hist.html#gaf4190090efa5c47cb367cf97a9a519bd
func CompareHist(hist1 Mat, hist2 Mat, method HistCompMethod) float32 {
	return 0
}

// EMD Computes the "minimal work" distance between two weighted point configurations.
//
// For further details, please see:
// https://docs.opencv.org/4.x/d6/dc7/group__imgproc__hist.html#ga902b8e60cc7075c8947345489221e0e0
func EMD(signature1, signature2 Mat, typ DistanceTypes) float32 {
	return 0
}

// ClipLine clips the line against the image rectangle.
// For further details, please see:
// https://docs.opencv.org/master/d6/d6e/group__imgproc__draw.html#gaf483cb46ad6b049bc35ec67052ef1c2c
func ClipLine(imgSize image.Point, pt1 image.Point, pt2 image.Point) bool {
	return false
}

// BilateralFilter applies a bilateral filter to an image.
//
// Bilateral filtering is described here:
// http://www.dai.ed.ac.uk/CVonline/LOCAL_COPIES/MANDUCHI1/Bilateral_Filtering.html
//
// BilateralFilter can reduce unwanted noise very well while keeping edges
// fairly sharp. However, it is very slow compared to most filters.
//
// For further details, please see:
// https://docs.opencv.org/master/d4/d86/group__imgproc__filter.html#ga9d7064d478c95d60003cf839430737ed
func BilateralFilter(src Mat, dst *Mat, diameter int, sigmaColor float64, sigmaSpace float64) {
	// Do nothing
}

// Blur blurs an image Mat using a normalized box filter.
//
// For further details, please see:
// https://docs.opencv.org/master/d4/d86/group__imgproc__filter.html#ga8c45db9afe636703801b0b2e440fce37
func Blur(src Mat, dst *Mat, ksize image.Point) {
	// Do nothing
}

// BoxFilter blurs an image using the box filter.
//
// For further details, please see:
// https://docs.opencv.org/master/d4/d86/group__imgproc__filter.html#gad533230ebf2d42509547d514f7d3fbc3
func BoxFilter(src Mat, dst *Mat, depth int, ksize image.Point) {
	// Do nothing
}

// SqBoxFilter calculates the normalized sum of squares of the pixel values overlapping the filter.
//
// For further details, please see:
// https://docs.opencv.org/master/d4/d86/group__imgproc__filter.html#ga045028184a9ef65d7d2579e5c4bff6c0
func SqBoxFilter(src Mat, dst *Mat, depth int, ksize image.Point) {
	// Do nothing
}

// Dilate dilates an image by using a specific structuring element.
//
// For further details, please see:
// https://docs.opencv.org/master/d4/d86/group__imgproc__filter.html#ga4ff0f3318642c4f469d0e11f242f3b6c
func Dilate(src Mat, dst *Mat, kernel Mat) {
	// Do nothing
}

// DilateWithParams dilates an image by using a specific structuring element.
//
// For further details, please see:
// https://docs.opencv.org/master/d4/d86/group__imgproc__filter.html#ga4ff0f3318642c4f469d0e11f242f3b6c
func DilateWithParams(src Mat, dst *Mat, kernel Mat, anchor image.Point, iterations, borderType BorderType, borderValue color.RGBA) {
	// Do nothing
}

// DistanceTransformLabelTypes are the types of the DistanceTransform algorithm flag
type DistanceTransformLabelTypes int

const (
	// DistanceLabelCComp assigns the same label to each connected component of zeros in the source image
	// (as well as all the non-zero pixels closest to the connected component).
	DistanceLabelCComp DistanceTransformLabelTypes = 0

	// DistanceLabelPixel assigns its own label to each zero pixel (and all the non-zero pixels closest to it).
	DistanceLabelPixel
)

// DistanceTransformMasks are the marsk sizes for distance transform
type DistanceTransformMasks int

const (
	// DistanceMask3 is a mask of size 3
	DistanceMask3 DistanceTransformMasks = 0

	// DistanceMask5 is a mask of size 3
	DistanceMask5

	// DistanceMaskPrecise is not currently supported
	DistanceMaskPrecise
)

// DistanceTransform Calculates the distance to the closest zero pixel for each pixel of the source image.
//
// For further details, please see:
// https://docs.opencv.org/master/d7/d1b/group__imgproc__misc.html#ga8a0b7fdfcb7a13dde018988ba3a43042
func DistanceTransform(src Mat, dst *Mat, labels *Mat, distType DistanceTypes, maskSize DistanceTransformMasks, labelType DistanceTransformLabelTypes) {
	// Do nothing
}

// Erode erodes an image by using a specific structuring element.
//
// For further details, please see:
// https://docs.opencv.org/master/d4/d86/group__imgproc__filter.html#gaeb1e0c1033e3f6b891a25d0511362aeb
func Erode(src Mat, dst *Mat, kernel Mat) {
	// Do nothing
}

// ErodeWithParams erodes an image by using a specific structuring element.
//
// For further details, please see:
// https://docs.opencv.org/master/d4/d86/group__imgproc__filter.html#gaeb1e0c1033e3f6b891a25d0511362aeb
func ErodeWithParams(src Mat, dst *Mat, kernel Mat, anchor image.Point, iterations, borderType int) {
	// Do nothing
}

// ErodeWithParamsAndBorderValue erodes an image by using a specific structuring
// element. Same as ErodeWithParams but requires an additional borderValue
// parameter.
//
// For further details, please see:
// https://docs.opencv.org/master/d4/d86/group__imgproc__filter.html#gaeb1e0c1033e3f6b891a25d0511362aeb
func ErodeWithParamsAndBorderValue(src Mat, dst *Mat, kernel Mat, anchor image.Point, iterations, borderType int, borderValue Scalar) {
	// Do nothing
}

// RetrievalMode is the mode of the contour retrieval algorithm.
type RetrievalMode int

const (
	// RetrievalExternal retrieves only the extreme outer contours.
	// It sets `hierarchy[i][2]=hierarchy[i][3]=-1` for all the contours.
	RetrievalExternal RetrievalMode = 0

	// RetrievalList retrieves all of the contours without establishing
	// any hierarchical relationships.
	RetrievalList RetrievalMode = 1

	// RetrievalCComp retrieves all of the contours and organizes them into
	// a two-level hierarchy. At the top level, there are external boundaries
	// of the components. At the second level, there are boundaries of the holes.
	// If there is another contour inside a hole of a connected component, it
	// is still put at the top level.
	RetrievalCComp RetrievalMode = 2

	// RetrievalTree retrieves all of the contours and reconstructs a full
	// hierarchy of nested contours.
	RetrievalTree RetrievalMode = 3

	// RetrievalFloodfill lacks a description in the original header.
	RetrievalFloodfill RetrievalMode = 4
)

// ContourApproximationMode is the mode of the contour approximation algorithm.
type ContourApproximationMode int

const (
	// ChainApproxNone stores absolutely all the contour points. That is,
	// any 2 subsequent points (x1,y1) and (x2,y2) of the contour will be
	// either horizontal, vertical or diagonal neighbors, that is,
	// max(abs(x1-x2),abs(y2-y1))==1.
	ChainApproxNone ContourApproximationMode = 1

	// ChainApproxSimple compresses horizontal, vertical, and diagonal segments
	// and leaves only their end points.
	// For example, an up-right rectangular contour is encoded with 4 points.
	ChainApproxSimple ContourApproximationMode = 2

	// ChainApproxTC89L1 applies one of the flavors of the Teh-Chin chain
	// approximation algorithms.
	ChainApproxTC89L1 ContourApproximationMode = 3

	// ChainApproxTC89KCOS applies one of the flavors of the Teh-Chin chain
	// approximation algorithms.
	ChainApproxTC89KCOS ContourApproximationMode = 4
)

// BoundingRect calculates the up-right bounding rectangle of a point set.
//
// For further details, please see:
// https://docs.opencv.org/3.3.0/d3/dc0/group__imgproc__shape.html#gacb413ddce8e48ff3ca61ed7cf626a366
func BoundingRect(contour PointVector) image.Rectangle {
	return image.Rectangle{}
}

// BoxPoints finds the four vertices of a rotated rect. Useful to draw the rotated rectangle.
//
// For further Details, please see:
// https://docs.opencv.org/3.3.0/d3/dc0/group__imgproc__shape.html#gaf78d467e024b4d7936cf9397185d2f5c
func BoxPoints(rect RotatedRect, pts *Mat) {
	// Do nothing
}

// BoxPoints finds the four vertices of a rotated rect. Useful to draw the rotated rectangle.
//
// For further Details, please see:
// https://docs.opencv.org/3.3.0/d3/dc0/group__imgproc__shape.html#gaf78d467e024b4d7936cf9397185d2f5c
func BoxPoints2f(rect RotatedRect2f, pts *Mat) {
	// Do nothing
}

// ContourArea calculates a contour area.
//
// For further details, please see:
// https://docs.opencv.org/3.3.0/d3/dc0/group__imgproc__shape.html#ga2c759ed9f497d4a618048a2f56dc97f1
func ContourArea(contour PointVector) float64 {
	return 0
}

type RotatedRect struct {
	Points       []image.Point
	BoundingRect image.Rectangle
	Center       image.Point
	Width        int
	Height       int
	Angle        float64
}

type RotatedRect2f struct {
	Points       []Point2f
	BoundingRect image.Rectangle
	Center       Point2f
	Width        float32
	Height       float32
	Angle        float64
}

// MinAreaRect finds a rotated rectangle of the minimum area enclosing the input 2D point set.
//
// For further details, please see:
// https://docs.opencv.org/master/d3/dc0/group__imgproc__shape.html#ga3d476a3417130ae5154aea421ca7ead9
func MinAreaRect(points PointVector) RotatedRect {
	return RotatedRect{}
}

// MinAreaRect finds a rotated rectangle of the minimum area enclosing the input 2D point set.
//
// For further details, please see:
// https://docs.opencv.org/master/d3/dc0/group__imgproc__shape.html#ga3d476a3417130ae5154aea421ca7ead9
func MinAreaRect2f(points PointVector) RotatedRect2f {
	return RotatedRect2f{}
}

// FitEllipse Fits an ellipse around a set of 2D points.
//
// For further details, please see:
// https://docs.opencv.org/master/d3/dc0/group__imgproc__shape.html#gaf259efaad93098103d6c27b9e4900ffa
func FitEllipse(pts PointVector) RotatedRect {
	return RotatedRect{}
}

// MinEnclosingCircle finds a circle of the minimum area enclosing the input 2D point set.
//
// For further details, please see:
// https://docs.opencv.org/3.4/d3/dc0/group__imgproc__shape.html#ga8ce13c24081bbc7151e9326f412190f1
func MinEnclosingCircle(pts PointVector) (x, y, radius float32) {
	return 0, 0, 0
}

// FindContours finds contours in a binary image.
//
// For further details, please see:
// https://docs.opencv.org/master/d3/dc0/group__imgproc__shape.html#ga95f5b48d01abc7c2e0732db24689837b
func FindContours(src Mat, mode RetrievalMode, method ContourApproximationMode) PointsVector {
	return PointsVector{}
}

// FindContoursWithParams finds contours in a binary image.
//
// For further details, please see:
// https://docs.opencv.org/master/d3/dc0/group__imgproc__shape.html#ga17ed9f5d79ae97bd4c7cf18403e1689a
func FindContoursWithParams(src Mat, hierarchy *Mat, mode RetrievalMode, method ContourApproximationMode) PointsVector {
	return PointsVector{}
}

// PointPolygonTest performs a point-in-contour test.
//
// For further details, please see:
// https://docs.opencv.org/master/d3/dc0/group__imgproc__shape.html#ga1a539e8db2135af2566103705d7a5722
func PointPolygonTest(pts PointVector, pt image.Point, measureDist bool) float64 {
	return 0
}

// ConnectedComponentsAlgorithmType specifies the type for ConnectedComponents
type ConnectedComponentsAlgorithmType int

const (
	// SAUF algorithm for 8-way connectivity, SAUF algorithm for 4-way connectivity.
	CCL_WU ConnectedComponentsAlgorithmType = 0

	// BBDT algorithm for 8-way connectivity, SAUF algorithm for 4-way connectivity.
	CCL_DEFAULT ConnectedComponentsAlgorithmType = 1

	// BBDT algorithm for 8-way connectivity, SAUF algorithm for 4-way connectivity
	CCL_GRANA ConnectedComponentsAlgorithmType = 2
)

// ConnectedComponents computes the connected components labeled image of boolean image.
//
// For further details, please see:
// https://docs.opencv.org/master/d3/dc0/group__imgproc__shape.html#gaedef8c7340499ca391d459122e51bef5
func ConnectedComponents(src Mat, labels *Mat) int {
	return 0
}

// ConnectedComponents computes the connected components labeled image of boolean image.
//
// For further details, please see:
// https://docs.opencv.org/master/d3/dc0/group__imgproc__shape.html#gaedef8c7340499ca391d459122e51bef5
func ConnectedComponentsWithParams(src Mat, labels *Mat, conn int, ltype MatType, ccltype ConnectedComponentsAlgorithmType) int {
	return 0
}

// ConnectedComponentsTypes are the connected components algorithm output formats
type ConnectedComponentsTypes int

const (
	//The leftmost (x) coordinate which is the inclusive start of the bounding box in the horizontal direction.
	CC_STAT_LEFT ConnectedComponentsTypes = 0

	//The topmost (y) coordinate which is the inclusive start of the bounding box in the vertical direction.
	CC_STAT_TOP ConnectedComponentsTypes = 1

	// The horizontal size of the bounding box.
	CC_STAT_WIDTH ConnectedComponentsTypes = 2

	// The vertical size of the bounding box.
	CC_STAT_HEIGHT ConnectedComponentsTypes = 3

	// The total area (in pixels) of the connected component.
	CC_STAT_AREA ConnectedComponentsTypes = 4

	CC_STAT_MAX ConnectedComponentsTypes = 5
)

// ConnectedComponentsWithStats computes the connected components labeled image of boolean
// image and also produces a statistics output for each label.
//
// For further details, please see:
// https://docs.opencv.org/master/d3/dc0/group__imgproc__shape.html#ga107a78bf7cd25dec05fb4dfc5c9e765f
func ConnectedComponentsWithStats(src Mat, labels *Mat, stats *Mat, centroids *Mat) int {
	return 0
}

// ConnectedComponentsWithStats computes the connected components labeled image of boolean
// image and also produces a statistics output for each label.
//
// For further details, please see:
// https://docs.opencv.org/master/d3/dc0/group__imgproc__shape.html#ga107a78bf7cd25dec05fb4dfc5c9e765f
func ConnectedComponentsWithStatsWithParams(src Mat, labels *Mat, stats *Mat, centroids *Mat,
	conn int, ltype MatType, ccltype ConnectedComponentsAlgorithmType) int {
	return 0
}

// TemplateMatchMode is the type of the template matching operation.
type TemplateMatchMode int

const (
	// TmSqdiff maps to TM_SQDIFF
	TmSqdiff TemplateMatchMode = 0
	// TmSqdiffNormed maps to TM_SQDIFF_NORMED
	TmSqdiffNormed TemplateMatchMode = 1
	// TmCcorr maps to TM_CCORR
	TmCcorr TemplateMatchMode = 2
	// TmCcorrNormed maps to TM_CCORR_NORMED
	TmCcorrNormed TemplateMatchMode = 3
	// TmCcoeff maps to TM_CCOEFF
	TmCcoeff TemplateMatchMode = 4
	// TmCcoeffNormed maps to TM_CCOEFF_NORMED
	TmCcoeffNormed TemplateMatchMode = 5
)

// MatchTemplate compares a template against overlapped image regions.
//
// For further details, please see:
// https://docs.opencv.org/master/df/dfb/group__imgproc__object.html#ga586ebfb0a7fb604b35a23d85391329be
func MatchTemplate(image Mat, templ Mat, result *Mat, method TemplateMatchMode, mask Mat) {
	// Do nothing
}

// Moments calculates all of the moments up to the third order of a polygon
// or rasterized shape.
//
// For further details, please see:
// https://docs.opencv.org/master/d3/dc0/group__imgproc__shape.html#ga556a180f43cab22649c23ada36a8a139
func Moments(src Mat, binaryImage bool) map[string]float64 {
	return nil
}

// PyrDown blurs an image and downsamples it.
//
// For further details, please see:
// https://docs.opencv.org/master/d4/d86/group__imgproc__filter.html#gaf9bba239dfca11654cb7f50f889fc2ff
func PyrDown(src Mat, dst *Mat, ksize image.Point, borderType BorderType) {
	// Do nothing
}

// PyrUp upsamples an image and then blurs it.
//
// For further details, please see:
// https://docs.opencv.org/master/d4/d86/group__imgproc__filter.html#gada75b59bdaaca411ed6fee10085eb784
func PyrUp(src Mat, dst *Mat, ksize image.Point, borderType BorderType) {
	// Do nothing
}

// MorphologyDefaultBorder returns "magic" border value for erosion and dilation.
// It is automatically transformed to Scalar::all(-DBL_MAX) for dilation.
//
// For further details, please see:
// https://docs.opencv.org/master/d4/d86/group__imgproc__filter.html#ga94756fad83d9d24d29c9bf478558c40a
func MorphologyDefaultBorderValue() Scalar {
	return Scalar{}
}

// MorphologyEx performs advanced morphological transformations.
//
// For further details, please see:
// https://docs.opencv.org/master/d4/d86/group__imgproc__filter.html#ga67493776e3ad1a3df63883829375201f
func MorphologyEx(src Mat, dst *Mat, op MorphType, kernel Mat) {
	// Do nothing
}

// MorphologyExWithParams performs advanced morphological transformations.
//
// For further details, please see:
// https://docs.opencv.org/master/d4/d86/group__imgproc__filter.html#ga67493776e3ad1a3df63883829375201f
func MorphologyExWithParams(src Mat, dst *Mat, op MorphType, kernel Mat, iterations int, borderType BorderType) {
	// Do nothing
}

// MorphShape is the shape of the structuring element used for Morphing operations.
type MorphShape int

const (
	// MorphRect is the rectangular morph shape.
	MorphRect MorphShape = 0

	// MorphCross is the cross morph shape.
	MorphCross MorphShape = 1

	// MorphEllipse is the ellipse morph shape.
	MorphEllipse MorphShape = 2
)

// GetStructuringElement returns a structuring element of the specified size
// and shape for morphological operations.
//
// For further details, please see:
// https://docs.opencv.org/master/d4/d86/group__imgproc__filter.html#gac342a1bb6eabf6f55c803b09268e36dc
func GetStructuringElement(shape MorphShape, ksize image.Point) Mat {
	return Mat{}
}

// MorphType type of morphological operation.
type MorphType int

const (
	// MorphErode operation
	MorphErode MorphType = 0

	// MorphDilate operation
	MorphDilate MorphType = 1

	// MorphOpen operation
	MorphOpen MorphType = 2

	// MorphClose operation
	MorphClose MorphType = 3

	// MorphGradient operation
	MorphGradient MorphType = 4

	// MorphTophat operation
	MorphTophat MorphType = 5

	// MorphBlackhat operation
	MorphBlackhat MorphType = 6

	// MorphHitmiss operation
	MorphHitmiss MorphType = 7
)

// BorderType type of border.
type BorderType int

const (
	// BorderConstant border type
	BorderConstant BorderType = 0

	// BorderReplicate border type
	BorderReplicate BorderType = 1

	// BorderReflect border type
	BorderReflect BorderType = 2

	// BorderWrap border type
	BorderWrap BorderType = 3

	// BorderReflect101 border type
	BorderReflect101 BorderType = 4

	// BorderTransparent border type
	BorderTransparent BorderType = 5

	// BorderDefault border type
	BorderDefault = BorderReflect101

	// BorderIsolated border type
	BorderIsolated BorderType = 16
)

// GaussianBlur blurs an image Mat using a Gaussian filter.
// The function convolves the src Mat image into the dst Mat using
// the specified Gaussian kernel params.
//
// For further details, please see:
// http://docs.opencv.org/master/d4/d86/group__imgproc__filter.html#gaabe8c836e97159a9193fb0b11ac52cf1
func GaussianBlur(src Mat, dst *Mat, ksize image.Point, sigmaX float64, sigmaY float64, borderType BorderType) {
	// Do nothing
}

// GetGaussianKernel returns Gaussian filter coefficients.
//
// For further details, please see:
// https://docs.opencv.org/master/d4/d86/group__imgproc__filter.html#gac05a120c1ae92a6060dd0db190a61afa
func GetGaussianKernel(ksize int, sigma float64) Mat {
	return Mat{}
}

// GetGaussianKernelWithParams returns Gaussian filter coefficients.
//
// For further details, please see:
// https://docs.opencv.org/master/d4/d86/group__imgproc__filter.html#gac05a120c1ae92a6060dd0db190a61afa
func GetGaussianKernelWithParams(ksize int, sigma float64, ktype MatType) Mat {
	return Mat{}
}

// Sobel calculates the first, second, third, or mixed image derivatives using an extended Sobel operator
//
// For further details, please see:
// https://docs.opencv.org/master/d4/d86/group__imgproc__filter.html#gacea54f142e81b6758cb6f375ce782c8d
func Sobel(src Mat, dst *Mat, ddepth MatType, dx, dy, ksize int, scale, delta float64, borderType BorderType) {
	// Do nothing
}

// SpatialGradient calculates the first order image derivative in both x and y using a Sobel operator.
//
// For further details, please see:
// https://docs.opencv.org/master/d4/d86/group__imgproc__filter.html#ga405d03b20c782b65a4daf54d233239a2
func SpatialGradient(src Mat, dx, dy *Mat, ksize MatType, borderType BorderType) {
	// Do nothing
}

// Laplacian calculates the Laplacian of an image.
//
// For further details, please see:
// https://docs.opencv.org/master/d4/d86/group__imgproc__filter.html#gad78703e4c8fe703d479c1860d76429e6
func Laplacian(src Mat, dst *Mat, dDepth MatType, size int, scale float64, delta float64, borderType BorderType) {
	// Do nothing
}

// Scharr calculates the first x- or y- image derivative using Scharr operator.
//
// For further details, please see:
// https://docs.opencv.org/master/d4/d86/group__imgproc__filter.html#gaa13106761eedf14798f37aa2d60404c9
func Scharr(src Mat, dst *Mat, dDepth MatType, dx int, dy int, scale float64, delta float64, borderType BorderType) {
	// Do nothing
}

// MedianBlur blurs an image using the median filter.
//
// For further details, please see:
// https://docs.opencv.org/master/d4/d86/group__imgproc__filter.html#ga564869aa33e58769b4469101aac458f9
func MedianBlur(src Mat, dst *Mat, ksize int) {
	// Do nothing
}

// Canny finds edges in an image using the Canny algorithm.
// The function finds edges in the input image image and marks
// them in the output map edges using the Canny algorithm.
// The smallest value between threshold1 and threshold2 is used
// for edge linking. The largest value is used to
// find initial segments of strong edges.
// See http://en.wikipedia.org/wiki/Canny_edge_detector
//
// For further details, please see:
// http://docs.opencv.org/master/dd/d1a/group__imgproc__feature.html#ga04723e007ed888ddf11d9ba04e2232de
func Canny(src Mat, edges *Mat, t1 float32, t2 float32) {
	// Do nothing
}

// CornerSubPix Refines the corner locations. The function iterates to find
// the sub-pixel accurate location of corners or radial saddle points.
//
// For further details, please see:
// https://docs.opencv.org/master/dd/d1a/group__imgproc__feature.html#ga354e0d7c86d0d9da75de9b9701a9a87e
func CornerSubPix(img Mat, corners *Mat, winSize image.Point, zeroZone image.Point, criteria TermCriteria) {
	// Do nothing
}

// GoodFeaturesToTrack determines strong corners on an image. The function
// finds the most prominent corners in the image or in the specified image region.
//
// For further details, please see:
// https://docs.opencv.org/master/dd/d1a/group__imgproc__feature.html#ga1d6bb77486c8f92d79c8793ad995d541
func GoodFeaturesToTrack(img Mat, corners *Mat, maxCorners int, quality float64, minDist float64) {
	// Do nothing
}

// GrabCutMode is the flag for GrabCut algorithm.
type GrabCutMode int

const (
	// GCInitWithRect makes the function initialize the state and the mask using the provided rectangle.
	// After that it runs the itercount iterations of the algorithm.
	GCInitWithRect GrabCutMode = 0
	// GCInitWithMask makes the function initialize the state using the provided mask.
	// GCInitWithMask and GCInitWithRect can be combined.
	// Then all the pixels outside of the ROI are automatically initialized with GC_BGD.
	GCInitWithMask GrabCutMode = 1
	// GCEval means that the algorithm should just resume.
	GCEval GrabCutMode = 2
	// GCEvalFreezeModel means that the algorithm should just run a single iteration of the GrabCut algorithm
	// with the fixed model
	GCEvalFreezeModel GrabCutMode = 3
)

// Grabcut runs the GrabCut algorithm.
// The function implements the GrabCut image segmentation algorithm.
// For further details, please see:
// https://docs.opencv.org/master/d7/d1b/group__imgproc__misc.html#ga909c1dda50efcbeaa3ce126be862b37f
func GrabCut(img Mat, mask *Mat, r image.Rectangle, bgdModel *Mat, fgdModel *Mat, iterCount int, mode GrabCutMode) {
	// Do nothing
}

// HoughMode is the type for Hough transform variants.
type HoughMode int

const (
	// HoughStandard is the classical or standard Hough transform.
	HoughStandard HoughMode = 0
	// HoughProbabilistic is the probabilistic Hough transform (more efficient
	// in case if the picture contains a few long linear segments).
	HoughProbabilistic HoughMode = 1
	// HoughMultiScale is the multi-scale variant of the classical Hough
	// transform.
	HoughMultiScale HoughMode = 2
	// HoughGradient is basically 21HT, described in: HK Yuen, John Princen,
	// John Illingworth, and Josef Kittler. Comparative study of hough
	// transform methods for circle finding. Image and Vision Computing,
	// 8(1):71–77, 1990.
	HoughGradient HoughMode = 3
)

// HoughCircles finds circles in a grayscale image using the Hough transform.
// The only "method" currently supported is HoughGradient. If you want to pass
// more parameters, please see `HoughCirclesWithParams`.
//
// For further details, please see:
// https://docs.opencv.org/master/dd/d1a/group__imgproc__feature.html#ga47849c3be0d0406ad3ca45db65a25d2d
func HoughCircles(src Mat, circles *Mat, method HoughMode, dp, minDist float64) {
	// Do nothing
}

// HoughCirclesWithParams finds circles in a grayscale image using the Hough
// transform. The only "method" currently supported is HoughGradient.
//
// For further details, please see:
// https://docs.opencv.org/master/dd/d1a/group__imgproc__feature.html#ga47849c3be0d0406ad3ca45db65a25d2d
func HoughCirclesWithParams(src Mat, circles *Mat, method HoughMode, dp, minDist, param1, param2 float64, minRadius, maxRadius int) {
	// Do nothing
}

// HoughLines implements the standard or standard multi-scale Hough transform
// algorithm for line detection. For a good explanation of Hough transform, see:
// http://homepages.inf.ed.ac.uk/rbf/HIPR2/hough.htm
//
// For further details, please see:
// http://docs.opencv.org/master/dd/d1a/group__imgproc__feature.html#ga46b4e588934f6c8dfd509cc6e0e4545a
func HoughLines(src Mat, lines *Mat, rho float32, theta float32, threshold int) {
	// Do nothing
}

// HoughLinesP implements the probabilistic Hough transform
// algorithm for line detection. For a good explanation of Hough transform, see:
// http://homepages.inf.ed.ac.uk/rbf/HIPR2/hough.htm
//
// For further details, please see:
// http://docs.opencv.org/master/dd/d1a/group__imgproc__feature.html#ga8618180a5948286384e3b7ca02f6feeb
func HoughLinesP(src Mat, lines *Mat, rho float32, theta float32, threshold int) {
	// Do nothing
}
func HoughLinesPWithParams(src Mat, lines *Mat, rho float32, theta float32, threshold int, minLineLength float32, maxLineGap float32) {
	// Do nothing
}

// HoughLinesPointSet implements the Hough transform algorithm for line
// detection on a set of points. For a good explanation of Hough transform, see:
// http://homepages.inf.ed.ac.uk/rbf/HIPR2/hough.htm
//
// For further details, please see:
// https://docs.opencv.org/master/dd/d1a/group__imgproc__feature.html#ga2858ef61b4e47d1919facac2152a160e
func HoughLinesPointSet(points Mat, lines *Mat, linesMax int, threshold int, minRho float32, maxRho float32, rhoStep float32, minTheta float32, maxTheta float32, thetaStep float32) {
	// Do nothing
}

// Integral calculates one or more integral images for the source image.
// For further details, please see:
// https://docs.opencv.org/master/d7/d1b/group__imgproc__misc.html#ga97b87bec26908237e8ba0f6e96d23e28
func Integral(src Mat, sum *Mat, sqsum *Mat, tilted *Mat) {
	// Do nothing
}

// ThresholdType type of threshold operation.
type ThresholdType int

const (
	// ThresholdBinary threshold type
	ThresholdBinary ThresholdType = 0

	// ThresholdBinaryInv threshold type
	ThresholdBinaryInv ThresholdType = 1

	// ThresholdTrunc threshold type
	ThresholdTrunc ThresholdType = 2

	// ThresholdToZero threshold type
	ThresholdToZero ThresholdType = 3

	// ThresholdToZeroInv threshold type
	ThresholdToZeroInv ThresholdType = 4

	// ThresholdMask threshold type
	ThresholdMask ThresholdType = 7

	// ThresholdOtsu threshold type
	ThresholdOtsu ThresholdType = 8

	// ThresholdTriangle threshold type
	ThresholdTriangle ThresholdType = 16
)

// Threshold applies a fixed-level threshold to each array element.
//
// For further details, please see:
// https://docs.opencv.org/3.3.0/d7/d1b/group__imgproc__misc.html#gae8a4a146d1ca78c626a53577199e9c57
func Threshold(src Mat, dst *Mat, thresh float32, maxvalue float32, typ ThresholdType) (threshold float32) {
	return 0
}

// AdaptiveThresholdType type of adaptive threshold operation.
type AdaptiveThresholdType int

const (
	// AdaptiveThresholdMean threshold type
	AdaptiveThresholdMean AdaptiveThresholdType = 0

	// AdaptiveThresholdGaussian threshold type
	AdaptiveThresholdGaussian AdaptiveThresholdType = 1
)

// AdaptiveThreshold applies a fixed-level threshold to each array element.
//
// For further details, please see:
// https://docs.opencv.org/master/d7/d1b/group__imgproc__misc.html#ga72b913f352e4a1b1b397736707afcde3
func AdaptiveThreshold(src Mat, dst *Mat, maxValue float32, adaptiveTyp AdaptiveThresholdType, typ ThresholdType, blockSize int, c float32) {
	// Do nothing
}

// ArrowedLine draws a arrow segment pointing from the first point
// to the second one.
//
// For further details, please see:
// https://docs.opencv.org/master/d6/d6e/group__imgproc__draw.html#ga0a165a3ca093fd488ac709fdf10c05b2
func ArrowedLine(img *Mat, pt1 image.Point, pt2 image.Point, c color.RGBA, thickness int) {
	// Do nothing
}

// Circle draws a circle.
//
// For further details, please see:
// https://docs.opencv.org/master/d6/d6e/group__imgproc__draw.html#gaf10604b069374903dbd0f0488cb43670
func Circle(img *Mat, center image.Point, radius int, c color.RGBA, thickness int) {
	// Do nothing
}

// CircleWithParams draws a circle.
//
// For further details, please see:
// https://docs.opencv.org/master/d6/d6e/group__imgproc__draw.html#gaf10604b069374903dbd0f0488cb43670
func CircleWithParams(img *Mat, center image.Point, radius int, c color.RGBA, thickness int, lineType LineType, shift int) {
	// Do nothing
}

// Ellipse draws a simple or thick elliptic arc or fills an ellipse sector.
//
// For further details, please see:
// https://docs.opencv.org/master/d6/d6e/group__imgproc__draw.html#ga28b2267d35786f5f890ca167236cbc69
func Ellipse(img *Mat, center, axes image.Point, angle, startAngle, endAngle float64, c color.RGBA, thickness int) {
	// Do nothing
}

// Ellipse draws a simple or thick elliptic arc or fills an ellipse sector.
//
// For further details, please see:
// https://docs.opencv.org/master/d6/d6e/group__imgproc__draw.html#ga28b2267d35786f5f890ca167236cbc69
func EllipseWithParams(img *Mat, center, axes image.Point, angle, startAngle, endAngle float64, c color.RGBA, thickness int, lineType LineType, shift int) {
	// Do nothing
}

// Line draws a line segment connecting two points.
//
// For further details, please see:
// https://docs.opencv.org/master/d6/d6e/group__imgproc__draw.html#ga7078a9fae8c7e7d13d24dac2520ae4a2
func Line(img *Mat, pt1 image.Point, pt2 image.Point, c color.RGBA, thickness int) {
	// Do nothing
}

// Rectangle draws a simple, thick, or filled up-right rectangle.
// It renders a rectangle with the desired characteristics to the target Mat image.
//
// For further details, please see:
// http://docs.opencv.org/master/d6/d6e/group__imgproc__draw.html#ga346ac30b5c74e9b5137576c9ee9e0e8c
func Rectangle(img *Mat, r image.Rectangle, c color.RGBA, thickness int) {
	// Do nothing
}

// RectangleWithParams draws a simple, thick, or filled up-right rectangle.
// It renders a rectangle with the desired characteristics to the target Mat image.
//
// For further details, please see:
// http://docs.opencv.org/master/d6/d6e/group__imgproc__draw.html#ga346ac30b5c74e9b5137576c9ee9e0e8c
func RectangleWithParams(img *Mat, r image.Rectangle, c color.RGBA, thickness int, lineType LineType, shift int) {
	// Do nothing
}

// FillPoly fills the area bounded by one or more polygons.
//
// For more information, see:
// https://docs.opencv.org/master/d6/d6e/group__imgproc__draw.html#gaf30888828337aa4c6b56782b5dfbd4b7
func FillPoly(img *Mat, pts PointsVector, c color.RGBA) {
	// Do nothing
}

// FillPolyWithParams fills the area bounded by one or more polygons.
//
// For more information, see:
// https://docs.opencv.org/master/d6/d6e/group__imgproc__draw.html#gaf30888828337aa4c6b56782b5dfbd4b7
func FillPolyWithParams(img *Mat, pts PointsVector, c color.RGBA, lineType LineType, shift int, offset image.Point) {
	// Do nothing
}

// Polylines draws several polygonal curves.
//
// For more information, see:
// https://docs.opencv.org/master/d6/d6e/group__imgproc__draw.html#ga1ea127ffbbb7e0bfc4fd6fd2eb64263c
func Polylines(img *Mat, pts PointsVector, isClosed bool, c color.RGBA, thickness int) {
	// Do nothing
}

// HersheyFont are the font libraries included in OpenCV.
// Only a subset of the available Hershey fonts are supported by OpenCV.
//
// For more information, see:
// http://sources.isc.org/utils/misc/hershey-font.txt
type HersheyFont int

const (
	// FontHersheySimplex is normal size sans-serif font.
	FontHersheySimplex HersheyFont = 0
	// FontHersheyPlain issmall size sans-serif font.
	FontHersheyPlain HersheyFont = 1
	// FontHersheyDuplex normal size sans-serif font
	// (more complex than FontHersheySIMPLEX).
	FontHersheyDuplex HersheyFont = 2
	// FontHersheyComplex i a normal size serif font.
	FontHersheyComplex HersheyFont = 3
	// FontHersheyTriplex is a normal size serif font
	// (more complex than FontHersheyCOMPLEX).
	FontHersheyTriplex HersheyFont = 4
	// FontHersheyComplexSmall is a smaller version of FontHersheyCOMPLEX.
	FontHersheyComplexSmall HersheyFont = 5
	// FontHersheyScriptSimplex is a hand-writing style font.
	FontHersheyScriptSimplex HersheyFont = 6
	// FontHersheyScriptComplex is a more complex variant of FontHersheyScriptSimplex.
	FontHersheyScriptComplex HersheyFont = 7
	// FontItalic is the flag for italic font.
	FontItalic HersheyFont = 16
)

// LineType are the line libraries included in OpenCV.
//
// For more information, see:
// https://vovkos.github.io/doxyrest-showcase/opencv/sphinx_rtd_theme/enum_cv_LineTypes.html
type LineType int

const (
	// Filled line
	Filled LineType = -1
	// Line4 4-connected line
	Line4 LineType = 4
	// Line8 8-connected line
	Line8 LineType = 8
	// LineAA antialiased line
	LineAA LineType = 16
)

// GetTextSize calculates the width and height of a text string.
// It returns an image.Point with the size required to draw text using
// a specific font face, scale, and thickness.
//
// For further details, please see:
// http://docs.opencv.org/master/d6/d6e/group__imgproc__draw.html#ga3d2abfcb995fd2db908c8288199dba82
func GetTextSize(text string, fontFace HersheyFont, fontScale float64, thickness int) image.Point {
	return image.Point{}
}

// GetTextSizeWithBaseline calculates the width and height of a text string including the basline of the text.
// It returns an image.Point with the size required to draw text using
// a specific font face, scale, and thickness as well as its baseline.
//
// For further details, please see:
// http://docs.opencv.org/master/d6/d6e/group__imgproc__draw.html#ga3d2abfcb995fd2db908c8288199dba82
func GetTextSizeWithBaseline(text string, fontFace HersheyFont, fontScale float64, thickness int) (image.Point, int) {
	return image.Point{}, 0
}

// PutText draws a text string.
// It renders the specified text string into the img Mat at the location
// passed in the "org" param, using the desired font face, font scale,
// color, and line thinkness.
//
// For further details, please see:
// http://docs.opencv.org/master/d6/d6e/group__imgproc__draw.html#ga5126f47f883d730f633d74f07456c576
func PutText(img *Mat, text string, org image.Point, fontFace HersheyFont, fontScale float64, c color.RGBA, thickness int) {
	// Do nothing
}

// PutTextWithParams draws a text string.
// It renders the specified text string into the img Mat at the location
// passed in the "org" param, using the desired font face, font scale,
// color, and line thinkness.
//
// For further details, please see:
// http://docs.opencv.org/master/d6/d6e/group__imgproc__draw.html#ga5126f47f883d730f633d74f07456c576
func PutTextWithParams(img *Mat, text string, org image.Point, fontFace HersheyFont, fontScale float64, c color.RGBA, thickness int, lineType LineType, bottomLeftOrigin bool) {
	// Do nothing
}

// InterpolationFlags are bit flags that control the interpolation algorithm
// that is used.
type InterpolationFlags int

const (
	// InterpolationNearestNeighbor is nearest neighbor. (fast but low quality)
	InterpolationNearestNeighbor InterpolationFlags = 0

	// InterpolationLinear is bilinear interpolation.
	InterpolationLinear InterpolationFlags = 1

	// InterpolationCubic is bicube interpolation.
	InterpolationCubic InterpolationFlags = 2

	// InterpolationArea uses pixel area relation. It is preferred for image
	// decimation as it gives moire-free results.
	InterpolationArea InterpolationFlags = 3

	// InterpolationLanczos4 is Lanczos interpolation over 8x8 neighborhood.
	InterpolationLanczos4 InterpolationFlags = 4

	// InterpolationDefault is an alias for InterpolationLinear.
	InterpolationDefault = InterpolationLinear

	// InterpolationMax indicates use maximum interpolation.
	InterpolationMax InterpolationFlags = 7

	// WarpFillOutliers fills all of the destination image pixels. If some of them correspond to outliers in the source image, they are set to zero.
	WarpFillOutliers = 8

	// WarpInverseMap, inverse transformation.
	WarpInverseMap = 16
)

// Resize resizes an image.
// It resizes the image src down to or up to the specified size, storing the
// result in dst. Note that src and dst may be the same image. If you wish to
// scale by factor, an empty sz may be passed and non-zero fx and fy. Likewise,
// if you wish to scale to an explicit size, a non-empty sz may be passed with
// zero for both fx and fy.
//
// For further details, please see:
// https://docs.opencv.org/master/da/d54/group__imgproc__transform.html#ga47a974309e9102f5f08231edc7e7529d
func Resize(src Mat, dst *Mat, sz image.Point, fx, fy float64, interp InterpolationFlags) {
	// Do nothing
}

// GetRectSubPix retrieves a pixel rectangle from an image with sub-pixel accuracy.
//
// For further details, please see:
// https://docs.opencv.org/master/da/d54/group__imgproc__transform.html#ga77576d06075c1a4b6ba1a608850cd614
func GetRectSubPix(src Mat, patchSize image.Point, center image.Point, dst *Mat) {
	// Do nothing
}

// GetRotationMatrix2D calculates an affine matrix of 2D rotation.
//
// For further details, please see:
// https://docs.opencv.org/master/da/d54/group__imgproc__transform.html#gafbbc470ce83812914a70abfb604f4326
func GetRotationMatrix2D(center image.Point, angle, scale float64) Mat {
	return Mat{}
}

// WarpAffine applies an affine transformation to an image. For more parameters please check WarpAffineWithParams
//
// For further details, please see:
// https://docs.opencv.org/master/da/d54/group__imgproc__transform.html#ga0203d9ee5fcd28d40dbc4a1ea4451983
func WarpAffine(src Mat, dst *Mat, m Mat, sz image.Point) {
	// Do nothing
}

// WarpAffineWithParams applies an affine transformation to an image.
//
// For further details, please see:
// https://docs.opencv.org/master/da/d54/group__imgproc__transform.html#ga0203d9ee5fcd28d40dbc4a1ea4451983
func WarpAffineWithParams(src Mat, dst *Mat, m Mat, sz image.Point, flags InterpolationFlags, borderType BorderType, borderValue color.RGBA) {
	// Do nothing
}

// WarpPerspective applies a perspective transformation to an image.
// For more parameters please check WarpPerspectiveWithParams.
//
// For further details, please see:
// https://docs.opencv.org/master/da/d54/group__imgproc__transform.html#gaf73673a7e8e18ec6963e3774e6a94b87
func WarpPerspective(src Mat, dst *Mat, m Mat, sz image.Point) {
	// Do nothing
}

// WarpPerspectiveWithParams applies a perspective transformation to an image.
//
// For further details, please see:
// https://docs.opencv.org/master/da/d54/group__imgproc__transform.html#gaf73673a7e8e18ec6963e3774e6a94b87
func WarpPerspectiveWithParams(src Mat, dst *Mat, m Mat, sz image.Point, flags InterpolationFlags, borderType BorderType, borderValue color.RGBA) {
	// Do nothing
}

// Watershed performs a marker-based image segmentation using the watershed algorithm.
//
// For further details, please see:
// https://docs.opencv.org/master/d7/d1b/group__imgproc__misc.html#ga3267243e4d3f95165d55a618c65ac6e1
func Watershed(image Mat, markers *Mat) {
	// Do nothing
}

// ColormapTypes are the 12 GNU Octave/MATLAB equivalent colormaps.
//
// For further details, please see:
// https://docs.opencv.org/master/d3/d50/group__imgproc__colormap.html
type ColormapTypes int

// List of the available color maps
//
// For further details, please see:
// https://docs.opencv.org/master/d3/d50/group__imgproc__colormap.html#ga9a805d8262bcbe273f16be9ea2055a65
const (
	ColormapAutumn  ColormapTypes = 0
	ColormapBone    ColormapTypes = 1
	ColormapJet     ColormapTypes = 2
	ColormapWinter  ColormapTypes = 3
	ColormapRainbow ColormapTypes = 4
	ColormapOcean   ColormapTypes = 5
	ColormapSummer  ColormapTypes = 6
	ColormapSpring  ColormapTypes = 7
	ColormapCool    ColormapTypes = 8
	ColormapHsv     ColormapTypes = 9
	ColormapPink    ColormapTypes = 10
	ColormapHot     ColormapTypes = 11
	ColormapParula  ColormapTypes = 12
)

// ApplyColorMap applies a GNU Octave/MATLAB equivalent colormap on a given image.
//
// For further details, please see:
// https://docs.opencv.org/master/d3/d50/group__imgproc__colormap.html#gadf478a5e5ff49d8aa24e726ea6f65d15
func ApplyColorMap(src Mat, dst *Mat, colormapType ColormapTypes) {
	// Do nothing
}

// ApplyCustomColorMap applies a custom defined colormap on a given image.
//
// For further details, please see:
// https://docs.opencv.org/master/d3/d50/group__imgproc__colormap.html#gacb22288ddccc55f9bd9e6d492b409cae
func ApplyCustomColorMap(src Mat, dst *Mat, customColormap Mat) {
	// Do nothing
}

// GetPerspectiveTransform returns 3x3 perspective transformation for the
// corresponding 4 point pairs as image.Point.
//
// For further details, please see:
// https://docs.opencv.org/master/da/d54/group__imgproc__transform.html#ga8c1ae0e3589a9d77fffc962c49b22043
func GetPerspectiveTransform(src, dst PointVector) Mat {
	return Mat{}
}

// GetPerspectiveTransform2f returns 3x3 perspective transformation for the
// corresponding 4 point pairs as gocv.Point2f.
//
// For further details, please see:
// https://docs.opencv.org/master/da/d54/group__imgproc__transform.html#ga8c1ae0e3589a9d77fffc962c49b22043
func GetPerspectiveTransform2f(src, dst Point2fVector) Mat {
	return Mat{}
}

// GetAffineTransform returns a 2x3 affine transformation matrix for the
// corresponding 3 point pairs as image.Point.
//
// For further details, please see:
// https://docs.opencv.org/master/da/d54/group__imgproc__transform.html#ga8f6d378f9f8eebb5cb55cd3ae295a999
func GetAffineTransform(src, dst PointVector) Mat {
	return Mat{}
}

// GetAffineTransform2f returns a 2x3 affine transformation matrix for the
// corresponding 3 point pairs as gocv.Point2f.
//
// For further details, please see:
// https://docs.opencv.org/master/da/d54/group__imgproc__transform.html#ga8f6d378f9f8eebb5cb55cd3ae295a999
func GetAffineTransform2f(src, dst Point2fVector) Mat {
	return Mat{}
}

type HomographyMethod int

const (
	HomographyMethodAllPoints HomographyMethod = 0
	HomographyMethodLMEDS     HomographyMethod = 4
	HomographyMethodRANSAC    HomographyMethod = 8
	HomographyMethodRHO       HomographyMethod = 16
)

// DrawContours draws contours outlines or filled contours.
//
// For further details, please see:
// https://docs.opencv.org/master/d6/d6e/group__imgproc__draw.html#ga746c0625f1781f1ffc9056259103edbc
func DrawContours(img *Mat, contours PointsVector, contourIdx int, c color.RGBA, thickness int) {
	// Do nothing
}

// DrawContoursWithParams draws contours outlines or filled contours.
//
// For further details, please see:
// https://docs.opencv.org/master/d6/d6e/group__imgproc__draw.html#ga746c0625f1781f1ffc9056259103edbc
func DrawContoursWithParams(img *Mat, contours PointsVector, contourIdx int, c color.RGBA, thickness int, lineType LineType, hierarchy Mat, maxLevel int, offset image.Point) {
	// Do nothing
}

// Remap applies a generic geometrical transformation to an image.
//
// For further details, please see:
// https://docs.opencv.org/master/da/d54/group__imgproc__transform.html#gab75ef31ce5cdfb5c44b6da5f3b908ea4
func Remap(src Mat, dst, map1, map2 *Mat, interpolation InterpolationFlags, borderMode BorderType, borderValue color.RGBA) {
	// Do nothing
}

// Filter2D applies an arbitrary linear filter to an image.
//
// For further details, please see:
// https://docs.opencv.org/master/d4/d86/group__imgproc__filter.html#ga27c049795ce870216ddfb366086b5a04
func Filter2D(src Mat, dst *Mat, ddepth MatType, kernel Mat, anchor image.Point, delta float64, borderType BorderType) {
	// Do nothing
}

// SepFilter2D applies a separable linear filter to the image.
//
// For further details, please see:
// https://docs.opencv.org/master/d4/d86/group__imgproc__filter.html#ga910e29ff7d7b105057d1625a4bf6318d
func SepFilter2D(src Mat, dst *Mat, ddepth MatType, kernelX, kernelY Mat, anchor image.Point, delta float64, borderType BorderType) {
	// Do nothing
}

// LogPolar remaps an image to semilog-polar coordinates space.
//
// For further details, please see:
// https://docs.opencv.org/master/da/d54/group__imgproc__transform.html#gaec3a0b126a85b5ca2c667b16e0ae022d
func LogPolar(src Mat, dst *Mat, center image.Point, m float64, flags InterpolationFlags) {
	// Do nothing
}

// LinearPolar remaps an image to polar coordinates space.
//
// For further details, please see:
// https://docs.opencv.org/master/da/d54/group__imgproc__transform.html#gaa38a6884ac8b6e0b9bed47939b5362f3
func LinearPolar(src Mat, dst *Mat, center image.Point, maxRadius float64, flags InterpolationFlags) {
	// Do nothing
}

// DistanceTypes types for Distance Transform and M-estimatorss
//
// For further details, please see:
// https://docs.opencv.org/master/d7/d1b/group__imgproc__misc.html#gaa2bfbebbc5c320526897996aafa1d8eb
type DistanceTypes int

const (
	DistUser   DistanceTypes = 0
	DistL1     DistanceTypes = 1
	DistL2     DistanceTypes = 2
	DistC      DistanceTypes = 3
	DistL12    DistanceTypes = 4
	DistFair   DistanceTypes = 5
	DistWelsch DistanceTypes = 6
	DistHuber  DistanceTypes = 7
)

// FitLine fits a line to a 2D or 3D point set.
//
// For further details, please see:
// https://docs.opencv.org/master/d3/dc0/group__imgproc__shape.html#gaf849da1fdafa67ee84b1e9a23b93f91f
func FitLine(pts PointVector, line *Mat, distType DistanceTypes, param, reps, aeps float64) {
	// Do nothing
}

// Shape matching methods.
//
// For further details, please see:
// https://docs.opencv.org/4.x/d3/dc0/group__imgproc__shape.html#gaadc90cb16e2362c9bd6e7363e6e4c317
type ShapeMatchModes int

const (
	ContoursMatchI1 ShapeMatchModes = 1
	ContoursMatchI2 ShapeMatchModes = 2
	ContoursMatchI3 ShapeMatchModes = 3
)

// Compares two shapes.
//
// For further details, please see:
// https://docs.opencv.org/4.x/d3/dc0/group__imgproc__shape.html#gaadc90cb16e2362c9bd6e7363e6e4c317
func MatchShapes(contour1 PointVector, contour2 PointVector, method ShapeMatchModes, parameter float64) float64 {
	return 0
}

// CLAHE is a wrapper around the cv::CLAHE algorithm.
type CLAHE struct {
	// Do nothing
}

// NewCLAHE returns a new CLAHE algorithm
//
// For further details, please see:
// https://docs.opencv.org/master/d6/db6/classcv_1_1CLAHE.html
func NewCLAHE() CLAHE {
	return CLAHE{}
}

// NewCLAHEWithParams returns a new CLAHE algorithm
//
// For further details, please see:
// https://docs.opencv.org/master/d6/db6/classcv_1_1CLAHE.html
func NewCLAHEWithParams(clipLimit float64, tileGridSize image.Point) CLAHE {
	return CLAHE{}
}

// Close CLAHE.
func (c *CLAHE) Close() error {
	return nil
}

// Apply CLAHE.
//
// For further details, please see:
// https://docs.opencv.org/master/d6/db6/classcv_1_1CLAHE.html#a4e92e0e427de21be8d1fae8dcd862c5e
func (c *CLAHE) Apply(src Mat, dst *Mat) {
	// Do nothing
}

func InvertAffineTransform(src Mat, dst *Mat) {
	// Do nothing
}

// Apply phaseCorrelate.
//
// For further details, please see:
// https://docs.opencv.org/master/d7/df3/group__imgproc__motion.html#ga552420a2ace9ef3fb053cd630fdb4952
func PhaseCorrelate(src1, src2, window Mat) (phaseShift Point2f, response float64) {
	return Point2f{}, 0
}

// CreateHanningWindow computes a Hanning window coefficients in two dimensions.
//
// For further details, please see:
// https://docs.opencv.org/4.x/d7/df3/group__imgproc__motion.html#ga80e5c3de52f6bab3a7c1e60e89308e1b
func CreateHanningWindow(img *Mat, size image.Point, typ MatType) {
	// Do nothing
}

// ToImage converts a Mat to a image.Image.
func (m *Mat) ToImage() (image.Image, error) {
	return nil, nil
}

// ToImageYUV converts a Mat to a image.YCbCr using image.YCbCrSubsampleRatio420 as default subsampling param.
func (m *Mat) ToImageYUV() (*image.YCbCr, error) {
	return nil, nil
}

// ToImageYUV converts a Mat to a image.YCbCr using provided YUV subsample ratio param.
func (m *Mat) ToImageYUVWithParams(ratio image.YCbCrSubsampleRatio) (*image.YCbCr, error) {
	return nil, nil
}

// ImageToMatRGBA converts image.Image to gocv.Mat,
// which represents RGBA image having 8bit for each component.
// Type of Mat is gocv.MatTypeCV8UC4.
func ImageToMatRGBA(img image.Image) (Mat, error) {
	return Mat{}, nil
}

// ImageToMatRGB converts image.Image to gocv.Mat,
// which represents RGB image having 8bit for each component.
// Type of Mat is gocv.MatTypeCV8UC3.
func ImageToMatRGB(img image.Image) (Mat, error) {
	return Mat{}, nil
}

// ImageGrayToMatGray converts image.Gray to gocv.Mat,
// which represents grayscale image 8bit.
// Type of Mat is gocv.MatTypeCV8UC1.
func ImageGrayToMatGray(img *image.Gray) (Mat, error) {
	return Mat{}, nil
}

// Adds the square of a source image to the accumulator image.
//
// For further details, please see:
// https://docs.opencv.org/master/d7/df3/group__imgproc__motion.html#ga1a567a79901513811ff3b9976923b199
//

func Accumulate(src Mat, dst *Mat) {
	// Do nothing
}

// Adds an image to the accumulator image with mask.
//
// For further details, please see:
// https://docs.opencv.org/master/d7/df3/group__imgproc__motion.html#ga1a567a79901513811ff3b9976923b199
func AccumulateWithMask(src Mat, dst *Mat, mask Mat) {
	// Do nothing
}

// Adds the square of a source image to the accumulator image.
//
// For further details, please see:
// https://docs.opencv.org/master/d7/df3/group__imgproc__motion.html#gacb75e7ffb573227088cef9ceaf80be8c
func AccumulateSquare(src Mat, dst *Mat) {
	// Do nothing
}

// Adds the square of a source image to the accumulator image with mask.
//
// For further details, please see:
// https://docs.opencv.org/master/d7/df3/group__imgproc__motion.html#gacb75e7ffb573227088cef9ceaf80be8c
func AccumulateSquareWithMask(src Mat, dst *Mat, mask Mat) {
	// Do nothing
}

// Adds the per-element product of two input images to the accumulator image.
//
// For further details, please see:
// https://docs.opencv.org/master/d7/df3/group__imgproc__motion.html#ga82518a940ecfda49460f66117ac82520
func AccumulateProduct(src1 Mat, src2 Mat, dst *Mat) {
	// Do nothing
}

// Adds the per-element product of two input images to the accumulator image with mask.
//
// For further details, please see:
// https://docs.opencv.org/master/d7/df3/group__imgproc__motion.html#ga82518a940ecfda49460f66117ac82520
func AccumulateProductWithMask(src1 Mat, src2 Mat, dst *Mat, mask Mat) {
	// Do nothing
}

// Updates a running average.
//
// For further details, please see:
// https://docs.opencv.org/master/d7/df3/group__imgproc__motion.html#ga4f9552b541187f61f6818e8d2d826bc7
func AccumulatedWeighted(src Mat, dst *Mat, alpha float64) {
	// Do nothing
}

// Updates a running average with mask.
//
// For further details, please see:
// https://docs.opencv.org/master/d7/df3/group__imgproc__motion.html#ga4f9552b541187f61f6818e8d2d826bc7
func AccumulatedWeightedWithMask(src Mat, dst *Mat, alpha float64, mask Mat) {
	// Do nothing
}
