/*
Package d3dmath provices vector and matrix functions for Direct3D. Vectors are
row vectors and matrices are stored in column-major order.
*/
package d3dmath

import (
	"fmt"
	"math"
)

// These factors can be used to convert between turns (as used for the rotation
// functions in this package), radians and degrees.
// For example, to convert a half rotation in degrees to radians, you can say
// 180 * DegToRad.
const (
	TurnsToRad = 2 * math.Pi
	RadToTurns = 1.0 / TurnsToRad
	RadToDeg   = 180.0 / math.Pi
	DegToRad   = 1.0 / RadToDeg
	TurnsToDeg = 360.0
	DegToTurns = 1.0 / TurnsToDeg
)

// Vec2 is a 2-element row vector. Elements are called x, y in the docs.
type Vec2 [2]float32

// Negate returns a vector with all elements of v negated.
func (v Vec2) Negate() Vec2 {
	return Vec2{-v[0], -v[1]}
}

// Add returns the sum of v + w.
func (v Vec2) Add(w Vec2) Vec2 {
	return Vec2{v[0] + w[0], v[1] + w[1]}
}

// Sub returns the difference of v - w.
func (v Vec2) Sub(w Vec2) Vec2 {
	return Vec2{v[0] - w[0], v[1] - w[1]}
}

// Dot returns the dot-product of v and w.
func (v Vec2) Dot(w Vec2) float32 {
	return v[0]*w[0] + v[1]*w[1]
}

// MulScalar returns a vector with all elements of v scaled by s.
func (v Vec2) MulScalar(s float32) Vec2 {
	return Vec2{v[0] * s, v[1] * s}
}

// MulMat returns the product of row vector v and matrix m.
func (v Vec2) MulMat(m Mat2) Vec2 {
	return Vec2{
		v[0]*m[0] + v[1]*m[1],
		v[0]*m[2] + v[1]*m[3],
	}
}

// SquareNorm returns the square of the length of v.
func (v Vec2) SquareNorm() float32 {
	return v[0]*v[0] + v[1]*v[1]
}

// Norm returns the length of v.
func (v Vec2) Norm() float32 {
	return float32(math.Hypot(float64(v[0]), float64(v[1])))
}

// Normalized returns a copy of v with elements normalized so the returned
// vector has length 1, or the zero vector if the length is 0.
func (v Vec2) Normalized() Vec2 {
	norm := v.Norm()
	if norm == 0 {
		return Vec2{}
	}
	f := 1.0 / norm
	return Vec2{f * v[0], f * v[1]}
}

// Homogeneous returns a 3-element vector where x and y are the same as in v and
// z is 1.
func (v Vec2) Homogeneous() Vec3 {
	return Vec3{v[0], v[1], 1}
}

func (v Vec2) String() string {
	return fmt.Sprintf("(%.2f %.2f)", v[0], v[1])
}

// AddVec2 returns the sum of all given vectors.
func AddVec2(v0 Vec2, v ...Vec2) Vec2 {
	if len(v) == 0 {
		return v0
	}
	return v0.Add(AddVec2(v[0], v[1:]...))
}

// Vec3 is a 3-element row vector. Elements are called x, y, z in the docs.
type Vec3 [3]float32

// Negate returns a vector with all elements of v negated.
func (v Vec3) Negate() Vec3 {
	return Vec3{-v[0], -v[1], -v[2]}
}

// Add returns the sum of v + w.
func (v Vec3) Add(w Vec3) Vec3 {
	return Vec3{v[0] + w[0], v[1] + w[1], v[2] + w[2]}
}

// Sub returns the difference of v - w.
func (v Vec3) Sub(w Vec3) Vec3 {
	return Vec3{v[0] - w[0], v[1] - w[1], v[2] - w[2]}
}

// Dot returns the dot-product of v and w.
func (v Vec3) Dot(w Vec3) float32 {
	return v[0]*w[0] + v[1]*w[1] + v[2]*w[2]
}

// Cross returns the cross-product of v and w.
func (v Vec3) Cross(w Vec3) Vec3 {
	return Vec3{
		v[1]*w[2] - v[2]*w[1],
		v[2]*w[0] - v[0]*w[2],
		v[0]*w[1] - v[1]*w[0],
	}
}

// MulScalar returns a vector with all elements of v scaled by s.
func (v Vec3) MulScalar(s float32) Vec3 {
	return Vec3{v[0] * s, v[1] * s, v[2] * s}
}

// MulMat returns the product of row vector v and matrix m.
func (v Vec3) MulMat(m Mat3) Vec3 {
	return Vec3{
		v[0]*m[0] + v[1]*m[1] + v[2]*m[2],
		v[0]*m[3] + v[1]*m[4] + v[2]*m[5],
		v[0]*m[6] + v[1]*m[7] + v[2]*m[8],
	}
}

// SquareNorm returns the square of the length of v.
func (v Vec3) SquareNorm() float32 {
	return v[0]*v[0] + v[1]*v[1] + v[2]*v[2]
}

// Norm returns the length of v.
func (v Vec3) Norm() float32 {
	return float32(math.Sqrt(float64(v.SquareNorm())))
}

// Normalized returns a copy of v with elements normalized so the returned
// vector has length 1, or the zero vector if the length is 0.
func (v Vec3) Normalized() Vec3 {
	norm := v.Norm()
	if norm == 0 {
		return Vec3{}
	}
	f := 1.0 / v.Norm()
	return Vec3{f * v[0], f * v[1], f * v[2]}
}

// Homogeneous returns a 4-element vector where x, y and z are the same as in v
// and w is 1.
func (v Vec3) Homogeneous() Vec4 {
	return Vec4{v[0], v[1], v[2], 1}
}

// DropZ returns a 2-element vector where x and y are the same as in v.
// This can be useful when going back from a homogeneous 3-element vector with z
// == 1, down one dimension to a 2-element vector.
// If z != 1 then use ByZ() to divide by z instead.
func (v Vec3) DropZ() Vec2 {
	return Vec2{v[0], v[1]}
}

// ByZ returns a 2-element vector created by dividing x and y by z. This can be
// useful when going back from a homogeneous 3-element vector with z != 1, down
// one dimension to a 2-element vector.
func (v Vec3) ByZ() Vec2 {
	f := float32(1.0)
	if v[2] != 0 {
		f = 1.0 / v[2]
	}
	return Vec2{f * v[0], f * v[1]}
}

func (v Vec3) String() string {
	return fmt.Sprintf("(%.2f %.2f %.2f)", v[0], v[1], v[2])
}

// AddVec3 returns the sum of all given vectors.
func AddVec3(v0 Vec3, v ...Vec3) Vec3 {
	if len(v) == 0 {
		return v0
	}
	return v0.Add(AddVec3(v[0], v[1:]...))
}

// Vec4 is a 4-element row vector. Elements are called x, y, z, w in the docs.
type Vec4 [4]float32

// Negate returns a vector with all elements of v negated.
func (v Vec4) Negate() Vec4 {
	return Vec4{-v[0], -v[1], -v[2], -v[3]}
}

// Add returns the sum of v + w.
func (v Vec4) Add(w Vec4) Vec4 {
	return Vec4{v[0] + w[0], v[1] + w[1], v[2] + w[2], v[3] + w[3]}
}

// Sub returns the difference of v - w.
func (v Vec4) Sub(w Vec4) Vec4 {
	return Vec4{v[0] - w[0], v[1] - w[1], v[2] - w[2], v[3] - w[3]}
}

// Dot returns the dot-product of v and w.
func (v Vec4) Dot(w Vec4) float32 {
	return v[0]*w[0] + v[1]*w[1] + v[2]*w[2] + v[3]*w[3]
}

// MulScalar returns a vector with all elements of v scaled by s.
func (v Vec4) MulScalar(s float32) Vec4 {
	return Vec4{v[0] * s, v[1] * s, v[2] * s, v[3] * s}
}

// MulMat returns the product of row vector v and matrix m.
func (v Vec4) MulMat(m Mat4) Vec4 {
	return Vec4{
		v[0]*m[0] + v[1]*m[1] + v[2]*m[2] + v[3]*m[3],
		v[0]*m[4] + v[1]*m[5] + v[2]*m[6] + v[3]*m[7],
		v[0]*m[8] + v[1]*m[9] + v[2]*m[10] + v[3]*m[11],
		v[0]*m[12] + v[1]*m[13] + v[2]*m[14] + v[3]*m[15],
	}
}

// SquareNorm returns the square of the length of v.
func (v Vec4) SquareNorm() float32 {
	return v[0]*v[0] + v[1]*v[1] + v[2]*v[2] + v[3]*v[3]
}

// Norm returns the length of v.
func (v Vec4) Norm() float32 {
	return float32(math.Sqrt(float64(v.SquareNorm())))
}

// Normalized returns a copy of v with elements normalized so the returned
// vector has length 1, or the zero vector if the length is 0.
func (v Vec4) Normalized() Vec4 {
	norm := v.Norm()
	if norm == 0 {
		return Vec4{}
	}
	f := 1.0 / norm
	return Vec4{f * v[0], f * v[1], f * v[2], f * v[3]}
}

// DropW returns a 3-element vector where x, y and z are the same as in v.
// This can be useful when going back from a homogeneous 4-element vector with w
// == 1, down one dimension to a 3-element vector.
// If w != 1 then use ByW() to divide by w instead.
func (v Vec4) DropW() Vec3 {
	return Vec3{v[0], v[1], v[2]}
}

// ByW returns a 3-element vector created by dividing x, y and z by w. This can
// be useful when going back from a homogeneous 4-element vector with w != 1,
// down one dimension to a 3-element vector.
func (v Vec4) ByW() Vec3 {
	f := float32(1.0)
	if v[3] != 0 {
		f = 1.0 / v[3]
	}
	return Vec3{f * v[0], f * v[1], f * v[2]}
}

func (v Vec4) String() string {
	return fmt.Sprintf("(%.2f %.2f %.2f %.2f)", v[0], v[1], v[2], v[3])
}

// AddVec4 returns the sum of all given vectors.
func AddVec4(v0 Vec4, v ...Vec4) Vec4 {
	if len(v) == 0 {
		return v0
	}
	return v0.Add(AddVec4(v[0], v[1:]...))
}

// Mat2 is a 2 by 2 matrix of float32s in column-major order.
type Mat2 [4]float32

// Add returns the sum of m + n.
func (m Mat2) Add(n Mat2) Mat2 {
	return Mat2{
		m[0] + n[0], m[1] + n[1],
		m[2] + n[2], m[3] + n[3],
	}
}

// Sub returns the difference of m - n.
func (m Mat2) Sub(n Mat2) Mat2 {
	return Mat2{
		m[0] - n[0], m[1] - n[1],
		m[2] - n[2], m[3] - n[3],
	}
}

// Mul returns the product of m * n.
func (m Mat2) Mul(n Mat2) Mat2 {
	return Mat2{
		m[0]*n[0] + m[2]*n[1],
		m[1]*n[0] + m[3]*n[1],

		m[0]*n[2] + m[2]*n[3],
		m[1]*n[2] + m[3]*n[3],
	}
}

// Identity2 returns the 2 by 2 identity matrix.
func Identity2() Mat2 {
	return Mat2{
		1, 0,
		0, 1,
	}
}

// Mul2 returns the product of the given matrices.
func Mul2(m0 Mat2, m ...Mat2) Mat2 {
	if len(m) == 0 {
		return m0
	}
	return m0.Mul(Mul2(m[0], m[1:]...))
}

// Transposed returns a transposed copy of m.
func (m Mat2) Transposed() Mat2 {
	return Mat2{
		m[0], m[2],
		m[1], m[3],
	}
}

// Homogeneous returns the homogeneous 3-dimensional equivalent of the
// 2-dimensional matrix.
func (m Mat2) Homogeneous() Mat3 {
	return Mat3{
		m[0], m[1], 0,
		m[2], m[3], 0,
		0, 0, 1,
	}
}

func (m Mat2) String() string {
	return fmt.Sprintf(`%.2f %.2f
%.2f %.2f`, m[0], m[2], m[1], m[3])
}

// Mat3 is a 3 by 3 matrix of float32s in column-major order.
type Mat3 [9]float32

// Add returns the sum of m + n.
func (m Mat3) Add(n Mat3) (sum Mat3) {
	for i := range sum {
		sum[i] = m[i] + n[i]
	}
	return
}

// Sub returns the difference of m - n.
func (m Mat3) Sub(n Mat3) (diff Mat3) {
	for i := range diff {
		diff[i] = m[i] - n[i]
	}
	return
}

// Mul returns the product of m * n.
func (m Mat3) Mul(n Mat3) Mat3 {
	return Mat3{
		m[0]*n[0] + m[3]*n[1] + m[6]*n[2],
		m[1]*n[0] + m[4]*n[1] + m[7]*n[2],
		m[2]*n[0] + m[5]*n[1] + m[8]*n[2],

		m[0]*n[3] + m[3]*n[4] + m[6]*n[5],
		m[1]*n[3] + m[4]*n[4] + m[7]*n[5],
		m[2]*n[3] + m[5]*n[4] + m[8]*n[5],

		m[0]*n[6] + m[3]*n[7] + m[6]*n[8],
		m[1]*n[6] + m[4]*n[7] + m[7]*n[8],
		m[2]*n[6] + m[5]*n[7] + m[8]*n[8],
	}
}

// Identity3 returns the 3 by 3 identity matrix.
func Identity3() Mat3 {
	return Mat3{
		1, 0, 0,
		0, 1, 0,
		0, 0, 1,
	}
}

// Mul3 returns the product of the given matrices.
func Mul3(m0 Mat3, m ...Mat3) Mat3 {
	if len(m) == 0 {
		return m0
	}
	return m0.Mul(Mul3(m[0], m[1:]...))
}

// Transposed returns a transposed copy of m.
func (m Mat3) Transposed() Mat3 {
	return Mat3{
		m[0], m[3], m[6],
		m[1], m[4], m[7],
		m[2], m[5], m[8],
	}
}

// Homogeneous returns the homogeneous 4-dimensional equivalent of the
// 3-dimensional matrix.
func (m Mat3) Homogeneous() Mat4 {
	return Mat4{
		m[0], m[1], m[2], 0,
		m[3], m[4], m[5], 0,
		m[6], m[7], m[8], 0,
		0, 0, 0, 1,
	}
}

func (m Mat3) String() string {
	return fmt.Sprintf(`%.2f %.2f %.2f
%.2f %.2f %.2f
%.2f %.2f %.2f`, m[0], m[3], m[6], m[1], m[4], m[7], m[2], m[5], m[8])
}

// Mat2x3 is a 2x3 matrix of float32s in column-major order. It represents a
// homogeneous 3x3 matrix where the last line is 0,0,1 implicitly.
type Mat2x3 [6]float32

// Add returns the sum of m + n.
func (m Mat2x3) Add(n Mat2x3) (sum Mat2x3) {
	for i := range sum {
		sum[i] = m[i] + n[i]
	}
	return
}

// Sub returns the difference of m - n.
func (m Mat2x3) Sub(n Mat2x3) (diff Mat2x3) {
	for i := range diff {
		diff[i] = m[i] - n[i]
	}
	return
}

// Mul returns the product of m * n.
func (m Mat2x3) Mul(n Mat2x3) Mat2x3 {
	return Mat2x3{
		m[0]*n[0] + m[2]*n[1],
		m[1]*n[0] + m[3]*n[1],

		m[0]*n[2] + m[2]*n[3],
		m[1]*n[2] + m[3]*n[3],

		m[0]*n[4] + m[2]*n[5] + m[4],
		m[1]*n[4] + m[3]*n[5] + m[5],
	}
}

// Identity2x3 returns the 2 by 3 homogeneous identity matrix.
func Identity2x3() Mat2x3 {
	return Mat2x3{
		1, 0,
		0, 1,
		0, 0,
	}
}

// Mul2x3 returns the product of the given matrices.
func Mul2x3(m0 Mat2x3, m ...Mat2x3) Mat2x3 {
	if len(m) == 0 {
		return m0
	}
	return m0.Mul(Mul2x3(m[0], m[1:]...))
}

// ToMat3 returns the 3 by 3 representation of m.
func (m Mat2x3) ToMat3() Mat3 {
	return Mat3{
		m[0], m[1], 0,
		m[2], m[3], 0,
		m[4], m[5], 1,
	}
}

func (m Mat2x3) String() string {
	return fmt.Sprintf(`%.2f %.2f %.2f
%.2f %.2f %.2f`, m[0], m[2], m[4], m[1], m[3], m[5])
}

// Mat4 is a 4 by 4 matrix of float32s in column-major order.
type Mat4 [16]float32

// Add returns the sum of m + n.
func (m Mat4) Add(n Mat4) (sum Mat4) {
	for i := range sum {
		sum[i] = m[i] + n[i]
	}
	return
}

// Sub returns the difference of m - n.
func (m Mat4) Sub(n Mat4) (diff Mat4) {
	for i := range diff {
		diff[i] = m[i] - n[i]
	}
	return
}

// Mul returns the product of m * n.
func (m Mat4) Mul(n Mat4) Mat4 {
	return Mat4{
		m[0]*n[0] + m[4]*n[1] + m[8]*n[2] + m[12]*n[3],
		m[1]*n[0] + m[5]*n[1] + m[9]*n[2] + m[13]*n[3],
		m[2]*n[0] + m[6]*n[1] + m[10]*n[2] + m[14]*n[3],
		m[3]*n[0] + m[7]*n[1] + m[11]*n[2] + m[15]*n[3],

		m[0]*n[4] + m[4]*n[5] + m[8]*n[6] + m[12]*n[7],
		m[1]*n[4] + m[5]*n[5] + m[9]*n[6] + m[13]*n[7],
		m[2]*n[4] + m[6]*n[5] + m[10]*n[6] + m[14]*n[7],
		m[3]*n[4] + m[7]*n[5] + m[11]*n[6] + m[15]*n[7],

		m[0]*n[8] + m[4]*n[9] + m[8]*n[10] + m[12]*n[11],
		m[1]*n[8] + m[5]*n[9] + m[9]*n[10] + m[13]*n[11],
		m[2]*n[8] + m[6]*n[9] + m[10]*n[10] + m[14]*n[11],
		m[3]*n[8] + m[7]*n[9] + m[11]*n[10] + m[15]*n[11],

		m[0]*n[12] + m[4]*n[13] + m[8]*n[14] + m[12]*n[15],
		m[1]*n[12] + m[5]*n[13] + m[9]*n[14] + m[13]*n[15],
		m[2]*n[12] + m[6]*n[13] + m[10]*n[14] + m[14]*n[15],
		m[3]*n[12] + m[7]*n[13] + m[11]*n[14] + m[15]*n[15],
	}
}

// Mul4 returns the product of the given matrices.
func Mul4(m0 Mat4, m ...Mat4) Mat4 {
	if len(m) == 0 {
		return m0
	}
	return m0.Mul(Mul4(m[0], m[1:]...))
}

// Transposed returns a transposed copy of m.
func (m Mat4) Transposed() Mat4 {
	return Mat4{
		m[0], m[4], m[8], m[12],
		m[1], m[5], m[9], m[13],
		m[2], m[6], m[10], m[14],
		m[3], m[7], m[11], m[15],
	}
}

// Identity4 returns the 4 by 4 identity matrix.
func Identity4() Mat4 {
	return Mat4{
		1, 0, 0, 0,
		0, 1, 0, 0,
		0, 0, 1, 0,
		0, 0, 0, 1,
	}
}

// Translate returns 4 by 4 matrix that, when multiplied with a homogeneous
// 4-element 3D vector, moves the vector by the given amounts in x, y and z.
func Translate(dx, dy, dz float32) Mat4 {
	return Mat4{
		1, 0, 0, dx,
		0, 1, 0, dy,
		0, 0, 1, dz,
		0, 0, 0, 1,
	}
}

// TranslateV is the same as Translate, but it takes a Vec3 as its argument
// instead of single x, y, z parameters.
func TranslateV(v Vec3) Mat4 {
	return Translate(v[0], v[1], v[2])
}

// ScaleUniform returns 4 by 4 matrix that, when multiplied with a homogeneous
// 4-element 3D vector, scales the vector by the given factor in x, y and z.
func ScaleUniform(s float32) Mat4 {
	return Scale(s, s, s)
}

// Scale returns 4 by 4 matrix that, when multiplied with a homogeneous
// 4-element 3D vector, scales the vector by the given factors in x, y and z.
func Scale(dx, dy, dz float32) Mat4 {
	return Mat4{
		dx, 0, 0, 0,
		0, dy, 0, 0,
		0, 0, dz, 0,
		0, 0, 0, 1,
	}
}

// ScaleV is the same as Scale, but it takes a Vec3 as its argument instead of
// single x, y, z parameters.
func ScaleV(v Vec3) Mat4 {
	return Scale(v[0], v[1], v[2])
}

func turnsToRadians(turns float32) float64 {
	return float64(turns) * 2 * math.Pi
}

// RotateLeftHandX returns 4 by 4 matrix that, when multiplied with a
// homogeneous 4-element 3D vector, rotates the vector about the x-axis,
// applying the left-handed rule, by the given number of turns. 1 turn is 2*Pi.
func RotateLeftHandX(turns float32) Mat4 {
	return RotateRightHandX(-turns)
}

// RotateRightHandX returns 4 by 4 matrix that, when multiplied with a
// homogeneous 4-element 3D vector, rotates the vector about the x-axis,
// applying the right-handed rule, by the given number of turns. 1 turn is
// 2*Pi.
func RotateRightHandX(turns float32) Mat4 {
	s, c := math.Sincos(turnsToRadians(turns))
	sin, cos := float32(s), float32(c)
	return Mat4{
		1, 0, 0, 0,
		0, cos, sin, 0,
		0, -sin, cos, 0,
		0, 0, 0, 1,
	}
}

// RotateLeftHandY returns 4 by 4 matrix that, when multiplied with a
// homogeneous 4-element 3D vector, rotates the vector about the y-axis,
// applying the left-handed rule, by the given number of turns. 1 turn is 2*Pi.
func RotateLeftHandY(turns float32) Mat4 {
	return RotateRightHandY(-turns)
}

// RotateRightHandY returns 4 by 4 matrix that, when multiplied with a
// homogeneous 4-element 3D vector, rotates the vector about the y-axis,
// applying the right-handed rule, by the given number of turns. 1 turn is
// 2*Pi.
func RotateRightHandY(turns float32) Mat4 {
	s, c := math.Sincos(turnsToRadians(turns))
	sin, cos := float32(s), float32(c)
	return Mat4{
		cos, 0, -sin, 0,
		0, 1, 0, 0,
		sin, 0, cos, 0,
		0, 0, 0, 1,
	}
}

// RotateLeftHandZ returns 4 by 4 matrix that, when multiplied with a
// homogeneous 4-element 3D vector, rotates the vector about the z-axis,
// applying the left-handed rule, by the given number of turns. 1 turn is 2*Pi.
func RotateLeftHandZ(turns float32) Mat4 {
	return RotateRightHandZ(-turns)
}

// RotateRightHandZ returns 4 by 4 matrix that, when multiplied with a
// homogeneous 4-element 3D vector, rotates the vector about the z-axis,
// applying the right-handed rule, by the given number of turns. 1 turn is
// 2*Pi.
func RotateRightHandZ(turns float32) Mat4 {
	s, c := math.Sincos(float64(turnsToRadians(turns)))
	sin, cos := float32(s), float32(c)
	return Mat4{
		cos, sin, 0, 0,
		-sin, cos, 0, 0,
		0, 0, 1, 0,
		0, 0, 0, 1,
	}
}

// RotateLeftHandAbout returns a 4 by 4 matrix that, when multiplied with a
// homogeneous 4-element 3D vector, rotates the vector about the given vector
// v, applying the left-handed rule, by the given number of turns. 1 turn is
// 2*Pi.
func RotateLeftHandAbout(v Vec3, turns float32) Mat4 {
	return RotateRightHandAbout(v, -turns)
}

// RotateRightHandAbout returns a 4 by 4 matrix that, when multiplied with a
// homogeneous 4-element 3D vector, rotates the vector about the given vector
// v, applying the right-handed rule, by the given number of turns. 1 turn is
// 2*Pi.
func RotateRightHandAbout(v Vec3, turns float32) Mat4 {
	sqLen := v.SquareNorm()
	if sqLen < 0.99999 || sqLen > 1.00001 {
		v = v.Normalized()
	}
	if sqLen == 0 {
		return Identity4()
	}
	s, c := math.Sincos(float64(turnsToRadians(turns)))
	sin, cos := float32(s), float32(c)
	x, y, z := v[0], v[1], v[2]
	return Mat4{
		cos + x*x*(1-cos), y*x*(1-cos) + z*sin, z*x*(1-cos) - y*sin, 0,
		x*y*(1-cos) - z*sin, cos + y*y*(1-cos), z*y*(1-cos) + x*sin, 0,
		x*z*(1-cos) + y*sin, y*z*(1-cos) - x*sin, cos + z*z*(1-cos), 0,
		0, 0, 0, 1,
	}
}

// Ortho returns an orthographic projection matrix.
func Ortho(left, right, bottom, top, near, far float32) Mat4 {
	return Mat4{
		2 / (right - left), 0, 0, (right + left) / (left - right),
		0, 2 / (top - bottom), 0, (top + bottom) / (bottom - top),
		0, 0, 2 / (far - near), (far + near) / (near - far),
		0, 0, 0, 1,
	}
}

// Perspective returns an perspective projection matrix.
func Perspective(fovRadians, aspect, near, far float32) Mat4 {
	f := 1 / float32(math.Tan(float64(fovRadians)/2))
	dz := far - near
	return Mat4{
		f / aspect, 0, 0, 0,
		0, f, 0, 0,
		0, 0, far / dz, -near * far / dz,
		0, 0, 1, 0,
	}
}

// LookAt returns a matrix that, when used for the camera, looks at target from
// position pos. Since you can tilt your head in infinite ways looking from one
// point at another, the up vector is used to specify which direction is up.
func LookAt(pos, target, up Vec3) Mat4 {
	z := target.Sub(pos).Normalized()
	x := up.Cross(z).Normalized()
	y := z.Cross(x)
	return Mat4{
		x[0], x[1], x[2], -x.Dot(pos),
		y[0], y[1], y[2], -y.Dot(pos),
		z[0], z[1], z[2], -z.Dot(pos),
		0, 0, 0, 1,
	}
}

func (m Mat4) String() string {
	return fmt.Sprintf(`%.2f %.2f %.2f %.2f
%.2f %.2f %.2f %.2f
%.2f %.2f %.2f %.2f
%.2f %.2f %.2f %.2f`, m[0], m[4], m[8], m[12], m[1], m[5], m[9], m[13], m[2],
		m[6], m[10], m[14], m[3], m[7], m[11], m[15])
}

// DecomposeAffineTransform decomposes the given matrix into scale, rotation and
// translation matrices that, when multiplied in that order, produce the
// original matrix. See this forum post for reference:
// https://math.stackexchange.com/questions/237369/given-this-transformation-matrix-how-do-i-decompose-it-into-translation-rotati
func DecomposeAffineTransform(m Mat4) (scale, rotation, translation Mat4) {
	translation = Mat4{
		1, 0, 0, m[12],
		0, 1, 0, m[13],
		0, 0, 1, m[14],
		0, 0, 0, 1,
	}
	sx := Vec3{m[0], m[4], m[8]}.Norm()
	sy := Vec3{m[1], m[5], m[9]}.Norm()
	sz := Vec3{m[2], m[6], m[10]}.Norm()
	scale = Mat4{
		sx, 0, 0, 0,
		0, sy, 0, 0,
		0, 0, sz, 0,
		0, 0, 0, 1,
	}
	fx, fy, fz := 1/sx, 1/sy, 1/sz
	rotation = Mat4{
		fx * m[0], fy * m[4], fz * m[8], 0,
		fx * m[1], fy * m[5], fz * m[9], 0,
		fx * m[2], fy * m[6], fz * m[10], 0,
		0, 0, 0, 1,
	}
	return
}
